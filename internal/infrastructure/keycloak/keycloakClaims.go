package keycloak

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type RealmAccess struct {
	Roles []string `json:"roles"`
}

type ResourceAccess struct {
	Roles []string `json:"roles"`
}

type KeycloakClaims struct {
	jwt.RegisteredClaims
	Typ              string                    `json:"typ"`
	Azp              string                    `json:"azp"`
	Sid              string                    `json:"sid"`
	SessionState     string                    `json:"session_state"`
	PreferedUsername string                    `json:"preferred_username"`
	Email            string                    `json:"email"`
	EmailVerified    bool                      `json:"email_verified"`
	GivenName        string                    `json:"given_name"`
	FamilyName       string                    `json:"family_name"`
	Scope            string                    `json:"scope"`
	RealmAccess      RealmAccess               `json:"realm_access"`
	ResourceAccess   map[string]ResourceAccess `json:"resource_access"`
}

var (
	ErrMissingSubject   = errors.New("missing subject in claims")
	ErrMissingJTI       = errors.New("missing jti in claims")
	ErrInvalidTokenType = errors.New("invalid token type in claims")
	ErrInvalidAzp       = errors.New("invalid azp in claims")
	ErrMissingUsername  = errors.New("missing preferred_username in claims")
)

func (c *KeycloakClaims) ValidateSemantics(expectedClientID string) error {
	if strings.TrimSpace(c.Subject) == "" {
		return ErrMissingSubject
	}

	if strings.TrimSpace(c.ID) == "" {
		return ErrMissingJTI
	}

	if c.Typ != "Bearer" {
		return ErrInvalidTokenType
	}

	if c.Azp != expectedClientID {
		return ErrInvalidAzp
	}

	if strings.TrimSpace(c.PreferedUsername) == "" {
		return ErrMissingUsername
	}

	return nil

}
