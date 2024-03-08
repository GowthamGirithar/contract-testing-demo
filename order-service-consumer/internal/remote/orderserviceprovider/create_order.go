package orderserviceprovider

import (
	"context"
	"fmt"
	order "github.com/GowthamGirithar/contract-testing-demo/proto/order-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	address string
	port    string
}

func NewClient(address string, port string) client {
	return client{
		address: address,
		port:    port,
	}
}

func (c client) CreateOrder(ctx context.Context) error {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", c.address, c.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return err
	}
	req := &order.CreateOrderRequest{
		OrderNumber:   "12",
		CustomerEmail: "test@gmail.com",
	}
	res, err := order.NewOrderClient(conn).CreateOrder(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
