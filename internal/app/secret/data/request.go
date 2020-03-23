package data

// Secret request data transfer object
type SecretRequest struct {
	Secret           string `json:"secret"`
	ExpireAfterViews int    `json:"expireAfterViews"`
	ExpireAfter      int    `json:"expireAfter"`
}

// Hash request data transfer object
type HashRequest string
