package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeHostingPaymentClose 托管交易关单
// POST https://api.huifu.com/v2/trade/hosting/payment/close
//
// 应用场景
// 托管支付关单接口，适用于微信、支付宝、云闪付交易，快捷网银待后续支持
func (t *Trade) V2TradeHostingPaymentClose(req V2TradeHostingPaymentCloseRequest) (res *V2TradeHostingPaymentCloseResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeHostingPaymentCloseRequest(BsPaySdk.V2TradeHostingPaymentCloseRequest{
		ReqSeqId:    req.ReqSeqId,
		ReqDate:     req.ReqDate,
		HuifuId:     req.HuifuId,
		OrgReqDate:  req.OrgReqDate,
		OrgReqSeqId: req.OrgReqSeqId,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeHostingPaymentCloseResponse](resp)
}

type V2TradeHostingPaymentCloseRequest struct {
	ReqSeqId    string `json:"req_seq_id"`     // 请求流水号
	ReqDate     string `json:"req_date"`       // 请求日期
	HuifuId     string `json:"huifu_id"`       // 商户号
	OrgReqDate  string `json:"org_req_date"`   // 原请求日期
	OrgReqSeqId string `json:"org_req_seq_id"` // 原请求流水号
}
type V2TradeHostingPaymentCloseResponse struct {
	Data struct {
		RespCode     string `json:"resp_code,omitempty"`
		RespDesc     string `json:"resp_desc,omitempty"`
		ReqSeqId     string `json:"req_seq_id,omitempty"`
		ReqDate      string `json:"req_date,omitempty"`
		HuifuId      string `json:"huifu_id,omitempty"`
		OrgReqDate   string `json:"org_req_date,omitempty"`
		OrgReqSeqId  string `json:"org_req_seq_id,omitempty"`
		OrgTransStat string `json:"org_trans_stat,omitempty"`
		CloseStat    string `json:"close_stat,omitempty"`
	} `json:"data"`
	Sign string `json:"sign"`
}

// V2TradeHostingPaymentCloseNotifyMessage 异步返回参数
type V2TradeHostingPaymentCloseNotifyMessage struct {
	RespCode string                                                               `json:"resp_code"`
	RespData common.StringObject[V2TradeHostingPaymentCloseNotifyMessageRespData] `json:"resp_data"`
	RespDesc string                                                               `json:"resp_desc"`
	Sign     string                                                               `json:"sign"`
}
type V2TradeHostingPaymentCloseNotifyMessageRespData struct {
	RespCode     string `json:"resp_code,omitempty"`      //业务响应码	String	8	Y	参见业务返回码；示例值：00000000
	RespDesc     string `json:"resp_desc,omitempty"`      //业务响应信息	String	512	Y	参见业务返回码；示例值：交易成功
	HuifuId      string `json:"huifu_id,omitempty"`       //商户号	String	32	Y	示例值：6666000123123123
	OrgReqDate   string `json:"org_req_date,omitempty"`   //原请求日期	String	8	Y	格式为yyyyMMdd，示例值：20221023
	OrgReqSeqId  string `json:"org_req_seq_id,omitempty"` //原请求流水号	String	64	Y	示例值：rQ20221010134649875651
	OrgTransStat string `json:"org_trans_stat,omitempty"` //原交易状态	String	1	N	P：处理中、S：成功、F：失败；示例值：S
	CloseStat    string `json:"close_stat,omitempty"`     //关单状态	String	1	Y	S：成功、F：失败；示例值：S
}
