package keycloak

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
	"learning-keycloak/internal/config"
	"learning-keycloak/internal/helper"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"` // modulus, base64url
	E   string `json:"e"` // exponent, base64url
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

func (j *JWK) ToRSAPublicKey() (*rsa.PublicKey, error) {

	if j.Kty != "RSA" {
		return nil, errors.New("Unsupported Key Type : " + j.Kty)
	}

	nBytes, errNBytes := base64.RawURLEncoding.DecodeString(j.N)

	if errNBytes != nil {
		return nil, errNBytes
	}

	eBytes, errEBytes := base64.RawURLEncoding.DecodeString(j.E)

	if errEBytes != nil {
		return nil, errEBytes
	}

	rsaPublicKey := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: int(new(big.Int).SetBytes(eBytes).Int64()),
	}

	return rsaPublicKey, nil

}

type TokenValidator struct {
	config    *config.KeycloakConfig
	endpoint  *Endpoint
	http      *http.Client
	my        sync.RWMutex
	keys      map[string]*rsa.PublicKey
	fetchedAt time.Time
}

func NewTokenValidator(cfg *config.KeycloakConfig, ep *Endpoint) *TokenValidator {
	return &TokenValidator{
		config:   cfg,
		endpoint: ep,
		http:     &http.Client{Timeout: 10 * time.Second},
		keys:     make(map[string]*rsa.PublicKey),
	}
}

func (v *TokenValidator) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("Unexpected Signing Method : %v", token.Header["alg"])
	}

	kid, ok := token.Header["kid"].(string)

	if !ok || kid == "" {
		return nil, errors.New("Token Header Has Not KID!")
	}

	return v.keys[kid], nil
}

func (v *TokenValidator) ValidateToken(ctx context.Context, raw string) (*KeycloakClaims, error) {
	claims := &KeycloakClaims{}

	_, err := jwt.ParseWithClaims(raw, claims, v.keyFunc,
		jwt.WithValidMethods([]string{"RS256"}),
		jwt.WithIssuer(v.endpoint.Issuer()),
		jwt.WithExpirationRequired(),
		jwt.WithLeeway(5*time.Second),
	)
	if err != nil {
		return nil, mapJWTError(err)
	}

	if err := claims.ValidateSemantics(v.config.KEYCLOAK_CLIENT_ID); err != nil {
		return nil, helper.NewUnauthorizedError("Invalid token claims",
			[]helper.ErrorDetail{{Message: err.Error()}})
	}
	return claims, nil
}

func mapJWTError(err error) error {
	switch {
	case errors.Is(err, jwt.ErrTokenExpired):
		return helper.NewUnauthorizedError("Token has expired",
			[]helper.ErrorDetail{{Message: "TOKEN_EXPIRED"}})
	case errors.Is(err, jwt.ErrTokenNotValidYet):
		return helper.NewUnauthorizedError("Token is not valid yet",
			[]helper.ErrorDetail{{Message: "TOKEN_NOT_YET_VALID"}})
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return helper.NewUnauthorizedError("Token signature is invalid",
			[]helper.ErrorDetail{{Message: "TOKEN_SIGNATURE_INVALID"}})
	case errors.Is(err, jwt.ErrTokenMalformed):
		return helper.NewUnauthorizedError("Token is malformed",
			[]helper.ErrorDetail{{Message: "TOKEN_MALFORMED"}})
	case errors.Is(err, jwt.ErrTokenInvalidIssuer):
		return helper.NewUnauthorizedError("Token issuer is not trusted",
			[]helper.ErrorDetail{{Message: "TOKEN_INVALID_ISSUER"}})
	default:
		return helper.NewUnauthorizedError("Token is invalid",
			[]helper.ErrorDetail{{Message: "TOKEN_INVALID"}})
	}
}
