package models

import "time"

type User struct {
	Id         string    `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Firstname  string    `json:"first_name"`
	Lastname   string    `json:"last_name"`
	Username   string    `json:"username"`
	Avatar     string    `json:"avatar"`
	Birthdate  time.Time `json:"birthdate"`
	Activated  bool      `json:"activated"`
	Created_at time.Time `json:"created_at"`
}
