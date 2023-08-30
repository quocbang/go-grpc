package assets

// generate proto file using gRPC framework
//go:generate protoc -I=. --go_out=paths=source_relative:../pkg protobuf/account/user.proto

//go:generate protoc -I=. --go_out=paths=source_relative:../pkg --go-grpc_out=paths=source_relative:../pkg protobuf/services.proto
