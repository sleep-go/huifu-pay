package user

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2UserBusiOpen 用户业务入驻
// POST https://api.huifu.com/v2/user/busi/open
//
// 最近更新时间：2025.06.16 作者: 宋少清
//
// 应用场景
// 服务商或商户为旗下的企业/个人用户配置结算费率及结算卡等信息。
//
// 适用对象
// 斗拱服务商或商户；
func (u *User) V2UserBusiOpen(req V2UserBusiOpenRequest) (res *V2UserBusiOpenResponse, raw string, err error) {
	resp, err := u.HuifuPay.BsPay.V2UserBusiOpenRequest(BsPaySdk.V2UserBusiOpenRequest{
		HuifuId:      req.HuifuId,
		ReqSeqId:     req.ReqSeqId,
		ReqDate:      req.ReqDate,
		UpperHuifuId: req.UpperHuifuId,
		ExtendInfos:  common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2UserBusiOpenResponse](resp)
}

type SettleConfig struct {
	SettleStatus       string `json:"settle_status"`        //开通状态	String	1	Y	0：关闭 1：开通；示例值：1
	SettleCycle        string `json:"settle_cycle"`         //结算周期	String	2	Y	T1：下个工作日到账；D1：下个自然日到账；示例值：T1
	MinAmt             string `json:"min_amt"`              //	起结金额	String	14	N	超过该金额后才会结算，单位为元，精确到小数点后两位。 取值范围[0.01,99999999999.99]；示例值：100.00
	RemainedAmt        string `json:"remained_amt"`         //留存金额	String	14	N	小于等于该金额不会结算，单位为元，精确到小数点后两位。 取值范围[0.01,99999999999.99]；示例值：100.00
	SettleAbstract     string `json:"settle_abstract"`      //结算摘要	String	128	N	如果需要自定义结算打款备注，请使用此字段传入，默认为空；支持配置格式化摘要内容，参见结算配置示例说明；
	OutSettleFlag      string `json:"out_settle_flag"`      //手续费外扣标记	String	1	N	1：外扣；2：内扣(为空时默认值)；示例值：1
	OutSettleHuifuid   string `json:"out_settle_huifuid"`   //结算手续费外扣时的汇付ID	String	18	C	外扣手续费承担方的汇付ID。外扣时必填；示例值：6666000123123123
	OutSettleAcctType  string `json:"out_settle_acct_type"` //结算手续费外扣时的账户类型	String	2	C	外扣手续费账户类型； 01：基本户（为空时默认值）， 05：充值户；外扣时必填； 示例值：01
	SettlePattern      string `json:"settle_pattern"`       //结算方式	String	2	N	P0：批次结算（为空时默认值），P2:批次定时结算；示例值：P0
	SettleBatchNo      string `json:"settle_batch_no"`      //结算批次号	String	32	C	settle_pattern为P0时必填；参见结算批次说明
	IsPriorityReceipt  string `json:"is_priority_receipt"`  //是否优先到账	String	1	C	settle_pattern为P0时选填， Y：是 N：否（为空默认取值）；示例值：Y
	SettleTime         string `json:"settle_time"`          //自定义结算处理时间	String	6	C	settle_pattern为P1/P2时必填，注意：00:00到00:30不能指定；格式：HHmmss；示例值：103000
	FixedRatio         string `json:"fixed_ratio"`          //节假日结算手续费率	String	6	C	settle_cycle为D1时必填。单位%，需保留小数点后两位。取值范围[0.00，100.00]，不收费请填写0.00； settle_cycle=T1时，不生效 ；settle_cycle为D1时，遇节假日按此费率结算 ；示例值：0.05
	WorkdayFixedRatio  string `json:"workday_fixed_ratio"`  //工作日结算手续费率	String	6	N	单位%，需保留小数点后两位。取值范围[0.00，100.00]，不填默认为0.00；示例值：0.05
	WorkdayConstantAmt string `json:"workday_constant_amt"` //工作日结算手续费固定金额	String	15	N	单位元，需保留小数点后两位。不填默认为0.00；示例值：1.00
	ConstantAmt        string `json:"constant_amt"`         //节假日结算手续费固定金额	String	15	C	settle_cycle为D1时必填。单位元，需保留小数点后两位。不收费请填写0.00； settle_cycle结算周期为D1时，遇节假日按此费率结算 ； 示例值：1.00
}
type CardInfo struct {
	SettleCycle        string `json:"settle_cycle"`         //结算周期	String	2	Y	T1：下个工作日到账；D1：下个自然日到账；示例值：T1
	MinAmt             string `json:"min_amt"`              //	起结金额	String	14	N	超过该金额后才会结算，单位为元，精确到小数点后两位。 取值范围[0.01,99999999999.99]；示例值：100.00
	RemainedAmt        string `json:"remained_amt"`         //留存金额	String	14	N	小于等于该金额不会结算，单位为元，精确到小数点后两位。 取值范围[0.01,99999999999.99]；示例值：100.00
	SettleAbstract     string `json:"settle_abstract"`      //结算摘要	String	128	N	如果需要自定义结算打款备注，请使用此字段传入，默认为空；支持配置格式化摘要内容，参见结算配置示例说明；
	OutSettleFlag      string `json:"out_settle_flag"`      //手续费外扣标记	String	1	N	1：外扣；2：内扣(为空时默认值)；示例值：1
	OutSettleHuifuid   string `json:"out_settle_huifuid"`   //结算手续费外扣时的汇付ID	String	18	C	外扣手续费承担方的汇付ID。外扣时必填；示例值：6666000123123123
	OutSettleAcctType  string `json:"out_settle_acct_type"` //结算手续费外扣时的账户类型	String	2	C	外扣手续费账户类型； 01：基本户（为空时默认值）， 05：充值户；外扣时必填； 示例值：01
	SettlePattern      string `json:"settle_pattern"`       //结算方式	String	2	N	P0：批次结算（为空时默认值），P2:批次定时结算；示例值：P0
	SettleBatchNo      string `json:"settle_batch_no"`      //结算批次号	String	32	C	settle_pattern为P0时必填；参见结算批次说明
	IsPriorityReceipt  string `json:"is_priority_receipt"`  //是否优先到账	String	1	C	settle_pattern为P0时选填， Y：是 N：否（为空默认取值）；示例值：Y
	SettleTime         string `json:"settle_time"`          //自定义结算处理时间	String	6	C	settle_pattern为P1/P2时必填，注意：00:00到00:30不能指定；格式：HHmmss；示例值：103000
	FixedRatio         string `json:"fixed_ratio"`          //节假日结算手续费率	String	6	C	settle_cycle为D1时必填。单位%，需保留小数点后两位。取值范围[0.00，100.00]，不收费请填写0.00； settle_cycle=T1时，不生效 ；settle_cycle为D1时，遇节假日按此费率结算 ；示例值：0.05
	WorkdayFixedRatio  string `json:"workday_fixed_ratio"`  //工作日结算手续费率	String	6	N	单位%，需保留小数点后两位。取值范围[0.00，100.00]，不填默认为0.00；示例值：0.05
	WorkdayConstantAmt string `json:"workday_constant_amt"` //工作日结算手续费固定金额	String	15	N	单位元，需保留小数点后两位。不填默认为0.00；示例值：1.00
	ConstantAmt        string `json:"constant_amt"`         //节假日结算手续费固定金额	String	15	C	settle_cycle为D1时必填。单位元，需保留小数点后两位。不收费请填写0.00； settle_cycle结算周期为D1时，遇节假日按此费率结算 ； 示例值：1.00
}
type CashConfig struct {
	SwitchState string `json:"switch_state"` //开通状态	String	1	Y	0:关闭 1:开通；示例值：1
	//业务类型	String	2	N	T1:下一工作日到银行账户；
	//D1：下一自然日到银行账户；
	//D0：当日到银行账户；默认D0；
	//DM：当日到账；到账资金不包括当天的交易资金；
	//示例值：T1
	CashType string `json:"cash_type"`
	FixAmt   string `json:"fix_amt"` //提现手续费（固定/元）	String	6	C	fix_amt与fee_rate至少填写一项， 需保留小数点后两位，不收费请填写0.00；示例值：1.00 注：当cash_type=D1时为节假日取现手续费
	//提现手续费率（%）	String	6	C	fix_amt与fee_rate至少填写一项，需保留小数点后两位，取值范围[0.00,100.00]，不收费请填写0.00；示例值：0.05
	//注：1、如果fix_amt与fee_rate都填写了则手续费=fix_amt+支付金额*fee_rate
	//2、当cash_type=D1时为节假日取现手续费
	FeeRate string `json:"fee_rate"`
	//D1工作日取现手续费固定金额	String	6	C	单位元，需保留小数点后两位。不收费请填写0.00；示例值：1.00
	//D1取现配置时选填，其他取现配置无效；cash_type取现类型为D1时，遇工作日按此费率结算，若未配置则默认按照节假日手续费计算
	WeekdayFixAmt string `json:"weekday_fix_amt"`
	//D1工作日取现手续费率	String	6	C	单位%，需保留小数点后两位。取值范围[0.00，100.00]，不收费请填写0.00；示例值：0.05
	//D1取现配置时选填，其他取现配置无效；cash_type取现类型为D1时，遇工作日按此费率结算 ，若未配置则默认按照节假日手续费计算
	WeekdayFeeRate string `json:"weekday_fee_rate"`
	//是否交易手续费外扣	String	1	N	1:外扣 2:内扣（默认2内扣）；示例值：1
	OutFeeFlag string `json:"out_fee_flag"`
	//手续费承担方	String	18	N	手续费外扣时必需指定手续费承担方ID；示例值：6666000123123123
	OutFeeHuifuId string `json:"out_fee_huifu_id"`
	//交易手续费外扣的账户类型	String	2	N	01-基本户，05-充值户；不填默认01； 示例值：01
	OutFeeAcctType    string `json:"out_fee_acct_type"`
	IsPriorityReceipt string `json:"is_priority_receipt"` //是否优先到账	String	1	N	Y：是 ，N：否。不填，默认值为否。仅在取现类型配置为D1 和 T1 时生效。示例值：Y
}
type LjhData struct {
	TaxAreaId string //税源地id	String	13	C	当合作平台为乐接活，必填
}
type ElecAcctConfig struct {
	SwitchState       string `json:"switch_state"`          //电子账户开关	String	1	Y	电子账户开通总开关：0:关闭 1:开通
	AcctType          string `json:"acct_type"`             //账户类型	String	2	Y	01：中信e管家
	CashFeeParty      string `json:"cash_fee_party"`        //电子账户提现手续费承担方	String	1	Y	1:总部 2:其他；
	Scene             string `json:"scene"`                 //场景	String	3	Y	与角色类型关联，中信定义的资金类型；示例值：001
	RoleType          string `json:"role_type"`             //角色类型(角色编号)	String	6	Y	必填；与场景关联，中信定义的角色类型;示例值：001001
	ElecCardList      string `json:"elec_card_list"`        //银行卡信息	Object		N	jsonArray字符串，如果开通斗拱E账户但不提供绑卡信息将无法取现，后续绑卡请调用电子账户绑卡接口
	UserType          string `json:"user_type"`             //用户类型	String	18	N	SPLIT-分账用户，默认； SETTLE-结算用户，不支持分账、余额支付；
	ElecAcctSignSeqId string `json:"elec_acct_sign_seq_id"` //中信签约短信流水号	String	64	N	示例值：EMSSBPG2504284593690058431260676
	SignSuccessFlag   string `json:"sign_success_flag"`     //签约成功标志	String	1	Y	Y：成功
}
type V2UserBusiOpenRequest struct {
	ReqSeqId     string `json:"req_seq_id"`
	ReqDate      string `json:"req_date"`
	HuifuId      string `json:"huifu_id"`       //汇付ID	String	18	Y	开户时返回的huifu_id；示例值：6666000123123123
	UpperHuifuId string `json:"upper_huifu_id"` //渠道商/商户汇付Id	String	18	Y	汇付分配的渠道商或商户编号；示例值：6666000123123123
	ExtendInfos  struct {
		SettleConfig SettleConfig `json:"settle_config"` //结算信息配置	String
		CardInfo     CardInfo     `json:"card_info"`     //结算卡信息
		CashConfig   []CashConfig `json:"cash_config"`   //取现配置列表
		FileList     []struct {
			FileId   string `json:"file_id"`
			FileName string `json:"file_name"`
			FileType string `json:"file_type"`
		} `json:"file_list"` //文件列表
		DelayFlag      string         `json:"delay_flag"`       //延迟入账开关	String	1	N	N：否 Y：是；示例值：Y
		ElecAcctConfig ElecAcctConfig `json:"elec_acct_config"` //斗拱e账户功能配置	String		N
		OpenTaxFlag    string         `json:"open_tax_flag"`    //灵活用工开关	String	1	N	N：否（默认） Y：是；示例值：Y 1、个人证件类型必须为身份证类型。2、结算卡信息可不填；若填写则结算类型不能为对公，且结算账户名与个人姓名一致。
		AsyncReturnUrl string         `json:"async_return_url"` //异步请求地址	String	128	N	为空时不推送异步消息 格式：http://消息接收地址，示例值：http://service.example.com/to/path
		LgPlatformType string         `json:"lg_platform_type"` //合作平台	String	3	N	LJH-乐接活，HYC-汇优财 灵工业务开关为Y，不填则默认汇优财
		LjhData        LjhData        `json:"ljh_data"`         //乐接活配置	String		C	当合作平台为乐接活，必填
	}
}
type V2UserBusiOpenResponse struct {
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
type RespBusiness struct {
	Type string `json:"type"` //配置类型	String	1	Y	1、绑卡信息；2、取现配置；3、结算信息配置；5、灵工业务配置；示例值：1
	Code string `json:"code"` //配置状态	String	1	Y	S:成功，F:失败；示例值：S
	Msg  string `json:"msg"`  //配置返回信息	String	512	N	业务响应信息；示例值：
}
type LjhResponse struct {
	AcId       string `json:"ac_id"`       //商事id
	Status     string `json:"status"`      //商事主体状态	String	1	N	0:待注册 1:注册中 2:注册成功 3:注册失败 4:园区处理中 5:已注销 6:注销失败
	StatusDesc string `json:"status_desc"` //注册状态描述
	UserId     string `json:"user_id"`     //乐接活用户id
}
