package webservice

import "net/http"

type Config struct {
	Client   *http.Client
	Username string
	Password string
	// @TODO fits here?
	PageSize int64
	// @TODO fits here?
	SearchLanguage string
}

func NewConfig(username, password string) *Config {
	return &Config{
		Username:       username,
		Password:       password,
		Client:         http.DefaultClient,
		PageSize:       10,
		SearchLanguage: "en",
	}
}
