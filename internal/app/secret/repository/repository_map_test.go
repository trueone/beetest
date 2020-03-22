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
	assert.Error(t, repo.Create(secret), "Create already existing secret")

	assert.NoError(t, repo.Get(secret), "Get first secret")
	assert.Error(t, repo.Get(&entity.Secret{}), "Get not existed")

	assert.NoError(t, repo.UpdateViewCount(secret), "Update view count")
	assert.Error(t, repo.UpdateViewCount(&entity.Secret{}), "Update not existed")

	assert.NoError(t, repo.Delete(secret), "Delete secret")
	assert.Error(t, repo.Delete(secret), "Delete not existed")
}
