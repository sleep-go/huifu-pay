package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeAcctpaymentPay 余额支付
// POST https://api.huifu.com/v2/trade/acctpayment/pay
//
// 应用场景
// 本接口用于商户使用预存金额进行交易。
//
// 适用对象
// 已开通余额支付相关功能的商户。
func (t *Trade) V2TradeAcctpaymentPay(req V2TradeAcctpaymentPayRequest) (res *V2TradeAcctpaymentPayResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeAcctpaymentPayRequest(BsPaySdk.V2TradeAcctpaymentPayRequest{
		ReqSeqId:         req.ReqSeqId,
		ReqDate:          req.ReqDate,
		OutHuifuId:       req.OutHuifuId,
		OrdAmt:           req.OrdAmt,
		FundType:         req.FundType,
		TransFeeTakeFlag: req.TransFeeTakeFlag,
		ExtendInfos:      common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeAcctpaymentPayResponse](resp)
}

type V2TradeAcctpaymentPayExtendInfo struct {
	NotifyUrl       string
	HuifuId         string
	GoodDesc        string
	Remark          string
	DelayAcctFlag   string
	OutAcctId       string
	AcctChannel     string
	HycFlag         string
	LgPlatformType  string
	SalaryModleType string
	BmemberId       string
	AcctSplitBunch  AcctSplitBunch
	RiskCheckData   RiskCheckData
	LjhData         common.LjhData
}
type V2TradeAcctpaymentPayRequest struct {
	ReqSeqId         string `json:"req_seq_id,omitempty"`          // 请求流水号
	ReqDate          string `json:"req_date,omitempty"`            // 请求日期
	OutHuifuId       string `json:"out_huifu_id,omitempty"`        // 出款方商户号
	OrdAmt           string `json:"ord_amt,omitempty"`             // 支付金额
	FundType         string `json:"fund_type,omitempty"`           // 资金类型资金类型。支付渠道为中信E管家时，资金类型必填（[详见说明](https://paas.huifu.com/open/doc/api/#/yuer/api_zxegjzllx)）
	TransFeeTakeFlag string `json:"trans_fee_take_flag,omitempty"` // 手续费承担方标识余额支付手续费承担方标识；商户余额支付扣收规则为接口指定承担方时必填！枚举值：&lt;br/&gt;OUT：出款方；&lt;br/&gt;IN：分账接受方。&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：IN&lt;/font&gt;
	ExtendInfos      V2TradeAcctpaymentPayExtendInfo
}
type V2TradeAcctpaymentPayResponse struct {
	Data struct {
		RespCode         string                              `json:"resp_code"`
		RespDesc         string                              `json:"resp_desc"`
		ReqSeqId         string                              `json:"req_seq_id"`
		ReqDate          string                              `json:"req_date"`
		HuifuId          string                              `json:"huifu_id"`
		OutHuifuId       string                              `json:"out_huifu_id"`
		OrdAmt           string                              `json:"ord_amt"`
		AcctSplitBunch   common.StringObject[AcctSplitBunch] `json:"acct_split_bunch"`
		HfSeqId          string                              `json:"hf_seq_id"`
		GoodDesc         string                              `json:"good_desc"`
		TransStat        string                              `json:"trans_stat"`
		TransFinishTime  string                              `json:"trans_finish_time"`
		Remark           string                              `json:"remark"`
		OutAcctId        string                              `json:"out_acct_id"`
		FundType         string                              //资金类型	String	16	C	资金类型。支付渠道为中信E管家时，资金类型必填（详见说明）
		AcctChannel      string                              //支付渠道	String	8	N	支付渠道。HUIFU=汇付账务(默认)；ZXE=中信E管家
		HycFlag          string                              //灵活用工标志	String	1	N	Y：灵活用工，N：非灵活用工(默认)；示例值：Y
		LgPlatformType   string                              //灵活用工平台	String	3	N	灵活用工平台 LJH：乐接活，HYC：汇优财(默认)。仅在hyc_flag=Y时起作用
		SalaryModleType  string                              //代发模式	String	1	N	1：普票，2：专票。灵活用工平台为汇优财时必填！示例值：1
		BmemberId        string                              //落地公司商户号	String	18	N	灵活用工平台为汇优财时必填！示例值：6666000109812124
		HycAttachId      string                              //灵活用工代发批次号	String	12	N	灵活用工平台为汇优财时必填！ 示例值：
		LjhResponse      common.StringObject[common.LjhData] //乐接活返回参数集合	String		N	jsonObject字符串
		TransFeeTakeFlag string                              //手续费承担方标识	String	4	C	余额支付手续费承担方标识；商户余额支付扣收规则为接口指定承担方时必填！枚举值：
		UnconfirmAmt     string                              //待确认总金额	String	14	N	单位元，需保留小数点后两位，示例值：1.00
		ConfirmedAmt     string                              //已确认总金额	String	14	N	单位元，需保留小数点后两位，示例值：1.00
	} `json:"data"`
	Sign string `json:"sign"`
}

type V2TradeAcctpaymentPayNotifyMessage struct {
	RespCode string                                                          `json:"resp_code"`
	RespData common.StringObject[V2TradeAcctpaymentPayNotifyMessageRespData] `json:"resp_data"`
	RespDesc string                                                          `json:"resp_desc"`
	Sign     string
}
type V2TradeAcctpaymentPayNotifyMessageRespData struct {
	RespCode       string         `json:"resp_code,omitempty"` //业务返回码	String	8	Y	业务返回码
	RespDesc       string         `json:"resp_desc,omitempty"` //业务返回描述	String	512	Y	业务返回描述
	ReqDate        string         `json:"req_date"`
	ReqSeqId       string         `json:"req_seq_id"`
	HuifuId        string         `json:"huifu_id"`
	ProductId      string         //产品号	String	64	Y	产品号
	OrdAmt         string         //订单金额	String	14	Y	订单金额
	TransType      string         //交易类型	String	32	Y	余额支付：ACCT_PAYMENT
	TransStat      string         //交易状态	String	1	Y	PSF P:处理中；S：成功；F：失败；C：完成。 状态为完成时，需查看分账对象中每个分账的具体状态
	HycFlag        string         //灵活用工标志	String	1	N	灵活用工标志 Y：灵活用工，N：非灵活用工(默认)
	HycAttachId    string         //灵活用工代发批次号	String	12	N	灵活用工代发批次号，灵活用工平台为汇优财时返回！
	AcctSplitBunch AcctSplitBunch //分账对象	Object		Y	分账对象，jsonObject字符串
}
