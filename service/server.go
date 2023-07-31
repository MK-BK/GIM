package service

import (
	"GIM/models"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

var DatacenterAPI *DataCenterManager

func Init() {
	DatacenterAPI = NewDataCenterManager(&models.GlobalEnvironment)
}
