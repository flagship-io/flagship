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

func TestWriteOptionals(t *testing.T) {
	cfg, err := WriteOptionals("test_configuration", "client_credentials", "*", 86400)

	assert.NotNil(t, cfg)
	assert.Nil(t, err)

	assert.Equal(t, cfg.GetString("grant_type"), GrantType)
	assert.Equal(t, cfg.GetString("scope"), Scope)
	assert.Equal(t, cfg.GetInt("expiration"), Expiration)
}

func TestInitLocalConfigureConfig(t *testing.T) {
	cfg := InitLocalConfigureConfig("/home/blackbeard/.flagship/configurations/test_configuration.yaml")

	assert.NotNil(t, cfg)

	assert.Equal(t, cfg.ConfigFileUsed(), "/home/blackbeard/.flagship/configurations/test_configuration.yaml")

}

func TestWriteToken(t *testing.T) {
	cfg, err := WriteToken("test_configuration", "token")

	assert.NotNil(t, cfg)
	assert.Nil(t, err)

	assert.Equal(t, cfg.GetString("token"), Token)

}
