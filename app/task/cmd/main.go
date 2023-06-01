package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"grpc-todolist/app/task/internal/repository/db/dao"
	"grpc-todolist/app/task/internal/service"
	"grpc-todolist/config"
	taskPb "grpc-todolist/idl/pb/task"
	"grpc-todolist/pkg/discovery"
	"net"
)

func main() {
	config.InitConfig()
	dao.InitDB()
	// etcd 地址
	etcdAddress := []string{config.Conf.Etcd.Address}
	// 服务注册
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	grpcAddress := config.Conf.Services["task"].Addr[0]
	defer etcdRegister.Stop()
	taskNode := discovery.Server{
		Name: config.Conf.Domain["task"].Name,
		Addr: grpcAddress,
	}
	server := grpc.NewServer()
	defer server.Stop()
	// 绑定service
	taskPb.RegisterTaskServiceServer(server, service.GetTaskSrv())
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err = etcdRegister.Register(taskNode, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %v", err))
	}
	logrus.Info("server started listen on ", grpcAddress)
	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
