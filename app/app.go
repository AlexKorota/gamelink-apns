package app

import (
	"context"
	"encoding/json"
	"fmt"
	"gamelink-apns/config"
	push "gamelink-go/proto_nats_msg"
	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/go-nats"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
	log "github.com/sirupsen/logrus"
)

type App struct {
	nc    *nats.Conn
	apns  *apns2.Client
	mchan chan push.PushMsgStruct
}

func NewApp() App {
	log.Warn("starting...")
	mchan := make(chan push.PushMsgStruct)
	return App{mchan: mchan}
}

func (a *App) ConnectNats() {
	log.Warn("connecting nats")
	nc, err := nats.Connect(config.NatsDialAddress)
	if err != nil {
		log.Fatal(err)
	}
	a.nc = nc
}

func (a *App) ConnectApns(ctx context.Context) {
	log.Warn("config apns connection")
	authKey, err := token.AuthKeyFromFile(config.ServiceKeyPath)
	if err != nil {
		log.Fatal("token error:", err)
	}
	t := &token.Token{
		AuthKey: authKey,
		// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
		KeyID: config.KeyID,
		// TeamID from developer account (View Account -> Membership)
		TeamID: config.TeamID,
	}
	a.apns = apns2.NewTokenClient(t)
	a.apns.Production() // Выставить в продакшн, когда приложуха будет в апсторе
	log.Println(a.apns)
}

func (a *App) GetAndPush() {
	var msgStruct push.PushMsgStruct
	// Subscribe to updates
	_, err := a.nc.Subscribe(config.NatsApnsChan, func(m *nats.Msg) {
		err := proto.Unmarshal(m.Data, &msgStruct)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Warn("getting new message struct")
		log.Warn(msgStruct)
		a.mchan <- msgStruct
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		m := <-a.mchan
		a.prepareMsg(m)
	}
}

func (a *App) prepareMsg(msg push.PushMsgStruct) {
	log.Warn("prepare push")
	p := map[string]interface{}{
		"aps": map[string]string{
			"alert": msg.Message,
		},
	}
	m, err := json.Marshal(p)
	if err != nil {
		log.Warn(err)
	}
	notification := &apns2.Notification{
		DeviceToken: msg.UserInfo.DeviceID,
		Topic:       config.BundleID,
		Payload:     m,
	}
	res, err := a.apns.Push(notification)
	if err != nil {
		log.Warn(err)
		return
	}
	if res != nil {
		if res.Sent() {
			log.Println("Sent:", res.ApnsID)
		} else {
			log.Warn(fmt.Sprintf("Not Sent: %d %s %s", res.StatusCode, res.ApnsID, res.Reason))
		}
	}
}
