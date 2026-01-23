package schedule

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew_EmptySpec(t *testing.T) {
	sched, err := NewSchedule("")

	require.NoError(t, err)
	require.NotNil(t, sched)

	assert.False(t, sched.enabled)
	assert.Nil(t, sched.cron)
}

func TestNew_ValidCronSpec_EveryMinute(t *testing.T) {
	sched, err := NewSchedule("* * * * *")

	require.NoError(t, err)
	require.NotNil(t, sched)

	assert.True(t, sched.enabled)
	assert.NotNil(t, sched.cron)
	assert.Equal(t, defaultLookback, sched.lookback)
	assert.Equal(t, time.Local, sched.location)
}

func TestNew_ValidCronSpec_EveryHour(t *testing.T) {
	sched, err := NewSchedule("0 * * * *")

	require.NoError(t, err)
	require.NotNil(t, sched)

	assert.True(t, sched.enabled)
	assert.NotNil(t, sched.cron)
}

func TestNew_InvalidCronSpec(t *testing.T) {
	sched, err := NewSchedule("invalid cron spec")

	assert.Error(t, err)
	assert.Nil(t, sched)
}

func TestNew_DefaultValues(t *testing.T) {
	sched, err := NewSchedule("* * * * *")

	require.NoError(t, err)
	require.NotNil(t, sched)

	assert.Equal(t, 2*time.Minute, sched.lookback)
	assert.Equal(t, time.Local, sched.location)
}

func TestSchedule_CanRun_Disabled(t *testing.T) {
	sched, err := NewSchedule("")
	require.NoError(t, err)

	// When disabled, should always return true
	assert.True(t, sched.CanRun())
}

func TestSchedule_CanRun_EveryMinute(t *testing.T) {
	sched, err := NewSchedule("* * * * *")
	require.NoError(t, err)

	// With "* * * * *" cron (every minute), CanRun should always be true
	// because there's always a run within the last 2 minutes
	assert.True(t, sched.CanRun())
}

func TestSchedule_CanRun_EveryHourAtZero(t *testing.T) {
	sched, err := NewSchedule("0 * * * *")
	require.NoError(t, err)

	// Should return true only if we're within 2 minutes of the hour mark
	now := time.Now()
	minute := now.Minute()

	canRun := sched.CanRun()

	// Should be able to run if we're within 2 minutes after the hour
	if minute <= 2 {
		assert.True(t, canRun)
	}
	// Otherwise it depends on when the last hour boundary was
}

func TestSchedule_CanRun_Daily(t *testing.T) {
	// Runs at 3:00 AM every day
	sched, err := NewSchedule("0 3 * * *")
	require.NoError(t, err)

	now := time.Now()
	canRun := sched.CanRun()

	// Should only return true if it's between 3:00 AM and 3:02 AM
	if now.Hour() == 3 && now.Minute() <= 2 {
		assert.True(t, canRun)
	} else {
		fmt.Println("canRun:", canRun)
		assert.False(t, canRun)
	}
	// Otherwise depends on current time
}

func TestSchedule_WorkDay(t *testing.T) {
	// Every minute of the workday (9am-5pm, Mon-Fri)
	sched, err := NewSchedule("* 9-16 * * 1-5")
	require.NoError(t, err)

	now := time.Now()
	canRun := sched.CanRun()

	// Should return false after 5pm local time
	if now.Hour() >= 17 || now.Hour() < 9 || now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		assert.False(t, canRun)
	} else {
		// During work hours, depends on minute
		fmt.Println("canRun:", canRun)
	}
}
