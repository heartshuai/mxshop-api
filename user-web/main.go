package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop-api/user-web/initialize"
)

func main() {
	//1.初始化logger
	initialize.InitLogger()
	Router := initialize.Routers()

	port := 8021

	/**
	1.S()可以获取一个全局的suger,可以让我们自己设置一个全局的logger
	2.日志是分级别的，debug,info,warn,error,fatal,panic
	3.S函数和L函数很有用 提供了一个全局的安全访问logger的途径
	*/
	zap.S().Debugf("启动服务器 端口:%d", port)

	err := Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}

}
