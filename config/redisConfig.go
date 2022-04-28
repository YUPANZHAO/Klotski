package config

type RedisConfig struct {
	Address  string // Redis服务器地址
	Password string // 密码
	DB       int    // 数据库编号
}

func NewRedisConfig() RedisConfig {
	return RedisConfig{
		Address:  "localhost:6379",
		Password: "",
		DB:       0,
	}
}
