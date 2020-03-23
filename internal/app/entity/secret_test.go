package entity_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/trueone/beetest/internal/app/entity"
)

func TestSecret_IsExceedLimit(t *testing.T) {
	secret := entity.Secret{
		RemainingViews: 1,
		ViewsCount:     0,
	}

	assert.Equal(t, false, secret.IsExceedLimit())

	secret.ViewsCount++
	assert.Equal(t, false, secret.IsExceedLimit())

	secret.ViewsCount++
	assert.Equal(t, true, secret.IsExceedLimit())
}

func TestSecret_IsExpired(t *testing.T) {
	now := time.Now()

	secret := entity.Secret{
		Created:     now,
		ExpireAfter: 1,
	}

	assert.Equal(t, true, secret.IsExpired(now.Add(2*time.Minute)))
	assert.Equal(t, false, secret.IsExpired(now))
}
