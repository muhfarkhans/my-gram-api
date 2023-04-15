package socialmedia

import (
	"my-gram/user"
	"time"
)

type Socialmedia struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           user.User `json:"user"`
}

func (Socialmedia) TableName() string {
	return "socialmedias"
}
