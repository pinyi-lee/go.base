package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pinyi-lee/go.base.git/internal/app/router"
	"github.com/pinyi-lee/go.base.git/internal/pkg/config"
	"github.com/pinyi-lee/go.base.git/internal/pkg/http/client"
	"github.com/pinyi-lee/go.base.git/internal/pkg/logger"
)

func Setup() {
	var err error

	if err = config.Setup(); err != nil {
		log.Fatal(err)
	}

	/*
		if err = elasticsearch.GetInstance().Setup(elasticsearch.Config{
			Url:         config.Env.ElasticsearchUrl,
			IndexPrefix: config.Env.ElasticsearchIndexPrefix,
		}); err != nil {
			log.Fatalf("elasticsearch Setup, error:%v", err)
		}
	*/

	if err = client.Setup(); err != nil {
		log.Fatal(err)
	}

	if err = logger.Setup(config.Env.LogLevel); err != nil {
		log.Fatal(err)
	}

	/*
		if err = mongo.GetInstance().Setup(mongo.Config{
			URI: config.Env.MongoURI,
		}); err != nil {
			log.Fatalf("mongo Setup, error:%v", err)
		}
	*/

	/*
		if err = nats.GetInstance().Setup(nats.Config{
			Url: config.Env.NatsUrl,
		}); err != nil {
			log.Fatalf("nats Setup, error:%v", err)
		}
	*/

	/*
		if err = postgres.GetInstance().Setup(postgres.Config{
			Username:                config.Env.PostgresUsername,
			Password:                config.Env.PostgresPassword,
			Host:                    config.Env.PostgresHost,
			Port:                    config.Env.PostgresPort,
			TableName:               config.Env.PostgresName,
			MinConnSize:             config.Env.PostgresMinConnSize,
			MaxConnSize:             config.Env.PostgresMaxConnSize,
			MaxConnIdleTimeBySecond: time.Duration(config.Env.PostgresMaxConnIdleTimeBySecond),
			MaxConnLifetimeBySecond: time.Duration(config.Env.PostgresMaxConnLifeTimeBySecond),
		}); err != nil {
			log.Fatalf("postgres Setup, error:%v", err)
		}
	*/

	/*
		if err = redis.GetInstance().Setup(redis.Config{
			Type:         config.Env.RedisType,
			EndpointList: config.Env.RedisEndpointList,
			Password:     config.Env.RedisPassword,
		}); err != nil {
			log.Fatalf("redis Setup, error:%v", err)
		}
	*/

	if err = router.Setup(); err != nil {
		log.Fatal(err)
	}
}

func Close() {
}

func RunServer() {
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Env.Port),
		Handler:      router.Router,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("%s\n", err)
	}
}

// @title        Community Service Swagger
// @description  this service is Community Service
func main() {
	Setup()
	defer Close()
	RunServer()
}
