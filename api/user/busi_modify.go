package user

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2UserBusiModify 用户业务入驻修改
// POST https://api.huifu.com/v2/user/busi/modify
//
// 应用场景
// 服务商或商户为旗下的企业/个人用户修改结算费率及结算卡等信息。
//
// 适用对象
// 斗拱服务商或商户；
func (u *User) V2UserBusiModify(req V2UserBusiModifyRequest) (res *V2UserBusiModifyReponse, raw string, err error) {
	resp, err := u.HuifuPay.BsPay.V2UserBusiModifyRequest(BsPaySdk.V2UserBusiModifyRequest{
		HuifuId:      req.HuifuId,
		ReqSeqId:     req.ReqSeqId,
		ReqDate:      req.ReqDate,
		UpperHuifuId: req.UpperHuifuId,
		ExtendInfos:  common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2UserBusiModifyReponse](resp)
}

type ElecReceiptConfig struct {
	SwitchState string `json:"switch_state,omitempty"` //电子回单开关	String	1	Y	0:关闭 1:开通
}
type SignUserInfo struct {
	Type     string //签约人类型	String		Y	LEGAL-法人
	MobileNo string //签约人手机号	String		Y
	CertNo   string //签约人身份证	String		N	企业用户不填默认为法人证件号，个人用户不填默认为个人证件号；当前仅支持法人
	Name     string //签约人姓名	String		N	企业用户不填默认为法人姓名，个人用户不填默认为个人姓名
}
type V2UserBusiModifyRequest struct {
	HuifuId      string `json:"huifu_id"` //汇付ID	String	18	Y	开户时返回的huifu_id；示例值：6666000123123123
	ReqSeqId     string `json:"req_seq_id"`
	ReqDate      string `json:"req_date"`
	UpperHuifuId string `json:"upper_huifu_id"` //渠道商/商户汇付Id	String	18	Y	汇付分配的渠道商或商户编号；示例值：6666000123123123
	ExtendInfos  V2UserBusiModifyExtendInfos
}
type V2UserBusiModifyExtendInfos struct {
	SettleConfig SettleConfig `json:"settle_config"` //结算信息配置	String
	CardInfo     CardInfo     `json:"card_info"`     //结算卡信息
	CashConfig   []CashConfig `json:"cash_config"`   //取现配置列表
	FileList     []struct {
		FileId   string `json:"file_id"`
		FileName string `json:"file_name"`
		FileType string `json:"file_type"`
	} `json:"file_list"` //文件列表
	DelayFlag         string            `json:"delay_flag"`       //延迟入账开关	String	1	N	N：否 Y：是；示例值：Y
	ElecAcctConfig    ElecAcctConfig    `json:"elec_acct_config"` //斗拱e账户功能配置	String		N
	OpenTaxFlag       string            `json:"open_tax_flag"`    //灵活用工开关	String	1	N	N：否（默认） Y：是；示例值：Y 1、个人证件类型必须为身份证类型。2、结算卡信息可不填；若填写则结算类型不能为对公，且结算账户名与个人姓名一致。
	AsyncReturnUrl    string            `json:"async_return_url"` //异步请求地址	String	128	N	为空时不推送异步消息 格式：http://消息接收地址，示例值：http://service.example.com/to/path
	LgPlatformType    string            `json:"lg_platform_type"` //合作平台	String	3	N	LJH-乐接活，HYC-汇优财 灵工业务开关为Y，不填则默认汇优财
	LjhData           common.LjhData    `json:"ljh_data"`         //乐接活配置	String		C	当合作平台为乐接活，必填
	ElecReceiptConfig ElecReceiptConfig `json:"elec_receipt_config"`
	SignUserInfo      SignUserInfo
}
type V2UserBusiModifyReponse struct {
	Data struct {
		RespDesc     string                              `json:"resp_desc"`
		RespCode     string                              `json:"resp_code"`
		HuifuId      string                              `json:"huifu_id"`      //汇付ID	String	18	N	示例值：6666000123123123
		TokenNo      string                              `json:"token_no"`      //取现卡序列号	String	20	N	取现卡序列号，交易时使用；示例值：10000406827
		RespBusiness common.StringObject[[]RespBusiness] `json:"resp_business"` //业务配置结果状态列表	String		N	jsonArray格式
		ApplyNo      string                              `json:"apply_no"`      //申请单号	String	18	N	返回审核中时有值，业务申请单号；示例值：2024022934731647
		LjhResponse  common.StringObject[LjhResponse]    `json:"ljh_response"`  //乐接活配置状态	String		N	灵工场景下，且合作平台为乐接活时，返回该参数。
	} `json:"data"`
	Sign string `json:"sign"`
}
