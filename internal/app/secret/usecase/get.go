package usecase

import (
	"errors"
	"time"

	"github.com/trueone/beetest/internal/app/entity"
	"github.com/trueone/beetest/internal/app/secret/data"
	"github.com/trueone/beetest/internal/app/secret/repository"
)

// Get scenario
type Get interface {
	find(hash data.HashRequest) (entity.Secret, error)
	update(secret *entity.Secret) error
	validate(secret *entity.Secret) error

	Execute(hash data.HashRequest) (entity.Secret, error)
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
func (g *get) find(hash data.HashRequest) (secret entity.Secret, err error) {
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
func (g *get) Execute(hash data.HashRequest) (secret entity.Secret, err error) {
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
