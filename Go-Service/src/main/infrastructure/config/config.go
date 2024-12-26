package config

import (
	"Go-Service/src/main/domain/config"
	"Go-Service/src/main/domain/interface/logger"
	"Go-Service/src/main/infrastructure/util"
	"context"
	"github.com/joho/godotenv"
	"os"
)

var AppConfig config.Config
func LoadConfig(log logger.Logger) (error) {
	projectRootPath, err := util.GetProjectRootPath()
	if err != nil {
		log.Fatal(context.Background(), "Error getting project root path: "+err.Error())
		panic(err)
	}
	err = godotenv.Load(projectRootPath + "/.env")
	if err != nil {
		log.Fatal(context.Background(), "Error loading .env file: "+err.Error())
		panic(err)
	}
	AppConfig.MongoDB.URI = os.Getenv("MONGO_URI")
	AppConfig.MongoDB.Database = os.Getenv("MONGO_DATABASE")
	return nil
}
