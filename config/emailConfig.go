package config

type EmailConfig struct {
	Host     string // 邮箱服务器
	Port     int    // 端口号
	UserName string // 发送邮箱
	Password string // 授权码
	Title    string // 邮件标题
}

func NewEmailConfig() EmailConfig {
	return EmailConfig{
		Host:     "smtp.qq.com",
		Port:     465,
		UserName: "xxx@xxx.com",
		Password: "xxx",
		Title:    "来自华容道解密网站的一封验证邮件",
	}
}
