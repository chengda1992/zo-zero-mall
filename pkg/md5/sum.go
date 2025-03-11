package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Sum(data string) string {
	//计算data的哈希值
	hash := md5.Sum([]byte(data))
	//使用hex.EncodeToString()将哈希值转换为字符串
	return hex.EncodeToString(hash[:])
}
