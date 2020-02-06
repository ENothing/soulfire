package rsp

type Errno struct {
	Code    int
	Message string
}

var (
	OK = &Errno{0, "OK"}

	//2 系统
	DatabaseErr = &Errno{10201, "Database Error"}
	UploadErr   = &Errno{10202, "文件上传失败，请重试"}

	//3 用户
	LoginFailed         = &Errno{10301, "登录失败"}
	GenerateTokenErr    = &Errno{10302, "Generare Token Error"}
	InvalidToken        = &Errno{10303, "Invalid token"}
	UserNotExits        = &Errno{10304, "用户信息获取失败"}
	PleaseLogin         = &Errno{10305, "请先授权登录"}
	FeedbackParamsEmpty = &Errno{10306, "缺少参数"}
	FeedbackPostFailed  = &Errno{10307, "反馈提交失败,请重试"}
	UserInfoGotFailed   = &Errno{10308, "用户信息获取失败"}
	UserCreateFailed    = &Errno{10309, "用户信息创建失败,请重试"}

	//4活动
	BannersNotExits          = &Errno{10401, "活动banner不存在"}
	ActivityCateNotExits     = &Errno{10402, "活动分类不存在"}
	VideoNotExits            = &Errno{10403, "活动banner-video不存在"}
	ActivityListNotExits     = &Errno{10404, "活动不存在"}
	ActivityNotExits         = &Errno{10405, "活动不存在"}
	EnterActivityRequired    = &Errno{10406, "报名信息填写不完整"}
	CreateActivityOrderFaild = &Errno{10407, "活动订单创建失败"}
	ActivityOrderNotExits    = &Errno{10408, "活动订单不存在"}
	ActivityPersonIsFull    = &Errno{10408, "活动名额不足"}

	//bbs
	ArticleNotExits      = &Errno{10501, "文章不存在"}
	ArticleCreateFailed  = &Errno{10502, "文章发表失败"}
	ArticleUpdateFailed  = &Errno{10502, "文章更新失败"}
	ArticleDeleteFailed  = &Errno{10502, "文章删除失败"}
	ArticleCateGetFailed = &Errno{10503, "文章分类获取失败"}

	//评论
	ArticleCommentCreateFailed = &Errno{10601, "文章评论失败"}
	ArticleCommentListNotExits = &Errno{10602, "文章评论获取失败"}

	//商城
	GoodsListNotExits      = &Errno{10701, "商品列表获取失败"}
	GoodsNotExits          = &Errno{10702, "商品获取失败"}
	ShopOrderGoodsNotExits = &Errno{10703, "已购买人获取失败"}
	ShopGoodsSpuNotExits   = &Errno{10704, "商品规格获取失败"}

	ShopOrderCreateFailed = &Errno{10705, "订单创建失败,请重试"}
	ShopGoodsNotEnough    = &Errno{10706, "商品库存不足"}
	ShopOrderNotExits     = &Errno{10707, "订单获取失败"}
	ShopOrderCancelFailed = &Errno{10708, "订单取消失败或已付款"}

	ShopOrderRefundRejected    = &Errno{10709, "当前订单不支持退款退货"}
	ShopOrderRefundFailed      = &Errno{10710, "发起退款失败,请重试"}
	ReShopOrderRefund          = &Errno{10711, "请勿重复发起退款"}
	ShopOrderRefundNotExits    = &Errno{10712, "退款订单获取失败,请重试"}
	ShopOrderRefundPostFailed  = &Errno{10713, "退款表单提交失败"}
	ShopOrderRefundReasonEmpty = &Errno{10714, "退款理由不能为空"}

	ShopHistoryDelFailed = &Errno{10715, "历史搜索记录删除失败"}


	//地址
	AddAddressRequired         = &Errno{10801, "地址添加缺少参数"}
	AddAddressFailed           = &Errno{10802, "地址添加失败"}
	EditAddressFailed          = &Errno{10803, "地址更新失败"}
	AddressNotExits            = &Errno{10804, "地址获取失败"}
	UpdateDefaultAddressFailed = &Errno{10805, "默认地址更新失败"}
	AddressDeleteFailed        = &Errno{10806, "地址删除失败"}

	//优惠券
	CouponsListNotExits     = &Errno{10902, "优惠券获取失败"}
	UserCouponsListNotExits = &Errno{10903, "用户优惠券获取失败"}

	//关注
	FollowCancelFailed = &Errno{11001, "取消关注失败"}
	FollowedFailed     = &Errno{11002, "关注失败"}
	FollowedSelfFailed = &Errno{11003, "不能关注自己"}

	//快递
	ExpressGetFailed = &Errno{12001, "暂未获取到快递信息"}
)
