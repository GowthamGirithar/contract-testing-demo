package grpc

import (
	"context"

	order "github.com/GowthamGirithar/contract-testing-demo/proto/order-service"
)

type server struct {
	order.UnimplementedOrderServer
}

func NewServer() server {
	return server{}
}

func (s server) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	return &order.CreateOrderResponse{
		OrderNumber: req.GetOrderNumber(),
	}, nil
}

func (s server) GetOrder(ctx context.Context, req *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	return nil, nil
}
