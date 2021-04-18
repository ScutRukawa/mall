package main

import (
	"fmt"
	"net"
	"strconv"
	v1 "usersrv/common/v1"
	"usersrv/core/user"
	_ "usersrv/db"
	"usersrv/grpc_health_v1"
	"usersrv/proto"
	"usersrv/utils"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
)

func main() {

	// 创建UUID
	u1 := uuid.Must(uuid.NewV4()).String() 
	fmt.Printf("UUIDv4: %s\n", u1)

	g := grpc.NewServer()
	proto.RegisterUserServer(g, &user.UserServicer{})
	grpc_health_v1.RegisterHealthServer(g, &v1.HealthImpl{})
	port, _ := utils.GetFreePort()
	listener, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	fmt.Println("Using port:", port)
	if err != nil {
		panic(err)
	}
	go func() {
		g.Serve(listener)
	}()
	v1.Register("127.0.0.1", port, "user_srv", []string{"mall", "wei"}, u1)
	v1.OnExit(u1)
}
