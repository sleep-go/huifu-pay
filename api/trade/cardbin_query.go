package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeCardbinQuery 卡bin信息查询
// POST https://api.huifu.com/v2/trade/cardbin/query
//
// 应用场景
// 用于通过银行卡号查询银行卡卡BIN信息，可以返回卡片类型等信息
//
// 适用对象
// 有限制交易卡片类型的交易流程，或单纯用户查询
func (t *Trade) V2TradeCardbinQuery(req V2TradeCardbinQueryRequest) (res *V2TradeCardbinQueryResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeCardbinQueryRequest(BsPaySdk.V2TradeCardbinQueryRequest{
		ReqDate:         req.ReqDate,
		ReqSeqId:        req.ReqSeqId,
		BankCardNoCrypt: req.BankCardNoCrypt,
		ExtendInfos:     nil,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeCardbinQueryResponse](resp)
}

type V2TradeCardbinQueryRequest struct {
	ReqDate         string `json:"req_date,omitempty"`           // 请求日期
	ReqSeqId        string `json:"req_seq_id,omitempty"`         // 请求流水号
	BankCardNoCrypt string `json:"bank_card_no_crypt,omitempty"` // 银行卡号密文
}
type V2TradeCardbinQueryResponse struct {
	Data struct {
		RespCode string `json:"resp_code"`
		RespDesc string `json:"resp_desc"`
		ReqDate  string `json:"req_date"`
		ReqSeqId string `json:"req_seq_id"`
		CardType string `json:"card_type"` //C-贷记卡;D-借记卡;S-准贷记卡;P-预付通卡；
		CardBin  string `json:"card_bin"`
		BankName string `json:"bank_name"`
		IssuerId string `json:"issuer_id"`
		BankCode string `json:"bank_code"`
	} `json:"data"`
	Sign string `json:"sign"`
}
