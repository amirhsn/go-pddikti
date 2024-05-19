package gopddikti

import (
	"errors"
	"net/http"
	"time"
)

const (
	baseAPIURL = "https://api-frontend.kemdikbud.go.id"
)

type ClientConfig struct {
	baseURL    string
	httpClient *http.Client

	// ContextTimeout in ms
	ContextTimeout int
	UseCorsPolicy  bool
	Origin         string
	Referer        string
}

func DefaultConfig() ClientConfig {
	return ClientConfig{
		baseURL:    baseAPIURL,
		httpClient: &http.Client{},
	}
}

func CustomConfig(cfg *ClientConfig) (ClientConfig, error) {
	if cfg.UseCorsPolicy {
		if cfg.Origin == "" || cfg.Referer == "" {
			return ClientConfig{}, errors.New("invalid cors schema")
		}
	}
	return ClientConfig{
		baseURL: baseAPIURL,
		httpClient: &http.Client{
			Timeout: time.Duration(cfg.ContextTimeout),
		},
		UseCorsPolicy: cfg.UseCorsPolicy,
		Origin:        cfg.Origin,
		Referer:       cfg.Referer,
	}, nil
}
