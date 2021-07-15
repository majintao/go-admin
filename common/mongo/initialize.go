package grpc

import (
	"github.com/douyu/jupiter/pkg/store/mongox"
	"github.com/go-admin-team/go-admin-core/sdk"
	toolsConfig "github.com/go-admin-team/go-admin-core/sdk/config"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Setup 配置Grpc
func Setup() {
	mongoSource := toolsConfig.MongoSourceObj
	// 植入mongo client
	mongoClients := make(map[string]*mongo.Client)
	for name, config := range mongoSource.MongoConfigs {
		defaultConfig := mongox.DefaultConfig()
		defaultConfig.DSN = config.Dsn
		defaultConfig.PoolLimit = config.PoolLimit
		defaultConfig.SocketTimeout = time.Duration(config.SocketTimeout) * time.Second
		mongoClients[name] = defaultConfig.Build()
	}
	sdk.Runtime.SetMongoClients(mongoClients)
}
