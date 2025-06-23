package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeAcctpaymentRefundQuery 余额支付退款查询
// POST https://api.huifu.com/v2/trade/acctpayment/refund/query
//
// 应用场景
// 本API为后台版接口，用于查询商户单笔余额支付退款状态及明细。
//
// 适用对象
// 已开通余额支付相关功能的商户。
func (t *Trade) V2TradeAcctpaymentRefundQuery(req V2TradeAcctpaymentRefundQueryRequest) (res *V2TradeAcctpaymentRefundQueryResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeAcctpaymentRefundQueryRequest(BsPaySdk.V2TradeAcctpaymentRefundQueryRequest{
		OrgReqSeqId: req.OrgReqSeqId,
		OrgReqDate:  req.OrgReqDate,
		HuifuId:     req.HuifuId,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeAcctpaymentRefundQueryResponse](resp)
}

type V2TradeAcctpaymentRefundQueryRequest struct {
	OrgReqSeqId string `json:"org_req_seq_id"`
	OrgReqDate  string `json:"org_req_date"`
	HuifuId     string `json:"huifu_id"`
}
type V2TradeAcctpaymentRefundQueryResponse struct {
	V2TradeAcctpaymentPayQueryResponse
}
