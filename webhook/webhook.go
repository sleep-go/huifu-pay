package webhook

import "github.com/sleep-go/huifu-pay/api/trade"

// PayWxLite 微信小程序支付
type PayWxLite struct {
	TransDate      string           `json:"trans_date"`
	BatchId        string           `json:"batch_id"`
	WxResponse     trade.WxResponse `json:"wx_response"`
	BankSeqId      string           `json:"bank_seq_id"`
	TransAmt       string           `json:"trans_amt"`
	Remark         string           `json:"remark"`
	EventDefineNo  string           `json:"event_define_no"`
	DiscountAmt    string           `json:"discount_amt"`
	AcctId         string           `json:"acct_id"`
	IsDiv          string           `json:"is_div"`
	FeeFlag        int              `json:"fee_flag"`
	TransOrderInfo struct {
		RefNum string `json:"ref_num"`
	} `json:"trans_order_info"`
	MerPriv         string                  `json:"mer_priv"`
	SettlementAmt   string                  `json:"settlement_amt"`
	AcctSplitBunch  trade.AcctSplitBunch    `json:"acct_split_bunch"`
	TransType       string                  `json:"trans_type"`
	MerOrdId        string                  `json:"mer_ord_id"`
	TransStat       string                  `json:"trans_stat"`
	CouponInfos     string                  `json:"coupon_infos"`
	TransTime       string                  `json:"trans_time"`
	HfSeqId         string                  `json:"hf_seq_id"`
	RefNo           string                  `json:"ref_no"`
	FeeAmount       string                  `json:"fee_amount"`
	OutTransId      string                  `json:"out_trans_id"`
	PartyOrderId    string                  `json:"party_order_id"`
	IsDelayAcct     string                  `json:"is_delay_acct"`
	Namespace       string                  `json:"namespace"`
	FeeFormulaInfos []trade.FeeFormulaInfos `json:"fee_formula_infos"`
	HuifuId         string                  `json:"huifu_id"`
	MerName         string                  `json:"mer_name"`
}

// PayaliQr 支付宝二维码支付
type PayaliQr struct {
	TransDate      string `json:"trans_date"`
	BatchId        string `json:"batch_id"`
	BankSeqId      string `json:"bank_seq_id"`
	TransAmt       string `json:"trans_amt"`
	Remark         string `json:"remark"`
	EventDefineNo  string `json:"event_define_no"`
	DiscountAmt    string `json:"discount_amt"`
	AcctId         string `json:"acct_id"`
	IsDiv          string `json:"is_div"`
	FeeFlag        int    `json:"fee_flag"`
	TransOrderInfo struct {
		RefNum string `json:"ref_num"`
	} `json:"trans_order_info"`
	MerPriv         string                  `json:"mer_priv"`
	SettlementAmt   string                  `json:"settlement_amt"`
	AcctSplitBunch  trade.AcctSplitBunch    `json:"acct_split_bunch"`
	TransType       string                  `json:"trans_type"`
	MerOrdId        string                  `json:"mer_ord_id"`
	TransStat       string                  `json:"trans_stat"`
	CouponInfos     string                  `json:"coupon_infos"`
	TransTime       string                  `json:"trans_time"`
	RefNo           string                  `json:"ref_no"`
	HfSeqId         string                  `json:"hf_seq_id"`
	FeeAmount       string                  `json:"fee_amount"`
	OutTransId      string                  `json:"out_trans_id"`
	PartyOrderId    string                  `json:"party_order_id"`
	AlipayResponse  trade.AlipayResponse    `json:"alipay_response"`
	IsDelayAcct     string                  `json:"is_delay_acct"`
	Namespace       string                  `json:"namespace"`
	FeeFormulaInfos []trade.FeeFormulaInfos `json:"fee_formula_infos"`
	HuifuId         string                  `json:"huifu_id"`
	MerName         string                  `json:"mer_name"`
}

// TransClose 交易关单
type TransClose struct {
	OrgReqDate    string `json:"org_req_date"`
	OrgReqSeqId   string `json:"org_req_seq_id"`
	TransStat     string `json:"trans_stat"`
	Namespace     string `json:"namespace"`
	HuifuId       string `json:"huifu_id"`
	EventDefineNo string `json:"event_define_no"`
	OrgMerOrdId   string `json:"org_mer_ord_id"`
	TransType     string `json:"trans_type"`
	OrgTransStat  string `json:"org_trans_stat"`
}

// RefundStandard 交易退款
type RefundStandard struct {
	DevsId         string               `json:"devs_id"`
	TransDate      string               `json:"trans_date"`
	OrgReqDate     string               `json:"org_req_date"`
	BankMessage    string               `json:"bank_message"`
	EventDefineNo  string               `json:"event_define_no"`
	ActualRefAmt   string               `json:"actual_ref_amt"`
	OrgOrdAmt      string               `json:"org_ord_amt"`
	OrgReqSeqId    string               `json:"org_req_seq_id"`
	TotalRefAmt    string               `json:"total_ref_amt"`
	MerPriv        string               `json:"mer_priv"`
	AcctSplitBunch trade.AcctSplitBunch `json:"acct_split_bunch"`
	TransType      string               `json:"trans_type"`
	OrgTermOrdId   string               `json:"org_term_ord_id"`
	BankCode       string               `json:"bank_code"`
	OrdAmt         string               `json:"ord_amt"`
	MerOrdId       string               `json:"mer_ord_id"`
	TransStat      string               `json:"trans_stat"`
	RespDesc       string               `json:"resp_desc"`
	TransTime      string               `json:"trans_time"`
	HfSeqId        string               `json:"hf_seq_id"`
	PartyOrderId   string               `json:"party_order_id"`
	OrgFeeAmt      string               `json:"org_fee_amt"`
	ReqSeqId       string               `json:"req_seq_id"`
	Namespace      string               `json:"namespace"`
	RespCode       string               `json:"resp_code"`
	HuifuId        string               `json:"huifu_id"`
	TotalRefFeeAmt string               `json:"total_ref_fee_amt"`
}

type HostingPay struct {
	BankCode       string `json:"bank_code"`
	BatchId        string `json:"batch_id"`
	TransStat      string `json:"trans_stat"`
	BankMessage    string `json:"bank_message"`
	BankSeqId      string `json:"bank_seq_id"`
	Remark         string `json:"remark"`
	TransAmt       string `json:"trans_amt"`
	GoodsDesc      string `json:"goods_desc"`
	EventDefineNo  string `json:"event_define_no"`
	HostingType    string `json:"hosting_type"`
	JsonData       string `json:"json_data"`
	PayWay         string `json:"pay_way"`
	PayTime        string `json:"pay_time"`
	RefNum         string `json:"ref_num"`
	PartyOrderId   string `json:"party_order_id"`
	ProjectId      string `json:"project_id"`
	TopOrderId     string `json:"top_order_id"`
	Namespace      string `json:"namespace"`
	TradeType      string `json:"trade_type"`
	HostingOrderId string `json:"hosting_order_id"`
	CreateDate     string `json:"create_date"`
	FeeAmt         string `json:"fee_amt"`
}

// HostingRefund 托管交易退款
type HostingRefund struct {
	RefundOrderId string `json:"refund_order_id"`
	ProjectId     string `json:"project_id"`
	TransStat     string `json:"trans_stat"`
	OrigOrderId   string `json:"orig_order_id"`
	Namespace     string `json:"namespace"`
	RefundAmt     string `json:"refund_amt"`
	Remark        string `json:"remark"`
	EventDefineNo string `json:"event_define_no"`
	CreateDate    string `json:"create_date"`
	RefundDate    string `json:"refund_date"`
}

// HostingAccounted 托管交易账务异步
type HostingAccounted struct {
	RespDesc       string `json:"resp_desc"`
	TransStat      string `json:"trans_stat"`
	ReqSeqId       string `json:"req_seq_id"`
	ReqDate        string `json:"req_date"`
	RespCode       string `json:"resp_code"`
	TransFinshTime string `json:"trans_finsh_time"`
	HfSeqId        string `json:"hf_seq_id"`
	HuifuId        string `json:"huifu_id"`
	AcctDate       string `json:"acct_date"`
	AcctStat       string `json:"acct_stat"`
	TransType      string `json:"trans_type"`
}
