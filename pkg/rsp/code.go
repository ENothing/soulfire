package rsp

type Errno struct {
	Code int
	Message string
}

var (
	OK				= &Errno{0,"OK"}
	DatabaseErr		= &Errno{10201,"Database Error"}
	GenerateTokenErr				= &Errno{10202,"Generare Token Error"}
	InvalidToken				= &Errno{10203,"Invalid token"}

	LoginFailed				= &Errno{10301,"登录失败"}

)



