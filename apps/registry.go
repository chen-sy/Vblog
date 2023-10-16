package apps

// 通过导包的方式，加载包里面的init()，将对象托管到ioc
import (
	// 从上往下依次注册
	_ "gitee.com/chensyi/vblog/apps/blog/impl"
	_ "gitee.com/chensyi/vblog/apps/comment/impl"
	_ "gitee.com/chensyi/vblog/apps/token/impl"
	_ "gitee.com/chensyi/vblog/apps/user/impl"

	// Api Handler注册
	_ "gitee.com/chensyi/vblog/apps/blog/api"
	_ "gitee.com/chensyi/vblog/apps/comment/api"
	_ "gitee.com/chensyi/vblog/apps/token/api"
	_ "gitee.com/chensyi/vblog/apps/user/api"
)
