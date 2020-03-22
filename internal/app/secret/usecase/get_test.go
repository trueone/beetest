package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/trueone/beetest/internal/app/entity"
	"github.com/trueone/beetest/internal/app/secret/dto"
	"github.com/trueone/beetest/internal/app/secret/repository"
	"github.com/trueone/beetest/internal/app/secret/usecase"
)

func TestGet_Execute(t *testing.T) {
	secret := entity.Secret{Hash: "secret", RemainingViews: 1}

	repo := repository.NewRepositoryMap()
	repo.Create(&secret)

	get := usecase.NewGet(repo)

	s, err := get.Execute(dto.HashRequest("secret"))
	assert.NoError(t, err, "Get secret")

	secret.ViewsCount++
	assert.Equal(t, secret, s, "Compare secret")

	_, err = get.Execute(dto.HashRequest("secret"))
	assert.Error(t, err, "Get limit exceeded")
}
