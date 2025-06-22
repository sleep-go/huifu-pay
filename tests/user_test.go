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
		ExtendInfos: struct {
			CertEndDate string //个人证件有效期截止日期	String	8	N	日期格式：yyyyMMdd; 示例值：20330909 ；长期有效时可不填，非长期有效必填
			Email       string //电子邮箱	String	64	N	示例值：carl.chen@huifu.com
			LoginName   string //管理员账号	String	32	N	示例值：Lg2022022201394910571
			SmsSendFlag string //是否发送短信标识	String	1	N	Y:发送短信通知，N：不发送短信通知。默认不发送短信通知。示例值：Y
			ExpandId    string //拓展方字段	String	18	N	如果该商户是第三方展业的可以填写拓展方的huifu_id;示例值：6666000123123123
			FileList    []struct {
				FileId   string //文件jfileID	String	128	Y	图片上传接口生成的fileId；示例值：57cc7f00-600a-33ab-b614-6221bbf2e529
				FileName string //文件名称	String	128	N	示例值：test42001.jpg
				FileType string //文件类型	String	8	Y	参见文件类型；示例值：F85
			} //文件列表
			Mcc        string //	所属行业	String	7	N	参考汇付MCC编码 ；示例值：5311； 当用户业务入驻修改，电子回单配置开关为开通时，需填写
			ProvId     string //省	String	6	N	参考地区编码；示例值：310000 ；如修改省市区要级联修改； 当用户业务入驻修改，电子回单配置开关为开通时，需填写
			AreaId     string //市	String	6	N	参考地区编码；示例值：310000 ；如修改省市区要级联修改； 当用户业务入驻修改，电子回单配置开关为开通时，需填写
			DistrictId string //区	String	6	N	参考地区编码；示例值：310101 ；如修改省市区要级联修改； 当用户业务入驻修改，电子回单配置开关为开通时，需填写
		}{
			CertEndDate: "",
			Email:       "",
			LoginName:   "",
			SmsSendFlag: "",
			ExpandId:    "",
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
			LjhData:        struct{ TaxAreaId string }{},
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
