package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "jannan"

// GetPassword 哈希获取密码MD5
func GetPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

// CheckPassword 判断密码是否一致
func CheckPassword(password string, newPassword string) bool {
	//password 原始密码
	//newPassword 登陆密码
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(newPassword))) == password
}
