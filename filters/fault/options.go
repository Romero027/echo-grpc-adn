package fault

import (
	"time"

	grpc "github.com/Romero027/grpc-go"
)

var (
	defaultOptions = &options{
		delay:            0,  // disabled
		delayProbability: -1, // disabled
		delayedCount:     0,
		abortProbability: -1, // disabled
		abortedCount:     0,
	}
)

func WithDelay(delay time.Duration) CallOption {
	if delay < 0 {
		panic("delay must be non-negative")
	}

	return CallOption{applyFunc: func(o *options) {
		o.delay = delay
	}}
}

func WithDelayProbability(probability float64) CallOption {
	if probability < 0 || probability > 1 {
		panic("probability must be between 0 and 1")
	}
	return CallOption{applyFunc: func(o *options) {
		o.delayProbability = probability
	}}
}

func WithAbortProbability(probability float64) CallOption {
	if probability < 0 || probability > 1 {
		panic("probability must be between 0 and 1")
	}
	return CallOption{applyFunc: func(o *options) {
		o.abortProbability = probability
	}}
}

type options struct {
	abortProbability float64
	delayProbability float64
	delay            time.Duration
	abortedCount     int32
	delayedCount     int32
}

type CallOption struct {
	grpc.EmptyCallOption // make sure we implement private after() and before() fields so we don't panic.
	applyFunc            func(opt *options)
}

func reuseOrNewWithCallOptions(opt *options, callOptions []CallOption) *options {
	if len(callOptions) == 0 {
		return opt
	}
	optCopy := &options{}
	*optCopy = *opt
	for _, f := range callOptions {
		f.applyFunc(optCopy)
	}
	return optCopy
}

func filterCallOptions(callOptions []grpc.CallOption) (grpcOptions []grpc.CallOption, faultOptions []CallOption) {
	for _, opt := range callOptions {
		if co, ok := opt.(CallOption); ok {
			faultOptions = append(faultOptions, co)
		} else {
			grpcOptions = append(grpcOptions, opt)
		}
	}
	return grpcOptions, faultOptions
}
