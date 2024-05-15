package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	UserID     string             `json:"user_id" bson:"user_id"`
	OptionBit  string             `json:"option_bit" bson:"option_bit"`
	ExpireDate time.Time          `json:"expire_date" bson:"expire_date"`
}
