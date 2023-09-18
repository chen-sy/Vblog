// 公用包
package common

// meta的构造函数，用于初始化meta数据，防止空指针
func NewMeta() *Meta {
	//meta数据如果只是下面三个，可以不初始化，因为gorm.Model也有以下字段，它会帮我们自动填充
	//GORM 定义一个 gorm.Model 结构体，其包括字段 ID、CreatedAt、UpdatedAt、DeletedAt
	//GORM 约定使用 CreatedAt、UpdatedAt 追踪创建/更新时间。如果您定义了这种字段，GORM 在创建、更新时会自动填充当前时间
	//如果您想要保存 UNIX（毫/纳）秒时间戳，而不是 time，您只需简单地将 time.Time 修改为 int 即可
	return &Meta{
		//CreatedAt: time.Now().Unix(),
	}
}

// 元数据
type Meta struct {
	// 用户id
	ID int64 `json:"id"`
	// 创建时间
	CreatedAt int `json:"created_at"`
	// 更新时间
	UpdatedAt int `json:"updated_at"`
}
