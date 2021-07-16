package grpc

import (
	"github.com/douyu/jupiter/pkg/client/grpc/resolver"
	"github.com/douyu/jupiter/pkg/registry/etcdv3"
	"github.com/go-admin-team/go-admin-core/sdk"
	toolsConfig "github.com/go-admin-team/go-admin-core/sdk/config"
	"time"
)

// Setup 配置Grpc
func Setup() {
	grpcConfig := toolsConfig.GrpcConfig
	// 注册统一的etcd服务源
	config := etcdv3.DefaultConfig()
	config.ConnectTimeout = time.Duration(grpcConfig.Registry.ConnectTimeout) * time.Second
	config.Endpoints = grpcConfig.Registry.Endpoints
	config.Secure = grpcConfig.Registry.Secure
	//config.Prefix = grpcConfig.Registry.Prefix
	resolver.Register("etcd", config.Build())

	// 植入grpc client配置
	sdk.Runtime.SetGrpcClients(grpcConfig.Clients)

}
