package token

import "context"

const (
	AppName = "token"
)

type Service interface {
	// 登录接口（颁发token）
	Login(context.Context, *LoginRequest) (*Token, error)
	// 退出接口（销毁token）
	Logout(context.Context, *LogoutRequest) error
	// 校验token
	ValidateToken(context.Context, *ValidateToken) (*Token, error)
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

type LoginRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func NewLogoutRequest() *LogoutRequest {
	return &LogoutRequest{}
}

type LogoutRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewValidateToken(token string) *ValidateToken {
	return &ValidateToken{
		AccessToken: token,
	}
}

type ValidateToken struct {
	AccessToken string `json:"access_token"`
}
