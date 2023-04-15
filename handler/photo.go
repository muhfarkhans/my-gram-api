package handler

import (
	"errors"
	"my-gram/auth"
	"my-gram/helper"
	"my-gram/photo"
	"my-gram/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type photoHandler struct {
	service     photo.Service
	authService auth.Service
}

func NewPhotoHandler(service photo.Service, authService auth.Service) *photoHandler {
	return &photoHandler{service, authService}
}

func (h *photoHandler) GetPhotos(c *gin.Context) {
	photos, err := h.service.GetAllPhoto()
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error get photo", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := photo.FormatPhotos(photos)
	response := helper.APIResponse("success get photo", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *photoHandler) FindPhotoById(c *gin.Context) {
	var input photo.GetIdUri
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
		response := helper.APIResponse("error get photo", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedPhoto, err := h.service.GetPhotoById(input.Id)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error get photo", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if updatedPhoto.Id == 0 {
		errors := gin.H{"errors": errors.New("photo not found")}

		response := helper.APIResponse("photo not found", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := photo.FormatPhoto(updatedPhoto)
	response := helper.APIResponse("success get photo", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *photoHandler) CreatePhoto(c *gin.Context) {
	var input photo.CreatePhotoInput
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
		response := helper.APIResponse("error create photo", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newPhoto, err := h.service.CreatePhoto(input)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error create photo", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success create photo", http.StatusOK, "success", newPhoto)
	c.JSON(http.StatusOK, response)
}

func (h *photoHandler) UpdatePhoto(c *gin.Context) {
	var input photo.CreatePhotoInput
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
		response := helper.APIResponse("error update photo", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputUri photo.GetIdUri
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
		response := helper.APIResponse("error update photo", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	updatedPhoto, err := h.service.UpdatePhoto(input, inputUri.Id)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error update photo", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := photo.FormatPhoto(updatedPhoto)
	response := helper.APIResponse("success update photo", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *photoHandler) DeletePhoto(c *gin.Context) {
	var inputUri photo.GetIdUri
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
		response := helper.APIResponse("error delete photo", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	_, err = h.service.DeletePhoto(inputUri.Id, currentUser.Id)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error delete photo", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := gin.H{"message": "deleted"}
	response := helper.APIResponse("success delete photo", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
