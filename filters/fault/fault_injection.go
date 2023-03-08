package fault

import (
	"math/rand"
	"time"

	codes "github.com/Romero027/grpc-go/codes"
	status "github.com/Romero027/grpc-go/status"
	"golang.org/x/net/context"

	grpc "github.com/Romero027/grpc-go"
)

// FaultInjectionDelay returns a gRPC client interceptor that injects faults into outgoing requests.
// optFuncs is a variable-length argument list of CallOption functions to customize the interceptor's behavior.
func FaultInjectionDelay(optFuncs ...CallOption) grpc.ADNClientProcessor {
	// Combine the default options with any user-provided options.
	intOpts := reuseOrNewWithCallOptions(defaultOptions, optFuncs)

	// Return a function that handles the outgoing request.
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.ADNInvoker, opts ...grpc.CallOption) (err error) {
		// Split the gRPC CallOptions into two sets: gRPC-specific options and fault injection options.
		grpcOpts, faultOptions := filterCallOptions(opts)

		// Combine the internal options with the fault injection options.
		callOpts := reuseOrNewWithCallOptions(intOpts, faultOptions)

		// Generate a random float between 0 and 1.
		rand.Seed(time.Now().UnixNano())
		p := rand.Float64()

		// If the random float is less than the delay probability, sleep for the specified delay duration.
		if p <= callOpts.delayProbability {
			time.Sleep(callOpts.delay)
			callOpts.delayedCount += 1
		}

		// If the random float is less than the abort probability, return an error with the Aborted code.
		if p <= callOpts.abortProbability {
			callOpts.abortedCount += 1
			return status.Error(codes.Aborted, "request aborted by fault injection filter.")
		}

		// Otherwise, invoke the original gRPC method with the specified options.
		return invoker(ctx, method, req, reply, cc, grpcOpts...)
	}
}
