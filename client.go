package main

/**
* @Author: super
* @Date: 2020-08-17 14:52
* @Description:
**/

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	proto "go-micro-tutorials/proto"
)


func main() {
	// 定义服务，可以传入其它可选参数
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// 创建新的客户端
	greeter := proto.NewGreeterService("greeter", service.Client())

	// 调用greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// 打印响应请求
	fmt.Println(rsp.Greeting)
}