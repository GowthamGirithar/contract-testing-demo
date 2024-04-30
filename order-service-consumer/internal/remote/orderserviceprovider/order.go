package orderserviceprovider

import (
	"context"
	"fmt"
	"github.com/GowthamGirithar/contract-testing-demo/order-service-consumer/internal/domain"
	order "github.com/GowthamGirithar/contract-testing-demo/proto/order-service"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

type Client struct {
	address   string
	port      string
	userAgent string
	timeout   time.Duration
}

func NewClient(address string, port string, userAgent string, timeout time.Duration) Client {
	return Client{
		address:   address,
		port:      port,
		userAgent: userAgent,
		timeout:   timeout,
	}
}

func (c Client) CreateOrder(ctx context.Context, o domain.Order) error {
	dialCtx, dialCancel := context.WithTimeout(ctx, 5*time.Second)
	defer dialCancel()
	fmt.Println(c.address)
	fmt.Println("the port is ", c.port)
	conn, err := grpc.DialContext(dialCtx, fmt.Sprintf("%s:%s", c.address, c.port),
		grpc.WithUserAgent(c.userAgent),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)),
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

	req := mapCreateOrderReqFrom(o)

	ctx = getContextWithMetadata(ctx, c.userAgent)
	contextCancel, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	res, err := order.NewOrderClient(conn).CreateOrder(contextCancel, req)
	if err != nil {
		fmt.Println(err)
		fmt.Println(status.Code(err))
		if status.Code(err) == codes.DeadlineExceeded {
			fmt.Println("hello yes deadline exceeded")
		}
		return err
	}

	fmt.Println("The order number is ", res.GetOrderNumber())

	return nil
}

func (c Client) GetOrder(ctx context.Context, orderCode string) (*order.GetOrderResponse, error) {
	dialCtx, dialCancel := context.WithTimeout(ctx, 1*time.Second)
	defer dialCancel()
	conn, err := grpc.DialContext(dialCtx, fmt.Sprintf("%s:%s", c.address, c.port),
		grpc.WithUserAgent(c.userAgent),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)),
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithMax(3),
				grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100*time.Millisecond)),
				grpc_retry.WithCodes(codes.Unavailable, codes.DeadlineExceeded),
			)),
	)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	ctx = getContextWithMetadata(ctx, c.userAgent)
	contextCancel, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	res, err := order.NewOrderClient(conn).GetOrder(contextCancel, &order.GetOrderRequest{OrderNumber: orderCode})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func mapCreateOrderReqFrom(o domain.Order) *order.CreateOrderRequest {
	return &order.CreateOrderRequest{
		OrderNumber:   o.OrderID,
		CustomerEmail: o.CustomerEmail,
	}
}

func getContextWithMetadata(ctx context.Context, agent string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs("user-agent", agent))
}
