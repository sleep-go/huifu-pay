package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeSettlementEncashment 取现
// POST https://api.huifu.com/v2/trade/settlement/encashment
//
// 应用场景
// 用于商户、用户、渠道商从汇付账户取现到事先绑定的取现银行卡中。
//
// 适用对象
// 商户需开通取现功能权限。
func (t *Trade) V2TradeSettlementEncashment(req V2TradeSettlementEncashmentRequest) (res *V2TradeSettlementEncashmentResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeSettlementEncashmentRequest(BsPaySdk.V2TradeSettlementEncashmentRequest{
		ReqDate:          req.ReqDate,
		ReqSeqId:         req.ReqSeqId,
		CashAmt:          req.CashAmt,
		HuifuId:          req.HuifuId,
		IntoAcctDateType: req.IntoAcctDateType,
		TokenNo:          req.TokenNo,
		ExtendInfos:      common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeSettlementEncashmentResponse](resp)
}

type V2TradeSettlementEncashmentExtendInfo struct {
	EnchashmentChannel string `json:"enchashment_channel"` //取现渠道	String	2	N	00：汇付（为空默认）； 10：中信e账通；示例值：00
	Remark             string `json:"remark"`              //备注	String	100	N	示例值：备注
	NotifyUrl          string `json:"notify_url"`          //异步通知地址	String	128	N
	AcctId             string `json:"acct_id"`             //可指定账户号，仅支持基本户、现金户，不填默认为基本户；
}
type V2TradeSettlementEncashmentRequest struct {
	ReqDate          string `json:"req_date" structs:"req_date"`                       // 请求日期
	ReqSeqId         string `json:"req_seq_id" structs:"req_seq_id"`                   // 请求流水号
	CashAmt          string `json:"cash_amt" structs:"cash_amt"`                       // 取现金额
	HuifuId          string `json:"huifu_id" structs:"huifu_id"`                       // 取现方ID号
	IntoAcctDateType string `json:"into_acct_date_type" structs:"into_acct_date_type"` // 到账日期类型
	TokenNo          string `json:"token_no" structs:"token_no"`                       // 取现卡序列号
	ExtendInfos      V2TradeSettlementEncashmentExtendInfo
}
type V2TradeSettlementEncashmentResponse struct {
	Data struct {
		RespCode  string `json:"resp_code,omitempty"`  //业务返回码	String	8	Y	参见业务返回码，示例值：00000000
		RespDesc  string `json:"resp_desc,omitempty"`  //业务返回描述	String	512	Y	业务返回信息，示例值：处理成功
		ReqDate   string `json:"req_date,omitempty"`   //请求日期	String	8	N	格式：yyyyMMdd；示例值：20211123
		ReqSeqId  string `json:"req_seq_id,omitempty"` //请求流水号	String	128	N	示例值：202109160899013231200005
		HfSeqId   string `json:"hf_seq_id,omitempty"`  //汇付全局流水号	String	128	N	示例值：002900TOP3A221112165433P410ac139c1300001
		TransStat string `json:"trans_stat,omitempty"` //交易状态	String	1	Y	S：成功 F：失败 P：处理中；示例值：S
		HuifuId   string `json:"huifu_id,omitempty"`   //商户号/机构号	String	32	N	汇付分配的商户号/机构号，示例值：6666000109812123
		AcctId    string `json:"acct_id,omitempty"`    //账户号	String	32	N	可指定账户号，仅支持基本户、现金户，不填默认为基本户；示例值：F00598600
	} `json:"data,omitempty"`
	Sign string `json:"sign,omitempty"`
}

type V2TradeSettlementEncashmentNotifyMessage struct {
	RespCode string                                                                `json:"resp_code"`
	RespData common.StringObject[V2TradeSettlementEncashmentNotifyMessageRespData] `json:"resp_data"`
	RespDesc string                                                                `json:"resp_desc"`
	Sign     string                                                                `json:"sign"`
}
type V2TradeSettlementEncashmentNotifyMessageRespData struct {
	SubRespCode    string `json:"sub_resp_code,omitempty"`     //业务返回码	String	8	Y	参见业务返回码，示例值：00000000
	SubRespDesc    string `json:"sub_resp_desc,omitempty"`     //业务返回描述	String	512	Y	业务返回信息，示例值：处理成功
	ReqSeqId       string `json:"req_seq_id,omitempty"`        //业务请求流水号	String	128	Y	示例值：202109160899013231200005
	ReqDate        string `json:"req_date,omitempty"`          //业务请求时间	String	8	Y	格式：yyyyMMdd；示例值：20211123
	HfSeqId        string `json:"hf_seq_id,omitempty"`         //汇付全局流水号	String	128	N	示例值：002900TOP3A221112165433P410ac139c1300001
	TransStatus    string `json:"trans_status,omitempty"`      //交易状态	String	1	N	S:成功；F:失败；P:处理中；示例值：S
	AcctStatus     string `json:"acct_status,omitempty"`       //账务状态	String	1	N	S:成功；F:失败；P:处理中；B:回账成功； 示例值：S
	ChannelStatus  string `json:"channel_status,omitempty"`    //通道状态	String	1	N	S:成功；F:失败；P:处理中；示例值：S
	FeeAmt         string `json:"fee_amt,omitempty"`           //手续费	String	14	Y	单位：元。示例值：1.23
	CashAmt        string `json:"cash_amt,omitempty"`          //取现金额	String	14	Y	单位：元。示例值：1.23
	MsgType        string `json:"msg_type,omitempty"`          //消息类型	String	8	N	01:通道；02:账务；示例值：01
	PaymentHfSeqId string `json:"payment_hf_seq_id,omitempty"` //原支付全局流水号	String	128	N	示例值：0030000topB230316111634P778c0a8314500000
	HuifuId        string `json:"huifu_id,omitempty"`          //取现方ID号	String	32	N	示例值：6666000109812123
	AcctId         string `json:"acct_id,omitempty"`           //账户号	String	32	N	示例值：F00598600
	TransType      string `json:"trans_type,omitempty"`        //交易类型	String	20	N	ENCHASHMENT-取现、TSENCHASHMENT-TS结算；示例值：ENCHASHMENT
}
