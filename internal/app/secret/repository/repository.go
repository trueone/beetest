package repository

import (
	"errors"

	"github.com/trueone/beetest/internal/app/entity"
)

type Repository interface {
	Create(secret *entity.Secret) error
	UpdateViewCount(secret *entity.Secret) error
	Delete(secret *entity.Secret) error
	Get(secret *entity.Secret) error
}

var (
	ErrAlreadyExist = errors.New("secret already exist")
	ErrNotExist     = errors.New("secret doesn't exist")
	ErrNotFound     = errors.New("secret not found")
)
