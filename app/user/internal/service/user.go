package service

import (
	"context"
	"grpc-todolist/app/user/internal/repository/db/dao"
	pb "grpc-todolist/idl/pb/user"
	"grpc-todolist/pkg/e"
	"sync"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
	pb.UnimplementedUserServiceServer
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (*UserSrv) UserLogin(ctx context.Context, req *pb.UserRequest) (resp *pb.UserDetailResponse, err error) {
	resp = new(pb.UserDetailResponse)
	resp.Code = e.SUCCESS
	user, err := dao.NewUserDao(ctx).ShowUserInfo(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	resp.UserDetail = &pb.UserResponse{
		UserID:   user.UserID,
		UserName: user.UserName,
		NickName: user.UserName,
	}
	return
}

func (*UserSrv) UserRegister(ctx context.Context, req *pb.UserRequest) (resp *pb.UserDetailResponse, err error) {
	resp = new(pb.UserDetailResponse)
	resp.Code = e.SUCCESS
	user, err := dao.NewUserDao(ctx).Create(req)
	if err != nil {

		resp.Code = e.ERROR
		return resp, err
	}
	resp.UserDetail = &pb.UserResponse{
		UserID:   user.UserID,
		UserName: user.UserName,
		NickName: user.UserName,
	}
	return resp, nil
}

func (*UserSrv) UserLogout(ctx context.Context, req *pb.UserRequest) (resp *pb.UserDetailResponse, err error) {
	resp = new(pb.UserDetailResponse)
	return resp, nil
}
