package storage

import (
	"time"

	"github.com/livekit/psrpc"
)

const (
	maxRetries = 5
	minDelay   = time.Millisecond * 100
	maxDelay   = time.Second * 5
)

type StorageClient interface {
	UploadFile(localFilepath string, storageFilepath string, contentType string) (string, int64, error)
}

type ProxyConfig struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func ErrUploadFailed(location string, err error) error {
	return psrpc.NewErrorf(psrpc.InvalidArgument, "%s upload failed: %v", location, err)
}
