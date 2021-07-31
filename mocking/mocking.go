package main

import (
	"io"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (spy *SpySleeper) Sleep() {
	spy.Calls++
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

const (
	SLEEP = "sleep"
	WRITE = "write"
)

type CountdownOperationSpy struct {
	Calls []string
}

func (spy *CountdownOperationSpy) Sleep() {
	spy.Calls = append(spy.Calls, SLEEP)
}

func (spy *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	spy.Calls = append(spy.Calls, WRITE)
	return
}

func Countdown(writer io.Writer, sleeper CountdownOperationSpy) {

	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		sleeper.Write(writer)
	}

	sleeper.Sleep()
	sleeper.Write(writer)
}

func main() {

	defaultSleeper := &DefaultSleeper{}
	Countdown(os.Stdout, defaultSleeper)

}
