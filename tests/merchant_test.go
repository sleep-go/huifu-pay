package tests

import (
	"fmt"
	"testing"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/sleep-go/huifu-pay/api/merchant"
	"github.com/sleep-go/huifu-pay/common"
)

var (
	bd *merchant.BasicData
)

func init() {
	huifuPay := common.NewHuifuPay(true, "../config/config.json")
	bd = merchant.NewBasicData(huifuPay)
}
func TestV2MerchantBasicdataQueryRequest(t *testing.T) {
	response, raw, err := bd.V2MerchantBasicdataQuery()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileutil.WriteStringToFile("./1.json", raw, false)
	for _, info := range response.Data.AgreementInfoList.Decode() {
		fmt.Printf("======info=======\n%+v\n", info)
	}
}
