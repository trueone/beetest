package repository

import (
	"github.com/trueone/beetest/internal/app/entity"
)

type Repository interface {
	Save(*entity.Secret) error
	Get(secret *entity.Secret) error
}
