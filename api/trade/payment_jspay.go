package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V3TradePaymentJspay 聚合正扫
// POST https://api.huifu.com/v3/trade/payment/jspay
//
// 应用场景
// 场景1-台牌码：用户通过扫描聚合静态二维码，输入订单金额完成支付操作。
// 场景2-公众号/小程序商城：用户通过公众号/小程序下单，输入密码完成支付操作。
// 场景3-商户根据用户选择支付方式，调用该接口生成支付宝/银联/数字人民币，用户使用对应app下“扫一扫”完成支付操作。
// 支持微信公众号-T_JSAPI、微信小程序-T_MINIAPP、支付宝JS-A_JSAPI、支付宝正扫-A_NATIVE、银联二维码正扫-U_NATIVE、银联二维码JS-U_JSAPI、数字人民币正扫-D_NATIVE、T_H5：微信直连H5支付、T_APP：微信APP支付、T_NATIVE：微信正扫交易；
//
// 适用对象
// 开通微信/支付宝/银联二维码/数字人民币权限的商户
func (t *Trade) V3TradePaymentJspay(req V3TradePaymentJspayRequest) (res *V3TradePaymentJspayResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V3TradePaymentJspayRequest(BsPaySdk.V3TradePaymentJspayRequest{
		ReqDate:     req.ReqDate,
		ReqSeqId:    req.ReqSeqId,
		HuifuId:     req.HuifuId,
		GoodsDesc:   req.GoodsDesc,
		TradeType:   string(req.TradeType),
		TransAmt:    req.TransAmt,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V3TradePaymentJspayResponse](resp)
}

type WxData struct {
	SubAppid  string `json:"sub_appid"`  //子商户应用ID	String	32	N	子商户在微信申请的应用ID，全局唯一。走聚合正扫发货管理的商户，使用的微信公众号/小程序支付 需要填写sub_appid+sub_openid 示例值：wxd678efh567hg6999
	SubOpenid string `json:"sub_openid"` //子商户用户标识	String	128	N	公众号和小程序场景必填。用户在子商户sub_appid下的唯一标识。下单前需获取到用户的sub_openid，sub_openid获取详见微信文档openid获取。 示例值：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o
	Attach    string `json:"attach"`     //附加数据	String	127	N	在查询api和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据；
	Body      string `json:"body"`       //商品描述 商品或支付单简要描述，格式要求：门店品牌名-城市分店名-实际商品名称；
	Detail    struct {
		CostPrice   string `json:"cost_price"` //订单原价(元)	String	12	N	1.商户侧一张小票订单可能被分多次支付，订单原价用于记录整张小票的交易金额。 2.当订单原价与支付金额不相等，则不享受优惠。3.该字段主要用于防止同一张小票分多次支付，以享受多次优惠的情况，正常支付订单不必上传此参数。
		ReceiptId   string `json:"receipt_id"` //商品小票ID
		GoodsDetail []struct {
			GoodsId      string `json:"goods_id"`       //商品编码	String	32	N	由半角的大小写字母、数字、中划线、下划线中的一种或几种组成；示例值：6934572310301
			GoodsName    string `json:"goods_name"`     //商品名称	String	256	N	商品的实际名称；示例值：太龙双黄连口服液
			Price        string `json:"price"`          //商品单价(元)	String	12	N	如果商户有优惠，需传输商户优惠后的单价； 例如：用户对一笔 100 元的订单使用了商场发的优惠券 100-50，则活动商品的单价应为原单价-50 示例值：43.00
			Quantity     string `json:"quantity"`       //用户购买的数量；示例值：1
			WxpayGoodsId string `json:"wxpay_goods_id"` //微信侧商品编码	String	32	N	示例值：12235413214070356458058
		} `json:"goods_detail"`
	} `json:"detail"`
	DeviceInfo string `json:"device_info"` //设备号	String	32	N	终端设备号(商户自定义，如门店编号)，示例值：660035730311000126101 注意：H5与小程序支付填以下值： 苹果APP：1；安卓APP：2；IOS手机网站：3；ANDROID手机网站：4
	GoodsTag   string `json:"goods_tag"`   //订单优惠标记	String	32	N	代金券或立减优惠功能的参数；示例值：WXG
	Identity   string `json:"identity"`    //实名支付	String	128	N	实名支付功能，用于公安和保险类商户使用, 包含类型、证件号、姓名三个子域。 示例值："{\"type":\"IDCARD\",\"number\":\"111111111111\",\"name\":\"张三\"}"
	Receipt    string `json:"receipt"`     //开发票入口开放标识	String	8	N	示例值：Y
	SceneInfo  struct {
		StoreInfo struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			AreaCode string `json:"area_code"`
			Address  string `json:"address"`
		}
	} `json:"scene_info"` //场景信息	Object		N	该字段用于上报场景信息，目前支持上报实际门店信息。
	SpbillCreateIp string `json:"spbill_create_ip"` //终端ip	String	64	N	调用微信支付API的机器IP；示例值：172.28.52.52
	PromotionFlag  string `json:"promotion_flag"`   //单品优惠标识	String	1	N	Y-是，N-否，默认否；直连模式需要填写；示例值：Y 若使用单品优惠，该字段必填，若该字段为Y，则商品详情【detail】必填
	ProductId      string `json:"product_id"`       //新增商品ID	String	32	N	直连模式【trade_type】=T_NATIVE支付的时候必填；示例值：
	LimitPayer     string `json:"limit_payer"`      //指定支付者	String	5	N	上传此参数，可限制用户只有是成年人才能支付，示例值：ADULT
}
type AlipayData struct {
	AlipayStoreId string `json:"alipay_store_id"` //支付宝的店铺编号	String	32	N	示例值：2016041400077000000003314986
	BuyerId       string `json:"buyer_id"`        //买家的支付宝唯一用户号	String	28	N	示例值：2088202954065786；
	BuyerLogonId  string `json:"buyer_logon_id"`  //买家支付宝账号
	ExtendParams  struct {
		CardType string `json:"card_type"` //卡类型	String	32	N	示例值：S0JP0000
		//支付宝点餐场景类型	String	20	N	QR_ORDER（店内扫码点餐）
		//PRE_ORDER（预点到店自提）
		//HOME_DELIVERY（外送到家）
		//DIRECT_PAYMENT（直接付款）
		//QR_FOOD_ORDER（点餐先付）
		//P_QR_FOOD_ORDER（点餐后付）
		//SELF_PICK（门店自提）
		//TAKE_OUT （餐饮外卖）
		//OTHER（其他）
		//该参数只适用于支付宝支付窗交易接口；示例值：TAKE_OUT
		FoodOrderType        string `json:"food_order_type"`
		HbFqNum              string `json:"hb_fq_num"`               //花呗分期数	String	5	N	使用花呗分期要进行的分期数；示例值：3
		HbFqSellerPercent    string `json:"hb_fq_seller_percent"`    //花呗卖家手续费百分比	String	3	N	使用花呗分期需要卖家承担的手续费比，单位比例的百分值。 花呗商贴支付默认传0，示例值：0
		IndustryRefluxInfo   string `json:"industry_reflux_info"`    //行业数据回流信息	String	64	N	示例值：{"scene_code":"metro_tradeorder","channel":"xxxx","scene_data":{"asset_name":"ALIPAY"}}
		FqChannels           string `json:"fq_channels"`             //信用卡分期资产方式	String	20	N	代表优先使用资产类型；alipayfq_cc：表示信⽤卡分期，为空时默认花呗。示例值：alipayfq_cc
		ParkingId            string `json:"parking_id"`              //停车场id	String	28	N	isv停车场id、向支付宝停车平台申请获得的支付宝停车场的唯一标识； 示例值：PI1504848980306666666
		SysServiceProviderId string `json:"sys_service_provider_id"` //系统商编号	String	64	N	该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的pid；
	} `json:"extend_params"` //业务扩展参数
	OperatorId  string `json:"operator_id"`
	ProductCode string `json:"product_code"`
	GoodsDetail []struct {
		GoodsId        string `json:"goods_id"`
		GoodsName      string `json:"goods_name"`
		Price          string `json:"price"`
		Quantity       string `json:"quantity"`
		Body           string `json:"body"`
		CategoriesTree string `json:"categories_tree"`
		ShowUrl        string `json:"show_url"`
		GoodsCategory  string `json:"goods_category"`
	} `json:"goods_detail"`
	MerchantOrderNo string `json:"merchant_order_no"`
	SellerId        string `json:"seller_id"`
	StoreId         string `json:"store_id"`
	ExtUserInfo     struct {
		Name   string `json:"name"`   //姓名	String	16	N	注：need_check_info=T时，该参数才有效；示例值：张三
		Mobile string `json:"mobile"` //手机号	String	20	N	注：该参数暂不校验；示例值：13911111111
		//证件类型	String	32	N	身份证：IDENTITY_CARD，
		//		护照：PASSPORT；
		//		军官证：OFFICER_CARD，
		//		士兵证：SOLDIER_CARD；
		//		户口本：HOKOU；
		//		示例值：IDENTITY_CARD
		//		注：need_check_info=T时，该参数才有效
		CertType      string `json:"cert_type"`
		CertNo        string `json:"cert_no"`         //证件号	String	64	N	示例值：Ly+fnExeyPOTzfOtgRRur77nJB9TAe4PGgK9M ， 需要密文传输，请参考加密解密说明使用汇付RSA公钥加密。 注：need_check_info=T时，该参数才有效
		MinAge        string `json:"min_age"`         //	允许的最小买家年龄	String	3	N	买家年龄必须大于等于所传数值。示例值：18 注：1. need_check_info = T 时该参数才有效，2. min_age 为整数，必须大于等于 0
		FixBuyer      string `json:"fix_buyer"`       //是否强制校验付款人身份信息	String	8	N	T：强制校验，F：不强制；示例值：T
		NeedCheckInfo string `json:"need_check_info"` //是否强制校验身份信息	String	1	N	T：强制校验，F：不强制；示例值：F
	} `json:"ext_user_info"`
	Subject           string `json:"subject"`
	StoreName         string `json:"store_name"`
	OpAppId           string `json:"op_app_id"`
	AliBusinessParams string `json:"ali_business_params"`
	Body              string `json:"body"`
	AliPromoParams    string `json:"ali_promo_params"`
}
type UnionpayData struct {
	QrCode        string `json:"qr_code"`        //	二维码	String	256	N	台牌码的url，交易类型为U_JSAPI:银联JS时必填
	AddnData      string `json:"addn_data"`      //收款方附加数据	String	3000	N	请参考银联收款方附加数据(addn_data)说明
	AreaInfo      string `json:"area_info"`      //地区信息	String	32	N	示例值：310000
	CustomerIp    string `json:"customer_ip"`    //持卡人ip	String	40	N	持卡人确认付款时的ip地址，用于防钓鱼。（js支付必填）示例值：127.1.1.1
	FrontUrl      string `json:"front_url"`      //前台通知地址	String	200	N	收款方向银联推送订单时上送的前台通知地址（仅允许为外网地址）； 用户完成支付点击“返回”后，银联通过浏览器post请求到该地址。 示例值：http://www.huifu.com
	OrderDesc     string `json:"order_desc"`     //订单描述	String	200	N	示例值：订单描述
	PayeeComments string `json:"payee_comments"` //收款方附言	String	100	N	示例值：业务收款
	PayeeInfo     struct {
		MerCatCode string `json:"mer_cat_code"` //商户类别	String	4	N	参考银联商户类别；示例值：0101
		SubId      string `json:"sub_id"`       //二级商户代码	String	20	N	示例值：823586070110039
		SubName    string `json:"sub_name"`     //二级商户名称	String	100	N	示例值：上海白乐门酒店
		TermId     string `json:"term_id"`      //终端号	String	8	N	示例值：58000001
	} `json:"payee_info"` //收款方信息	Object		N
	PnrInsIdCd  string `json:"pnr_ins_id_cd"` //银联分配的服务商机构标识码	String	11	N	示例值：01008330
	ReqReserved string `json:"req_reserved"`  //请求方自定义域	String	500	N	示例值：
	TermInfo    string `json:"term_info"`     //终端信息	String	32	N	示例值：
	UserId      string `json:"user_id"`       //银联用户标识	String	128	N	调用获取银联用户标识接口会返回【user_id】； 示例值：gaqiMrRnKwwOZO7dNtUc349YTMaa3HkRZg+OMU+46ysDzn6flfomHP88qOvH+6yG
}
type DcData struct {
	DigitalBankNo string `json:"digital_bank_no"` //数字货币银行编号	String	5	N	01002: 工行，01004: 中行， 01005: 建行，01008：邮储
}
type AcctSplitBunch struct {
	AcctInfos []struct {
		DivAmt        string `json:"div_amt"`  //分账金额	String	14	N	单位元，需保留小数点后两位，示例值：1.00 ，最低传入0.01
		HuifuId       string `json:"huifu_id"` //分账接收方ID	String	32	Y	斗拱开户时生成；示例值：6666000123120001
		AcctId        string `json:"acct_id"`  //账户号	String	9	N	可指定账户号，仅支持基本户、现金户，不填默认为基本户；示例值：F00598600
		AcctDate      string `json:"acct_date"`
		PercentageDiv string `json:"percentage_div"` //分账百分比%	String	6	N	示例值：23.50，表示23.50%。仅在percentage_flag=Y时起作用 acct_infos中全部分账百分比之和必须为100.00%。
	} `json:"acct_infos"` //分账明细	Array		N	jsonArray分账明细
	PercentageFlag string `json:"percentage_flag"` //百分比分账标志	String	1	N	Y:使用百分比分账；示例值：Y
	IsCleanSplit   string `json:"is_clean_split"`  //是否净值分账	String	1	N	Y:使用净值分账，仅在交易手续费内扣且使用百分比分账时起作用；示例值：Y
	FeeAmt         string `json:"fee_amt"`
	FeeHuifuId     string `json:"fee_huifu_id"`
	FeeAcctId      string `json:"fee_acct_id"`
	FeeAcctDate    string `json:"fee_acct_date"`
}
type CombinedpayData struct {
	HuifuId  string `json:"huifu_id"`  //补贴方汇付商户号	String	32	Y	补贴方汇付ID；示例值：6666000123123123
	UserType string `json:"user_type"` //补贴方类型	String	32	Y	补贴方类型：channel-渠道，agent-代理；示例值：agent
	AcctId   string `json:"acct_id"`   //补贴方账户号	String	32	Y	营销补贴方账户号；示例值：F00900982
	Amount   string `json:"amount"`    //补贴金额	String	14	Y	单位元，需保留小数点后两位，例如：1.00,最低传入0.01； 示例值：1.01
}
type RiskCheckData struct {
	IpAddr string `json:"ip_addr"` //ip地址	String	32	N	IP地址、经纬度、基站地址最少要送其中一项；示例值：172.28.52.52
	//基站地址	String	32	N	【mcc】+【mnc】+【location_cd】+【lbs_num】
	//	- mcc:移动国家代码，460代表中国；3位长
	//	- mnc：移动网络号码；2位长；
	//	- location_cd：位置区域码，16进制，5位长
	//	- lbs_num：基站编号，16进制，5位长
	//	- 注意若位数不足用空格补足；
	//	示例值：460001039217563，其中460（mcc)， 00(mnc)，10392(location_cd)， 17563(lbs_num)
	BaseStation  string `json:"base_station"`
	Latitude     string `json:"latitude"`  //纬度	String	32	N	格式：+表示北纬，-表示南纬。纬度整数位不超过2位，小数位不超过6位。示例值：+37.12； IP地址、经纬度、基站地址最少要送其中一项
	Longitude    string `json:"longitude"` //经度	String	32	N	格式：+表示东经，-表示西经；经度整数位不超过3位，小数位不超过5位；示例值：-121.213；IP地址、经纬度、基站地址最少要送其中一项
	SubProduct   string `json:"sub_product"`
	TransferType string `json:"transfer_type"`
}
type TerminalDeviceData struct {
	DeviceType    string `json:"device_type"`     //设备类型	String	2	N	1: 手机，2: 平板，3: 手表，4: PC；示例值：1
	DeviceIp      string `json:"device_ip"`       //交易设备IP	String	64	N	绑卡设备所在的公网IP，可用于定位所属地区，不是wifi连接时的局域网IP。 ABCD:EF01:2345:6789:ABCD:EF01:2345:6789（IPv6）； 目前暂传IPv4格式。示例值：10.10.0.1（IPv4）；
	DeviceMac     string `json:"device_mac"`      //交易设备MAC	String	64	N	示例值：F0E1D2C3B4A5
	DeviceImei    string `json:"device_imei"`     //交易设备IMEI	String	64	N	移动终端设备的唯一标识；示例值：460030912121001
	DeviceImsi    string `json:"device_imsi"`     //交易设备IMSI	String	64	N	示例值：460030912121001
	DeviceIccId   string `json:"device_icc_id"`   //交易设备ICCID	String	64	N	示例值：898600680113F0123014
	DeviceWifiMac string `json:"device_wifi_mac"` //交易设备WIFIMAC	String	64	N	示例值：968778695A4B
	DeviceGps     string `json:"device_gps"`      //交易设备GPS	String	64	N	示例值：20.346790,-4.654321
	AppVersion    string `json:"app_version"`     //商户终端应用程序版本	String	8	N	终端应用程序的版本号。应用程序变更应保证版本号不重复；示例值：3.2.5
	IccId         string `json:"icc_id"`          //SIM卡卡号	String	20	N	ICCID，SIM 卡卡号；示例值：898600680113F0123014
	//商户终端实时经纬度信息	String	32	N	受理终端设备实时经纬度信息；格式为纬度/经度，示例值：+37.12/-121.213；
	//	+表示北纬、东经，-表示南纬、西经。
	//	经度整数位不超过3位，小数位不超过5位；纬度整数位不超过2位，小数位不超过6位。
	//	注：银联AT交易时，location和mer_device_IP二选一必填其一；
	//	银联交易必填
	Location    string `json:"location"`
	MerDeviceIp string `json:"mer_device_ip"` //	商户交易设备IP	String	64	N	示例值：10.10.0.1；注：银联AT交易时，location和mer_device_IP二选一必填其一
	//	商户设备类型	String	2	Y	01：自动柜员机（含 ATM 和 CDM）和多媒体自助终端，02：传统 POS，
	//	03：mPOS，04：智能 POS，05：II 型固定电话，
	//	06：银联闪付终端，08：手机 POS，09：刷脸付终端，
	//	10：条码支付受理终端，11：条码支付辅助受理终端（如：台牌码），
	//	12：行业终端（公交、地铁用于指 定行业的终端），13：MIS 终端；
	//	示例值：11
	MerDeviceType   string `json:"mer_device_type"`
	MobileCountryCd string `json:"mobile_country_cd"` //移动国家代码	String	3	N	基站信息，由国际电联(ITU)统一分配的移动国家代码（MCC），中国为 460（默认）； 示例值：460
	MobileNetNum    string `json:"mobile_net_num"`    //移动网络号码	String	2	N	由国际电联(ITU) 统 一 分 配 的 移 动 网 络 号 码（MNC）；259号文新增字段 移动：00、02、04、07；联通：01、06、09； 电信：03、05、11； 示例值：00
	NetworkLicense  string `json:"network_license"`   //商户终端入网认证编号	String	5	N	银行卡受理终端产品入网认证编号。 该编号由“中国银联标识产品企业资质认证办公室”为通过入网认证的终端进行分配; 格式：5 位字符，示例值：P3100
	SerialNum       string `json:"serial_num"`        //商户终端序列号	String	50	N	终端设备的硬件序列号；示例值：000001041804CA878530 可查看机身条形码下方的设备序列号，或在操作员的版本菜单里查看；
	DevsId          string `json:"devs_id"`           //汇付机具号	String	32	Y	通过汇付报备的机具必传；示例值：660035730311000126101
}
type TransFeeAllowanceInfo struct {
	//补贴手续费金额	String	16	N	金额以元为单位，最少1分，示例值：0.01；
	//	1.补贴手续费金额小于或等于该笔手续费金额时，按照补贴手续费金额补贴；
	//	2.补贴手续费金额大于该笔手续费金额时，按照该笔交易实际手续费金额补贴；
	AllowanceFeeAmt string `json:"allowance_fee_amt"`
}
type V3TradePaymentJspayExtendInfos struct {
	TimeExpire            string                `json:"time_expire"`
	WxData                WxData                `json:"wx_data"`
	AlipayData            AlipayData            `json:"alipay_data"`
	UnionpayData          UnionpayData          `json:"unionpay_data"`
	DcData                DcData                `json:"dc_data"` //数字人民币参数集合
	DelayAcctFlag         string                `json:"delay_acct_flag"`
	FeeFlag               string                `json:"fee_flag"`
	AcctSplitBunch        AcctSplitBunch        `json:"acct_split_bunch"`
	TermDivCouponType     string                `json:"term_div_coupon_type"` //传入分账遇到优惠的处理规则	Integer	1	N	1: 按比例分，2: 按分账明细顺序保障，3: 只给交易商户（默认)；示例值：1
	CombinedpayData       []CombinedpayData     `json:"combinedpay_data"`
	LimitPayType          string                `json:"limit_pay_type"` //禁用信用卡标记	String	128	N	本次交易禁止使用的支付方式，默认不禁用；取值参见说明；示例值：NO_CREDIT
	FqMerDiscountFlag     string                `json:"fq_mer_discount_flag"`
	ChannelNo             string                `json:"channel_no"`
	PayScene              string                `json:"pay_scene"`
	Remark                string                `json:"remark"`
	RiskCheckData         RiskCheckData         `json:"risk_check_data"`
	TerminalDeviceData    TerminalDeviceData    `json:"terminal_device_data"`
	NotifyUrl             string                `json:"notify_url"`
	TransFeeAllowanceInfo TransFeeAllowanceInfo `json:"trans_fee_allowance_info"`
}
type V3TradePaymentJspayRequest struct {
	ReqDate     string           `json:"req_date"`
	ReqSeqId    string           `json:"req_seq_id"`
	HuifuId     string           `json:"huifu_id"`
	AcctId      string           `json:"acct_id"`    //账户号	String	9	N	可指定收款账户号，仅支持基本户、现金户，不填默认为基本户； 示例值：F00598600
	GoodsDesc   string           `json:"goods_desc"` //商品描述	String	127	Y	示例值：XX商品
	TradeType   common.TradeType `json:"trade_type"` //交易类型	String	16	Y
	TransAmt    string           `json:"trans_amt"`
	ExtendInfos V3TradePaymentJspayExtendInfos
}
type V3TradePaymentJspayResponse struct {
	Data struct {
		RespCode     string           `json:"resp_code"`
		RespDesc     string           `json:"resp_desc"`
		ReqDate      string           `json:"req_date"`
		ReqSeqId     string           `json:"req_seq_id"`
		HfSeqId      string           `json:"hf_seq_id"`
		TradeType    common.TradeType `json:"trade_type"` //交易类型	String	16	Y	T_JSAPI: 微信公众号
		TransAmt     string           `json:"trans_amt"`
		TransStat    string           `json:"trans_stat"`
		HuifuId      string           `json:"huifu_id"`
		BankMessage  string           `json:"bank_message"`
		PayInfo      string           `json:"pay_info"`
		Remark       string           `json:"remark"`
		AcctId       string           `json:"acct_id"`
		DeviceType   string           `json:"device_type"`
		PartyOrderId string           `json:"party_order_id"` //用户账单上的商户订单号
		AtuSubMerId  string           `json:"atu_sub_mer_id"` //ATU真实商户号	String	32	N	微信、支付宝、银联真实商户号；示例值：411111141
		UnconfirmAmt string           `json:"unconfirm_amt"`  //待确认金额	String	14	N	待确认金额；单位元。示例值：1.00
	} `json:"data"`
	Sign string `json:"sign"`
}

// V3TradePaymentJspayNotifyMessage 异步返回参数
type V3TradePaymentJspayNotifyMessage struct {
	RespCode string                                                        `json:"resp_code"`
	RespData common.StringObject[V3TradePaymentJspayNotifyMessageRespData] `json:"resp_data"`
	RespDesc string                                                        `json:"resp_desc"`
	Sign     string                                                        `json:"sign"`
}
type V3TradePaymentJspayNotifyMessageRespData struct {
	RespCode              string            `json:"resp_code"`
	RespDesc              string            `json:"resp_desc"`
	HuifuId               string            `json:"huifu_id"`
	ReqSeqId              string            `json:"req_seq_id"`
	ReqDate               string            `json:"req_date"`
	TransType             string            `json:"trans_type"`
	HfSeqId               string            `json:"hf_seq_id"`
	OutTransId            string            `json:"out_trans_id"`
	PartyOrderId          string            `json:"party_order_id"`
	TransAmt              string            `json:"trans_amt"`
	PayAmt                string            `json:"pay_amt"`
	SettlementAmt         string            `json:"settlement_amt"`
	EndTime               string            `json:"end_time"`
	AcctDate              string            `json:"acct_date"`
	TransStat             string            `json:"trans_stat"`
	FeeFlag               int               `json:"fee_flag"`
	FeeFormulaInfos       []FeeFormulaInfos `json:"fee_formula_infos"`
	FeeAmount             string            `json:"fee_amount"`
	CombinedpayFeeAmt     string            `json:"combinedpay_fee_amt"`
	TransFeeAllowanceInfo struct {
		ReceivableFeeAmt        string `json:"receivable_fee_amt"`
		ActualFeeAmt            string `json:"actual_fee_amt"`
		AllowanceFeeAmt         string `json:"allowance_fee_amt"`
		AllowanceType           string `json:"allowance_type"`
		NoAllowanceDesc         string `json:"no_allowance_desc"`
		CurAllowanceConfigInfos string `json:"cur_allowance_config_infos"`
	} `json:"trans_fee_allowance_info"`
	CombinedpayData  []CombinedpayData       `json:"combinedpay_data"`
	DebitType        string                  `json:"debit_type"`
	IsDiv            string                  `json:"is_div"`
	AcctSplitBunch   AcctSplitBunch          `json:"acct_split_bunch"`
	IsDelayAcct      string                  `json:"is_delay_acct"`
	WxUserId         string                  `json:"wx_user_id"`
	WxResponse       WxResponse              `json:"wx_response"`
	AlipayResponse   AlipayResponse          `json:"alipay_response"`
	DcResponse       DcResponse              `json:"dc_response"`
	UnionpayResponse common.UnionpayResponse `json:"unionpay_response"`
	DeviceType       string                  `json:"device_type"`
	MerDevLocation   MerDevLocation          `json:"mer_dev_location"`
	BankCode         string                  `json:"bank_code"`
	Remark           string                  `json:"remark"`
	FqChannels       string                  `json:"fq_channels"`
	NotifyType       string                  `json:"notify_type"`
	SplitFeeInfo     SplitFeeInfo            `json:"split_fee_info"`
	AtuSubMerId      string                  `json:"atu_sub_mer_id"`
	DevsId           string                  `json:"devs_id"`
	FundFreezeStat   string                  `json:"fund_freeze_stat"`
	AcctStat         string                  `json:"acct_stat"`
	AcctId           string                  `json:"acct_id"`
	AvoidSmsFlag     string                  `json:"avoid_sms_flag"`
	BagentId         string                  `json:"bagent_id"`
	BankDesc         string                  `json:"bank_desc"`
	BankMessage      string                  `json:"bank_message"`
	BankOrderNo      string                  `json:"bank_order_no"`
	BankSeqId        string                  `json:"bank_seq_id"`
	BaseAcctId       string                  `json:"base_acct_id"`
	BatchId          string                  `json:"batch_id"`
	ChannelType      string                  `json:"channel_type"`
	DelayAcctFlag    string                  `json:"delay_acct_flag"`
	DivFlag          string                  `json:"div_flag"`
	FeeAmt           string                  `json:"fee_amt"`
	FeeRecType       string                  `json:"fee_rec_type"`
	FeeType          string                  `json:"fee_type"`
	GateId           string                  `json:"gate_id"`
	MazeRespCode     string                  `json:"maze_resp_code"`
	MerName          string                  `json:"mer_name"`
	MerOrdId         string                  `json:"mer_ord_id"`
	MypaytsfDiscount string                  `json:"mypaytsf_discount"`
	NeedBigObject    bool                    `json:"need_big_object"`
	OrgAuthNo        string                  `json:"org_auth_no"`
	OrgHuifuSeqId    string                  `json:"org_huifu_seq_id"`
	OrgTransDate     string                  `json:"org_trans_date"`
	OutOrdId         string                  `json:"out_ord_id"`
	PayScene         string                  `json:"pay_scene"`
	PospSeqId        string                  `json:"posp_seq_id"`
	ProductId        string                  `json:"product_id"`
	RefNo            string                  `json:"ref_no"`
	RiskCheckData    RiskCheckData           `json:"risk_check_data"`
	RiskCheckInfo    string                  `json:"risk_check_info"`
	SubRespCode      string                  `json:"sub_resp_code"`
	SubRespDesc      string                  `json:"sub_resp_desc"`
	SubsidyStat      string                  `json:"subsidy_stat"`
	SysId            string                  `json:"sys_id"`
	TransDate        string                  `json:"trans_date"`
	TransTime        string                  `json:"trans_time"`
	TradeType        string                  `json:"trade_type"`
	TransFinishTime  string                  `json:"trans_finish_time"`
}

// V3TradePaymentJspayUnfreezeNotifyMessage 解冻异步返回参数
type V3TradePaymentJspayUnfreezeNotifyMessage struct {
	RespCode string                                                                `json:"resp_code"`
	RespData common.StringObject[V3TradePaymentJspayUnfreezeNotifyMessageRespData] `json:"resp_data"`
	RespDesc string                                                                `json:"resp_desc"`
	Sign     string
}
type V3TradePaymentJspayUnfreezeNotifyMessageRespData struct {
	RespCode       string `json:"resp_code"`        //业务返回码	String	8	Y	业务返回码
	RespDesc       string `json:"resp_desc"`        //业务返回描述	String	512	Y	业务返回描述
	HfSeqId        string `json:"hf_seq_id"`        //交易的汇付全局流水号	String	40	Y	示例值：00470topo1A221019132207P068ac1362af00000
	ReqSeqId       string `json:"req_seq_id"`       //交易请求流水号	String	128	Y	交易时传入，原样返回；示例值：rQ2021121311173944134649875651
	ReqDate        string `json:"req_date"`         //交易请求日期	String	8	Y	交易时传入，原样返回,格式为yyyyMMdd，示例值：20091225
	HuifuId        string `json:"huifu_id"`         //商户号	String	32	Y	示例值：6666000123120000
	NotifyType     string `json:"notify_type"`      //通知类型	Integer	1	Y	3：资金解冻通知；示例值：3
	FundFreezeStat string `json:"fund_freeze_stat"` //资金冻结状态	String	16	Y	UNFREEZE：解冻；示例值：UNFREEZE
	UnfreezeAmt    string `json:"unfreeze_amt"`     //解冻金额	String	14	Y	单元：元。示例值：1.23
	FreezeTime     string `json:"freeze_time"`      //冻结时间	String	14	Y	格式为yyyyMMddHHMMSS，示例值：20091225091010
	UnfreezeTime   string `json:"unfreeze_time"`    //解冻时间	String	14	Y	格式为yyyyMMddHHMMSS，示例值：20091225091010
}
