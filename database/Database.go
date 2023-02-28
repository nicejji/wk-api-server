package database

import "time"

type User struct {
	Id          int          `json:"id"`
	Nickname    string       `json:"nickname"`
	Password    string       `json:"password"`
	Bio         string       `json:"bio"`
	Image       string       `json:"image"`
	Created     time.Time    `json:"created"`
	Memberships []Membership `json:"memberships"`
}

type Chat struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Private  bool      `json:"private"`
	Created  time.Time `json:"created"`
	Users    []User    `json:"users"`
	Messages []Message `json:"messages"`
}

type Membership struct {
	Id      int       `json:"id"`
	Chat    Chat      `json:"chat"`
	Blocked bool      `json:"blocked"`
	Created time.Time `json:"created"`
}

type Message struct {
	Id      int       `json:"id"`
	User    User      `json:"user"`
	Chat    Chat      `json:"chat"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}
