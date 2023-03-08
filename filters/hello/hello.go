package filters

import (
	"log"

	"golang.org/x/net/context"

	grpc "github.com/Romero027/grpc-go"
)

// Hello is a client-side unary interceptor function that logs a greeting message.
func HelloClient(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.ADNInvoker, opts ...grpc.CallOption) (err error) {
	log.Println("Hello from UnaryClientInterceptor")
	return invoker(ctx, method, req, reply, cc, opts...)
}

// Hello is a server-side unary interceptor function that logs a greeting message.
func HelloServer(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("Hello from UnaryServerInterceptor")
	resp, err = handler(ctx, req)
	return
}
