package db

import (
	"github.com/suvamsingh/bookstore_oauth-api/src/domain/access_token"
	"github.com/suvamsingh/bookstore_oauth-api/src/utils/errors"
)

//DBRepository ...
type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

//New ...
func NewRepository() DBRepository {
	return &dbRepository{}
}

//GetByID ...
func (r *dbRepository) GetByID(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")

}
