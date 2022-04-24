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

//获取所有metric元数据
func GetAllMetricMetadata(params model.GetAllMetricMetadataRequest) mgresult.Result {
	//GET zstack/v1/zwatch/metrics/meta-data
	url := fmt.Sprintf("zstack/v1/zwatch/metrics/meta-data")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetAllMetricMetadataFailed, err)
	}
	var result model.GetAllMetricMetadataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取metric的标签值
func GetMetricLabelValue(params model.GetMetricLabelValueRequest) mgresult.Result {
	//GET zstack/v1/zwatch/metrics/label-values
	url := fmt.Sprintf("zstack/v1/zwatch/metrics/label-values")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetMetricLabelValueFailed, err)
	}
	var result model.GetMetricLabelValueResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取metric数据
func GetMetricData(params model.GetMetricDataRequest) mgresult.Result {
	//GET zstack/v1/zwatch/metrics
	url := fmt.Sprintf("zstack/v1/zwatch/metrics")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetMetricDataFailed, err)
	}
	var result model.GetMetricDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//存入自定义metric数据
func PutMetricData(params model.PutMetricDataRequest) mgresult.Result {
	//POST zstack/v1/zwatch/metrics
	url := fmt.Sprintf("zstack/v1/zwatch/metrics")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.PutMetricDataFailed, err)
	}
	var result model.PutMetricDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取所有event元数据
func GetAllEventMetadata(params model.GetAllEventMetadataRequest) mgresult.Result {
	//GET zstack/v1/zwatch/events/meta-data
	url := fmt.Sprintf("zstack/v1/zwatch/events/meta-data")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetAllEventMetadataFailed, err)
	}
	var result model.GetAllEventMetadataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取事件
func GetEventData(params model.GetEventDataRequest) mgresult.Result {
	//GET zstack/v1/zwatch/events
	url := fmt.Sprintf("zstack/v1/zwatch/events")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetEventDataFailed, err)
	}
	var result model.GetEventDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新事件
func UpdateEventData(params model.UpdateEventDataRequest) mgresult.Result {
	//PUT zstack/v1/zwatch/events/actions
	url := fmt.Sprintf("zstack/v1/zwatch/events/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateEventDataFailed, err)
	}
	var result model.UpdateEventDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询事件报警器报警记录
func QueryEventRecord(params model.QueryEventRecordRequest) mgresult.Result {
	//GET zstack/v1/zwatch/event-records
	url := fmt.Sprintf("zstack/v1/zwatch/event-records")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryEventRecordFailed, err)
	}
	var result model.QueryEventRecordResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建报警器
func CreateAlarm(params model.CreateAlarmRequest) mgresult.Result {
	//POST zstack/v1/zwatch/alarms
	url := fmt.Sprintf("zstack/v1/zwatch/alarms")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateAlarmFailed, err)
	}
	var result model.CreateAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除报警器
func DeleteAlarm(params model.DeleteAlarmRequest) mgresult.Result {
	//DELETE zstack/v1/zwatch/alarms/{uuid}
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteAlarmFailed, err)
	}
	var result model.DeleteAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//修改报警器
func UpdateAlarm(params model.UpdateAlarmRequest) mgresult.Result {
	//PUT zstack/v1/zwatch/alarms/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateAlarmFailed, err)
	}
	var result model.UpdateAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//修改事件报警器
func UpdateSubscribeEvent(params model.UpdateSubscribeEventRequest) mgresult.Result {
	//PUT zstack/v1/zwatch/events/subscriptions/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/events/subscriptions/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateSubscribeEventFailed, err)
	}
	var result model.UpdateSubscribeEventResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//增加报警动作
func AddActionToAlarm(params model.AddActionToAlarmRequest) mgresult.Result {
	//POST zstack/v1/zwatch/alarms/{alarmUuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/%s/actions", params.AlarmUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddActionToAlarmFailed, err)
	}
	var result model.AddActionToAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除报警动作
func RemoveActionFromAlarm(params model.RemoveActionFromAlarmRequest) mgresult.Result {
	//DELETE zstack/v1/zwatch/alarms/{alarmUuid}/actions/{actionUuid}
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/%s/actions/%s", params.AlarmUuid, params.ActionUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveActionFromAlarmFailed, err)
	}
	var result model.RemoveActionFromAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//增加标签到报警器
func AddLabelToAlarm(params model.AddLabelToAlarmRequest) mgresult.Result {
	//POST zstack/v1/zwatch/alarms/{alarmUuid}/labels
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/%s/labels", params.AlarmUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddLabelToAlarmFailed, err)
	}
	var result model.AddLabelToAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//修改报警器标签
func UpdateAlarmLabel(params model.UpdateAlarmLabelRequest) mgresult.Result {
	//PUT zstack/v1/zwatch/alarms/labels/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/labels/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateAlarmLabelFailed, err)
	}
	var result model.UpdateAlarmLabelResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从报警器上删除标签
func RemoveLabelFromAlarm(params model.RemoveLabelFromAlarmRequest) mgresult.Result {
	//DELETE zstack/v1/zwatch/alarms/labels/{uuid}
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/labels/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveLabelFromAlarmFailed, err)
	}
	var result model.RemoveLabelFromAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更改报警器状态
func ChangeAlarmState(params model.ChangeAlarmStateRequest) mgresult.Result {
	//PUT zstack/v1/zwatch/alarms/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeAlarmStateFailed, err)
	}
	var result model.ChangeAlarmStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询报警器
func QueryAlarm(params model.QueryAlarmRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryAlarmFailed, err)
	}
	var result model.QueryAlarmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取报警器历史
func GetAlarmData(params model.GetAlarmDataRequest) mgresult.Result {
	//GET zstack/v1/zwatch/alarm-histories
	url := fmt.Sprintf("zstack/v1/zwatch/alarm-histories")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetAlarmDataFailed, err)
	}
	var result model.GetAlarmDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新报警器历史消息
func UpdateAlarmData(params model.UpdateAlarmDataRequest) mgresult.Result {
	//PUT zstack/v1/zwatch/alarm-histories/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarm-histories/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateAlarmDataFailed, err)
	}
	var result model.UpdateAlarmDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//统计报警消息数量
func GetZWatchAlertHistogram(params model.GetZWatchAlertHistogramRequest) mgresult.Result {
	//GET zstack/v1/zwatch/alert-histories/histogram
	url := fmt.Sprintf("zstack/v1/zwatch/alert-histories/histogram")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetZWatchAlertHistogramFailed, err)
	}
	var result model.GetZWatchAlertHistogramResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询资源报警器报警记录
func QueryAlarmRecord(params model.QueryAlarmRecordRequest) mgresult.Result {
	//GET zstack/v1/zwatch/alarm-histories
	url := fmt.Sprintf("zstack/v1/zwatch/alarm-histories")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryAlarmRecordFailed, err)
	}
	var result model.QueryAlarmRecordResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//订阅事件
func SubscribeEvent(params model.SubscribeEventRequest) mgresult.Result {
	//POST zstack/v1/zwatch/events/subscriptions
	url := fmt.Sprintf("zstack/v1/zwatch/events/subscriptions")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SubscribeEventFailed, err)
	}
	var result model.SubscribeEventResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//退订事件
func UnsubscribeEvent(params model.UnsubscribeEventRequest) mgresult.Result {
	//DELETE zstack/v1/zwatch/events/subscriptions/{uuid}
	url := fmt.Sprintf("zstack/v1/zwatch/events/subscriptions/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UnsubscribeEventFailed, err)
	}
	var result model.UnsubscribeEventResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询事件订阅
func QueryEventSubscription(params model.QueryEventSubscriptionRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryEventSubscriptionFailed, err)
	}
	var result model.QueryEventSubscriptionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新事件订阅的标签
func UpdateEventSubscriptionLabel(params model.UpdateEventSubscriptionLabelRequest) mgresult.Result {
	//PUT zstack/v1/zwatch/events/subscriptions/labels/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/events/subscriptions/labels/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateEventSubscriptionLabelFailed, err)
	}
	var result model.UpdateEventSubscriptionLabelResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建报警消息模板
func CreateSNSTextTemplate(params model.CreateSNSTextTemplateRequest) mgresult.Result {
	//POST zstack/v1/zwatch/alarms/sns/text-templates
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateSNSTextTemplateFailed, err)
	}
	var result model.CreateSNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除报警消息模板
func DeleteSNSTextTemplate(params model.DeleteSNSTextTemplateRequest) mgresult.Result {
	//DELETE zstack/v1/zwatch/alarms/sns/text-templates/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteSNSTextTemplateFailed, err)
	}
	var result model.DeleteSNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新报警消息模板
func UpdateSNSTextTemplate(params model.UpdateSNSTextTemplateRequest) mgresult.Result {
	//PUT zstack/v1/zwatch/alarms/sns/text-templates/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateSNSTextTemplateFailed, err)
	}
	var result model.UpdateSNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询报警消息模板
func QuerySNSTextTemplate(params model.QuerySNSTextTemplateRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QuerySNSTextTemplateFailed, err)
	}
	var result model.QuerySNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建SNS监控阿里云短信模板
func CreateAliyunSmsSNSTextTemplate(params model.CreateAliyunSmsSNSTextTemplateRequest) mgresult.Result {
	//POST zstack/v1/zwatch/alarms/sns/text-templates/aliyun-sms
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates/aliyun-sms")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateAliyunSmsSNSTextTemplateFailed, err)
	}
	var result model.CreateAliyunSmsSNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新SNS阿里云短信文本模板
func UpdateAliyunSmsSNSTextTemplate(params model.UpdateAliyunSmsSNSTextTemplateRequest) mgresult.Result {
	//PUT zstack/v1/zwatch/alarms/sns/text-templates/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/sns/text-templates/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateAliyunSmsSNSTextTemplateFailed, err)
	}
	var result model.UpdateAliyunSmsSNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询SNS监控阿里云短信模板
func QueryAliyunSmsSNSTextTemplate(params model.QueryAliyunSmsSNSTextTemplateRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryAliyunSmsSNSTextTemplateFailed, err)
	}
	var result model.QueryAliyunSmsSNSTextTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取审计数据
func GetAuditData(params model.GetAuditDataRequest) mgresult.Result {
	//GET zstack/v1/zwatch/audits
	url := fmt.Sprintf("zstack/v1/zwatch/audits")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetAuditDataFailed, err)
	}
	var result model.GetAuditDataResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
