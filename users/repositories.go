package users

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindByEmailLike(email string) (User, error)
	Register(user User) (User, error)
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByEmailLike(email string) (User, error) {
	var user User
    err := r.db.Where("email LIKE ?", email).First(&user).Error
    return user, err
}

func (r *repository) Register(user User) (User, error) {
	err := r.db.Create(&user).Error
    return user, err
}