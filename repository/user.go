package repository

import (
	"context"
	"fmt"
	"go-workspace/ext-crypt-backend/database"
	"go-workspace/ext-crypt-backend/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCol *mongo.Collection

func GetUsersColl() {
	userCol = database.DB.Collection("users")
}

func SignIn(user *models.User) (string, error) {
	log.Println("Call signin api")
	update := bson.M{
		"$set": bson.M{
			"remember_token": user.RememberToken,
			"updated_at":     time.Now(),
		},
		"$setOnInsert": bson.M{
			"name":       user.Name,
			"user_id":    user.UserId,
			"email":      user.Email,
			"role":       "USER",
			"created_at": time.Now(),
		},
	}

	// Find one document and update it or insert if not exists (upsert)
	opts := options.Update().SetUpsert(true)

	// Update the document or insert if it does not exist (upsert)
	updateResult, err := userCol.UpdateOne(context.TODO(), bson.M{"email": user.Email}, update, opts)
	if err != nil {
		return "", err
	}

	// Check if a new document was inserted or an existing document was updated
	if updateResult.UpsertedCount > 0 {
		fmt.Println("A new document was inserted with ID:", updateResult.UpsertedID)
	} else {
		fmt.Println("An existing document was updated.")
	}

	// Retrieve the updated or newly inserted document
	var updatedUser models.User
	err = userCol.FindOne(context.TODO(), bson.M{"email": user.Email}, options.FindOne().SetProjection(bson.M{"_id": 0})).Decode(&updatedUser)
	if err != nil {
		return "", err
	}

	// Print the updated or inserted document
	fmt.Printf("User: %+v\n", updatedUser)
	return updatedUser.UserId, nil
}
