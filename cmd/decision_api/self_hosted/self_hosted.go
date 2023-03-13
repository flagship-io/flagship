/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package self_hosted

import (
	"context"
	"crypto/tls"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/flagship-io/decision-api/pkg/connectors"
	"github.com/flagship-io/decision-api/pkg/connectors/assignments_managers"
	"github.com/flagship-io/decision-api/pkg/connectors/environment_loaders"
	"github.com/flagship-io/decision-api/pkg/connectors/hits_processors"
	"github.com/flagship-io/decision-api/pkg/models"
	"github.com/flagship-io/decision-api/pkg/server"
	"github.com/flagship-io/decision-api/pkg/utils/config"
	"github.com/flagship-io/decision-api/pkg/utils/logger"
	"github.com/spf13/cobra"
)

var shutdownTimeout = 3 * time.Second

func getAssignmentsManager(cfg *config.Config) (assignmentsManager connectors.AssignmentsManager, err error) {
	switch cfg.GetStringDefault("cache.type", "") {
	case "memory":
		assignmentsManager = assignments_managers.InitMemoryManager()
	case "local":
		assignmentsManager, err = assignments_managers.InitLocalCacheManager(assignments_managers.LocalOptions{
			DbPath: cfg.GetStringDefault("cache.options.dbpath", "cache_data"),
		})
	case "redis":
		var tlsConfig *tls.Config
		if cfg.GetBool("cache.options.redisTls") {
			tlsConfig = &tls.Config{}
		}
		assignmentsManager, err = assignments_managers.InitRedisManager(assignments_managers.RedisOptions{
			Host:      cfg.GetStringDefault("cache.options.redisHost", "localhost:6379"),
			Username:  cfg.GetStringDefault("cache.options.redisUsername", ""),
			Password:  cfg.GetStringDefault("cache.options.redisPassword", ""),
			Db:        cfg.GetIntDefault("cache.options.redisDb", 0),
			TTL:       cfg.GetDurationDefault("cache.options.redisTtl", 3*30*24*time.Hour),
			LogLevel:  cfg.GetStringDefault("log.level", config.LoggerLevel),
			LogFormat: logger.LogFormat(cfg.GetStringDefault("log.format", config.LoggerFormat)),
			TLSConfig: tlsConfig,
		})
	case "dynamo":
		session, _ := session.NewSession()
		client := dynamodb.New(session)
		assignmentsManager = assignments_managers.InitDynamoManager(assignments_managers.DynamoManagerOptions{
			Client:              client,
			TableName:           cfg.GetStringDefault("cache.options.dynamoTableName", "visitor-assignments"),
			PrimaryKeySeparator: cfg.GetStringDefault("cache.options.dynamoPKSeparator", "."),
			PrimaryKeyField:     cfg.GetStringDefault("cache.options.dynamoPKField", "id"),
			GetItemTimeout:      cfg.GetDurationDefault("cache.options.dynamoGetTimeout", 1*time.Second),
			LogLevel:            cfg.GetStringDefault("log.level", config.LoggerLevel),
			LogFormat:           logger.LogFormat(cfg.GetStringDefault("log.format", config.LoggerFormat)),
		})
	default:
		assignmentsManager = &assignments_managers.EmptyManager{}
	}

	return assignmentsManager, err
}

func createLogger(cfg *config.Config) *logger.Logger {
	lvl := cfg.GetStringDefault("log.level", config.LoggerLevel)
	format := cfg.GetStringDefault("log.format", config.LoggerFormat)

	return logger.New(lvl, logger.LogFormat(format), "Server")
}

func createServer(cfg *config.Config, log *logger.Logger) (*server.Server, error) {
	logLvl := cfg.GetStringDefault("log.level", config.LoggerLevel)
	logFmt := cfg.GetStringDefault("log.format", config.LoggerFormat)

	log.Info("initializing assignment cache manager from configuration")
	assignmentManager, err := getAssignmentsManager(cfg)
	if err != nil {
		log.Fatalf("error occurred when initializing assignment cache manager: %v", err)
	}

	return server.CreateServer(
		cfg.GetString("env_id"),
		cfg.GetString("api_key"),
		cfg.GetString("address"),
		server.WithLogger(log),
		server.WithEnvironmentLoader(
			environment_loaders.NewCDNLoader(
				environment_loaders.WithLogger(logLvl, logger.LogFormat(logFmt)),
				environment_loaders.WithPollingInterval(cfg.GetDuration("polling_interval"))),
		),
		server.WithHitsProcessor(hits_processors.NewDataCollectProcessor(hits_processors.WithLogger(logLvl, logger.LogFormat(logFmt)))),
		server.WithAssignmentsManager(assignmentManager),
		server.WithCorsOptions(&models.CorsOptions{
			Enabled:        cfg.GetBool("cors.enabled"),
			AllowedOrigins: cfg.GetStringDefault("cors.allowed_origins", config.ServerCorsAllowedOrigins),
			AllowedHeaders: cfg.GetStringDefault("cors.allowed_headers", config.ServerCorsAllowedHeaders),
		}),
	)
}

// DecisionSelfHostedCmd represents the self hosted decision api command
var DecisionSelfHostedCmd = &cobra.Command{
	Use:   "serve",
	Short: "use of decision api self-hosted",
	Long:  `use of the decison api self-hosted in the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		cfgFilename := flag.String("decision-config", "config.yaml", "Path the configuration file")
		flag.Parse()

		cfg, errCfg := config.NewFromFilename(*cfgFilename)
		logger := createLogger(cfg)

		if errCfg != nil {
			logger.Warn(errCfg)
		}

		srv, err := createServer(cfg, logger)
		if err != nil {
			logger.Fatalf("error when creating server: %v", err)
		}

		// Run server
		go func() {
			logger.Infof("Flagship Decision API server [%s] listening on %s", models.Version, cfg.GetStringDefault("address", ":8080"))
			if err := srv.Listen(); err != http.ErrServerClosed {
				logger.Fatalf("error when starting server: %v", err)
			}
		}()

		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)

		<-signalChannel

		// Try to gracefully shutdown the server
		ctx, cancelFunc := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancelFunc()
		srv.Shutdown(ctx)
	},
}
