// Package qrcodeutil 二维码工具包
package qrcodeutil

import (
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"

	"github.com/dajinkuang/errors"
)

// GenQRCodeBase64PNG 字符串生成二维码png图片base64
//  src    源字符串
//  width  图片宽
//  height 图片高
func GenQRCodeBase64PNG(src string, width, height int) (qrCodeBase64 string, err error) {
	if src == "" {
		return
	}
	code, err := qr.Encode(src, qr.L, qr.Unicode)
	if err != nil {
		err = errors.Wrap(err)
		return
	}
	code, err = barcode.Scale(code, width, height)
	if err != nil {
		err = errors.Wrap(err)
		return
	}
	tmp := bytes.NewBufferString("")
	if err = png.Encode(tmp, code); err != nil {
		err = errors.Wrap(err)
		return
	}
	qrCodeBase64 = base64.StdEncoding.EncodeToString(tmp.Bytes())
	return
}
