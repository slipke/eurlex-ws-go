package webservice

import (
	"log"
	"net/http"
)

type Config struct {
	Client   *http.Client
	Username string
	Password string
	Logger   *log.Logger
	// @TODO fits here?
	PageSize int64
	// @TODO fits here?
	SearchLanguage string
}

func NewConfig(username, password string) *Config {
	return &Config{
		Username:       username,
		Password:       password,
		Logger:         log.Default(),
		Client:         http.DefaultClient,
		PageSize:       10,
		SearchLanguage: "en",
	}
}
