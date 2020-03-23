package data

import (
	"time"

	"github.com/trueone/beetest/internal/app/entity"
)

// Secret response data transfer object
type SecretResponse struct {
	Hash           string    `json:"hash"`
	SecretText     string    `json:"secretText"`
	CreatedAt      time.Time `json:"createdAt"`
	ExpiresAt      time.Time `json:"expiresAt"`
	RemainingViews int       `json:"remainingViews"`
}

// Deserialize data from Secret entity
func (sr *SecretResponse) Deserialize(secret entity.Secret) {
	sr.Hash = secret.Hash
	sr.SecretText = secret.Text
	sr.CreatedAt = secret.Created
	if secret.ExpireAfter != 0 {
		sr.ExpiresAt = secret.Created.Add(time.Minute * time.Duration(secret.ExpireAfter))
	}
	sr.RemainingViews = secret.RemainingViews
}
