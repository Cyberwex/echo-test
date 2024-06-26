package service

import "time"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d).Hours() / 24

	return int64(dur)
}
