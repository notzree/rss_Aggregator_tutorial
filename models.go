package main

import (
	"strconv"
	"time"

	nanoid "github.com/matoous/go-nanoid/v2"

	"gorm.io/gorm"
)

const (
	alphabet = "0123456789"
	length   = 12
)

type Handler struct {
	db *gorm.DB
}
type Users struct {
	Id         int       `json:"id"`
	Created_at time.Time `json:"created_at" gorm:"type:datetime"`
	Updated_at time.Time `json:"updated_at" gorm:"type:datetime"`
	Name       string    `json:"name"`
	Feeds      []Feeds   `json:"feeds" gorm:"foreignKey:UserId"`
}

type Feeds struct {
	Id         int       `json:"id"`
	Created_at time.Time `json:"created_at" gorm:"type:datetime"`
	Updated_at time.Time `json:"updated_at" gorm:"type:datetime"`
	Name       string    `json:"name"`
	Url        string    `json:"url" gorm:"not null"`
	UserId     int       `json:"user_id"`
}
type QueryFeed struct {
	Id int `json:"id"`
}

func databaseUserToUser(user Users) Users {
	return Users{
		Id:         user.Id,
		Created_at: user.Created_at,
		Updated_at: user.Updated_at,
		Name:       user.Name,
		Feeds:      user.Feeds,
	}

}
func databaseFeedToFeed(feed Feeds) Feeds {
	return Feeds{
		Id:         feed.Id,
		Created_at: feed.Created_at,
		Updated_at: feed.Updated_at,
		Name:       feed.Name,
		Url:        feed.Url,
		UserId:     feed.UserId,
		//dont need to return user object everytime we create a feed.
	}
}

func NewId() (int, error) {
	temp, _ := nanoid.Generate(alphabet, length)
	return strconv.Atoi(temp)

}
