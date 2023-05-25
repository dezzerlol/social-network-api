package models

type Like struct {
	UserId int64 `json:"user_id"`
	PostId int64 `json:"post_id"`
}
