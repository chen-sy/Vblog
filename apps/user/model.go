package user

// 定义用户实体对象
type User struct {
	// 用户id
	Id int64 `json:"id"`
	// 创建时间
	CreatedAt int64 `json:"create_at"`
	// 更新时间
	UpdatedAt int64 `json:"update_at"`
	// 状态
	State int `json:"state"`
	*CreateUserRequest
}
