package user_repository

import (
	"context"
	"main/database"
	m "main/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

func Create(user m.User) error {

	var err error

	_, err = collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func Read() (m.Users, error) {

	var users m.Users
	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var user m.User
		err := cur.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}
	return users, nil
}

func Update(user m.User, userId string) error {

	var err error
	uid, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": uid}

	Update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"telephone":  user.Telephone,
			"last_name":  user.LastName,
			"password":   user.Password,
			"skills":     user.Skills,
			"projects":   user.Projects,
			"updated_at": time.Now(),
		}}

	_, err = collection.UpdateOne(ctx, filter, Update)
	if err != nil {
		return err
	}

	return nil
}

func Delete(userId string) error {

	var err error
	var uid primitive.ObjectID

	uid, err = primitive.ObjectIDFromHex(userId)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": uid}

	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}

func ReadById(userId string) (*m.User, error) {

	var user m.User
	var err error

	uid, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": uid}

	err = collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func ReadByEmail(email string) (*m.User, error) {

	var user m.User
	var err error

	filter := bson.M{"email": email}

	err = collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
