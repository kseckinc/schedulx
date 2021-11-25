package tool

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"unicode/utf8"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func PickDomainFromUrl(s string) (string, error) {
	if !strings.Contains(s, "http") {
		s = "https://" + s
	}

	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	return u.Host, nil
}

func ToJson(val interface{}) string {
	bytes, _ := jsoniter.Marshal(val)
	return string(bytes)
}

func Interface2String(value interface{}) string {
	key := ""
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	case json.Number:
		key = value.(json.Number).String()
	}

	return key
}

// StrAppend 字符串拼接
func StrAppend(str1 string, str2 ...string) string {
	var builder strings.Builder
	builder.WriteString(str1)
	for _, str := range str2 {
		builder.WriteString(str)
	}
	return builder.String()
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func SubStr(s string, length int) string {
	var size, n int
	for i := 0; i < length && n < len(s); i++ {
		_, size = utf8.DecodeRuneInString(s[n:])
		n += size
	}

	return s[:n]
}
