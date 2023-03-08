package fault

import (
	"math/rand"
	"time"

	"golang.org/x/net/context"

	grpc "github.com/Romero027/grpc-go"
)

func FaultInjectionDelay(opts ...grpc.CallOption) grpc.ADNClientProcessor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.ADNInvoker, opts ...grpc.CallOption) (err error) {

		rand.Seed(time.Now().UnixNano())
		p := rand.Float64()
		if p >= opts.probability {
			time.Sleep(opts.delay)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
