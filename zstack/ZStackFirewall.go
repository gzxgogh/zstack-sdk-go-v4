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

//创建防火墙
func CreateVpcFirewall(params model.CreateVpcFirewallRequest) models.Result {
	//POST zstack/v1/vpcfirewalls
	url := fmt.Sprintf("zstack/v1/vpcfirewalls")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateVpcFirewallFailed, err.Error())
	}
	var result model.CreateVpcFirewallResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询防火墙
func QueryVpcFirewall(params model.QueryVpcFirewallRequest) models.Result {
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
		return models.Error(errcode.QueryVpcFirewallFailed, err.Error())
	}
	var result model.QueryVpcFirewallResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新防火墙
func UpdateVpcFirewall(params model.UpdateVpcFirewallRequest) models.Result {
	//PUT zstack/v1/vpcfirewalls/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateVpcFirewallFailed, err.Error())
	}
	var result model.UpdateVpcFirewallResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//刷新防火墙配置
func RefreshFirewall(params model.RefreshFirewallRequest) models.Result {
	//PUT zstack/v1/vpcfirewalls/refresh/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/refresh/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.RefreshFirewallFailed, err.Error())
	}
	var result model.RefreshFirewallResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除防火墙
func DeleteFirewall(params model.DeleteFirewallRequest) models.Result {
	//DELETE zstack/v1/vpcfirewalls/{uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteFirewallFailed, err.Error())
	}
	var result model.DeleteFirewallResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建防火墙规则集
func CreateFirewallRuleSet(params model.CreateFirewallRuleSetRequest) models.Result {
	//POST zstack/v1/vpcfirewalls/ruleSets
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ruleSets")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateFirewallRuleSetFailed, err.Error())
	}
	var result model.CreateFirewallRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询防火墙规则集
func QueryFirewallRuleSet(params model.QueryFirewallRuleSetRequest) models.Result {
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
		return models.Error(errcode.QueryFirewallRuleSetFailed, err.Error())
	}
	var result model.QueryFirewallRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新防火墙规则集
func UpdateFirewallRuleSet(params model.UpdateFirewallRuleSetRequest) models.Result {
	//PUT zstack/v1/vpcfirewalls/ruleSets/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ruleSets/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateFirewallRuleSetFailed, err.Error())
	}
	var result model.UpdateFirewallRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除防火墙规则集
func DeleteFirewallRuleSet(params model.DeleteFirewallRuleSetRequest) models.Result {
	//DELETE zstack/v1/vpcfirewalls/ruleSets/{uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ruleSets/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteFirewallRuleSetFailed, err.Error())
	}
	var result model.DeleteFirewallRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//网络加载防火墙规则集
func AttachFirewallRuleSetToL3(params model.AttachFirewallRuleSetToL3Request) models.Result {
	//POST zstack/v1/vpcfirewalls/ruleSets/{ruleSetUuid}/l3networks/{l3Uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ruleSets/%s/l3networks/%s", params.RuleSetUuid, params.L3Uuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachFirewallRuleSetToL3Failed, err.Error())
	}
	var result model.AttachFirewallRuleSetToL3Response
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//网络卸载防火墙规则集
func DetachFirewallRuleSetFromL3(params model.DetachFirewallRuleSetFromL3Request) models.Result {
	//POST zstack/v1/vpcfirewalls/l3networks/{l3Uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/l3networks/%s", params.L3Uuid, params.L3Uuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.DetachFirewallRuleSetFromL3Failed, err.Error())
	}
	var result model.DetachFirewallRuleSetFromL3Response
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询防火墙规则集与三层网络关联关系
func QueryFirewallRuleSetL3Ref(params model.QueryFirewallRuleSetL3RefRequest) models.Result {
	//GET zstack/v1/vpcfirewalls/l3networks/rulesets/refs
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/l3networks/rulesets/refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryFirewallRuleSetL3RefFailed, err.Error())
	}
	var result model.QueryFirewallRuleSetL3RefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建防火墙规则
func CreateFirewallRule(params model.CreateFirewallRuleRequest) models.Result {
	//POST zstack/v1/vpcfirewalls/rules
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateFirewallRuleFailed, err.Error())
	}
	var result model.CreateFirewallRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询防火墙规则
func QueryFirewallRule(params model.QueryFirewallRuleRequest) models.Result {
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
		return models.Error(errcode.QueryFirewallRuleFailed, err.Error())
	}
	var result model.QueryFirewallRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新防火墙规则
func UpdateFirewallRule(params model.UpdateFirewallRuleRequest) models.Result {
	//PUT zstack/v1/vpcfirewalls/rules/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateFirewallRuleFailed, err.Error())
	}
	var result model.UpdateFirewallRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除防火墙规则
func DeleteFirewallRule(params model.DeleteFirewallRuleRequest) models.Result {
	//DELETE zstack/v1/vpcfirewalls/rules/{uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteFirewallRuleFailed, err.Error())
	}
	var result model.DeleteFirewallRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更改防火墙规则状态
func ChangeFirewallRuleState(params model.ChangeFirewallRuleStateRequest) models.Result {
	//PUT zstack/v1/vpcfirewalls/rules/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeFirewallRuleStateFailed, err.Error())
	}
	var result model.ChangeFirewallRuleStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建防火墙规则模板
func CreateFirewallRuleTemplate(params model.CreateFirewallRuleTemplateRequest) models.Result {
	//POST zstack/v1/vpcfirewalls/rules/template
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules/template")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateFirewallRuleTemplateFailed, err.Error())
	}
	var result model.CreateFirewallRuleTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除防火墙规则模板
func DeleteFirewallRuleTemplate(params model.DeleteFirewallRuleTemplateRequest) models.Result {
	//DELETE zstack/v1/vpcfirewalls/rules/templates/{uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules/templates/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteFirewallRuleTemplateFailed, err.Error())
	}
	var result model.DeleteFirewallRuleTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新防火墙规则模板
func UpdateFirewallRuleTemplate(params model.UpdateFirewallRuleTemplateRequest) models.Result {
	//PUT zstack/v1/vpcfirewalls/rules/template/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/rules/template/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateFirewallRuleTemplateFailed, err.Error())
	}
	var result model.UpdateFirewallRuleTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询防火墙规则模板
func QueryFirewallRuleTemplate(params model.QueryFirewallRuleTemplateRequest) models.Result {
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
		return models.Error(errcode.QueryFirewallRuleTemplateFailed, err.Error())
	}
	var result model.QueryFirewallRuleTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询防火墙与单节点路由器关联关系
func QueryVpcFirewallVRouterRef(params model.QueryVpcFirewallVRouterRefRequest) models.Result {
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
		return models.Error(errcode.QueryVpcFirewallVRouterRefFailed, err.Error())
	}
	var result model.QueryVpcFirewallVRouterRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//下发规则集更改内容到路由器
func ApplyRuleSetChanges(params model.ApplyRuleSetChangesRequest) models.Result {
	//PUT zstack/v1/vpcfirewalls/ruleSets/apply/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ruleSets/apply/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ApplyRuleSetChangesFailed, err.Error())
	}
	var result model.ApplyRuleSetChangesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建防火墙IP或端口模板
func CreateFirewallIpSetTemplate(params model.CreateFirewallIpSetTemplateRequest) models.Result {
	//POST zstack/v1/vpcfirewalls/ipset/templates
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ipset/templates")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateFirewallIpSetTemplateFailed, err.Error())
	}
	var result model.CreateFirewallIpSetTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除防火墙IP或端口模板
func DeleteFirewallIpSetTemplate(params model.DeleteFirewallIpSetTemplateRequest) models.Result {
	//DELETE zstack/v1/vpcfirewalls/ipset/templates/{uuid}
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ipset/templates/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteFirewallIpSetTemplateFailed, err.Error())
	}
	var result model.DeleteFirewallIpSetTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新防火墙IP或端口模板
func UpdateFirewallIpSetTemplate(params model.UpdateFirewallIpSetTemplateRequest) models.Result {
	//PUT zstack/v1/vpcfirewalls/ipset/templates/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpcfirewalls/ipset/templates/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateFirewallIpSetTemplateFailed, err.Error())
	}
	var result model.UpdateFirewallIpSetTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询防火墙IP或端口模板
func QueryFirewallIpSetTemplate(params model.QueryFirewallIpSetTemplateRequest) models.Result {
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
		return models.Error(errcode.QueryFirewallIpSetTemplateFailed, err.Error())
	}
	var result model.QueryFirewallIpSetTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
