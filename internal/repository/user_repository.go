package repository

import (
	"github.com/heshify/redoed/internal/db"
	"github.com/heshify/redoed/internal/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(user *models.User) (string, error) {
	result := db.DB.Create(user)
	return user.ID.String(), result.Error
}

func (r *UserRepository) GetUser(id string) (models.User, error) {
	var user models.User
	result := db.DB.First(&user, "id = ?", id)
	return user, result.Error
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := db.DB.First(&user, "email = ?", email)
	return user, result.Error
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	result := db.DB.Find(&users)
	return users, result.Error
}

func (r *UserRepository) UpdateUser(id string, user *models.User) error {
	var updatedUser models.User
	result := db.DB.Model(&updatedUser).Where("id = ?", id).Updates(user)
	return result.Error
}

func (r *UserRepository) DeleteUser(id string) error {
	result := db.DB.Delete(&models.User{}, "id = ?", id)
	return result.Error
}
