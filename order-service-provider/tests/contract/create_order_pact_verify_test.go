package contract_test

import (
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	pb "github.com/GowthamGirithar/contract-testing-demo/proto/order-service"

	l "github.com/pact-foundation/pact-go/v2/log"
	"github.com/pact-foundation/pact-go/v2/provider"
	"google.golang.org/grpc"

	order "github.com/GowthamGirithar/contract-testing-demo/order-service-provider/internal/adapter/grpc"
)

func TestPactProvider(t *testing.T) {
	go startProvider()
	l.SetLogLevel("INFO")

	verifier := provider.NewVerifier()
	// Verify the Provider - From file
	err := verifier.VerifyProvider(t, provider.VerifyRequest{
		Provider:        "create-order-provider",
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", 8222),
		Transports: []provider.Transport{
			{
				Protocol: "grpc",
				Port:     8222,
			},
		},
		ProviderBranch:     os.Getenv("VERSION_BRANCH"),
		FailIfNoPactsFound: false,
		PactFiles:          []string{""}, // TODO: MENTION PACT PATH
		ProviderVersion:    os.Getenv("VERSION_COMMIT"),
	})

	if err != nil {
		fmt.Println(err)
	}
}

// Starts the provider API with hooks for provider states.
// This essentially mirrors the main.go file, with extra routes added.
func startProvider() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8222))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterOrderServer(grpcServer, order.NewServer())
	grpcServer.Serve(lis)
}
