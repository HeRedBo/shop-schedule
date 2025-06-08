package main

import (
	"fmt"
	"github.com/HeRedBo/pkg/cache"
	"github.com/HeRedBo/pkg/es"
	"github.com/HeRedBo/pkg/logger"
	"github.com/HeRedBo/pkg/nosql"
	"github.com/HeRedBo/pkg/shutdown"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	"shop-schedule/global"
	"shop-schedule/internal/script/order_rebuild"
)

func init() {
	global.LoadConfig()
	global.LOG = global.SetupLogger()
	initRedisClient()
	initMongoClient()
	initESClient()
}

func initRedisClient() {
	redisCfg := global.CONFIG.Redis
	opt := redis.Options{
		Addr:        redisCfg.Host,
		Password:    redisCfg.Password,
		IdleTimeout: redisCfg.IdleTimeout,
	}
	err := cache.InitRedis(cache.DefaultRedisClient, &opt)
	if err != nil {
		global.LOG.Error("redis init error", zap.Error(err), "client", cache.DefaultRedisClient)
		panic("initRedisClient error")
	}
	global.Redis = cache.GetRedisClient(cache.DefaultRedisClient)
}

// 初始化ES
func initESClient() {
	err := es.InitClientWithOptions(es.DefaultClient, global.CONFIG.Elasticsearch.Hosts,
		global.CONFIG.Elasticsearch.Username,
		global.CONFIG.Elasticsearch.Password,
		es.WithScheme("https"))
	if err != nil {
		global.LOG.Error("InitClientWithOptions error", err, "client", es.DefaultClient)
		panic(err)
	}
	global.ES = es.GetClient(es.DefaultClient)
}

func initMongoClient() {
	mongoCfg := global.CONFIG.MongoDB
	err := nosql.InitMongoClient(nosql.DefaultMongoClient, mongoCfg.User,
		mongoCfg.Password, mongoCfg.Host, 200)
	if err != nil {
		logger.Error("InitMongoClient error", zap.Error(err), zap.String("client", nosql.DefaultMongoClient))
		panic(err)
	}
	global.Mongo = nosql.GetMongoClient(nosql.DefaultMongoClient)
}

func main() {
	switch global.CONFIG.App.Task {
	case "product-search-log":
		//product_searchlog.AnalysisSearchLog()
	case "order-rebuild":
		order_rebuild.Rebuild("3")

	}
	//优雅关闭
	shutdown.NewHook().Close(
		func() {
			// mongo
			global.Mongo.Close()
		},
		func() {
			//es
			es.CloseAll()
		},
		func() {
			err := global.Redis.Close()
			if err != nil {
				global.LOG.Error("redis close error", zap.Error(err), zap.String("client", cache.DefaultRedisClient))
			}
		},
		func() {
			if global.Mongo != nil {
				global.Mongo.Close()
			}
		},
		func() {
			err := global.LOG.Sync()
			if err != nil {
				fmt.Println(err)
			}
		},
	)
}
