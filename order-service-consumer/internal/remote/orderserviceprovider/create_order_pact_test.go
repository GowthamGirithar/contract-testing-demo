package orderserviceprovider

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	message "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	mockProvider, err := message.NewSynchronousPact(message.Config{
		Consumer: "create-order-consumer",
		Provider: "create-order-provider",
	})

	assert.NoError(t, err)
	dir := `https://raw.githubusercontent.com/GowthamGirithar/contract-testing-demo/main/proto/order-service/order.proto`

	grpcInteraction := `{
		"pact:proto": "` + filepath.ToSlash(dir) + `",
		"pact:proto-service": "Order/CreateOrder",
		"pact:content-type": "application/protobuf",
		"request": {
			"order_number" :   "matching(number, 3)",
			"customer_email":  "notEmpty('test@gmail.com')"
		},
		"response": {
			"order_number": "matching(number, 4)"
		}
	}`

	err = mockProvider.
		AddSynchronousMessage("create order request").
		UsingPlugin(message.PluginConfig{
			Plugin: "protobuf",
		}).
		WithContents(grpcInteraction, "application/grpc").
		// Start the gRPC mock server
		StartTransport("grpc", "127.0.0.1", nil).
		// Execute the test
		ExecuteTest(t, func(transport message.TransportConfig, m message.SynchronousMessage) error {
			err := NewClient(transport.Address, fmt.Sprintf("%d", transport.Port)).CreateOrder(context.Background())
			fmt.Println(err)
			// Assert: check the result
			assert.NoError(t, err)
			return err
		})

	assert.NoError(t, err)
}
