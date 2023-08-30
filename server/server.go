package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/quocbang/go-grpc/middleware"
	"github.com/quocbang/go-grpc/middleware/logger"
	"github.com/quocbang/go-grpc/pkg/config"
	"github.com/quocbang/go-grpc/pkg/protobuf"
	"github.com/quocbang/go-grpc/pkg/protobuf/account"
	a "github.com/quocbang/go-grpc/server/account"
)

type server struct {
	protobuf.UnimplementedProductServer
}

func Run() {
	flag.BoolVar(&config.C.IsDev, "dev-mode", false, "is development")
	flag.StringVar(&config.C.GrpcHost, "grpc-host", "localhost", "grpc server host")
	flag.IntVar(&config.C.GrpcPort, "grpc-port", 8081, "grpc server port")
	flag.Parse()

	// init logger.
	logger.InitLogger(config.C.IsDev)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.C.GrpcHost, config.C.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen tcp, error: %v", err)
	}
	log.Printf("start listen grpc server %s:%d", config.C.GrpcHost, config.C.GrpcPort)

	// init logging middleware
	grpcLogger := grpc.UnaryInterceptor(middleware.Logging)

	s := grpc.NewServer(grpcLogger)
	sv := server{}
	protobuf.RegisterProductServer(s, &sv)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve, error: %v", err)
	}
}

func (s *server) Login(ctx context.Context, req *account.LoginRequest) (*account.LoginReply, error) {
	return a.Login(ctx, req)
}
