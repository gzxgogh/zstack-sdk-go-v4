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

//添加资源栈模板
func AddStackTemplate(params model.AddStackTemplateRequest) mgresult.Result {
	//POST zstack/v1/cloudformation/template
	url := fmt.Sprintf("zstack/v1/cloudformation/template")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddStackTemplateFailed, err)
	}
	var result model.AddStackTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除资源栈模板
func DeleteStackTemplate(params model.DeleteStackTemplateRequest) mgresult.Result {
	//DELETE zstack/v1/cloudformation/template/{uuid}
	url := fmt.Sprintf("zstack/v1/cloudformation/template/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteStackTemplateFailed, err)
	}
	var result model.DeleteStackTemplateRequest
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询资源栈模板
func QueryStackTemplate(params model.QueryStackTemplateRequest) mgresult.Result {
	//GET zstack/v1/cloudformation/template
	//GET zstack/v1/cloudformation/template/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/cloudformation/template")
	} else {
		url = fmt.Sprintf("zstack/v1/cloudformation/template/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryStackTemplateFailed, err)
	}
	var result model.QueryStackTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//修改资源栈模板
func UpdateStackTemplate(params model.UpdateStackTemplateRequest) mgresult.Result {
	//PUT zstack/v1/cloudformation/template/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/cloudformation/template/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateStackTemplateFailed, err)
	}
	var result model.UpdateStackTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取支持的资源列表
func GetSupportedCloudFormationResources(params model.GetSupportedCloudFormationResourcesRequest) mgresult.Result {
	//GET zstack/v1/cloudformation/resources
	url := fmt.Sprintf("zstack/v1/cloudformation/resources")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetSupportedCloudFormationResourcesFailed, err)
	}
	var result model.GetSupportedCloudFormationResourcesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建资源栈
func CreateResourceStack(params model.CreateResourceStackRequest) mgresult.Result {
	//POST zstack/v1/cloudformation/stack
	url := fmt.Sprintf("zstack/v1/cloudformation/stack")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateResourceStackFailed, err)
	}
	var result model.CreateResourceStackResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//预览资源栈
func PreviewResourceStack(params model.PreviewResourceStackRequest) mgresult.Result {
	//POST zstack/v1/cloudformation/stack/preview
	url := fmt.Sprintf("zstack/v1/cloudformation/stack/preview")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.PreviewResourceStackFailed, err)
	}
	var result model.PreviewResourceStackResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除资源栈
func DeleteResourceStack(params model.DeleteResourceStackRequest) mgresult.Result {
	//DELETE zstack/v1/cloudformation/stack/{uuid}
	url := fmt.Sprintf("zstack/v1/cloudformation/stack/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteResourceStackFailed, err)
	}
	var result model.DeleteResourceStackResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//修改资源栈
func UpdateResourceStack(params model.UpdateResourceStackRequest) mgresult.Result {
	//PUT zstack/v1/cloudformation/stack/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/cloudformation/stack/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateResourceStackFailed, err)
	}
	var result model.UpdateResourceStackResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询资源栈
func QueryResourceStack(params model.QueryResourceStackRequest) mgresult.Result {
	//GET zstack/v1/cloudformation/stack
	//GET zstack/v1/cloudformation/stack/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/cloudformation/stack")
	} else {
		url = fmt.Sprintf("zstack/v1/cloudformation/stack/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryResourceStackFailed, err)
	}
	var result model.QueryResourceStackResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取资源栈内资源列表
func GetResourceFromResourceStack(params model.GetResourceFromResourceStackRequest) mgresult.Result {
	//GET zstack/v1/cloudformation/stack/resources
	url := fmt.Sprintf("zstack/v1/cloudformation/stack/resources")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetResourceFromResourceStackFailed, err)
	}
	var result model.GetResourceFromResourceStackResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询资源栈内事件列表
func QueryEventFromResourceStack(params model.QueryEventFromResourceStackRequest) mgresult.Result {
	//GET zstack/v1/cloudformation/event
	//GET zstack/v1/cloudformation/event/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/cloudformation/event")
	} else {
		url = fmt.Sprintf("zstack/v1/cloudformation/event/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryEventFromResourceStackFailed, err)
	}
	var result model.QueryEventFromResourceStackResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//重启资源栈
func RestartResourceStackRequestStack(params model.RestartResourceStackRequest) mgresult.Result {
	//PUT zstack/v1/cloudformation/stack/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/cloudformation/stack/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RestartResourceStackRequestStackFailed, err)
	}
	var result model.RestartResourceStackResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//检查资源栈模板参数
func CheckStackTemplateParameters(params model.CheckStackTemplateParametersRequest) mgresult.Result {
	//POST zstack/v1/cloudformation/stack/check
	url := fmt.Sprintf("zstack/v1/cloudformation/stack/check")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CheckStackTemplateParametersFailed, err)
	}
	var result model.CheckStackTemplateParametersResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从资源编排模板解析成资源关系图
func DecodeStackTemplate(params model.DecodeStackTemplateRequest) mgresult.Result {
	//POST zstack/v1/cloudformation/stack/preview/resource
	url := fmt.Sprintf("zstack/v1/cloudformation/stack/preview/resource")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DecodeStackTemplateFailed, err)
	}
	var result model.DecodeStackTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取资源对应的资源栈
func GetResourceStackFromResource(params model.GetResourceStackFromResourceRequest) mgresult.Result {
	//GET zstack/v1/cloudformation/resources/stack
	url := fmt.Sprintf("zstack/v1/cloudformation/resources/stack")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetResourceStackFromResourceFailed, err)
	}
	var result model.GetResourceStackFromResourceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取资源栈中云主机端口监控状态
func GetResourceStackVmStatus(params model.GetResourceStackVmStatusRequest) mgresult.Result {
	//GET zstack/v1/cloudformation/stack/monitor/vmstatus
	url := fmt.Sprintf("zstack/v1/cloudformation/stack/monitor/vmstatus")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetResourceStackVmStatusFailed, err)
	}
	var result model.GetResourceStackVmStatusResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加资源栈中云主机端口监控
func AddResourceStackVmPortMonitor(params model.AddResourceStackVmPortMonitorRequest) mgresult.Result {
	//POST zstack/v1/cloudformation/stack/monitor/addvm
	url := fmt.Sprintf("zstack/v1/cloudformation/stack/monitor/addvm")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddResourceStackVmPortMonitorFailed, err)
	}
	var result model.AddResourceStackVmPortMonitorResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除资源栈中云主机端口监控
func DeleteResourceStackVmPortMonitor(params model.DeleteResourceStackVmPortMonitorRequest) mgresult.Result {
	//DELETE zstack/v1/cloudformation/stack/monitor/delvm
	url := fmt.Sprintf("zstack/v1/cloudformation/stack/monitor/delvm")

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteResourceStackVmPortMonitorFailed, err)
	}
	var result model.DeleteResourceStackVmPortMonitorResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
