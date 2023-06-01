package handler

import (
	"github.com/gin-gonic/gin"
	"grpc-todolist/app/gateway/rpc"
	pb "grpc-todolist/idl/pb/task"
	"grpc-todolist/pkg/ctl"
	"grpc-todolist/pkg/util/jwt"
	"net/http"
)

func GetTaskList(ginCtx *gin.Context) {
	var req pb.TaskRequest
	if err := ginCtx.Bind(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, ctl.RespError(ginCtx, err, "绑定参数错误"))
		return
	}
	claim, err := jwt.ParseToken(ginCtx.GetHeader("Authorization"))
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "获取用户信息错误"))
		return
	}
	req.UserID = int64(claim.UserID)
	r, err := rpc.TaskShow(ginCtx, &req)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "TaskShow RPC调用错误"))
		return
	}

	ginCtx.JSON(http.StatusOK, ctl.RespSuccess(ginCtx, r))
}

func CreateTask(ginCtx *gin.Context) {
	var req pb.TaskRequest
	if err := ginCtx.Bind(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, ctl.RespError(ginCtx, err, "绑定参数错误"))
		return
	}
	claim, err := jwt.ParseToken(ginCtx.GetHeader("Authorization"))
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "获取用户信息错误"))
		return
	}
	req.UserID = int64(claim.UserID)
	r, err := rpc.TaskCreate(ginCtx, &req)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "TaskShow RPC调用错误"))
		return
	}

	ginCtx.JSON(http.StatusOK, ctl.RespSuccess(ginCtx, r))
}

func SearchTask(ginCtx *gin.Context) {
	var req pb.TaskRequest
	if err := ginCtx.Bind(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, ctl.RespError(ginCtx, err, "绑定参数错误"))
		return
	}
	claim, err := jwt.ParseToken(ginCtx.GetHeader("Authorization"))
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "获取用户信息错误"))
		return
	}
	req.UserID = int64(claim.UserID)
	r, err := rpc.TaskSearch(ginCtx, &req)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "TaskShow RPC调用错误"))
		return
	}

	ginCtx.JSON(http.StatusOK, ctl.RespSuccess(ginCtx, r))
}

func UpdateTask(ginCtx *gin.Context) {
	var req pb.TaskRequest
	if err := ginCtx.Bind(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, ctl.RespError(ginCtx, err, "绑定参数错误"))
		return
	}
	claim, err := jwt.ParseToken(ginCtx.GetHeader("Authorization"))
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "获取用户信息错误"))
		return
	}
	req.UserID = int64(claim.UserID)
	r, err := rpc.TaskUpdate(ginCtx, &req)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "TaskShow RPC调用错误"))
		return
	}

	ginCtx.JSON(http.StatusOK, ctl.RespSuccess(ginCtx, r))
}

func UpdateAllTask(ginCtx *gin.Context) {
	var req pb.TaskRequest
	if err := ginCtx.Bind(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, ctl.RespError(ginCtx, err, "绑定参数错误"))
		return
	}
	claim, err := jwt.ParseToken(ginCtx.GetHeader("Authorization"))
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "获取用户信息错误"))
		return
	}
	req.UserID = int64(claim.UserID)
	r, err := rpc.TaskUpdateAll(ginCtx, &req)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "TaskShow RPC调用错误"))
		return
	}

	ginCtx.JSON(http.StatusOK, ctl.RespSuccess(ginCtx, r))
}

func DeleteTask(ginCtx *gin.Context) {
	var req pb.TaskRequest
	if err := ginCtx.Bind(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, ctl.RespError(ginCtx, err, "绑定参数错误"))
		return
	}
	claim, err := jwt.ParseToken(ginCtx.GetHeader("Authorization"))
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "获取用户信息错误"))
		return
	}
	req.UserID = int64(claim.UserID)
	r, err := rpc.TaskDelete(ginCtx, &req)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "TaskShow RPC调用错误"))
		return
	}

	ginCtx.JSON(http.StatusOK, ctl.RespSuccess(ginCtx, r))
}

func DeleteAllTask(ginCtx *gin.Context) {
	var req pb.TaskRequest
	if err := ginCtx.Bind(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, ctl.RespError(ginCtx, err, "绑定参数错误"))
		return
	}
	claim, err := jwt.ParseToken(ginCtx.GetHeader("Authorization"))
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "获取用户信息错误"))
		return
	}
	req.UserID = int64(claim.UserID)
	r, err := rpc.TaskDeleteAll(ginCtx, &req)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, ctl.RespError(ginCtx, err, "TaskShow RPC调用错误"))
		return
	}

	ginCtx.JSON(http.StatusOK, ctl.RespSuccess(ginCtx, r))
}
