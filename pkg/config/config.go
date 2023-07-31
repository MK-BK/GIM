package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"GIM/models"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
)

func InitMysql() error {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", models.Conf.MysqlUser, models.Conf.MysqlPwd, models.Conf.MysqlHost, models.Conf.MysqlDB)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Error,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				Colorful:                  true,
			},
		),
		DisableAutomaticPing: true,
	})
	if err != nil {
		return err
	}

	return nil
}

func InitRedis() error {
	options := &redis.Options{
		Addr:     models.Conf.RedisHost,
		Password: models.Conf.RedisPwd,
		DB:       0,
	}

	Redis = redis.NewClient(options)
	if err := Redis.Ping().Err(); err != nil {
		return err
	}

	return nil
}

func LoadConfig(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, models.Conf)
}
