package repositories

import (
	"github.com/ronnyscale/layerarch/internal/models"
	"github.com/ronnyscale/layerarch/pkg/database"
)

func SaveUser(user *models.User) error {
	result := db.DB.Create(user)
	return result.Error
}
