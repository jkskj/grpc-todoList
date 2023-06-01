package rpc

import (
	"context"
	"errors"
	"grpc-todolist/pkg/e"

	userPb "grpc-todolist/idl/pb/user"
)

func UserLogin(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserResponse, err error) {
	r, err := UserClient.UserLogin(ctx, req)
	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New("登陆失败")
		return
	}

	return r.UserDetail, nil
}

func UserRegister(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserResponse, err error) {
	r, err := UserClient.UserRegister(ctx, req)
	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New("注册失败")
		return
	}

	return r.UserDetail, nil
}
