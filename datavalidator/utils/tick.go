package utils

import (
	"context"
	"time"
)

func Ticker(ctx context.Context, fn func(ctx context.Context), duration time.Duration) {
	tick := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-tick.C:
				fn(ctx)
			case <-ctx.Done():
				return
			}
		}
	}()
}
