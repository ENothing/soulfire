package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"soulfire/utils"
)

type Decrypt struct{}

func ToDecrypt(sessionKey, encryptData, iv string) (map[string]interface{}, error) {

	decodeBytes, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, err
	}
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	dataBytes, err := AesDecrypt(decodeBytes, sessionKeyBytes, ivBytes)

	m := utils.JsonDecode(string(dataBytes))

	return m, nil

}

func AesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	//获取的数据尾端有'/x0e'占位符,去除它
	for i, ch := range origData {
		if ch == '\x03' {
			origData[i] = ' '
		}
	}
	//{"phoneNumber":"15082726017","purePhoneNumber":"15082726017","countryCode":"86","watermark":{"timestamp":1539657521,"appid":"wx4c6c3ed14736228c"}}//<nil>
	return origData, nil
}
