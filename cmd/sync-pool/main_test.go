package main

import (
	"context"
	"runtime/debug"
	"sync"
	"testing"
	"time"
)

const testsTimeout = time.Second * 10

func TestNoPool(t *testing.T) {
	const workersCount = 100

	const data = `{"count": 42, "text": "some-text"}`

	start := make(chan struct{})
	var ctx context.Context
	var cancelCtx context.CancelFunc
	wg := &sync.WaitGroup{}

	wr := &Writer{}

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				NoPool(wr, []byte(data))
			}
		}()
	}

	ctx, cancelCtx = context.WithTimeout(context.Background(), testsTimeout)
	defer cancelCtx()

	close(start)

	wg.Wait()

	var gcStats debug.GCStats
	debug.ReadGCStats(&gcStats)

	t.Logf("GC pause total: %d ms", gcStats.PauseTotal/time.Millisecond)
}

func TestWithPool(t *testing.T) {
	const workersCount = 100

	const data = `{"count": 42, "text": "some-text"}`

	start := make(chan struct{})
	var ctx context.Context
	var cancelCtx context.CancelFunc
	wg := &sync.WaitGroup{}

	wr := &Writer{}

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				WithPool(wr, []byte(data))
			}
		}()
	}

	ctx, cancelCtx = context.WithTimeout(context.Background(), testsTimeout)
	defer cancelCtx()

	close(start)

	wg.Wait()

	var gcStats debug.GCStats
	debug.ReadGCStats(&gcStats)

	t.Logf("GC pause total: %d ms", gcStats.PauseTotal/time.Millisecond)
}

type Writer struct {
}

func (w *Writer) Write(b []byte) (int, error) {
	return len(b), nil
}
