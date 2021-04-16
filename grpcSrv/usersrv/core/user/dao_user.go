package user

import (
	"fmt"
	"usersrv/proto"

	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

type UserDao struct {
	runner *dbx.TxRunner
}

func (dao *UserDao) GetUserList(page int, offset int) []*User {

	var userInfo []*User
	sql := "select * from user  limit ? offset ?"
	err := dao.runner.Find(&userInfo, sql, page, offset)
	fmt.Printf("%+v", userInfo[0])
	//dao.runner.Query()
	if err != nil {
		logrus.Error(err)
	}
	return userInfo
}
func (dao *UserDao) GetOne(mobile string) (*User, error) {
	po := &User{Mobile: mobile}
	_, err := dao.runner.GetOne(po)
	if err != nil {
		return nil, err
	}
	return po, err
}
func (dao *UserDao) GetById(id int32) (*User, error) {
	po := &User{ID: id}
	_, err := dao.runner.GetOne(po)
	if err != nil {
		return nil, err
	}
	return po, err
}
func (dao *UserDao) ToDTO(po *User) *proto.UserInfoResponse {
	userInfo := &proto.UserInfoResponse{}
	userInfo.Birthday = po.BirthDay.String() //todo
	userInfo.Gender = int32(po.Gender)
	userInfo.Id = po.ID
	userInfo.Gender = int32(po.Gender)
	userInfo.NickName = po.NickName
	userInfo.Role = int32(po.Role)
	userInfo.Password = po.Password
	userInfo.Mobile = po.Mobile
	return userInfo
}

// func ParseRows(rows *sql.Rows) []map[string]interface{} {
// 	columns, _ := rows.Columns()
// 	scanArgs := make([]interface{}, len(columns))
// 	values := make([]interface{}, len(columns))
// 	for j := range values {
// 		scanArgs[j] = &values[j]
// 	}

// 	record := make(map[string]interface{})
// 	records := make([]map[string]interface{}, 0)
// 	for rows.Next() {
// 		//将行数据保存到record字典
// 		err := rows.Scan(scanArgs...)
// 		checkErr(err)

// 		for i, col := range values {
// 			if col != nil {
// 				record[columns[i]] = col
// 			}
// 		}
// 		records = append(records, record)
// 	}
// 	return records
// }

// func checkErr(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 		panic(err)
// 	}
// }
