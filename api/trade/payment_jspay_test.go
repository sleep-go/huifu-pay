package trade

import (
	"testing"

	"github.com/huifurepo/bspay-go-sdk/ut/tool"
	"github.com/sleep-go/huifu-pay/common"
)

var tr *Trade

func init() {
	huifuPay := common.NewHuifuPay(true, "../../config/config.json")
	tr = NewTrade(huifuPay)
}
func TestJSpay(t *testing.T) {
	jspay, raw, err := tr.V3TradePaymentJspay(V3TradePaymentJspayRequest{
		ReqDate:   tool.GetCurrentDate(),
		ReqSeqId:  tool.GetReqSeqId(),
		HuifuId:   tr.HuifuPay.BsPay.Msc.SysId,
		AcctId:    "",
		GoodsDesc: "GoodsDesc",
		TradeType: common.TradeType_A_NATIVE,
		TransAmt:  "0.01",
		ExtendInfos: struct {
			TimeExpire            string                `json:"time_expire"`
			WxData                WxData                `json:"wx_data"`
			AlipayData            AlipayData            `json:"alipay_data"`
			UnionpayData          UnionpayData          `json:"unionpay_data"`
			DcData                DcData                `json:"dc_data"` //数字人民币参数集合
			DelayAcctFlag         string                `json:"delay_acct_flag"`
			FeeFlag               string                `json:"fee_flag"`
			AcctSplitBunch        AcctSplitBunch        `json:"acct_split_bunch,omitempty"`
			TermDivCouponType     string                `json:"term_div_coupon_type"` //传入分账遇到优惠的处理规则	Integer	1	N	1: 按比例分，2: 按分账明细顺序保障，3: 只给交易商户（默认)；示例值：1
			CombinedpayData       CombinedpayData       `json:"combinedpay_data"`
			LimitPayType          string                `json:"limit_pay_type"` //禁用信用卡标记	String	128	N	本次交易禁止使用的支付方式，默认不禁用；取值参见说明；示例值：NO_CREDIT
			FqMerDiscountFlag     string                `json:"fq_mer_discount_flag"`
			ChannelNo             string                `json:"channel_no"`
			PayScene              string                `json:"pay_scene"`
			Remark                string                `json:"remark"`
			RiskCheckData         RiskCheckData         `json:"risk_check_data"`
			TerminalDeviceData    TerminalDeviceData    `json:"terminal_device_data"`
			NotifyUrl             string                `json:"notify_url"`
			TransFeeAllowanceInfo TransFeeAllowanceInfo `json:"trans_fee_allowance_info"`
		}{
			NotifyUrl: "",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(raw, jspay)
}
