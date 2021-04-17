package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/slipke/eurlex-ws-go/pkg/webservice"
)

func main() {
	cfg := webservice.NewConfig("testuser", "testpass")
	ws := webservice.NewWebservice(cfg)

	sr, err := ws.Search(
		webservice.NewSearchRequestWithConfig(
			cfg,
			"DN=3*",
		),
	)
	if err != nil {
		log.Errorf("Failed to issue request: %s", err)
	}

	log.Info(sr)
}

func init() {
	log.SetLevel(log.DebugLevel)
}
