package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		Countdown(buffer, spySleeper)
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
	t.Run("latihan cobain sendiri heeh", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleepTime := 5 * time.Second
		spyTime := &SpyTime{}
		sleeper := &ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}
		Countdown(buffer, sleeper)

		if spyTime.durationSlept != sleepTime {
			t.Errorf("shoud have sleep for %v but sleep for %v", sleepTime, spyTime.durationSlept)
		}
	})
}

// spy sleeper
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// spy countdown operations
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, er error) {
	s.Calls = append(s.Calls, write)
	return n, er
}

// configurable sleeper
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

const (
	sleep = "sleep"
	write = "write"
)
