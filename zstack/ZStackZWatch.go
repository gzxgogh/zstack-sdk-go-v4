package zstack

import (
	"fmt"
	"github.com/gzxgogh/zstack-sdk-go-v4/errcode"
	"github.com/gzxgogh/zstack-sdk-go-v4/model"
	"github.com/gzxgogh/zstack-sdk-go-v4/request"
	"github.com/maczh/mgin/logs"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
)

//获取所有metric元数据
func GetAllMetricMetadata(params model.GetAllMetricMetadataRequest) models.Result {
	//GET zstack/v1/zwatch/metrics/meta-data
	url := fmt.Sprintf("zstack/v1/zwatch/metrics/meta-data")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetAllMetricMetadataFailed, err.Error())
	}
	var result model.GetAllMetricMetadataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取metric的标签值
func GetMetricLabelValue(params model.GetMetricLabelValueRequest) models.Result {
	//GET zstack/v1/zwatch/metrics/label-values
	url := fmt.Sprintf("zstack/v1/zwatch/metrics/label-values")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetMetricLabelValueFailed, err.Error())
	}
	var result model.GetMetricLabelValueResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取metric数据
func GetMetricData(params model.GetMetricDataRequest) models.Result {
	//GET zstack/v1/zwatch/metrics
	url := fmt.Sprintf("zstack/v1/zwatch/metrics")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetMetricDataFailed, err.Error())
	}
	var result model.GetMetricDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//存入自定义metric数据
func PutMetricData(params model.PutMetricDataRequest) models.Result {
	//POST zstack/v1/zwatch/metrics
	url := fmt.Sprintf("zstack/v1/zwatch/metrics")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.PutMetricDataFailed, err.Error())
	}
	var result model.PutMetricDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取所有event元数据
func GetAllEventMetadata(params model.GetAllEventMetadataRequest) models.Result {
	//GET zstack/v1/zwatch/events/meta-data
	url := fmt.Sprintf("zstack/v1/zwatch/events/meta-data")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetAllEventMetadataFailed, err.Error())
	}
	var result model.GetAllEventMetadataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取事件
func GetEventData(params model.GetEventDataRequest) models.Result {
	//GET zstack/v1/zwatch/events
	url := fmt.Sprintf("zstack/v1/zwatch/events")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetEventDataFailed, err.Error())
	}
	var result model.GetEventDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新事件
func UpdateEventData(params model.UpdateEventDataRequest) models.Result {
	//PUT zstack/v1/zwatch/events/actions
	url := fmt.Sprintf("zstack/v1/zwatch/events/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateEventDataFailed, err.Error())
	}
	var result model.UpdateEventDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询事件报警器报警记录
func QueryEventRecord(params model.QueryEventRecordRequest) models.Result {
	//GET zstack/v1/zwatch/event-records
	url := fmt.Sprintf("zstack/v1/zwatch/event-records")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryEventRecordFailed, err.Error())
	}
	var result model.QueryEventRecordResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建报警器
func CreateAlarm(params model.CreateAlarmRequest) models.Result {
	//POST zstack/v1/zwatch/alarms
	url := fmt.Sprintf("zstack/v1/zwatch/alarms")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateAlarmFailed, err.Error())
	}
	var result model.CreateAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除报警器
func DeleteAlarm(params model.DeleteAlarmRequest) models.Result {
	//DELETE zstack/v1/zwatch/alarms/{uuid}
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteAlarmFailed, err.Error())
	}
	var result model.DeleteAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改报警器
func UpdateAlarm(params model.UpdateAlarmRequest) models.Result {
	//PUT zstack/v1/zwatch/alarms/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateAlarmFailed, err.Error())
	}
	var result model.UpdateAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改事件报警器
func UpdateSubscribeEvent(params model.UpdateSubscribeEventRequest) models.Result {
	//PUT zstack/v1/zwatch/events/subscriptions/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/events/subscriptions/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateSubscribeEventFailed, err.Error())
	}
	var result model.UpdateSubscribeEventResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//增加报警动作
func AddActionToAlarm(params model.AddActionToAlarmRequest) models.Result {
	//POST zstack/v1/zwatch/alarms/{alarmUuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/%s/actions", params.AlarmUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.AddActionToAlarmFailed, err.Error())
	}
	var result model.AddActionToAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除报警动作
func RemoveActionFromAlarm(params model.RemoveActionFromAlarmRequest) models.Result {
	//DELETE zstack/v1/zwatch/alarms/{alarmUuid}/actions/{actionUuid}
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/%s/actions/%s", params.AlarmUuid, params.ActionUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveActionFromAlarmFailed, err.Error())
	}
	var result model.RemoveActionFromAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//增加标签到报警器
func AddLabelToAlarm(params model.AddLabelToAlarmRequest) models.Result {
	//POST zstack/v1/zwatch/alarms/{alarmUuid}/labels
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/%s/labels", params.AlarmUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddLabelToAlarmFailed, err.Error())
	}
	var result model.AddLabelToAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改报警器标签
func UpdateAlarmLabel(params model.UpdateAlarmLabelRequest) models.Result {
	//PUT zstack/v1/zwatch/alarms/labels/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/labels/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateAlarmLabelFailed, err.Error())
	}
	var result model.UpdateAlarmLabelResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从报警器上删除标签
func RemoveLabelFromAlarm(params model.RemoveLabelFromAlarmRequest) models.Result {
	//DELETE zstack/v1/zwatch/alarms/labels/{uuid}
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/labels/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveLabelFromAlarmFailed, err.Error())
	}
	var result model.RemoveLabelFromAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更改报警器状态
func ChangeAlarmState(params model.ChangeAlarmStateRequest) models.Result {
	//PUT zstack/v1/zwatch/alarms/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeAlarmStateFailed, err.Error())
	}
	var result model.ChangeAlarmStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询报警器
func QueryAlarm(params model.QueryAlarmRequest) models.Result {
	//GET zstack/v1/zwatch/alarms
	//GET zstack/v1/zwatch/alarms/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/zwatch/alarms")
	} else {
		url = fmt.Sprintf("zstack/v1/zwatch/alarms/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAlarmFailed, err.Error())
	}
	var result model.QueryAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取报警器历史
func GetAlarmData(params model.GetAlarmDataRequest) models.Result {
	//GET zstack/v1/zwatch/alarm-histories
	url := fmt.Sprintf("zstack/v1/zwatch/alarm-histories")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetAlarmDataFailed, err.Error())
	}
	var result model.GetAlarmDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新报警器历史消息
func UpdateAlarmData(params model.UpdateAlarmDataRequest) models.Result {
	//PUT zstack/v1/zwatch/alarm-histories/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarm-histories/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateAlarmDataFailed, err.Error())
	}
	var result model.UpdateAlarmDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//统计报警消息数量
func GetZWatchAlertHistogram(params model.GetZWatchAlertHistogramRequest) models.Result {
	//GET zstack/v1/zwatch/alert-histories/histogram
	url := fmt.Sprintf("zstack/v1/zwatch/alert-histories/histogram")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetZWatchAlertHistogramFailed, err.Error())
	}
	var result model.GetZWatchAlertHistogramResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询资源报警器报警记录
func QueryAlarmRecord(params model.QueryAlarmRecordRequest) models.Result {
	//GET zstack/v1/zwatch/alarm-histories
	url := fmt.Sprintf("zstack/v1/zwatch/alarm-histories")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAlarmRecordFailed, err.Error())
	}
	var result model.QueryAlarmRecordResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//订阅事件
func SubscribeEvent(params model.SubscribeEventRequest) models.Result {
	//POST zstack/v1/zwatch/events/subscriptions
	url := fmt.Sprintf("zstack/v1/zwatch/events/subscriptions")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.SubscribeEventFailed, err.Error())
	}
	var result model.SubscribeEventResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//退订事件
func UnsubscribeEvent(params model.UnsubscribeEventRequest) models.Result {
	//DELETE zstack/v1/zwatch/events/subscriptions/{uuid}
	url := fmt.Sprintf("zstack/v1/zwatch/events/subscriptions/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.UnsubscribeEventFailed, err.Error())
	}
	var result model.UnsubscribeEventResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询事件订阅
func QueryEventSubscription(params model.QueryEventSubscriptionRequest) models.Result {
	//GET zstack/v1/zwatch/events/subscriptions
	//GET zstack/v1/zwatch/events/subscriptions/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/zwatch/events/subscriptions")
	} else {
		url = fmt.Sprintf("zstack/v1/zwatch/events/subscriptions/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryEventSubscriptionFailed, err.Error())
	}
	var result model.QueryEventSubscriptionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新事件订阅的标签
func UpdateEventSubscriptionLabel(params model.UpdateEventSubscriptionLabelRequest) models.Result {
	//PUT zstack/v1/zwatch/events/subscriptions/labels/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/events/subscriptions/labels/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateEventSubscriptionLabelFailed, err.Error())
	}
	var result model.UpdateEventSubscriptionLabelResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建报警消息模板
func CreateSNSTextTemplate(params model.CreateSNSTextTemplateRequest) models.Result {
	//POST zstack/v1/zwatch/alarms/sns/text-templates
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateSNSTextTemplateFailed, err.Error())
	}
	var result model.CreateSNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除报警消息模板
func DeleteSNSTextTemplate(params model.DeleteSNSTextTemplateRequest) models.Result {
	//DELETE zstack/v1/zwatch/alarms/sns/text-templates/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteSNSTextTemplateFailed, err.Error())
	}
	var result model.DeleteSNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新报警消息模板
func UpdateSNSTextTemplate(params model.UpdateSNSTextTemplateRequest) models.Result {
	//PUT zstack/v1/zwatch/alarms/sns/text-templates/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateSNSTextTemplateFailed, err.Error())
	}
	var result model.UpdateSNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询报警消息模板
func QuerySNSTextTemplate(params model.QuerySNSTextTemplateRequest) models.Result {
	//GET zstack/v1/zwatch/alarms/sns/text-templates
	//GET zstack/v1/zwatch/alarms/sns/text-templates/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates")
	} else {
		url = fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QuerySNSTextTemplateFailed, err.Error())
	}
	var result model.QuerySNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建SNS监控阿里云短信模板
func CreateAliyunSmsSNSTextTemplate(params model.CreateAliyunSmsSNSTextTemplateRequest) models.Result {
	//POST zstack/v1/zwatch/alarms/sns/text-templates/aliyun-sms
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates/aliyun-sms")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateAliyunSmsSNSTextTemplateFailed, err.Error())
	}
	var result model.CreateAliyunSmsSNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新SNS阿里云短信文本模板
func UpdateAliyunSmsSNSTextTemplate(params model.UpdateAliyunSmsSNSTextTemplateRequest) models.Result {
	//PUT zstack/v1/zwatch/alarms/sns/text-templates/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateAliyunSmsSNSTextTemplateFailed, err.Error())
	}
	var result model.UpdateAliyunSmsSNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询SNS监控阿里云短信模板
func QueryAliyunSmsSNSTextTemplate(params model.QueryAliyunSmsSNSTextTemplateRequest) models.Result {
	//GET zstack/v1/zwatch/alarms/sns/text-templates/aliyun-sms
	//GET zstack/v1/zwatch/alarms/sns/text-templates/aliyun-sms/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates/aliyun-sms")
	} else {
		url = fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates/aliyun-sms/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAliyunSmsSNSTextTemplateFailed, err.Error())
	}
	var result model.QueryAliyunSmsSNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取审计数据
func GetAuditData(params model.GetAuditDataRequest) models.Result {
	//GET zstack/v1/zwatch/audits
	url := fmt.Sprintf("zstack/v1/zwatch/audits")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetAuditDataFailed, err.Error())
	}
	var result model.GetAuditDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
