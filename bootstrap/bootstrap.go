package bootstrap

import (
	"go-framework/config"
	"go-framework/library/cache"
	"go-framework/library/logger"
	"go-framework/library/mysql"
	"go-framework/library/redisx"

	"go-framework/library/client"
)

func MustInit(c config.Config) {
	logger.New(c.Log.Level, 2)

	initMySQL(c)

	initRedis(c)

	initCache(c)
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
