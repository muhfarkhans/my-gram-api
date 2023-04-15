package socialmedia

import (
	"my-gram/user"
	"time"
)

type SocialMediaFormatter struct {
	Id             int           `json:"id"`
	Name           string        `json:"name"`
	SocialMediaUrl string        `json:"social_media_url"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	User           user.UserSafe `json:"user"`
}

func FormatSocialMedia(socialmedia Socialmedia) SocialMediaFormatter {
	formatter := SocialMediaFormatter{}
	formatter.Id = socialmedia.Id
	formatter.Name = socialmedia.Name
	formatter.SocialMediaUrl = socialmedia.SocialMediaUrl
	formatter.CreatedAt = socialmedia.CreatedAt
	formatter.UpdatedAt = socialmedia.UpdatedAt
	formatter.User.Id = socialmedia.User.Id
	formatter.User.Username = socialmedia.User.Username
	formatter.User.Email = socialmedia.User.Email

	return formatter
}

func FormatSocialMedias(socialmedias []Socialmedia) []SocialMediaFormatter {
	socialmediasFormatter := []SocialMediaFormatter{}

	for _, socialmedia := range socialmedias {
		socialmediaFormatter := FormatSocialMedia(socialmedia)
		socialmediasFormatter = append(socialmediasFormatter, socialmediaFormatter)
	}

	return socialmediasFormatter
}
