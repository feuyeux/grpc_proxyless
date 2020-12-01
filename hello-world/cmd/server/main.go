package main

import (
	"github.com/feuyeux/grpc_proxyless/common/pkg/config"
	"github.com/feuyeux/grpc_proxyless/common/pkg/logger"
	"github.com/feuyeux/grpc_proxyless/hello-world/internal/app/server"
	"go.uber.org/zap"
)

func main() {

	config, err := config.ReadConfig()
	if err != nil {
		logger.Logger.Fatal("Unable to read config", zap.Error(err))
	}

	server.StartServer(config.GetInt("port"))
}
