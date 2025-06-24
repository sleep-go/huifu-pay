package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeOnlinepaymentQuickpaySmscheck 快捷支付短信预校验
// POST https://api.huifu.com/v2/trade/onlinepayment/quickpay/smscheck
//
// 应用场景
// 使用快捷支付付款时，对短信验证码进行预校验，无需等待银行返回报错才能识别，提升用户使用体验。
//
// 适用对象
// 开通快捷支付商户；
func (t *Trade) V2TradeOnlinepaymentQuickpaySmscheck(req V2TradeOnlinepaymentQuickpaySmscheckRequest) (res *V2TradeOnlinepaymentQuickpaySmscheckResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeOnlinepaymentQuickpaySmscheckRequest(BsPaySdk.V2TradeOnlinepaymentQuickpaySmscheckRequest{
		ReqSeqId:    req.ReqSeqId,
		ReqDate:     req.ReqDate,
		HuifuId:     req.HuifuId,
		OrgReqDate:  req.OrgReqDate,
		OrgReqSeqId: req.OrgReqSeqId,
		SmsCode:     req.SmsCode,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeOnlinepaymentQuickpaySmscheckResponse](resp)
}

type V2TradeOnlinepaymentQuickpaySmscheckRequest struct {
	ReqSeqId    string `json:"req_seq_id"`
	ReqDate     string `json:"req_date"`
	HuifuId     string `json:"huifu_id"`
	OrgReqDate  string `json:"org_req_date"`
	OrgReqSeqId string `json:"org_req_seq_id"`
	SmsCode     string `json:"sms_code"`
}
type V2TradeOnlinepaymentQuickpaySmscheckResponse struct {
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
