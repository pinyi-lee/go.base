package test

import (
	"log"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/pinyi-lee/go.base.git/internal/app/router"
	"github.com/pinyi-lee/go.base.git/internal/pkg/config"
	"github.com/pinyi-lee/go.base.git/internal/pkg/http/client"
	"github.com/pinyi-lee/go.base.git/internal/pkg/logger"
)

func TestMain(m *testing.M) {
	//ctx := context.Background()

	/*
		mongoContainer, setupMongoErr := container.SetupMongo(ctx)

		if setupMongoErr != nil {
			log.Fatalf("SetupMongo Fail, %s\n", setupMongoErr)
		}
		defer mongoContainer.Terminate(ctx)
	*/

	/*
		postgresContainer, setupPostgresErr := container.SetupPostgres(ctx)
		if setupPostgresErr != nil {
			log.Fatalf("SetupPostgres Fail, %s\n", setupPostgresErr)
		}
		defer postgresContainer.Terminate(ctx)
	*/

	/*
		redisContainer, setupRedisErr := container.SetupRedis(ctx)
		if setupRedisErr != nil {
			log.Fatalf("SetupRedis Fail, %s\n", setupRedisErr)
		}
		defer redisContainer.Terminate(ctx)
	*/

	Setup()
	defer Close()

	httpmock.ActivateNonDefault(client.Get().GetClient())

	r := m.Run()
	os.Exit(r)
}

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
