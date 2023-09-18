// 公用包
package common

import "time"

// meta的构造函数，用于初始化meta数据，防止空指针
func NewMeta() *Meta {
	return &Meta{
		CreatedAt: time.Now().Unix(),
	}
}

// 元数据
type Meta struct {
	// 用户id
	Id int64 `json:"id"`
	// 创建时间
	CreatedAt int64 `json:"create_at"`
	// 更新时间
	UpdatedAt int64 `json:"update_at"`
}
