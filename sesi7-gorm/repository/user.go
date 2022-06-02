package repository

import (
	"sesi7-gorm/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(*models.User) error
	GetAllUsers() (*[]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	GetUsersWithProducts() (*[]models.User, error)
	//Latihan
	UpdateUser(*models.User) error
	DeleteUser(id uint) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(request *models.User) error {
	err := r.db.Create(request).Error
	return err
}

func (r *userRepo) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return &users, err
}

func (r *userRepo) GetUserByID(id uint) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, "id=?", id).Error
	return &user, err
}

func (r *userRepo) GetUsersWithProducts() (*[]models.User, error) {
	var users []models.User
	err := r.db.Preload("Products").Find(&users).Error
	return &users, err
}

//Latihan
func (r *userRepo) UpdateUser(request *models.User) error {
	err := r.db.Exec("UPDATE users SET email = ? WHERE id = ?", request.Email, request.ID).Error
	return err
}

func (r *userRepo) DeleteUser(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Where("id=?", id).Delete(&user).Error
	return &user, err
}
