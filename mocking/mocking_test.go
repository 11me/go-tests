package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleepPrinter := &CountdownOperationSpy{}

	Countdown(buffer, spySleeperPrinter)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4, got %d", spySleeper.Calls)
	}

	t.Run("sleep before every writing", func(t *testing.T) {

		spySleepPrinter := &CountdownOperationSpy{}
		buffer := &bytes.Buffer{}

		Countdown(buffer, spySleepPrinter)

		want := []string{
			SLEEP,
			WRITE,
			SLEEP,
			WRITE,
			SLEEP,
			WRITE,
			SLEEP,
			WRITE,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}

	})
}
