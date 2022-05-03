package config

type AlipayConfig struct {
	Host            string  //支付宝接口地址
	AppId           string  //应用ID
	ReturnUrl       string  //付款成功后的同步回调地址
	NotifyUrl       string  //付款成功后的异步回调地址
	PayAmount       float64 //付款金额，单位:元
	OrderName       string  //订单标题
	SignType        string  //签名算法类型，支持RSA2和RSA，推荐使用RSA2
	RsaPrivateKey   string  //应用私钥
	AlipayPublicKey string  //支付宝公钥
}

func NewAlipayConfig() AlipayConfig {
	return AlipayConfig{
		Host:            "https://openapi.alipaydev.com/gateway.do",
		AppId:           "xxx",
		ReturnUrl:       "http://127.0.0.1:8080/callback",
		NotifyUrl:       "http://127.0.0.1:8080/notify",
		PayAmount:       1.00,
		OrderName:       "华融道解密网站",
		SignType:        "RSA2",
		RsaPrivateKey:   "xxx",
		AlipayPublicKey: "xxx",
	}
}
