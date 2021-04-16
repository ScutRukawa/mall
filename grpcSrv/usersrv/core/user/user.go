package user

import (
	"context"
	"usersrv/db"
	"usersrv/proto"

	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type UserServicer struct {
}

func (u *UserServicer) GetUserList(ctx context.Context, request *proto.PageInfo) (res *proto.UserListResponse, errRsp error) {
	logrus.Info("recieve a  request")
	res = new(proto.UserListResponse)
	logrus.Info("res: %v", res)

	userDao := UserDao{}
	err := db.Tx(func(runner *dbx.TxRunner) error {
		page := 0
		pageSize := 10
		if request.PSize != 0 {
			pageSize = int(request.PSize)
		}
		if request.Pn != 0 {
			page = int(request.Pn-1) * pageSize
		}
		userDao.runner = runner

		user := userDao.GetUserList(pageSize, page)
		userInfos := make([]*proto.UserInfoResponse, 0)
		for _, u := range user {
			userInfos = append(userInfos, userDao.ToDTO(u))
		}
		res.Data = userInfos
		return nil
	})
	if err != nil {
		return res, nil
	}
	res.Total = int32(len(res.Data))
	return res, nil
}

func (u *UserServicer) GetUserByMobile(ctx context.Context, request *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	userDao := UserDao{}
	userInfo := &proto.UserInfoResponse{}
	err := db.Tx(func(runner *dbx.TxRunner) error {
		userDao.runner = runner
		user, err := userDao.GetOne(request.Mobile)
		if err != nil {
			return err
		}
		userInfo = userDao.ToDTO(user)
		return nil
	})
	if err != nil {
		return &proto.UserInfoResponse{}, err
	}
	return userInfo, err
}

func (u *UserServicer) GetUserById(ctx context.Context, request *proto.IDRequest) (*proto.UserInfoResponse, error) {
	userDao := UserDao{}
	userInfo := &proto.UserInfoResponse{}
	err := db.Tx(func(runner *dbx.TxRunner) error {
		userDao.runner = runner
		user, err := userDao.GetById(request.Id)
		if err != nil {
			return err
		}
		userInfo = userDao.ToDTO(user)
		return nil
	})
	if err != nil {
		return &proto.UserInfoResponse{}, err
	}
	return userInfo, err
}

func (u *UserServicer) CreateUser(ctx context.Context, request *proto.CreateUserInfo) (*proto.UserInfoResponse, error) { //todo

	return &proto.UserInfoResponse{}, nil
}
func (u *UserServicer) UpdateUser(ctx context.Context, request *proto.UpdateUserInfo) (*emptypb.Empty, error) { //todo
	return &emptypb.Empty{}, nil
}

// func (u *UserServicer) CheckPassword(ctx context.Context, in *proto.PasswordCheckInfo) (proto.CheckResponse, error) {
// 	return
// }

// func (u *UserServicer) GetUserList(ctx context.Context, request *proto.PageInfo) (res *proto.UserListResponse, errRsp error) {
// 	logrus.Info("recieve a  request")
// 	userDao := UserDao{}
// 	var queryResult []map[string]interface{}
// 	userInfo := make([]*proto.UserInfoResponse, request.PSize)
// 	err := db.Tx(func(runner *dbx.TxRunner) error {
// 		page := 0
// 		pageSize := 10
// 		if request.PSize != 0 {
// 			pageSize = int(request.PSize)
// 		}
// 		if request.Pn != 0 {
// 			page = int(request.Pn-1) * pageSize
// 		}
// 		userDao.runner = runner
// 		queryResult = userDao.GetUserList(pageSize, page)
// 		return nil
// 	})
// 	if err != nil {
// 		return res, nil
// 	}
// 	for _, m := range queryResult {
// 		oneUser := new(proto.UserInfoResponse)
// 		logrus.Info(m["id"].([]byte))
// 		val, _ := m["id"].([]byte)
// 		id, err := strconv.ParseInt(string(val), 10, 32)
// 		logrus.Error(err)
// 		logrus.Info(id)
// 		oneUser.Id = int32(id)

// 		oneUser.NickName = m["nickname"].(string)
// 		oneUser.Gender = m["gender"].(int32)
// 		oneUser.Birthday = m["birthday"].(uint64)
// 		userInfo = append(userInfo, oneUser)
// 	}
// 	res.Data = userInfo
// 	res.Total = int32(len(queryResult))
// 	return res, nil
// }
