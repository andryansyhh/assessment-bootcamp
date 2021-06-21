package handler

import (
	"assesment/auth"
	"assesment/entity"
	"assesment/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service     user.Service
	authService auth.Service
}

func NewUserHandler(service user.Service, authService auth.Service) *userHandler {
	return &userHandler{service, authService}

}

func (h *userHandler) GetAllUser(c *gin.Context) {

	users, err := h.service.GetAllUser()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, users)
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	paramID := c.Params.ByName("user_id")
	user, err := h.service.GetUserByID(paramID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, user)
}

func (h *userHandler) UserRegister(c *gin.Context) {
	var input entity.UserRegister

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.UserRegister(input)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, user)
}

func (h *userHandler) UserLogin(c *gin.Context) {
	var input entity.UserLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.UserLogin(input)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, user)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	paramID := c.Params.ByName("user_id")
	var input entity.UpdateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	user, err := h.service.UpdateUser(paramID, input)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, user)
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	paramID := c.Params.ByName("user_id")

	user, err := h.service.DeleteUser(paramID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, user)
}
