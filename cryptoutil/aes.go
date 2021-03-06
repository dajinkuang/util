package cryptoutil

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/dajinkuang/util/cryptoutil/padding"
)

func AesEncrypt(origData, key []byte) (crypted []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("panic:%v", e)
		}
	}()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = padding.PKCS7.Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted = make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) (original []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("panic:%v", e)
		}
	}()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	original, _ = padding.PKCS7.Unpadding(origData, blockSize)
	return original, nil
}
