# Vblog

前后端分离的博客项目，用到的技术：
+ API Server
  + DDD：使用领域驱动设计
  + Ioc：使用ioc/container，做解耦、注入、依赖倒置
  + Gin
  + MySQL(Grom)
+ Vue3 + ArcoDesign

项目原型
+ 用户
  * 访客无需登陆，就能浏览文章, 登录后才能进行评论
  * 创作者登录后, 才能发布文章(Markdown编辑器)
+ 流程：发布博客, 访客可以在界面搜索并且查看博客

## 分层介绍
各目录的作用：
- /apps/：业务处理模块
  - /blog/：文章管理模块
  - /comment/：评论管理模块
  - /token/：鉴权管理模块
  - /user/：用户管理模块
    - /api/
      - /http.go/：用户对外暴露的接口
    - /impl/
      - /impl.go/：用户业务实现
    - /enum.go/：枚举
    - /interface.go/：用户接口定义
    - /model.go/：用户数据结构定义
- /common/：通用公共包
  - /meta.go/
- /conf/：配置管理
  - config.go：全局配置文件
  - load.go：加载配置
- /etc/：程序加载的配置文件目录
- /exception/：异常处理
  - /bussiness/：业务异常
  - /exception/：异常定义
- /ioc/：对象托管
- /middleware/：中间件
  - /auth.go/：鉴权
- /protocol/：协议服务
- /response/：web响应
- /test/：测试模块
- /web/：项目前端