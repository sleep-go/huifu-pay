package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V3TradeOnlinepaymentQuickpayPay 快捷支付
// POST https://api.huifu.com/v3/trade/onlinepayment/quickpay/pay
//
// 应用场景
// 本接口用于快捷支付流程中，商户系统调用快捷支付申请接口后，斗拱平台会向持卡人在银行预留的手机号发送短信验证码，商户系统调用快捷支付确认接口，上送短信验证码，完成支付。
//
// 适用对象
// 商户需要开通快捷支付业务，并且持卡人已在斗拱平台完成绑卡签约操作，已获取绑卡id。
//
// 接口说明
// 支持格式：JSON
func (t *Trade) V3TradeOnlinepaymentQuickpayPay(req V3TradeOnlinepaymentQuickpayPayRequest) (res *V3TradeOnlinepaymentQuickpayPayResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V3TradeOnlinepaymentQuickpayPayRequest(BsPaySdk.V3TradeOnlinepaymentQuickpayPayRequest{
		ReqSeqId:    req.ReqSeqId,
		ReqDate:     req.ReqDate,
		HuifuId:     req.HuifuId,
		SmsCode:     req.SmsCode,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V3TradeOnlinepaymentQuickpayPayResponse](resp)
}

type V3TradeOnlinepaymentQuickpayPayExtendInfo struct {
	NotifyUrl string `json:"notify_url"`
	Remark    string `json:"remark"`
}
type V3TradeOnlinepaymentQuickpayPayRequest struct {
	ReqSeqId    string `json:"req_seq_id"`
	ReqDate     string `json:"req_date"`
	HuifuId     string `json:"huifu_id"`
	SmsCode     string `json:"sms_code"`
	ExtendInfos V3TradeOnlinepaymentQuickpayPayExtendInfo
}
type V3TradeOnlinepaymentQuickpayPayResponse struct {
	Data struct {
		RespCode    string `json:"resp_code"`
		RespDesc    string `json:"resp_desc"`
		ReqDate     string `json:"req_date"`
		ReqSeqId    string `json:"req_seq_id"`
		HuifuId     string `json:"huifu_id"`
		HfSeqId     string `json:"hf_seq_id"`
		UserHuifuId string `json:"user_huifu_id"`
		TransAmt    string `json:"trans_amt"`
		TransType   string `json:"trans_type"`
		TransStat   string `json:"trans_stat"`
		Remark      string `json:"remark"`
	} `json:"data"`
	Sign string `json:"sign"`
}
type PyerResponse struct {
	GateType         string `json:"gate_type"`           //网关支付类型	String	2	N	01: 个人网关02:企业网关
	PyerBankId       string `json:"pyer_bank_id"`        //付款方银行号	String	32	N
	PyerBankCardNo   string `json:"pyer_bank_card_no"`   //付款方银行卡号	String	1024	N	返回密文
	PyerBankAcctName string `json:"pyer_bank_acct_name"` //付款方银行账户名	String	128	N
}
type V3TradeOnlinepaymentQuickpayPayNotifyMessageData struct {
	RespCode          string            `json:"resp_code"`
	RespDesc          string            `json:"resp_desc"`
	ReqDate           string            `json:"req_date"`
	ReqSeqId          string            `json:"req_seq_id"`
	HuifuId           string            `json:"huifu_id"`
	UserHuifuId       string            `json:"user_huifu_id"`
	HfSeqId           string            `json:"hf_seq_id"`
	TransAmt          string            `json:"trans_amt"`
	TransType         string            `json:"trans_type"`
	TransStat         string            `json:"trans_stat"`
	OrderType         string            `json:"order_type"`
	AcctStat          string            `json:"acct_stat"`
	BankCode          string            `json:"bank_code"`
	BankMessage       string            `json:"bank_message"`
	DelayAcctFlag     string            `json:"delay_acct_flag"`
	Remark            string            `json:"remark"`
	AcctFinishTime    string            `json:"acct_finish_time"`
	TransFinishTime   string            `json:"trans_finish_time"`
	EndTime           string            `json:"end_time"`
	FeeAmt            string            `json:"fee_amt"`
	FeeFlag           string            `json:"fee_flag"`
	NotifyType        string            `json:"notify_type"`
	CombinedpayFeeAmt string            `json:"combinedpay_fee_amt"`
	PyerResponse      PyerResponse      `json:"pyer_response"`
	SplitFeeInfo      SplitFeeInfo      `json:"split_fee_info"`
	AcctSplitBunch    AcctSplitBunch    `json:"acct_split_bunch"`
	CombinedpayData   []CombinedpayData `json:"combinedpay_data"`
}
type V3TradeOnlinepaymentQuickpayPayNotifyMessage struct {
	RespDesc string                                                                `json:"resp_desc"`
	RespCode string                                                                `json:"resp_code"`
	Data     common.StringObject[V3TradeOnlinepaymentQuickpayPayNotifyMessageData] `json:"data"`
	Sign     string                                                                `json:"sign"`
}
