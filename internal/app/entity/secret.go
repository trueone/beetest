package entity

import (
	"time"
)

type Secret struct {
	Hash           string
	Text           string
	Created        time.Time
	Expires        time.Time
	RemainingViews int
	ViewsCount     int
}

func (s *Secret) IsExceedLimit() bool {
	return s.ViewsCount > s.RemainingViews
}

func (s *Secret) IsExpired(now time.Time) bool {
	return now.After(s.Expires)
}
