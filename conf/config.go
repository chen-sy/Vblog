package conf

import (
	"encoding/json"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局配置
type Config struct {
	MySQL *MySQL `json:"mysql" toml:"mysql"`
}

func DefaultConfig() *Config {
	return &Config{
		MySQL: &MySQL{
			User:   "admin",
			Pass:   "123456",
			Host:   "127.0.0.1",
			Port:   3306,
			DbName: "vblog",
		},
	}
}

func (c *Config) string() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type MySQL struct {
	User   string `json:"user" toml:"user" env:"MYSQL_USER"`
	Pass   string `json:"pass" toml:"pass" env:"MYSQL_PASS"`
	Host   string `json:"host" toml:"host" env:"MYSQL_HOST"`
	Port   int    `json:"port" toml:"port" env:"MYSQL_PORT"`
	DbName string `json:"dbname" toml:"dbname" env:"MYSQL_DBNAME"`
	// 缓存一个连接对象，避免占用数据库连接池
	conn *gorm.DB
	// 加锁避免数据竞争
	lock sync.Mutex
}

func (m *MySQL) DSN() string {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User,
		m.Pass,
		m.Host,
		m.Port,
		m.DbName,
	)
}

func (m *MySQL) GetConn() *gorm.DB {
	if m.conn == nil {
		// 加锁
		m.lock.Lock()
		// 释放锁
		defer m.lock.Unlock()

		db, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		m.conn = db
	}

	return m.conn
}
