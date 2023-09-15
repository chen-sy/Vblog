package impl_test

import (
	"context"
	"testing"

	"gitee.com/chensyi/vblog/apps/user"
	"gitee.com/chensyi/vblog/apps/user/impl"
	"gitee.com/chensyi/vblog/test"
)

var (
	userSvc *impl.UserServiceImpl
	ctx     = context.Background()
)

func TestCreateUser(t *testing.T) {
	u, err := userSvc.CreateUser(ctx, &user.CreateUserRequest{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestDeleteUser(t *testing.T) {
	err := userSvc.DeleteUser(ctx, &user.DeleteUserRequest{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserToDB(t *testing.T) {
	err := impl.MySqlAutoMigrate()
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	test.DevelopmentSetup()
	userSvc = impl.NewUserServiceImpl()

}
