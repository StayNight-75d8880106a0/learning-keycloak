package config

import "os"

type MySqlConfig struct {
	MYSQL_HOST      string
	MYSQL_USER      string
	MYSQL_PORT      string
	MYSQL_CHARSET   string
	MYSQL_COLLATION string
	MYSQL_DATABASE  string
	MYSQL_PASSWORD  string
}

func NewMySqlConfig() *MySqlConfig {
	return &MySqlConfig{
		MYSQL_HOST:      os.Getenv("MYSQL_HOST"),
		MYSQL_USER:      os.Getenv("MYSQL_USER"),
		MYSQL_PORT:      os.Getenv("MYSQL_PORT"),
		MYSQL_CHARSET:   os.Getenv("MYSQL_CHARSET"),
		MYSQL_COLLATION: os.Getenv("MYSQL_COLLATION"),
		MYSQL_DATABASE:  os.Getenv("MYSQL_DATABASE"),
		MYSQL_PASSWORD:  os.Getenv("MYSQL_PASSWORD"),
	}
}

func NewAppConfig() string {
	port := os.Getenv("APP_PORT")

	return port
}

type RedisConfig struct {
	REDIS_HOST      string
	REDIS_PORT      string
	REDIS_PASSWORD  string
	REDIS_DATABASES int
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		REDIS_HOST:      os.Getenv("REDIS_HOST"),
		REDIS_PORT:      os.Getenv("REDIS_PORT"),
		REDIS_PASSWORD:  os.Getenv("REDIS_PASSWORD"),
		REDIS_DATABASES: 0,
	}
}

type KeycloakConfig struct {
	KEYCLOAK_URL           string
	KEYCLOAK_CLIENT_ID     string
	KEYCLOAK_CLIENT_SECRET string
}

func NewKeycloakConfig() *KeycloakConfig {
	return &KeycloakConfig{
		KEYCLOAK_URL:           os.Getenv("KEYCLOAK_URL"),
		KEYCLOAK_CLIENT_ID:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		KEYCLOAK_CLIENT_SECRET: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
	}
}

func LoadConfig() {
	NewMySqlConfig()
	NewRedisConfig()
	NewAppConfig()
	NewKeycloakConfig()
}
