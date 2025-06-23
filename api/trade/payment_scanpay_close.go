package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradePaymentScanpayClose 扫码交易关单
// POST https://api.huifu.com/v2/trade/payment/scanpay/close
//
// 应用场景
// 服务商/商户系统通过本接口发起订单关闭请求。
//
// 适用对象
// 开通微信/支付宝权限的商户。
// 注：银联、数字货币订单不支持关单；原交易已是终态（成功/失败）的，关单会失败。
func (t *Trade) V2TradePaymentScanpayClose(req V2TradePaymentScanpayCloseRequest) (res *V2TradePaymentScanpayCloseResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradePaymentScanpayCloseRequest(BsPaySdk.V2TradePaymentScanpayCloseRequest{
		ReqDate:     req.ReqDate,
		ReqSeqId:    req.ReqSeqId,
		HuifuId:     req.HuifuId,
		OrgReqDate:  req.OrgReqDate,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradePaymentScanpayCloseResponse](resp)
}

type V3TradePaymentScanpayCloseExtendInfos struct {
	OrgHfSeqId  string `json:"org_hf_seq_id"`
	OrgReqSeqId string `json:"org_req_seq_id"`
}
type V2TradePaymentScanpayCloseRequest struct {
	ReqDate     string `json:"req_date"`
	ReqSeqId    string `json:"req_seq_id"`
	HuifuId     string `json:"huifu_id"`
	OrgReqDate  string `json:"org_req_date"`
	ExtendInfos V3TradePaymentScanpayCloseExtendInfos
}
type V2TradePaymentScanpayCloseResponse struct {
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
