package users

import (
	"gorm.io/gorm"
)

type Repository interface {
	// GetAll() ([]User, error)
	FindByEmailLike(email string) (User, error)
	// FindByUsernameLike(email string) (User, error)
	Register(user User) (User, error)
	// FindByID(ID uuid.UUID) (User, error)
	// Update(book User) (User, error)
	// Delete(book User) (User, error)
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

// func (r *repository) FindByUsernameLike(username string) (User, error) {
// 	var user User
//     err := r.db.Where("username LIKE ?", username).First(&user).Error
//     return user, err
// }

func (r *repository) Register(user User) (User, error) {
	err := r.db.Create(&user).Error
    return user, err
}

// func (r *repository) FindByID(ID uuid.UUID) (User, error) {
// 	var user User
//     err := r.db.Find(&user, ID).Error
//     return user, err
// }

// func (r *repository) GetAll() ([]User, error) {
//     var users []User

//     query := r.db.Model(&User{})
// 	// if queries.Title != "" {
//     //     query = query.Where("title ILIKE ?", "%"+queries.Title+"%")
//     // }
// 	// if queries.Description != "" {
//     //     query = query.Or("description ILIKE ?", "%"+queries.Description+"%")
//     // }
// 	err := query.Find(&users).Error

//     return users, err
// }