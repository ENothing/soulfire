package decrypt

type UserInfo struct {
	HeadUrl  string
	NickName string
	Gender   int64
}

func (d *Decrypt) UserInfo(sessionKey, encryptedData, iv string) *UserInfo {

	res, err := ToDecrypt(sessionKey, encryptedData, iv)
	if err != nil {
		return nil
	}

	return &UserInfo{
		HeadUrl:  (res["avatarUrl"]).(string),
		NickName: (res["nickName"]).(string),
		Gender:   int64(res["gender"].(float64)),
	}

}
