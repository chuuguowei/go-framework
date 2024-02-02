package user

import (
	"context"
	"fmt"
	"go-framework/library/client"
	"go-framework/library/common/errorx"
	"go-framework/library/logger"
	"go-framework/model/types"
	"time"
)

type IndexLogic struct {
	logger.Logger
	ctx context.Context
}

func NewIndexLogic(ctx context.Context) *IndexLogic {
	return &IndexLogic{
		ctx:    ctx,
		Logger: logger.WithContext(ctx),
	}
}

func (that *IndexLogic) Handle(req *types.UserIndexReq) (resp *types.UserIndexReply, err error) {
	val, found := client.CacheClient.Get(req.Name)
	var message string
	found = false
	if found {
		message = " get value from cache:" + val.(string)
	} else {
		if req.Message == "" {
			return nil, errorx.NewDefaultError("消息不能为空")
		}
		result, err := client.RedisClient.SetCtx(that.ctx, "name", "测试命令行344", time.Hour)
		fmt.Println(result, err)
		client.CacheClient.Set(req.Name, req.Message)
	}

	return &types.UserIndexReply{
		Message: req.Name + " say:" + message,
	}, nil
}
