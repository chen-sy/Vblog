package blog

// 定义领域模型
type Blog struct {
	// 文章id
	ID int
	// 文章名称
	BlogName string
	// 文章概要
	BlogSummary string
	// 文章内容
	BlogContent string
	// 作者
	Author string
	// 创建时间戳
	CreateAt int
}
