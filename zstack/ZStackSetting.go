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

//查询全局设置
func QueryGlobalConfig(params model.QueryGlobalConfigRequest) models.Result {
	//GET zstack/v1/global-configurations
	url := fmt.Sprintf("zstack/v1/global-configurations")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryGlobalConfigFailed, err.Error())
	}
	var respResult model.QueryGlobalConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//更新全局设置
func UpdateGlobalConfig(params model.UpdateGlobalConfigRequest) models.Result {
	//PUT zstack/v1/global-configurations/{category}/{name}/actions
	url := fmt.Sprintf("zstack/v1/global-configurations/%s/%s/actions", params.Category, params.Name)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateGlobalConfigFailed, err.Error())
	}
	var respResult model.UpdateGlobalConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//重置全局配置
func ResetGlobalConfig(params model.ResetGlobalConfigRequest) models.Result {
	//PUT zstack/v1/global-configurations/actions
	url := fmt.Sprintf("zstack/v1/global-configurations/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ResetGlobalConfigFailed, err.Error())
	}
	var respResult model.ResetGlobalConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取资源的资源高级设置
func GetResourceConfig(params model.GetResourceConfigRequest) models.Result {
	//GET zstack/v1/resource-configurations/{resourceUuid}/{category}/{name}
	url := fmt.Sprintf("zstack/v1/resource-configurations/%s/%s/%s", params.ResourceUuid, params.Category, params.Name)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetResourceConfigFailed, err.Error())
	}
	var respResult model.GetResourceConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取可配置的资源高级设置
func GetResourceBindableConfig(params model.GetResourceBindableConfigRequest) models.Result {
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
		return models.Error(errcode.GetResourceBindableConfigFailed, err.Error())
	}
	var respResult model.GetResourceBindableConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//删除资源的资源高级设置
func DeleteResourceConfig(params model.DeleteResourceConfigRequest) models.Result {
	//DELETE zstack/v1/resource-configurations/{category}/{name}/{resourceUuid}
	url := fmt.Sprintf("zstack/v1/resource-configurations/%s/%s/%s", params.Category, params.Name, params.ResourceUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteResourceConfigFailed, err.Error())
	}
	var respResult model.DeleteResourceConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询资源高级设置
func QueryResourceConfig(params model.QueryResourceConfigRequest) models.Result {
	//GET zstack/v1/resource-configurations
	url := fmt.Sprintf("zstack/v1/resource-configurations")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryResourceConfigFailed, err.Error())
	}
	var respResult model.QueryResourceConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//更新资源高级设置
func UpdateResourceConfig(params model.UpdateResourceConfigRequest) models.Result {
	//PUT zstack/v1/resource-configurations/{category}/{name}/{resourceUuid}/actions
	url := fmt.Sprintf("zstack/v1/resource-configurations/%s/%s/%s/actions", params.Category, params.Name, params.ResourceUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateResourceConfigFailed, err.Error())
	}
	var respResult model.UpdateResourceConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//应用模板
func ApplyTemplateConfig(params model.ApplyTemplateConfigRequest) models.Result {
	//PUT zstack/v1/template-configurations/{templateUuid}/actions
	url := fmt.Sprintf("zstack/v1/template-configurations/%s/actions", params.TemplateUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ApplyTemplateConfigFailed, err.Error())
	}
	var respResult model.ApplyTemplateConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询所有模板信息
func QueryGlobalConfigTemplate(params model.QueryGlobalConfigTemplateRequest) models.Result {
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
		return models.Error(errcode.QueryGlobalConfigTemplateFailed, err.Error())
	}
	var respResult model.QueryGlobalConfigTemplateResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询模板具体配置
func QueryTemplateConfig(params model.QueryTemplateConfigRequest) models.Result {
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
		return models.Error(errcode.QueryTemplateConfigFailed, err.Error())
	}
	var respResult model.QueryTemplateConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//更新模板值
func UpdateTemplateConfig(params model.UpdateTemplateConfigRequest) models.Result {
	//PUT zstack/v1/template-configurations/{templateUuid}/actions
	url := fmt.Sprintf("zstack/v1/template-configurations/%s/actions", params.TemplateUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateTemplateConfigFailed, err.Error())
	}
	var respResult model.UpdateTemplateConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//重置模板到初始状态
func ResetTemplateConfig(params model.ResetTemplateConfigRequest) models.Result {
	//PUT zstack/v1/template-configurations/{templateUuid}/actions
	url := fmt.Sprintf("zstack/v1/template-configurations/%s/actions", params.TemplateUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ResetTemplateConfigFailed, err.Error())
	}
	var respResult model.ResetTemplateConfigResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}
