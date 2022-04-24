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

//创建端口转发规则
func CreatePortForwardingRule(params model.CreatePortForwardingRuleRequest) mgresult.Result {
	//POST zstack/v1/port-forwarding
	url := fmt.Sprintf("zstack/v1/port-forwarding")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreatePortForwardingRuleFailed, err)
	}
	var result model.CreatePortForwardingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除端口转发规则
func DeletePortForwardingRule(params model.DeletePortForwardingRuleRequest) mgresult.Result {
	//DELETE zstack/v1/port-forwarding/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/port-forwarding/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeletePortForwardingRuleFailed, err)
	}
	var result model.DeletePortForwardingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询端口转发规则
func QueryPortForwardingRule(params model.QueryPortForwardingRuleRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryPortForwardingRuleFailed, err)
	}
	var result model.QueryPortForwardingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新端口转发规则
func UpdatePortForwardingRule(params model.UpdatePortForwardingRuleRequest) mgresult.Result {
	//PUT zstack/v1/port-forwarding/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/port-forwarding/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdatePortForwardingRuleFailed, err)
	}
	var result model.UpdatePortForwardingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//改变端口转发规则的状态
func ChangePortForwardingRuleState(params model.ChangePortForwardingRuleStateRequest) mgresult.Result {
	//PUT zstack/v1/port-forwarding/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/port-forwarding/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangePortForwardingRuleStateFailed, err)
	}
	var result model.ChangePortForwardingRuleStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取云主机网卡列表
func GetPortForwardingAttachableVmNics(params model.GetPortForwardingAttachableVmNicsRequest) mgresult.Result {
	//GET zstack/v1/port-forwarding/{ruleUuid}/vm-instances/candidate-nics
	url := fmt.Sprintf("zstack/v1/port-forwarding/%s/vm-instances/candidate-nics", params.RuleUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetPortForwardingAttachableVmNicsFailed, err)
	}
	var result model.GetPortForwardingAttachableVmNicsResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//挂载规则到虚拟机网卡上
func AttachPortForwardingRule(params model.AttachPortForwardingRuleRequest) mgresult.Result {
	//POST zstack/v1/port-forwarding/{ruleUuid}/vm-instances/nics/{vmNicUuid}
	url := fmt.Sprintf("zstack/v1/port-forwarding/%s/vm-instances/nics/%s", params.RuleUuid, params.VmNicUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachPortForwardingRuleFailed, err)
	}
	var result model.AttachPortForwardingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从虚拟机网卡卸载规则
func DetachPortForwardingRule(params model.DetachPortForwardingRuleRequest) mgresult.Result {
	//DELETE zstack/v1/port-forwarding/{uuid}/vm-instances/nics
	url := fmt.Sprintf("zstack/v1/port-forwarding/%s/vm-instances/nics", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachPortForwardingRuleFailed, err)
	}
	var result model.DetachPortForwardingRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
