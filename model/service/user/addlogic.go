package user

import (
	"context"
	"go-framework/library/logger"
	"go-framework/model/dao"
	"go-framework/model/types"
)

type AddLogic struct {
	logger.Logger
	ctx       context.Context
	UserModel dao.UserModel
}

func NewAddLogic(ctx context.Context) *AddLogic {
	return &AddLogic{
		ctx:       ctx,
		Logger:    logger.WithContext(ctx),
		UserModel: dao.NewUserModel(ctx),
	}
}

func (that *AddLogic) Handle(req *types.UserAddReq) (resp *types.UserAddReply, err error) {

	userid, err := that.UserModel.Insert(&dao.User{
		Username: req.Username,
		Age:      req.Age,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserAddReply{
		UserId: userid,
	}, nil

}
