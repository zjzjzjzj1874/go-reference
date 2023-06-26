package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"time"
)

// Client custom business
func pingLoop(conn ziface.IConnection) {
	for {
		err := conn.SendMsg(1, []byte("hello server:Ping...Ping...Ping...[FromClient]"))
		if err != nil {
			fmt.Println(err)
			break
		}

		time.Sleep(1 * time.Second)
	}
}

// Executed when a connection is created
func onClientStart(conn ziface.IConnection) {
	fmt.Println("onClientStart is Called ... ")
	go pingLoop(conn)
}

func main() {
	//Create a client client
	client := znet.NewClient("127.0.0.1", 8999)

	//Set the hook function after the link is successfully established
	client.SetOnConnStart(onClientStart)

	//start the client
	client.Start()

	//Prevent the process from exiting, waiting for an interrupt signal
	select {}
}
