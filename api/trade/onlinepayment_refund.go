package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeOnlinepaymentRefund 线上交易退款
// POST https://api.huifu.com/v2/trade/onlinepayment/refund
//
// 应用场景
// 支持线上交易全额或部分退款。支持快捷、网银、手机wap、银联统一在线收银台、分期支付、银行大额支付、代扣交易退款。
// 补贴支付退款：支持全额退款、银行卡退款，不支持单独退补贴金额；
// 支持最大退款期限：
//
// 快捷：180天
// 网银：180天
// 手机WAP：180天
// 银联统一在线收银台：180天
// 分期支付：180天
// 银行大额支付：180天
// 代扣：180天
// 适用对象
// 支持单笔且状态为成功的线上交易的全额或部分退款。
func (t *Trade) V2TradeOnlinepaymentRefund(req V2TradeOnlinepaymentRefundRequest) (res *V2TradeOnlinepaymentRefundResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeOnlinepaymentRefundRequest(BsPaySdk.V2TradeOnlinepaymentRefundRequest{
		ReqDate:     req.ReqDate,
		ReqSeqId:    req.ReqSeqId,
		HuifuId:     req.HuifuId,
		OrdAmt:      req.OrdAmt,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeOnlinepaymentRefundResponse](resp)
}

type V2TradeOnlinepaymentRefundExtendInfo struct {
	OrgReqSeqId        string             `json:"org_req_seq_id"`
	OrgReqDate         string             `json:"org_req_date"`
	OrgHfSeqId         string             `json:"org_hf_seq_id"`
	AcctSplitBunch     AcctSplitBunch     `json:"acct_split_bunch"`
	TerminalDeviceData TerminalDeviceData `json:"terminal_device_data"`
	RiskCheckData      RiskCheckData      `json:"risk_check_data"`
	Remark             string             `json:"remark"`
	NotifyUrl          string             `json:"notify_url"`
	CombinedpayData    []CombinedpayData  `json:"combinedpay_data"`
	BankInfoData       BankInfoData       `json:"bank_info_data"`
}
type V2TradeOnlinepaymentRefundRequest struct {
	HuifuId     string `json:"huifu_id"`
	OrdAmt      string `json:"ord_amt"`
	ReqSeqId    string `json:"req_seq_id"`
	ReqDate     string `json:"req_date"`
	ExtendInfos V2TradeOnlinepaymentRefundExtendInfo
}
type V2TradeOnlinepaymentRefundResponse struct {
	Data struct {
		RespCode          string                                 `json:"resp_code"`
		RespDesc          string                                 `json:"resp_desc"`
		ProductId         string                                 `json:"product_id"`
		HuifuId           string                                 `json:"huifu_id"`
		OrgHfSeqId        string                                 `json:"org_hf_seq_id"`
		OrgReqSeqId       string                                 `json:"org_req_seq_id"`
		OrgReqDate        string                                 `json:"org_req_date"`
		ReqDate           string                                 `json:"req_date"`
		ReqSeqId          string                                 `json:"req_seq_id"`
		OrdAmt            string                                 `json:"ord_amt"`
		FeeAmt            string                                 `json:"fee_amt"`
		TransFinishTime   string                                 `json:"trans_finish_time"`
		TransStat         string                                 `json:"trans_stat"`
		Remark            string                                 `json:"remark"`
		BankCode          string                                 `json:"bank_code"`
		BankMessage       string                                 `json:"bank_message"`
		UnconfirmAmt      string                                 `json:"unconfirm_amt"`
		ConfirmedAmt      string                                 `json:"confirmed_amt"`
		FqMerDiscountFlag string                                 `json:"fq_mer_discount_flag"`
		CombinedpayData   common.StringObject[[]CombinedpayData] `json:"combinedpay_data"`
		AcctSplitBunch    common.StringObject[AcctSplitBunch]    `json:"acct_split_bunch"`
	} `json:"data"`
}
type V2TradeOnlinepaymentRefundNotifyMessage struct {
	RespCode string                                                           `json:"resp_code"`
	RespDesc string                                                           `json:"resp_desc"`
	Data     common.StringObject[V2TradeOnlinepaymentRefundNotifyMessageData] `json:"data"`
	Sign     string                                                           `json:"sign"`
}
type V2TradeOnlinepaymentRefundNotifyMessageData struct {
	RespCode          string `json:"resp_code"`
	RespDesc          string `json:"resp_desc"`
	ProductId         string `json:"product_id"`
	HuifuId           string `json:"huifu_id"`
	MerName           string `json:"mer_name"`
	ReqSeqId          string `json:"req_seq_id"`
	ReqDate           string `json:"req_date"`
	HfSeqId           string `json:"hf_seq_id"`
	TransType         string `json:"trans_type"`
	TransDate         string `json:"trans_date"`
	TransTime         string `json:"trans_time"`
	OrgOrdAmt         string `json:"org_ord_amt"`
	OrgFeeAmt         string `json:"org_fee_amt"`
	TotalRefAmt       string `json:"total_ref_amt"`
	TotalRefFeeAmt    string `json:"total_ref_fee_amt"`
	OrdAmt            string `json:"ord_amt"`
	FeeAmt            string `json:"fee_amt"`
	TransFinishTime   string `json:"trans_finish_time"`
	TransStat         string `json:"trans_stat"`
	RefCut            string `json:"ref_cut"`
	Remark            string `json:"remark"`
	BankCode          string `json:"bank_code"`
	BankMessage       string `json:"bank_message"`
	FqFeeAmt          string `json:"fq_fee_amt"`
	FqMerDiscountFlag string `json:"fq_mer_discount_flag"`
	AcctSplitBunch    AcctSplitBunch
	SplitFeeInfo      SplitFeeInfo
	combinedPayData   []CombinedpayData
	YljfqResponse     YljfqResponse
}
type YljfqResponse struct {
	ActualRefAmt    string
	PromotionDetail []struct {
		PromotionId        string `json:"promotion_id,omitempty"`        //优惠券ID	String	32	N
		Name               string `json:"name,omitempty"`                //优惠名称	String	64	N
		Scope              string `json:"scope,omitempty"`               //优惠范围	String	32	N
		Type               string `json:"type,omitempty"`                //优惠类型	String	32	N
		Amount             string `json:"amount,omitempty"`              //优惠券面额	String	14	N
		ActivityId         string `json:"activity_id,omitempty"`         //活动ID	String	32	N
		MerchantContribute string `json:"merchant_contribute,omitempty"` //商户出资	String	1	N	单位元，需保留小数点后两位，示例值：1000.00
		OtherContribute    string `json:"other_contribute,omitempty"`    //其他出资	String	14	N	单位元，需保留小数点后两位，示例值：1000.00
	}
}
