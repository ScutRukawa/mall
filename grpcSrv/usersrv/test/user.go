package main

import (
	"context"
	"usersrv/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}
	defer conn.Close()
	userClient := proto.NewUserClient(conn)
	userRequest := &proto.PageInfo{
		Pn:    1,
		PSize: 10,
	}
	mobileRequest := &proto.MobileRequest{
		Mobile: "17306674706",
	}
	IDRequest := &proto.IDRequest{
		Id: 171,
	}
	userList, err := userClient.GetUserList(context.Background(), userRequest)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(userList)

	userInfo, err := userClient.GetUserByMobile(context.Background(), mobileRequest)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(userInfo)

	userInfo2, err := userClient.GetUserById(context.Background(), IDRequest)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(userInfo2)
}
