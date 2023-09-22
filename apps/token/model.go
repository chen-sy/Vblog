package token

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/xid"
)

func NewToken() *Token {
	return &Token{
		// xid是一个全局唯一的id生成器库，使用xid生成一个UUID的字符串
		AccessToken:           xid.New().String(),
		AccessTokenExpiredAt:  time.Now().Add(2 * time.Hour).Unix(),
		RefreshToken:          xid.New().String(),
		RefreshTokenExpiredAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
		CreatedAt:             time.Now().Unix(),
	}
}

type Token struct {
	// 用户ID
	UserID int64 `json:"user_id"`
	// 颁发给用户的访问令牌（用户需要携带 token 才能访问服务器）
	AccessToken string `json:"access_token"`
	// token 的过期时间（暂定2h，程序写死，后续可使用配置），单位是秒
	AccessTokenExpiredAt int64 `json:"access_token_expired_at"`
	// 用来在 token 过期以后重新获取 accessToken
	RefreshToken string `json:"refresh_token"`
	// 刷新 token 的过期时间（7d），单位是秒
	RefreshTokenExpiredAt int64 `json:"refresh_token_expired_at"`
	// 创建时间
	CreatedAt int64 `json:"created_at"`
}

// 判断token是否过期
func (t *Token) IsExpired() error {
	// 将时间戳转成当前时间格式
	t1 := time.Unix(t.AccessTokenExpiredAt, 0)
	// 使用Since函数来计算指定时间与当前时间的时间间隔，按秒显示
	expSeconds := time.Since(t1).Seconds()
	// 正数说明当前时间已经超过了过期时间，即token已过期
	if expSeconds > 0 {
		return fmt.Errorf("token已过期")
	}
	return nil
}

func (t *Token) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}
