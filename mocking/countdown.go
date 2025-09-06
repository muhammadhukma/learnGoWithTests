package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

<<<<<<< HEAD
// DefaultSleeper
=======
// DefaultSleep tai
>>>>>>> refs/remotes/origin/main
type DefaultSleep struct{}

func (d *DefaultSleep) Sleep() {
	time.Sleep(time.Second * 1)
}

<<<<<<< HEAD
// configurabelsleeper
=======
// ConfigurableSleeper 2025-09-06
>>>>>>> refs/remotes/origin/main
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{time.Second * 1, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
