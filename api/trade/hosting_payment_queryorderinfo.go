package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeHostingPaymentQueryorderinfo 托管交易查询
// POST https://api.huifu.com/v2/trade/hosting/payment/queryorderinfo
//
// 应用场景
// 通过托管支付预下单接口支付的交易都可以通过本接口查询。
func (t *Trade) V2TradeHostingPaymentQueryorderinfo(req V2TradeHostingPaymentQueryorderinfoRequest) (res *V2TradeHostingPaymentQueryorderinfoResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeHostingPaymentQueryorderinfoRequest(BsPaySdk.V2TradeHostingPaymentQueryorderinfoRequest{
		ReqDate:     req.ReqDate,
		ReqSeqId:    req.ReqSeqId,
		HuifuId:     req.HuifuId,
		OrgReqDate:  req.OrgReqDate,
		OrgReqSeqId: req.OrgReqSeqId,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeHostingPaymentQueryorderinfoResponse](resp)
}

type V2TradeHostingPaymentQueryorderinfoRequest struct {
	ReqDate     string `json:"req_date" structs:"req_date"`             // 请求日期
	ReqSeqId    string `json:"req_seq_id" structs:"req_seq_id"`         // 请求流水号
	HuifuId     string `json:"huifu_id" structs:"huifu_id"`             // 商户号
	OrgReqDate  string `json:"org_req_date" structs:"org_req_date"`     // 原交易请求日期
	OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原交易请求流水号
}
type V2TradeHostingPaymentQueryorderinfoResponse struct {
	RespCode              string                  //业务响应码	String	8	Y	参见业务返回码；示例值：00000000
	RespDesc              string                  //业务响应信息	String	512	Y	参见业务返回码；示例值：操作成功
	HuifuId               string                  //商户号	String	32	Y	示例值：6666000123123123
	ReqDate               string                  //请求日期	String	8	Y	请求格式：yyyyMMdd；示例值：20221023
	ReqSeqId              string                  //请求流水号	String	64	Y	同一huifu_id下当天唯一；示例值：rQ2021121311173944
	OrgReqDate            string                  //原机构请求日期	String	8	Y	格式为yyyyMMdd，示例值：20220125
	OrgReqSeqId           string                  //原机构请求流水号	String	128	N	示例值：202110210012100005
	OrgHfSeqId            string                  //斗拱返回的全局流水号	String	128	N	示例值：00290TOP1GR210919004230P853ac13262200000
	PreOrderId            string                  //预下单订单号	String	64	Y	示例值：H2022050711225100912503288
	OrderStat             string                  //预下单状态	String	1	N	1:支付成功,2:支付中,3:已退款,4:处理中,5:支付失败,6-部分退款 参见状态说明文档；示例值：1
	PartyOrderId          string                  //用户账单上的商户订单号	String	64	N	示例值：03232109190255105603561；参见用户账单说明
	OrdAmt                string                  //订单金额(元)	String	14	Y	保留小数点后两位，示例值：1.00，最低传入0.01
	TransDate             string                  //订单日期	String	8	Y	格式：yyyymmdd；示例值：20221023
	TransAmt              string                  //交易金额(元)	String	14	Y	保留小数点后两位，示例值：1.00，最低传入0.01
	PayType               string                  ////交易类型	String	16	N	T_JSAPI：微信公众号支付
	TransStat             string                  //交易状态	String	1	N	P：处理中；S：成功；F：失败；I: 初始；示例值：S 初始状态很罕见，请联系汇付技术人员处理
	TransTime             string                  //交易时间	String	14	N	格式：yyyymmddHHMMSS；示例值：20231112200913
	CloseStat             string                  //关单状态	String	1	N	P：处理中、S：成功、F：失败；示例值：S
	FeeFlag               string                  //手续费扣款标志	Int	1	N	1: 外扣，2: 内扣；默认返回控台配置方式；示例值：2
	FeeAmt                string                  //手续费金额(元)	String	14	N	保留小数点后两位，最低传入0.01。示例值：1.00
	RefAmt                string                  //可退金额(元)	String	14	N	示例值：1.00
	GoodsDesc             string                  //商品描述	String	40	N	示例值：电脑PC
	Remark                string                  //备注	String	255	N	原样返回；示例值：备注
	MerPriv               string                  //商户私有域	String	1500	N	示例值：商户私有域
	BankCode              string                  //外部通道返回码	String	32	N	示例值：TRADE_SUCCESS
	BankDesc              string                  //外部通道返回描述	String	200	N	示例值：TRADE_SUCCESS
	WxResponse            WxResponse              //微信返回的响应报文	String	6000	N	jsonObject格式
	AlipayResponse        AlipayResponse          //支付宝返回的响应报文	String	6000	N	jsonObject格式
	UnionpayResponse      common.UnionpayResponse //银联返回的响应报文	String	6000	N	jsonObject格式
	WxUserId              string                  //微信用户唯一标识码	String	128	N	示例值：W6NYVcMwXDfAT+3LXuLSMx+UH5AXx1kG7JzTiTEomdk=
	IsDiv                 string                  //是否分账交易	String	1	Y	Y: 分账交易, N: 非分账交易；示例值：N
	AcctSplitBunch        AcctSplitBunch          //分账对象	String	2048	N	参见现在H5/pc托管异步中的字段
	IsDelayAcct           string                  //是否延时交易	String	1	Y	Y: 延迟 N: 不延迟；示例值：N
	TransFeeAllowanceInfo TransFeeAllowanceInfo   //手续费补贴信息	String	6000	N	参见现在H5/pc托管异步中的字段
	QuickOnlineResponse   QuickOnlineResponse     //快捷网银响应	String	6000	N	新设计的域字段
	BankExtendParam       BankExtendParam         //银行扩展信息	String		N	jsonObject格式；网银返回
	DevsId                string                  //汇付机具号	String	32	Y	通过汇付报备的机具必传；示例值：660035730311000126101
	OutTransId            string                  //交易单号	String	64	N	用户账单上的交易订单号，示例值：092021091922001451301445517582；参见用户账单说明
	RequestIp             string                  //请求IP	String	15	N	付款方IP,仅在支付成功后返回;示例：192.168.1.1
}
type QuickOnlineResponse struct {
	DebitFlag   string `json:"debit_flag,omitempty"`    //借贷记标识	String	1	N	D-借记卡 C-信用卡 Z-借贷合一卡；示例值：C
	UserHuifuId string `json:"user_huifu_id,omitempty"` //用户号	String	32	N	汇付分配的用户号快捷支付时才有值；示例值：6666000103905031
	OrderType   string `json:"order_type,omitempty"`    //订单类型	String	1	N	P-支付 R-充值 默认：P-支付；示例值：P
	BankId      string `json:"bank_id,omitempty"`       //银行代号	String	8	N	示例值：30200000
}
type BankExtendParam struct {
	GateType   string `json:"gate_type,omitempty"`    //网关支付类型	String	2	N	01: 个人网关02:企业网关；示例值：02
	BankId     string `json:"bank_id,omitempty"`      //付款方银行号	String	32	N	示例值：01040000
	PyerAcctId string `json:"pyer_acct_id,omitempty"` //付款方银行账户	String	1024	N	B2B支付成功后可能返回密文；示例值：6226560119005276
	PyerAcctNm string `json:"pyer_acct_nm,omitempty"` //付款方银行账户名	String	128	N	示例值：上海汇付支付有限公司
}
