package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeOnlinepaymentUnionpay 银联统一在线收银台
// POST https://api.huifu.com/v2/trade/onlinepayment/unionpay
//
// 应用场景
// 第一种商户APP场景：调用银联统一在线收银台接口获取“银联受理订单号”，随后调用SDK唤起消费者在手机端已安装的云闪付APP或银行APP完成支付。
// 第二种场景：商户小程序、H5页面可以通过拉起银联云闪付小程序方式进行支付。 云闪付小程序支付演示视频
//
// 注意：
//
// 使用云闪付小程序支付每笔都需要验证支付密码+短信验证码；
// 有些手机型号跳入银联收银台之后不展示银行列表选择页面，是因为手机端未安装银联可信安全服务组件。下载银联可信安全服务组件即可。
// 银联线上收银台虽然同时支持APP支付与云闪付小程序支付，但需要分别申请，且会生成两个不同的银联商户号。目前斗拱不支持一个斗拱商户号同时绑定多个银联商户号，如果同时开通需要入驻两个斗拱商户。
// 使用云闪付小程序支付的，在获取pay_info、union_app_id、union_path后商户前端调用方式参见银联指引文档；
// 适用对象
// 商户需要开通银联统一在线收银台权限。
// 银联统一在线收银台支持的银行列表；
// 开发联调需要集成银联线上收银台SDK；
func (t *Trade) V2TradeOnlinepaymentUnionpay(req V2TradeOnlinepaymentUnionpayRequest) (res *V2TradeOnlinepaymentUnionpayResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeOnlinepaymentUnionpayRequest(BsPaySdk.V2TradeOnlinepaymentUnionpayRequest{
		HuifuId:       req.HuifuId,
		ReqDate:       req.ReqDate,
		ReqSeqId:      req.ReqSeqId,
		TransAmt:      req.TransAmt,
		OrderDesc:     req.GoodsDesc,
		RiskCheckData: req.RiskCheckData.Encode(),
		ThirdPayData:  req.ThirdPayData.Encode(),
		ExtendInfos:   common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeOnlinepaymentUnionpayResponse](resp)
}

type V2TradeOnlinepaymentUnionpayExtendInfo struct {
	CardNumberLock string         `json:"card_number_lock"`
	EbankEnAbbr    string         `json:"ebank_en_abbr"`
	OrderDesc      string         `json:"order_desc"`
	PayCardNo      string         `json:"pay_card_no"`
	PayCardType    string         `json:"pay_card_type"`
	TimeExpire     string         `json:"time_expire"`
	AcctSplitBunch AcctSplitBunch `json:"acct_split_bunch"`
	FrontUrl       string         `json:"front_url"`
	NotifyUrl      string         `json:"notify_url"`
	Remark         string         `json:"remark"`
	PayScene       string         `json:"pay_scene"`
	SignTokenNo    string         `json:"sign_token_no"`
	DelayAcctFlag  string         `json:"delay_acct_flag"`
	FeeFlag        string         `json:"fee_flag"`
}
type ThirdPayData struct {
	AppId string `json:"app_id"`
}
type V2TradeOnlinepaymentUnionpayRequest struct {
	ReqSeqId      string                             `json:"req_seq_id,omitempty"` // 请求流水号
	ReqDate       string                             `json:"req_date,omitempty"`   // 请求日期
	HuifuId       string                             `json:"huifu_id,omitempty"`   // 收款方商户号
	TransAmt      string                             `json:"trans_amt,omitempty"`  // 交易金额
	GoodsDesc     string                             `json:"goods_desc,omitempty"` // 商品描述
	RiskCheckData common.StringObject[RiskCheckData] `json:"risk_check_data,omitempty"`
	ThirdPayData  common.StringObject[ThirdPayData]  `json:"third_pay_data,omitempty"`
	ExtendInfos   V2TradeOnlinepaymentUnionpayExtendInfo
}
type V2TradeOnlinepaymentUnionpayResponse struct {
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

type V2TradeOnlinepaymentUnionpayNotifyMessage struct {
	RespCode string                                                                 `json:"resp_code"`
	RespData common.StringObject[V2TradeOnlinepaymentUnionpayNotifyMessageRespData] `json:"resp_data"`
	RespDesc string                                                                 `json:"resp_desc"`
	Sign     string                                                                 `json:"sign"`
}
type V2TradeOnlinepaymentUnionpayNotifyMessageRespData struct {
	SubRespCode    string         `json:"sub_resp_code,omitempty"`    //返回码	String	8	Y	参见业务返回码，示例值：00000000
	SubRespDesc    string         `json:"sub_resp_desc,omitempty"`    //返回描述	String	512	Y	参见业务返回信息，示例值：交易成功
	ReqSeqId       string         `json:"req_seq_id,omitempty"`       //请求流水号	String	64	Y	示例值：2021091708126665001
	ReqDate        string         `json:"req_date,omitempty"`         //请求日期	String	8	Y	格式：yyyyMMdd；示例值：20221120
	ProductId      string         `json:"product_id,omitempty"`       //汇付产品号	String	32	Y	汇付分配的产品号；示例值：YYZY
	HuifuId        string         `json:"huifu_id,omitempty"`         //汇付商户号	String	32	Y	汇付分配的商户号；示例值：6666000108854952
	TransAmt       string         `json:"trans_amt,omitempty"`        //订单金额	String	14	N	单位：元；示例值：1.23
	UnionOrderNo   string         `json:"union_order_no,omitempty"`   //银联受理订单	String	32	N	银联受理订单,银联移动支付系统返回该流水号，调用支付控件时使用 示例值：562616342321571143110
	PayCardNo      string         `json:"pay_card_no,omitempty"`      //交易银行卡卡号	String	1024	N	卡号是密文(公私钥加解密)；示例值：
	IsDelayAcct    string         `json:"is_delay_acct,omitempty"`    //是否延时交易	String	1	N	0：实时；1：延时；示例值：0
	IsDiv          string         `json:"is_div,omitempty"`           //是否分账交易	String	1	N	0：非分账交易；1：是分账交易；示例值：0
	FeeFlag        string         `json:"fee_flag,omitempty"`         //手续费扣款标志	String	1	N	1：外扣；2：内扣；示例值：2
	AcctSplitBunch AcctSplitBunch `json:"acct_split_bunch,omitempty"` //分账对象	Object		N	参见分账对象
	SplitFeeInfo   SplitFeeInfo   `json:"split_fee_info,omitempty"`   //分账手续费信息	Object	2048	N	jsonObject格式
	ChannelType    string         `json:"channel_type,omitempty"`     //支付渠道类型	String	1	N	U：银联；示例值：U
	BankOrderNo    string         `json:"bank_order_no,omitempty"`    //支付渠道订单编号	String	32	N	汇付发往银行（银联）的订单编号，即银联流水号； 示例值：202111050936518a4wzkU0A0
	TransStat      string         `json:"trans_stat,omitempty"`       //交易状态	String	1	N	I:处理中，P：处理中，S：成功，F：失败；示例值：S
	MerPriv        string         `json:"mer_priv,omitempty"`         //商户私有域	String	1024	N	原样返回请求参数中的“remark”内容；示例值：备注
	Remark         string         `json:"remark,omitempty"`           //备注	String	1024	N	示例值：
	NotifySyncUrl  string         `json:"notify_sync_url,omitempty"`  //应答跳转的地址	String	4096	N	示例值：https://www.service.com/getresp
	BankCode       string         `json:"bank_code,omitempty"`        //银行返回码	String	32	N	示例值：00
	BankMessage    string         `json:"bank_message,omitempty"`     //银行返回信息	String	32	N	示例值：成功[0000000]
	PayScene       string         `json:"pay_scene,omitempty"`        //支付场景	String	16	N	U_APP：银联云闪付APP；
	AppId          string         `json:"app_id,omitempty"`           //小程序id	String	128	N	示例值：wx8c7d23e805a33666
	UnionAppId     string         `json:"union_app_id,omitempty"`     //云闪付小程序id	String	128	N	示例值：wx3cbe919f36710d1c
	UnionPath      string         `json:"union_path,omitempty"`       //云闪付小程序path	String	1024	N	云闪付小程序path，示例值：/pages/CQPApplet/index?tn=517962138420828734108
	SignTokenNo    string         `json:"sign_token_no,omitempty"`    //签约令牌号	String	32	N	签约令牌号：0029021fe3262200000
}
