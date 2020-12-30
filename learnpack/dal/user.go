package dal

import (
    "model"
    "errors"
)

type UserDal interface {
     NewUser(user* User) (*User, error)
     GetByName(name* string) (*User, error)
     GetByEmail(email* string) (*User, error)
     SetEmailVerified(id UserId, verified bool) error
     SetPassReset(id UserId) error
     UpdatePassword(name* string, pass* string) error
}
