# 汇付天下api文档

https://paas.huifu.com/open/doc/api 


### 三方库
https://www.golancet.cn/

### 下单
```go
package tests

import (
	"testing"

	"github.com/huifurepo/bspay-go-sdk/ut/tool"
	"github.com/sleep-go/huifu-pay/api/trade"
	"github.com/sleep-go/huifu-pay/common"
)

var tr *trade.Trade

func init() {
	huifuPay := common.NewHuifuPay(true, "../config/config.json")
	tr = trade.NewTrade(huifuPay)
}
func TestJSpay(t *testing.T) {
	jspay, raw, err := tr.V3TradePaymentJspay(trade.V3TradePaymentJspayRequest{
		ReqDate:   tool.GetCurrentDate(),
		ReqSeqId:  tool.GetReqSeqId(),
		HuifuId:   tr.HuifuPay.BsPay.Msc.SysId,
		AcctId:    "",
		GoodsDesc: "GoodsDesc",
		TradeType: common.TradeType_A_NATIVE,
		TransAmt:  "0.01",
		ExtendInfos: trade.V3TradePaymentJspayExtendInfos{
			NotifyUrl: "",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(raw, jspay)
}
```
