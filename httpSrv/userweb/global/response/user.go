package response

type UserResponse struct {
	ID       int32  `json:"id"`
	Mobile   string `json:"mobile"` //必须添加uni 唯一索引查找？ 如何实现？
	NickName string `json:"nickname"`
	BirthDay string `json:"birthday"` //todo timestamp
	Gender   int    `json:"gender"`
}
