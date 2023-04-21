package Data

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
)

// InterfaceToStr 转字符串
func InterfaceToStr(i interface{}) string {
	switch i.(type) {
	case int64:
		return fmt.Sprintf("%d", i)
	case string:
		return i.(string)
	case float64:
		return fmt.Sprintf("%.0f", i)
	default:
		return ""
	}
}

// StrToMd5 字符串加密md5
func StrToMd5(data string, size int8) string {
	t := md5.New()
	_, err := io.WriteString(t, data)
	if err != nil {
		return ""
	}
	r := fmt.Sprintf("%x", t.Sum(nil))
	return r[:size]
}

// GetUUid 获取随机uuid
func GetUUid(size int8) string {
	_, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	//uuid做md5转换
	u4 := uuid.New()
	return StrToMd5(u4.String(), size)
}
