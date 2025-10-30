package webhook

import (
	"crypto/tls"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

func configureRetryableHTTPClient(client *retryablehttp.Client, params HTTPClientParams) {
	if params.RetryWaitMin > 0 {
		client.RetryWaitMin = params.RetryWaitMin
	}
	if params.RetryWaitMax > 0 {
		client.RetryWaitMax = params.RetryWaitMax
	}
	if params.MaxRetries > 0 {
		client.RetryMax = params.MaxRetries
	}
	if params.ClientTimeout > 0 {
		ensureHTTPClient(client).Timeout = params.ClientTimeout
	}
	if params.InsecureSkipVerify {
		setInsecureSkipVerify(client)
	}
}

func ensureHTTPClient(client *retryablehttp.Client) *http.Client {
	if client.HTTPClient == nil {
		client.HTTPClient = &http.Client{}
	}
	return client.HTTPClient
}

func setInsecureSkipVerify(client *retryablehttp.Client) {
	httpClient := ensureHTTPClient(client)

	var transport *http.Transport
	switch t := httpClient.Transport.(type) {
	case *http.Transport:
		transport = t.Clone()
	case nil:
		transport = cloneDefaultTransport()
	default:
		transport = cloneDefaultTransport()
	}

	if transport.TLSClientConfig == nil {
		transport.TLSClientConfig = &tls.Config{}
	}
	transport.TLSClientConfig.InsecureSkipVerify = true // #nosec G402

	httpClient.Transport = transport
}

func cloneDefaultTransport() *http.Transport {
	if dt, ok := http.DefaultTransport.(*http.Transport); ok && dt != nil {
		return dt.Clone()
	}
	return &http.Transport{}
}
