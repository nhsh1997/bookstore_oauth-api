package access_token

import (
	"github.com/nhsh1997/bookstore_oauth-api/src/repository/rest"
	"github.com/nhsh1997/bookstore_oauth-api/src/utils/errors"
	"strings"
)


type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}


type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service  {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestError){
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request AccessTokenRequest) ( *AccessToken, *errors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// Authenticate the user against the Users API:
	//Generate a new access token
	at := GetNewAccessToken()
	//Save the new access token in Cassandra:


	return &at, nil
}

func (s *service) UpdateExpirationTime( at AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}