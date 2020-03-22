package usecase

import (
	"errors"
	"time"

	"github.com/trueone/beetest/internal/app/entity"
	"github.com/trueone/beetest/internal/app/secret/dto"
	"github.com/trueone/beetest/internal/app/secret/repository"
)

// Get scenario
type Get interface {
	find(hash dto.HashRequest) (entity.Secret, error)
	update(*entity.Secret) error
	validate(*entity.Secret) error

	Execute(hash dto.HashRequest) (entity.Secret, error)
}

// Get inter actor
type get struct {
	repository repository.Repository
}

// Get constructor
func NewGet(sr repository.Repository) Get {
	g := new(get)
	g.repository = sr

	return g
}

// Find secret
func (g *get) find(hash dto.HashRequest) (secret entity.Secret, err error) {
	secret.Hash = string(hash)
	err = g.repository.Get(&secret)

	return
}

// Update view count
func (g *get) update(secret *entity.Secret) error {
	return g.repository.UpdateViewCount(secret)
}

// Validate max view count & expire time
func (g *get) validate(secret *entity.Secret) (err error) {
	if secret.IsExceedLimit() {
		err = errors.New("view limit is exceeded")
	}

	if secret.IsExpired(time.Now()) {
		err = errors.New("secret expired")
	}

	if err != nil {
		_ = g.repository.Delete(secret)
	}

	return err
}

// Execute entry point
func (g *get) Execute(hash dto.HashRequest) (secret entity.Secret, err error) {
	secret, err = g.find(hash)
	if err != nil {
		return
	}

	err = g.update(&secret)
	if err != nil {
		return
	}

	err = g.validate(&secret)
	if err != nil {
		return
	}

	return
}
