package request

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/maczh/mgin/logs"
	"strings"
	"time"
)

func hmacSHA1(key string, data string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func GetSign(method, uri, accessKeyId, accessKeySecret string) map[string]string {
	uri = strings.ReplaceAll(uri, "zstack/", "/")

	data := map[string]string{
		"accesskey_id":     accessKeyId,
		"accesskey_secret": accessKeySecret,
		"method":           method,
		"date":             time.Now().Format("Mon, 02 Jan 2006 15:04:05") + " PRC",
		"uri":              uri,
	}
	str := fmt.Sprintf("%s\n%s\n%s", method, data["date"], uri)
	logs.Debug("签名源数据:{}", str)
	sign := hmacSHA1(accessKeySecret, str)
	signStr := fmt.Sprintf("ZStack %s:%s", accessKeyId, sign)

	result := map[string]string{
		"Authorization": signStr,
		"date":          data["date"],
		"Content-type":  "application/json; charset=utf-8",
	}
	return result
}

func IsList(params interface{}) bool {
	switch params.(type) {
	case []interface{}:
		return true
	}
	return false
}

func IsBool(params interface{}) bool {
	switch params.(type) {
	case bool:
		return true
	}
	return false
}
