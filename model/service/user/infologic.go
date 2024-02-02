package user

import (
	"context"
	"fmt"
	"go-framework/library/client"
	"go-framework/library/common/errorx"
	"go-framework/library/logger"
	"go-framework/model/dao"
	"go-framework/model/types"
	"time"
)

type InfoLogic struct {
	logger.Logger
	ctx       context.Context
	UserModel dao.UserModel
}

func NewInfoLogic(ctx context.Context) *InfoLogic {
	return &InfoLogic{
		ctx:       ctx,
		Logger:    logger.WithContext(ctx),
		UserModel: dao.NewUserModel(ctx),
	}
}

func (that *InfoLogic) Handle(req *types.UserInfoReq) (resp *types.UserInfoReply, err error) {
	_, err = client.RedisClient.Expire("test", time.Nanosecond/1000)
	if err != nil {
		return nil, err
	}

	that.Logger.Debug("这是debug消息")
	that.Logger.Info("这是info消息", "这是info日志22")
	that.Logger.Warn("这是warn消息")
	that.Logger.Error(req)

	that.Logger.DebugF("debug测试 %s %d", "姓名", 43)
	that.Logger.InfoF("info测试 %s %d", "姓名", 43)
	that.Logger.WarnF("warn测试 %s %d", "姓名", 43)
	that.Logger.ErrorF("error测试 %s %d", "姓名", 43)
	user, err := that.UserModel.List(100)
	that.Logger.Info(user)
	if err != nil {
		return
	}
	fmt.Println(user)

	if user == nil {
		err = errorx.NewDefaultError("数据不存在")
		return
	}
	resp = &types.UserInfoReply{
		Message: "hello",
		Ctime:   user[0].Ctime,
	}
	return

}
