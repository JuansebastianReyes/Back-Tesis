package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `json:"name"`
	LeaderId    string             `json:"leaderid"`
	Skills      []string           `json:"skills"`
	MembersId   []string           `json:"membersid"`
	Description string             `json:"description"`
	Phases      []string           `json:"phases"`
	State       string             `json:"state"`
	CreateAt    time.Time          `bson:"created_at" json:"created_at"`
	StartDate   time.Time          `bson:"start_date" json:"start_date"`
	EndDate     time.Time          `bson:"end_date" json:"end_date"`
	UpdateAt    time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

type Projects []*Project
