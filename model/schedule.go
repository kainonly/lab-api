package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Schedule struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	ClusterId  primitive.ObjectID `bson:"cluster_id" json:"cluster_id"`
	Name       string             `bson:"name" json:"name"`
	Image      string             `bson:"image" json:"image"`
	CreateTime time.Time          `bson:"create_time" json:"create_time"`
	UpdateTime time.Time          `bson:"update_time" json:"update_time"`
}
