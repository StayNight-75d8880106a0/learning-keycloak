package keycloak

import (
	"fmt"
	"learning-keycloak/internal/config"
	"strings"
)

type Endpoint struct {
	issuer string
	oidc   string
	admin  string
}

func NewEndpoint(cfg *config.KeycloakConfig) *Endpoint {
	public := strings.TrimSuffix(cfg.KEYCLOAK_URL, "/")
	internal := strings.TrimSuffix(cfg.KEYCLOAK_INTERNAL_URL, "/")

	return &Endpoint{
		issuer: fmt.Sprintf("%s/realms/%s", public, cfg.KEYCLOAK_REALM),
		oidc:   fmt.Sprintf("%s/realms/%s/protocol/openid-connect", internal, cfg.KEYCLOAK_REALM),
		admin:  fmt.Sprintf("%s/admin/realms/%s", internal, cfg.KEYCLOAK_REALM),
	}
}

func (e *Endpoint) Issuer() string {
	return e.issuer
}

func (e *Endpoint) Token() string {
	return e.issuer + "/token"
}

func (e *Endpoint) Logout() string {
	return e.oidc + "/logout"
}

func (e *Endpoint) Certs() string {
	return e.oidc + "/certs"
}

func (e *Endpoint) Users() string {
	return e.admin + "/users"
}

func (e *Endpoint) UserActionEmail(userID string) string {
	return fmt.Sprintf("%s/users/%s/execute-actions-email", e.admin, userID)
}
