package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeAcctpaymentRefund 余额支付退款
// POST https://api.huifu.com/v2/trade/acctpayment/refund
//
// 应用场景
// 本API为后台版接口，用于支持余额支付成功交易的全额或部分退款。
//
// 适用对象
// 已开通余额支付相关功能的商户。
func (t *Trade) V2TradeAcctpaymentRefund(req V2TradeAcctpaymentRefundRequest) (res *V2TradeAcctpaymentRefundResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeAcctpaymentRefundRequest(BsPaySdk.V2TradeAcctpaymentRefundRequest{
		ReqSeqId:    req.ReqSeqId,
		ReqDate:     req.ReqDate,
		HuifuId:     req.HuifuId,
		OrgReqDate:  req.OrgReqDate,
		OrgReqSeqId: req.OrgReqSeqId,
		OrgHfSeqId:  req.OrgHfSeqId,
		OrdAmt:      req.OrdAmt,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeAcctpaymentRefundResponse](resp)
}

type V2TradeAcctpaymentRefundRequest struct {
	ReqSeqId    string `json:"req_seq_id"`
	ReqDate     string `json:"req_date"`
	HuifuId     string `json:"huifu_id"`
	OrgReqDate  string `json:"org_req_date"`
	OrgReqSeqId string `json:"org_req_seq_id"`
	OrgHfSeqId  string `json:"org_hf_seq_id"`
	OrdAmt      string `json:"ord_amt"`
	ExtendInfos struct {
		AcctSplitBunch AcctSplitBunch `json:"acct_split_bunch"`
		RiskCheckData  RiskCheckData  `json:"risk_check_data"`
		Remark         string         `json:"remark"`
	}
}
type V2TradeAcctpaymentRefundResponse struct {
	Data struct {
		RespCode        string                              `json:"resp_code"`
		RespDesc        string                              `json:"resp_desc"`
		ReqDate         string                              `json:"req_date"`
		ReqSeqId        string                              `json:"req_seq_id"`
		HuifuId         string                              `json:"huifu_id"`
		ProductId       string                              `json:"product_id"`
		OrdAmt          string                              `json:"ord_amt"`
		OrgReqDate      string                              `json:"org_req_date"`
		OrgReqSeqId     string                              `json:"org_req_seq_id"`
		OrgHfSeqId      string                              `json:"org_hf_seq_id"`
		TransStat       string                              `json:"trans_stat"`
		TransFinishTime string                              `json:"trans_finish_time"`
		AcctSplitBunch  common.StringObject[AcctSplitBunch] `json:"acct_split_bunch"`
		Remark          string                              `json:"remark"`
	} `json:"data"`
	Sign string `json:"sign"`
}
