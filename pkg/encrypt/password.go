package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "jannan"

func GetPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
