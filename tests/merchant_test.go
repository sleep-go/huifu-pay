package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
	"github.com/sleep-go/huifu-pay/api/merchant"
	"github.com/sleep-go/huifu-pay/common"
)

var (
	bd *merchant.Merchant
)

func init() {
	huifuPay := common.NewHuifuPay(true, "../config/config.json")
	bd = merchant.NewMerchant(huifuPay)
}
func TestV2MerchantBasicdataQueryRequest(t *testing.T) {
	response, raw, err := bd.V2MerchantBasicdataQuery()
	if err != nil {
		log.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	for _, info := range response.Data.AgreementInfoList.Decode() {
		fmt.Printf("======info=======\n%+v\n", info)
	}
}
func TestV2MerchantUrlForward(t *testing.T) {
	response, raw, err := bd.V2MerchantUrlForward(merchant.V2MerchantUrlForwardRequest{
		ReqSeqId:        tool.GetReqSeqId(),
		ReqDate:         tool.GetCurrentDate(),
		Phone:           "18502243993",
		StoreId:         "HS001",
		Expires:         "50000",
		BackPageUrl:     "",
		AsyncReceiveUrl: "",
		TemplateId:      "",
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======response======\n%+v\n", response.Data)
}
