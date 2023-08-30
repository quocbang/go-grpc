package config

var C Config

type Config struct {
	IsDev    bool
	GrpcHost string
	GrpcPort int
}
