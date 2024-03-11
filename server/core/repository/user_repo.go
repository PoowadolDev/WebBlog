package repository

import (
	"Web_Blog/core/domain"
)

type UserRepository interface {
	Save(user domain.User) error
}
