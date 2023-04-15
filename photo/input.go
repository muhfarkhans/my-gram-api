package photo

import "my-gram/user"

type CreatePhotoInput struct {
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url" binding:"required"`
	User     user.User
}

type GetIdUri struct {
	Id int `uri:"id" binding:"required"`
}
