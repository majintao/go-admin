package rpc

import (
	"git.nonolive.co/FE/mildom-video/pkg/cms/mildomapi"
	"github.com/douyu/jupiter/pkg/client/grpc"
	"github.com/go-admin-team/go-admin-core/sdk"
	"time"
)

var MildomVideoCmsServiceClient mildomapi.MildomVideoCmsServiceClient

func Init() error {
	client := sdk.Runtime.GetGrpcClients()["mildomVideo"]
	config := grpc.DefaultConfig()
	config.Address = client.Address
	config.Block = client.Block
	config.DialTimeout = time.Duration(client.DialTimeout) * time.Second
	config.BalancerName = client.BalancerName

	MildomVideoCmsServiceClient = mildomapi.NewMildomVideoCmsServiceClient(config.Build())

	return nil
}
