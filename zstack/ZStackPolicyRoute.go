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

//创建策略路由规则集
func CreatePolicyRouteRuleSet(params model.CreatePolicyRouteRuleSetRequest) models.Result {
	//POST zstack/v1/policy-routes/rulesets
	url := fmt.Sprintf("zstack/v1/policy-routes/rulesets")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreatePolicyRouteRuleSetFailed, err.Error())
	}
	var result model.CreatePolicyRouteRuleSetRequest
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除策略路由规则集
func DeletePolicyRouteRuleSet(params model.DeletePolicyRouteRuleSetRequest) models.Result {
	//DELETE zstack/v1/policy-routes/ruleSets/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/policy-routes/ruleSets/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeletePolicyRouteRuleSetFailed, err.Error())
	}
	var result model.DeletePolicyRouteRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新策略路由规则集属性
func UpdatePolicyRouteRuleSet(params model.UpdatePolicyRouteRuleSetRequest) models.Result {
	//PUT zstack/v1/policy-routes/ruleSets/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/policy-routes/ruleSets/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdatePolicyRouteRuleSetFailed, err.Error())
	}
	var result model.UpdatePolicyRouteRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询策略路由规则集
func QueryPolicyRouteRuleSet(params model.QueryPolicyRouteRuleSetRequest) models.Result {
	//GET zstack/v1/policy-routes/rulesets
	//GET zstack/v1/policy-routes/rulesets/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/policy-routes/rulesets")
	} else {
		url = fmt.Sprintf("zstack/v1/policy-routes/rulesets/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPolicyRouteRuleSetFailed, err.Error())
	}
	var result model.QueryPolicyRouteRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询策略路由规则集网络关联
func QueryPolicyRouteRuleSetL3Ref(params model.QueryPolicyRouteRuleSetL3RefRequest) models.Result {
	//GET zstack/v1/policy-routes/rulesets/l3networdks/refs
	url := fmt.Sprintf("zstack/v1/policy-routes/rulesets/l3networdks/refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPolicyRouteRuleSetL3RefFailed, err.Error())
	}
	var result model.QueryPolicyRouteRuleSetL3RefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询策略路由规则集与单节点云路由关联
func QueryPolicyRouteRuleSetVRouterRef(params model.QueryPolicyRouteRuleSetVRouterRefRequest) models.Result {
	//GET zstack/v1/policy-routes/rulesets/vrouters/refs
	url := fmt.Sprintf("zstack/v1/policy-routes/rulesets/vrouters/refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPolicyRouteRuleSetVRouterRefFailed, err.Error())
	}
	var result model.QueryPolicyRouteRuleSetVRouterRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建策略路由规则集规则
func CreatePolicyRouteRule(params model.CreatePolicyRouteRuleRequest) models.Result {
	//POST zstack/v1/policy-routes/rules
	url := fmt.Sprintf("zstack/v1/policy-routes/rules")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreatePolicyRouteRuleFailed, err.Error())
	}
	var result model.CreatePolicyRouteRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除策略路由规则集规则
func DeletePolicyRouteRule(params model.DeletePolicyRouteRuleRequest) models.Result {
	//DELETE zstack/v1/policy-routes/rules/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/policy-routes/rules/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeletePolicyRouteRuleFailed, err.Error())
	}
	var result model.DeletePolicyRouteRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建策略路由表
func CreatePolicyRouteTable(params model.CreatePolicyRouteTableRequest) models.Result {
	//POST zstack/v1/policy-routes/tables
	url := fmt.Sprintf("zstack/v1/policy-routes/tables")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreatePolicyRouteTableFailed, err.Error())
	}
	var result model.CreatePolicyRouteTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除策略路由表
func DeletePolicyRouteTable(params model.DeletePolicyRouteTableRequest) models.Result {
	//DELETE zstack/v1/policy-routes/tables/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/policy-routes/tables/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeletePolicyRouteTableFailed, err.Error())
	}
	var result model.DeletePolicyRouteTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询策略路由表
func QueryPolicyRouteTable(params model.QueryPolicyRouteTableRequest) models.Result {
	//GET zstack/v1/policy-routes/tables
	//GET zstack/v1/policy-routes/tables/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/policy-routes/tables")
	} else {
		url = fmt.Sprintf("zstack/v1/policy-routes/tables/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPolicyRouteTableFailed, err.Error())
	}
	var result model.QueryPolicyRouteTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询策略路由表与单节点云路由关联
func QueryPolicyRouteTableVRouterRef(params model.QueryPolicyRouteTableVRouterRefeRequest) models.Result {
	//GET zstack/v1/policy-routes/tables/vrouters/refs
	url := fmt.Sprintf("zstack/v1/policy-routes/tables/vrouters/refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPolicyRouteTableVRouterRefFailed, err.Error())
	}
	var result model.QueryPolicyRouteTableVRouterRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建策略路由
func CreatePolicyRouteTableRouteEntry(params model.CreatePolicyRouteTableRouteEntryRequest) models.Result {
	//POST zstack/v1/policy-routes/routes
	url := fmt.Sprintf("zstack/v1/policy-routes/routes")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreatePolicyRouteTableRouteEntryFailed, err.Error())
	}
	var result model.CreatePolicyRouteTableRouteEntryResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除策略路由
func DeletePolicyRouteTableRouteEntry(params model.DeletePolicyRouteTableRouteEntryRequest) models.Result {
	//DELETE zstack/v1/policy-routes/routes/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/policy-routes/routes/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeletePolicyRouteTableRouteEntryFailed, err.Error())
	}
	var result model.DeletePolicyRouteTableRouteEntryResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询策略路由
func QueryPolicyRouteTableRouteEntry(params model.QueryPolicyRouteTableRouteEntryRequest) models.Result {
	//GET zstack/v1/policy-routes/routes
	//GET zstack/v1/policy-routes/routes/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/policy-routes/routes")
	} else {
		url = fmt.Sprintf("zstack/v1/policy-routes/routes/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPolicyRouteTableRouteEntryFailed, err.Error())
	}
	var result model.QueryPolicyRouteTableRouteEntryResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询策略路由规则
func QueryPolicyRouteRule(params model.QueryPolicyRouteRuleRequest) models.Result {
	//GET zstack/v1/policy-routes/rules
	//GET zstack/v1/policy-routes/rules/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/policy-routes/rules")
	} else {
		url = fmt.Sprintf("zstack/v1/policy-routes/rules/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPolicyRouteRuleFailed, err.Error())
	}
	var result model.QueryPolicyRouteRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//网络加载策略路由规则集
func AttachPolicyRouteRuleSetToL3(params model.AttachPolicyRouteRuleSetToL3Request) models.Result {
	//POST zstack/v1/policy-routes/rulesets/{ruleSetUuid}/l3networks/{l3Uuid}
	url := fmt.Sprintf("zstack/v1/policy-routes/rulesets/%s/l3networks/%s", params.RuleSetUuid, params.L3Uuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachPolicyRouteRuleSetToL3Failed, err.Error())
	}
	var result model.AttachPolicyRouteRuleSetToL3Response
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//网络解绑策略路由
func DetachPolicyRouteRuleSetFromL3(params model.DetachPolicyRouteRuleSetFromL3Request) models.Result {
	//DELETE zstack/v1/policy-routes/rulesets/{ruleSetUuid}/l3networks/{l3Uuid}
	url := fmt.Sprintf("zstack/v1/policy-routes/rulesets/%s/l3networks/%s", params.RuleSetUuid, params.L3Uuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachPolicyRouteRuleSetFromL3Failed, err.Error())
	}
	var result model.DetachPolicyRouteRuleSetFromL3Response
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取虚拟路由器的策略路由集合
func GetPolicyRouteRuleSetFromVirtualRouter(params model.GetPolicyRouteRuleSetFromVirtualRouterRequest) models.Result {
	//GET zstack/v1/policy-routes/rulesets/virtualrouter/{vmInstanceUuid}
	url := fmt.Sprintf("zstack/v1/policy-routes/rulesets/virtualrouter/%s", params.VmInstanceUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetPolicyRouteRuleSetFromVirtualRouterFailed, err.Error())
	}
	var result model.GetPolicyRouteRuleSetFromVirtualRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
