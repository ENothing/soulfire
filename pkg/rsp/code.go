package rsp

type Errno struct {
	Code int
	Message string
}

var (



	OK				= &Errno{0,"OK"}


	//2 系统
	DatabaseErr		= &Errno{10101,"Database Error"}





	//3 用户
	LoginFailed				= &Errno{10301,"登录失败"}
	GenerateTokenErr				= &Errno{10302,"Generare Token Error"}
	InvalidToken				= &Errno{10303,"Invalid token"}

	//4活动
	BannersNotExits		= &Errno{10401,"活动banner不存在"}
	ActivityCateNotExits		= &Errno{10401,"活动分类不存在"}

)



