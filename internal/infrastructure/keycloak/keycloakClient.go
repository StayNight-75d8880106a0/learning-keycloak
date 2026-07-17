package keycloak

import (
	"learning-keycloak/internal/config"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type KeycloakClient struct {
	config   *config.KeycloakConfig
	endpoint *Endpoint
	rds      *redis.Client
	http     *http.Client
}

func NewKeycloanClient(cfg *config.KeycloakConfig, ep *Endpoint, rds *redis.Client) *KeycloakClient {
	return &KeycloakClient{
		config:   cfg,
		endpoint: ep,
		rds:      rds,
		http:     &http.Client{Timeout: 10 * time.Second},
	}
}

type KeycloakTokenResponse struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	SessionState     string `json:"session_state"`
}

type KeycloakErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
