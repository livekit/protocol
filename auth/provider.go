package auth

import (
	"io"

	"gopkg.in/yaml.v3"
)

type FileBasedKeyProvider struct {
	keys map[string]string
}

func NewFileBasedKeyProviderFromReader(r io.Reader) (p *FileBasedKeyProvider, err error) {
	keys := make(map[string]string)
	decoder := yaml.NewDecoder(r)
	if err = decoder.Decode(&keys); err != nil {
		return
	}
	p = &FileBasedKeyProvider{
		keys: keys,
	}

	return
}

func NewFileBasedKeyProviderFromMap(keys map[string]string) *FileBasedKeyProvider {
	return &FileBasedKeyProvider{
		keys: keys,
	}
}

func (p *FileBasedKeyProvider) GetSecret(key string) string {
	return p.keys[key]
}

func (p *FileBasedKeyProvider) NumKeys() int {
	return len(p.keys)
}

type SimpleKeyProvider struct {
	apiKey    string
	apiSecret string
}

func NewSimpleKeyProvider(apiKey, apiSecret string) *SimpleKeyProvider {
	return &SimpleKeyProvider{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (p *SimpleKeyProvider) GetSecret(key string) string {
	if key == p.apiKey {
		return p.apiSecret
	}
	return ""
}

func (p *SimpleKeyProvider) NumKeys() int {
	return 1
}
