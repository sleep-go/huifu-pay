package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeOnlinepaymentTransferAccount 银行大额支付
// POST https://api.huifu.com/v2/trade/onlinepayment/transfer/account
//
// 应用场景
// 现有快捷、网银基于银行、个人限额的限制，无法满足商户大额充值、支付的需求。银行大额支付接口可以满足商户、用户大额充值、支付需求。
//
// 适用对象
// 因其它支付方式的交易限额原因无法实现一次性大额交易的商户或用户。目前银行大额支付有两种模式，支持的银行不同；
//
// 备付金模式支持银行清单
// 苏宁银行模式支持银行清单
// 注：银行大额支付不支持实时分账；
//
// 接口说明
// 支持格式：JSON
// 加签验签说明：请参考接入指引-开发指南
// 调用流程说明
// 大额支付分为五种模式：
//
// 备付金订单下单模式：调用银行大额支付接口创建订单，汇付返回的入账标识；详见说明；
// 备付金绑卡充值/支付模式：线上网银/线下银行柜台转账充值；详见说明；
// 备付金固定转账标识模式：生成固定收款账号进行转账支付；详见说明；
// 银行(H5页面下单)模式：调用银行大额支付接口，创建订单，汇付返回H5下单页面；详见说明；
// 银行订单下单模式：调用银行大额支付接口，获取动态/固态转账账号；详见说明；
func (t *Trade) V2TradeOnlinepaymentTransferAccount(req V2TradeOnlinepaymentTransferAccountRequest) (res *V2TradeOnlinepaymentTransferAccountResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeOnlinepaymentTransferAccountRequest(BsPaySdk.V2TradeOnlinepaymentTransferAccountRequest{
		ReqDate:     req.ReqDate,
		ReqSeqId:    req.ReqSeqId,
		HuifuId:     req.HuifuId,
		GoodsDesc:   req.GoodsDesc,
		TransAmt:    req.TransAmt,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeOnlinepaymentTransferAccountResponse](resp)
}

type V2TradeOnlinepaymentTransferAccountExtendInfo struct {
	UserHuifuId          string `json:"user_huifu_id"`
	AcctId               string `json:"acct_id"`
	CertificateName      string `json:"certificate_name"`        //付款方名称	String	64	N	bank_mode=SNYH，且page_flag=N时必填；示例值：上海汇付支付有限公司
	BankCardNo           string `json:"bank_card_no"`            //付款方银行卡号	String	2048	N	bank_mode=SNYH，且page_flag=N时必填；原文最大为19位，密文最大长度为2048；使用斗拱公钥做RSA加密； 示例值：b9LE5RccVVLChrHgo9lvp……PhWhjKrWg2NPfbe0mkQ== 支持的银行见清单；
	Remark               string `json:"remark"`                  //备注	String	1024	N	示例值：采购原料
	PageFlag             string `json:"page_flag"`               //页面标识	String	1	N	Y-返回页面url，N-不返回页面url，默认：Y bank_mode=SNYH生效，示例值：Y
	OrderType            string `json:"order_type"`              //订单类型	String	1	N	P-支付，R-充值，默认：P；示例值：P
	NotifyUrl            string `json:"notify_url"`              //异步通知地址	String	512	N	以'http'或者'https'开头， 示例值：http://www.huifu.com/getResp
	BankMode             string `json:"bank_mode"`               // BFJ-备付金模式(默认模式)，SNYH-银行模式
	AcctMode             string `json:"acct_mode"`               //入账模式	String	8	N	bank_mode为银行模式时，必填
	DelayAcctFlag        string `json:"delay_acct_flag"`         //延时标记	String	1	N	Y：延迟，银行大额支付仅支持延迟分账；示例值：Y，不传时则不分账 延时交易要调【交易确认】接口资金才能进入收款方账户，否则会停留在延时账户中。
	OrderMode            string `json:"order_mode"`              //订单模式	String	2	N	O-订单下单模式，
	OrgRemittanceOrderId string `json:"org_remittance_order_id"` //原汇款订单号	String	128	N	order_mode=B或order_mode=FD时，必填； 示例值：2022012614120615000
	DynamicFlag          string `json:"dynamic_flag"`            //动态码标识	String	1	N	bank_mode=SNYH，page_flag=N 时生效 Y-动态，N-静态，默认：N
	TimeExpire           string `json:"time_expire"`             //订单失效时间	String	14	N	格式：yyyyMMddHHmmss；示例值：20220925102245，若不填写：2个工作日自动关单
	FeeFlag              string `json:"fee_flag"`                //手续费扣款标志	String	1	N	1: 外扣 2: 内扣 (默认取控台配置值)；示例值：1
}
type V2TradeOnlinepaymentTransferAccountRequest struct {
	ReqSeqId    string `json:"req_seq_id,omitempty"` // 请求流水号
	ReqDate     string `json:"req_date,omitempty"`   // 请求日期
	HuifuId     string `json:"huifu_id,omitempty"`   // 收款方商户号
	TransAmt    string `json:"trans_amt,omitempty"`  // 交易金额
	GoodsDesc   string `json:"goods_desc,omitempty"` // 商品描述
	ExtendInfos V2TradeOnlinepaymentTransferAccountExtendInfo
}
type V2TradeOnlinepaymentTransferAccountResponse struct {
	Data struct {
		RespCode          string `json:"resp_code"`
		RespDesc          string `json:"resp_desc"`
		ReqSeqId          string `json:"req_seq_id"`
		ReqDate           string `json:"req_date"`
		HfSeqId           string `json:"hf_seq_id"`
		HuifuId           string `json:"huifu_id"`
		TransAmt          string `json:"trans_amt"`
		FeeAmt            string `json:"fee_amt"`
		TransStat         string `json:"trans_stat"`
		GoodsDesc         string `json:"goods_desc"`
		Remark            string `json:"remark"`
		InAcctFlag        string `json:"in_acct_flag"`
		FormUrl           string `json:"form_url"`
		SubBaseAcctNo     string `json:"sub_base_acct_no"`
		SubBaseAcctNoName string `json:"sub_base_acct_no_name"`
		SubBankName       string `json:"sub_bank_name"`
		DelayAcctFlag     string `json:"delay_acct_flag"`
		DynamicFlag       string `json:"dynamic_flag"`
		FeeFlag           string `json:"fee_flag"`
	} `json:"data"`
	Sign string `json:"sign"`
}

// V2TradeOnlinepaymentTransferAccountNotifyMessage 异步返回参数
type V2TradeOnlinepaymentTransferAccountNotifyMessage struct {
	RespCode string                                                                        `json:"resp_code"`
	RespData common.StringObject[V2TradeOnlinepaymentTransferAccountNotifyMessageRespData] `json:"resp_data"`
	RespDesc string                                                                        `json:"resp_desc"`
	Sign     string                                                                        `json:"sign"`
}
type V2TradeOnlinepaymentTransferAccountNotifyMessageRespData struct {
	V3TradePaymentJspayNotifyMessageRespData
	DebitFlag       string          `json:"debit_flag"`
	UserHuifuId     string          `json:"user_huifu_id"`
	BankId          string          `json:"bank_id"`
	BankExtendParam BankExtendParam `json:"bank_extend_param"`
	OrderType       string          `json:"order_type"` //订单类型	String	1	N	P-支付 R-充值 默认：P-支付；示例值：P
	RequestIp       string          `json:"request_ip"`
	CertificateName string          //收款方名称	String	64	Y	示例值：上海汇付支付支付公司
	BankCardNo      string          //银行卡号	String	19	Y	密文最大长度为2048；使用斗拱公钥做RSA加密；
}
