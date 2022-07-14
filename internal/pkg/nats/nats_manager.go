package nats

import (
	"encoding/json"
	"fmt"

	Nats "github.com/nats-io/nats.go"
)

var (
	instance *Manager
)

func GetInstance() *Manager {
	return instance
}

type Manager struct {
	connect *Nats.Conn
}

type Config struct {
	Url string
}

func (manager *Manager) Setup(config Config) error {

	nc, err := Nats.Connect(config.Url)
	if err != nil {
		fmt.Printf("nats connect fail, error : %+v\n", err)
		return err
	}

	instance = &Manager{
		connect: nc,
	}

	return nil
}

func (manager *Manager) Publish(sub string, msg interface{}) error {

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = manager.connect.Publish(sub, data)
	if err != nil {
		return err
	}

	return nil
}

func (manager *Manager) Subscribe(sub string, msgHandler func(sub string, msg []byte)) (*Nats.Subscription, error) {

	subscription, err := manager.connect.Subscribe(sub, func(msg *Nats.Msg) {
		msgHandler(msg.Subject, msg.Data)
	})
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
