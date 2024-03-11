package grpc

import (
	"context"
	"fmt"
	order "github.com/GowthamGirithar/contract-testing-demo/order-service-provider/internal/adapter/grpc"
	pb "github.com/GowthamGirithar/contract-testing-demo/proto/order-service"
	gmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	gctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"net"
)

func RunServer(ctx context.Context, port string) error {
	server := grpc.NewServer(grpc.UnaryInterceptor(
		gmiddleware.ChainUnaryServer(
			gctxtags.UnaryServerInterceptor(gctxtags.WithFieldExtractor(gctxtags.CodeGenRequestFieldExtractor)),
		),
	))
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("error in listening")
	}

	orderServer := order.NewServer()
	pb.RegisterOrderServer(server, orderServer)
	return server.Serve(listen)
}
