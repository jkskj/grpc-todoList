package handler

import (
	"context"
	"task/internal/repository"
	"task/internal/service"
	"task/pkg/e"
)

type TaskService struct {
}

func NewTaskService() *TaskService {
	return &TaskService{}
}
func (*TaskService) TaskCreate(ctx context.Context, req *service.TaskRequest) (resp *service.TaskCommonResponse, err error) {
	var task repository.Task
	resp = new(service.TaskCommonResponse)
	resp.Code = e.SUCCESS
	err = task.CreateTask(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return
}
func (*TaskService) TaskShow(ctx context.Context, req *service.TaskRequest) (resp *service.TasksDetailResponse, err error) {
	var task repository.Task
	resp = new(service.TasksDetailResponse)
	resp.Code = e.SUCCESS
	taskList, err := task.ShowTask(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	resp.TaskDetail = repository.BuildTasks(taskList)
	return
}
func (*TaskService) TaskSearch(ctx context.Context, req *service.TaskRequest) (resp *service.TasksDetailResponse, err error) {
	var task repository.Task
	resp = new(service.TasksDetailResponse)
	resp.Code = e.SUCCESS
	taskList, err := task.SearchTask(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	resp.TaskDetail = repository.BuildTasks(taskList)
	return
}
func (*TaskService) TaskUpdate(ctx context.Context, req *service.TaskRequest) (resp *service.TaskCommonResponse, err error) {
	var task repository.Task
	resp = new(service.TaskCommonResponse)
	resp.Code = e.SUCCESS
	err = task.UpdateTask(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return
}
func (*TaskService) TaskUpdateAll(ctx context.Context, req *service.TaskRequest) (resp *service.TaskCommonResponse, err error) {
	var task repository.Task
	resp = new(service.TaskCommonResponse)
	resp.Code = e.SUCCESS
	err = task.UpdateAllTask(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return
}
func (*TaskService) TaskDelete(ctx context.Context, req *service.TaskRequest) (resp *service.TaskCommonResponse, err error) {
	var task repository.Task
	resp = new(service.TaskCommonResponse)
	resp.Code = e.SUCCESS
	err = task.DeleteTask(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return
}
func (*TaskService) TaskDeleteAll(ctx context.Context, req *service.TaskRequest) (resp *service.TaskCommonResponse, err error) {
	var task repository.Task
	resp = new(service.TaskCommonResponse)
	resp.Code = e.SUCCESS
	err = task.DeleteAllTask(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return
}
