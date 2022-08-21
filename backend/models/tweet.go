package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Tweet struct {
	Author      string
	Data        string
	PublishedOn string `bson:"published_on"`
}
