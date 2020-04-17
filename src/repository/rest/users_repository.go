package rest

import (
	"github.com/nhsh1997/bookstore_oauth-api/src/domain/users"
	"github.com/nhsh1997/bookstore_oauth-api/src/utils/errors"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

type usersRepository struct {

}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestError) {
	return nil, nil
}