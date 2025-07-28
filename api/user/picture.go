package user

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2SupplementaryPicture 图片上传
// POST https://api.huifu.com/v2/supplementary/picture
//
// 应用场景
// 提供商户图片上传服务，返回的图片文件ID可用于商户进件、业务开通、变更等接口。
//
// 适用对象
// 调用商户进件、业务开通或其他需要上传图片时；
//
// 接口说明
// 请求方式：该接口与其他接口请求格式不同 ，该接口使用POST表单格式（multipart/form-data）
// 支持格式：JSON
// 加签验签说明：图片上传没有加签，不需要验签；
func (u *User) V2SupplementaryPicture(req V2SupplementaryPictureRequest) (response *V2SupplementaryPictureResponse, raw string, err error) {
	// 请求参数 data 字段（嵌套 JSON）
	dataMap := BsPaySdk.ToMap(req)
	var url = BsPaySdk.BASE_API_TEST_URL_V2
	if u.HuifuPay.BsPay.IsProdMode {
		url = BsPaySdk.BASE_API_URL_V2
	}
	reqUrl := url + BsPaySdk.V2_SUPPLEMENTARY_PICTURE
	resp, err := u.doUploadByFileUrl(reqUrl, dataMap)
	if err != nil {
		return nil, "", err
	}
	var res map[string]any
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2SupplementaryPictureResponse](res)
}

type V2SupplementaryPictureRequest struct {
	ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 业务请求流水号
	ReqDate  string `json:"req_date" structs:"req_date"`     // 业务请求日期
	HuifuId  string `json:"huifu_id" structs:"huifu_id"`
	FileType string `json:"file_type" structs:"file_type"` // 图片类型
	FileUrl  string `json:"file_url" structs:"file_url"`   // 图片http地址
}
type V2SupplementaryPictureResponse struct {
	Data struct {
		RespCode string `json:"resp_code"`
		RespDesc string `json:"resp_desc"`
		FileId   string `json:"file_id"`
	} `json:"data"`
}

// doUploadByFileUrl 上传远程非本地文件方法
func (u *User) doUploadByFileUrl(url string, params map[string]interface{}) (*http.Response, error) {
	// 构建参数字符串
	signMap := BsPaySdk.DeepCopy(BsPaySdk.DeleteEmptyValue(params)).(map[string]interface{})
	delete(signMap, "extend_infos")
	if params["extend_infos"] != nil {
		signMap = BsPaySdk.AddMapValue(params["extend_infos"].(map[string]interface{}), signMap)
	}
	paramText, _ := BsPaySdk.FormatSignSrcText(signMap)
	sign, _ := BsPaySdk.RsaSign(paramText, u.HuifuPay.BsPay.Msc)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("data", paramText)
	respData := make(map[string]interface{})
	respData["sign"] = sign
	respData["sys_id"] = u.HuifuPay.BsPay.Msc.SysId
	respData["product_id"] = u.HuifuPay.BsPay.Msc.ProductId
	respData["data"] = signMap
	for key, val := range respData {
		if str, ok := val.(string); ok {
			_ = writer.WriteField(key, str)
		}
	}
	writerErr := writer.Close()
	if writerErr != nil {
		BsPaySdk.BspayPrintln("writer file error: " + writerErr.Error())
		return nil, writerErr
	}
	request, _ := http.NewRequest("POST", url, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return http.DefaultClient.Do(request)
}
