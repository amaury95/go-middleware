package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	transactionv1 "middleware.go/domain/transaction/v1"
	"middleware.go/internal/transaction/v1"
)

func main() {
	// application context
	ctx := context.Background()

	// server handler
	mux := runtime.NewServeMux()

	// services
	transactionServer := transaction.NewServer()
	transactionv1.RegisterTransactionServiceHandlerServer(ctx, mux, transactionServer)

	// run server on port 8000
	print("Running middleware on port http://localhost:8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		panic(err)
	}
}
