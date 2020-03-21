package repository

import (
	"github.com/trueone/beetest/internal/app/entity"
)

type Repository interface {
	Create(*entity.Secret) error
	Get(secret *entity.Secret) error
}
