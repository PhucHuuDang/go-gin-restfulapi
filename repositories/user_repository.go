package repositories

import (
	"fmt"
	"golang-restapi/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.UserModel{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *UserRepository) ExistsById(id uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.UserModel{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err

}

func (r *UserRepository) FindUserById(id uint) (*models.UserModel, error) {
	var user models.UserModel

	err := r.db.Preload("Posts").First(&user, id).Error
	return &user, err
}

func (r *UserRepository) CreatUser(user *models.UserModel) error {

	return r.db.Create(user).Error
}

func (r *UserRepository) UpdateUser(id uint, data map[string]interface{}) error {
	return r.db.Model(&models.UserModel{}).Where("id = ?", id).Updates(data).Error
}

func (r *UserRepository) DeleteUser(user *models.UserModel, id uint) error {
	fmt.Println("id: ", id)
	return r.db.Delete(user, id).Error
}

func (r *UserRepository) GetAllUsers() ([]models.UserModel, error) {
	var users []models.UserModel

	err := r.db.Find(&users).Error

	return users, err

}
