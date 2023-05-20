package models

type Follower struct {
	UserId     string `json:"user_id"`
	FollowerId string `json:"follower_id"`
	FollowedAt string `json:"followed_at"`
}
