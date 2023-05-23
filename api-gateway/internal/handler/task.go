package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTaskList(ginCtx *gin.Context) {
	var req service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&req))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	req.UserID = int64(claim.UserID)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskShow(context.Background(), &req)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func CreateTask(ginCtx *gin.Context) {
	var req service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&req))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	req.UserID = int64(claim.UserID)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskCreate(context.Background(), &req)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UpdateTask(ginCtx *gin.Context) {
	var req service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&req))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	req.UserID = int64(claim.UserID)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskUpdate(context.Background(), &req)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func DeleteTask(ginCtx *gin.Context) {
	var req service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&req))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	req.UserID = int64(claim.UserID)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskDelete(context.Background(), &req)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}
