package main

import (
	"context"

	"github.com/GowthamGirithar/contract-testing-demo/order-service-consumer/internal/remote/orderserviceprovider"
)

func main() {
	s := orderserviceprovider.NewClient("127.0.0.1", "8082")
	s.CreateOrder(context.Background())
}
