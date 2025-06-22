package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/duke-git/lancet/v2/fileutil"
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
	response, raw, err := tr.V3TradePaymentJspay(trade.V3TradePaymentJspayRequest{
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
		log.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", response)
}
func TestV3TradePaymentScanpayQuery(t *testing.T) {
	response, raw, err := tr.V3TradePaymentScanpayQuery(trade.V3TradePaymentScanpayQueryRequest{
		HuifuId:     tr.HuifuPay.BsPay.Msc.SysId,
		OrgReqDate:  "20250618",
		OutOrdId:    "",
		OrgHfSeqId:  "0056default250618150144P870ac1367c100000",
		OrgReqSeqId: "RH20250618150110504170",
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", response)
}
