package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetOptionalsDefault(t *testing.T) {
	cfg, err := SetOptionalsDefault("client_credentials", "*", 86400)

	assert.NotNil(t, cfg)
	assert.Nil(t, err)

	assert.Equal(t, cfg.GetString("grant_type"), GrantType)
	assert.Equal(t, cfg.GetString("scope"), Scope)
	assert.Equal(t, cfg.GetInt("expiration"), Expiration)
}

func TestWriteCredentials(t *testing.T) {
	cfg, err := WriteCredentials("credentialsTest.yaml", "clientID", "clientSecret", "accountID", "accountEnvironmentID")

	assert.NotNil(t, cfg)
	assert.Nil(t, err)

	assert.Equal(t, cfg.GetString("client_id"), ClientID)
	assert.Equal(t, cfg.GetString("client_secret"), ClientSecret)
	assert.Equal(t, cfg.GetString("account_id"), AccountID)
	assert.Equal(t, cfg.GetString("account_environment_id"), AccountEnvironmentID)

}

func TestWriteOptionals(t *testing.T) {
	cfg, err := WriteOptionals("credentialsTest.yaml", "client_credentials", "*", 86400)

	assert.NotNil(t, cfg)
	assert.Nil(t, err)

	assert.Equal(t, cfg.GetString("grant_type"), GrantType)
	assert.Equal(t, cfg.GetString("scope"), Scope)
	assert.Equal(t, cfg.GetInt("expiration"), Expiration)
}

func TestInitLocalConfigureConfig(t *testing.T) {
	cfg := InitLocalConfigureConfig("/home/blackbeard/.flagship/credentialsTest.yaml")

	assert.NotNil(t, cfg)

	assert.Equal(t, cfg.ConfigFileUsed(), "/home/blackbeard/.flagship/credentialsTest.yaml")

}

func TestWriteToken(t *testing.T) {
	cfg, err := WriteToken("credentialsTest.yaml", "token")

	assert.NotNil(t, cfg)
	assert.Nil(t, err)

	assert.Equal(t, cfg.GetString("token"), Token)

}
