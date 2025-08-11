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
	response, raw, err := u.V2UserBasicdataIndv(user.V2UserBasicdataIndvRequest{
		ReqSeqId:         tool.GetReqSeqId(),
		ReqDate:          tool.GetCurrentDate(),
		Name:             "张三",
		CertType:         common.CertType00,
		CertNo:           "",
		CertValidityType: "",
		CertNationality:  "",
		CertBeginDate:    "",
		MobileNo:         "",
		Address:          "",
		ExtendInfos: user.V2UserBasicdataIndvExtendInfos{
			CertEndDate: "",
			Email:       "",
			LoginName:   "",
			SmsSendFlag: "",
			ExpandId:    "",
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
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", response)
}
func TestV2UserBasicdataIndvModify(t *testing.T) {
	response, raw, err := u.V2UserBasicdataIndvModify(user.V2UserBasicdataIndvModifyRequest{
		ReqSeqId: tool.GetReqSeqId(),
		ReqDate:  tool.GetCurrentDate(),
		HuifuId:  u.HuifuPay.BsPay.Msc.SysId,
		ExtendInfos: user.V2UserBasicdataIndvModifyExtendInfos{
			CertValidityType: "",
			CertBeginDate:    "",
			CertEndDate:      "",
			Email:            "",
			MobileNo:         "",
			FileList: []struct {
				FileId   string
				FileName string
				FileType string
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
		},
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
		ExtendInfos: user.V2UserBusiOpenExtendInfos{
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
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", response)
}

func TestV2UserBusiModify(t *testing.T) {
	response, raw, err := u.V2UserBusiModify(user.V2UserBusiModifyRequest{
		HuifuId:      "",
		ReqSeqId:     tool.GetReqSeqId(),
		ReqDate:      tool.GetCurrentDate(),
		UpperHuifuId: u.HuifuPay.BsPay.Msc.SysId,
		ExtendInfos:  user.V2UserBusiModifyExtendInfos{},
	})
	if err != nil {
		t.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", response)
}

func TestV2SupplementaryPicture(t *testing.T) {
	response, raw, err := u.V2SupplementaryPicture(user.V2SupplementaryPictureRequest{
		ReqSeqId: tool.GetReqSeqId(),
		ReqDate:  tool.GetCurrentDate(),
		HuifuId:  u.HuifuPay.BsPay.Msc.SysId,
		FileType: common.FileTypeF01,
		FileUrl:  "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png",
	})
	if err != nil {
		t.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", response)
}
