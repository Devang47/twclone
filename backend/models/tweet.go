package models

type Likes struct {
	TotalLikes int      `json:"total_likes" bson:"total_likes"`
	LikedBy    []string `json:"liked_by" bson:"liked_by"`
}

type Tweet struct {
	Id          string `json:"id"`
	Author      string `json:"author"`
	AuthorUid   string `json:"author_uid" bson:"author_uid"`
	Tweet       string `json:"tweet"`
	PublishedOn string `json:"published_on" bson:"published_on"`
	Likes       Likes
}
