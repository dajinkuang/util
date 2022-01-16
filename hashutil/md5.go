package hashutil

import (
	"crypto/md5"
)

// MD5 -
func MD5(data []byte) []byte {
	sum := md5.Sum(data)
	return sum[:]
}
