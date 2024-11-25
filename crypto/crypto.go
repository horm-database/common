package crypto

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/horm-database/common/types"
)

// MD5Str 计算字符串的 md5 值
func MD5Str(str string) string {
	sumBytes := md5.Sum(types.StringToBytes(str))
	return hex.EncodeToString(sumBytes[:])
}

// MD5Bytes 计算字符串的 md5 值
func MD5Bytes(b []byte) []byte {
	src := md5.Sum(b)
	dst := make([]byte, hex.EncodedLen(len(src[:])))
	hex.Encode(dst, src[:])
	return dst
}
