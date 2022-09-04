package models

type Likes struct {
	TotalLikes int      `json:"total_likes"`
	LikedBy    []string `json:"liked_by"`
}

type Tweet struct {
	Author      string `json:"author"`
	Tweet       string `json:"tweet"`
	PublishedOn string `json:"published_on"`
	Likes       Likes
}
