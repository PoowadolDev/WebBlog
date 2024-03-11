package adapters

import (
	"Web_Blog/core/domain"
	"Web_Blog/core/repository"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Save(user domain.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
