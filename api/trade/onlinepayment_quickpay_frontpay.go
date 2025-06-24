package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeOnlinepaymentQuickpayFrontpay 快捷支付页面版
// POST https://api.huifu.com/v2/trade/onlinepayment/quickpay/frontpay
//
// 最近更新时间：2025.03.11 作者: 魏溪
//
// 应用场景
// 本接口直接返回快捷支付页面，封装了快捷支付流程。用户在页面上输入银行卡等信息，斗拱平台会向持卡人在银行预留的手机号发送短信验证码，系统自动完成绑卡支付。
//
// 适用对象
// 斗拱平台提供快捷支付页面，用户在快捷支付页面上可以完成绑卡、支付等操作。适合想快速接入快捷支付方式的服务商或渠道商；
//
// 接口说明
// 支持格式：JSON
func (t *Trade) V2TradeOnlinepaymentQuickpayFrontpay(req V2TradeOnlinepaymentQuickpayFrontpayRequest) (res *V2TradeOnlinepaymentQuickpayFrontpayResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeOnlinepaymentQuickpayFrontpayRequest(BsPaySdk.V2TradeOnlinepaymentQuickpayFrontpayRequest{
		ReqSeqId:    req.ReqSeqId,
		ReqDate:     req.ReqDate,
		HuifuId:     req.HuifuId,
		TransAmt:    req.TransAmt,
		NotifyUrl:   req.NotifyUrl,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeOnlinepaymentQuickpayFrontpayResponse](resp)
}

type ExtendPayData struct {
	GoodsShortName string `json:"goods_short_name,omitempty"` //商品简称	String	40	Y	示例值：个人电脑，不能包含特殊字符<、>、&、'、"、
	BizTp          string `json:"biz_tp,omitempty"`           //业务种类	String	6	Y	说明详见附录《业务种类字典》；示例值：100099
	GwChnnlTp      string `json:"gw_chnnl_tp,omitempty"`      //网关支付受理渠道	String	2	Y	01：电脑浏览器，02：手机浏览器，03：手机应用程序，99：其他，(默认)；示例值：01
}
type V2TradeOnlinepaymentQuickpayFrontpayExtendInfo struct {
	TimeExpire         string             `json:"time_expire"`
	RequestType        string             `json:"request_type"`
	GoodsDesc          string             `json:"goods_desc"`
	FeeFlag            string             `json:"fee_flag"`
	OrderType          string             `json:"order_type"`
	DelayAcctFlag      string             `json:"delay_acct_flag"`
	AcctSplitBunch     AcctSplitBunch     `json:"acct_split_bunch"`
	ExtendPayData      ExtendPayData      `json:"extend_pay_data"`
	TerminalDeviceData TerminalDeviceData `json:"terminal_device_data"`
	RiskCheckData      RiskCheckData      `json:"risk_check_data"`
	Remark             string             `json:"remark"`
	FrontUrl           string             `json:"front_url"`
}
type V2TradeOnlinepaymentQuickpayFrontpayRequest struct {
	ReqSeqId    string `json:"req_seq_id"`
	ReqDate     string `json:"req_date"`
	HuifuId     string `json:"huifu_id"`
	TransAmt    string `json:"trans_amt"`
	NotifyUrl   string `json:"notify_url"`
	ExtendInfos V2TradeOnlinepaymentQuickpayFrontpayExtendInfo
}
type V2TradeOnlinepaymentQuickpayFrontpayResponse struct {
	Data struct {
		RespCode       string                              `json:"resp_code"`
		RespDesc       string                              `json:"resp_desc"`
		ReqSeqId       string                              `json:"req_seq_id"`
		ReqDate        string                              `json:"req_date"`
		MerName        string                              `json:"mer_name"`
		UserHuifuId    string                              `json:"user_huifu_id"`
		GoodsDesc      string                              `json:"goods_desc"`
		TransAmt       string                              `json:"trans_amt"`
		TransStat      string                              `json:"trans_stat"`
		HfSeqId        string                              `json:"hf_seq_id"`
		RequestType    string                              `json:"request_type"`
		FormUrl        string                              `json:"form_url"`
		Remark         string                              `json:"remark"`
		AcctSplitBunch common.StringObject[AcctSplitBunch] `json:"acct_split_bunch"`
	} `json:"data"`
	Sign string `json:"sign"`
}
type V2TradeOnlinepaymentQuickpayFrontpayNotifyMessage struct {
	RespCode string                                                                     `json:"resp_code"`
	RespDesc string                                                                     `json:"resp_desc"`
	Data     common.StringObject[V2TradeOnlinepaymentQuickpayFrontpayNotifyMessageData] `json:"data"`
	Sign     string
}
type V2TradeOnlinepaymentQuickpayFrontpayNotifyMessageData struct {
	RespCode       string         `json:"resp_code"`
	RespDesc       string         `json:"resp_desc"`
	ReqSeqId       string         `json:"req_seq_id"`
	ReqDate        string         `json:"req_date"`
	UserHuifuId    string         `json:"user_huifu_id"`
	GoodsDesc      string         `json:"goods_desc"`
	TransAmt       string         `json:"trans_amt"`
	TransStat      string         `json:"trans_stat"`
	HfSeqId        string         `json:"hf_seq_id"`
	RequestType    string         `json:"request_type"`
	FormUrl        string         `json:"form_url"`
	Remark         string         `json:"remark"`
	ProductId      string         //产品号	String	32	Y	汇付分配的产品号；示例值：YYZY
	HuifuId        string         //商户号	String	32	Y	汇付分配的商户号；示例值：6666000108854952
	MerName        string         //商户名称	String	256	N	示例值：上海汇付支付有限公司
	MerPriv        string         //商户私有域	String	1024	N	原样返回请求参数中的“remark备注”内容；示例值：商户私有域
	TransDate      string         //交易日期	String	8	N	格式：YYYYMMDD；示例值：20230413
	TransTime      string         //交易时间	String	6	N	格式：HHMMSS；示例值：121009
	BankId         string         //银行代号	String	8	N	示例值：30200000
	DebitFlag      string         //借贷标识	String	1	N	D：借记；C：贷记；示例值：D
	FeeFlag        string         //手续费扣款标志	String	1	N	1：外扣；2：内扣；示例值：2
	FeeAmt         string         //手续费	String	12	N	单位元，需保留小数点后两位，示例值：1.00
	IsDiv          string         //是否分账交易	String	1	N	0：非分账交易；1：是分账交易示例值：0
	IsDelayAcct    string         //是否延时分账	String	1	N	0：实时；1：延时；示例值：1 注意延时交易要调交易确认接口资金才能进入收款方账户，否则会停留在延时账户中。
	AcctSplitBunch AcctSplitBunch //分账串	Object		N	格式：jsonObject；参见分账串
	SplitFeeInfo   SplitFeeInfo   //分账手续费信息	Object	2048	N	jsonObject格式
	TransType      string         //交易类型	String	30	Y	QUICK_PAY：快捷支付，QUICK_RECHARGE：快捷充值；示例值：QUICK_PAY
	OrderType      string         //订单类型	String	1	N	P-支付 R-充值 默认：P-支付；示例值：P
	BankCode       string         //通道返回码	String	8	N	示例值：00
	BankMessage    string         //通道返回描述	String	128	N	示例值：成功[0000000]
}
