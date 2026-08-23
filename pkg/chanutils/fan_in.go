package chanutils

import (
	"context"
	"sync"
)

// FanIn merges channels into one.
//
// FanIn manage output channel.
// If ctx is canceled, returned channel is closed and process is stopped.
//
// Returns channel with buffer of len(chs) that merges all channels data.
func FanIn[T any](ctx context.Context, chs ...<-chan T) chan T {
	resCh := make(chan T, len(chs))
	var wg sync.WaitGroup

	for _, ch := range chs {
		wg.Go(func() {
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-ch:
					if !ok {
						return
					}

					select {
					case <-ctx.Done():
						return
					case resCh <- v:
					}
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	return resCh
}
