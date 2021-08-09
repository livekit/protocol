package auth

import (
	"time"

	"gopkg.in/square/go-jose.v2/jwt"
)

type APIKeyTokenVerifier struct {
	token    *jwt.JSONWebToken
	identity string
	apiKey   string
}

// ParseAPIToken parses an encoded JWT token and
func ParseAPIToken(raw string) (*APIKeyTokenVerifier, error) {
	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		return nil, err
	}

	out := jwt.Claims{}
	if err := tok.UnsafeClaimsWithoutVerification(&out); err != nil {
		return nil, err
	}

	v := &APIKeyTokenVerifier{
		token:    tok,
		apiKey:   out.Issuer,
		identity: out.Subject,
	}
	if v.identity == "" {
		v.identity = out.ID
	}
	return v, nil
}

// APIKey returns the API key this token was signed with
func (v *APIKeyTokenVerifier) APIKey() string {
	return v.apiKey
}

func (v *APIKeyTokenVerifier) Identity() string {
	return v.identity
}

func (v *APIKeyTokenVerifier) Verify(key interface{}) (*ClaimGrants, error) {
	if key == nil || key == "" {
		return nil, ErrKeysMissing
	}
	if s, ok := key.(string); ok {
		key = []byte(s)
	}
	out := jwt.Claims{}
	claims := ClaimGrants{}
	if err := v.token.Claims(key, &out, &claims); err != nil {
		return nil, err
	}
	if err := out.Validate(jwt.Expected{Issuer: v.apiKey, Time: time.Now()}); err != nil {
		return nil, err
	}

	// copy over identity
	claims.Identity = v.identity
	return &claims, nil
}
