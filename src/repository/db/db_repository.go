package db

import (
	"github.com/nhsh1997/bookstore_oauth-api/src/domain/access_token"
	"github.com/nhsh1997/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}

func NewRepository() DbRepository  {
	return &dbRepository{}
}

type dbRepository struct {

}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}