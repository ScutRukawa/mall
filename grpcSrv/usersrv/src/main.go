package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"usersrv/core/user"
	_ "usersrv/db"
	"usersrv/proto"

	"google.golang.org/grpc"
)

func main() {
	g := grpc.NewServer()
	proto.RegisterUserServer(g, &user.UserServicer{})
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}
	go func() {
		g.Serve(listener)
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT)
	<-quit
	fmt.Println("xxxxxx")
}
