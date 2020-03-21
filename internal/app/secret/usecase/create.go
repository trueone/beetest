package usecase

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"github.com/trueone/beetest/internal/app/entity"
	"github.com/trueone/beetest/internal/app/secret/dto"
	"github.com/trueone/beetest/internal/app/secret/repository"
)

// Create scenario
type Create interface {
	validate(request dto.SecretRequest) error
	construct(request dto.SecretRequest) entity.Secret
	save(secret *entity.Secret) error

	Execute(request dto.SecretRequest) (entity.Secret, error)
}

// Create inter actor
type create struct {
	repository repository.Repository
}

// Create constructor
func NewCreate(sr repository.Repository) Create {
	c := new(create)
	c.repository = sr

	return c
}

// Validate fields
func (c *create) validate(request dto.SecretRequest) error {
	if request.Secret == "" {
		return errors.New("secret text is required")
	}

	if request.ExpireAfterViews == 0 {
		return errors.New("expire after view must be greater than 0")
	}

	return nil
}

// Create Secret object
func (c *create) construct(request dto.SecretRequest) entity.Secret {
	hash := md5.Sum([]byte(request.Secret))
	now := time.Now()

	return entity.Secret{
		Hash:           hex.EncodeToString(hash[:]),
		Text:           request.Secret,
		Created:        now,
		Expires:        now.Add(time.Minute * time.Duration(request.ExpireAfter)),
		RemainingViews: request.ExpireAfterViews,
	}
}

// Save secret
func (c *create) save(secret *entity.Secret) error {
	return c.repository.Create(secret)
}

// Execute entry point
func (c *create) Execute(request dto.SecretRequest) (secret entity.Secret, err error) {
	if err = c.validate(request); err != nil {
		return
	}

	secret = c.construct(request)

	if err = c.save(&secret); err != nil {
		return
	}

	return
}
