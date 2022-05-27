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

//创建安全组
func CreateSecurityGroup(params model.CreateSecurityGroupRequest) mgresult.Result {
	//POST zstack/v1/security-groups
	url := fmt.Sprintf("zstack/v1/security-groups")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateSecurityGroupFailed, err)
	}
	var result model.CreateSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除安全组
func DeleteSecurityGroup(params model.DeleteSecurityGroupRequest) mgresult.Result {
	//DELETE zstack/v1/security-groups/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/security-groups/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteSecurityGroupFailed, err)
	}
	var result model.DeleteSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询安全组
func QuerySecurityGroup(params model.QuerySecurityGroupRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QuerySecurityGroupFailed, err)
	}
	var result model.QuerySecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新安全组
func UpdateSecurityGroup(params model.UpdateSecurityGroupRequest) mgresult.Result {
	//PUT zstack/v1/security-groups/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/security-groups/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateSecurityGroupFailed, err)
	}
	var result model.UpdateSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//改变安全组状态
func ChangeSecurityGroupState(params model.ChangeSecurityGroupStateRequest) mgresult.Result {
	//PUT zstack/v1/security-groups/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/security-groups/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeSecurityGroupStateFailed, err)
	}
	var result model.ChangeSecurityGroupStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//挂载安全组到L3网络
func AttachSecurityGroupToL3Network(params model.AttachSecurityGroupToL3NetworkRequest) mgresult.Result {
	//POST zstack/v1/security-groups/{securityGroupUuid}/l3-networks/{l3NetworkUuid}
	url := fmt.Sprintf("zstack/v1/security-groups/%s/l3-networks/%s", params.SecurityGroupUuid, params.L3NetworkUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachSecurityGroupToL3NetworkFailed, err)
	}
	var result model.AttachSecurityGroupToL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从L3网络卸载安全组
func DetachSecurityGroupFromL3Network(params model.DetachSecurityGroupFromL3NetworkRequest) mgresult.Result {
	//DELETE zstack/v1/security-groups/{securityGroupUuid}/l3-networks/{l3NetworkUuid}
	url := fmt.Sprintf("zstack/v1/security-groups/%s/l3-networks/%s", params.SecurityGroupUuid, params.L3NetworkUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachSecurityGroupFromL3NetworkFailed, err)
	}
	var result model.DetachSecurityGroupFromL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取网卡列表清单
func GetCandidateVmNicForSecurityGroup(params model.GetCandidateVmNicForSecurityGroupRequest) mgresult.Result {
	//GET zstack/v1/security-groups/{securityGroupUuid}/vm-instances/candidate-nics
	url := fmt.Sprintf("zstack/v1/security-groups/%s/vm-instances/candidate-nics", params.SecurityGroupUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCandidateVmNicForSecurityGroupFailed, err)
	}
	var result model.GetCandidateVmNicForSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加虚拟机网卡到安全组
func AddVmNicToSecurityGroup(params model.AddVmNicToSecurityGroupRequest) mgresult.Result {
	//POST zstack/v1/security-groups/{securityGroupUuid}/vm-instances/nics
	url := fmt.Sprintf("zstack/v1/security-groups/%s/vm-instances/nics", params.SecurityGroupUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddVmNicToSecurityGroupFailed, err)
	}
	var result model.AddVmNicToSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从安全组删除虚拟机网卡
func DeleteVmNicFromSecurityGroup(params model.DeleteVmNicFromSecurityGroupRequest) mgresult.Result {
	//DELETE zstack/v1/security-groups/{securityGroupUuid}/vm-instances/nics?vmNicUuids={vmNicUuids}
	url := fmt.Sprintf("zstack/v1/security-groups/%s/vm-instances/nics?vmNicUuids=%s", params.SecurityGroupUuid, params.VmNicUuids)

	dataStr, err := request.DeleteUrlWithParams(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVmNicFromSecurityGroupFailed, err)
	}
	var result model.DeleteVmNicFromSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询应用了安全组的网卡列表
func QueryVmNicInSecurityGroup(params model.QueryVmNicInSecurityGroupRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryVmNicInSecurityGroupFailed, err)
	}
	var result model.QueryVmNicInSecurityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加规则到安全组
func AddSecurityGroupRule(params model.AddSecurityGroupRuleRequest) mgresult.Result {
	//POST zstack/v1/security-groups/{securityGroupUuid}/rules
	url := fmt.Sprintf("zstack/v1/security-groups/%s/rules", params.SecurityGroupUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddSecurityGroupRuleFailed, err)
	}
	var result model.AddSecurityGroupRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询安全组规则
func QuerySecurityGroupRule(params model.QuerySecurityGroupRuleRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QuerySecurityGroupRuleFailed, err)
	}
	var result model.QuerySecurityGroupRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除安全组规则
func DeleteSecurityGroupRule(params model.DeleteSecurityGroupRuleRequest) mgresult.Result {
	//DELETE zstack/v1/security-groups/rules?ruleUuids={ruleUuids}
	url := fmt.Sprintf("zstack/v1/security-groups/rules?ruleUuids=%s", params.RuleUuids)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteSecurityGroupRuleFailed, err)
	}
	var result model.DeleteSecurityGroupRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
