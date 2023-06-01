package main

import (
	"fmt"
	"grpc-todolist/app/gateway/routes"
	"grpc-todolist/app/gateway/rpc"
	"grpc-todolist/config"
	"grpc-todolist/pkg/util/shutdown"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.InitConfig()
	rpc.Init()

	go startListen() // 转载路由
	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignals
		fmt.Println("exit! ", s)
	}
	fmt.Println("gateway listen on :3000")
}

func startListen() {
	r := routes.NewRouter()
	server := &http.Server{
		Addr:           config.Conf.Server.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("gateway启动失败, err: ", err)
	}
	go func() {
		// 优雅关闭
		shutdown.GracefullyShutdown(server)
	}()
}
