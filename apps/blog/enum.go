package blog

// 文章类型，原创/转载
type BlogType uint8

const (
	TYPE_ORIGINAL  BlogType = iota // 原创
	TYPE_REPRINTED                 // 转载
	end                            // 私有常量，用于检查枚举值的有效性
)

func (t BlogType) isValid() bool {
	return t < end
}

func (t BlogType) Set(x BlogType) *BlogType {
	return &x
}

// 文章可见范围，全部可见/仅我可见
type VisibleRange uint8

const (
	Range_ALL VisibleRange = iota // 全部可见
	Range_OWN                     // 仅我可见
	range_end
)

func (v VisibleRange) isValid() bool {
	return v < range_end
}

func (v VisibleRange) Set(x VisibleRange) *VisibleRange {
	return &x
}

// 文章状态，草稿/已发布
type Status uint32

const (
	STATUS_DRAFT     Status = iota // 草稿
	STATUS_PUBLISHED               // 已发布
)

func (s Status) Set(x Status) *Status {
	return &x
}

// 更新模式，全量更新/增量更新
type UpdateMode int

const (
	UPDATE_MODE_PUT   UpdateMode = iota // 全量更新
	UPDATE_MODE_PATCH                   // 增量更新
)

// 文章列表的关键字查询条件，当前支持Title、Author
type QueryBy int

const (
	QUERY_BY_TITLE QueryBy = iota
	QUERY_BY_AUTHOR
)
