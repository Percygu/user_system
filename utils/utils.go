package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Contains(source []string, tg string) bool {
	for _, s := range source {
		if s == tg {
			return true
		}
	}
	return false
}

func Md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	str := hex.EncodeToString(h.Sum(nil))
	return str
}

//func GenerateSession(uname string) string {
//    return Md5String(fmt.Sprintf("%s:%d", uname, rand.Intn(999999)))
//}

func GenerateSession(userName string) string {
	return Md5String(fmt.Sprintf("%s:%s", userName, "session"))
}
