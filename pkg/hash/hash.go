package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) string {
	hash := md5.New()

	_, err := hash.Write([]byte(s))
	if err != nil {
		return ""
	}

	return hex.EncodeToString(hash.Sum(nil))
}
