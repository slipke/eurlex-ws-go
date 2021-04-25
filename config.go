package eurlex

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Username string
	Password string
	Logger   *log.Logger
	Client   *http.Client
}

func NewConfig(username, password string) *Config {
	return &Config{
		Username: username,
		Password: password,
		Logger:   log.StandardLogger(),
		Client:   http.DefaultClient,
	}
}
