package schedule

import (
	"time"

	"github.com/robfig/cron/v3"
)

const defaultLookback = 2 * time.Minute

type Schedule struct {
	cron     cron.Schedule
	lookback time.Duration
	location *time.Location
	enabled  bool
}

// NewSchedule creates a new Schedule from a cron spec string.
// If spec is empty, the schedule is always allowed (disabled).
func NewSchedule(spec string) (*Schedule, error) {
	if spec == "" {
		return &Schedule{
			enabled: false, // always allowed
		}, nil
	}

	parser := cron.NewParser(
		cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow,
	)

	sched, err := parser.Parse(spec)
	if err != nil {
		return nil, err
	}

	return &Schedule{
		cron:     sched,
		lookback: defaultLookback,
		location: time.Local,
		enabled:  true,
	}, nil
}

// CanRun returns true if the current time is within the allowed schedule.
func (s *Schedule) CanRun() bool {
	if !s.enabled {
		return true
	}

	now := time.Now().In(s.location)
	start := now.Add(-s.lookback)

	next := s.cron.Next(start)

	return !next.After(now)
}
