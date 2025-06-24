package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V3QuickbuckleApply 快捷绑卡申请接口
//
// 应用场景
// 支持快捷与代扣绑卡申请。上送持卡人四要素信息（银行卡、户名、证件号、银行留存手机号）进行验证，验证成功后银行会向持卡人的银行留存手机号发送短信验证码。
//
// 各银行绑卡短验码有效期不尽相同，最短3分钟。有些银行的短信中会提示有效时间。
// 快捷支持银行参见斗拱支持的快捷银行列表。
// 适用对象
// 开通快捷支付商户
//
// 接口说明
// 请求URL：https://api.huifu.com/v3/quickbuckle/apply
// 请求方式：POST
// 支持格式：JSON
// 加签验签说明：请参考接入指引-开发指南
func (t *Trade) V3QuickbuckleApply(req V3QuickbuckleApplyRequest) (res *V3QuickbuckleApplyResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V3QuickbuckleApplyRequest(BsPaySdk.V3QuickbuckleApplyRequest{
		ReqDate:          req.ReqDate,
		ReqSeqId:         req.ReqSeqId,
		HuifuId:          req.HuifuId,
		OutCustId:        req.OutCustId,
		CardNo:           req.CardNo,
		CardName:         req.CardName,
		CertNo:           req.CertNo,
		CertValidityType: req.CertValidityType,
		CertBeginDate:    req.CertBeginDate,
		CertEndDate:      req.CertEndDate,
		MobileNo:         req.MobileNo,
		ProtocolNo:       req.ProtocolNo,
		ExtendInfos:      common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V3QuickbuckleApplyResponse](resp)
}

type RiskInfo struct {
	DeviceId   string `json:"device_id"`
	DeviceType string `json:"device_type"`
	IpType     string `json:"ip_type"`
	SourceIp   string `json:"source_ip"`
	Mobile     string `json:"mobile"`
}
type TrxDeviceInf struct {
	TrxMobileNum      string `json:"trx_mobile_num"`
	TrxDeviceType     string `json:"trx_device_type"`
	TrxDeviceIp       string `json:"trx_device_ip"`
	TrxDeviceMac      string `json:"trx_device_mac"`
	TrxDeviceImei     string `json:"trx_device_imei"`
	TrxDeviceImsi     string `json:"trx_device_imsi"`
	TrxDeviceIccId    string `json:"trx_device_icc_id"`
	TrxDeviceWfifiMac string `json:"trx_device_wfifi_mac"`
	TrxDeviceGps      string `json:"trx_device_gps"`
}
type V3QuickbuckleApplyExtendInfo struct {
	RiskInfo     RiskInfo     `json:"risk_info"`
	TrxDeviceInf TrxDeviceInf `json:"trx_device_inf"`
}
type V3QuickbuckleApplyRequest struct {
	ReqSeqId         string `json:"req_seq_id"`
	ReqDate          string `json:"req_date"`
	HuifuId          string `json:"huifu_id"`
	OutCustId        string `json:"out_cust_id"`
	CardNo           string `json:"card_no"`
	CardName         string `json:"card_name"`
	CertNo           string `json:"cert_no"`
	MobileNo         string `json:"mobile_no"`
	CertValidityType string `json:"cert_validity_type"`
	CertBeginDate    string `json:"cert_begin_date"`
	CertEndDate      string `json:"cert_end_date"`
	DcType           string `json:"dc_type"`
	Cvv2             string `json:"cvv2"`
	Expiration       string `json:"expiration"`
	ProtocolNo       string `json:"protocol_no"`
	ExtendInfos      V3QuickbuckleApplyExtendInfo
}
type V3QuickbuckleApplyResponse struct {
	Data struct {
		RespDesc string `json:"resp_desc"`
		ReqSeqId string `json:"req_seq_id"`
		ReqDate  string `json:"req_date"`
		RespCode string `json:"resp_code"`
	} `json:"data"`
	Sign string `json:"sign"`
}
