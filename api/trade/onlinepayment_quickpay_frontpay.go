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
}
