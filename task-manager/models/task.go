package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
    ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Title       string             `json:"title,omitempty" bson:"title,omitempty"`
    Description string             `json:"description,omitempty" bson:"description,omitempty"`
    Category    string             `json:"category,omitempty" bson:"category,omitempty"`
    Status      string             `json:"status,omitempty" bson:"status,omitempty"`
    UserID      primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
}
