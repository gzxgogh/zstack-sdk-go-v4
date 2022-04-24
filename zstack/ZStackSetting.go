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

//查询全局设置
func QueryGlobalConfig(params model.QueryGlobalConfigRequest) mgresult.Result {
	//GET zstack/v1/global-configurations
	url := fmt.Sprintf("zstack/v1/global-configurations")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryGlobalConfigFailed, err)
	}
	var respResult model.QueryGlobalConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//更新全局设置
func UpdateGlobalConfig(params model.UpdateGlobalConfigRequest) mgresult.Result {
	//PUT zstack/v1/global-configurations/{category}/{name}/actions
	url := fmt.Sprintf("zstack/v1/global-configurations/%s/%s/actions", params.Category, params.Name)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateGlobalConfigFailed, err)
	}
	var respResult model.UpdateGlobalConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//重置全局配置
func ResetGlobalConfig(params model.ResetGlobalConfigRequest) mgresult.Result {
	//PUT zstack/v1/global-configurations/actions
	url := fmt.Sprintf("zstack/v1/global-configurations/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ResetGlobalConfigFailed, err)
	}
	var respResult model.ResetGlobalConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//获取资源的资源高级设置
func GetResourceConfig(params model.GetResourceConfigRequest) mgresult.Result {
	//GET zstack/v1/resource-configurations/{resourceUuid}/{category}/{name}
	url := fmt.Sprintf("zstack/v1/resource-configurations/%s/%s/%s", params.ResourceUuid, params.Category, params.Name)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetResourceConfigFailed, err)
	}
	var respResult model.GetResourceConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//获取可配置的资源高级设置
func GetResourceBindableConfig(params model.GetResourceBindableConfigRequest) mgresult.Result {
	//GET zstack/v1/resource-configurations/bindable
	//GET zstack/v1/resource-configurations/bindable/{category}
	var url string
	if params.Category == "" {
		url = fmt.Sprintf("zstack/v1/resource-configurations/bindable")
	} else {
		url = fmt.Sprintf("zstack/v1/resource-configurations/bindable/%s", params.Category)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetResourceBindableConfigFailed, err)
	}
	var respResult model.GetResourceBindableConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//删除资源的资源高级设置
func DeleteResourceConfig(params model.DeleteResourceConfigRequest) mgresult.Result {
	//DELETE zstack/v1/resource-configurations/{category}/{name}/{resourceUuid}
	url := fmt.Sprintf("zstack/v1/resource-configurations/%s/%s/%s", params.Category, params.Name, params.ResourceUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteResourceConfigFailed, err)
	}
	var respResult model.DeleteResourceConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//查询资源高级设置
func QueryResourceConfig(params model.QueryResourceConfigRequest) mgresult.Result {
	//GET zstack/v1/resource-configurations
	url := fmt.Sprintf("zstack/v1/resource-configurations")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryResourceConfigFailed, err)
	}
	var respResult model.QueryResourceConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//更新资源高级设置
func UpdateResourceConfig(params model.UpdateResourceConfigRequest) mgresult.Result {
	//PUT zstack/v1/resource-configurations/{category}/{name}/{resourceUuid}/actions
	url := fmt.Sprintf("zstack/v1/resource-configurations/%s/%s/%s/actions", params.Category, params.Name, params.ResourceUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateResourceConfigFailed, err)
	}
	var respResult model.UpdateResourceConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//应用模板
func ApplyTemplateConfig(params model.ApplyTemplateConfigRequest) mgresult.Result {
	//PUT zstack/v1/template-configurations/{templateUuid}/actions
	url := fmt.Sprintf("zstack/v1/template-configurations/%s/actions", params.TemplateUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ApplyTemplateConfigFailed, err)
	}
	var respResult model.ApplyTemplateConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//查询所有模板信息
func QueryGlobalConfigTemplate(params model.QueryGlobalConfigTemplateRequest) mgresult.Result {
	//GET zstack/v1/template-configurations/templates
	//GET zstack/v1/template-configurations/templates/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/template-configurations/templates")
	} else {
		url = fmt.Sprintf("zstack/v1/template-configurations/templates/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryGlobalConfigTemplateFailed, err)
	}
	var respResult model.QueryGlobalConfigTemplateResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//查询模板具体配置
func QueryTemplateConfig(params model.QueryTemplateConfigRequest) mgresult.Result {
	//GET zstack/v1/template-configurations/configs
	//GET zstack/v1/template-configurations/configs/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/template-configurations/configs")
	} else {
		url = fmt.Sprintf("zstack/v1/template-configurations/configs/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryTemplateConfigFailed, err)
	}
	var respResult model.QueryTemplateConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//更新模板值
func UpdateTemplateConfig(params model.UpdateTemplateConfigRequest) mgresult.Result {
	//PUT zstack/v1/template-configurations/{templateUuid}/actions
	url := fmt.Sprintf("zstack/v1/template-configurations/%s/actions", params.TemplateUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateTemplateConfigFailed, err)
	}
	var respResult model.UpdateTemplateConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//重置模板到初始状态
func ResetTemplateConfig(params model.ResetTemplateConfigRequest) mgresult.Result {
	//PUT zstack/v1/template-configurations/{templateUuid}/actions
	url := fmt.Sprintf("zstack/v1/template-configurations/%s/actions", params.TemplateUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ResetTemplateConfigFailed, err)
	}
	var respResult model.ResetTemplateConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}
