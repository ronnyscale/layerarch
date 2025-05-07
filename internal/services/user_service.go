package services

import (
    "golang.org/x/crypto/bcrypt"
    "github.com/ronnyscale/layerarch/internal/dto"
    "github.com/ronnyscale/layerarch/internal/models"
    "github.com/ronnyscale/layerarch/internal/repositories"
)

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CreateUser(dto dto.CreateUserDTO) (*models.User, error) {
    hashedPassword, err := hashPassword(dto.Password)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        Email:    dto.Email,
        Password: hashedPassword,
    }

    err = repositories.SaveUser(user)
    if err != nil {
        return nil, err
    }

    return user, nil
}