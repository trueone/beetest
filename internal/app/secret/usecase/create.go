package usecase

import (
	"github.com/trueone/beetest/internal/app/entity"
	"github.com/trueone/beetest/internal/app/secret/repository"
)

// Create scenario
type Create interface {
	validate(*entity.Secret) error
	hash(*entity.Secret) error
	save(*entity.Secret) error

	Execute(*entity.Secret) error
}

// create inter actor
type create struct {
	repository repository.Repository
}

// Create constructor
func NewCreate(sr repository.Repository) Create {
	c := new(create)
	c.repository = sr

	return c
}

// validate fields
func (c *create) validate(secret *entity.Secret) error {
	return nil
}

// get hash from text
func (c *create) hash(secret *entity.Secret) error {
	return nil
}

func (c *create) save(secret *entity.Secret) error {
	return nil
}

// Execute entry point
func (c *create) Execute(secret *entity.Secret) error {
	return nil
}
