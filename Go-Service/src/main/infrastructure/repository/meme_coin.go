package repository

import (
	"Go-Service/src/main/domain/entity"
	"Go-Service/src/main/domain/interface/logger"
	"Go-Service/src/main/domain/interface/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *MemeCoinRepository) Create(ctx context.Context, memeCoin *entity.MemeCoin) (string, error) {
	result, err := r.collection.InsertOne(ctx, memeCoin)
	if err != nil {
		r.logger.Error(ctx, "Error creating meme coin")
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *MemeCoinRepository) GetByID(ctx context.Context, id string) (*entity.MemeCoin, error) {
	var memeCoin entity.MemeCoin
	
	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Error(ctx, "Invalid ID format")
		return nil, err
	}
	
	filter := bson.D{{"_id", objectID}}
	err = r.collection.FindOne(ctx, filter).Decode(&memeCoin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Warn(ctx, "Meme coin not found")
			return nil, nil
		}
		r.logger.Error(ctx, "Error retrieving meme coin")
		return nil, err
	}
	return &memeCoin, nil
}

func (r *MemeCoinRepository) Update(ctx context.Context, description string, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Error(ctx, "Invalid ID format")
		return err
	}
	
	filter := bson.D{{"_id", objectID}}
	update := bson.D{{"$set", bson.D{{"description", description}}}}
	_, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		r.logger.Error(ctx, "Error updating meme coin")
		return err
	}
	return nil
}

func (r *MemeCoinRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Error(ctx, "Invalid ID format")
		return err
	}
	
	filter := bson.D{{"_id", objectID}}
	_, err = r.collection.DeleteOne(ctx, filter)
	if err != nil {
		r.logger.Error(ctx, "Error deleting meme coin")
		return err
	}
	return nil
}

func (r *MemeCoinRepository) Poke(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Error(ctx, "Invalid ID format")
		return err
	}
	
	filter := bson.D{{"_id", objectID}}
	update := bson.D{
		{"$inc", bson.D{{"popularity_score", 1}}},
	}
	_, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		r.logger.Error(ctx, "Error poking meme coin")
		return err
	}
	return nil
}
