package usecase

import (
	"context"
	"github.com/GowthamGirithar/contract-testing-demo/order-service-consumer/internal/domain"
	"github.com/GowthamGirithar/contract-testing-demo/order-service-consumer/internal/remote/orderserviceprovider"
)

type Service struct {
	client orderserviceprovider.Client
}

func NewService(client orderserviceprovider.Client) *Service {
	return &Service{
		client: client,
	}
}

func (s Service) CreateOrder(ctx context.Context, orderID string, email string) error {
	err := s.client.CreateOrder(ctx, domain.Order{
		OrderID:       orderID,
		CustomerEmail: email,
	})
	return err
}

func (s Service) GetOrder(ctx context.Context, orderID string) (*domain.Order, error) {
	res, err := s.client.GetOrder(ctx, orderID)
	order := &domain.Order{
		OrderID:       res.GetOrderNumber(),
		CustomerEmail: res.GetCustomerEmail(),
	}
	return order, err
}
