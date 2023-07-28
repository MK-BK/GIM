package main

import (
	"flag"
	"log"
	"os"

	"GIM/api"
	"GIM/models"
	"GIM/pkg/config"
	"GIM/service"
	"GIM/service/user"
)

var env = &models.GlobalEnvironment

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.json", "config.json path")
	flag.Parse()

	if err := config.LoadConfig(configPath); err != nil {
		panic(err)
	}

	if err := config.InitMysql(); err != nil {
		panic(err)
	}

	if err := config.InitRedis(); err != nil {
		panic(err)
	}

	initDataDir()

	initDatabase()

	initEnvironment()

	if err := api.InitRouter(); err != nil {
		log.Fatal(err)
	}
}

func initDataDir() {
	_, err := os.Open(models.Conf.DataPath)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	if err != nil && os.IsNotExist(err) {
		if err := os.Mkdir(models.Conf.DataPath, 0777); err != nil {
			panic(err)
		}
	}
}

func initDatabase() {
	db := config.GetDB()
	db.AutoMigrate(&models.RoleAssignment{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Group{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.Request{})
	db.AutoMigrate(&models.Moment{})
	db.AutoMigrate(&models.Comment{})
}

func initEnvironment() {
	env.Users = user.NewUserManager()
	env.Groups = service.NewGroupManager()
	env.Comments = service.NewCommonManager()
	env.Requests = service.NewRequestManager()
	env.Moments = service.NewMomentManager()
	env.Sessions = service.NewSessionManager()
	env.RoleAssignments = service.NewRoleAssignmentManager()

	service.Init()
}
