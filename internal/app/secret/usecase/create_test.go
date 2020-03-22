package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/trueone/beetest/internal/app/secret/dto"
	"github.com/trueone/beetest/internal/app/secret/repository"
	"github.com/trueone/beetest/internal/app/secret/usecase"
)

var (
	repo = repository.NewRepositoryMap()
)

func TestCreate_Execute(t *testing.T) {
	cases := []struct {
		name  string
		valid bool
		data  dto.SecretRequest
	}{
		{
			name:  "Create secret",
			valid: true,
			data: dto.SecretRequest{
				Secret:           "secret",
				ExpireAfterViews: 1,
				ExpireAfter:      0,
			},
		},
		{
			name:  "Create existing secret",
			valid: false,
			data: dto.SecretRequest{
				Secret:           "secret",
				ExpireAfterViews: 1,
				ExpireAfter:      0,
			},
		},
		{
			name:  "Create empty secret",
			valid: false,
			data: dto.SecretRequest{
				Secret:           "",
				ExpireAfterViews: 1,
				ExpireAfter:      0,
			},
		},
		{
			name:  "Create zero view secret",
			valid: false,
			data: dto.SecretRequest{
				Secret:           "",
				ExpireAfterViews: 1,
				ExpireAfter:      0,
			},
		},
	}

	create := usecase.NewCreate(repo)
	for _, c := range cases {
		_, err := create.Execute(c.data)

		if c.valid {
			assert.NoError(t, err, c.name)
		} else {
			assert.Error(t, err, c.name)
		}
	}
}
