package main

import (
	"fmt"
	"ktpay/common"
	"ktpay/user"
	"testing"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
)

var (
	pay *BsPaySdk.BsPay
	bd  user.BasicData
)

func init() {
	var err error
	pay, err = BsPaySdk.NewBsPay(true, "./config/config.json")
	if err != nil {
		BsPaySdk.BspayPrintln(err)
	}
	bd = user.BasicData{
		HuifuPay: common.NewHuifuPay(pay),
	}

}
func TestV2MerchantBasicdataQueryRequest(t *testing.T) {
	response, raw, err := bd.V2MerchantBasicdataQueryRequest()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileutil.WriteStringToFile("./1.json", raw, false)
	for _, info := range response.Data.QryWxConfList.Decode() {
		fmt.Printf("======info=======\n%+v\n", info)
	}
}
