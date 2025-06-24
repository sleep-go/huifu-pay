package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeHostingPaymentQueryrefundinfo 托管交易关单
// POST https://api.huifu.com/v2/trade/hosting/payment/close
//
// 应用场景
// 托管支付关单接口，适用于微信、支付宝、云闪付交易，快捷网银待后续支持
func (t *Trade) V2TradeHostingPaymentQueryrefundinfo(req V2TradeHostingPaymentQueryrefundinfoRequest) (res *V2TradeHostingPaymentQueryrefundinfoResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeHostingPaymentQueryrefundinfoRequest(BsPaySdk.V2TradeHostingPaymentQueryrefundinfoRequest{
		ReqDate:     req.ReqDate,
		ReqSeqId:    req.ReqSeqId,
		HuifuId:     req.HuifuId,
		OrgReqDate:  req.OrgReqDate,
		OrgHfSeqId:  req.OrgHfSeqId,
		OrgReqSeqId: req.OrgReqSeqId,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeHostingPaymentQueryrefundinfoResponse](resp)
}

type V2TradeHostingPaymentQueryrefundinfoRequest struct {
	ReqDate     string `json:"req_date"`
	ReqSeqId    string `json:"req_seq_id"`
	HuifuId     string `json:"huifu_id"`
	OrgHfSeqId  string `json:"org_hf_seq_id"`
	OrgReqSeqId string `json:"org_req_seq_id"`
	OrgReqDate  string `json:"org_req_date"`
}
type V2TradeHostingPaymentQueryrefundinfoResponse struct {
	Data struct {
		RespCode         string                                       `json:"resp_code"`
		RespDesc         string                                       `json:"resp_desc"`
		ReqDate          string                                       `json:"req_date"`
		ReqSeqId         string                                       `json:"req_seq_id"`
		HuifuId          string                                       `json:"huifu_id"`
		OrgHfSeqId       string                                       `json:"org_hf_seq_id"`
		OrgReqDate       string                                       `json:"org_req_date"`
		OrgReqSeqId      string                                       `json:"org_req_seq_id"`
		OrdAmt           string                                       `json:"ord_amt"`
		ActualRefAmt     string                                       `json:"actual_ref_amt"`
		TransStat        string                                       `json:"trans_stat"`
		BankCode         string                                       `json:"bank_code"`
		BankMessage      string                                       `json:"bank_message"`
		FeeAmt           string                                       `json:"fee_amt"`
		AcctSplitBunch   common.StringObject[AcctSplitBunch]          `json:"acct_split_bunch"`
		SplitFeeInfo     common.StringObject[SplitFeeInfo]            `json:"split_fee_info"`
		OrgPartyOrderId  string                                       `json:"org_party_order_id"`
		OrgOutOrderId    string                                       `json:"org_out_order_id"`
		UnionpayResponse common.StringObject[common.UnionpayResponse] `json:"unionpay_response"`
		TransFinishTime  string                                       `json:"trans_finish_time"`
	} `json:"data"`
	Sign string `json:"sign"`
}
