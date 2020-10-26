package concurrency

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incremnting the counter 3 time leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, &counter, 3)
	})

	t.Run("runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCount(t, &counter, wantedCount)
	})
}

func assertCount(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, wanted %d", got.Value(), want)
	}

}
