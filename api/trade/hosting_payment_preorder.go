package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeHostingPaymentPreorderH5 H5、PC预下单接口
// POST https://api.huifu.com/v2/trade/hosting/payment/preorder
//
// 应用场景
// 场景1-H5：接入支付托管接口并获取链接，通过主动跳转链接完成支付。
// 场景2-PC：接入支付托管接口并获取链接，通过新窗口跳转链接完成支付。点击链接体验
// 场景3-动态二维码：接入支付托管接口并获取链接转换成二维码，支持付款人用手机微信、支付宝、云闪付、浏览器扫码支付；
func (t *Trade) V2TradeHostingPaymentPreorderH5(req V2TradeHostingPaymentPreorderH5Request) (res *V2TradeHostingPaymentPreorderH5Response, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeHostingPaymentPreorderH5Request(BsPaySdk.V2TradeHostingPaymentPreorderH5Request{
		ReqDate:      req.ReqDate,
		ReqSeqId:     req.ReqSeqId,
		HuifuId:      req.HuifuId,
		TransAmt:     req.TransAmt,
		GoodsDesc:    req.GoodsDesc,
		PreOrderType: "1",
		HostingData:  req.HostingData.Encode(),
		ExtendInfos:  common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeHostingPaymentPreorderH5Response](resp)
}

type V2TradeHostingPaymentPreorderH5Request struct {
	ReqDate     string                           `json:"req_date"`     // 请求日期
	ReqSeqId    string                           `json:"req_seq_id"`   // 请求流水号
	HuifuId     string                           `json:"huifu_id"`     // 商户号
	TransAmt    string                           `json:"trans_amt"`    // 交易金额
	GoodsDesc   string                           `json:"goods_desc"`   // 商品描述
	HostingData common.StringObject[HostingData] `json:"hosting_data"` // 半支付托管扩展参数集合
	ExtendInfos V2TradeHostingPaymentPreorderExtendInfo
}
type V2TradeHostingPaymentPreorderH5Response struct {
	Data struct {
		RespCode     string                           `json:"resp_code"`
		RespDesc     string                           `json:"resp_desc"`
		ReqDate      string                           `json:"req_date"`
		ReqSeqId     string                           `json:"req_seq_id"`
		HuifuId      string                           `json:"huifu_id"`
		PreOrderType string                           `json:"pre_order_type"`
		PreOrderId   string                           `json:"pre_order_id"`
		GoodsDesc    string                           `json:"goods_desc"`
		JumpUrl      string                           `json:"jump_url"` //支付跳转链接	String
		UsageType    string                           `json:"usage_type"`
		TransType    common.TradeType                 `json:"trans_type"`
		HostingData  common.StringObject[HostingData] `json:"hosting_data"`
		CurrentTime  string                           `json:"current_time"`
		TimeExpire   string                           `json:"time_expire"`
	} `json:"data"`
	Sign string `json:"sign"`
}

// V2TradeHostingPaymentPreorderAli 支付宝小程序预下单接口
// POST https://api.huifu.com/v2/trade/hosting/payment/preorder
//
// 应用场景
// 支付宝小程序预下单接口，可以满足商户或服务商在app拉起支付宝支付的需求。
func (t *Trade) V2TradeHostingPaymentPreorderAli(req V2TradeHostingPaymentPreorderAliRequest) (res *V2TradeHostingPaymentPreorderAliResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeHostingPaymentPreorderAliRequest(BsPaySdk.V2TradeHostingPaymentPreorderAliRequest{
		HuifuId:      req.HuifuId,
		ReqDate:      req.ReqDate,
		ReqSeqId:     req.ReqSeqId,
		PreOrderType: "2",
		TransAmt:     req.TransAmt,
		GoodsDesc:    req.GoodsDesc,
		AppData:      req.AppData.Encode(),
		ExtendInfos:  common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeHostingPaymentPreorderAliResponse](resp)
}

type V2TradeHostingPaymentPreorderAliRequest struct {
	HuifuId     string                       `json:"huifu_id"`   // 商户号
	ReqDate     string                       `json:"req_date"`   // 请求日期
	ReqSeqId    string                       `json:"req_seq_id"` // 请求流水号
	TransAmt    string                       `json:"trans_amt"`
	GoodsDesc   string                       `json:"goods_desc"` // 商品描述
	AppData     common.StringObject[AppData] `json:"app_data"`
	ExtendInfos V2TradeHostingPaymentPreorderExtendInfo
}
type V2TradeHostingPaymentPreorderAliResponse struct {
	Data struct {
		RespCode     string           `json:"resp_code"`
		RespDesc     string           `json:"resp_desc"`
		ReqDate      string           `json:"req_date"`
		ReqSeqId     string           `json:"req_seq_id"`
		HuifuId      string           `json:"huifu_id"`
		TransAmt     string           `json:"trans_amt"`
		JumpUrl      string           `json:"jump_url"` //支付跳转链接	String
		PreOrderType string           `json:"pre_order_type"`
		PreOrderId   string           `json:"pre_order_id"`
		GoodsDesc    string           `json:"goods_desc"`
		UsageType    string           `json:"usage_type"`
		TransType    common.TradeType `json:"trans_type"`
		CurrentTime  string           `json:"current_time"`
		TimeExpire   string           `json:"time_expire"`
	} `json:"data"`
	Sign string `json:"sign"`
}

// V2TradeHostingPaymentPreorderWx 微信小程序预下单接口
// POST https://api.huifu.com/v2/trade/hosting/payment/preorder
//
// 应用场景
// 微信小程序预下单接口，可以满足商户或服务商在app拉起微信小程序支付的需求
func (t *Trade) V2TradeHostingPaymentPreorderWx(req V2TradeHostingPaymentPreorderWxRequest) (res *V2TradeHostingPaymentPreorderWxResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeHostingPaymentPreorderWxRequest(BsPaySdk.V2TradeHostingPaymentPreorderWxRequest{
		PreOrderType: "3",
		ReqDate:      req.ReqDate,
		ReqSeqId:     req.ReqSeqId,
		HuifuId:      req.HuifuId,
		TransAmt:     req.TransAmt,
		GoodsDesc:    req.GoodsDesc,
		MiniappData:  req.MiniappData.Encode(),
		ExtendInfos:  common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeHostingPaymentPreorderWxResponse](resp)
}

type V2TradeHostingPaymentPreorderWxRequest struct {
	HuifuId     string                           `json:"huifu_id"`   // 商户号
	ReqDate     string                           `json:"req_date"`   // 请求日期
	ReqSeqId    string                           `json:"req_seq_id"` // 请求流水号
	TransAmt    string                           `json:"trans_amt"`
	GoodsDesc   string                           `json:"goods_desc"` // 商品描述
	MiniappData common.StringObject[MiniappData] `json:"miniapp_data"`
	ExtendInfos V2TradeHostingPaymentPreorderExtendInfo
}
type V2TradeHostingPaymentPreorderWxResponse struct {
	Data struct {
		RespCode    string                           `json:"resp_code"`
		RespDesc    string                           `json:"resp_desc"`
		ReqDate     string                           `json:"req_date"`
		ReqSeqId    string                           `json:"req_seq_id"`
		HuifuId     string                           `json:"huifu_id"`
		TransAmt    string                           `json:"trans_amt"`
		PreOrderId  string                           `json:"pre_order_id"`
		MiniappData common.StringObject[MiniappData] `json:"miniapp_data"`
	} `json:"data"`
	Sign string `json:"sign"`
}

// V2TradeHostingPaymentPreorder 预下单统一接口 (上面3个接口的整合)
func (t *Trade) V2TradeHostingPaymentPreorder(req V2TradeHostingPaymentPreorderRequest) (res *V2TradeHostingPaymentPreorderResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeHostingPaymentPreorderRequest(BsPaySdk.V2TradeHostingPaymentPreorderRequest{
		HuifuId:      req.HuifuId,
		ReqDate:      req.ReqDate,
		ReqSeqId:     req.ReqSeqId,
		PreOrderType: req.PreOrderType,
		TransAmt:     req.TransAmt,
		GoodsDesc:    req.GoodsDesc,
		AppData:      req.AppData.Encode(),
		ExtendInfos:  common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeHostingPaymentPreorderResponse](resp)
}

type MiniappData struct {
	Appid       string `json:"appid"`
	GhId        string `json:"gh_id"`
	NeedScheme  string `json:"need_scheme"`
	Path        string `json:"path"`
	PrivateInfo string `json:"private_info"`
	SchemeCode  string `json:"scheme_code"`
	SeqId       string `json:"seq_id"`
}
type AppData struct {
	AppSchema   string //小程序返回码	String	100	Y	小程序完成支付后需要返回所填写的AppScheme，返回App必填信息。 示例值：https://callback.service.com/xx 注：参数过长容易导致浏览器截取跳转地址，无法唤起收银台，限100字符。 如果地址中有“+/?%#&=”字符需要进行编码操作
	PrivateInfo string //	私有信息	String	255	N	对应异步通知和主动查询接口中的remark字段；示例值：收取服务费
}
type HostingData struct {
	ProjectTitle string `json:"project_title"` //项目标题	String	64	Y	用于账单页面显示；示例值：汇付收银台
	ProjectId    string `json:"project_id"`    //半支付托管项目号	String	32	Y	商户创建的半支付托管项目号；示例值：PROJECTID2022032912492559
	PrivateInfo  string `json:"private_info"`  //商户私有信息	String	255	N	对应异步通知和主动查询接口中的remark字段；示例值：收取服务费
	CallbackUrl  string `json:"callback_url"`  //回调地址	String	512	N	若不填，支付成功后停留在当前页面，填写后跳转回指定地址；示例值：https://paas.huifu.com
	RequestType  string `json:"request_type"`  //请求类型	String	1	C	P:PC页面版，默认：P；M:H5页面版；指定交易类型时必填；示例值：M
}
type BizInfo struct {
	PayerCheckAli struct {
		NeedCheckInfo string //是否提供校验身份信息	String	1	N	T：强制校验，需要填写person_payer字段；F：不强制；示例值：T
		MinAge        string //允许的最小买家年龄	String	3	N	买家年龄必须大于等于所传数值。示例值：18 注：1. need_check_info=T 时该参数才有效， 2. min_age 为整数，必须大于等于 0；
		FixBuyer      string //是否强制校验付款人身份信息	String	8	N	用户进行实名认证校验；T：强制校验，F：不强制；示例值：T
	} //付款人验证（支付宝）	Object		N	支付宝特殊交易需验证买家信息；如彩票行业等； 当前只支持AT类交易有验证功能；
	PayerCheckWx struct {
		LimitPayer   string
		RealNameFlag string
	} //付款人验证（微信）	Object		N	微信实名支付需验证买家信息；如彩票行业等；当前只支持AT类交易有验证功能；
	PersonPayer struct {
		Name     string //姓名	String	16	N	注：支付宝交易need_check_info=T时，该参数才有效；示例值：张三
		CertType string //证件类型	String	32	N	身份证：IDENTITY_CARD，（微信只支持身份证）
		CertNo   string //证件号	String	64	N	注：支付宝交易need_check_info=T时，该参数才有效；需要密文传输，请参考加密解密说明使用汇付RSA公钥加密。 示例值：Mc5pjf+b/Keyi/t/wnH……MfYQnK7Lzw==
		Mobile   string //手机号	String	20	N	支付宝字段；注：该参数暂不校验；示例值：13911111111
	} //个人付款人信息	Object		N	付款人验证打开后需要填写付款人信息，但非必填；微信/支付宝共用字段；
}
type LargeamtData struct {
	CertificateName string
	BankCardNo      string
}
type V2TradeHostingPaymentPreorderExtendInfo struct {
	AcctId             string             `json:"acct_id"`
	StyleId            string             `json:"style_id"`             //收银台样式	String	30	N	指定收银台的风格样式；取值登录控台在【开发设置】->【品牌样式】页面查看，参见图片；示例值：128
	DelayAcctFlag      string             `json:"delay_acct_flag"`      //是否延迟交易	String	1	N	Y 为延迟，N为不延迟，不传默认N；示例值：N
	AcctSplitBunch     AcctSplitBunch     `json:"acct_split_bunch"`     //分账对象	String	2000	N	分账对象，jsonObject字符串
	HostingData        HostingData        `json:"hosting_data"`         //半支付托管扩展参数集合	String	2000	Y	jsonObject半支付托管扩展参数集合
	TimeExpire         string             `json:"time_expire"`          //交易失效时间	String	14	N	请求格式：yyyyMMddHHmmss；示例值：20220912111230 注意:为空默认失效时间为10分钟；用户在交易失效时间后完成交易有可能被关单。最终结果以异步为准； 建议商户在交易量大时，或在搞营销活动时将失效时间设置短一些。
	BizInfo            BizInfo            `json:"biz_info"`             //业务信息	String	2000	N	jsonObject格式；交易相关的信息
	NotifyUrl          string             `json:"notify_url"`           //交易异步通知地址	String	512	N	http或https开头，示例值：https://callback.service.com/xx
	UsageType          string             `json:"usage_type"`           //使用类型	String	1	N	P-支付（默认）； R-充值；示例值：P
	TransType          common.TradeType   `json:"trans_type"`           //交易类型	String	256	N	支持同时上送多个支付类型（多个时，使用英文逗号分割），上送多个或未传值时进入收银台，上送单个时直接到支付页；
	WxData             WxData             `json:"wx_data"`              //微信参数集合	String	2048	N	jsonObject字符串
	AlipayData         AlipayData         `json:"alipay_data"`          //支付宝参数集合	String	2048	N	jsonObject字符串
	UnionpayData       UnionpayData       `json:"unionpay_data"`        //银联参数集合	String	2048	N	jsonObject字符串
	TerminalDeviceData TerminalDeviceData `json:"terminal_device_data"` //设备信息	String	2048	N	jsonObject字符串
}
type V2TradeHostingPaymentPreorderRequest struct {
	HuifuId      string                       `json:"huifu_id"`       // 商户号
	ReqDate      string                       `json:"req_date"`       // 请求日期
	ReqSeqId     string                       `json:"req_seq_id"`     // 请求流水号
	PreOrderType string                       `json:"pre_order_type"` // 预下单类型
	TransAmt     string                       `json:"trans_amt"`      // 交易金额
	GoodsDesc    string                       `json:"goods_desc"`     // 商品描述
	AppData      common.StringObject[AppData] `json:"app_data"`       // app扩展参数集合
	ExtendInfos  V2TradeHostingPaymentPreorderExtendInfo
}
type V2TradeHostingPaymentPreorderResponse struct {
	Data struct {
		RespCode     string                           `json:"resp_code"`
		RespDesc     string                           `json:"resp_desc"`
		ReqDate      string                           `json:"req_date"`
		ReqSeqId     string                           `json:"req_seq_id"`
		HuifuId      string                           `json:"huifu_id"`
		PreOrderType string                           `json:"pre_order_type"`
		TransAmt     string                           `json:"trans_amt"`
		PreOrderId   string                           `json:"pre_order_id"`
		HostingData  common.StringObject[HostingData] `json:"hosting_data"`
		MiniappData  common.StringObject[MiniappData] `json:"miniapp_data"`
		JumpUrl      string                           `json:"jump_url"`
		GoodsDesc    string                           `json:"goods_desc"`
		UsageType    string                           `json:"usage_type"`
		TransType    common.TradeType                 `json:"trans_type"`
		CurrentTime  string                           `json:"current_time"`
		TimeExpire   string                           `json:"time_expire"`
	} `json:"data"`
	Sign string `json:"sign"`
}

// V2TradeHostingPaymentPreorderMessage 异步返回参数
// resp_data
// 注意：成功发起支付才会有异步返回；
type V2TradeHostingPaymentPreorderMessage struct {
	RespCode string                                                            `json:"resp_code"`
	RespData common.StringObject[V2TradeHostingPaymentPreorderMessageRespData] `json:"resp_data"`
	RespDesc string                                                            `json:"resp_desc"`
	Sign     string                                                            `json:"sign"`
}
type V2TradeHostingPaymentPreorderMessageRespData struct {
	V3TradePaymentJspayNotifyMessageRespData
	DebitFlag       string `json:"debit_flag"`
	UserHuifuId     string `json:"user_huifu_id"`
	BankId          string `json:"bank_id"`
	BankExtendParam string `json:"bank_extend_param"`
	OrderType       string `json:"order_type"` //订单类型	String	1	N	P-支付 R-充值 默认：P-支付；示例值：P
	RequestIp       string `json:"request_ip"`
}
