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

//创建弹性伸缩组
func CreateAutoScalingGroup(params model.CreateAutoScalingGroupRequest) models.Result {
	//POST zstack/v1/autoscaling/groups
	url := fmt.Sprintf("zstack/v1/autoscaling/groups")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateAutoScalingGroupFailed, err.Error())
	}
	var result model.CreateAutoScalingGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除弹性伸缩组
func DeleteAutoScalingGroup(params model.DeleteAutoScalingGroupRequest) models.Result {
	//DELETE zstack/v1/autoscaling/groups/{uuid}
	url := fmt.Sprintf("zstack/v1/autoscaling/groups/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteAutoScalingGroupFailed, err.Error())
	}
	var result model.DeleteAutoScalingGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改弹性伸缩组
func UpdateAutoScalingGroup(params model.UpdateAutoScalingGroupRequest) models.Result {
	//PUT zstack/v1/autoscaling/groups/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/autoscaling/groups/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateAutoScalingGroupFailed, err.Error())
	}
	var result model.UpdateAutoScalingGroupParams
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询弹性伸缩组
func QueryAutoScalingGroup(params model.QueryAutoScalingGroupRequest) models.Result {
	//GET zstack/v1/autoscaling/groups
	//GET zstack/v1/autoscaling/groups/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/autoscaling/groups")
	} else {
		url = fmt.Sprintf("zstack/v1/autoscaling/groups/%s", params.UUID)

	}
	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAutoScalingGroupFailed, err.Error())
	}
	var result model.QueryAutoScalingGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询弹性伸缩组活动列表
func QueryAutoScalingGroupActivity(params model.QueryAutoScalingGroupActivityRequest) models.Result {
	//GET zstack/v1/autoscaling/groups/activities
	//GET zstack/v1/autoscaling/groups/activities/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/autoscaling/groups/activities")
	} else {
		url = fmt.Sprintf("zstack/v1/autoscaling/groups/activities/%s", params.UUID)

	}
	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAutoScalingGroupActivityFailed, err.Error())
	}
	var result model.QueryAutoScalingGroupActivityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建伸缩组云主机扩容规则
func CreateAutoScalingGroupAddingNewInstanceRule(params model.CreateAutoScalingGroupAddingNewInstanceRuleRequest) models.Result {
	//POST zstack/v1/autoscaling/rules/adding-new-instance
	url := fmt.Sprintf("zstack/v1/autoscaling/rules/adding-new-instance")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateAutoScalingGroupAddingNewInstanceRuleFailed, err.Error())
	}
	var result model.CreateAutoScalingGroupAddingNewInstanceRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建伸缩组云主机缩容规则
func CreateAutoScalingGroupRemovalInstanceRule(params model.CreateAutoScalingGroupRemovalInstanceRuleRequest) models.Result {
	//POST zstack/v1/autoscaling/rules/removal-instance
	url := fmt.Sprintf("zstack/v1/autoscaling/rules/removal-instance")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateAutoScalingGroupRemovalInstanceRuleFailed, err.Error())
	}
	var result model.CreateAutoScalingGroupRemovalInstanceRuleParams
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除伸缩规则
func DeleteAutoScalingRule(params model.DeleteAutoScalingRuleRequest) models.Result {
	//DELETE zstack/v1/autoscaling/rules/{uuid}
	url := fmt.Sprintf("zstack/v1/autoscaling/rules/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteAutoScalingRuleFailed, err.Error())
	}
	var result model.DeleteAutoScalingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改弹性伸缩组规则
func UpdateAutoScalingRule(params model.UpdateAutoScalingRuleRequest) models.Result {
	//PUT zstack/v1/autoscaling/rules/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/autoscaling/rules/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateAutoScalingRuleFailed, err.Error())
	}
	var result model.UpdateAutoScalingRuleParams
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改伸缩组扩容规则
func UpdateAutoScalingGroupAddingNewInstanceRule(params model.UpdateAutoScalingGroupAddingNewInstanceRuleRequest) models.Result {
	//PUT zstack/v1/autoscaling/rules/adding-new-instance/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/autoscaling/rules/adding-new-instance/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateAutoScalingGroupAddingNewInstanceRuleFailed, err.Error())
	}
	var result model.UpdateAutoScalingGroupAddingNewInstanceRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改伸缩组缩容规则
func UpdateAutoScalingGroupRemovalInstanceRule(params model.UpdateAutoScalingGroupRemovalInstanceRuleRequest) models.Result {
	//PUT zstack/v1/autoscaling/rules/removal-instance/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/autoscaling/rules/removal-instance/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateAutoScalingGroupRemovalInstanceRuleFailed, err.Error())
	}
	var result model.UpdateAutoScalingGroupRemovalInstanceRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询伸缩规则
func QueryAutoScalingRule(params model.QueryAutoScalingRuleRequest) models.Result {
	//GET zstack/v1/autoscaling/groups/rules
	//GET zstack/v1/autoscaling/groups/rules/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/autoscaling/groups/rules")
	} else {
		url = fmt.Sprintf("zstack/v1/autoscaling/groups/rules/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAutoScalingRuleFailed, err.Error())
	}
	var result model.QueryAutoScalingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建伸缩规则触发器
func CreateAutoScalingRuleAlarmTrigger(params model.CreateAutoScalingRuleAlarmTriggerRequest) models.Result {
	//POST zstack/v1/zwatch/alarms/{alarmUuid}/autoscaling/rules/{ruleUuid}
	url := fmt.Sprintf("zstack/v1/zwatch/alarms/%s/autoscaling/rules/%s", params.AlarmUuid, params.RuleUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateAutoScalingRuleAlarmTriggerFailed, err.Error())
	}
	var result model.CreateAutoScalingRuleAlarmTriggerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建伸缩规则定时任务触发器
func CreateAutoScalingRuleSchedulerJobTrigger(params model.CreateAutoScalingRuleSchedulerJobTriggeRequest) models.Result {
	//POST zstack/v1/scheduler/jobs/{schedulerJobUuid}/autoscaling/rules/{ruleUuid}
	url := fmt.Sprintf("zstack/v1/scheduler/jobs/%s/autoscaling/rules/%s", params.SchedulerJobUuid, params.RuleUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateAutoScalingRuleSchedulerJobTriggerFailed, err.Error())
	}
	var result model.CreateAutoScalingRuleSchedulerJobTriggerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除伸缩规则触发器
func DeleteAutoScalingRuleTrigger(params model.DeleteAutoScalingRuleTriggerRequest) models.Result {
	//DELETE zstack/v1/autoscaling/groups/rules/triggers/{uuid}
	url := fmt.Sprintf("zstack/v1/autoscaling/groups/rules/triggers/%s", params.Uuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteAutoScalingRuleTriggerFailed, err.Error())
	}
	var result model.DeleteAutoScalingRuleTriggerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询伸缩规则触发器列表
func QueryAutoScalingRuleTrigger(params model.QueryAutoScalingRuleTriggerRequest) models.Result {
	//GET zstack/v1/autoscaling/groups/rules/trigger
	//GET zstack/v1/autoscaling/groups/rules/trigger/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/autoscaling/groups/rules/trigger")
	} else {
		url = fmt.Sprintf("zstack/v1/autoscaling/groups/rules/trigger/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAutoScalingRuleTriggerFailed, err.Error())
	}
	var result model.QueryAutoScalingRuleTriggerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建伸缩组云主机模块
func CreateAutoScalingVmTemplate(params model.CreateAutoScalingVmTemplateRequest) models.Result {
	//POST zstack/v1/autoscaling/vmtemplate
	url := fmt.Sprintf("zstack/v1/autoscaling/vmtemplate")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateAutoScalingVmTemplateFailed, err.Error())
	}
	var result model.CreateAutoScalingVmTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//将云主机模块添加到弹性伸缩组
func AttachAutoScalingTemplateToGroup(params model.AttachAutoScalingTemplateToGroupRequest) models.Result {
	//POST zstack/v1/autoscaling/template/{uuid}/groups/{groupUuid}
	url := fmt.Sprintf("zstack/v1/autoscaling/template/%s/groups/%s", params.UUID, params.GroupUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachAutoScalingTemplateToGroupFailed, err.Error())
	}
	var result model.AttachAutoScalingTemplateToGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除伸缩组模板
func DeleteAutoScalingTemplate(params model.DeleteAutoScalingTemplateRequest) models.Result {
	//DELETE zstack/v1/autoscaling/template/{uuid}
	url := fmt.Sprintf("zstack/v1/autoscaling/template/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteAutoScalingTemplateFailed, err.Error())
	}
	var result model.DeleteAutoScalingTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//卸载伸缩组模板
func DetachAutoScalingTemplateFromGroup(params model.DetachAutoScalingTemplateFromGroupRequest) models.Result {
	//DELETE zstack/v1/autoscaling/template/{uuid}/groups/{groupUuid}
	url := fmt.Sprintf("zstack/v1/autoscaling/template/%s/groups/%s", params.TemplateUuid, params.GroupUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachAutoScalingTemplateFromGroupFailed, err.Error())
	}
	var result model.DetachAutoScalingTemplateFromGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询伸缩组云主机模板
func QueryAutoScalingVmTemplate(params model.QueryAutoScalingVmTemplateRequest) models.Result {
	//GET zstack/v1/autoscaling/vmtemplate
	//GET zstack/v1/autoscaling/vmtemplate/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/autoscaling/vmtemplate")
	} else {
		url = fmt.Sprintf("zstack/v1/autoscaling/vmtemplate/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAutoScalingVmTemplateFailed, err.Error())
	}
	var result model.QueryAutoScalingVmTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//手动删除弹性伸缩组内云主机
func DeleteAutoScalingGroupInstance(params model.DeleteAutoScalingGroupInstanceRequest) models.Result {
	//DELETE zstack/v1/autoscaling/groups/instances/{instanceUuid}
	url := fmt.Sprintf("zstack/v1/autoscaling/groups/instances/%s", params.InstanceUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteAutoScalingGroupInstanceFailed, err.Error())
	}
	var result model.DeleteAutoScalingGroupInstanceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询弹性伸缩组内云主机列表
func QueryAutoScalingGroupInstance(params model.QueryAutoScalingGroupInstanceRequest) models.Result {
	//GET zstack/v1/autoscaling/groups/instances
	//GET zstack/v1/autoscaling/groups/instances/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/autoscaling/groups/instances")
	} else {
		url = fmt.Sprintf("zstack/v1/autoscaling/groups/instances/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAutoScalingGroupInstanceFailed, err.Error())
	}
	var result model.QueryAutoScalingGroupInstanceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改弹性伸缩组启用状态
func ChangeAutoScalingGroupState(params model.ChangeAutoScalingGroupStateRequest) models.Result {
	//PUT zstack/v1/autoscaling/groups/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/autoscaling/groups/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeAutoScalingGroupStateFailed, err.Error())
	}
	var result model.ChangeAutoScalingGroupStateParams
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//手动执行伸缩组规则
func ExecuteAutoScalingRule(params model.ExecuteAutoScalingRuleRequest) models.Result {
	//PUT zstack/v1/autoscaling/rules/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/autoscaling/rules/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ExecuteAutoScalingRuleFailed, err.Error())
	}
	var result model.ExecuteAutoScalingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新伸缩组实例信息
func UpdateAutoScalingGroupInstance(params model.UpdateAutoScalingGroupInstanceRequest) models.Result {
	//PUT zstack/v1/autoscaling/groups/instances/{instanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/autoscaling/groups/instances/%s/actions", params.InstanceUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateAutoScalingGroupInstanceFailed, err.Error())
	}
	var result model.UpdateAutoScalingGroupInstanceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新伸缩组云主机模板
func UpdateAutoScalingVmTemplate(params model.UpdateAutoScalingVmTemplateRequest) models.Result {
	//PUT zstack/v1/autoscaling/vmtemplate/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/autoscaling/vmtemplate/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateAutoScalingVmTemplateFailed, err.Error())
	}
	var result model.UpdateAutoScalingVmTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
