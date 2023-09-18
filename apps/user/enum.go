package user

//使用类型别名来定义枚举类型，sex 是 int 别名
type Sex int

const (
	SEX_MAN     Sex = iota //男
	SEX_WOMAN              //女
	SEX_UNKNOWN            //未知
)

//通过对枚举类型重写 String 方法，可以对枚举值进行自定义，可以清晰地描述该枚举变量的意义和作用
func (s Sex) String() string {
	return [...]string{"男", "女", "未知"}[s]
}

// 声明一个 Role 类型
type Role int

const (
	ROLE_MEMBER Role = iota // 普通用户
	ROLE_ADMIN              // 管理员
)

// 声明一个 State 类型
type State int

const (
	STATE_NORMAL State = 1  // 正常
	STATE_CANCEL State = -1 // 失效
)

// 单个用户的查询条件，当前支持id、UserName
type QueryBy int

const (
	QUERY_BY_ID QueryBy = iota
	QUERY_BY_USERNAME
)
