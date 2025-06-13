// Copyright 2024 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"cloud.google.com/go/storage"
	"github.com/googleapis/gax-go/v2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

const (
	storageScope = "https://www.googleapis.com/auth/devstorage.read_write"
)

type GCPConfig struct {
	CredentialsJSON string       `yaml:"credentials_json"` // (env GOOGLE_APPLICATION_CREDENTIALS)
	Bucket          string       `yaml:"bucket"`
	ProxyConfig     *ProxyConfig `yaml:"proxy_config"`
}

type GCPClient struct {
	conf   *GCPConfig
	prefix string
	client *storage.Client
}

func NewGCPClient(c *GCPConfig, prefix string) (*GCPClient, error) {
	u := &GCPClient{
		conf:   c,
		prefix: prefix,
	}

	var opts []option.ClientOption
	if c.CredentialsJSON != "" {
		jwtConfig, err := google.JWTConfigFromJSON([]byte(c.CredentialsJSON), storageScope)
		if err != nil {
			return nil, ErrUploadFailed("GCP", err)
		}
		opts = append(opts, option.WithTokenSource(jwtConfig.TokenSource(context.Background())))
	}

	defaultTransport := http.DefaultTransport.(*http.Transport)
	transportClone := defaultTransport.Clone()

	if c.ProxyConfig != nil {
		proxyUrl, err := url.Parse(c.ProxyConfig.Url)
		if err != nil {
			return nil, err
		}
		defaultTransport.Proxy = http.ProxyURL(proxyUrl)
		if c.ProxyConfig.Username != "" && c.ProxyConfig.Password != "" {
			auth := fmt.Sprintf("%s:%s", c.ProxyConfig.Username, c.ProxyConfig.Password)
			basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
			defaultTransport.ProxyConnectHeader = http.Header{}
			defaultTransport.ProxyConnectHeader.Add("Proxy-Authorization", basicAuth)
		}
	}

	client, err := storage.NewClient(context.Background(), opts...)
	// restore default transport
	http.DefaultTransport = transportClone
	if err != nil {
		return nil, ErrUploadFailed("GCP", err)
	}

	u.client = client
	return u, nil
}

func (u *GCPClient) UploadFile(localFilepath, storageFilepath string, _ string) (string, int64, error) {
	storageFilepath = path.Join(u.prefix, storageFilepath)

	file, err := os.Open(localFilepath)
	if err != nil {
		return "", 0, ErrUploadFailed("GCP", err)
	}
	defer func() {
		_ = file.Close()
	}()

	stat, err := file.Stat()
	if err != nil {
		return "", 0, ErrUploadFailed("GCP", err)
	}

	wc := u.client.Bucket(u.conf.Bucket).Object(storageFilepath).Retryer(
		storage.WithBackoff(gax.Backoff{
			Initial:    minDelay,
			Max:        maxDelay,
			Multiplier: 2,
		}),
		storage.WithMaxAttempts(maxRetries),
		storage.WithPolicy(storage.RetryAlways),
	).NewWriter(context.Background())
	wc.ChunkRetryDeadline = 0

	if _, err = io.Copy(wc, file); err != nil {
		return "", 0, ErrUploadFailed("GCP", err)
	}

	if err = wc.Close(); err != nil {
		return "", 0, ErrUploadFailed("GCP", err)
	}

	return fmt.Sprintf("https://%s.storage.googleapis.com/%s", u.conf.Bucket, storageFilepath), stat.Size(), nil
}
