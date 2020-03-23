package entity

import (
	"time"
)

type Secret struct {
	Hash           string
	Text           string
	Created        time.Time
	ExpireAfter    int
	RemainingViews int
	ViewsCount     int
}

func (s *Secret) IsExceedLimit() bool {
	return s.ViewsCount > s.RemainingViews
}

func (s *Secret) IsExpired(now time.Time) bool {
	if s.ExpireAfter == 0 {
		return false
	}

	expire := s.Created.Add(time.Minute * time.Duration(s.ExpireAfter))
	return now.After(expire)
}
