package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// MemeCoin represents a meme coin entity
type MemeCoin struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name" validate:"required,unique"`
	Description     string             `bson:"description"`
	CreatedAt       time.Time          `bson:"created_at"`
	PopularityScore int                `bson:"popularity_score"`
}
