package wechat

import (
	"soulfire/pkg/wechat/decrypt"
	"soulfire/pkg/wechat/login"
	"soulfire/pkg/wechat/token"
)

type Wc struct{}

func (wc *Wc) Login() *login.Login {
	return new(login.Login)
}

func (wc *Wc) Token() *token.Token {
	return new(token.Token)
}

func (wc *Wc) Decrypt() *decrypt.Decrypt {
	return new(decrypt.Decrypt)
}
