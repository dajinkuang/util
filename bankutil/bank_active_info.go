package bankutil

import (
	"encoding/json"
	"strings"
)

// BankActiveInfo 可用银行信息
type BankActiveInfo struct {
	Name   string `json:"name"`
	Active int64  `json:"active"`
}

var __AllBankActiveInfo map[string]BankActiveInfo

func init() {
	var allBankInfoJSON = `{
    "ccb": {
        "name": "建设银行",
        "active": 1
    },
    "icbc": {
        "name": "工商银行", 
        "active": 1
    },
    "abchina": {
        "name": "农业银行",
        "active": 1
    },
    "abc": {
        "name": "农业银行",
        "active": 0
    },
    "bankcomm": {
        "name": "交通银行",
        "active": 1
    },
    "boc": {
        "name": "中国银行",
        "active": 1
    },
    "psbc": {
        "name": "邮政银行",
        "active": 1
    },
    "cebbank": {
        "name": "光大银行",
        "active": 1
    },
    "ceb": {
        "name": "光大银行",
        "active": 0
    },
    "cmbchina": {
        "name": "招商银行",
        "active": 1
    },
    "ecitic": {
        "name": "中信银行",
        "active": 1
    },
    "citic": {
        "name": "中信银行",
        "active": 0
    },
    "cmbc": {
        "name": "民生银行",
        "active": 1
    },
    "cib": {
        "name": "兴业银行",
        "active": 1
    },
    "cgb": {
        "name": "广发银行",
        "active": 1
    },
    "spdb": {
        "name": "浦发银行",
        "active": 1
    },
    "spabank": {
        "name": "平安银行",
        "active": 1
    },
    "bjbank": {
        "name": "北京银行",
        "active": 0
    }
}`
	if err := json.Unmarshal([]byte(allBankInfoJSON), &__AllBankActiveInfo); err != nil {
		panic("[初始化]json解析可用银行信息出错")
	}
}

// FetchBankActiveInfoByKey 根据银行简称获取可用银行信息
func FetchBankActiveInfoByKey(key string) *BankActiveInfo {
	if v, ok := __AllBankActiveInfo[strings.ToLower(key)]; ok {
		return &v
	}
	return nil
}
