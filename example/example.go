package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/agonzalezro/logrus_honeybadger"
)

func init() {
	log.AddHook(honeybadger.NewHook("apikey", "live", 5*time.Second))
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}
