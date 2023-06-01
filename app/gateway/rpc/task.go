package rpc

import (
	"context"
	"errors"
	taskPb "grpc-todolist/idl/pb/task"
	"grpc-todolist/pkg/e"
)

func TaskCreate(ctx context.Context, req *taskPb.TaskRequest) (resp *taskPb.TaskCommonResponse, err error) {
	r, err := TaskClient.TaskCreate(ctx, req)

	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New(r.Msg)
		return
	}

	return r, nil
}

func TaskSearch(ctx context.Context, req *taskPb.TaskRequest) (resp *taskPb.TasksDetailResponse, err error) {
	r, err := TaskClient.TaskSearch(ctx, req)

	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New("获取失败")
		return
	}

	return r, nil
}

func TaskUpdate(ctx context.Context, req *taskPb.TaskRequest) (resp *taskPb.TaskCommonResponse, err error) {
	r, err := TaskClient.TaskUpdate(ctx, req)

	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New(r.Msg)
		return
	}

	return r, nil
}

func TaskUpdateAll(ctx context.Context, req *taskPb.TaskRequest) (resp *taskPb.TaskCommonResponse, err error) {
	r, err := TaskClient.TaskUpdateAll(ctx, req)

	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New(r.Msg)
		return
	}

	return r, nil
}

func TaskDelete(ctx context.Context, req *taskPb.TaskRequest) (resp *taskPb.TaskCommonResponse, err error) {
	r, err := TaskClient.TaskDelete(ctx, req)

	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New(r.Msg)
		return
	}

	return r, nil
}

func TaskDeleteAll(ctx context.Context, req *taskPb.TaskRequest) (resp *taskPb.TaskCommonResponse, err error) {
	r, err := TaskClient.TaskDeleteAll(ctx, req)

	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New(r.Msg)
		return
	}

	return r, nil
}

func TaskShow(ctx context.Context, req *taskPb.TaskRequest) (resp *taskPb.TasksDetailResponse, err error) {
	r, err := TaskClient.TaskShow(ctx, req)

	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New("获取失败")
		return
	}

	return r, nil
}
