package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

//
//func TestCountdown(t *testing.T) {
//	buffer := &bytes.Buffer{}
//	Countdown(buffer, )
//	want := "3"
//	got := buffer.String()
//	if got != want {
//		t.Errorf("got %q want %q",got,want)
//	}
//}

func TestCountdown1(t *testing.T) {
	t.Run("wihout order checking", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		SpySleeper := &SpySleeper{}
		Countdown(buffer, SpySleeper)
		want := `3
2
1
Go!`

		got := buffer.String()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if SpySleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", SpySleeper.Calls)
		}

	})

	t.Run("with order checking", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}

	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime)
	}
}
