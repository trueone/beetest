package repository

import (
	"github.com/trueone/beetest/internal/app/entity"
)

type Repository interface {
	Create(secret *entity.Secret) error
	Get(secret *entity.Secret) error
}
