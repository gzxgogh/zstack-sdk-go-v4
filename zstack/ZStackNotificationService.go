package zstack

import (
	"fmt"
	"github.com/gzxgogh/zstack-sdk-go-v4/errcode"
	"github.com/gzxgogh/zstack-sdk-go-v4/model"
	"github.com/gzxgogh/zstack-sdk-go-v4/request"
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/logs"
	"github.com/maczh/mgerr"
	"github.com/maczh/utils"
)

//创建主题
func CreateSNSTopic(params model.CreateSNSTopicRequest) mgresult.Result {
	//POST zstack/v1/sns/topics
	url := fmt.Sprintf("zstack/v1/sns/topics")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateSNSTopicFailed, err)
	}
	var result model.CreateSNSTopicResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除主题
func DeleteSNSTopic(params model.DeleteSNSTopicRequest) mgresult.Result {
	//DELETE zstack/v1/sns/topics/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/sns/topics/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteSNSTopicFailed, err)
	}
	var result model.DeleteSNSTopicResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更改主题状态
func ChangeSNSTopicState(params model.ChangeSNSTopicStateRequest) mgresult.Result {
	//PUT zstack/v1/zwatch/topics/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/topics/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeSNSTopicStateFailed, err)
	}
	var result model.ChangeSNSTopicStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新主题
func UpdateSNSTopic(params model.UpdateSNSTopicRequest) mgresult.Result {
	//PUT zstack/v1/sns/topics/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/sns/topics/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateSNSTopicFailed, err)
	}
	var result model.UpdateSNSTopicResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询主题
func QuerySNSTopic(params model.QuerySNSTopicRequest) mgresult.Result {
	//GET zstack/v1/sns/topics
	//GET zstack/v1/sns/topics/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/sns/topics")
	} else {
		url = fmt.Sprintf("zstack/v1/sns/topics/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QuerySNSTopicFailed, err)
	}
	var result model.QuerySNSTopicResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建邮件服务器
func CreateSNSEmailPlatform(params model.CreateSNSEmailPlatformRequest) mgresult.Result {
	//POST zstack/v1/sns/application-platforms/email
	url := fmt.Sprintf("zstack/v1/sns/application-platforms/email")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateSNSEmailPlatformFailed, err)
	}
	var result model.CreateSNSEmailPlatformResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//测试邮件服务器
func ValidateSNSEmailPlatform(params model.ValidateSNSEmailPlatformRequest) mgresult.Result {
	//PUT zstack/v1/sns/application-platforms/email/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/sns/application-platforms/email/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ValidateSNSEmailPlatformFailed, err)
	}
	var result model.ValidateSNSEmailPlatformResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除SNS应用平台
func DeleteSNSApplicationPlatform(params model.DeleteSNSApplicationPlatformRequest) mgresult.Result {
	//DELETE zstack/v1/sns/application-platforms/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/sns/application-platforms/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteSNSApplicationPlatformFailed, err)
	}
	var result model.DeleteSNSApplicationPlatformResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询邮件服务器
func QuerySNSEmailPlatform(params model.QuerySNSEmailPlatformRequest) mgresult.Result {
	//GET zstack/v1/sns/application-platforms/email
	//GET zstack/v1/sns/application-platforms/email/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/sns/application-platforms/email")
	} else {
		url = fmt.Sprintf("zstack/v1/sns/application-platforms/email/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QuerySNSEmailPlatformFailed, err)
	}
	var result model.QuerySNSEmailPlatformResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新SNS应用平台
func UpdateSNSApplicationPlatform(params model.UpdateSNSApplicationPlatformRequest) mgresult.Result {
	//PUT zstack/v1/sns/application-platforms/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/sns/application-platforms/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateSNSApplicationPlatformFailed, err)
	}
	var result model.UpdateSNSApplicationPlatformResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询SNS应用平台
func QuerySNSApplicationPlatform(params model.QuerySNSApplicationPlatformRequest) mgresult.Result {
	//GET zstack/v1/sns/application-platforms
	//GET zstack/v1/sns/application-platforms/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/sns/application-platforms")
	} else {
		url = fmt.Sprintf("zstack/v1/sns/application-platforms/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QuerySNSApplicationPlatformFailed, err)
	}
	var result model.QuerySNSApplicationPlatformResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更改SNS应用平台状态
func ChangeSNSApplicationPlatformState(params model.ChangeSNSApplicationPlatformStateRequest) mgresult.Result {
	//PUT zstack/v1/sns/application-platforms/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/sns/application-platforms/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeSNSApplicationPlatformStateFailed, err)
	}
	var result model.ChangeSNSApplicationPlatformStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建邮件接收端
func CreateSNSEmailEndpoint(params model.CreateSNSEmailEndpointRequest) mgresult.Result {
	//POST zstack/v1/sns/application-endpoints/emails
	url := fmt.Sprintf("zstack/v1/sns/application-endpoints/emails")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateSNSEmailEndpointFailed, err)
	}
	var result model.CreateSNSEmailEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询邮件接收端
func QuerySNSEmailEndpoint(params model.QuerySNSEmailEndpointRequest) mgresult.Result {
	//GET zstack/v1/sns/application-endpoints/emails
	//GET zstack/v1/sns/application-endpoints/emails/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/sns/application-endpoints/emails")
	} else {
		url = fmt.Sprintf("zstack/v1/sns/application-endpoints/emails/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QuerySNSEmailEndpointFailed, err)
	}
	var result model.QuerySNSEmailEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建HTTP接收端
func CreateSNSHttpEndpoint(params model.CreateSNSHttpEndpointRequest) mgresult.Result {
	//POST zstack/v1/sns/application-endpoints/http
	url := fmt.Sprintf("zstack/v1/sns/application-endpoints/http")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateSNSHttpEndpointFailed, err)
	}
	var result model.CreateSNSHttpEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询HTTP接收端
func QuerySNSHttpEndpoint(params model.QuerySNSHttpEndpointRequest) mgresult.Result {
	//GET zstack/v1/sns/application-endpoints/http
	//GET zstack/v1/sns/application-endpoints/http/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/sns/application-endpoints/http")
	} else {
		url = fmt.Sprintf("zstack/v1/sns/application-endpoints/http/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QuerySNSHttpEndpointFailed, err)
	}
	var result model.QuerySNSHttpEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建钉钉接收端
func CreateSNSDingTalkEndpoint(params model.CreateSNSDingTalkEndpointRequest) mgresult.Result {
	//POST zstack/v1/sns/application-endpoints/ding-talk
	url := fmt.Sprintf("zstack/v1/sns/application-endpoints/ding-talk")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateSNSDingTalkEndpointFailed, err)
	}
	var result model.CreateSNSDingTalkEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加钉钉用户
func AddSNSDingTalkAtPerson(params model.AddSNSDingTalkAtPersonRequest) mgresult.Result {
	//POST zstack/v1/sns/application-endpoints/ding-talk/at-persons
	url := fmt.Sprintf("zstack/v1/sns/application-endpoints/ding-talk/at-persons")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddSNSDingTalkAtPersonFailed, err)
	}
	var result model.AddSNSDingTalkAtPersonResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除钉钉用户
func RemoveSNSDingTalkAtPerson(params model.RemoveSNSDingTalkAtPersonRequest) mgresult.Result {
	//DELETE zstack/v1/sns/application-endpoints/ding-talk/{endpointUuid}/at-persons/{phoneNumber}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/sns/application-endpoints/ding-talk/%s/at-persons/%s", params.EndpointUuid, params.PhoneNumber)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveSNSDingTalkAtPersonFailed, err)
	}
	var result model.RemoveSNSDingTalkAtPersonResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询钉钉终端
func QuerySNSDingTalkEndpoint(params model.QuerySNSDingTalkEndpointRequest) mgresult.Result {
	//GET zstack/v1/sns/application-endpoints/ding-talk
	//GET zstack/v1/sns/application-endpoints/ding-talk/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/sns/application-endpoints/ding-talk")
	} else {
		url = fmt.Sprintf("zstack/v1/sns/application-endpoints/ding-talk/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QuerySNSDingTalkEndpointFailed, err)
	}
	var result model.QuerySNSDingTalkEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除接收端
func DeleteSNSApplicationEndpoint(params model.DeleteSNSApplicationEndpointRequest) mgresult.Result {
	//DELETE zstack/v1/sns/application-endpoints/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/sns/application-endpoints/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteSNSApplicationEndpointFailed, err)
	}
	var result model.DeleteSNSApplicationEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新接收端
func UpdateSNSApplicationEndpoint(params model.UpdateSNSApplicationEndpointRequest) mgresult.Result {
	//PUT zstack/v1/sns/application-endpoints/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/sns/application-endpoints/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateSNSApplicationEndpointFailed, err)
	}
	var result model.UpdateSNSApplicationEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询接收端
func QuerySNSApplicationEndpoint(params model.QuerySNSApplicationEndpointRequest) mgresult.Result {
	//GET zstack/v1/sns/application-endpoints
	//GET zstack/v1/sns/application-endpoints/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/sns/application-endpoints")
	} else {
		url = fmt.Sprintf("zstack/v1/sns/application-endpoints/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QuerySNSApplicationEndpointFailed, err)
	}
	var result model.QuerySNSApplicationEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更改接收端
func ChangeSNSApplicationEndpointState(params model.ChangeSNSApplicationEndpointStateRequest) mgresult.Result {
	//PUT zstack/v1/sns/application-endpoints/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/sns/application-endpoints/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeSNSApplicationEndpointStateFailed, err)
	}
	var result model.ChangeSNSApplicationEndpointStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建阿里云短信接收端
func CreateSNSAliyunSmsEndpoint(params model.CreateSNSAliyunSmsEndpointRequest) mgresult.Result {
	//POST zstack/v1/sns/sms-endpoints/aliyunsms
	url := fmt.Sprintf("zstack/v1/sns/sms-endpoints/aliyunsms")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateSNSAliyunSmsEndpointFailed, err)
	}
	var result model.CreateSNSAliyunSmsEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//验证阿里云短信接收端
func ValidateSNSAliyunSmsEndpoint(params model.ValidateSNSAliyunSmsEndpointRequest) mgresult.Result {
	//PUT zstack/v1/sns/sms-endpoints/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/sns/sms-endpoints/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ValidateSNSAliyunSmsEndpointFailed, err)
	}
	var result model.ValidateSNSAliyunSmsEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加短信接收者
func AddSNSSmsReceiver(params model.AddSNSSmsReceiverRequest) mgresult.Result {
	//POST zstack/v1/sns/sms-endpoints/receivers
	url := fmt.Sprintf("zstack/v1/sns/sms-endpoints/receivers")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddSNSSmsReceiverFailed, err)
	}
	var result model.AddSNSSmsReceiverResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除短信接收者
func RemoveSNSSmsReceiver(params model.RemoveSNSSmsReceiverRequest) mgresult.Result {
	//DELETE zstack/v1/sns/sms-endpoints/{endpointUuid}/receivers/{phoneNumber}
	url := fmt.Sprintf("zstack/v1/sns/sms-endpoints/%s/receivers/%s", params.EndpointUuid, params.PhoneNumber)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveSNSSmsReceiverFailed, err)
	}
	var result model.RemoveSNSSmsReceiverResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询短信接收端
func QuerySNSSmsEndpoint(params model.QuerySNSSmsEndpointRequest) mgresult.Result {
	//GET zstack/v1/sns/sms-endpoints
	//GET zstack/v1/sns/sms-endpoints/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/sns/sms-endpoints")
	} else {
		url = fmt.Sprintf("zstack/v1/sns/sms-endpoints/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QuerySNSSmsEndpointFailed, err)
	}
	var result model.QuerySNSSmsEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加邮箱地址到邮箱接收端
func AddEmailAddressToSNSEmailEndpoint(params model.AddEmailAddressToSNSEmailEndpointRequest) mgresult.Result {
	//POST zstack/v1/sns/application-endpoints/emails/email-addresses
	url := fmt.Sprintf("zstack/v1/sns/application-endpoints/emails/email-addresses")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddEmailAddressToSNSEmailEndpointFailed, err)
	}
	var result model.AddEmailAddressToSNSEmailEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新接收端邮箱地址
func UpdateEmailAddressOfSNSEmailEndpoint(params model.UpdateEmailAddressOfSNSEmailEndpointRequest) mgresult.Result {
	//PUT zstack/v1/sns/application-endpoints/emails/email-addresses
	url := fmt.Sprintf("zstack/v1/sns/application-endpoints/emails/email-addresses")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateEmailAddressOfSNSEmailEndpointFailed, err)
	}
	var result model.UpdateEmailAddressOfSNSEmailEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除邮箱接收端的地址
func DeleteEmailAddressOfSNSEmailEndpoint(params model.DeleteEmailAddressOfSNSEmailEndpointRequest) mgresult.Result {
	//DELETE zstack/v1/sns/application-endpoints/emails/{endpointUuid}/email-addresses/{emailAddressUuid}
	url := fmt.Sprintf("zstack/v1/sns/application-endpoints/emails/%s/email-addresses/%s", params.EndpointUuid, params.EmailAddressUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteEmailAddressOfSNSEmailEndpointFailed, err)
	}
	var result model.DeleteEmailAddressOfSNSEmailEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询接收端邮箱地址
func QuerySNSEmailAddress(params model.QuerySNSEmailAddressRequest) mgresult.Result {
	//GET zstack/v1/sns/application-endpoints/emails/email-addresses
	//GET zstack/v1/sns/application-endpoints/emails/email-addresses/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/sns/application-endpoints/emails/email-addresses")
	} else {
		url = fmt.Sprintf("zstack/v1/sns/application-endpoints/emails/email-addresses/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QuerySNSEmailAddressFailed, err)
	}
	var result model.QuerySNSEmailAddressResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建微软Teams接收端
func CreateSNSMicrosoftTeamsEndpointRequest(params model.CreateSNSMicrosoftTeamsEndpointRequest) mgresult.Result {
	//POST zstack/v1/sns/application-endpoints/microsoft-teams
	url := fmt.Sprintf("zstack/v1/sns/application-endpoints/microsoft-teams")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateSNSMicrosoftTeamsEndpointRequestFailed, err)
	}
	var result model.CreateSNSMicrosoftTeamsEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询微软Teams接收端
func QuerySNSMicrosoftTeamsEndpoint(params model.QuerySNSMicrosoftTeamsEndpointRequest) mgresult.Result {
	//GET zstack/v1/sns/application-endpoints/microsoft-teams
	//GET zstack/v1/sns/application-endpoints/microsoft-teams/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/sns/application-endpoints/microsoft-teams")
	} else {
		url = fmt.Sprintf("zstack/v1/sns/application-endpoints/microsoft-teams/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QuerySNSMicrosoftTeamsEndpointFailed, err)
	}
	var result model.QuerySNSMicrosoftTeamsEndpointResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//订阅主题
func SubscribeSNSTopic(params model.SubscribeSNSTopicRequest) mgresult.Result {
	//POST zstack/v1/sns/topics/{topicUuid}/endpoints/{endpointUuid}
	url := fmt.Sprintf("zstack/v1/sns/topics/%s/endpoints/%s", params.TopicUuid, params.EndpointUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SubscribeSNSTopicFailed, err)
	}
	var result model.SubscribeSNSTopicResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询订阅主题的终端
func QuerySNSTopicSubscriber(params model.QuerySNSTopicSubscriberRequest) mgresult.Result {
	//GET zstack/v1/sns/topics/subscribers
	//GET zstack/v1/sns/topics/subscribers/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/sns/topics/subscribers")
	} else {
		url = fmt.Sprintf("zstack/v1/sns/topics/subscribers/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QuerySNSTopicSubscriberFailed, err)
	}
	var result model.QuerySNSTopicSubscriberResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//退订主题
func UnsubscribeSNSTopic(params model.UnsubscribeSNSTopicRequest) mgresult.Result {
	//DELETE zstack/v1/sns/topics/{topicUuid}/endpoints/{endpointUuid}
	url := fmt.Sprintf("zstack/v1/sns/topics/%s/endpoints/%s", params.TopicUuid, params.EndpointUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UnsubscribeSNSTopicFailed, err)
	}
	var result model.UnsubscribeSNSTopicResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取管理节点目录容量
func GetManagementNodeDirCapacity(params model.GetManagementNodeDirCapacityRequest) mgresult.Result {
	//GET zstack/v1/zwatch/mn
	url := fmt.Sprintf("zstack/v1/zwatch/mn")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetManagementNodeDirCapacityFailed, err)
	}
	var result model.GetManagementNodeDirCapacityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
