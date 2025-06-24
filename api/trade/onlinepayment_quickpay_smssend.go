package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V3TradeOnlinepaymentQuickpaySmssend 快捷短信发送
// POST https://api.huifu.com/v3/trade/onlinepayment/quickpay/smssend
//
// 应用场景
// 本接口用于快捷支付流程中，商户系统调用该接口发起快捷支付申请，斗拱平台会向持卡人在银行预留的手机号发送短信验证码。短验码的有效期10分钟，建议开发者前端交互设计压缩一下订单有效期，时间过长会降低付款的成功率。
//
// 适用对象
// 商户需要开通快捷支付业务，并且持卡人已在斗拱平台完成绑卡签约操作，已获取绑卡id。
//
// 接口说明
// 支持格式：JSON
func (t *Trade) V3TradeOnlinepaymentQuickpaySmssend(req V3TradeOnlinepaymentQuickpaySmssendRequest) (res *V3TradeOnlinepaymentQuickpaySmssendResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V3TradeOnlinepaymentQuickpaySmssendRequest(BsPaySdk.V3TradeOnlinepaymentQuickpaySmssendRequest{
		ReqDate:     req.ReqDate,
		ReqSeqId:    req.ReqSeqId,
		HuifuId:     req.HuifuId,
		UserHuifuId: req.UserHuifuId,
		CardBindId:  req.CardBindId,
		TransAmt:    req.TransAmt,
		NotifyUrl:   req.NotifyUrl,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V3TradeOnlinepaymentQuickpaySmssendResponse](resp)
}

type NuccData struct {
	GoodsShortName string //商品简称	String	40	Y	(网联银联快捷支付使用)示例值：个人电脑
	GwChnnlTp      string //网关支付受理渠道	String	2	Y	网关支付受理渠道：01：电脑浏览器，02：手机浏览器，03：手机应用程序，99：其他
	BizTp          string //业务种类	String	6	Y	说明详见附录《业务种类字典》，示例值：100099
}
type V3TradeOnlinepaymentQuickpaySmssendExtendInfo struct {
	OrderType          string             `json:"order_type"`
	Remark             string             `json:"remark"`
	TimeExpire         string             `json:"time_expire"`
	AcctSplitBunch     AcctSplitBunch     `json:"acct_split_bunch"`
	DelayAcctFlag      string             `json:"delay_acct_flag"`
	AcctId             string             `json:"acct_id"`
	FeeFlag            string             `json:"fee_flag"`
	NuccData           NuccData           `json:"nucc_data"`
	TerminalDeviceData TerminalDeviceData `json:"terminal_device_data"`
	RiskCheckData      RiskCheckData      `json:"risk_check_data"`
	CombinedpayData    []CombinedpayData  `json:"combinedpay_data"`
}
type V3TradeOnlinepaymentQuickpaySmssendRequest struct {
	ReqSeqId    string `json:"req_seq_id"`
	ReqDate     string `json:"req_date"`
	UserHuifuId string `json:"user_huifu_id"`
	CardBindId  string `json:"card_bind_id"`
	TransAmt    string `json:"trans_amt"`
	HuifuId     string `json:"huifu_id"`
	NotifyUrl   string `json:"notify_url"`
	ExtendInfos V3TradeOnlinepaymentQuickpaySmssendExtendInfo
}
type V3TradeOnlinepaymentQuickpaySmssendResponse struct {
	Data struct {
		RespCode    string `json:"resp_code"`
		RespDesc    string `json:"resp_desc"`
		ReqDate     string `json:"req_date"`
		ReqSeqId    string `json:"req_seq_id"`
		HuifuId     string `json:"huifu_id"`
		UserHuifuId string `json:"user_huifu_id"`
		HfSeqId     string `json:"hf_seq_id"`
		CardBindId  string `json:"card_bind_id"`
		TransAmt    string `json:"trans_amt"`
		TransType   string `json:"trans_type"`
		SmsSendStat string `json:"sms_send_stat"`
		Remark      string `json:"remark"`
	} `json:"data"`
	Sign string `json:"sign"`
}
type V3TradeOnlinepaymentQuickpaySmssendNotifyMessage struct {
	RespDesc string                                                                    `json:"resp_desc"`
	RespCode string                                                                    `json:"resp_code"`
	Data     common.StringObject[V3TradeOnlinepaymentQuickpaySmssendNotifyMessageData] `json:"data"`
	Sign     string                                                                    `json:"sign"`
}
type V3TradeOnlinepaymentQuickpaySmssendNotifyMessageData struct {
	RespCode    string `json:"resp_code"`
	RespDesc    string `json:"resp_desc"`
	ReqDate     string `json:"req_date"`
	ReqSeqId    string `json:"req_seq_id"`
	ProductId   string `json:"product_id"`
	HuifuId     string `json:"huifu_id"`
	UserHuifuId string `json:"user_huifu_id"`
	HfSeqId     string `json:"hf_seq_id"`
	TransAmt    string `json:"trans_amt"`
	TransType   string `json:"trans_type"`
	SmsSendStat string `json:"sms_send_stat"`
	Remark      string `json:"remark"`
	OrderType   string `json:"order_type"`   //订单类型	String	1	N	P：支付，R：充值，默认：P（支付）；示例值：P
	BankCode    string `json:"bank_code"`    //通道返回码	String	32	N	请勿根据此字段判断交易状态，此字段建议在交易失败时配合bank_message使用。示例值：00
	BankMessage string `json:"bank_message"` //通道返回描述	String	200	N	示例值：成功[0000000]
}
