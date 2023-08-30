package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/quocbang/go-grpc/client/config"
	"github.com/quocbang/go-grpc/pkg/protobuf"
	"github.com/quocbang/go-grpc/pkg/protobuf/account"
)

func init() {
	flag.StringVar(&config.C.GrpcServerHost, "grpc-server-host", "", "grpc server host")
	flag.IntVar(&config.C.GrpcServerPort, "grpc-server-port", 0, "grpc server port")
	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to initialize zap, error: %v", err)
	}
	zap.ReplaceGlobals(logger)
	zap.RedirectStdLog(logger)
}

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.C.GrpcServerHost, config.C.GrpcServerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect, error: %v", err)
	}
	defer conn.Close()

	client := protobuf.NewProductClient(conn)

	resp, err := client.Login(context.Background(), &account.LoginRequest{
		UserID:   "quocbang",
		Password: "quocbang",
	})
	if err != nil {
		log.Printf("login failed, error: %v", err)
	}

	log.Println(resp)
}
