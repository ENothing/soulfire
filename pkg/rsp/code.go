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
	ActivityCateNotExits		= &Errno{10402,"活动分类不存在"}
	VideoNotExits		= &Errno{10403,"活动banner-video不存在"}
	ActivityListNotExits		= &Errno{10404,"活动不存在"}
	ActivityNotExits		= &Errno{10405,"活动不存在"}
	EnterActivityRequired		= &Errno{10406,"报名信息填写不完整"}
	CreateActivityOrderFaild		= &Errno{10407,"活动订单创建失败"}
	ActivityOrderNotExits		= &Errno{10408,"活动订单不存在"}

)



