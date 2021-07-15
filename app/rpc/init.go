package rpc

import (
	"git.nonolive.co/FE/mildom-video/pkg/app/mildomapi"
	"github.com/douyu/jupiter/pkg/client/grpc"
	"github.com/go-admin-team/go-admin-core/sdk"
	"time"
)

var MildomVideoServiceClient mildomapi.MildomVideoServiceClient

func Init() error {
	client := sdk.Runtime.GetGrpcClients()["mildomVideo"]
	config := grpc.DefaultConfig()
	config.Address = client.Address
	config.Block = client.Block
	config.DialTimeout = time.Duration(client.DialTimeout) * time.Second
	config.BalancerName = client.BalancerName

	MildomVideoServiceClient = mildomapi.NewMildomVideoServiceClient(config.Build())
	return nil
}
