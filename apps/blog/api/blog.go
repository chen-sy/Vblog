package api

import (
	"gitee.com/chensyi/vblog/apps/blog"
	"gitee.com/chensyi/vblog/apps/user"
	"gitee.com/chensyi/vblog/middleware"
	"gitee.com/chensyi/vblog/response"
	"github.com/gin-gonic/gin"
)

// 注册路由
func (h *apiHandler) Registry(r gin.IRouter) {
	// 创建一个路由组
	v1 := r.Group("v1").Group("blogs")
	// 设置请求url对应的函数体
	v1.GET("/:search", h.SearchBlogs)
	v1.GET("/:id", h.GetBlogDetails)

	// 需要鉴权的接口
	v1.Use(middleware.NewTokenAuther().Auth)
	v1.POST("/", middleware.Required(user.ROLE_MEMBER), h.CreateBlog)
	v1.DELETE("/:id", middleware.Required(user.ROLE_MEMBER), h.DeleteBlog)
	v1.PUT("/:id", middleware.Required(user.ROLE_MEMBER), h.UpdateBlog)
	v1.PATCH("/:id", middleware.Required(user.ROLE_MEMBER), h.PatchBlog)
	v1.GET("/", middleware.Required(user.ROLE_MEMBER), h.GetBlogList)

}

// 创建博客
func (h *apiHandler) CreateBlog(c *gin.Context) {
	in := blog.NewCreateBlogRequest()
	// 将请求的json，转换成需要的对象
	err := c.BindJSON(in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	// 将http协议的请求 ---> 控制器的请求
	ins, err := h.svc.CreateBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// 删除博客
func (h *apiHandler) DeleteBlog(c *gin.Context) {
	err := h.svc.DeleteBlog(c.Request.Context(), &blog.DeleteBlogRequest{Id: c.Param("id")})
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "delete success")
}

// 更新博客(全量)
func (h *apiHandler) UpdateBlog(c *gin.Context) {
	in := blog.NewPutUpdateBlogRequest(c.Param("id"))
	err := c.BindJSON(in.CreateBlogRequest)
	if err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := h.svc.UpdateBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, ins)
}

// 更新博客(增量)
func (h *apiHandler) PatchBlog(c *gin.Context) {
	in := blog.NewPatchUpdateBlogRequest(c.Param("id"))
	err := c.BindJSON(in.CreateBlogRequest)
	if err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := h.svc.UpdateBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, ins)
}

// 获取博客详细内容
func (h *apiHandler) GetBlogDetails(c *gin.Context) {
	ins, err := h.svc.GetBlogDetails(c.Request.Context(), &blog.GetBlogDetailsRequest{Id: c.Param("id")})
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, ins)
}

// 获取创建者的博客列表
func (h *apiHandler) GetBlogList(c *gin.Context) {
	in := blog.NewGetBlogListRequest()
	in.ParsePageIndex(c.Query("page_index"))
	in.ParsePageSize(c.Query("page_size"))
	switch c.Query("blog_type") {
	case "original":
		in.BlogType.Set(blog.TYPE_ORIGINAL)
	case "reprinted":
		in.BlogType.Set(blog.TYPE_REPRINTED)
	}
	switch c.Query("visible_range") {
	case "all":
		in.VisibleRange.Set(blog.Range_ALL)
	case "own":
		in.VisibleRange.Set(blog.Range_OWN)
	}
	switch c.Query("states") {
	case "draft":
		in.States.Set(blog.STATES_DRAFT)
	case "published":
		in.States.Set(blog.STATES_PUBLISHED)
	}
	list, err := h.svc.GetBlogList(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, list)
}

// 搜索博客
func (h *apiHandler) SearchBlogs(c *gin.Context) {
	in := blog.NewSearchBlogsRequest()
	in.ParsePageIndex(c.Query("page_index"))
	in.ParsePageSize(c.Query("page_size"))
	switch c.Query("param") {
	case "title":
		in.Param = blog.QUERY_BY_TITLE
	case "author":
		in.Param = blog.QUERY_BY_AUTHOR
	}
	in.Keywords = c.Query("keywords")
	list, err := h.svc.SearchBlogs(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, list)
}
