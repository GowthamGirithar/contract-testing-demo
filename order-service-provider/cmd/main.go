package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	order "github.com/GowthamGirithar/contract-testing-demo/order-service-provider/internal/adapter/grpc"

	pb "github.com/GowthamGirithar/contract-testing-demo/proto/order-service"
)

func main() {
	server := grpc.NewServer()
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("error in listening")
	}

	orderServer := order.NewServer()
	pb.RegisterOrderServer(server, orderServer)
	server.Serve(listen)
}
