export PACT_DIR = $(PWD)/pacts

run-order-service-consumer:
	@go run order-service-consumer/cmd/main.go

run-order-service-provider:
	@go run order-service-provider/cmd/main.go

consumer: export PACT_TEST := true
provider: export PACT_TEST := true

test-consumer:
	go test github.com/GowthamGirithar/contract-testing-demo/order-service-consumer/tests/contract -run 'TestCreateOrder' -v

test-provider:
	go test github.com/GowthamGirithar/contract-testing-demo/order-service-provider/tests/contract -run 'TestCreateOrder' -v