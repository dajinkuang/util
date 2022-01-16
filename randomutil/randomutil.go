// Package randomutil 随机字符串工具包
package randomutil

import (
	"math/rand"
	"time"
)

type (
	// Random 随机对象
	Random struct {
		charset Charset
	}
	// Charset 字符集
	Charset string
)

const (
	// Alphanumeric 数字和大小写字母字符集
	Alphanumeric Charset = Alphabetic + Numeric
	// Alphabetic 大小写字母字符集
	Alphabetic Charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Numeric 数字字符集
	Numeric Charset = "0123456789"
	// Hex 十六进制字符集
	Hex Charset = Numeric + "abcdef"
)

var global = NewRandom() // 全局随机对象

// NewRandom 新建一个随机对象
func NewRandom() *Random {
	rand.Seed(time.Now().UnixNano())
	return &Random{
		charset: Alphanumeric,
	}
}

// SetCharset 设置字符集
func (r *Random) SetCharset(c Charset) {
	r.charset = c
}

// String 得到随机字符串
func (r *Random) String(length uint8) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = r.charset[rand.Int63()%int64(len(r.charset))]
	}
	return string(b)
}

// SetCharset 给global设置字符集
func SetCharset(c Charset) {
	global.SetCharset(c)
}

// String global获取随机字符串
func String(length uint8) string {
	return global.String(length)
}

// RandNumStr 获取随机数字符串
func RandNumStr(len uint8) string {
	r := NewRandom()
	r.SetCharset(Numeric)
	return r.String(len)
}
