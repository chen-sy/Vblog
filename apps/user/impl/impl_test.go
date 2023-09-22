package impl_test

import (
	"context"
	"crypto/md5"
	"encoding/hex"
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
	u, err := userSvc.CreateUser(ctx, &user.CreateUserRequest{UserName: "test001", PassWord: "123456"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestGetSingleUser(t *testing.T) {
	u, err := userSvc.GetSingleUser(ctx, &user.GetSingleUserRequest{Param: user.QUERY_BY_USERNAME, Value: "chensy"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestDeleteUser(t *testing.T) {
	err := userSvc.DeleteUser(ctx, &user.DeleteUserRequest{Id: 1})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateUser(t *testing.T) {
	i, err := userSvc.UpdateUser(ctx, &user.UpdateUserRequest{Id: 5, PassWord: "666666"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(i)
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

func TestMD5(t *testing.T) {
	h := md5.New()
	_, err := h.Write([]byte("123456"))
	if err != nil {
		t.Fatal(err)
	}
	s := hex.EncodeToString(h.Sum(nil))
	t.Log(s)
}
