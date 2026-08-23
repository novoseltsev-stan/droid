package chanutils

import "context"

// FanOut split channel to many.
//
// FanOut manage output channels.
// If ctx is canceled or input channel closed, returned channels also is closed and process is stopped.
func FanOut[T any](ctx context.Context, ch <-chan T, nums int) []chan T {
	chs := make([]chan T, nums)
	for i := range nums {
		chs[i] = make(chan T)
	}

	for _, outCh := range chs {
		go func() {
			defer close(outCh)

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
					case outCh <- v:
					}
				}
			}
		}()
	}

	return chs
}
