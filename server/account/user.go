package account

import (
	"context"

	"github.com/quocbang/go-grpc/pkg/protobuf/account"
)

func Login(ctx context.Context, req *account.LoginRequest) (*account.LoginReply, error) {
	return nil, nil
}
