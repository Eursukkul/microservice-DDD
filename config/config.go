package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App      App
		Db       Db
		Jwt      Jwt
		Grpc     Grpc
		Paginate Paginate
		Rabbit   Rabbit
		Redis    Redis
	}

	App struct {
		Name  string
		Url   string
		Stage string
	}

	Db struct {
		Url string
	}

	Jwt struct {
		AccessSecretKey  string
		RefreshSecretKey string
		ApiSceretKey     string
		AccessDuration   int64
		RefreshDuration  int64
		ApiDuration      int64
	}

	Grpc struct {
		AuthUrl          string
		UsersUrl         string
		KnowledgeUrl     string
		ElasticSearchUrl string
		RoleUrl          string
	}

	Paginate struct {
		ItemNextPageBasedUrl      string
		InventoryNextPageBasedUrl string
	}

	Rabbit struct {
		Url string
	}

	Redis struct {
		Url string
	}
)

func LoadConfig(path string) Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		App: App{
			Name:  os.Getenv("APP_NAME"),
			Url:   os.Getenv("APP_URL"),
			Stage: os.Getenv("APP_STAGE"),
		},
		Db: Db{
			Url: os.Getenv("DB_URL"),
		},
		Jwt: Jwt{
			AccessSecretKey:  os.Getenv("JWT_ACCESS_SECRET_KEY"),
			RefreshSecretKey: os.Getenv("JWT_REFRESH_SECRET_KEY"),
			ApiSceretKey:     os.Getenv("JWT_API_SECRET_KEY"),
			AccessDuration: func() int64 {
				result, err := strconv.ParseInt(os.Getenv("JWT_ACCESS_DURATION"), 10, 64)
				if err != nil {
					log.Fatal("Error loading access duration failed")
				}
				return result
			}(),
			RefreshDuration: func() int64 {
				result, err := strconv.ParseInt(os.Getenv("JWT_REFRESH_DURATION"), 10, 64)
				if err != nil {
					log.Fatal("Error loading refresh duration failed")
				}
				return result
			}(),
		},
		Paginate: Paginate{
			ItemNextPageBasedUrl:      os.Getenv("PAGINATE_ITEM_NEXT_PAGE_BASED_URL"),
			InventoryNextPageBasedUrl: os.Getenv("PAGINATE_INVENTORY_NEXT_PAGE_BASED_URL"),
		},
		Rabbit: Rabbit{
			Url: os.Getenv("RABBITMQ_URL"),
		},
		Redis: Redis{
			Url: os.Getenv("REDIS_URL"),
		},
	}
}
