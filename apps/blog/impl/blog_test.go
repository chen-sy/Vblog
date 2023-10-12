package impl_test

import (
	"context"
	"fmt"
	"testing"

	"gitee.com/chensyi/vblog/apps/blog"
	"gitee.com/chensyi/vblog/apps/blog/impl"
	"gitee.com/chensyi/vblog/ioc"
	"gitee.com/chensyi/vblog/test"
)

var (
	svc blog.Service
	ctx = context.Background()
)

func TestCreateBlog(t *testing.T) {
	in := blog.NewCreateBlogRequest()
	in.Title = "Test Create Blog"
	in.Tags["分类"] = "Golang"
	ins, err := svc.CreateBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
func TestDeleteUser(t *testing.T) {
	err := svc.DeleteBlog(ctx, &blog.DeleteBlogRequest{Id: 1})
	if err != nil {
		t.Fatal(err)
	}
}
func TestUpdateBlogPut(t *testing.T) {
	in := blog.NewCreateBlogRequest()
	in.Title = "我被更新了2222"
	i, err := svc.UpdateBlog(ctx, &blog.UpdateBlogRequest{Id: 1, UpdateMode: blog.UPDATE_MODE_PUT, CreateBlogRequest: in})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(i)
}
func TestUpdateBlogPatch(t *testing.T) {
	in := blog.NewCreateBlogRequest()
	in.Content = "我又被更新了"
	i, err := svc.UpdateBlog(ctx, &blog.UpdateBlogRequest{Id: 1, UpdateMode: blog.UPDATE_MODE_PATCH, CreateBlogRequest: in})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(i)
}

func TestGetBlogDetails(t *testing.T) {
	ins, err := svc.GetBlogDetails(ctx, &blog.GetBlogDetailsRequest{Id: 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestGetBlogList(t *testing.T) {
	in := blog.NewGetBlogListRequest()
	in.Keywords = "4"
	u, err := svc.GetBlogList(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestSearchBlogs(t *testing.T) {
	in := blog.NewSearchBlogsRequest()
	in.Param = blog.QUERY_BY_AUTHOR
	in.Keywords = "chensy"
	u, err := svc.SearchBlogs(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestBlogToDB(t *testing.T) {
	err := impl.MySqlAutoMigrate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestData(t *testing.T) {
	for i := 0; i < 20; i++ {
		in := blog.NewCreateBlogRequest()
		in.Title = fmt.Sprintf("Test Title %d", i+1)
		_, err := svc.CreateBlog(ctx, in)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func init() {
	test.DevelopmentSetup()
	svc = ioc.Controller().Get(blog.AppName).(blog.Service)
}
