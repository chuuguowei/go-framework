package bootstrap

import (
	"go-framework/config"
	"go-framework/library/cache"
	"go-framework/library/client"
	"go-framework/library/configs"
	"go-framework/library/logger"
	"go-framework/library/mysql"
	"go-framework/library/redisx"
)

func MustInit(c config.Config) {
	logger.New(c.Log.Level, 2)

	initMySQL(c)

	initRedis(c)

	initCache(c)

	initBizConfig(c)
}

func initMySQL(c config.Config) {
	// 初始化mysql
	client.MySqlClient = mysql.NewMysql(c.Mysql.DataSource)
}

func initRedis(c config.Config) {
	// 初始化redis
	client.RedisClient = redisx.New(&redisx.Options{
		Network:  c.Redis.Network,
		Addr:     c.Redis.Addr,
		Username: c.Redis.Username,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})
}

func initCache(c config.Config) {
	// 初始化缓存
	client.CacheClient = library.NewCacheX(c)
}

func initBizConfig(c config.Config) {
	// 初始化业务配置
	configs.BizConfig = c
}
