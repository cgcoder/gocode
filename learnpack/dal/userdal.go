package dal

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/gocode/learnpack/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userDalImpl struct {
	*pgxpool.Pool
}

// NewUserDal initialize dal layer
func NewUserDal() (UserDal, error) {
	pool := GetPool()
	if pool == nil {
		return nil, errors.New("pool not initialized")
	}
	return &userDalImpl{Pool: pool}, nil
}

func (dal *userDalImpl) NewUser(user *models.User) (*models.User, error) {

	sql := `
	INSERT INTO USERS(ID,NAME,EMAIL,PASS)
	VALUES ($1,$2,$3,$4)
	`

	conn, err := dal.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	done := false
	for !done {
		user.ID = models.UserId(time.Now().UnixNano())
		_, err = dal.Exec(context.Background(), sql, user.ID, user.Name, user.Email, user.Pass)
		if err == nil {
			done = true
		} else if strings.Index(err.Error(), "duplicate key") == -1 {
			return nil, err
		}
	}

	return user, nil
}

func (dal *userDalImpl) GetByName(name *string) (*models.User, error) {
	return nil, nil
}

func (dal *userDalImpl) GetByEmail(email *string) (*models.User, error) {
	return nil, nil
}

func (dal *userDalImpl) SetEmailVerified(id models.UserId, verified bool) error {
	return nil
}

func (dal *userDalImpl) SetPassReset(id models.UserId) error {
	return nil
}

func (dal *userDalImpl) UpdatePassword(name *string, pass *string) error {
	return nil
}

func (dal *userDalImpl) DeleteUser(id models.UserId) error {
	sql := `DELETE FROM USERS WHERE ID=$1`

	conn, err := dal.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = dal.Exec(context.Background(), sql, int64(id))

	return err
}
