package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// Feecalc 手续费试算
//
// 应用场景
// 正式交易前通过手续费试算接口可以获取手续费。接口根据试算金额，返回对应交易类型的借贷记手续费，及通用手续费。支持交易类型：聚合正扫（微信、支付宝等）、快捷、网银、手机WAP支付、余额支付、银联APP、取现的手续费试算。
//
// 适用对象
// 斗拱的商户。
//
// 接口说明
// 请求URL：https://api.huifu.com/v2/trade/feecalc
// 请求方式：POST
func (t *Trade) Feecalc(req V2TradeFeecalcRequest) (res *V2TradeFeecalcResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeFeecalcRequest(BsPaySdk.V2TradeFeecalcRequest{
		HuifuId:     req.HuifuId,
		ReqDate:     req.ReqDate,
		ReqSeqId:    req.ReqSeqId,
		TradeType:   string(req.TradeType),
		TransAmt:    req.TransAmt,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeFeecalcResponse](resp)
}

type V2TradeFeecalcRequest struct {
	HuifuId     string           `json:"huifu_id" structs:"huifu_id"`     // 商户号
	ReqDate     string           `json:"req_date" structs:"req_date"`     // 请求日期
	ReqSeqId    string           `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
	TradeType   common.TradeType `json:"trade_type" structs:"trade_type"` // 交易类型
	TransAmt    string           `json:"trans_amt" structs:"trans_amt"`   // 交易金额
	ExtendInfos V2TradeFeecalcExtendInfo
}
type V2TradeFeecalcExtendInfo struct {
	OnlineTransType string `json:"online_trans_type"` //网银交易类型	String	3	N	网银交易必传；3000:B2B网银支付，3001:B2C网银支付，3002:B2B网银充值，3003:B2C网银充值 示例值：3001
	TransAmt        string `json:"trans_amt"`         //交易金额	String	14	Y	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
	BankId          string `json:"bank_id"`           //付款方银行编号	String	8	N	网银、快捷和wap手机支付类交易需传入； 参考： 银行编码；示例值：01020000
	CardType        string `json:"card_type"`         //卡类型	String	8	N	D：借记卡，C：信用卡，默认：D（借记卡）；示例值：C
	ChannelNo       string `json:"channel_no"`        //渠道号	String	32	N	扫码交易必传；如果交易走自有渠道请联系运维人员获取；示例值:10000001
	DigitalBankNo   string `json:"digital_bank_no"`   //数字货币银行编号	String	5	N	01002: 工行，01004: 中行， 01005: 建行，01008：邮储 示例值：01002
	EncashType      string `json:"encash_type"`       //取现到账类型	String	3	N	取现交易必传；T0：当日到账，T1：次工作日到账，D1：次自然日到账； 示例值:T0
	PayScene        string `json:"pay_scene"`         //场景类型	String	2	N
}
type V2TradeFeecalcResponse struct {
	Data struct {
		RespCode     string `json:"resp_code"`
		RespDesc     string `json:"resp_desc"`
		ReqSeqId     string `json:"req_seq_id"`
		ReqDate      string `json:"req_date"`
		HfSeqId      string `json:"hfSeqId"`
		HuifuId      string `json:"huifu_id"`
		FeeAmt       string `json:"fee_amt"`
		CreditFeeAmt string `json:"credit_fee_amt"`
		DebitFeeAmt  string `json:"debit_fee_amt"`
	} `json:"data"`
	Sign string `json:"sign"`
}
