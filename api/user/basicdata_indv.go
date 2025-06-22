package user

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
	"github.com/sleep-go/huifu-pay/common"
)

type User struct {
	HuifuPay *common.HuifuPay
}

func NewUser(huifuPay *common.HuifuPay) *User {
	return &User{HuifuPay: huifuPay}
}

// V2UserBasicdataIndv 个人用户基本信息开户
// POST https://api.huifu.com/v2/user/basicdata/indv
//
// 应用场景
// 服务商或商户为旗下的个人用户开户，开户后该用户可使用多方分账及结算功能。
//
// 适用对象
// 需要为旗下个人用户开户的服务商或商户
func (u *User) V2UserBasicdataIndv(req V2UserBasicdataIndvRequest) (res *V2UserBasicdataIndvResponse, raw string, err error) {
	resp, err := u.HuifuPay.BsPay.V2UserBasicdataIndvRequest(BsPaySdk.V2UserBasicdataIndvRequest{
		ReqSeqId:         tool.GetReqSeqId(),
		ReqDate:          tool.GetCurrentDate(),
		Name:             req.Name,
		CertType:         string(req.CertType),
		CertNo:           req.CertNo,
		CertValidityType: req.CertValidityType,
		CertBeginDate:    req.CertBeginDate,
		CertNationality:  req.CertNationality,
		MobileNo:         req.MobileNo,
		Address:          req.Address,
		ExtendInfos: map[string]interface{}{
			"cert_end_date": req.CertEndDate,
			"mobile_no":     req.MobileNo,
			"email":         req.Email,
			"login_name":    req.LoginName,
			"sms_send_flag": req.SmsSendFlag,
			"expand_id":     req.ExpandId,
			"file_list":     req.FileList,
			"mcc":           req.Mcc,
			"prov_id":       req.ProvId,
			"area_id":       req.AreaId,
			"district_id":   req.DistrictId,
		},
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2UserBasicdataIndvResponse](resp)
}

type V2UserBasicdataIndvRequest struct {
	ReqSeqId         string          `json:"req_seq_id"`
	ReqDate          string          `json:"req_date"`
	Name             string          `json:"name"`               //个人姓名	String	32	Y	示例值：张三
	CertType         common.CertType `json:"cert_type"`          //个人证件类型	String	2	Y	参见《自然人证件类型》说明；示例值：00
	CertNo           string          `json:"cert_no"`            //个人证件号码	String	32	Y	示例值：320926198312024023
	CertValidityType string          `json:"cert_validity_type"` //个人证件有效期类型	String	1	Y	1:长期有效 0:非长期有效；示例值：0
	CertBeginDate    string          `json:"cert_begin_date"`    //个人证件有效期开始日期	String	8	Y	日期格式：yyyyMMdd；示例值：20220909
	CertEndDate      string          `json:"cert_end_date"`      //个人证件有效期截止日期	String	8	N	日期格式：yyyyMMdd; 示例值：20330909 ；长期有效时可不填，非长期有效必填
	CertNationality  string          `json:"cert_nationality"`   //个人国籍	String	50	C	个人证件类型为外国人居留证时，必填，参见《国籍编码》示例值：CHN
	MobileNo         string          `json:"mobile_no"`          //手机号	String	11	Y	示例值：13917354627
	Email            string          `json:"email"`              //电子邮箱	String	64	N	示例值：carl.chen@huifu.com
	LoginName        string          `json:"login_name"`         //管理员账号	String	32	N	示例值：Lg2022022201394910571
	SmsSendFlag      string          `json:"sms_send_flag"`      //是否发送短信标识	String	1	N	Y:发送短信通知，N：不发送短信通知。默认不发送短信通知。示例值：Y
	ExpandId         string          `json:"expand_id"`          //拓展方字段	String	18	N	如果该商户是第三方展业的可以填写拓展方的huifu_id;示例值：6666000123123123
	FileList         []struct {
		FileId   string `json:"file_id"`   //文件jfileID	String	128	Y	图片上传接口生成的fileId；示例值：57cc7f00-600a-33ab-b614-6221bbf2e529
		FileName string `json:"file_name"` //文件名称	String	128	N	示例值：test42001.jpg
		FileType string `json:"file_type"` //文件类型	String	8	Y	参见文件类型；示例值：F85
	} `json:"file_list"` //文件列表
	Address    string `json:"address"`     //地址	String	256	C	开通中信E管家必填
	Mcc        string `json:"mcc"`         //	所属行业	String	7	N	参考汇付MCC编码 ；示例值：5311； 当用户业务入驻修改，电子回单配置开关为开通时，需填写
	ProvId     string `json:"prov_id"`     //省	String	6	N	参考地区编码；示例值：310000 ；如修改省市区要级联修改； 当用户业务入驻修改，电子回单配置开关为开通时，需填写
	AreaId     string `json:"area_id"`     //市	String	6	N	参考地区编码；示例值：310000 ；如修改省市区要级联修改； 当用户业务入驻修改，电子回单配置开关为开通时，需填写
	DistrictId string `json:"district_id"` //区	String	6	N	参考地区编码；示例值：310101 ；如修改省市区要级联修改； 当用户业务入驻修改，电子回单配置开关为开通时，需填写
}
type V2UserBasicdataIndvResponse struct {
	Data struct {
		RespDesc      string `json:"resp_desc"`
		RespCode      string `json:"resp_code"`
		HuifuId       string `json:"huifu_id"`
		LoginName     string `json:"login_name"`
		LoginPassword string `json:"login_password"`
	} `json:"data"`
	Sign string
}

// V2UserBasicdataIndvModify 个人用户基本信息修改
// POST https://api.huifu.com/v2/user/basicdata/indv/modify
//
// 应用场景
// 使用本接口可完成个人用户信息修改，可用于维护分账用户，快捷交易用户等。
//
// 适用对象
// 渠道商或商户开通的个人用户
func (u *User) V2UserBasicdataIndvModify(req V2UserBasicdataIndvModifyRequest) (res *V2UserBasicdataIndvModifyResponse, raw string, err error) {
	resp, err := u.HuifuPay.BsPay.V2UserBasicdataIndvModifyRequest(BsPaySdk.V2UserBasicdataIndvModifyRequest{
		ReqSeqId: req.ReqSeqId,
		ReqDate:  req.ReqDate,
		HuifuId:  req.HuifuId,
		ExtendInfos: map[string]any{
			"cert_validity_type": req.CertValidityType,
			"cert_begin_date":    req.CertBeginDate,
			"cert_end_date":      req.CertEndDate,
			"email":              req.Email,
			"mobile_no":          req.MobileNo,
			"file_list":          req.FileList,
			"address":            req.Address,
			"mcc":                req.Mcc,
			"prov_id":            req.ProvId,
			"district_id":        req.DistrictId,
		},
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2UserBasicdataIndvModifyResponse](resp)
}

type V2UserBasicdataIndvModifyRequest struct {
	ReqSeqId         string `json:"req_seq_id"`
	ReqDate          string `json:"req_date"`
	HuifuId          string `json:"huifu_id"`
	CertValidityType string `json:"cert_validity_type"` //个人证件有效期类型	String	1	Y	1:长期有效 0:非长期有效；示例值：0
	CertBeginDate    string `json:"cert_begin_date"`    //个人证件有效期开始日期	String	8	Y	日期格式：yyyyMMdd；示例值：20220909
	CertEndDate      string `json:"cert_end_date"`      //个人证件有效期截止日期	String	8	N	日期格式：yyyyMMdd; 示例值：20330909 ；长期有效时可不填，非长期有效必填
	Email            string `json:"email"`              //电子邮箱	String	64	N	示例值：carl.chen@huifu.com
	MobileNo         string `json:"mobile_no"`          //手机号	String	11	N	手机号，11位数字 示例值：13917145190
	FileList         []struct {
		FileId   string `json:"file_id"`   //文件jfileID	String	128	Y	图片上传接口生成的fileId；示例值：57cc7f00-600a-33ab-b614-6221bbf2e529
		FileName string `json:"file_name"` //文件名称	String	128	N	示例值：test42001.jpg
		FileType string `json:"file_type"` //文件类型	String	8	Y	参见文件类型；示例值：F85
	} `json:"file_list"` //文件列表
	Address    string `json:"address"`     //地址	String	256	C	开通中信E管家必填
	Mcc        string `json:"mcc"`         //	所属行业	String	7	N	参考汇付MCC编码 ；示例值：5311； 当用户业务入驻修改，电子回单配置开关为开通时，需填写
	ProvId     string `json:"prov_id"`     //省	String	6	N	参考地区编码；示例值：310000 ；如修改省市区要级联修改； 当用户业务入驻修改，电子回单配置开关为开通时，需填写
	AreaId     string `json:"area_id"`     //市	String	6	N	参考地区编码；示例值：310000 ；如修改省市区要级联修改； 当用户业务入驻修改，电子回单配置开关为开通时，需填写
	DistrictId string `json:"district_id"` //区	String	6	N	参考地区编码；示例值：310101 ；如修改省市区要级联修改； 当用户业务入驻修改，电子回单配置开关为开通时，需填写
}
type V2UserBasicdataIndvModifyResponse struct {
	Data struct {
		RespDesc string `json:"resp_desc"`
		RespCode string `json:"resp_code"`
		HuifuId  string `json:"huifu_id"`
	} `json:"data"`
	Sign string `json:"sign"`
}
