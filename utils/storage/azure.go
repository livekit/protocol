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
	"fmt"
	"net/url"
	"os"
	"path"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type AzureConfig struct {
	AccountName   string `yaml:"account_name"` // (env AZURE_STORAGE_ACCOUNT)
	AccountKey    string `yaml:"account_key"`  // (env AZURE_STORAGE_KEY)
	ContainerName string `yaml:"container_name"`
}

type AzureClient struct {
	conf      *AzureConfig
	prefix    string
	container string
}

func NewAzureClient(c *AzureConfig, prefix string) (*AzureClient, error) {
	return &AzureClient{
		conf:      c,
		prefix:    prefix,
		container: fmt.Sprintf("https://%s.blob.core.windows.net/%s", c.AccountName, c.ContainerName),
	}, nil
}

func (u *AzureClient) UploadFile(localFilepath, storageFilepath string, contentType string) (string, int64, error) {
	storageFilepath = path.Join(u.prefix, storageFilepath)

	credential, err := azblob.NewSharedKeyCredential(
		u.conf.AccountName,
		u.conf.AccountKey,
	)
	if err != nil {
		return "", 0, ErrUploadFailed("Azure", err)
	}

	azUrl, err := url.Parse(u.container)
	if err != nil {
		return "", 0, ErrUploadFailed("Azure", err)
	}

	pipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{
		Retry: azblob.RetryOptions{
			Policy:        azblob.RetryPolicyExponential,
			MaxTries:      maxRetries,
			RetryDelay:    minDelay,
			MaxRetryDelay: maxDelay,
		},
	})
	containerURL := azblob.NewContainerURL(*azUrl, pipeline)
	blobURL := containerURL.NewBlockBlobURL(storageFilepath)

	file, err := os.Open(localFilepath)
	if err != nil {
		return "", 0, ErrUploadFailed("Azure", err)
	}
	defer func() {
		_ = file.Close()
	}()

	stat, err := file.Stat()
	if err != nil {
		return "", 0, ErrUploadFailed("Azure", err)
	}

	// upload blocks in parallel for optimal performance
	// it calls PutBlock/PutBlockList for files larger than 256 MBs and PutBlob for smaller files
	_, err = azblob.UploadFileToBlockBlob(context.Background(), file, blobURL, azblob.UploadToBlockBlobOptions{
		BlobHTTPHeaders: azblob.BlobHTTPHeaders{ContentType: string(contentType)},
		BlockSize:       4 * 1024 * 1024,
		Parallelism:     16,
	})
	if err != nil {
		return "", 0, ErrUploadFailed("Azure", err)
	}

	return fmt.Sprintf("%s/%s", u.container, storageFilepath), stat.Size(), nil
}
