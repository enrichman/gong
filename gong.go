package gong

import (
	"context"
	"time"
)

// Interval schedules a time.Ticker that will execute the callback at the interval
// specified by the delay. This can be canceled providing an appropriate context.Context.
func Interval(
	ctx context.Context,
	callback func(),
	delay time.Duration,
) {
	ticker := time.NewTicker(delay)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-done(ctx):
				return
			case <-ticker.C:
				callback()
			}
		}
	}()
}

// Timeout schedules a time.Timer that will execute the callback after the timeout
// specified by the delay. This can be canceled providing an appropriate context.Context.
func Timeout(
	ctx context.Context,
	callback func(),
	delay time.Duration,
) {
	timer := time.NewTimer(delay)
	go func() {
		for {
			select {
			case <-done(ctx):
				timer.Stop()
				return
			case <-timer.C:
				callback()
			}
		}
	}()
}

func done(ctx context.Context) <-chan struct{} {
	if ctx == nil {
		ctx = context.Background()
	}
	return ctx.Done()
}
