package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V3TradePaymentScanpayQuery 扫码交易查询
// POST https://api.huifu.com/v3/trade/payment/scanpay/query
//
// 应用场景
// 服务商/商户系统因网络原因未收到交易状态，可以通过本接口主动查询订单状态。支持微信公众号-T_JSAPI、小程序-T_MINIAPP、微信APP支付-T_APP、支付宝JS-A_JSAPI、支付宝正扫-A_NATIVE、银联二维码正扫-U_NATIVE、银联二维码JS-U_JSAP、数字货币二维码支付-D_NATIVE，以及微信反扫、支付宝反扫、银联二维码反扫、数字人民币反扫交易查询。
//
// 适用对象
// 开通微信/支付宝/银联二维码/数字人民币聚合扫码功能的商户。
func (t *Trade) V3TradePaymentScanpayQuery(req V3TradePaymentScanpayQueryRequest) (res *V3TradePaymentScanpayQueryResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V3TradePaymentScanpayQueryRequest(BsPaySdk.V3TradePaymentScanpayQueryRequest{
		HuifuId:     req.HuifuId,
		OrgReqDate:  req.OrgReqDate,
		OutOrdId:    req.OutOrdId,
		OrgHfSeqId:  req.OrgHfSeqId,
		OrgReqSeqId: req.OrgReqSeqId,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V3TradePaymentScanpayQueryResponse](resp)
}

type V3TradePaymentScanpayQueryExtendInfos struct {
}
type V3TradePaymentScanpayQueryRequest struct {
	HuifuId     string `json:"huifu_id"`
	OrgReqDate  string `json:"org_req_date"`
	OutOrdId    string `json:"out_ord_id"`
	OrgHfSeqId  string `json:"org_hf_seq_id"`
	OrgReqSeqId string `json:"org_req_seq_id"`
	ExtendInfos V3TradePaymentScanpayQueryExtendInfos
}
type V3TradePaymentScanpayQueryResponse struct {
	Data struct {
		RespCode                  string                                                               `json:"resp_code"`
		RespDesc                  string                                                               `json:"resp_desc"`
		BagentId                  string                                                               `json:"bagent_id"`
		HuifuId                   string                                                               `json:"huifu_id"`
		OrgReqDate                string                                                               `json:"org_req_date"`
		OrgHfSeqId                string                                                               `json:"org_hf_seq_id"`
		OrgReqSeqId               string                                                               `json:"org_req_seq_id"`
		OutTransId                string                                                               `json:"out_trans_id"`
		PartyOrderId              string                                                               `json:"party_order_id"`
		TransAmt                  string                                                               `json:"trans_amt"`
		PayAmt                    string                                                               `json:"pay_amt"`
		SettlementAmt             string                                                               `json:"settlement_amt"` //结算金额	String	14	N	单位元；示例值：1000.00
		UnconfirmAmt              string                                                               `json:"unconfirm_amt"`  //待确认总金额	String	14	N	单位元，需保留小数点后两位，示例值：1.00
		ConfirmedAmt              string                                                               `json:"confirmed_amt"`  //已确认总金额	String	14	N	单位元，需保留小数点后两位，示例值：1.00
		TransType                 string                                                               `json:"trans_type"`
		TransStat                 string                                                               `json:"trans_stat"`
		TransTime                 string                                                               `json:"trans_time"`
		EndTime                   string                                                               `json:"end_time"`
		DelayAcctFlag             string                                                               `json:"delay_acct_flag"`
		AcctId                    string                                                               `json:"acct_id"`
		AcctDate                  string                                                               `json:"acct_date"`
		AcctStat                  string                                                               `json:"acct_stat"`
		DebitType                 string                                                               `json:"debit_type"`
		FeeHuifuId                string                                                               `json:"fee_huifu_id"`
		FeeFormulaInfos           common.StringObject[FeeFormulaInfos]                                 `json:"fee_formula_infos"`
		FeeAmt                    string                                                               `json:"fee_amt"`
		FeeType                   string                                                               `json:"fee_type"`
		WxUserId                  string                                                               `json:"wx_user_id"`
		WxResponse                common.StringObject[WxResponse]                                      `json:"wx_response"`
		AlipayResponse            common.StringObject[AlipayResponse]                                  `json:"alipay_response"`
		UnionpayResponse          string                                                               `json:"unionpay_response"`
		DivFlag                   string                                                               `json:"div_flag"`
		AcctSplitBunch            common.StringObject[V3TradePaymentScanpayQueryAcctSplitBunch]        `json:"acct_split_bunch"`
		SplitFeeInfo              common.StringObject[SplitFeeInfo]                                    `json:"split_fee_info"`
		CombinedpayData           string                                                               `json:"combinedpay_data"`
		CombinedpayFeeAmt         string                                                               `json:"combinedpay_fee_amt"`
		TransFeeAllowanceInfo     common.StringObject[V3TradePaymentScanpayQueryTransFeeAllowanceInfo] `json:"trans_fee_allowance_info"`
		Remark                    string                                                               `json:"remark"`
		DeviceType                string                                                               `json:"device_type"`
		MerDevLocation            common.StringObject[MerDevLocation]                                  `json:"mer_dev_location"`
		MerPriv                   string                                                               `json:"mer_priv"`
		AuthNo                    string                                                               `json:"auth_no"`                        //授权号	String	6	N	同一商户当天，同一终端，同一批次号唯一；示例值：727902
		PasswordTrade             string                                                               `json:"password_trade"`                 //输入密码提示	String	1	N	Y-等待用户输入密码状态；示例值：Y
		MerName                   string                                                               `json:"mer_name"`                       //商户名称	String	100	N	示例值：上海汇付支付服务公司
		ShopName                  string                                                               `json:"shop_name"`                      //店铺名称	String	100	N	示例值：上海汇付支付宝山分公司
		PreAuthAmt                string                                                               `json:"pre_auth_amt"`                   //预授权金额	String	14	N	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
		PreAuthPayAmt             string                                                               `json:"pre_auth_pay_amt"`               //预授权完成金额	String	14	N	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
		OrgAuthNo                 string                                                               `json:"org_auth_no"`                    //	原授权号	String	6	N	同一商户当天，同一终端，同一批次号唯一；示例值：727902
		PreAuthHfSeqId            string                                                               `json:"pre_auth_hf_seq_id"`             //预授权汇付全局流水号	String	128	N	示例值：0035000topA220628152651P306c0a8217a00000
		PreAuthPayFeeAmount       string                                                               `json:"pre_auth_pay_fee_amount"`        //预授权完成手续费	String	14	N	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
		PreAuthPayRefundFeeAmount string                                                               `json:"pre_auth_pay_refund_fee_amount"` //预授权完成退还手续费	String	14	N	单位元，需保留小数点后两位，示例值：1.00，最低传入0.01
		OrgFeeFlag                string                                                               `json:"org_fee_flag"`                   //原手续费扣款标志	String	8	N	INNER：内扣， OUTSIDE：外扣；示例值：OUTSIDE
		OrgFeeRecType             string                                                               `json:"org_fee_rec_type"`               //原手续费扣取方式	String	1	N	1:实收,2:后收；示例值：1
		OrgAllowanceType          string                                                               `json:"org_allowance_type"`             //原补贴类型	String	1	N	0-不补贴，1-补贴,2-部分补贴；示例值：1
		FqChannels                string                                                               `json:"fq_channels"`                    //信用卡分期资产方式	String	20	N	代表优先使用资产类型；alipayfq_cc：表示信⽤卡分期。示例值：alipayfq_cc
		BankDesc                  string                                                               `json:"bank_desc"`                      //外部通道返回描述	String	200	N	示例值：TRADE_SUCCESS
		AtuSubMerId               string                                                               `json:"atu_sub_mer_id"`                 //ATU真实商户号	String	32	N	示例值：411111141
		DcResponse                common.StringObject[DcResponse]                                      `json:"dc_response"`                    //数字货币返回的响应报文	String	6000	N	Json to String
		FundFreezeStat            string                                                               `json:"fund_freeze_stat"`               //资金冻结状态	String	16	N	FREEZE：冻结；UNFREEZE：解冻；示例值：UNFREEZE
		FreezeTime                string                                                               `json:"freeze_time"`                    //冻结时间	String	14	Y	格式为yyyyMMddHHMMSS，示例值：20091225091010
		UnfreezeAmt               string                                                               `json:"unfreeze_amt"`                   //解冻金额	String	14	Y	单元：元。示例值：1.23
		UnfreezeTime              string                                                               `json:"unfreeze_time"`                  //解冻时间	String	14	Y	格式为yyyyMMddHHMMSS，示例值：20091225091010
	} `json:"data"`
	Sign string `json:"sign"`
}

// FeeFormulaInfos 手续费费率信息
type FeeFormulaInfos struct {
	FeeFormula string `json:"fee_formula,omitempty"` //手续费计算公式	String	512	Y	示例值：AMT*0.003
	FeeType    string `json:"fee_type,omitempty"`    //手续费类型	String	32	Y	TRANS_FEE：交易手续费；ACCT_FEE：组合支付账户补贴手续费 示例值：ACCT_FEE
	HuifuId    string `json:"huifu_id,omitempty"`    //商户号	String	32	N	组合支付账户补贴时，补贴账户的huifuId；示例值：6666000123123123
}
type WxResponse struct {
	SubOpenid       string `json:"sub_openid,omitempty"` //子商户公众账号ID	String	16	Y	用户在子商户appid下的唯一标识；示例值：oWNHX5RNaCUmZR
	Openid          string `json:"openid,omitempty"`     //用户标识	String	128	Y	用户在商户appid下的唯一标识；示例值：oGhiSxIAPtEnPfe9Xo000000B
	CashFee         string `json:"cash_fee,omitempty"`   //现金支付金额	String	14	Y	现金支付金额订单现金支付金额；单位格式为0.00；示例值：1.00
	Attach          string `json:"attach,omitempty"`     //商家数据包	String	128	N	原样返回；示例值：附加数据
	CouponFee       string `json:"coupon_fee,omitempty"` //代金券金额	String	14	N	代金券或立减优惠金额<=订单总金额， 订单总金额-代金券或立减优惠金额=现金支付金额； 示例值：1.00
	PromotionDetail []struct {
		ActivityId  string `json:"activity_id,omitempty"`  //活动id	String	32	Y	在微信商户后台配置的批次ID；示例值：458626534790131712
		Amount      string `json:"amount,omitempty"`       //优惠券面额	String	5	Y	用户享受优惠的金额；示例值：10.00
		PromotionId string `json:"promotion_id,omitempty"` //券或者立减优惠id	String	32	Y	示例值：2345234235
		GoodsDetail []struct {
			DiscountAmount string `json:"discount_amount,omitempty"` //商品优惠金额	String	32	Y	单品的总优惠金额，单位为：元；示例值：20.00
			GoodsId        string `json:"goods_id,omitempty"`        //商品编码	String	32	Y	由半角的大小写字母、数字、中划线、下划线中的一种或几种组成；示例值：6934572310301
			Price          string `json:"price,omitempty"`           //商品价格	String	32	Y	单位为：元。示例值：50.00 如果商户有优惠，需传输商户优惠后的单价。 例如：用户对一笔100元的订单使用了商场发的纸质优惠券100-50，则活动商品的单价应为原单价-50
			Quantity       string `json:"quantity,omitempty"`        //商品数量	String	32	Y	用户购买的数量；示例值：10
			GoodsRemark    string `json:"goods_remark,omitempty"`    //商品备注	String	32	N	按照配置原样返回，字段内容在微信后台配置券时进行设置。示例值：商品备注
		} `json:"goods_detail,omitempty"` //单品列表	Array	3000	N	单品信息，Json格式
		MerchantContribute string `json:"merchant_contribute,omitempty"` //商户出资(元)	String	32	N	特指商户自己创建的优惠，出资金额等于本项优惠总金额；示例值：1000
		Name               string `json:"name,omitempty"`                //优惠名称	String	64	N	示例值：八折券
		OtherContribute    string `json:"other_contribute,omitempty"`    //其他出资(元)	String	32	N	其他出资方出资金额；示例值：2000
		Scope              string `json:"scope,omitempty"`               //优惠范围	String	32	N	GLOBAL- 全场代金券，SINGLE- 单品优惠；示例值：GLOBAL
		Type               string `json:"type,omitempty"`                //优惠类型	String	32	N	COUPON- 代金券，需要走结算资金的充值型代金券,（境外商户券币种与支付币种一致） DISCOUNT- 优惠券，不走结算资金的免充值型优惠券，（境外商户券币种与标价币种一致） 示例值：DISCOUNT
	} `json:"promotion_detail,omitempty"` //营销详情列表	Array	6000	N	营销详情列表，使返回值为Json格式；
	BankType   string `json:"bank_type,omitempty"`   //付款银行	String	16	Y	银行类型，采用字符串类型的银行标识；参见微信银行类型说明；示例值：CMB
	SubAppid   string `json:"sub_appid,omitempty"`   //商户公众号APPID	String	32	N	直连交易字段；示例值：wx5934540532
	DeviceInfo string `json:"device_info,omitempty"` //交易终端设备信息	String	32	N	直连交易字段；示例值：
}
type AlipayResponse struct {
	BuyerId      string `json:"buyer_id,omitempty"`
	BuyerLogonId string `json:"buyer_logon_id,omitempty"`
	FundBillList []struct {
		BankCode    string `json:"bank_code" json:"bank_code,omitempty"`
		Amount      string `json:"amount" json:"amount,omitempty"`
		FundChannel string `json:"fund_channel" json:"fund_channel,omitempty"`
		FundType    string `json:"fund_type" json:"fund_type,omitempty"`
		RealAmount  string `json:"real_amount" json:"real_amount,omitempty"`
	} `json:"fund_bill_list,omitempty"`
	HbFqNum           string `json:"hb_fq_num,omitempty"`
	VoucherDetailList []struct {
		Amount             string `json:"amount,omitempty" json:"amount,omitempty"`                           //优惠券面额（元）	String	32	Y	用户享受优惠的金额；示例值：10.00
		Id                 string `json:"id,omitempty" json:"id,omitempty"`                                   //券id	String	32	Y	优惠券号；示例值：6934572310301
		Name               string `json:"name,omitempty" json:"name,omitempty"`                               //券名称	String	32	Y	优惠名称；示例值：实体店付款通用立减券
		Type               string `json:"type,omitempty" json:"type,omitempty"`                               //券类型	String	32	Y	COUPON- 代金券，需要走结算资金的充值型代金券,（境外商户券币种与支付币种一致） DISCOUNT- 优惠券，不走结算资金的免充值型优惠券，（境外商户券币种与标价币种一致）； 示例值：DISCOUNT
		MerchantContribute string `json:"merchant_contribute,omitempty" json:"merchant_contribute,omitempty"` //商家出资	String	32	N	特指商户自己创建的优惠，出资金额等于本项优惠总金额，单位为元。示例值：0.10
		OtherContribute    string `json:"other_contribute,omitempty" json:"other_contribute,omitempty"`       //其他出资方出资金额	String	14	N	其他出资方出资金额，单位为元。示例值：0.20
	} `json:"voucher_detail_list,omitempty"`
	FqOrdAmt        string `json:"fq_ord_amt,omitempty"`       //花呗组合支付分期订单金额	String	14	N	单位元 格式:0.00；示例值：1.00
	FqAcqFeeAmt     string `json:"fq_acq_fee_amt,omitempty"`   //花呗组合支付分期收单手续费	String	14	N	单位元 格式:0.00；示例值：1.00
	OthOrdAmt       string `json:"oth_ord_amt,omitempty"`      //花呗组合支付非分期订单金额	String	14	N	单位元 格式:0.00；示例值：1.00
	OthFeeAmt       string `json:"oth_fee_amt,omitempty"`      //花呗组合支付非分期手续费金额	String	14	N	单位元 格式:0.00；示例值：1.00
	BuyerPayAmount  string `json:"buyer_pay_amount,omitempty"` //买家实付金额	String	11	N	直连模式字段，单位为元，两位小数。示例值：1.00 该金额代表该笔交易买家实际支付的金额，不包含商户折扣等金额
	PointAmount     string `json:"point_amount,omitempty"`     //积分支付的金额	String	11	N	直连模式字段，单位为元，两位小数。示例值：1.00 该金额代表该笔交易中用户。使用积分支付的金额，比如集分宝或者支付宝实时优惠等
	InvoiceAmount   string `json:"invoice_amount,omitempty"`   //交易中用户支付的可开具发票的金额	String	11	N	直连模式字段，单位为元，两位小数。示例值：1.00 该金额代表该笔交易中可以给用户开具发票的金额。
	SendPayDate     string `json:"send_pay_date,omitempty"`    //本次交易打款给卖家的时间	String	32	N	直连模式字段；示例值：2014-11-27 15:45:57
	StoreId         string `json:"store_id,omitempty"`         //商户门店编号	String	32	N	直连模式字段；示例值：sh001
	TerminalId      string `json:"terminal_id,omitempty"`      //商户机具终端编号	String	32	N	直连模式字段；示例值：898600680113F0123014
	StoreName       string `json:"store_name,omitempty"`       //请求交易支付中的商户店铺的名称	String	512	N	直连模式字段；示例值：汇付宝山路门店
	BuyerUserType   string `json:"buyer_user_type,omitempty"`  //买家用户类型	String	18	N	直连模式字段。CORPORATE:企业用户；PRIVATE:个人用户。；示例值：PRIVATE
	MdiscountAmount string `json:"mdiscount_amount,omitempty"` //商家优惠金额	String	11	N	直连模式字段；示例值：2.00
	DiscountAmount  string `json:"discount_amount,omitempty"`  //平台优惠金额	String	11	N	直连模式字段；示例值：3.00
	ExtInfos        string `json:"ext_infos,omitempty"`        //交易额外信息	Json	1024	N	直连模式字段，特殊场景下与支付宝约定返回；示例值：交易额外信息
}
type V3TradePaymentScanpayQueryAcctSplitBunch struct {
	AcctInfos []struct {
		AcctId  string `json:"acct_id"`
		DivAmt  string `json:"div_amt"`
		HuifuId string `json:"huifu_id"`
	} `json:"acct_infos"`
	IsCleanSplit string
}
type SplitFeeInfo struct {
	TotalSplitFeeAmt string `json:"total_split_fee_amt,omitempty"` //分账手续费总金额(元)	String	14	N	示例值：0.10
	SplitFeeFlag     string `json:"split_fee_flag,omitempty"`      //分账手续费扣款标志	Integer	1	Y	1: 外扣 2: 内扣；示例值：2
	SplitFeeDetails  []struct {
		SplitFeeAmt     string `json:"split_fee_amt,omitempty"`      //分账手续费金额(元)	String	14	Y	示例值：0.10
		SplitFeeHuifuId string `json:"split_fee_huifu_id,omitempty"` //分账手续费承担方商户号	String	32	Y	示例值：6666000123120001
		SplitFeeAcctId  string `json:"split_fee_acct_id,omitempty"`  //分账手续费承担方账号	String	32	N	示例值：F00598600
	} `json:"split_fee_details,omitempty"` //分账手续费明细	Array		Y	分账手续费明细
}
type V3TradePaymentScanpayQueryTransFeeAllowanceInfo struct {
	ActualFeeAmt            string `json:"actual_fee_amt"`
	AllowanceFeeAmt         string `json:"allowance_fee_amt"`
	AllowanceType           string `json:"allowance_type"`
	CurAllowanceConfigInfos struct {
		MerchantGroup     string //商户号	String	64	N	示例值：6666000123120000
		AcctId            string //门店	String	64	N	示例值：sh002
		ActivityId        string //活动号	String	64	Y	示例值：223402342
		ActivityName      string //活动描述	String	128	N	示例值：开业大促
		Status            string //活动是否有效	String	4	Y	1:生效 0：失效；示例值：1
		TotalLimitAmt     string //活动总补贴额度	String	16	Y	示例值：10.00
		StartTime         string //活动开始时间	String	8	Y	yyyyMMdd；示例值：20220909
		EndTime           string //活动结束时间	String	8	Y	yyyyMMdd；示例值：202209011
		HumanFlag         string //是否人工操作	String	4	Y	N：自动 Y：人工；示例值：N
		AllowanceSys      string //补贴方	String	64	Y	1:银行 2:服务商 3:汇来米；示例值：1
		AllowanceSysId    string //补贴方ID	String	64	Y	对应补贴方的id；示例值：6666000123120000
		IsDelayAllowance  string //补贴类型	String	2	Y	1:实补,2:后补,默认实补；示例值：1
		IsShare           string //是不是共享额度	String	4	N	示例值：
		MarketId          string //自定义活动编号	String	64	Y	示例值：ISFE00232
		MarketName        string //自定义活动名称	String	128	N	示例值：开业大促
		MarketDesc        string //自定义活动描述	String	64	N	示例值：新店开业大促
		PosCreditLimitAmt string //pos贷记卡补贴额度	String	16	Y	示例值：5.00
		PosDebitLimitAmt  string //pos借记卡补贴额度	String	16	Y	示例值：2.00
		PosLimitAmt       string //pos补贴额度	String	16	Y	示例值：4.00
		QrLimitAmt        string //扫码补贴额度	String	16	Y	示例值：1.00
		CreateBy          string //创建人	String	32	Y	示例值：Lg2022022201394910571
		CreateTime        string //创建时间	String	32	Y	示例值：2022-04-14 22:00:30
		UpdateTime        string //更新时间	String	32	Y	示例值：2022-04-14 23:00:30
	} `json:"cur_allowance_config_infos"`
	NoAllowanceDesc  string `json:"no_allowance_desc"`
	ReceivableFeeAmt string `json:"receivable_fee_amt"`
}
type MerDevLocation struct {
	TerminalIp       string //	交易设备IP	String	64	N	绑卡设备（付款 APP） 所在的公网IP，可用于定位所属地区，不是 wifi 连接时的局域网 IP。局域网 IP 包括：
	TerminalLocation string //	终端实时经纬度信息	String	32	N
}
type DcResponse struct {
	MerchantId     string `json:"merchant_id,omitempty"`      //商户号	String	35	N	数字货币: 工行；示例值：S5088295305
	TermId         string `json:"term_id,omitempty"`          //终端号	String	32	N	数字货币: 工行；示例值：58000001
	TradeType      string `json:"trade_type,omitempty"`       //交易类型	String	14	N	数字货币: 工行；示例值：SCANPAY
	CustomBankCode string `json:"custom_bank_code,omitempty"` //客户所属运营机构代码	String	14	N	数字货币: 工行；示例值：
	CustomBankName string `json:"custom_bank_name,omitempty"` //客户所属运营机构名称	String	70	N	数字货币: 工行；示例值：
	Openid         string `json:"openid,omitempty"`           //用户标识	String	64	N	数字货币: 工行；示例值：oGhiSxIAPtEnPfe9Xo000000B
	SubOpenid      string `json:"sub_openid,omitempty"`       //用户子标识	String	64	N	数字货币: 工行；示例值：oWNHX5RNaCUmZR
	CouponAmount   string `json:"coupon_amount,omitempty"`    //代金券金额	String	14	N	数字货币: 工行，单位：元；示例值：0.01
	CouponCount    string `json:"coupon_count,omitempty"`     //代金券数量	String	3	N	数字货币: 工行；示例值：1
	CouponList     []struct {
		CouponId     string `json:"coupon_id,omitempty" json:"coupon_id,omitempty"`         //代金券ID	String	40	N	示例值：2020060702001
		CouponType   string `json:"coupon_type,omitempty" json:"coupon_type,omitempty"`     //	代金券类型	String	8	N	示例值：5
		CouponAmount string `json:"coupon_amount,omitempty" json:"coupon_amount,omitempty"` //	单个代金券支付金额	String	14	N	单位：元；示例值：50.00
	} `json:"coupon_list,omitempty"` //代金券集合	Array	2048	N	数字货币: 工行； 示例值：[{"coupon_type":"5","coupon_id":"2020060702001","coupon_amount":"50.00"}]
	CreditorWalletName        string `json:"creditorWalletName,omitempty"`        //收款钱包名称	String	60	N	数字货币: 邮储；示例值：上海汇涵信息科技服务有限公司
	CreditorWalletId          string `json:"creditorWalletId,omitempty"`          //收款钱包ID	String	68	N	数字货币: 邮储；示例值：002250******5081
	CreditorWalletType        string `json:"creditorWalletType,omitempty"`        //收款钱包类型	String	4	N	数字货币: 邮储；收款钱包类型： WT01: 个人钱包  WT02: 子个人钱包 WT09: 对公钱包 WT10: 子对公钱包；
	CreditorWalletLevel       string `json:"creditorWalletLevel,omitempty"`       //收款钱包等级	String	4	N	数字货币: 邮储；收款钱包等级： WL01: 一类钱包 WL02: 二类钱包 WL03: 三类钱包 WL04: 四类钱包； 示例值：WT01
	DebtorWalletName          string `json:"debtorWalletName,omitempty"`          //付款钱包名称	String	60	N	数字货币: 邮储；示例值：我的钱包
	DebtorWalletId            string `json:"debtorWalletId,omitempty"`            //付款钱包ID	String	68	N	数字货币: 邮储数字货币-邮储；示例值：004100******0032
	DebtorWalletType          string `json:"debtorWalletType,omitempty"`          //付款钱包类型	String	4	N	数字货币: 邮储；付款钱包类型： WT01: 个人钱包 WT02: 子个人钱包 WT09: 对公钱包 WT10: 子对公钱包 示例值：WT01
	DebtorWalletLevel         string `json:"debtorWalletLevel,omitempty"`         //付款钱包等级	String	4	N	数字货币: 邮储；付款钱包等级： WL01: 一类钱包 WL02: 二类钱包 WL03: 三类钱包 WL04: 四类钱包 示例值：WT02
	DebtorPartyIdentification string `json:"debtorPartyIdentification,omitempty"` //付款运营机构	String	14	N	数字货币: 邮储；示例值：
	BusinessType              string `json:"businessType,omitempty"`              //支付类型	String	1	N	数字货币: 邮储；支付类型：O-本贷他，I-本贷本；示例值：0
}
