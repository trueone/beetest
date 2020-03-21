package dto

// Secret data transfer object
type SecretRequest struct {
	Secret           string `json:"secret"`
	ExpireAfterViews int    `json:"expireAfterViews"`
	ExpireAfter      int    `json:"expireAfter"`
}
