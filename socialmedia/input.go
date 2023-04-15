package socialmedia

import "my-gram/user"

type CreateSocialMediaInput struct {
	Name           string `json:"name" form:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" binding:"required"`
	User           user.User
}

type GetIdUri struct {
	Id int `uri:"id" binding:"required"`
}
