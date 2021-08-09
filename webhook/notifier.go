package webhook

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-logr/logr"
	"github.com/livekit/protocol/auth"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Notifier struct {
	apiKey    string
	apiSecret string
	urls      []string
	Logger    logr.Logger
}

func NewNotifier(apiKey, apiSecret string, urls []string) *Notifier {
	return &Notifier{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		urls:      urls,
		Logger:    logr.Discard(),
	}
}

func (n *Notifier) Notify(payload interface{}) error {
	var encoded []byte
	var err error
	if message, ok := payload.(proto.Message); ok {
		// use proto marshaler to ensure lowerCaseCamel
		encoded, err = protojson.Marshal(message)
	} else {
		// encode as JSON
		encoded, err = json.Marshal(payload)
	}
	if err != nil {
		return err
	}

	// sign payload
	sum := sha256.Sum256(encoded)
	b64 := base64.StdEncoding.EncodeToString(sum[:])

	at := auth.NewAccessToken(n.apiKey, n.apiSecret).
		SetValidFor(5 * time.Minute).
		SetSha256(b64)
	token, err := at.ToJWT()
	if err != nil {
		return err
	}

	for _, url := range n.urls {
		r, err := http.NewRequest("POST", url, bytes.NewReader(encoded))
		if err != nil {
			// ignore and continue
			n.Logger.Error(err, "could not create request", "url", url)
			continue
		}
		r.Header.Set(authHeader, token)
		_, err = http.DefaultClient.Do(r)
		if err != nil {
			n.Logger.Error(err, "could not post to webhook", "url", url)
		}
	}

	return nil
}
