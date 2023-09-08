package mocks

import "time"

type SleeperMock struct {
	TimeSlept time.Duration
}

func (s *SleeperMock) Sleep(d time.Duration) {
	s.TimeSlept = d
}
