package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/slipke/eurlex-ws-go/pkg/webservice"
)

func main() {
	cfg := webservice.NewConfig("", "")
	ws := webservice.NewWebservice(cfg)

	sr, err := ws.Search(
		webservice.NewSearchRequestWithConfig(
			cfg,
			"DN~32014R1338  OR  DN~32014R1337",
		),
	)
	if err != nil {
		log.Errorf("Failed to issue request: %s", err)
		return
	}

	log.Infof("%+v", sr)
}

func init() {
	log.SetLevel(log.DebugLevel)
}
