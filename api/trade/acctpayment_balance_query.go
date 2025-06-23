package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeAcctpaymentBalanceQuery 账户余额信息查询
// POST https://api.huifu.com/v2/trade/acctpayment/balance/query
//
// 应用场景
// 本接口用于查询商户各类账户的余额信息。
//
// 适用对象
// 商户完成基本信息进件或业务入驻流程创建账户后才能查询账户余额信息。
func (t *Trade) V2TradeAcctpaymentBalanceQuery(req V2TradeAcctpaymentBalanceQueryRequest) (res *V2TradeAcctpaymentBalanceQueryResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeAcctpaymentBalanceQueryRequest(BsPaySdk.V2TradeAcctpaymentBalanceQueryRequest{
		ReqDate:     req.ReqDate,
		HuifuId:     req.HuifuId,
		ExtendInfos: nil,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeAcctpaymentBalanceQueryResponse](resp)
}

type V2TradeAcctpaymentBalanceQueryRequest struct {
	ReqDate string `json:"req_date"`
	HuifuId string `json:"huifu_id"`
}
type V2TradeAcctpaymentBalanceQueryResponse struct {
	Data struct {
		AcctInfoList common.StringObject[AcctInfoList] `json:"acctInfo_list"`
		ReqDate      string                            `json:"req_date"`
		ReqSeqId     string                            `json:"req_seq_id"`
		RespCode     string                            `json:"resp_code"`
		RespDesc     string                            `json:"resp_desc"`
	} `json:"data"`
	Sign string `json:"sign"`
}
type AcctInfoList struct {
	HuifuId    string `json:"huifu_id"`
	AcctId     string `json:"acct_id"`
	AcctType   string `json:"acct_type"`
	BalanceAmt string `json:"balance_amt"`
	AvlBal     string `json:"avl_bal"`
	FrzBal     string `json:"frz_bal"`
	LastAvlBal string `json:"last_avl_bal"`
	AcctStat   string `json:"acct_stat"`
	ControlBal string `json:"control_bal"`
}
