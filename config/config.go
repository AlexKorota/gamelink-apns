package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)

var (
	//NatsDialAddress - dial address for nats
	NatsDialAddress string
	//NatsApnsChan - nats chan for apns push
	NatsApnsChan string
	//ServiceKeyPAth - path to json service google key
	ServiceKeyPath string
	//KeyID - from developer account (Certificates, Identifiers & Profiles -> Keys)
	KeyID string
	//TeamID -  from developer account (View Account -> Membership)
	TeamID string
	//BundleID - app address? com.smedialink.testsigningapp
	BundleID string
)

const (
	modeKey         = "MODE"
	devMode         = "development"
	natsDial        = "NATSDIAL"
	natsApnsChannel = "NATSCHANAPNS"
	servicekeypath  = "SKEYPATH"
	keyID           = "KEYID"
	teamID          = "TEAMID"
	bundleID        = "TOPIC"
)

func init() {
	LoadEnvironment()
}

//GetEnvironment - this function returns mode string of the os environment or "development" mode if empty or not defined
func GetEnvironment() string {
	var env string
	if env = os.Getenv(modeKey); env == "" {
		return devMode
	}
	return env
}

//IsDevelopmentEnv - this function try to get mode environment and check it is development
func IsDevelopmentEnv() bool { return GetEnvironment() == devMode }

func LoadEnvironment() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = godotenv.Load(path.Join(wd, strings.ToLower(GetEnvironment())+".env"))
	if err != nil {
		log.Warning(err.Error())
	}
	NatsDialAddress = os.Getenv(natsDial)
	if NatsDialAddress == "" {
		log.Fatal("nats dial address must be set")
	}
	NatsApnsChan = os.Getenv(natsApnsChannel)
	if NatsApnsChan == "" {
		log.Fatal("nats android chan must be set")
	}
	ServiceKeyPath = os.Getenv(servicekeypath)
	if ServiceKeyPath == "" {
		log.Fatal("path to service key must be set")
	}
	KeyID = os.Getenv(keyID)
	if KeyID == "" {
		log.Fatal("keyID must be set")
	}
	TeamID = os.Getenv(teamID)
	if TeamID == "" {
		log.Fatal("keyID must be set")
	}
	BundleID = os.Getenv(bundleID)
	if TeamID == "" {
		log.Fatal("topic must be set")
	}
}
