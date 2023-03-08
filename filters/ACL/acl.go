package acl

import (
	codes "github.com/Romero027/grpc-go/codes"
	status "github.com/Romero027/grpc-go/status"
	"golang.org/x/net/context"

	echo "github.com/Romero027/echo-grpc-adn/pb"
	grpc "github.com/Romero027/grpc-go"
)

// ContentBasedACL is a server-side unary interceptor function that filters requests based on their content and blocks requests with a specific body.
func ContentBasedACL(content string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if m, ok := req.(*echo.Msg); ok {
			if m.GetBody() == content {
				return nil, status.Error(codes.InvalidArgument, "request blocked by ACL filter.")
			}
		}
		return handler(ctx, req)
	}
}

// // ContentBasedACL is a server-side unary interceptor function that filters requests based on their content and blocks requests with a specific body.
// func ContentBasedACL(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

// 	if m, ok := req.(*echo.Msg); ok {
// 		if m.GetBody() == "/test" {
// 			return nil, status.Error(codes.InvalidArgument, "request blocked by ACL filter.")
// 		}
// 	}

// 	resp, err = handler(ctx, req)

// 	return
// }
