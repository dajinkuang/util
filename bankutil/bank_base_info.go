package bankutil

import (
	"encoding/json"
	"strings"
)

// BankBaseInfo 银行基本信息
type BankBaseInfo struct {
	Name       string `json:"name"`
	Bank       string `json:"bank"`
	Logo       string `json:"logo"`
	BankCode   string `json:"bank_code"`
	BankBranch string `json:"bank_branch"`
}

var __AllBankBaseInfo map[string]BankBaseInfo

func init() {
	var allBankBaseInfoJSON string = `{
	"ccb": {
		"name": "建设银行",
		"bank": "ccb",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-ccb.png",
		"bank_code": "105100000017",
		"bank_branch": "中国建设银行总行(不受理个人业务)"
	},
	"icbc": {
		"name": "工商银行",
		"bank": "icbc",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-icbc.png",
		"bank_code": "102100099996",
		"bank_branch": "中国工商银行总行清算中心"
	},
	"abchina": {
		"name": "农业银行",
		"bank": "abchina",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-abchina.png",
		"bank_code": "103100000026",
		"bank_branch": "中国农业银行资金清算中心"
	},
	"bankcomm": {
		"name": "交通银行",
		"bank": "bankcomm",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-bankcomm.png",
		"bank_code": "301290000007",
		"bank_branch": "交通银行"
	},
	"boc": {
		"name": "中国银行",
		"bank": "boc",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-boc.png",
		"bank_code": "104100000004",
		"bank_branch": "中国银行总行"
	},
	"cebbank": {
		"name": "光大银行",
		"bank": "cebbank",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-cebbank.png",
		"bank_code": "303100000006",
		"bank_branch": "中国光大银行"
	},
	"cmbchina": {
		"name": "招商银行",
		"bank": "cmbchina",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-cmbchina.png",
		"bank_code": "308584000013",
		"bank_branch": "招商银行"
	},
	"ecitic": {
		"name": "中信银行",
		"bank": "ecitic",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-ecitic.png",
		"bank_code": "302100011000",
		"bank_branch": "中信银行总行管理部（不受理储蓄业务）"
	},
	"cmbc": {
		"name": "民生银行",
		"bank": "cmbc",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-cmbc.png",
		"bank_code": "305100000013",
		"bank_branch": "中国民生银行总行"
	},
	"cib": {
		"name": "兴业银行",
		"bank": "cib",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-cib.png",
		"bank_code": "309391000011",
		"bank_branch": "兴业银行总行"
	},
	"cgb": {
		"name": "广发银行",
		"bank": "cgb",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-cgb.png",
		"bank_code": "306581000003",
		"bank_branch": "广发银行股份有限公司"
	},
	"spdb": {
		"name": "浦发银行",
		"bank": "spdb",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-spdb.png",
		"bank_code": "310290000013",
		"bank_branch": "上海浦东发展银行"
	},
	"spabank": {
		"name": "平安银行",
		"bank": "spabank",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-spabank.png",
		"bank_code": "307584007998",
		"bank_branch": "平安银行"
	},
	"psbc": {
		"name": "邮政银行",
		"bank": "psbc",
		"logo": "https://cdn.qingsongchou.com/app/bank-icon/bank-psbc.png",
		"bank_code": "403100000004",
		"bank_branch": "中国邮政储蓄银行总行"
	}
}`
	if err := json.Unmarshal([]byte(allBankBaseInfoJSON), &__AllBankBaseInfo); err != nil {
		panic("[初始化]json解析银行基本信息出错")
	}
}

// FetchBankBaseInfoByKey 根据银行简称获取银行基本信息
func FetchBankBaseInfoByKey(key string) *BankBaseInfo {
	if v, ok := __AllBankBaseInfo[strings.ToLower(key)]; ok {
		return &v
	}
	return nil
}
