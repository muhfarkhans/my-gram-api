package comment

import (
	"my-gram/photo"
	"my-gram/user"
	"time"
)

type Comment struct {
	Id        int         `json:"id"`
	Message   string      `json:"message"`
	UserId    int         `json:"user_id"`
	PhotoId   int         `json:"photo_id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	User      user.User   `json:"user"`
	Photo     photo.Photo `json:"photo"`
}
