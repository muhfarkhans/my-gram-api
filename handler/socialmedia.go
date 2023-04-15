package handler

import (
	"errors"
	"my-gram/auth"
	"my-gram/helper"
	"my-gram/socialmedia"
	"my-gram/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type socialmediaHandler struct {
	service     socialmedia.Service
	authService auth.Service
}

func NewSocialMediaHandler(service socialmedia.Service, authService auth.Service) *socialmediaHandler {
	return &socialmediaHandler{service, authService}
}

func (h *socialmediaHandler) GetSocialMedias(c *gin.Context) {
	socialmedias, err := h.service.GetAllSocialmedia()
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error get socialmedia", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := socialmedia.FormatSocialMedias(socialmedias)
	response := helper.APIResponse("success get socialmedia", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *socialmediaHandler) FindSocialMediaById(c *gin.Context) {
	var input socialmedia.GetIdUri
	err := c.ShouldBindUri(&input)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error get socialmedia", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedSocialMedia, err := h.service.GetSocialmediaById(input.Id)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error get socialmedia", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if updatedSocialMedia.Id == 0 {
		errors := gin.H{"errors": errors.New("socialmedia not found")}

		response := helper.APIResponse("socialmedia not found", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := socialmedia.FormatSocialMedia(updatedSocialMedia)
	response := helper.APIResponse("success get socialmedia", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *socialmediaHandler) CreateSocialMedia(c *gin.Context) {
	var input socialmedia.CreateSocialMediaInput
	err := c.ShouldBind(&input)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error create socialmedia", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newSocialMedia, err := h.service.CreateSocialmedia(input)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error create socialmedia", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success create socialmedia", http.StatusOK, "success", newSocialMedia)
	c.JSON(http.StatusOK, response)
}

func (h *socialmediaHandler) UpdateSocialMedia(c *gin.Context) {
	var input socialmedia.CreateSocialMediaInput
	err := c.ShouldBind(&input)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error update socialmedia", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputUri socialmedia.GetIdUri
	err = c.ShouldBindUri(&inputUri)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error update socialmedia", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	updatedSocialMedia, err := h.service.UpdateSocialmedia(input, inputUri.Id)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error update socialmedia", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := socialmedia.FormatSocialMedia(updatedSocialMedia)
	response := helper.APIResponse("success update socialmedia", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *socialmediaHandler) DeleteSocialMedia(c *gin.Context) {
	var inputUri socialmedia.GetIdUri
	err := c.ShouldBindUri(&inputUri)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error delete socialmedia", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	_, err = h.service.DeleteSocialmedia(inputUri.Id, currentUser.Id)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error delete socialmedia", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := gin.H{"message": "deleted"}
	response := helper.APIResponse("success delete socialmedia", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
