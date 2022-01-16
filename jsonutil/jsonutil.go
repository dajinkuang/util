// Package jsonutil JSON工具包
package jsonutil

import "encoding/json"

// DumpJSON 将v转为字符串
func DumpJSON(v interface{}) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}
