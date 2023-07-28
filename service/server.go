package service

import (
	"GIM/models"
	"GIM/pkg/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var log = logrus.New()

var (
	db            *gorm.DB
	redisClient   *redis.Client
	DatacenterAPI *DataCenterManager
)

func Init() {
	db = config.GetDB()
	redisClient = config.GetRedisClient()
	DatacenterAPI = NewDataCenterManager(&models.GlobalEnvironment)
}
