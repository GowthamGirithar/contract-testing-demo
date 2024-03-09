package orderserviceprovider

import (
	"context"
	"fmt"
	order "github.com/GowthamGirithar/contract-testing-demo/proto/order-service"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"time"
)

type client struct {
	address   string
	port      string
	userAgent string
	timeout   time.Duration
}

func NewClient(address string, port string, userAgent string, timeout time.Duration) client {
	return client{
		address:   address,
		port:      port,
		userAgent: userAgent,
		timeout:   timeout,
	}
}

func (c client) CreateOrder(ctx context.Context, req interface{}) error {
	dialCtx, dialCancel := context.WithTimeout(ctx, 1*time.Second)
	defer dialCancel()
	conn, err := grpc.DialContext(dialCtx, fmt.Sprintf("%s:%s", c.address, c.port),
		grpc.WithUserAgent(c.userAgent),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)), - check for pact test
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithMax(3),
				grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100*time.Millisecond)),
				grpc_retry.WithCodes(codes.Unavailable, codes.DeadlineExceeded),
			)),
	)
	if err != nil {
		return err
	}
	defer conn.Close()

	if req == nil {
		req = &order.CreateOrderRequest{
			OrderNumber:   "12",
			CustomerEmail: "test@gmail.com",
		}
	}

	ctx = getContextWithMetadata(ctx, c.userAgent)
	contextCancel, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	res, err := order.NewOrderClient(conn).CreateOrder(contextCancel, req.(*order.CreateOrderRequest))
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}

func getContextWithMetadata(ctx context.Context, agent string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs("user-agent", agent))
}
