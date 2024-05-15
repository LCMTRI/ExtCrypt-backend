package repository

import (
	"context"
	"go-workspace/ext-crypt-backend/database"
	"go-workspace/ext-crypt-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()
var coll *mongo.Collection

func GetSubscriptionsColl() {
	coll = database.DB.Collection("subscriptions")
}

func GetAllSubscriptions() ([]*models.Subscription, error) {
	var subscriptions []*models.Subscription
	c, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	err = c.All(ctx, &subscriptions)
	if err != nil {
		return nil, err
	}

	return subscriptions, err
}
