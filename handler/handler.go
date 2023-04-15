package handler

import (
	"fmt"
	"my-gram/auth"
	"my-gram/comment"
	"my-gram/config"
	"my-gram/db"
	"my-gram/helper"
	"my-gram/photo"
	"my-gram/socialmedia"
	"my-gram/user"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func NewHandler(router *gin.Engine) {
	dbConfig := config.PostgreConfig()
	database := db.InitDatabase(dbConfig)

	api := router.Group("v1")

	authService := auth.NewService()

	userRepository := user.NewRepository(database)
	userService := user.NewService(userRepository)
	userHandler := NewUserHandler(userService, authService)
	routerUser := api.Group("users")
	routerUser.POST("/register", userHandler.RegisterNewUser)
	routerUser.POST("/login", userHandler.Login)

	auth := authMiddleware(authService, userService)

	photoRepository := photo.NewRepository(database)
	photoService := photo.NewService(photoRepository)
	photoHandler := NewPhotoHandler(photoService, authService)
	routerPhoto := api.Group("/photos")
	routerPhoto.GET("/", photoHandler.GetPhotos)
	routerPhoto.GET("/:id", photoHandler.FindPhotoById)
	routerPhoto.POST("/", auth, photoHandler.CreatePhoto)
	routerPhoto.PUT("/:id", auth, photoHandler.UpdatePhoto)
	routerPhoto.DELETE("/:id", auth, photoHandler.DeletePhoto)

	socialmediaRepository := socialmedia.NewRepository(database)
	socialmediaService := socialmedia.NewService(socialmediaRepository)
	socialmediaHandler := NewSocialMediaHandler(socialmediaService, authService)
	routerSocialMedia := api.Group("/socialmedias")
	routerSocialMedia.GET("/", socialmediaHandler.GetSocialMedias)
	routerSocialMedia.GET("/:id", socialmediaHandler.FindSocialMediaById)
	routerSocialMedia.POST("/", auth, socialmediaHandler.CreateSocialMedia)
	routerSocialMedia.PUT("/:id", auth, socialmediaHandler.UpdateSocialMedia)
	routerSocialMedia.DELETE("/:id", auth, socialmediaHandler.DeleteSocialMedia)

	commentRepository := comment.NewRepository(database)
	commentService := comment.NewService(commentRepository)
	commentHandler := NewCommentHandler(commentService, authService)
	routerComment := api.Group("/comments")
	routerComment.GET("/", commentHandler.GetComments)
	routerComment.GET("/:id", commentHandler.FindCommentById)
	routerComment.POST("/", auth, commentHandler.CreateComment)
	routerComment.PUT("/:id", auth, commentHandler.UpdateComment)
	routerComment.DELETE("/:id", auth, commentHandler.DeleteComment)
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			tokenString = tokenArray[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		fmt.Println(tokenArray)
		fmt.Println(tokenString)

		claim, ok := token.Claims.(jwt.MapClaims)
		fmt.Println(ok)
		fmt.Println(token)
		fmt.Println(claim)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))
		user, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
