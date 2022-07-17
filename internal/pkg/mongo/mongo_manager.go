package mongo

import (
	"context"
	"errors"
	"fmt"

	Mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

var (
	instance *Manager
)

func GetInstance() *Manager {
	return instance
}

type Manager struct {
	context      context.Context
	client       *Mongo.Client
	databaseName string
}

type Config struct {
	URI string
}

var ERROR_DATA_NOT_FOUND = errors.New("data not found")

func (manager *Manager) Setup(config Config) error {
	background := context.Background()

	connstring, err := connstring.ParseAndValidate(config.URI)
	if err != nil {
		fmt.Printf("mongo connstring parse and validate fail, error : %+v\n", err)
		return err
	}

	databaseName := connstring.Database

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
		context:      background,
		client:       client,
		databaseName: databaseName,
	}

	return nil
}

func (manager *Manager) Close() {
	manager.client.Disconnect(manager.context)
}
