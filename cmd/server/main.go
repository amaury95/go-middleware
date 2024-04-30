package main

import (
	"context"
	"net/http"

	transactionv1 "github.com/amaury95/go-middleware/domain/transaction/v1"
	"github.com/amaury95/go-middleware/internal/transaction/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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
