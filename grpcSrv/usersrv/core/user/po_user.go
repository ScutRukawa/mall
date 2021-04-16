package user

import "time"

type UserGender int
type UserRole int

const (
	Male   UserGender = 1
	Female UserGender = 0
)
const (
	Admin   UserRole = 1
	General UserRole = 0
)

type User struct {
	ID       int32     `db:"id,uni"`
	Mobile   string    `db:"mobile,uni"` //必须添加uni 唯一索引查找？ 如何实现？
	Password string    `db:"password"`
	NickName string    `db:"nickname"`
	Headurl  string    `db:"headurl"`
	BirthDay time.Time `db:"birthday"`
	Address  string    `db:"address"`
	Desc     string    `db:"desc"`
	Gender   int       `db:"gender"`
	Role     int       `db:"role"`
}
