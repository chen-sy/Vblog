package impl_test

import (
	"context"
	"testing"

	"gitee.com/chensyi/vblog/apps/token"
	"gitee.com/chensyi/vblog/apps/token/impl"
	userImpl "gitee.com/chensyi/vblog/apps/user/impl"
	"gitee.com/chensyi/vblog/test"
)

var (
	tokenSvc *impl.TokenServiceImpl
	ctx      = context.Background()
)

func TestLogin(t *testing.T) {
	u, err := tokenSvc.Login(ctx, &token.LoginRequest{UserName: "chensy", PassWord: "123456"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestLogout(t *testing.T) {
	err := tokenSvc.Logout(ctx, &token.LogoutRequest{AccessToken: "ck6ve5f067qhpj63pbeg", RefreshToken: "ck6ve5f067qhpj63pbf0"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestValidateToken(t *testing.T) {
	u, err := tokenSvc.ValidateToken(ctx, &token.ValidateToken{AccessToken: "ck6ve5f067qhpj63pbeg"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestUserToDB(t *testing.T) {
	err := impl.MySqlAutoMigrate()
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	test.DevelopmentSetup()
	tokenSvc = impl.NewTokenServiceImpl(userImpl.NewUserServiceImpl())

}
