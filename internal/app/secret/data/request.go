package data

// Secret request data transfer object
type SecretRequest struct {
	Secret           string `json:"secret" validate:"required"`
	ExpireAfterViews int    `json:"expireAfterViews" validate:"required"`
	ExpireAfter      int    `json:"expireAfter" validate:"required"`
}

// Hash request data transfer object
type HashRequest string
