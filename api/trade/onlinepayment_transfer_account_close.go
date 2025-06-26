package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeOnlinepaymentTransferAccountClose 银行大额支付关单
// POST https://api.huifu.com/v2/trade/onlinepayment/transfer/account/close
//
// 应用场景
// 1、调用银行大额支付接口后发现传参有误，可调用关单接口进行关单，关单后不可进行支付。若有支付需通过企业微信、邮件等方式联系售后经理、汇付运营，进行人工线下流程处理。
// 2、如2个工作日未完成打款，系统将自动关闭订单。
// 注：针对订单标识或金额错误，且支付卡号与金额正确，可调用差错申请接口做差错处理。
//
// 适用对象
// 已调银行大额支付接口的SAAS服务商、商户。
func (t *Trade) V2TradeOnlinepaymentTransferAccountClose(req V2TradeOnlinepaymentTransferAccountCloseRequest) (res *V2TradeOnlinepaymentTransferAccountCloseResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeOnlinepaymentTransferAccountCloseRequest(BsPaySdk.V2TradeOnlinepaymentTransferAccountCloseRequest{
		ReqSeqId:    req.ReqSeqId,
		ReqDate:     req.ReqDate,
		HuifuId:     req.HuifuId,
		OrgReqSeqId: req.ReqSeqId,
		OrgReqDate:  req.OrgReqDate,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeOnlinepaymentTransferAccountCloseResponse](resp)
}

type V2TradeOnlinepaymentTransferAccountCloseExtendInfo struct {
	BankCardNo string `json:"bank_card_no"`
	InAcctFlag string `json:"in_acct_flag"`
}
type V2TradeOnlinepaymentTransferAccountCloseRequest struct {
	ReqSeqId    string `json:"req_seq_id"` // 请求流水号
	ReqDate     string `json:"req_date"`   // 请求日期
	HuifuId     string `json:"huifu_id"`   // 收款方商户号
	OrgReqSeqId string `json:"org_req_seq_id"`
	OrgReqDate  string `json:"org_req_date"`
	ExtendInfos V2TradeOnlinepaymentTransferAccountCloseExtendInfo
}
type V2TradeOnlinepaymentTransferAccountCloseResponse struct {
	Data struct {
		RespCode       string `json:"resp_code"`
		RespDesc       string `json:"resp_desc"`
		HuifuId        string `json:"huifu_id"`
		ReqSeqId       string `json:"req_seq_id"`
		ReqDate        string `json:"req_date"`
		OrgReqSeqId    string `json:"org_req_seq_id"`
		OrgReqDate     string `json:"org_req_date"`
		InAcctFlag     string `json:"in_acct_flag"`
		OrgHfSeqId     string `json:"org_hf_seq_id"`
		OrgTransStat   string `json:"org_trans_stat"`
		TransCloseStat string `json:"trans_close_stat"`
	} `json:"data"`
	Sign string `json:"sign"`
}
