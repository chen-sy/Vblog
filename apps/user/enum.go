package user

//定义枚举时，常常会显示的声明其类型，比如这样：
type Sex int

//这样做有一个好处，我们可以对我们的 Sex 类型进行扩展
func (this Sex) String() string {
	switch this {
	case 0:
		return "男"
	case 1:
		return "女"
	default:
		return "未知"
	}
}

const (
	SEX_MAN     Sex = iota //男
	SEX_WOMAN              //女
	SEX_UNKNOWN            //未知
)

// 使用Role类型来表现枚举类型
type Role int

const (
	ROLE_MEMBER = iota // 普通用户
	ROLE_ADMIN         // 管理员
)
