package user

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/api/merchant"
	"github.com/sleep-go/huifu-pay/common"
)

// V2UserBasicdataQuery 用户信息查询
// POST https://api.huifu.com/v2/user/basicdata/query
//
// 应用场景
// 查询为用户基本信息、结算卡及配置费率等信息。
//
// 适用对象 斗拱服务商或商户；
func (u *User) V2UserBasicdataQuery(req V2UserBasicdataQueryRequest) (res *V2UserBasicdataQueryResponse, raw string, err error) {
	resp, err := u.HuifuPay.BsPay.V2UserBasicdataQueryRequest(BsPaySdk.V2UserBasicdataQueryRequest{
		HuifuId:  req.HuifuId,
		ReqSeqId: req.ReqSeqId,
		ReqDate:  req.ReqDate,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2UserBasicdataQueryResponse](resp)
}

type V2UserBasicdataQueryRequest struct {
	ReqSeqId string `json:"req_seq_id"`
	ReqDate  string `json:"req_date"`
	HuifuId  string `json:"huifu_id"`
}
type IndvBaseInfo struct {
	CertBeginDate    string `json:"cert_begin_date"`
	CertEndDate      string `json:"cert_end_date"`
	CertNo           string `json:"cert_no"`
	CertType         string `json:"cert_type"`
	CertValidityType string `json:"cert_validity_type"`
	Email            string `json:"email"`
	FileList         []struct {
		FileId   string `json:"file_id"`
		FileName string `json:"file_name"`
		FileType string `json:"file_type"`
	} `json:"file_list"`
	LoginName string `json:"login_name"`
	MobileNo  string `json:"mobile_no"`
	Name      string `json:"name"`
}
type V2UserBasicdataQueryResponse struct {
	Data struct {
		RespCode            string                                            `json:"resp_code"`
		RespDesc            string                                            `json:"resp_desc"`
		EntBaseInfo         string                                            `json:"ent_base_info"`
		IndvBaseInfo        common.StringObject[IndvBaseInfo]                 `json:"indv_base_info"`
		CardInfo            common.StringObject[CardInfo]                     `json:"card_info"`
		SettleConfig        common.StringObject[SettleConfig]                 `json:"settle_config"`
		QryCashConfigList   common.StringObject[[]merchant.QryCashConfigList] `json:"qry_cash_config_list"`
		QryCashCardInfoList string                                            `json:"qry_cash_card_info_list"`
		ElecAcctConfig      common.StringObject[ElecAcctConfig]               `json:"elec_acct_config"`
		ElecReceiptConfig   common.StringObject[ElecReceiptConfig]            `json:"elec_receipt_config"`
		SignUserInfo        common.StringObject[SignUserInfo]                 `json:"sign_user_info"`
	} `json:"data"`
	Sign string `json:"sign"`
}
