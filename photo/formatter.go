package photo

import (
	"my-gram/user"
	"time"
)

type PhotoFormatter struct {
	Id        int           `json:"id"`
	Title     string        `json:"title"`
	Caption   string        `json:"caption"`
	PhotoUrl  string        `json:"photo_url"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      user.UserSafe `json:"user"`
}

func FormatPhoto(photo Photo) PhotoFormatter {
	formatter := PhotoFormatter{}
	formatter.Id = photo.Id
	formatter.Title = photo.Title
	formatter.Caption = photo.Caption
	formatter.PhotoUrl = photo.PhotoUrl
	formatter.CreatedAt = photo.CreatedAt
	formatter.UpdatedAt = photo.UpdatedAt
	formatter.User.Id = photo.User.Id
	formatter.User.Username = photo.User.Username
	formatter.User.Email = photo.User.Email

	return formatter
}

func FormatPhotos(photos []Photo) []PhotoFormatter {
	photosFormatter := []PhotoFormatter{}

	for _, photo := range photos {
		photoFormatter := FormatPhoto(photo)
		photosFormatter = append(photosFormatter, photoFormatter)
	}

	return photosFormatter
}
