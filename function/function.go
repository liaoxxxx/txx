package function

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

//获取固定长度的随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//获取加密后的
func GetEncPasswd(salt, passworld string) string {
	h := md5.New()
	h.Write([]byte(salt + passworld))
	//密码加密后
	return hex.EncodeToString(h.Sum(nil))
}
