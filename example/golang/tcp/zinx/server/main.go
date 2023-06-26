package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

// PingRouter MsgId=1
type PingRouter struct {
	znet.BaseRouter
}

//Ping Handle MsgId=1
func (r *PingRouter) Handle(request ziface.IRequest) {
	//read client data
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
}

func main() {
	//1 Create a server service
	s := znet.NewServer()

	//2 configure routing
	s.AddRouter(1, &PingRouter{})

	//3 start service
	s.Serve()
}