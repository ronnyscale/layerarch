package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ronnyscale/layerarch/internal/dto"
	"github.com/ronnyscale/layerarch/internal/services"
)

func CreateUser(c *gin.Context) {
	var dto dto.CreateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.CreateUser(dto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
