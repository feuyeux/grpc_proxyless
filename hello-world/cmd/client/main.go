package main

import (
	"fmt"
	"github.com/feuyeux/grpc_proxyless/common/pkg/config"
	"github.com/feuyeux/grpc_proxyless/common/pkg/logger"
	"github.com/feuyeux/grpc_proxyless/hello-world/internal/app/client"

	"google.golang.org/grpc"

	"go.uber.org/zap"
)

const ()

var (
	conn *grpc.ClientConn
)

func main() {

	config, err := config.ReadConfig()
	if err != nil {
		logger.Logger.Fatal("Unable to read config", zap.Error(err))
	}

	client.StartClient(fmt.Sprintf("%s", config.GetString("hello.host")))
}
