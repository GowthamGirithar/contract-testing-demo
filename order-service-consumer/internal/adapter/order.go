package adapter

import (
	"fmt"
	"github.com/GowthamGirithar/contract-testing-demo/order-service-consumer/internal/usecase"
	"net/http"
)

type Server struct {
	service usecase.Service
}

func NewServer(service *usecase.Service) *Server {
	return &Server{service: *service}
}

func (s *Server) CreateOrder(res http.ResponseWriter, req *http.Request) {
	err := s.service.CreateOrder(req.Context(), "123", "test@gmail.com")
	if err != nil {
		res.WriteHeader(500)
		fmt.Fprintf(res, "Internal Error")
		return
	}
	res.WriteHeader(201)
	fmt.Fprintf(res, "Success")
}

func (s Server) GetOrder(res http.ResponseWriter, req *http.Request) {
	_, err := s.service.GetOrder(req.Context(), "12")
	if err != nil {
		fmt.Println(err)
	}
}
