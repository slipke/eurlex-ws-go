package webservice

import "net/http"

type Config struct {
	Client   *http.Client
	Username string
	Password string
}

func NewConfig(username, password string) *Config {
	return &Config{
		Username: username,
		Password: password,
	}
}
