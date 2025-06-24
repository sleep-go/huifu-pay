package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2QuickbuckleOnekeyCardbind 一键绑卡
//
// 应用场景
// 用户只要输入个人身份信息（姓名、身份证号、手机号）选择好银行，免输卡号也可以绑卡。斗拱系统会将个人身份信息送到银行，由银行系统负责绑卡。绑卡之后可以通过快捷方式进行支付。
//
// 适用对象
// 开通快捷支付的商户。
//
// 接口说明
// 请求URL：https://api.huifu.com/v2/quickbuckle/onekey/cardbind
// 请求方式：POST
// 支持格式：JSON
func (t *Trade) V2QuickbuckleOnekeyCardbind(req V2QuickbuckleOnekeyCardbindRequest) (res *V2QuickbuckleOnekeyCardbindResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2QuickbuckleOnekeyCardbindRequest(BsPaySdk.V2QuickbuckleOnekeyCardbindRequest{
		ReqSeqId:       req.ReqSeqId,
		ReqDate:        req.ReqDate,
		HuifuId:        req.HuifuId,
		OutCustId:      req.OutCustId,
		BankId:         req.BankId,
		CardName:       req.CardName,
		CertId:         req.CertId,
		CertType:       req.CertType,
		CertEndDate:    req.CertEndDate,
		CardMp:         req.CardMp,
		DcType:         req.DcType,
		AsyncReturnUrl: req.AsyncReturnUrl,
		ExtendInfos:    common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2QuickbuckleOnekeyCardbindResponse](resp)
}

type V2QuickbuckleOnekeyCardbindExtendInfo struct {
	CertValidityType string       `json:"cert_validity_type"`
	CertBeginDate    string       `json:"cert_begin_date"`
	ProtocolNo       string       `json:"protocol_no"`
	RiskInfo         RiskInfo     `json:"risk_info"`
	TrxDeviceInf     TrxDeviceInf `json:"trx_device_inf"`
	VerifyFrontUrl   string       `json:"verify_front_url"`
}
type V2QuickbuckleOnekeyCardbindRequest struct {
	ReqSeqId       string `json:"req_seq_id"`
	ReqDate        string `json:"req_date"`
	HuifuId        string `json:"huifu_id"`
	OutCustId      string `json:"out_cust_id"`
	BankId         string `json:"bank_id"`
	CardName       string `json:"card_name"`
	CertType       string `json:"cert_type"`
	CertEndDate    string `json:"cert_end_date"`
	CertId         string `json:"cert_id"`
	DcType         string `json:"dc_type"`
	CardMp         string `json:"card_mp"`
	AsyncReturnUrl string `json:"async_return_url"`
	ExtendInfos    V2QuickbuckleOnekeyCardbindExtendInfo
}
type V2QuickbuckleOnekeyCardbindResponse struct {
	Data struct {
		RespCode     string `json:"resp_code"`
		RespDesc     string `json:"resp_desc"`
		HuifuId      string `json:"huifu_id"`
		ReqDate      string `json:"req_date"`
		ReqSeqId     string `json:"req_seq_id"`
		TokenNo      string `json:"token_no"`
		TransStatus  string `json:"trans_status"`
		BankCode     string `json:"bank_code"`
		BankDesc     string `json:"bank_desc"`
		VerifyDetail string `json:"verify_detail"`
		CardIdMask   string `json:"card_id_mask"`
	} `json:"data"`
	Sign string `json:"sign"`
}
type V2QuickbuckleOnekeyCardbindNotifyMessageData struct {
	V2QuickbuckleOnekeyCardbindResponse
}
type V2QuickbuckleOnekeyCardbindNotifyMessage struct {
	RespCode string                                                            `json:"resp_code"`
	RespDesc string                                                            `json:"resp_desc"`
	Data     common.StringObject[V2QuickbuckleOnekeyCardbindNotifyMessageData] `json:"data"`
	Sign     string
}
