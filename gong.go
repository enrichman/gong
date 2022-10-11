package gong

import (
	"context"
	"time"
)

// Interval schedules a time.Ticker that will execute the callback at the interval
// specified by the delay. This can be canceled providing an appropriate context.Context.
// Callback is a func that takes no argument. To schedule a callback with arguments
// check the other IntervalX func.
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

// Interval1 is like Interval, but with the callback taking one argument
func Interval1[A any](
	ctx context.Context,
	callback func(A),
	delay time.Duration,
	arg1 A,
) {
	ticker := time.NewTicker(delay)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-done(ctx):
				return
			case <-ticker.C:
				callback(arg1)
			}
		}
	}()
}

// Interval2 is like Interval, but with the callback taking two arguments
func Interval2[A, B any](
	ctx context.Context,
	callback func(A, B),
	delay time.Duration,
	arg1 A,
	arg2 B,
) {
	ticker := time.NewTicker(delay)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-done(ctx):
				return
			case <-ticker.C:
				callback(arg1, arg2)
			}
		}
	}()
}

// Interval3 is like Interval, but with the callback taking three arguments
func Interval3[A, B, C any](
	ctx context.Context,
	callback func(A, B, C),
	delay time.Duration,
	arg1 A,
	arg2 B,
	arg3 C,
) {
	ticker := time.NewTicker(delay)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-done(ctx):
				return
			case <-ticker.C:
				callback(arg1, arg2, arg3)
			}
		}
	}()
}

// Interval4 is like Interval, but with the callback taking four arguments
func Interval4[A, B, C, D any](
	ctx context.Context,
	callback func(A, B, C, D),
	delay time.Duration,
	arg1 A,
	arg2 B,
	arg3 C,
	arg4 D,
) {
	ticker := time.NewTicker(delay)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-done(ctx):
				return
			case <-ticker.C:
				callback(arg1, arg2, arg3, arg4)
			}
		}
	}()
}

// Interval5 is like Interval, but with the callback taking five arguments
func Interval5[A, B, C, D, E any](
	ctx context.Context,
	callback func(A, B, C, D, E),
	delay time.Duration,
	arg1 A,
	arg2 B,
	arg3 C,
	arg4 D,
	arg5 E,
) {
	ticker := time.NewTicker(delay)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-done(ctx):
				return
			case <-ticker.C:
				callback(arg1, arg2, arg3, arg4, arg5)
			}
		}
	}()
}

// Timeout schedules a time.Timer that will execute the callback after the timeout
// specified by the delay. This can be canceled providing an appropriate context.Context.
// Callback is a func that takes no argument. To schedule a callback with arguments
// check the other TimeoutX func.
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

// Timeout1 is like Timeout, but with the callback taking one argument
func Timeout1[A any](
	ctx context.Context,
	callback func(A),
	delay time.Duration,
	arg1 A,
) {
	timer := time.NewTimer(delay)
	go func() {
		defer timer.Stop()
		for {
			select {
			case <-done(ctx):
				return
			case <-timer.C:
				callback(arg1)
			}
		}
	}()
}

// Timeout2 is like Timeout, but with the callback taking two arguments
func Timeout2[A, B any](
	ctx context.Context,
	callback func(A, B),
	delay time.Duration,
	arg1 A,
	arg2 B,
) {
	timer := time.NewTimer(delay)
	go func() {
		defer timer.Stop()
		for {
			select {
			case <-done(ctx):
				return
			case <-timer.C:
				callback(arg1, arg2)
			}
		}
	}()
}

// Timeout3 is like Timeout, but with the callback taking three arguments
func Timeout3[A, B, C any](
	ctx context.Context,
	callback func(A, B, C),
	delay time.Duration,
	arg1 A,
	arg2 B,
	arg3 C,
) {
	timer := time.NewTimer(delay)
	go func() {
		defer timer.Stop()
		for {
			select {
			case <-done(ctx):
				return
			case <-timer.C:
				callback(arg1, arg2, arg3)
			}
		}
	}()
}

// Timeout4 is like Timeout, but with the callback taking four arguments
func Timeout4[A, B, C, D any](
	ctx context.Context,
	callback func(A, B, C, D),
	delay time.Duration,
	arg1 A,
	arg2 B,
	arg3 C,
	arg4 D,
) {
	timer := time.NewTimer(delay)
	go func() {
		defer timer.Stop()
		for {
			select {
			case <-done(ctx):
				return
			case <-timer.C:
				callback(arg1, arg2, arg3, arg4)
			}
		}
	}()
}

// Timeout5 is like Timeout, but with the callback taking five arguments
func Timeout5[A, B, C, D, E any](
	ctx context.Context,
	callback func(A, B, C, D, E),
	delay time.Duration,
	arg1 A,
	arg2 B,
	arg3 C,
	arg4 D,
	arg5 E,
) {
	timer := time.NewTimer(delay)
	go func() {
		defer timer.Stop()
		for {
			select {
			case <-done(ctx):
				return
			case <-timer.C:
				callback(arg1, arg2, arg3, arg4, arg5)
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
