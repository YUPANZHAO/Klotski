package model

import (
	"KlotskiWeb/config"
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type AlipayService struct {
	Host            string          //支付宝接口地址
	AppId           string          //应用ID
	ReturnUrl       string          //付款成功后的同步回调地址
	NotifyUrl       string          //付款成功后的异步回调地址
	Charset         string          //字符编码
	RsaPrivateKey   *rsa.PrivateKey //应用私钥
	AlipayPublicKey *rsa.PublicKey  //支付宝公钥
	TotalFee        float64         //付款金额，单位:元
	OutTradeNo      string          //订单编号
	OrderName       string          //订单标题
}

func NewAlipayService() AlipayService {
	conf := config.NewAlipayConfig()
	privateKey, _ := ParsePrivateKey(FormatPrivateKey(conf.RsaPrivateKey))
	alipayPublicKey, _ := ParsePublicKey(FormatPublicKey(conf.AlipayPublicKey))
	return AlipayService{
		Host:            conf.Host,
		AppId:           conf.AppId,
		ReturnUrl:       conf.ReturnUrl,
		NotifyUrl:       conf.NotifyUrl,
		Charset:         "utf-8",
		RsaPrivateKey:   privateKey,
		AlipayPublicKey: alipayPublicKey,
		TotalFee:        conf.PayAmount,
		OutTradeNo:      Uniqid(),
		OrderName:       conf.OrderName,
	}
}

//生成订单号
func Uniqid() string {
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", "", now.Unix(), now.UnixNano()%0x100000)
}

func ParsePrivateKey(data []byte) (key *rsa.PrivateKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, errors.New("private key failed to load")
	}

	var priInterface interface{}
	priInterface, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := priInterface.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key failed to load")
	}

	return key, err
}

func ParsePublicKey(data []byte) (key *rsa.PublicKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, errors.New("alipay public key failed to load")
	}

	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("alipay public key failed to load")
	}

	return key, err
}

func FormatPublicKey(raw string) []byte {
	return formatKey(raw, "-----BEGIN PUBLIC KEY-----", "-----END PUBLIC KEY-----", 64)
}

func FormatPrivateKey(raw string) []byte {
	return formatKey(raw, "-----BEGIN RSA PRIVATE KEY-----", "-----END RSA PRIVATE KEY-----", 64)
}

func formatKey(raw, prefix, suffix string, lineCount int) []byte {
	if raw == "" {
		return nil
	}
	raw = strings.Replace(raw, prefix, "", 1)
	raw = strings.Replace(raw, suffix, "", 1)
	raw = strings.Replace(raw, " ", "", -1)
	raw = strings.Replace(raw, "\n", "", -1)
	raw = strings.Replace(raw, "\r", "", -1)
	raw = strings.Replace(raw, "\t", "", -1)

	var sl = len(raw)
	var c = sl / lineCount
	if sl%lineCount > 0 {
		c = c + 1
	}

	var buf bytes.Buffer
	buf.WriteString(prefix + "\n")
	for i := 0; i < c; i++ {
		var b = i * lineCount
		var e = b + lineCount
		if e > sl {
			buf.WriteString(raw[b:])
		} else {
			buf.WriteString(raw[b:e])
		}
		buf.WriteString("\n")
	}
	buf.WriteString(suffix)
	return buf.Bytes()
}

func (this *AlipayService) DoPay() (string, error) {
	//请求参数
	var bizContent = make(map[string]string)
	bizContent["out_trade_no"] = this.OutTradeNo
	bizContent["product_code"] = "FAST_INSTANT_TRADE_PAY"
	bizContent["total_amount"] = strconv.FormatFloat(float64(this.TotalFee), 'f', 2, 64) //2表示保留2位小数
	bizContent["subject"] = this.OrderName
	bizContentJson, err := json.Marshal(bizContent)
	if err != nil {
		return "", errors.New("json.Marshal: " + err.Error())
	}

	//公共参数
	var m = make(map[string]string)
	m["app_id"] = this.AppId
	m["method"] = "alipay.trade.page.pay" //接口名称
	m["format"] = "JSON"
	m["return_url"] = this.ReturnUrl
	m["charset"] = this.Charset
	m["sign_type"] = "RSA2"
	m["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	m["version"] = "1.0"
	m["notify_url"] = this.NotifyUrl
	m["biz_content"] = string(bizContentJson)

	//获取签名
	sign := this.genSign(m)
	m["sign"] = sign
	return this.buildRequestForm(m), nil
}

// GenSign 产生签名
func (this *AlipayService) genSign(m map[string]string) string {
	var data []string
	var encryptedBytes []byte
	for k, v := range m {
		if v != "" && k != "sign" {
			data = append(data, fmt.Sprintf(`%s=%s`, k, v))
		}
	}
	sort.Strings(data)
	signData := strings.Join(data, "&")
	s := sha256.New()
	_, err := s.Write([]byte(signData))
	if err != nil {
		panic(err)
	}
	hashByte := s.Sum(nil)
	hashs := crypto.SHA256
	rsaPrivateKey := this.RsaPrivateKey
	if encryptedBytes, err = rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, hashs, hashByte); err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(encryptedBytes)
}

func (this *AlipayService) buildRequestForm(m map[string]string) string {
	var buf []string
	buf = append(buf, fmt.Sprintf("正在跳转至支付页面...<form id='alipaysubmit' name='alipaysubmit' action='%s?charset=%s' method='POST'>", this.Host, this.Charset))
	for k, v := range m {
		buf = append(buf, fmt.Sprintf("<input type='hidden' name='%s' value='%s'/>", k, v))
	}
	buf = append(buf, "<input type='submit' value='ok' style='display:none;''></form>")
	buf = append(buf, "<script>document.forms['alipaysubmit'].submit();</script>")
	return strings.Join(buf, "")
}

func (this *AlipayService) VerifySign(data url.Values) (ok bool, err error) {
	return verifySign(data, this.AlipayPublicKey)
}

func verifySign(data url.Values, key *rsa.PublicKey) (ok bool, err error) {
	sign := data.Get("sign")

	var keys = make([]string, 0, 0)
	for key := range data {
		if key == "sign" || key == "sign_type" {
			continue
		}
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var pList = make([]string, 0, 0)
	for _, key := range keys {
		pList = append(pList, key+"="+data.Get(key))
	}
	var s = strings.Join(pList, "&")

	return verifyData([]byte(s), sign, key)
}

func verifyData(data []byte, sign string, key *rsa.PublicKey) (ok bool, err error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	var h = crypto.SHA256.New()
	h.Write(data)
	var hashed = h.Sum(nil)

	if err = rsa.VerifyPKCS1v15(key, crypto.SHA256, hashed, signBytes); err != nil {
		return false, err
	}
	return true, nil
}
