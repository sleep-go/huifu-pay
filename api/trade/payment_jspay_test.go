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
		ExtendInfos: V3TradePaymentJspayExtendInfos{
			NotifyUrl: "",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(raw, jspay)
}
