package repository

import (
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
	defer r.mu.Unlock()

	if _, exist := r.m[secret.Hash]; !exist {
		r.m[secret.Hash] = *secret
		return nil
	}

	return ErrAlreadyExist
}

// Update view count
func (r *repositoryMap) UpdateViewCount(secret *entity.Secret) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exist := r.m[secret.Hash]; exist {
		secret.ViewsCount++
		r.m[secret.Hash] = *secret

		return nil
	}

	return ErrNotExist
}

// Delete
func (r *repositoryMap) Delete(secret *entity.Secret) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exist := r.m[secret.Hash]; exist {
		delete(r.m, secret.Hash)

		return nil
	}

	return ErrNotExist
}

// Get secret
func (r *repositoryMap) Get(secret *entity.Secret) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if s, exist := r.m[secret.Hash]; exist {
		*secret = s
		return nil
	}

	return ErrNotFound
}
