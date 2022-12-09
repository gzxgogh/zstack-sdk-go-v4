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

//创建端口转发规则
func CreatePortForwardingRule(params model.CreatePortForwardingRuleRequest) models.Result {
	//POST zstack/v1/port-forwarding
	url := fmt.Sprintf("zstack/v1/port-forwarding")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreatePortForwardingRuleFailed, err.Error())
	}
	var result model.CreatePortForwardingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除端口转发规则
func DeletePortForwardingRule(params model.DeletePortForwardingRuleRequest) models.Result {
	//DELETE zstack/v1/port-forwarding/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/port-forwarding/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeletePortForwardingRuleFailed, err.Error())
	}
	var result model.DeletePortForwardingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询端口转发规则
func QueryPortForwardingRule(params model.QueryPortForwardingRuleRequest) models.Result {
	//GET zstack/v1/port-forwarding
	//GET zstack/v1/port-forwarding/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/port-forwarding")
	} else {
		url = fmt.Sprintf("zstack/v1/port-forwarding/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPortForwardingRuleFailed, err.Error())
	}
	var result model.QueryPortForwardingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新端口转发规则
func UpdatePortForwardingRule(params model.UpdatePortForwardingRuleRequest) models.Result {
	//PUT zstack/v1/port-forwarding/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/port-forwarding/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdatePortForwardingRuleFailed, err.Error())
	}
	var result model.UpdatePortForwardingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//改变端口转发规则的状态
func ChangePortForwardingRuleState(params model.ChangePortForwardingRuleStateRequest) models.Result {
	//PUT zstack/v1/port-forwarding/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/port-forwarding/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangePortForwardingRuleStateFailed, err.Error())
	}
	var result model.ChangePortForwardingRuleStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取云主机网卡列表
func GetPortForwardingAttachableVmNics(params model.GetPortForwardingAttachableVmNicsRequest) models.Result {
	//GET zstack/v1/port-forwarding/{ruleUuid}/vm-instances/candidate-nics
	url := fmt.Sprintf("zstack/v1/port-forwarding/%s/vm-instances/candidate-nics", params.RuleUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetPortForwardingAttachableVmNicsFailed, err.Error())
	}
	var result model.GetPortForwardingAttachableVmNicsResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//挂载规则到虚拟机网卡上
func AttachPortForwardingRule(params model.AttachPortForwardingRuleRequest) models.Result {
	//POST zstack/v1/port-forwarding/{ruleUuid}/vm-instances/nics/{vmNicUuid}
	url := fmt.Sprintf("zstack/v1/port-forwarding/%s/vm-instances/nics/%s", params.RuleUuid, params.VmNicUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachPortForwardingRuleFailed, err.Error())
	}
	var result model.AttachPortForwardingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从虚拟机网卡卸载规则
func DetachPortForwardingRule(params model.DetachPortForwardingRuleRequest) models.Result {
	//DELETE zstack/v1/port-forwarding/{uuid}/vm-instances/nics
	url := fmt.Sprintf("zstack/v1/port-forwarding/%s/vm-instances/nics", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachPortForwardingRuleFailed, err.Error())
	}
	var result model.DetachPortForwardingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
