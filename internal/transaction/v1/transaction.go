package transaction

import (
	"context"

	commonv1 "middleware.go/domain/common/v1"
	transactionv1 "middleware.go/domain/transaction/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	transactionv1.UnsafeTransactionServiceServer
}

func NewServer() *server {
	return &server{}
}

func (*server) Ping(context.Context, *emptypb.Empty) (*commonv1.BoolResponse, error) {
	return &commonv1.BoolResponse{Success: true}, nil
}
func (*server) Submit(context.Context, *transactionv1.SubmitRequest) (*transactionv1.SubmitResponse, error) {
	panic("implementation missing")
}

var (
	_ transactionv1.TransactionServiceServer = &server{}
)
