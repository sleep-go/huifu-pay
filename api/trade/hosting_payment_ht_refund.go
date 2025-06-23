package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeHostingPaymentHtrefund 托管交易退款
// POST https://api.huifu.com/v2/trade/hosting/payment/htRefund
//
// 应用场景
// 通过托管支付预下单接口发起的交易可以通过本接口退款。
func (t *Trade) V2TradeHostingPaymentHtrefund(req V2TradeHostingPaymentHtrefundRequest) (res *V2TradeHostingPaymentHtrefundResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeHostingPaymentHtrefundRequest(BsPaySdk.V2TradeHostingPaymentHtrefundRequest{
		ReqDate:            req.ReqDate,
		ReqSeqId:           req.ReqSeqId,
		HuifuId:            req.HuifuId,
		OrdAmt:             req.OrdAmt,
		OrgReqDate:         req.OrgReqDate,
		RiskCheckData:      req.RiskCheckData.Encode(),
		TerminalDeviceData: req.TerminalDeviceData.Encode(),
		BankInfoData:       req.BankInfoData,
		ExtendInfos:        common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeHostingPaymentHtrefundResponse](resp)
}

type BankInfoData struct {
	Province          string `json:"province,omitempty"`           //省份	String	4	C	付款方为对公账户时必填，参见省市地区码；示例值：0013
	Area              string `json:"area,omitempty"`               //地区	String	4	C	付款方为对公账户时必填，参见省市地区码；示例值：1301
	BankCode          string `json:"bank_code,omitempty"`          //银行编号	String	8	C	付款方为对公账户时必填，参考：银行编码； 示例值：01040000
	CorrespondentCode string `json:"correspondent_code,omitempty"` //联行号	String	30	C	付款方为对公账户时必填，参见：银行支行编码； 示例值：102290026507
	CardAcctType      string `json:"card_acct_type,omitempty"`     //付款方账户类型	String	1	N	对公:E，对私:P；默认：P，示例值：P
}
type V2TradeHostingPaymentHtrefundExtendInfo struct {
	OrgHfSeqId      string         `json:"org_hf_seq_id"`
	OrgPartyOrderId string         `json:"org_party_order_id"`
	OrgReqSeqId     string         `json:"org_req_seq_id"`
	AcctSplitBunch  AcctSplitBunch `json:"acct_split_bunch"`
	Remark          string         `json:"remark"`
	NotifyUrl       string         `json:"notify_url"`
	BankInfoData    BankInfoData   `json:"bank_info_data"`
}
type V2TradeHostingPaymentHtrefundRequest struct {
	ReqDate            string                                  `json:"req_date" structs:"req_date"`                         // 请求日期
	ReqSeqId           string                                  `json:"req_seq_id" structs:"req_seq_id"`                     // 请求流水号
	HuifuId            string                                  `json:"huifu_id" structs:"huifu_id"`                         // 商户号
	OrdAmt             string                                  `json:"ord_amt" structs:"ord_amt"`                           // 申请退款金额
	OrgReqDate         string                                  `json:"org_req_date" structs:"org_req_date"`                 // 原交易请求日期
	RiskCheckData      common.StringObject[RiskCheckData]      `json:"risk_check_data" structs:"risk_check_data"`           // 安全信息线上交易退款必填，参见线上退款接口；jsonObject字符串
	TerminalDeviceData common.StringObject[TerminalDeviceData] `json:"terminal_device_data" structs:"terminal_device_data"` // 设备信息线上交易退款必填，参见线上退款接口；jsonObject字符串
	BankInfoData       string                                  `json:"bank_info_data" structs:"bank_info_data"`             // 大额转账支付账户信息数据jsonObject格式；银行大额转账支付交易退款申请时必填
	ExtendInfos        V2TradeHostingPaymentHtrefundExtendInfo
}
type V2TradeHostingPaymentHtrefundResponse struct {
	Data struct {
		RespCode         string                                       `json:"resp_code"`
		RespDesc         string                                       `json:"resp_desc"`
		HuifuId          string                                       `json:"huifu_id"`
		MerName          string                                       `json:"mer_name"`
		ReqDate          string                                       `json:"req_date"`
		ReqSeqId         string                                       `json:"req_seq_id"`
		HfSeqId          string                                       `json:"hf_seq_id"`
		OrgReqDate       string                                       `json:"org_req_date"`
		OrgReqSeqId      string                                       `json:"org_req_seq_id"`
		OrgHfSeqId       string                                       `json:"org_hf_seq_id"`
		TransTime        string                                       `json:"trans_time"`
		TransStat        string                                       `json:"trans_stat"`
		OrdAmt           string                                       `json:"ord_amt"`
		ActualRefAmt     string                                       `json:"actual_ref_amt"`
		AcctSplitBunch   common.StringObject[AcctSplitBunch]          `json:"acct_split_bunch"`
		UnionpayResponse common.StringObject[common.UnionpayResponse] `json:"unionpay_response"`
		Remark           string                                       `json:"remark"`
		BankCode         string                                       `json:"bank_code"`
		BankMessage      string                                       `json:"bank_message"`
		FeeAmt           string                                       `json:"fee_amt"`
	} `json:"data"`
	Sign string `json:"sign"`
}

type V2TradeHostingPaymentHtrefundNotifyMessage struct {
	RespCode string                                                                  `json:"resp_code"`
	RespData common.StringObject[V2TradeHostingPaymentHtrefundNotifyMessageRespData] `json:"resp_data"`
	RespDesc string                                                                  `json:"resp_desc"`
	Sign     string                                                                  `json:"sign"`
}
type V2TradeHostingPaymentHtrefundNotifyMessageRespData struct {
	RespCode         string                  `json:"resp_code,omitempty"`
	RespDesc         string                  `json:"resp_desc,omitempty"`
	HuifuId          string                  `json:"huifu_id,omitempty"`
	ReqDate          string                  `json:"req_date,omitempty"`          //请求日期	String	8	Y	格式为yyyyMMdd，示例值：20220925
	ReqSeqId         string                  `json:"req_seq_id,omitempty"`        //请求流水号	String	128	Y	示例值：rQ202112131117394413651
	HfSeqId          string                  `json:"hf_seq_id,omitempty"`         //全局流水号	String	128	N	示例值：0030default220825182711P099ac1f343f00000
	OrgReqDate       string                  `json:"org_req_date,omitempty"`      //原交易请求日期	String	8	N	格式为yyyyMMdd，示例值：20220925
	OrgReqSeqId      string                  `json:"org_req_seq_id,omitempty"`    //原交易请求流水号	String	128	N	示例值：rQ202112131149875651
	OrgOrdAmt        string                  `json:"org_ord_amt,omitempty"`       //	原交易订单金额	String	14	Y	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
	OrgFeeAmt        string                  `json:"org_fee_amt,omitempty"`       //原交易手续费	String	14	Y	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
	TransDate        string                  `json:"trans_date,omitempty"`        //退款交易发生日期	String	8	Y	格式为yyyyMMdd，示例值：20220925
	TransTime        string                  `json:"trans_time,omitempty"`        //退款交易发生时间	String	6	N	格式：HHMMSS，示例值：091010；9点10分10秒
	TransFinishTime  string                  `json:"trans_finish_time,omitempty"` //退款完成时间	String	14	N	格式yyyyMMddHHmmss；示例值：20091225091010
	TransType        string                  `json:"trans_type,omitempty"`        //交易类型	String	40	Y	TRANS_REFUND：交易退款；目前仅该一个枚举值； 示例值：TRANS_REFUND
	TransStat        string                  `json:"trans_stat,omitempty"`        //交易状态	String	1	N	P：处理中、S：成功、F：失败；示例值：S
	OrdAmt           string                  `json:"ord_amt,omitempty"`           //退款金额	String	14	Y	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
	ActualRefAmt     string                  `json:"actual_ref_amt,omitempty"`    //实际退款金额	String	14	N	扫码交易返回；示例值：111.00
	TotalRefAmt      string                  `json:"total_ref_amt,omitempty"`     //原交易累计退款金额	String	14	Y	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
	TotalRefFeeAmt   string                  `json:"total_ref_fee_amt,omitempty"` //原交易累计退款手续费金额	String	14	Y	单位元，示例值：1.00；注意：退还手续费规则参见说明文档
	RefCut           string                  `json:"ref_cut,omitempty"`           //累计退款次数	String	14	Y	示例值：1
	AcctSplitBunch   AcctSplitBunch          `json:"acct_split_bunch,omitempty"`  //分账信息	String	4000	Y	与同步返参的相同
	SplitFeeInfo     SplitFeeInfo            `json:"split_fee_info,omitempty"`    //分账手续费信息	String	2048	N	线上交易返回
	PartyOrderId     string                  `json:"party_order_id,omitempty"`    //微信支付宝的商户单号	String	64	N	示例值：03232109190255105603561；参见用户账单说明
	FeeAmt           string                  `json:"fee_amt,omitempty"`           //退款返还手续费	String	14	N	线上交易返回
	Remark           string                  `json:"remark,omitempty"`            //备注	String	84	N	原样返回；示例值：备注
	BankCode         string                  `json:"bank_code,omitempty"`         //通道返回码	String	64	N	示例值：01020000
	BankMessage      string                  `json:"bank_message,omitempty"`      //通道返回描述	String	256	N	示例值：SUCCESS
	UnionpayResponse common.UnionpayResponse `json:"unionpay_response,omitempty"` //银联返回的响应报文	String	6000	N	JsonObject格式；
	BankId           string                  `json:"bank_id,omitempty"`           //银行编号	String	32	N	线上交易返回；示例值：
	BankName         string                  `json:"bank_name,omitempty"`         //银行名称	String	128	N	线上交易返回；示例值：
}
