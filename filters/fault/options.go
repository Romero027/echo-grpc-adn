package fault

import (
	"time"

	grpc "github.com/Romero027/grpc-go"
)

func WithDelay(delay time.Duration) CallOption {
	return CallOption{applyFunc: func(o *options) {
		o.delay = delay
	}}
}

func WithProbability(probability float64) CallOption {
	return CallOption{applyFunc: func(o *options) {
		o.probability = probability
	}}
}

type options struct {
	probability float64
	delay       time.Duration
}

type CallOption struct {
	grpc.EmptyCallOption // make sure we implement private after() and before() fields so we don't panic.
	applyFunc            func(opt *options)
}
