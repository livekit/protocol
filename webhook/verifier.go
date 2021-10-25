package webhook

import (
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"github.com/livekit/protocol/auth"
)

// Receive reads and verifies incoming webhook is signed with key/secret pair
// closes body after reading
func Receive(r *http.Request, provider auth.KeyProvider) ([]byte, error) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	authToken := r.Header.Get(authHeader)
	if authToken == "" {
		return nil, ErrNoAuthHeader
	}

	v, err := auth.ParseAPIToken(authToken)
	if err != nil {
		return nil, err
	}

	secret := provider.GetSecret(v.APIKey())
	if secret == "" {
		return nil, ErrSecretNotFound
	}

	claims, err := v.Verify(secret)
	if err != nil {
		return nil, err
	}

	// verify checksum
	sha := sha256.Sum256(data)
	hash := base64.StdEncoding.EncodeToString(sha[:])

	if claims.Sha256 != hash {
		return nil, ErrInvalidChecksum
	}

	return data, nil
}
