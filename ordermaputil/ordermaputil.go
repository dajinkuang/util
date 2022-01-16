// Package ordermaputil 有序map工具包
package ordermaputil

import (
	"bytes"
	"encoding/json"
)

// KeyValue 键值对
type KeyValue struct {
	Key   string
	Value interface{}
}

// OrderMap 定义一个有序map
type OrderMap struct {
	valueList []KeyValue
	keyMap    map[string]int
}

// NewOrderMap 新建一个OrderedMap
func NewOrderMap() *OrderMap {
	return &OrderMap{
		valueList: make([]KeyValue, 0),
		keyMap:    make(map[string]int),
	}
}

// MarshalJSON Implement the json.Marshaler interface
func (o OrderMap) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString("{")
	for i, kv := range o.valueList {
		if i != 0 {
			buf.WriteString(",")
		}
		// marshal key
		key, err := json.Marshal(kv.Key)
		if err != nil {
			return nil, err
		}
		buf.Write(key)
		buf.WriteString(":")
		// marshal value
		if tmp, ok := kv.Value.(error); ok {
			kv.Value = tmp.Error() // 对于error类型特殊处理
		}
		val, err := json.Marshal(kv.Value)
		if err != nil {
			return nil, err
		}
		buf.Write(val)
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}

// Set 设置kv
func (o *OrderMap) Set(key string, val interface{}) {
	newValue := KeyValue{Key: key, Value: val}
	if v, ok := o.keyMap[key]; ok {
		o.valueList[v] = newValue
		return
	}
	o.valueList = append(o.valueList, newValue)
	o.keyMap[key] = len(o.valueList) - 1
}

// Get 根据k取得v
func (o OrderMap) Get(key string) (val interface{}, ok bool) {
	if v, ok := o.keyMap[key]; ok {
		return o.valueList[v].Value, true
	}
	return nil, false
}

// AddValues add values
func (o *OrderMap) AddValues(values *OrderMap) {
	if values == nil {
		return
	}
	for _, v := range values.valueList {
		o.Set(v.Key, v.Value)
	}
}
