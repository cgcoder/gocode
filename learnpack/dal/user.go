package dal

import (
	"github.com/gocode/learnpack/models"
)

type UserDal interface {
	NewUser(user *models.User) (*models.User, error)
	GetByName(name *string) (*models.User, error)
	GetByEmail(email *string) (*models.User, error)
	SetEmailVerified(id models.UserId, verified bool) error
	SetPassReset(id models.UserId) error
	UpdatePassword(name *string, pass *string) error
	DeleteUser(id models.UserId) error
}
