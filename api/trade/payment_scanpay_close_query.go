package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

func (t *Trade) V2TradePaymentScanpayCloseQuery(req V2TradePaymentScanpayCloseQueryRequest) (res *V2TradePaymentScanpayCloseQueryResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradePaymentScanpayClosequeryRequest(BsPaySdk.V2TradePaymentScanpayClosequeryRequest{
		ReqDate:     req.ReqDate,
		ReqSeqId:    req.ReqSeqId,
		HuifuId:     req.HuifuId,
		OrgReqDate:  req.OrgReqDate,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradePaymentScanpayCloseQueryResponse](resp)
}

type V2TradePaymentScanpayClosequeryExtendInfos struct {
	OrgHfSeqId  string `json:"org_hf_seq_id"`
	OrgReqSeqId string `json:"org_req_seq_id"`
}
type V2TradePaymentScanpayCloseQueryRequest struct {
	ReqDate     string `json:"req_date"`
	ReqSeqId    string `json:"req_seq_id"`
	HuifuId     string `json:"huifu_id"`
	OrgReqDate  string `json:"org_req_date"`
	ExtendInfos V2TradePaymentScanpayClosequeryExtendInfos
}
type V2TradePaymentScanpayCloseQueryResponse struct {
	Data struct {
		RespCode     string `json:"resp_code"`
		RespDesc     string `json:"resp_desc"`
		HuifuId      string `json:"huifu_id"`
		ReqDate      string `json:"req_date"`
		ReqSeqId     string `json:"req_seq_id"`
		OrgReqDate   string `json:"org_req_date"`
		OrgReqSeqId  string `json:"org_req_seq_id"`
		OrgHfSeqId   string `json:"org_hf_seq_id"`
		OrgTransStat string `json:"org_trans_stat"`
		TransStat    string `json:"trans_stat"`
	} `json:"data"`
	Sign string `json:"sign"`
}
