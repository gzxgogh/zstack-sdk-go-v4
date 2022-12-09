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

//创建安全组
func CreateSecurityGroup(params model.CreateSecurityGroupRequest) models.Result {
	//POST zstack/v1/security-groups
	url := fmt.Sprintf("zstack/v1/security-groups")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateSecurityGroupFailed, err.Error())
	}
	var result model.CreateSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除安全组
func DeleteSecurityGroup(params model.DeleteSecurityGroupRequest) models.Result {
	//DELETE zstack/v1/security-groups/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/security-groups/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteSecurityGroupFailed, err.Error())
	}
	var result model.DeleteSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询安全组
func QuerySecurityGroup(params model.QuerySecurityGroupRequest) models.Result {
	//GET zstack/v1/security-groups
	//GET zstack/v1/security-groups/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/security-groups")
	} else {
		url = fmt.Sprintf("zstack/v1/security-groups/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QuerySecurityGroupFailed, err.Error())
	}
	var result model.QuerySecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新安全组
func UpdateSecurityGroup(params model.UpdateSecurityGroupRequest) models.Result {
	//PUT zstack/v1/security-groups/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/security-groups/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateSecurityGroupFailed, err.Error())
	}
	var result model.UpdateSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//改变安全组状态
func ChangeSecurityGroupState(params model.ChangeSecurityGroupStateRequest) models.Result {
	//PUT zstack/v1/security-groups/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/security-groups/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeSecurityGroupStateFailed, err.Error())
	}
	var result model.ChangeSecurityGroupStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//挂载安全组到L3网络
func AttachSecurityGroupToL3Network(params model.AttachSecurityGroupToL3NetworkRequest) models.Result {
	//POST zstack/v1/security-groups/{securityGroupUuid}/l3-networks/{l3NetworkUuid}
	url := fmt.Sprintf("zstack/v1/security-groups/%s/l3-networks/%s", params.SecurityGroupUuid, params.L3NetworkUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachSecurityGroupToL3NetworkFailed, err.Error())
	}
	var result model.AttachSecurityGroupToL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从L3网络卸载安全组
func DetachSecurityGroupFromL3Network(params model.DetachSecurityGroupFromL3NetworkRequest) models.Result {
	//DELETE zstack/v1/security-groups/{securityGroupUuid}/l3-networks/{l3NetworkUuid}
	url := fmt.Sprintf("zstack/v1/security-groups/%s/l3-networks/%s", params.SecurityGroupUuid, params.L3NetworkUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachSecurityGroupFromL3NetworkFailed, err.Error())
	}
	var result model.DetachSecurityGroupFromL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取网卡列表清单
func GetCandidateVmNicForSecurityGroup(params model.GetCandidateVmNicForSecurityGroupRequest) models.Result {
	//GET zstack/v1/security-groups/{securityGroupUuid}/vm-instances/candidate-nics
	url := fmt.Sprintf("zstack/v1/security-groups/%s/vm-instances/candidate-nics", params.SecurityGroupUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetCandidateVmNicForSecurityGroupFailed, err.Error())
	}
	var result model.GetCandidateVmNicForSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加虚拟机网卡到安全组
func AddVmNicToSecurityGroup(params model.AddVmNicToSecurityGroupRequest) models.Result {
	//POST zstack/v1/security-groups/{securityGroupUuid}/vm-instances/nics
	url := fmt.Sprintf("zstack/v1/security-groups/%s/vm-instances/nics", params.SecurityGroupUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddVmNicToSecurityGroupFailed, err.Error())
	}
	var result model.AddVmNicToSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从安全组删除虚拟机网卡
func DeleteVmNicFromSecurityGroup(params model.DeleteVmNicFromSecurityGroupRequest) models.Result {
	//DELETE zstack/v1/security-groups/{securityGroupUuid}/vm-instances/nics?vmNicUuids={vmNicUuids}
	url := fmt.Sprintf("zstack/v1/security-groups/%s/vm-instances/nics", params.SecurityGroupUuid)

	dataStr, err := request.DeleteUrlWithParams(url, params)
	if err != nil {
		return models.Error(errcode.DeleteVmNicFromSecurityGroupFailed, err.Error())
	}
	var result model.DeleteVmNicFromSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询应用了安全组的网卡列表
func QueryVmNicInSecurityGroup(params model.QueryVmNicInSecurityGroupRequest) models.Result {
	//GET zstack/v1/security-groups/vm-instances/nics
	//GET zstack/v1/security-groups/vm-instances/nics/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/security-groups/vm-instances/nics")
	} else {
		url = fmt.Sprintf("zstack/v1/security-groups/vm-instances/nics/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVmNicInSecurityGroupFailed, err.Error())
	}
	var result model.QueryVmNicInSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加规则到安全组
func AddSecurityGroupRule(params model.AddSecurityGroupRuleRequest) models.Result {
	//POST zstack/v1/security-groups/{securityGroupUuid}/rules
	url := fmt.Sprintf("zstack/v1/security-groups/%s/rules", params.SecurityGroupUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddSecurityGroupRuleFailed, err.Error())
	}
	var result model.AddSecurityGroupRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询安全组规则
func QuerySecurityGroupRule(params model.QuerySecurityGroupRuleRequest) models.Result {
	//GET zstack/v1/security-groups/rules
	//GET zstack/v1/security-groups/rules/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/security-groups/rules")
	} else {
		url = fmt.Sprintf("zstack/v1/security-groups/rules/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QuerySecurityGroupRuleFailed, err.Error())
	}
	var result model.QuerySecurityGroupRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除安全组规则
func DeleteSecurityGroupRule(params model.DeleteSecurityGroupRuleRequest) models.Result {
	//DELETE zstack/v1/security-groups/rules?ruleUuids={ruleUuids}
	url := fmt.Sprintf("zstack/v1/security-groups/rules")

	dataStr, err := request.DeleteUrlWithParams(url, params)
	if err != nil {
		return models.Error(errcode.DeleteSecurityGroupRuleFailed, err.Error())
	}
	var result model.DeleteSecurityGroupRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
