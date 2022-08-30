package models

// import (
// "go.mongodb.org/mongo-driver/bson"
// )

type Likes struct {
	TotalLikes int      `bson:"total_likes"`
	LikedBy    []string `bson:"liked_by"`
}

type Tweet struct {
	Author      string
	Tweet       string
	PublishedOn string `bson:"published_on"`
	Likes       Likes
}
