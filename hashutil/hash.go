package hashutil

import (
	"crypto/sha1"
	"crypto/sha256"
	"hash/crc32"
)

// SHA1Sum -
func SHA1Sum(data []byte) []byte {
	sum := sha1.Sum(data)
	return sum[:]
}

// CRC32Sum -
func CRC32Sum(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

// SHA256Sum -
func SHA256Sum(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}
