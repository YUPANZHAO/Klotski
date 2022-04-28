package config

type MysqlConfig struct {
	User     string // 用户名
	Password string // 密码
	Host     string // 数据库服务器地址
	Port     int    // 端口号
	DB       string // 数据库名
}

func NewMysqlConfig() MysqlConfig {
	return MysqlConfig{
		User:     "root",
		Password: "xxx",
		Host:     "127.0.0.1",
		Port:     3306,
		DB:       "test",
	}
}
