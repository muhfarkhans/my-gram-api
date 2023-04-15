package handler

import (
	"errors"
	"my-gram/auth"
	"my-gram/comment"
	"my-gram/helper"
	"my-gram/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type commentHandler struct {
	service     comment.Service
	authService auth.Service
}

func NewCommentHandler(service comment.Service, authService auth.Service) *commentHandler {
	return &commentHandler{service, authService}
}

func (h *commentHandler) GetComments(c *gin.Context) {
	comments, err := h.service.GetAllComment()
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error get comment", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := comment.FormatComments(comments)
	response := helper.APIResponse("success get comment", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *commentHandler) FindCommentById(c *gin.Context) {
	var input comment.GetIdUri
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
		response := helper.APIResponse("error get comment", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedComment, err := h.service.GetCommentById(input.Id)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error get comment", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if updatedComment.Id == 0 {
		errors := gin.H{"errors": errors.New("comment not found")}

		response := helper.APIResponse("comment not found", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := comment.FormatComment(updatedComment)
	response := helper.APIResponse("success get comment", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *commentHandler) CreateComment(c *gin.Context) {
	var input comment.CreateCommentInput
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
		response := helper.APIResponse("error create comment", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newComment, err := h.service.CreateComment(input)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error create comment", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success create comment", http.StatusOK, "success", newComment)
	c.JSON(http.StatusOK, response)
}

func (h *commentHandler) UpdateComment(c *gin.Context) {
	var input comment.CreateCommentInput
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
		response := helper.APIResponse("error update comment", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputUri comment.GetIdUri
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
		response := helper.APIResponse("error update comment", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	updatedComment, err := h.service.UpdateComment(input, inputUri.Id)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error update comment", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := comment.FormatComment(updatedComment)
	response := helper.APIResponse("success update comment", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *commentHandler) DeleteComment(c *gin.Context) {
	var inputUri comment.GetIdUri
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
		response := helper.APIResponse("error delete comment", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	_, err = h.service.DeleteComment(inputUri.Id, currentUser.Id)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error delete comment", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := gin.H{"message": "deleted"}
	response := helper.APIResponse("success delete comment", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
