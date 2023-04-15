package comment

import (
	"my-gram/photo"
	"my-gram/user"
)

type CreateCommentInput struct {
	Message string `json:"message" form:"message" binding:"required"`
	PhotoId int    `json:"photo_id" form:"photo_id" binding:"required"`
	User    user.User
	Photo   photo.Photo
}

type GetIdUri struct {
	Id int `uri:"id" binding:"required"`
}
