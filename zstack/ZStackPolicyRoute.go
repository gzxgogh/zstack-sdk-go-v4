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

//创建策略路由规则集
func CreatePolicyRouteRuleSet(params model.CreatePolicyRouteRuleSetRequest) mgresult.Result {
	//POST zstack/v1/policy-routes/rulesets
	url := fmt.Sprintf("zstack/v1/policy-routes/rulesets")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreatePolicyRouteRuleSetFailed, err)
	}
	var result model.CreatePolicyRouteRuleSetRequest
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除策略路由规则集
func DeletePolicyRouteRuleSet(params model.DeletePolicyRouteRuleSetRequest) mgresult.Result {
	//DELETE zstack/v1/policy-routes/ruleSets/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/policy-routes/ruleSets/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeletePolicyRouteRuleSetFailed, err)
	}
	var result model.DeletePolicyRouteRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新策略路由规则集属性
func UpdatePolicyRouteRuleSet(params model.UpdatePolicyRouteRuleSetRequest) mgresult.Result {
	//PUT zstack/v1/policy-routes/ruleSets/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/policy-routes/ruleSets/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdatePolicyRouteRuleSetFailed, err)
	}
	var result model.UpdatePolicyRouteRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询策略路由规则集
func QueryPolicyRouteRuleSet(params model.QueryPolicyRouteRuleSetRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryPolicyRouteRuleSetFailed, err)
	}
	var result model.QueryPolicyRouteRuleSetResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询策略路由规则集网络关联
func QueryPolicyRouteRuleSetL3Ref(params model.QueryPolicyRouteRuleSetL3RefRequest) mgresult.Result {
	//GET zstack/v1/policy-routes/rulesets/l3networdks/refs
	url := fmt.Sprintf("zstack/v1/policy-routes/rulesets/l3networdks/refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryPolicyRouteRuleSetL3RefFailed, err)
	}
	var result model.QueryPolicyRouteRuleSetL3RefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询策略路由规则集与单节点云路由关联
func QueryPolicyRouteRuleSetVRouterRef(params model.QueryPolicyRouteRuleSetVRouterRefRequest) mgresult.Result {
	//GET zstack/v1/policy-routes/rulesets/vrouters/refs
	url := fmt.Sprintf("zstack/v1/policy-routes/rulesets/vrouters/refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryPolicyRouteRuleSetVRouterRefFailed, err)
	}
	var result model.QueryPolicyRouteRuleSetVRouterRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建策略路由规则集规则
func CreatePolicyRouteRule(params model.CreatePolicyRouteRuleRequest) mgresult.Result {
	//POST zstack/v1/policy-routes/rules
	url := fmt.Sprintf("zstack/v1/policy-routes/rules")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreatePolicyRouteRuleFailed, err)
	}
	var result model.CreatePolicyRouteRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除策略路由规则集规则
func DeletePolicyRouteRule(params model.DeletePolicyRouteRuleRequest) mgresult.Result {
	//DELETE zstack/v1/policy-routes/rules/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/policy-routes/rules/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeletePolicyRouteRuleFailed, err)
	}
	var result model.DeletePolicyRouteRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建策略路由表
func CreatePolicyRouteTable(params model.CreatePolicyRouteTableRequest) mgresult.Result {
	//POST zstack/v1/policy-routes/tables
	url := fmt.Sprintf("zstack/v1/policy-routes/tables")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreatePolicyRouteTableFailed, err)
	}
	var result model.CreatePolicyRouteTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除策略路由表
func DeletePolicyRouteTable(params model.DeletePolicyRouteTableRequest) mgresult.Result {
	//DELETE zstack/v1/policy-routes/tables/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/policy-routes/tables/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeletePolicyRouteTableFailed, err)
	}
	var result model.DeletePolicyRouteTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询策略路由表
func QueryPolicyRouteTable(params model.QueryPolicyRouteTableRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryPolicyRouteTableFailed, err)
	}
	var result model.QueryPolicyRouteTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询策略路由表与单节点云路由关联
func QueryPolicyRouteTableVRouterRef(params model.QueryPolicyRouteTableVRouterRefeRequest) mgresult.Result {
	//GET zstack/v1/policy-routes/tables/vrouters/refs
	url := fmt.Sprintf("zstack/v1/policy-routes/tables/vrouters/refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryPolicyRouteTableVRouterRefFailed, err)
	}
	var result model.QueryPolicyRouteTableVRouterRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建策略路由
func CreatePolicyRouteTableRouteEntry(params model.CreatePolicyRouteTableRouteEntryRequest) mgresult.Result {
	//POST zstack/v1/policy-routes/routes
	url := fmt.Sprintf("zstack/v1/policy-routes/routes")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreatePolicyRouteTableRouteEntryFailed, err)
	}
	var result model.CreatePolicyRouteTableRouteEntryResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除策略路由
func DeletePolicyRouteTableRouteEntry(params model.DeletePolicyRouteTableRouteEntryRequest) mgresult.Result {
	//DELETE zstack/v1/policy-routes/routes/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/policy-routes/routes/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeletePolicyRouteTableRouteEntryFailed, err)
	}
	var result model.DeletePolicyRouteTableRouteEntryResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询策略路由
func QueryPolicyRouteTableRouteEntry(params model.QueryPolicyRouteTableRouteEntryRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryPolicyRouteTableRouteEntryFailed, err)
	}
	var result model.QueryPolicyRouteTableRouteEntryResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询策略路由规则
func QueryPolicyRouteRule(params model.QueryPolicyRouteRuleRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryPolicyRouteRuleFailed, err)
	}
	var result model.QueryPolicyRouteRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//网络加载策略路由规则集
func AttachPolicyRouteRuleSetToL3(params model.AttachPolicyRouteRuleSetToL3Request) mgresult.Result {
	//POST zstack/v1/policy-routes/rulesets/{ruleSetUuid}/l3networks/{l3Uuid}
	url := fmt.Sprintf("zstack/v1/policy-routes/rulesets/%s/l3networks/%s", params.RuleSetUuid, params.L3Uuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachPolicyRouteRuleSetToL3Failed, err)
	}
	var result model.AttachPolicyRouteRuleSetToL3Response
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//网络解绑策略路由
func DetachPolicyRouteRuleSetFromL3(params model.DetachPolicyRouteRuleSetFromL3Request) mgresult.Result {
	//DELETE zstack/v1/policy-routes/rulesets/{ruleSetUuid}/l3networks/{l3Uuid}
	url := fmt.Sprintf("zstack/v1/policy-routes/rulesets/%s/l3networks/%s", params.RuleSetUuid, params.L3Uuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachPolicyRouteRuleSetFromL3Failed, err)
	}
	var result model.DetachPolicyRouteRuleSetFromL3Response
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取虚拟路由器的策略路由集合
func GetPolicyRouteRuleSetFromVirtualRouter(params model.GetPolicyRouteRuleSetFromVirtualRouterRequest) mgresult.Result {
	//GET zstack/v1/policy-routes/rulesets/virtualrouter/{vmInstanceUuid}
	url := fmt.Sprintf("zstack/v1/policy-routes/rulesets/virtualrouter/%s", params.VmInstanceUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetPolicyRouteRuleSetFromVirtualRouterFailed, err)
	}
	var result model.GetPolicyRouteRuleSetFromVirtualRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
