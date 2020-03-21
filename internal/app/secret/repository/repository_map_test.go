package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/trueone/beetest/internal/app/entity"
	"github.com/trueone/beetest/internal/app/secret/repository"
)

var (
	repo   = repository.NewRepositoryMap()
	secret = &entity.Secret{Hash: "secret"}
)

func TestNewRepositoryMap(t *testing.T) {
	assert.NoError(t, repo.Create(secret), "Create secret")
	assert.NoError(t, repo.Get(secret), "Get first secret")
	assert.Error(t, repo.Get(&entity.Secret{}), "Get not existed")
}
