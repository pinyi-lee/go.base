package mongo

import (
	"context"
	"fmt"

	Mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	instance *Manager
)

func GetInstance() *Manager {
	return instance
}

type Manager struct {
	Client  *Mongo.Client
	context context.Context
}

type Config struct {
	URI string
}

func (manager *Manager) Setup(config Config) error {
	background := context.Background()

	clientOptions := options.Client().ApplyURI(config.URI)
	client, err := Mongo.Connect(background, clientOptions)
	if err != nil {
		fmt.Printf("mongo connect fail, error : %+v\n", err)
		return err
	}

	err = client.Ping(background, nil)
	if err != nil {
		fmt.Printf("mongo ping fail, error : %+v\n", err)
		return err
	}

	instance = &Manager{
		Client:  client,
		context: background,
	}

	return nil
}

func (manager *Manager) Close() {
	manager.Client.Disconnect(manager.context)
}
