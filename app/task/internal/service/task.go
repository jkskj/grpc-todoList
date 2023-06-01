package service

import (
	"context"
	"grpc-todolist/app/task/internal/repository/db/dao"
	pb "grpc-todolist/idl/pb/task"
	"grpc-todolist/pkg/e"
	"sync"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
	pb.UnimplementedTaskServiceServer
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (*TaskSrv) TaskCreate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskCommonResponse, err error) {
	resp = new(pb.TaskCommonResponse)
	resp.Code = e.SUCCESS
	err = dao.NewTaskDao(ctx).CreateTask(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return
}

func (*TaskSrv) TaskShow(ctx context.Context, req *pb.TaskRequest) (resp *pb.TasksDetailResponse, err error) {
	resp = new(pb.TasksDetailResponse)
	resp.Code = e.SUCCESS
	taskList, err := dao.NewTaskDao(ctx).ShowTask(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	resp.TaskDetail = dao.BuildTasks(taskList)
	return
}

func (*TaskSrv) TaskSearch(ctx context.Context, req *pb.TaskRequest) (resp *pb.TasksDetailResponse, err error) {
	resp = new(pb.TasksDetailResponse)
	resp.Code = e.SUCCESS
	taskList, err := dao.NewTaskDao(ctx).SearchTask(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	resp.TaskDetail = dao.BuildTasks(taskList)
	return
}

func (*TaskSrv) TaskUpdate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskCommonResponse, err error) {
	resp = new(pb.TaskCommonResponse)
	resp.Code = e.SUCCESS
	err = dao.NewTaskDao(ctx).UpdateTask(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return
}
func (*TaskSrv) TaskUpdateAll(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskCommonResponse, err error) {
	resp = new(pb.TaskCommonResponse)
	resp.Code = e.SUCCESS
	err = dao.NewTaskDao(ctx).UpdateAllTask(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return
}
func (*TaskSrv) TaskDelete(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskCommonResponse, err error) {
	resp = new(pb.TaskCommonResponse)
	resp.Code = e.SUCCESS
	err = dao.NewTaskDao(ctx).DeleteTask(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return
}
func (*TaskSrv) TaskDeleteAll(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskCommonResponse, err error) {
	resp = new(pb.TaskCommonResponse)
	resp.Code = e.SUCCESS
	err = dao.NewTaskDao(ctx).DeleteAllTask(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return
}
