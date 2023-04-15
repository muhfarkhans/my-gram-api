package comment

import (
	"my-gram/photo"
	"my-gram/user"
	"time"
)

type CommentFormatter struct {
	Id        int           `json:"id"`
	Message   string        `json:"title"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      user.UserSafe `json:"user"`
	Photo     photo.Photo   `json:"photo"`
}

func FormatComment(comment Comment) CommentFormatter {
	formatter := CommentFormatter{}
	formatter.Id = comment.Id
	formatter.Message = comment.Message
	formatter.CreatedAt = comment.CreatedAt
	formatter.UpdatedAt = comment.UpdatedAt
	formatter.User.Id = comment.User.Id
	formatter.User.Username = comment.User.Username
	formatter.User.Email = comment.User.Email
	formatter.Photo.Title = comment.Photo.Title
	formatter.Photo.Caption = comment.Photo.Caption
	formatter.Photo.PhotoUrl = comment.Photo.PhotoUrl
	formatter.Photo.User.Id = comment.Photo.User.Id
	formatter.Photo.User.Username = comment.Photo.User.Username
	formatter.Photo.User.Email = comment.Photo.User.Email

	return formatter
}

func FormatComments(comments []Comment) []CommentFormatter {
	commentsFormatter := []CommentFormatter{}

	for _, comment := range comments {
		commentFormatter := FormatComment(comment)
		commentsFormatter = append(commentsFormatter, commentFormatter)
	}

	return commentsFormatter
}
