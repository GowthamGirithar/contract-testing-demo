package main

import (
	"context"
	"fmt"
	"github.com/GowthamGirithar/contract-testing-demo/order-service-consumer/internal/domain"
	"time"

	"github.com/GowthamGirithar/contract-testing-demo/order-service-consumer/internal/remote/orderserviceprovider"
)

func main() {
	s := orderserviceprovider.NewClient("127.0.0.1", "8082", "order-service-consumer", time.Duration(1*time.Second))
	err := s.CreateOrder(context.Background(), domain.Order{
		CustomerEmail: "test@gmail.com",
	})
	fmt.Println(err)
}
