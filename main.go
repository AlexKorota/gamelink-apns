package main

import (
	"context"
	"gamelink-apns/app"
	"gamelink-apns/config"
	"gamelink-apns/version"
	log "github.com/sirupsen/logrus"
)

func init() {
	config.LoadEnvironment()
	if config.IsDevelopmentEnv() {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

func main() {
	log.Printf(
		"Starting the service: commit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)
	ctx := context.Background()
	a := app.NewApp()
	a.ConnectNats()
	a.ConnectApns(ctx)
	a.GetAndPush()
}
