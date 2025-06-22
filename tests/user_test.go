package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
	"github.com/sleep-go/huifu-pay/api/user"
	"github.com/sleep-go/huifu-pay/common"
)

var (
	u *user.User
)

func init() {
	huifuPay := common.NewHuifuPay(true, "../config/config.json")
	u = user.NewUser(huifuPay)
}

func TestV2UserBasicdataIndv(t *testing.T) {
	req := user.V2UserBasicdataIndvRequest{
		ReqSeqId:         tool.GetReqSeqId(),
		ReqDate:          tool.GetCurrentDate(),
		Name:             "张三",
		CertType:         common.CertType00,
		CertNo:           "",
		CertValidityType: "",
		CertNationality:  "",
		CertBeginDate:    "",
		CertEndDate:      "",
		MobileNo:         "",
		Address:          "",
		Email:            "",
		LoginName:        "",
		SmsSendFlag:      "",
		ExpandId:         "",
		FileList: []struct {
			FileId   string `json:"file_id"`
			FileName string `json:"file_name"`
			FileType string `json:"file_type"`
		}{
			{
				FileId:   "123",
				FileName: "123",
				FileType: "123",
			},
		},
		Mcc:        "Mcc",
		ProvId:     "",
		AreaId:     "",
		DistrictId: "",
	}
	response, raw, err := u.V2UserBasicdataIndv(req)
	if err != nil {
		log.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", response)
}
func TestV2UserBasicdataIndvModify(t *testing.T) {
	response, raw, err := u.V2UserBasicdataIndvModify(user.V2UserBasicdataIndvModifyRequest{
		ReqSeqId:         tool.GetReqSeqId(),
		ReqDate:          tool.GetCurrentDate(),
		HuifuId:          u.HuifuPay.BsPay.Msc.SysId,
		CertValidityType: "",
		CertBeginDate:    "",
		CertEndDate:      "",
		Email:            "",
		MobileNo:         "",
		FileList: []struct {
			FileId   string `json:"file_id"`
			FileName string `json:"file_name"`
			FileType string `json:"file_type"`
		}{
			{
				FileId:   "123",
				FileName: "123",
				FileType: "123",
			},
		},
		Address:    "",
		Mcc:        "Mcc",
		ProvId:     "",
		AreaId:     "",
		DistrictId: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", response)
}

func TestV2UserBusiOpen(t *testing.T) {
	response, raw, err := u.V2UserBusiOpen(user.V2UserBusiOpenRequest{
		ReqSeqId:     "",
		ReqDate:      "",
		HuifuId:      "",
		UpperHuifuId: "",
		SettleConfig: user.SettleConfig{
			SettleCycle:        "",
			MinAmt:             "",
			RemainedAmt:        "",
			SettleAbstract:     "",
			OutSettleFlag:      "",
			OutSettleHuifuid:   "",
			OutSettleAcctType:  "",
			SettlePattern:      "",
			SettleBatchNo:      "",
			IsPriorityReceipt:  "",
			SettleTime:         "",
			FixedRatio:         "",
			WorkdayFixedRatio:  "",
			WorkdayConstantAmt: "",
			ConstantAmt:        "",
		},
		CardInfo:   user.CardInfo{},
		CashConfig: nil,
		FileList: []struct {
			FileId   string `json:"file_id"`
			FileName string `json:"file_name"`
			FileType string `json:"file_type"`
		}{},
		DelayFlag:      "",
		ElecAcctConfig: user.ElecAcctConfig{},
		OpenTaxFlag:    "",
		AsyncReturnUrl: "",
		LgPlatformType: "",
		LjhData:        struct{ TaxAreaId string }{},
	})
	if err != nil {
		t.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", response)
}

func TestV2UserBusiModify(t *testing.T) {
	response, raw, err := u.V2UserBusiModify(user.V2UserBusiModifyRequest{
		ReqSeqId: "",
		ReqDate:  "",
	})
	if err != nil {
		t.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", response)
}
