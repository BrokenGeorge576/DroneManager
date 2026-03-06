package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Drone struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Model     string             `bson:"model"`
	Serial    string             `bson:"serial"`
	Charge    int                `bson:"charge"`
	Status    string             `bson:"status"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
