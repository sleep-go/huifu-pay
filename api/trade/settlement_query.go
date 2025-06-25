package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeSettlementQuery 出金交易查询
// POST https://api.huifu.com/v2/trade/settlement/query
//
// 最近更新时间：2024.12.24 作者: 张卿
//
// 应用场景
// 用于查询单笔出金交易状态及明细。
//
// 适用对象
// 商户需开通取现功能权限。
//
// 接口说明
// 支持格式：JSON
func (t *Trade) V2TradeSettlementQuery(req V2TradeSettlementQueryRequest) (res *V2TradeSettlementQueryResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeSettlementQueryRequest(BsPaySdk.V2TradeSettlementQueryRequest{
		HuifuId:     req.HuifuId,
		OrgReqDate:  req.OrgReqDate,
		OrgHfSeqId:  req.OrgHfSeqId,
		OrgReqSeqId: req.OrgReqSeqId,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeSettlementQueryResponse](resp)
}

type V2TradeSettlementQueryRequest struct {
	OrgReqDate  string
	OrgHfSeqId  string
	OrgReqSeqId string
	HuifuId     string `json:"huifu_id"`
}
type V2TradeSettlementQueryResponse struct {
	Data struct {
		RespCode    string `json:"resp_code"`
		RespDesc    string `json:"resp_desc"`
		OrgHfSeqId  string `json:"org_hf_seq_id"`
		OrgReqDate  string `json:"org_req_date"`
		TransStatus string `json:"trans_status"`
		OrgReqSeqId string `json:"org_req_seq_id"`
		TransDesc   string `json:"trans_desc"`
		CashAmt     string `json:"cash_amt"`
		FeeAmt      string `json:"fee_amt"`
	} `json:"data"`
	Sign string `json:"sign"`
}
