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

//创建防火墙
func CreateVpcFirewall(params model.CreateVpcFirewallRequest) mgresult.Result {
	//POST zstack/v1/vpcfirewalls
	url := fmt.Sprintf("zstack/v1/vpcfirewalls")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateVpcFirewallFailed, err)
	}
	var result model.CreateVpcFirewallResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询防火墙
func QueryVpcFirewall(params model.QueryVpcFirewallRequest) mgresult.Result {
	//GET zstack/v1/vpcfirewalls
	//GET zstack/v1/vpcfirewalls/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vpcfirewalls")
	} else {
		url = fmt.Sprintf("zstack/v1/vpcfirewalls/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVpcFirewallFailed, err)
	}
	var result model.QueryVpcFirewallResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新防火墙
func UpdateVpcFirewall(params model.UpdateVpcFirewallRequest) mgresult.Result {
	//PUT zstack/v1/vpcfirewalls/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateVpcFirewallFailed, err)
	}
	var result model.UpdateVpcFirewallResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//刷新防火墙配置
func RefreshFirewall(params model.RefreshFirewallRequest) mgresult.Result {
	//PUT zstack/v1/vpcfirewalls/refresh/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/refresh/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RefreshFirewallFailed, err)
	}
	var result model.RefreshFirewallResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除防火墙
func DeleteFirewall(params model.DeleteFirewallRequest) mgresult.Result {
	//DELETE zstack/v1/vpcfirewalls/{uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteFirewallFailed, err)
	}
	var result model.DeleteFirewallResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建防火墙规则集
func CreateFirewallRuleSet(params model.CreateFirewallRuleSetRequest) mgresult.Result {
	//POST zstack/v1/vpcfirewalls/ruleSets
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ruleSets")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateFirewallRuleSetFailed, err)
	}
	var result model.CreateFirewallRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询防火墙规则集
func QueryFirewallRuleSet(params model.QueryFirewallRuleSetRequest) mgresult.Result {
	//GET zstack/v1/vpcfirewalls/ruleSets
	//GET zstack/v1/vpcfirewalls/ruleSets/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vpcfirewalls/ruleSets")
	} else {
		url = fmt.Sprintf("zstack/v1/vpcfirewalls/ruleSets/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryFirewallRuleSetFailed, err)
	}
	var result model.QueryFirewallRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新防火墙规则集
func UpdateFirewallRuleSet(params model.UpdateFirewallRuleSetRequest) mgresult.Result {
	//PUT zstack/v1/vpcfirewalls/ruleSets/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ruleSets/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateFirewallRuleSetFailed, err)
	}
	var result model.UpdateFirewallRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除防火墙规则集
func DeleteFirewallRuleSet(params model.DeleteFirewallRuleSetRequest) mgresult.Result {
	//DELETE zstack/v1/vpcfirewalls/ruleSets/{uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ruleSets/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteFirewallRuleSetFailed, err)
	}
	var result model.DeleteFirewallRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//网络加载防火墙规则集
func AttachFirewallRuleSetToL3(params model.AttachFirewallRuleSetToL3Request) mgresult.Result {
	//POST zstack/v1/vpcfirewalls/ruleSets/{ruleSetUuid}/l3networks/{l3Uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ruleSets/%s/l3networks/%s", params.RuleSetUuid, params.L3Uuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachFirewallRuleSetToL3Failed, err)
	}
	var result model.AttachFirewallRuleSetToL3Response
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//网络卸载防火墙规则集
func DetachFirewallRuleSetFromL3(params model.DetachFirewallRuleSetFromL3Request) mgresult.Result {
	//POST zstack/v1/vpcfirewalls/l3networks/{l3Uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/l3networks/%s", params.L3Uuid, params.L3Uuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachFirewallRuleSetFromL3Failed, err)
	}
	var result model.DetachFirewallRuleSetFromL3Response
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询防火墙规则集与三层网络关联关系
func QueryFirewallRuleSetL3Ref(params model.QueryFirewallRuleSetL3RefRequest) mgresult.Result {
	//GET zstack/v1/vpcfirewalls/l3networks/rulesets/refs
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/l3networks/rulesets/refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryFirewallRuleSetL3RefFailed, err)
	}
	var result model.QueryFirewallRuleSetL3RefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建防火墙规则
func CreateFirewallRule(params model.CreateFirewallRuleRequest) mgresult.Result {
	//POST zstack/v1/vpcfirewalls/rules
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateFirewallRuleFailed, err)
	}
	var result model.CreateFirewallRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询防火墙规则
func QueryFirewallRule(params model.QueryFirewallRuleRequest) mgresult.Result {
	//GET zstack/v1/vpcfirewalls/rules
	//GET zstack/v1/vpcfirewalls/rules/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vpcfirewalls/rules")
	} else {
		url = fmt.Sprintf("zstack/v1/vpcfirewalls/rules/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryFirewallRuleFailed, err)
	}
	var result model.QueryFirewallRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新防火墙规则
func UpdateFirewallRule(params model.UpdateFirewallRuleRequest) mgresult.Result {
	//PUT zstack/v1/vpcfirewalls/rules/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateFirewallRuleFailed, err)
	}
	var result model.UpdateFirewallRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除防火墙规则
func DeleteFirewallRule(params model.DeleteFirewallRuleRequest) mgresult.Result {
	//DELETE zstack/v1/vpcfirewalls/rules/{uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteFirewallRuleFailed, err)
	}
	var result model.DeleteFirewallRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更改防火墙规则状态
func ChangeFirewallRuleState(params model.ChangeFirewallRuleStateRequest) mgresult.Result {
	//PUT zstack/v1/vpcfirewalls/rules/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeFirewallRuleStateFailed, err)
	}
	var result model.ChangeFirewallRuleStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建防火墙规则模板
func CreateFirewallRuleTemplate(params model.CreateFirewallRuleTemplateRequest) mgresult.Result {
	//POST zstack/v1/vpcfirewalls/rules/template
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules/template")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateFirewallRuleTemplateFailed, err)
	}
	var result model.CreateFirewallRuleTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除防火墙规则模板
func DeleteFirewallRuleTemplate(params model.DeleteFirewallRuleTemplateRequest) mgresult.Result {
	//DELETE zstack/v1/vpcfirewalls/rules/templates/{uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules/templates/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteFirewallRuleTemplateFailed, err)
	}
	var result model.DeleteFirewallRuleTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新防火墙规则模板
func UpdateFirewallRuleTemplate(params model.UpdateFirewallRuleTemplateRequest) mgresult.Result {
	//PUT zstack/v1/vpcfirewalls/rules/template/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules/template/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateFirewallRuleTemplateFailed, err)
	}
	var result model.UpdateFirewallRuleTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询防火墙规则模板
func QueryFirewallRuleTemplate(params model.QueryFirewallRuleTemplateRequest) mgresult.Result {
	//GET zstack/v1/vpcfirewalls/rules/templates
	//GET zstack/v1/vpcfirewalls/rules/templates/uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vpcfirewalls/rules/templates")
	} else {
		url = fmt.Sprintf("zstack/v1/vpcfirewalls/rules/templates/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryFirewallRuleTemplateFailed, err)
	}
	var result model.QueryFirewallRuleTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询防火墙与单节点路由器关联关系
func QueryVpcFirewallVRouterRef(params model.QueryVpcFirewallVRouterRefRequest) mgresult.Result {
	//GET zstack/v1/vpcfirewalls/vrouters/refs
	//GET zstack/v1/vpcfirewalls/vrouters/refs/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vpcfirewalls/vrouters/refs")
	} else {
		url = fmt.Sprintf("zstack/v1/vpcfirewalls/vrouters/refs/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVpcFirewallVRouterRefFailed, err)
	}
	var result model.QueryVpcFirewallVRouterRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//下发规则集更改内容到路由器
func ApplyRuleSetChanges(params model.ApplyRuleSetChangesRequest) mgresult.Result {
	//PUT zstack/v1/vpcfirewalls/ruleSets/apply/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ruleSets/apply/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ApplyRuleSetChangesFailed, err)
	}
	var result model.ApplyRuleSetChangesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建防火墙IP或端口模板
func CreateFirewallIpSetTemplate(params model.CreateFirewallIpSetTemplateRequest) mgresult.Result {
	//POST zstack/v1/vpcfirewalls/ipset/templates
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ipset/templates")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateFirewallIpSetTemplateFailed, err)
	}
	var result model.CreateFirewallIpSetTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除防火墙IP或端口模板
func DeleteFirewallIpSetTemplate(params model.DeleteFirewallIpSetTemplateRequest) mgresult.Result {
	//DELETE zstack/v1/vpcfirewalls/ipset/templates/{uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ipset/templates/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteFirewallIpSetTemplateFailed, err)
	}
	var result model.DeleteFirewallIpSetTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新防火墙IP或端口模板
func UpdateFirewallIpSetTemplate(params model.UpdateFirewallIpSetTemplateRequest) mgresult.Result {
	//PUT zstack/v1/vpcfirewalls/ipset/templates/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ipset/templates/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateFirewallIpSetTemplateFailed, err)
	}
	var result model.UpdateFirewallIpSetTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询防火墙IP或端口模板
func QueryFirewallIpSetTemplate(params model.QueryFirewallIpSetTemplateRequest) mgresult.Result {
	//GET zstack/v1/vpcfirewalls/ipset/templates
	//GET zstack/v1/vpcfirewalls/ipset/templates/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vpcfirewalls/ipset/templates")
	} else {
		url = fmt.Sprintf("zstack/v1/vpcfirewalls/ipset/templates/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryFirewallIpSetTemplateFailed, err)
	}
	var result model.QueryFirewallIpSetTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
