package repository

import (
	"errors"
	"sync"

	"github.com/trueone/beetest/internal/app/entity"
)

// Repository map driver
type repositoryMap struct {
	mu sync.RWMutex
	m  map[string]entity.Secret
}

// NewRepositoryMap constructor
func NewRepositoryMap() Repository {
	r := new(repositoryMap)
	r.mu = sync.RWMutex{}
	r.m = make(map[string]entity.Secret)

	return r
}

// Save secret
func (r *repositoryMap) Create(secret *entity.Secret) error {
	r.mu.Lock()
	r.m[secret.Hash] = *secret
	r.mu.Unlock()

	return nil
}

// Get secret
func (r *repositoryMap) Get(secret *entity.Secret) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if s, ok := r.m[secret.Hash]; ok {
		*secret = s
		return nil
	}

	return errors.New("not found")
}
