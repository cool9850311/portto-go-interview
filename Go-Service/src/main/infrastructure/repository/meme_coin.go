package repository

import (
	"Go-Service/src/main/domain/entity"
	"Go-Service/src/main/domain/interface/logger"
	"Go-Service/src/main/domain/interface/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MemeCoinRepository struct {
	collection *mongo.Collection
	logger     logger.Logger
}

func NewMemeCoinRepository(db *mongo.Database, logger logger.Logger) repository.MemeCoinRepository {
	collection := db.Collection("meme_coin")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"name", 1}}, // 1 for ascending order
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		logger.Fatal(context.TODO(), "Error creating index")
		panic(err)
	}
	return &MemeCoinRepository{collection: collection, logger: logger}
}

func (r *MemeCoinRepository) Create(ctx context.Context, memeCoin *entity.MemeCoin) error {
	_, err := r.collection.InsertOne(ctx, memeCoin)
	if err != nil {
		r.logger.Error(ctx, "Error creating meme coin")
		return err
	}
	return nil
}
