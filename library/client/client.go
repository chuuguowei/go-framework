package client

import (
	"go-framework/library/cache"
	"go-framework/library/mysql"
	"go-framework/library/redisx"
)

var MySqlClient mysql.SqlConn

var RedisClient *redisx.Client

var CacheClient *library.CacheX
