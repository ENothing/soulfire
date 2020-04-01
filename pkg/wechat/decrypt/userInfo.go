package decrypt

import "soulfire/pkg/logging"

type UserInfo struct {
	HeadUrl  string
	NickName string
	Gender   int64
}

func (d *Decrypt) UserInfo(sessionKey, encryptedData, iv string) *UserInfo {

	res, err := ToDecrypt(sessionKey, encryptedData, iv)
	logging.Logging(logging.INFO, sessionKey)
	logging.Logging(logging.INFO, encryptedData)
	logging.Logging(logging.INFO, iv)
	logging.Logging(logging.INFO, res)
	logging.Logging(logging.ERR, err)
	if err != nil {

		return nil
	}

	return &UserInfo{
		HeadUrl:  (res["avatarUrl"]).(string),
		NickName: (res["nickName"]).(string),
		Gender:   int64(res["gender"].(float64)),
	}

}
