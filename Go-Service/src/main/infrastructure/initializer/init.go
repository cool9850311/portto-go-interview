package initializer

import (
	domainLogger "Go-Service/src/main/domain/interface/logger"
	"Go-Service/src/main/infrastructure/config"
	infraLogger "Go-Service/src/main/infrastructure/logger"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Client *mongo.Client
var DB *mongo.Database
var Log domainLogger.Logger

func InitLog() {
	var err error
	Log, err = infraLogger.NewLogger("application.log")
	if err != nil {
		panic(err)
	}
}
func InitConfig(log domainLogger.Logger) {
	err := config.LoadConfig(log)
	if err != nil {
		panic(err)
	}
}
func InitMongoClient() {

	// Create client options

	clientOptions := options.Client().ApplyURI(config.AppConfig.MongoDB.URI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Context with timeout to use for ping and initial connection check
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping the primary to verify connectivity
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")

	// Assign the client and database to global variables
	Client = client
	DB = client.Database(config.AppConfig.MongoDB.Database)
}
