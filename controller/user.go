package controller

import (
	"github.com/gin-gonic/gin"
	"go-framework/library/common/errorx"
	"go-framework/library/common/response"
	"go-framework/library/logger"
	"go-framework/model/service/user"
	"go-framework/model/types"
)

func AddHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.UserAddReq
		if err := ctx.ShouldBind(&req); err != nil {
			logger.WithContext(ctx).ErrorF("Handler ShouldBind Parse %v", err)
			response.Result(ctx, errorx.NewCodeError(1, err.Error()))

			return
		}

		logic := user.NewAddLogic(ctx)
		resp, err := logic.Handle(&req)
		if err != nil {
			response.Result(ctx, err)
		} else {
			response.Result(ctx, resp)
		}
	}
}

func IndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.UserIndexReq
		if err := ctx.ShouldBind(&req); err != nil {
			logger.WithContext(ctx).ErrorF("Handler ShouldBind Parse %v", err)
			response.Result(ctx, errorx.NewCodeError(1, err.Error()))
			return
		}

		logic := user.NewIndexLogic(ctx)
		resp, err := logic.Handle(&req)
		if err != nil {
			response.Result(ctx, err)
		} else {
			response.Result(ctx, resp)
		}
	}
}

func InfoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserInfoReq
		if err := ctx.ShouldBind(&req); err != nil {
			logger.WithContext(ctx).ErrorF("Handler ShouldBind Parse %v", err)
			response.Result(ctx, errorx.NewCodeError(1, err.Error()))
			return
		}

		logic := user.NewInfoLogic(ctx)
		resp, err := logic.Handle(&req)
		if err != nil {
			response.Result(ctx, err)
		} else {
			response.Result(ctx, resp)
		}
	}
}
