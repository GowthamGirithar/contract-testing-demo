package main

import (
	"fmt"
	"github.com/GowthamGirithar/contract-testing-demo/order-service-consumer/internal/adapter"
	"github.com/GowthamGirithar/contract-testing-demo/order-service-consumer/internal/remote/orderserviceprovider"
	"github.com/GowthamGirithar/contract-testing-demo/order-service-consumer/internal/usecase"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// for integration test pass address port as container name and "4770" which the grpc mock server port
	// TODO
	remoteClient := orderserviceprovider.NewClient("localhost", "9090", "order-service-consumer", time.Duration(5*time.Second))
	service := usecase.NewService(remoteClient)
	ser := adapter.NewServer(service)
	mux := http.NewServeMux()
	mux.HandleFunc("/order", ser.CreateOrder)

	go func() {
		fmt.Println("server started")
		if err := http.ListenAndServe(":8080", mux); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
