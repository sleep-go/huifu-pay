package merchant

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
	"github.com/sleep-go/huifu-pay/common"
)

// V2MerchantBasicdataQuery 商户详细信息查询
func (bd *Merchant) V2MerchantBasicdataQuery() (res *V2MerchantBasicdataQueryResponse, raw string, err error) {
	resp, err := bd.HuifuPay.BsPay.V2MerchantBasicdataQueryRequest(BsPaySdk.V2MerchantBasicdataQueryRequest{
		ReqSeqId: tool.GetReqSeqId(),
		ReqDate:  tool.GetCurrentDate(),
		HuifuId:  bd.HuifuPay.BsPay.Msc.SysId,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2MerchantBasicdataQueryResponse](resp)
}

type V2MerchantBasicdataQueryResponse struct {
	Data struct {
		AgreementInfoList     common.StringObject[[]AgreementInfoList]   `json:"agreement_info_list"` //协议信息
		AreaId                string                                     `json:"area_id"`             //银行所在市
		BalancePayFlag        string                                     `json:"balance_pay_flag"`    //是否开通余额支付	String	1	N	1:开通 0:未开通，为空未配置(未开通)；示例值：1
		BankBigAmtPayConfig   common.StringObject[BankBigAmtPayConfig]   `json:"bank_big_amt_pay_config"`
		BeneficiaryInfoList   common.StringObject[[]BeneficiaryInfoList] `json:"beneficiary_info_list"`
		BusiType              string                                     `json:"busi_type"` //经营类型	String	1	N	1:实体 2:虚拟，企业商户有值；示例值：1
		ContactEmail          string                                     `json:"contact_email"`
		ContactMobileNo       string                                     `json:"contact_mobile_no"`
		ContactName           string                                     `json:"contact_name"`
		CrossBorderPayConfig  string                                     `json:"cross_border_pay_config"`
		CustType              string                                     `json:"cust_type"`  //商户类型	String	1	Y	1：企业 2：个人 ；示例值：1
		DelayFlag             string                                     `json:"delay_flag"` //是否开通延迟入账	String	1	N	N：否 Y：是，为空则为否；示例值：Y
		DetailAddr            string                                     `json:"detail_addr"`
		DistrictId            string                                     `json:"district_id"`
		EntType               string                                     `json:"ent_type"` //商户种类	String	1	Y	1：政府机构、2：国营企业、3：私营企业、4：外资企业、5：个体工商户、6：其他组织、7：事业单位、9：业主委员会；
		EnterFeeFlag          string                                     `json:"enter_fee_flag"`
		ExtMerId              string                                     `json:"ext_mer_id"`
		FileInfoList          common.StringObject[[]FileInfoList]        `json:"file_info_list"`
		ForcedDelayFlag       string                                     `json:"forced_delay_flag"` //商户开通强制延迟标记	String		N	Y:开通,N:关闭；示例值：N
		FundCollectionFlag    string                                     `json:"fund_collection_flag"`
		HalfPayHostFlag       string                                     `json:"half_pay_host_flag"`
		HeadOfficeFlag        string                                     `json:"head_office_flag"`
		LargeAmtPayConfig     common.StringObject[LargeAmtPayConfig]     `json:"large_amt_pay_config"`
		LegalAddr             string                                     `json:"legal_addr"`
		LegalCertBeginDate    string                                     `json:"legal_cert_begin_date"`
		LegalCertEndDate      string                                     `json:"legal_cert_end_date"`
		LegalCertNo           string                                     `json:"legal_cert_no"`
		LegalCertType         string                                     `json:"legal_cert_type"`
		LegalCertValidityType string                                     `json:"legal_cert_validity_type"`
		LegalName             string                                     `json:"legal_name"`
		LegalNationality      string                                     `json:"legal_nationality"`
		LicenseBeginDate      string                                     `json:"license_begin_date"`
		LicenseCode           string                                     `json:"license_code"` //营业执照编号	String	18	N	工商营业执照编号，示例值：92650109MA79R8E308
		LicenseEndDate        string                                     `json:"license_end_date"`
		LicenseType           string                                     `json:"license_type"`          //证照类型	String	32	N	参见机构证照类型说明； 示例值：NATIONAL_LEGAL
		LicenseValidityType   string                                     `json:"license_validity_type"` //证照有效期类型	String	1	N	企业类商户有值，1:长期有效 0:非长期有效；示例值：0
		LoginName             string                                     `json:"login_name"`
		Mcc                   string                                     `json:"mcc"`         //所属行业（MCC）	String	7	N	银联MCC编码，企业商户有值；示例值：5411
		MerEnName             string                                     `json:"mer_en_name"` //商户英文名称
		MerIcp                string                                     `json:"mer_icp"`
		MerLevel              string                                     `json:"mer_level"`
		MerStat               string                                     `json:"mer_stat"`
		MerUrl                string                                     `json:"mer_url"`
		OnlineBusiType        string                                     `json:"online_busi_type"`
		OnlineFeeConfList     common.StringObject[[]OnlineFeeConfList]   `json:"online_fee_conf_list"`
		OnlineFlag            string                                     `json:"online_flag"`
		OnlineMediaInfo       common.StringObject[[]OnlineMediaInfo]     `json:"online_media_info"`
		OnlineRefund          string                                     `json:"online_refund"`
		OpenLicenceNo         string                                     `json:"open_licence_no"`
		OutFeeFlag            string                                     `json:"out_fee_flag"` //交易手续费外扣标记	String	1	N	1：外扣，2：内扣；（默认2内扣）；示例值：2
		OutOrderFundsMerge    string                                     `json:"out_order_funds_merge"`
		ParentHuifuId         string                                     `json:"parent_huifu_id"`
		PlatformRefund        string                                     `json:"platform_refund"`
		PreAuthorizationFlag  string                                     `json:"pre_authorization_flag"`
		ProductId             string                                     `json:"product_id"`
		ProvId                string                                     `json:"prov_id"` //银行所在省
		QryAliConfList        common.StringObject[[]QryAliConfList]      `json:"qry_ali_conf_list"`
		QryBalancePayConfig   common.StringObject[QryBalancePayConfig]   `json:"qry_balance_pay_config"`
		QryBankCardConf       common.StringObject[QryBankCardConf]       `json:"qry_bank_card_conf"`
		QryCashCardInfoList   common.StringObject[[]QryCashCardInfoList] `json:"qry_cash_card_info_list"`
		QryCashConfigList     common.StringObject[[]QryCashConfigList]   `json:"qry_cash_config_list"`
		QryUnionConf          common.StringObject[QryUnionConf]          `json:"qry_union_conf"`
		QryWxConfList         common.StringObject[[]QryWxConfList]       `json:"qry_wx_conf_list"`
		QuickFlag             string                                     `json:"quick_flag"`
		ReceiptName           string                                     `json:"receipt_name"` //小票名称
		ReconRespAddr         string                                     `json:"recon_resp_addr"`
		RegAreaId             string                                     `json:"reg_area_id"`     //注册市
		RegDetail             string                                     `json:"reg_detail"`      // 册详细地址	String	256	N	与证照保持一致；
		RegDistrictId         string                                     `json:"reg_district_id"` //注册区
		RegName               string                                     `json:"reg_name"`
		RegProvId             string                                     `json:"reg_prov_id"` //注册省	String	6	N	参考地区码；示例值：310000；与证照保持一致
		Remarks               string                                     `json:"remarks"`     //备注
		RespCode              string                                     `json:"resp_code"`
		RespDesc              string                                     `json:"resp_desc"`
		ServicePhone          string                                     `json:"service_phone"`
		ShareHolderInfoList   common.StringObject[[]ShareHolderInfoList] `json:"share_holder_info_list"`
		ShortName             string                                     `json:"short_name"` //商户简称
		SignUserInfoList      common.StringObject[[]SignUserInfoList]    `json:"sign_user_info_list"`
		SmsSendFlag           string                                     `json:"sms_send_flag"`
		SupportRevoke         string                                     `json:"support_revoke"`
		TaxConfig             common.StringObject[TaxConfig]             `json:"tax_config"`
		UniAppPaymentConfig   common.StringObject[UniAppPaymentConfig]   `json:"uni_app_payment_config"`
		UpperHuifuId          string                                     `json:"upper_huifu_id"`
		WalletConfig          string                                     `json:"wallet_config"`
		WebFlag               string                                     `json:"web_flag"`
		WithholdFlag          string                                     `json:"withhold_flag"`
		WxZlConf              string                                     `json:"wx_zl_conf"`
	} `json:"data"`
	Sign string `json:"sign"`
}
type AgreementInfoList struct {
	AgreeBeginDate    string `json:"agree_begin_date"`
	AgreeEndDate      string `json:"agree_end_date"`
	AgreementId       string `json:"agreement_id"`
	AgreementLink     string `json:"agreement_link"`
	AgreementModel    string `json:"agreement_model"`
	AgreementName     string `json:"agreement_name"`
	AgreementNo       string `json:"agreement_no"`
	AgreementQueryUrl string `json:"agreement_query_url,omitempty"`
	AgreementType     string `json:"agreement_type"`
	ConStat           string `json:"con_stat"`
	SignDate          string `json:"sign_date"`
	ArgeementQueryUrl string `json:"argeement_query_url,omitempty"`
}
type QrySettleConfigList struct {
	FeeRate          string `json:"fee_rate"`
	MinAmt           string `json:"min_amt"`
	OutSettleFlag    string `json:"out_settle_flag"`
	OutSettleHuifuid string `json:"out_settle_huifuid"`
	RemainedAmt      string `json:"remained_amt"`
	SettleAbstract   string `json:"settle_abstract"`
	SettleCycle      string `json:"settle_cycle"`
	SettleStatus     string `json:"settle_status"`
	SettleType       string `json:"settle_type"`
}
type QryAliConfList struct {
	AliMerInfos []struct {
		AlipayClsCode   string `json:"alipay_cls_code"`
		BankMerCode     string `json:"bank_mer_code"`
		IsDefault       string `json:"is_default"`
		NuccBankMerCode string `json:"nucc_bank_mer_code"`
		PayChannelId    string `json:"pay_channel_id"`
		SubMerId        string `json:"sub_mer_id"`
	} `json:"ali_mer_infos"`
	FeeChargeType    string   `json:"fee_charge_type"`
	FeeMinAmt        string   `json:"fee_min_amt"`
	FeeRate          string   `json:"fee_rate"`
	PayChannelIdList []string `json:"pay_channel_id_list"`
	PayScene         string   `json:"pay_scene"`
}
type BankBigAmtPayConfig struct {
	BigAmtLimitPerTime      float64 `json:"big_amt_limit_per_time"`
	BizType                 string  `json:"biz_type"`
	FeeFixAmt               float64 `json:"fee_fix_amt"`
	FeeRate                 float64 `json:"fee_rate"`
	MerSameCardRechargeFlag string  `json:"mer_same_card_recharge_flag"`
	OutFeeAcctId            string  `json:"out_fee_acct_id"`
	OutFeeFlag              string  `json:"out_fee_flag"`
	OutFeeHuifuId           string  `json:"out_fee_huifu_id"`
	SwitchState             string  `json:"switch_state"`
}
type BeneficiaryInfoList struct {
	BoAddress            string `json:"bo_address"`
	BoDateStart          string `json:"bo_date_start"`
	BoDeadLine           string `json:"bo_dead_line"`
	BoId                 string `json:"bo_id"`
	BoName               string `json:"bo_name"`
	BoNationality        string `json:"bo_nationality"`
	BoNo                 string `json:"bo_no"`
	BoStat               string `json:"bo_stat"`
	BoType               string `json:"bo_type"`
	FinalBeneficiaryMode string `json:"final_beneficiary_mode"`
}
type FileInfoList struct {
	FileId     string `json:"file_id"`
	FileName   string `json:"file_name"`
	FileType   string `json:"file_type"`
	FileUrl    string `json:"file_url"`
	UpdateTime string `json:"update_time"`
}
type LargeAmtPayConfig struct {
	LargeAmtPayConfigInfoList []struct {
		AllowUserDepositFlag     string  `json:"allow_user_deposit_flag"`
		BizType                  string  `json:"biz_type"`
		BusinessModel            string  `json:"business_model"`
		FeeFixAmt                float64 `json:"fee_fix_amt"`
		FeeRate                  float64 `json:"fee_rate"`
		MerSameCardRechargeFlag  string  `json:"mer_same_card_recharge_flag"`
		ProvisionsAutoRefundFlag string  `json:"provisions_auto_refund_flag"`
		SwitchState              string  `json:"switch_state"`
	} `json:"large_amt_pay_config_info_list"`
	MerPayerCardType string `json:"mer_payer_card_type"`
	OutFeeAcctId     string `json:"out_fee_acct_id"`
	OutFeeFlag       string `json:"out_fee_flag"`
	OutFeeHuifuId    string `json:"out_fee_huifu_id"`
}
type OnlineFeeConfList struct {
	BankId       string  `json:"bank_id"`
	BankName     string  `json:"bank_name"`
	BankShortChn string  `json:"bank_short_chn"`
	DcFlag       string  `json:"dc_flag"`
	FeeRate      string  `json:"fee_rate"`
	StatFlag     string  `json:"stat_flag"`
	FeeType      string  `json:"fee_type,omitempty"`
	FixAmt       string  `json:"fix_amt,omitempty"`
	PayEveryDeal float64 `json:"pay_every_deal,omitempty"`
}
type OnlineMediaInfo struct {
	AuthorizeMaterials string `json:"authorize_materials"`
	MediaFrontPage     string `json:"media_front_page"`
	MediaMerCommonFlag string `json:"media_mer_common_flag"`
	MediaName          string `json:"media_name"`
	MediaOrderPage     string `json:"media_order_page"`
	MediaPayPage       string `json:"media_pay_page"`
	MediaServicePage   string `json:"media_service_page"`
	MediaType          string `json:"media_type"`
	OtherInfo          string `json:"other_info"`
	TestAccount        string `json:"test_account"`
	TestSecret         string `json:"test_secret"`
}
type QryBalancePayConfig struct {
	BalanceModel      string `json:"balance_model"`
	BalancePayExtInfo struct {
		BusiInstruction     string `json:"busi_instruction"`
		CapitalInstruction  string `json:"capital_instruction"`
		FunctionInstruction string `json:"function_instruction"`
	} `json:"balance_pay_ext_info"`
	FeeRate string `json:"fee_rate"`
	FixAmt  string `json:"fix_amt"`
}
type QryCashCardInfoList struct {
	AreaId           string `json:"area_id"`
	BankCode         string `json:"bank_code"`
	BankName         string `json:"bank_name"`
	BranchCode       string `json:"branch_code"`
	BranchName       string `json:"branch_name"`
	CardName         string `json:"card_name"`
	CardNo           string `json:"card_no"`
	CardType         string `json:"card_type"`
	CertBeginDate    string `json:"cert_begin_date"`
	CertEndDate      string `json:"cert_end_date"`
	CertNo           string `json:"cert_no"`
	CertType         string `json:"cert_type"`
	CertValidityType string `json:"cert_validity_type"`
	ProvId           string `json:"prov_id"`
	Status           string `json:"status"`
	TokenNo          string `json:"token_no"`
}
type QryBankCardConf struct {
	IsOpenSmallFlag string `json:"is_open_small_flag"`
}
type QryCashConfigList struct {
	CashType          string `json:"cash_type"`
	FeeRate           string `json:"fee_rate"`
	FixAmt            string `json:"fix_amt"`
	IsPriorityReceipt string `json:"is_priority_receipt,omitempty"`
	OutCashFlag       string `json:"out_cash_flag"`
	SwitchState       string `json:"switch_state"`
}
type QryUnionConf struct {
	BankMerCode       string //银联商户号	String	18	N	银联报备生成的银行商户号；示例值：66345234544
	QrCodeInfo        string //银联小微二维码	String	64	N
	FeeChargeType     string //交易手续费收取类型	String	1	N	1:实收,2:后收；示例值：1
	DebitFeeRateUp    string //借记卡手续费1000以上（%）	String	9	N	数字类型，大于0保留2位小数；取值范围（0.00,100.00]；示例值：0.60
	DebitFeeLimitUp   string //借记卡封顶1000以上	String	9	N	数字类型，大于0保留2位小数；示例值：20.00
	CreditFeeRateUp   string //银联二维码业务贷记卡手续费1000以上（%）	String	9	Y	数字类型，大于0保留2位小数；取值范围（0.00,100.00];示例值：0.60
	DebitFeeRateDown  string //借记卡手续费1000以下（%）	String	9	Y	数字类型，大于0保留2位小数；取值范围（0.00,100.00];示例值：0.60
	DebitFeeLimitDown string //借记卡封顶1000以下	String	4	N	数字类型，大于0保留2位小数；示例值：20.00
	CreditFeeRateDown string //银联二维码业务贷记卡手续费1000以下（%）	String	9	N	数字类型，大于0保留2位小数；取值范围（0.00,100.00]；示例值：0.60
}
type QryWxConfList struct {
	FeeChargeType    string   `json:"fee_charge_type"`
	FeeMinAmt        string   `json:"fee_min_amt"`
	FeeRate          string   `json:"fee_rate"`
	PayChannelIdList []string `json:"pay_channel_id_list"`
	PayScene         string   `json:"pay_scene"`
	WxMerInfos       []struct {
		BankMerCode     string `json:"bank_mer_code"`
		IsDefault       string `json:"is_default"`
		NuccBankMerCode string `json:"nucc_bank_mer_code"`
		PayChannelId    string `json:"pay_channel_id"`
		SubMerId        string `json:"sub_mer_id"`
	} `json:"wx_mer_infos"`
}
type ShareHolderInfoList struct {
	CertBeginDate    string `json:"cert_begin_date"`
	CertEndDate      string `json:"cert_end_date"`
	CertType         string `json:"cert_type"`
	CertValidityType string `json:"cert_validity_type"`
	Name             string `json:"name"`
	ShareHolderId    string `json:"share_holder_id"`
	State            string `json:"state"`
}
type SignUserInfoList struct {
	AuthFileId          string `json:"auth_file_id"`
	CertNo              string `json:"cert_no"`
	IdentityBackFileId  string `json:"identity_back_file_id"`
	IdentityFrontFileId string `json:"identity_front_file_id"`
	MobileNo            string `json:"mobile_no"`
	Name                string `json:"name"`
	Type                string `json:"type"`
}
type TaxConfig struct {
	OpenTaxFlag      string
	OpenTaxState     string
	OpenTaxStateDesc string
}
type UniAppPaymentConfig struct {
	SwitchState                  string //状态开关	String	1	N	0-关闭; 1-开通, 为空默认开通；示例值：1
	DebitFeeRate                 string //借记手续费（%）	String	6	C	借记卡费率与贷记卡费率不能同时为空; 保留2位小数，最大值100.00，最小值0.00；示例值：0.02
	CreditFeeRate                string //贷记手续费（%）	String	6	C	借记卡费率与贷记卡费率不能同时为空; 保留2位小数，最大值100.00，最小值0.00；示例值：0.06
	OutFeeFlag                   string //是否交易手续费外扣	String	1	N	1:外扣 2:内扣（默认2内扣）；示例值：1 如果不为空，则out_fee_acct_type，out_fee_huifuid，out_fee_flag取单独配置，否则取公共配置；
	OutFeeHuifuid                string //交易手续费外扣汇付ID	String	18	N	开通交易手续费外扣业务时必填；示例值：6666000123123123
	OutFeeAcctType               string //交易手续费外扣时的账户类型	String	2	N	01-基本户，05-充值户，09-营销户；不填默认01；示例值：01
	FeeChargeType                string //	交易手续费收取类型	String	1	N	1:实收 2:后收；示例值：1
	CloudQuickPassSecretFreeFlag string //云闪付免密支付开通标识	String	1	N	Y-开通/N-关闭 ；示例值：Y；
}
