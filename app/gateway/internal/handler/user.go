package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"grpc-todolist/app/gateway/rpc"
	pb "grpc-todolist/idl/pb/user"
	"grpc-todolist/pkg/ctl"
	"grpc-todolist/pkg/res"
	"grpc-todolist/pkg/util/jwt"
	"net/http"
)

// UserRegister 用户注册
func UserRegister(ginCtx *gin.Context) {
	var userReq pb.UserRequest
	if err := ginCtx.Bind(&userReq); err != nil {
		ginCtx.JSON(http.StatusBadRequest, ctl.RespError(ginCtx, err, "绑定参数错误"))
		return
	}
	r, err := rpc.UserRegister(context.Background(), &userReq)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "UserRegister RPC服务调用错误"))
		return
	}

	ginCtx.JSON(http.StatusOK, ctl.RespSuccess(ginCtx, r))
}

// UserLogin 用户登录
func UserLogin(ginCtx *gin.Context) {
	var req pb.UserRequest
	if err := ginCtx.Bind(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, ctl.RespError(ginCtx, err, "绑定参数错误"))
		return
	}

	userResp, err := rpc.UserLogin(context.Background(), &req)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "UserLogin RPC服务调用错误"))
		return
	}

	token, err := jwt.GenerateToken(uint(userResp.UserID))
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "加密错误"))
		return
	}

	ginCtx.JSON(http.StatusOK, ctl.RespSuccess(ginCtx, res.TokenData{User: userResp, Token: token}))
}
