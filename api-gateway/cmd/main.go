package main

import (
	"api-gateway/config"
	"api-gateway/discovery"
	"api-gateway/internal/service"
	"api-gateway/routes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.InitConfig()
	// etcd 地址
	etcdAddress := []string{viper.GetString("etcd.address")}
	// 服务注册
	etcdRegister := discovery.NewResolver(etcdAddress, logrus.New())
	resolver.Register(etcdRegister)
	defer etcdRegister.Close()
	fmt.Println("gateway listen on ", viper.GetString("server.port"))
	go startListen() // 转载路由
	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignals
		fmt.Println("exit! ", s)
	}

}
func startListen() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	userConn, err := grpc.Dial(viper.GetString("user.address"), opts...)
	if err != nil {
		panic(err)
	}
	userService := service.NewUserServiceClient(userConn)

	taskConn, err1 := grpc.Dial(viper.GetString("task.address"), opts...)
	if err1 != nil {
		panic(err1)
	}
	taskService := service.NewTaskServiceClient(taskConn)

	ginRouter := routes.NewRouter(userService, taskService)
	server := &http.Server{
		Addr:           viper.GetString("server.port"),
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err = server.ListenAndServe(); err != nil {
		fmt.Println("gateway启动失败, err: ", err)
	}
}
