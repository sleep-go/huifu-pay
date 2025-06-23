package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
)

type StringDecoder[T any] interface {
	Decode() T
	Encode() string
}

type StringObject[T any] string

func (o StringObject[T]) Encode() string {
	marshal, _ := json.Marshal(o)
	return string(marshal)
}
func (o StringObject[T]) Decode() T {
	var res T
	_ = json.Unmarshal([]byte(o), &res)
	return res
}
func HandleResponse[T any](resp map[string]any) (out *T, raw string, err error) {
	marshal, err := json.Marshal(resp)
	if err != nil {
		return nil, "", err
	}
	err = json.Unmarshal(marshal, &out)
	if err != nil {
		return nil, "", err
	}
	return out, string(marshal), nil
}

func StructToMapClean(v any) map[string]any {
	out := make(map[string]interface{})
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	t := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := t.Field(i)
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "-" || jsonTag == "" {
			continue
		}
		key := jsonTag
		if idx := len(key); idx > 0 && key[idx-1] == ',' {
			key = key[:idx-1]
		}
		if field.IsZero() {
			continue
		}
		out[key] = field.Interface()
	}
	return out
}

// ParseCallbackRespData 解析回调通知
func ParseCallbackRespData[T any](r *http.Request, msc *BsPaySdk.MerchSysConfig) (resp *T, raw string, err error) {
	res, body, err := parseQueryToMap(r)
	if err != nil {
		return nil, "", err
	}
	// 验签
	sign, err := BsPaySdk.RsaSignVerify(res["sign"].(string), res["resp_data"].(string), msc)
	if err != nil {
		return nil, body, err
	}
	if !sign {
		return nil, body, errors.New("check signature error")
	}
	return HandleResponse[T](res)
}
func ParseCallbackData[T any](r *http.Request, msc *BsPaySdk.MerchSysConfig) (resp *T, raw string, err error) {
	res, body, err := parseQueryToMap(r)
	if err != nil {
		return nil, "", err
	}
	// 验签
	sign, err := BsPaySdk.RsaSignVerify(res["sign"].(string), res["data"].(string), msc)
	if err != nil {
		return nil, body, err
	}
	if !sign {
		return nil, body, errors.New("check signature error")
	}
	return HandleResponse[T](res)
}
func parseQueryToMap(r *http.Request) (res map[string]any, raw string, err error) {
	if r.Method != http.MethodPost {
		return nil, "", errors.New("method Not Allowed")
	}
	// 读取原始 body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, "", err
	}
	defer r.Body.Close()
	params, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, "", errors.New(fmt.Sprintf("解析参数字符串失败: %s", err))
	}
	res = make(map[string]any)
	for key, values := range params {
		if len(values) > 0 {
			res[key] = values[0]
		}
	}
	return res, string(body), nil
}

func Encrypt(content string, msc *BsPaySdk.MerchSysConfig) (string, error) {
	pubKey, err := getPublicKey(msc)
	if err != nil {
		BsPaySdk.BspayPrintln("error in get public key: " + err.Error())
		return "", err
	} else if pubKey == nil {
		BsPaySdk.BspayPrintln("error in get public key")
		return "", errors.New("error in get public key")
	}
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(content))
	if err != nil {
		BsPaySdk.BspayPrintln("check sign error: " + err.Error())
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherText), nil
}
func Decrypt(content string, msc *BsPaySdk.MerchSysConfig) (string, error) {
	privateKey, err := getPrivateKey(msc)
	if err != nil {
		BsPaySdk.BspayPrintln("error in get private key: " + err.Error())
		return "", err
	} else if privateKey == nil {
		BsPaySdk.BspayPrintln("error in get private key")
		return "", errors.New("error in get private key")
	}
	decodeString, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		BsPaySdk.BspayPrintln("base64 DecodeString error: " + err.Error())
		return "", err
	}
	cipherText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decodeString)
	if err != nil {
		BsPaySdk.BspayPrintln("DecryptPKCS1v15 error: " + err.Error())
		return "", err
	}
	return string(cipherText), nil
}
func getPublicKey(msc *BsPaySdk.MerchSysConfig) (*rsa.PublicKey, error) {
	pubKeyString := msc.RsaHuifuPublicKey
	if pubKeyString == "" {
		BsPaySdk.BspayPrintln("unknow public key data")
		return nil, errors.New("unknow public key data")
	}

	// 加头尾
	if !strings.Contains(pubKeyString, "-----BEGIN PUBLIC KEY-----") {
		pubKeyString = "-----BEGIN PUBLIC KEY-----\n" + pubKeyString
	}
	if !strings.Contains(pubKeyString, "-----END PUBLIC KEY-----") {
		pubKeyString = pubKeyString + "\n-----END PUBLIC KEY-----"
	}

	// 拦截处理 key 时可能的错误
	// 拦截到以后，两个返回参数都是空了
	defer func() {
		err := recover()
		if err != nil {
			BsPaySdk.BspayPrintln("error in get public key")
		}
	}()
	keyByts, _ := pem.Decode([]byte(pubKeyString))
	pubKey, err := x509.ParsePKIXPublicKey(keyByts.Bytes)
	if err != nil {
		BsPaySdk.BspayPrintln("error in get public key: " + err.Error())
		return nil, err
	}

	// 正常返回
	return pubKey.(*rsa.PublicKey), nil
}
func getPrivateKey(msc *BsPaySdk.MerchSysConfig) (*rsa.PrivateKey, error) {
	priKeyString := msc.RsaMerchPrivateKey
	if priKeyString == "" {
		BsPaySdk.BspayPrintln("unknow private key data")
		return nil, errors.New("unknow private key data")
	}

	// 加头尾
	if !strings.Contains(priKeyString, "-----BEGIN PRIVATE KEY-----") {
		priKeyString = "-----BEGIN PRIVATE KEY-----\n" + priKeyString
	}
	if !strings.Contains(priKeyString, "-----END PRIVATE KEY-----") {
		priKeyString = priKeyString + "\n-----END PRIVATE KEY-----"
	}

	// 拦截处理 key 时可能的错误
	// 拦截到以后，两个返回参数都是空了
	defer func() {
		err := recover()
		if err != nil {
			BsPaySdk.BspayPrintln("error in get private key")
		}
	}()
	keyByts, _ := pem.Decode([]byte(priKeyString))
	privateKey, err := x509.ParsePKCS8PrivateKey(keyByts.Bytes)
	if err != nil {
		BsPaySdk.BspayPrintln("error in get private key: " + err.Error())
		return nil, err
	}
	// 正常返回
	return privateKey.(*rsa.PrivateKey), nil
}
