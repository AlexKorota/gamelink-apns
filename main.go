package main

import (
	"context"
	"gamelink-apns/app"
	"gamelink-apns/config"
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
	ctx := context.Background()
	a := app.NewApp()
	a.ConnectNats()
	a.ConnectApns(ctx)
	a.GetAndPush()
}
