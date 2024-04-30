FROM golang:1.22.0 as builder

WORKDIR /app

COPY . ./

RUN go mod tidy && go build -o /order-service-provider ./cmd/main.go

CMD ["/order-service-provider"]

