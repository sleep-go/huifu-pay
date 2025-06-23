package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeAcctpaymentPayQuery 余额支付查询
// POST https://api.huifu.com/v2/trade/acctpayment/pay/query
//
// 应用场景
// 本接口用于查询商户单笔余额支付交易状态及明细。
//
// 适用对象
// 已开通余额支付相关功能的商户。
func (t *Trade) V2TradeAcctpaymentPayQuery(req V2TradeAcctpaymentPayQueryRequest) (res *V2TradeAcctpaymentPayQueryResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeAcctpaymentPayQueryRequest(BsPaySdk.V2TradeAcctpaymentPayQueryRequest{
		HuifuId:     req.HuifuId,
		OrgReqDate:  req.OrgReqDate,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeAcctpaymentPayQueryResponse](resp)
}

type V2TradeAcctpaymentPayQueryRequest struct {
	HuifuId     string `json:"huifu_id,omitempty"`
	OrgReqDate  string `json:"org_req_date,omitempty"`
	ExtendInfos struct {
		OrgReqSeqId string `json:"org_req_seq_id,omitempty"`
		OrgHfSeqId  string `json:"org_hf_seq_id" json:"org_hf_seq_id,omitempty"`
	}
}
type V2TradeAcctpaymentPayQueryResponse struct {
	Data struct {
		RespCode         string                              `json:"resp_code"`
		RespDesc         string                              `json:"resp_desc"`
		OrgReqSeqId      string                              `json:"org_req_seq_id"`
		OrgReqDate       string                              `json:"org_req_date"`
		HuifuId          string                              `json:"huifu_id"`
		OrdAmt           string                              `json:"ord_amt"`
		AcctSplitBunch   common.StringObject[AcctSplitBunch] `json:"acct_split_bunch"`
		OutHuifuId       string                              `json:"out_huifu_id"`
		OrgHfSeqId       string                              `json:"org_hf_seq_id"`
		TransStat        string                              `json:"trans_stat"`
		TransFinishTime  string                              `json:"trans_finish_time"`
		Remark           string                              `json:"remark"`
		GoodDesc         string                              `json:"good_desc"`
		DelayAcctFlag    string                              `json:"delay_acct_flag"`     //延时标记	String	1	N	是否延时交易，Y为延迟，N为不延迟，不传默认N；示例值：N
		HycFlag          string                              `json:"hyc_flag"`            //灵活用工标志	String	1	N	灵活用工标志 Y：灵活用工，N：非灵活用工(默认)
		LgPlatformType   string                              `json:"lg_platform_type"`    //灵活用工平台	String	3	N	灵活用工平台 LJH：乐接活，HYC：汇优财(默认)。仅在hyc_flag=Y时起作用
		HycAttachId      string                              `json:"hyc_attach_id"`       //灵活用工代发批次号	String	12	N	灵活用工代发批次号，灵活用工平台为汇优财时返回！
		TransFeeTakeFlag string                              `json:"trans_fee_take_flag"` //手续费承担方标识	String	4	C	余额支付手续费承担方标识；商户余额支付扣收规则为接口指定承担方时必填！枚举值：
		UnconfirmAmt     string                              `json:"unconfirm_amt"`       //待确认总金额	String	14	N	单位元，需保留小数点后两位，示例值：1.00
		ConfirmedAmt     string                              `json:"confirmed_amt"`       //已确认总金额	String	14	N	单位元，需保留小数点后两位，示例值：1.00
	} `json:"data"`
	Sign string `json:"sign"`
}
