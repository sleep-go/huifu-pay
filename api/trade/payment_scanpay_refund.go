package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradePaymentScanpayRefund 扫码交易退款查询
// POST https://api.huifu.com/v3/trade/payment/scanpay/refundquery
//
// 应用场景
// 商家调用退款接口未收到状态返回，调用本接口查询退款结果。支持微信公众号-T_JSAPI、小程序-T_MINIAPP、微信APP支付-T_APP、支付宝JS-A_JSAPI、支付宝正扫-A_NATIVE、银联二维码正扫-U_NATIVE、银联二维码JS-U_JSAP、数字货币二维码支付-D_NATIVE，以及微信反扫、支付宝反扫、银联二维码反扫、数字人民币反扫交易退款的查询。
//
// 适用对象
// 开通微信/支付宝/银联二维码/数字人民币聚合扫码功能的商户。
func (t *Trade) V2TradePaymentScanpayRefund(req V2TradePaymentScanpayRefundRequest) (res *V2TradePaymentScanpayRefundResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradePaymentScanpayRefundRequest(BsPaySdk.V2TradePaymentScanpayRefundRequest{
		ReqDate:     req.ReqDate,
		ReqSeqId:    req.ReqSeqId,
		HuifuId:     req.HuifuId,
		OrdAmt:      req.OrdAmt,
		OrgReqDate:  req.OrgReqDate,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradePaymentScanpayRefundResponse](resp)
}

type V2TradePaymentScanpayRefundExtendInfos struct {
	OrgHfSeqId          string         `json:"org_hf_seq_id"`
	OrgPartyOrderId     string         `json:"org_party_order_id"`
	OrgReqSeqId         string         `json:"org_req_seq_id"`
	AcctSplitBunch      AcctSplitBunch `json:"acct_split_bunch"`
	WxData              WxData         `json:"wx_data"`
	DigitalCurrencyData struct {
		RefundDesc string `json:"refund_desc"`
	} `json:"digital_currency_data"`
	CombinedpayData    []CombinedpayData
	Remark             string `json:"remark"`
	LoanFlag           string `json:"loan_flag"`
	LoanUndertaker     string `json:"loan_undertaker"` //垫资承担者	String	32	N	垫资方的huifu_id；示例值：6666000123123123 为空则各自承担。不为空走第三方垫资，目前支持商户垫资；
	LoanAcctType       string `json:"loan_acct_type"`  //垫资账户类型	String	2	N	01:基本户, 05: 充值户, 默认充值户；示例值：01
	RiskCheckData      RiskCheckData
	TerminalDeviceData TerminalDeviceData
	NotifyUrl          string
	UnionpayData       UnionpayData
}
type V2TradePaymentScanpayRefundRequest struct {
	ReqDate     string `json:"req_date"`
	ReqSeqId    string `json:"req_seq_id"`
	HuifuId     string `json:"huifu_id"`
	OrdAmt      string `json:"ord_amt"`
	OrgReqDate  string `json:"org_req_date"`
	ExtendInfos V2TradePaymentScanpayRefundExtendInfos
}
type V2TradePaymentScanpayRefundResponse struct {
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

type V2TradePaymentScanpayRefundNotifyMessage struct {
	RespCode string                                                                `json:"resp_code"`
	RespData common.StringObject[V2TradePaymentScanpayRefundNotifyMessageRespData] `json:"resp_data"`
	RespDesc string                                                                `json:"resp_desc"`
	Sign     string
}
type V2TradePaymentScanpayRefundNotifyMessageRespData struct {
	RespCode                 string                  `json:"resp_code,omitempty"`
	RespDesc                 string                  `json:"resp_desc,omitempty"`
	ProductId                string                  `json:"product_id,omitempty"`
	HuifuId                  string                  `json:"huifu_id,omitempty"`
	ReqDate                  string                  `json:"req_date,omitempty"`
	ReqSeqId                 string                  `json:"req_seq_id,omitempty"`
	HfSeqId                  string                  `json:"hf_seq_id,omitempty"`                    //全局流水号	String	128	N	示例值：00470topo1A221019132207P068ac1362af00000
	OrgReqDate               string                  `json:"org_req_date,omitempty"`                 //原交易请求日期	String	8	N	格式为yyyyMMdd，示例值：20220925
	OrgReqSeqId              string                  `json:"org_req_seq_id,omitempty"`               //原交易请求流水号	String	128	N	示例值：rQ202112131149875651
	OrgOrdAmt                string                  `json:"org_ord_amt,omitempty"`                  //原交易订单金额	String	14	Y	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
	OrgFeeAmt                string                  `json:"org_fee_amt,omitempty"`                  //原交易手续费	String	14	Y	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
	TransDate                string                  `json:"trans_date,omitempty"`                   //退款交易发生日期	String	8	Y	格式为yyyyMMdd，示例值：20220925
	TransTime                string                  `json:"trans_time,omitempty"`                   //退款交易发生时间	String	6	N	格式：HHMMSS，示例值：0910109点10分10秒
	TransFinishTime          string                  `json:"trans_finish_time,omitempty"`            //退款完成时间	String	14	N	格式yyyyMMddHHmmss；示例值：20091225091010
	TransType                string                  `json:"trans_type,omitempty"`                   //交易类型	String	40	Y	TRANS_REFUND：交易退款；目前仅该一个枚举值； 示例值：TRANS_REFUND
	TransStat                string                  `json:"trans_stat,omitempty"`                   //交易状态	String	1	N	P：处理中、S：成功、F：失败；示例值：
	OrdAmt                   string                  `json:"ord_amt,omitempty"`                      //退款金额	String	14	Y	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
	ActualRefAmt             string                  `json:"actual_ref_amt,omitempty"`               //实际退款金额	String	14	N	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
	TotalRefAmt              string                  `json:"total_ref_amt,omitempty"`                //原交易累计退款金额	String	14	Y	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
	TotalRefFeeAmt           string                  `json:"total_ref_fee_amt,omitempty"`            //原交易累计退款手续费金额	String	14	Y	单位元，示例值：1.00；注意：退还手续费规则参见说明文档
	RefCut                   string                  `json:"ref_cut,omitempty"`                      //累计退款次数	String	14	Y	示例值：1
	AcctSplitBunch           AcctSplitBunch          `json:"acct_split_bunch,omitempty"`             //分账信息	Object	4000	Y	分账信息
	PartyOrderId             string                  `json:"party_order_id,omitempty"`               //微信支付宝的商户单号	String	64	N	示例值：03232109190255105603561；参见用户账单说明
	WxResponse               WxResponse              `json:"wx_response,omitempty"`                  //微信返回的响应报文	Object	6000	N
	DcResponse               DcResponse              `json:"dc_response,omitempty"`                  //数字人民币响应报文	Object	2048	N	jsonObject格式
	AlipayResponse           AlipayResponse          `json:"alipay_response,omitempty"`              //支付宝返回的响应报文
	CombinedpayData          []CombinedpayData       `json:"combinedpay_data,omitempty"`             //补贴支付信息	Array		N	参见《补贴支付信息》
	Remark                   string                  `json:"remark,omitempty"`                       //备注	String	1500	N	原样返回；示例值：备注
	BankMessage              string                  `json:"bank_message,omitempty"`                 //通道返回描述	String	256	N	示例值：SUCCESS
	UnionpayResponse         common.UnionpayResponse `json:"unionpay_response,omitempty"`            //银联返回的响应报文	Object	6000	N	Json格式
	FundFreezeStat           string                  `json:"fund_freeze_stat,omitempty"`             //资金冻结状态	String	16	N	FREEZE：冻结；UNFREEZE：解冻；示例值：UNFREEZE 退款发生时，对应原交易的资金冻结状态。
	TransFeeRefAllowanceInfo TransFeeAllowanceInfo   `json:"trans_fee_ref_allowance_info,omitempty"` //手续费补贴返还信息	Object		N	手续费补贴返还信息对象，jsonObject字符串
	PayChannel               string                  `json:"pay_channel,omitempty"`                  //交易通道	String	1	N	枚举值：A-支付宝、T-微信、U-银联二维码、D-数字货币
}
