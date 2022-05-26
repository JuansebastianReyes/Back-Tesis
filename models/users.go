package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `json:"name"`
	LastName  string             `json:"last_name"`
	Email     string             `json:"email"`
	Telephone string             `json:"telephone"`
	Password  string             `json:"password"`
	Skills    []string           `json:"skills"`
	Projects  []string           `json:"projects"`
	CreateAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt  time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

type Users []*User
