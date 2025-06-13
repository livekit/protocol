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
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/livekit/psrpc"

	"github.com/livekit/protocol/logger"
)

const (
	defaultBucketLocation = "us-east-1"
)

type S3Config struct {
	AccessKey      string       `yaml:"access_key"`    // (env AWS_ACCESS_KEY_ID)
	Secret         string       `yaml:"secret"`        // (env AWS_SECRET_ACCESS_KEY)
	SessionToken   string       `yaml:"session_token"` // (env AWS_SESSION_TOKEN)
	Region         string       `yaml:"region"`        // (env AWS_DEFAULT_REGION)
	Endpoint       string       `yaml:"endpoint"`
	Bucket         string       `yaml:"bucket"`
	ForcePathStyle bool         `yaml:"force_path_style"`
	ProxyConfig    *ProxyConfig `yaml:"proxy_config"`

	MaxRetries    int           `yaml:"max_retries"`
	MaxRetryDelay time.Duration `yaml:"max_retry_delay"`
	MinRetryDelay time.Duration `yaml:"min_retry_delay"`

	Metadata             map[string]string `yaml:"metadata"`
	Tagging              string            `yaml:"tagging"`
	ContentDisposition   string            `yaml:"content_disposition"`
	GeneratePresignedUrl bool              `yaml:"generate_presigned_url"`
}

type S3Client struct {
	conf    *S3Config
	prefix  string
	awsConf *aws.Config
}

func NewS3Client(c *S3Config, prefix string) (*S3Client, error) {
	awsConf, err := getAWSConfig(c)
	if err != nil {
		return nil, err
	}

	return &S3Client{
		conf:    c,
		prefix:  prefix,
		awsConf: awsConf,
	}, nil
}

func getAWSConfig(conf *S3Config) (*aws.Config, error) {
	opts := func(o *awsConfig.LoadOptions) error {
		if conf.Region != "" {
			o.Region = conf.Region
		} else {
			o.Region = defaultBucketLocation
		}

		if conf.AccessKey != "" && conf.Secret != "" {
			o.Credentials = credentials.StaticCredentialsProvider{
				Value: aws.Credentials{
					AccessKeyID:     conf.AccessKey,
					SecretAccessKey: conf.Secret,
					SessionToken:    conf.SessionToken,
				},
			}
		}

		o.Retryer = func() aws.Retryer {
			return retry.NewStandard(func(o *retry.StandardOptions) {
				o.MaxAttempts = conf.MaxRetries
				o.MaxBackoff = conf.MaxRetryDelay
				o.Retryables = append(o.Retryables, &s3Retryer{})
			})
		}

		if conf.ProxyConfig != nil {
			proxyUrl, err := url.Parse(conf.ProxyConfig.Url)
			if err != nil {
				return err
			}
			s3Transport := http.DefaultTransport.(*http.Transport).Clone()
			s3Transport.Proxy = http.ProxyURL(proxyUrl)
			if conf.ProxyConfig.Username != "" && conf.ProxyConfig.Password != "" {
				auth := fmt.Sprintf("%s:%s", conf.ProxyConfig.Username, conf.ProxyConfig.Password)
				basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
				s3Transport.ProxyConnectHeader = http.Header{}
				s3Transport.ProxyConnectHeader.Add("Proxy-Authorization", basicAuth)
			}
			o.HTTPClient = &http.Client{Transport: s3Transport}
		}

		return nil
	}

	awsConf, err := awsConfig.LoadDefaultConfig(context.Background(), opts)
	if err != nil {
		return nil, psrpc.NewErrorf(psrpc.InvalidArgument, "failed to load aws config: %v", err)
	}

	if conf.Endpoint != "" {
		awsConf.BaseEndpoint = &conf.Endpoint
	} else if conf.Region == "" {
		if err = updateRegion(&awsConf, conf.Bucket); err != nil {
			return nil, err
		}
	}

	return &awsConf, nil
}

func updateRegion(awsConf *aws.Config, bucket string) error {
	req := &s3.GetBucketLocationInput{
		Bucket: &bucket,
	}

	resp, err := s3.NewFromConfig(*awsConf).GetBucketLocation(context.Background(), req)
	if err != nil {
		return psrpc.NewErrorf(psrpc.InvalidArgument, "failed to retrieve upload bucket region: %v", err)
	}

	if resp.LocationConstraint != "" {
		awsConf.Region = string(resp.LocationConstraint)
	}

	return nil
}

func (u *S3Client) UploadFile(
	localFilepath, storageFilepath string,
	contentType string,
) (string, int64, error) {

	storageFilepath = path.Join(u.prefix, storageFilepath)

	file, err := os.Open(localFilepath)
	if err != nil {
		return "", 0, ErrUploadFailed("S3", err)
	}
	defer func() {
		_ = file.Close()
	}()

	stat, err := file.Stat()
	if err != nil {
		return "", 0, ErrUploadFailed("S3", err)
	}

	l := logger.NewS3Logger()
	client := s3.NewFromConfig(*u.awsConf, func(o *s3.Options) {
		o.Logger = l
		o.ClientLogMode = aws.LogRequest | aws.LogResponse | aws.LogRetries
		o.UsePathStyle = u.conf.ForcePathStyle
	})

	input := &s3.PutObjectInput{
		Body:        file,
		Bucket:      aws.String(u.conf.Bucket),
		ContentType: aws.String(contentType),
		Key:         aws.String(storageFilepath),
		Metadata:    u.conf.Metadata,
	}
	if u.conf.Tagging != "" {
		input.Tagging = &u.conf.Tagging
	}
	if u.conf.ContentDisposition != "" {
		input.ContentDisposition = &u.conf.ContentDisposition
	} else {
		contentDisposition := "inline"
		input.ContentDisposition = &contentDisposition
	}

	if _, err = manager.NewUploader(client).Upload(context.Background(), input); err != nil {
		l.WriteLogs()
		return "", 0, ErrUploadFailed("S3", err)
	}

	endpoint := "s3.amazonaws.com"
	if u.conf.Endpoint != "" {
		endpoint = u.conf.Endpoint
	}

	var location string
	if u.conf.ForcePathStyle {
		location = fmt.Sprintf("https://%s/%s/%s", endpoint, u.conf.Bucket, storageFilepath)
	} else {
		location = fmt.Sprintf("https://%s.%s/%s", u.conf.Bucket, endpoint, storageFilepath)
	}

	if !u.conf.GeneratePresignedUrl {
		return location, stat.Size(), nil
	}

	res, err := s3.NewPresignClient(client).PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(u.conf.Bucket),
		Key:    aws.String(storageFilepath),
	})
	if err != nil {
		return "", 0, ErrUploadFailed("S3", err)
	}

	return res.URL, stat.Size(), nil
}

type s3Retryer struct{}

func (r *s3Retryer) IsErrorRetryable(_ error) aws.Ternary {
	return aws.TrueTernary
}
