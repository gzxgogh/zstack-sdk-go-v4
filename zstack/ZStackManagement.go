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

//请求控制台访问地址
func RequestConsoleAccess(params model.RequestConsoleAccessRequest) models.Result {
	//POST zstack/v1/consoles
	url := fmt.Sprintf("zstack/v1/consoles")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.RequestConsoleAccessFailed, err.Error())
	}
	var respResult model.RequestConsoleAccessResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询控制台代理
func QueryConsoleProxyAgent(params model.QueryConsoleProxyAgentRequest) models.Result {
	//GET zstack/v1/consoles/agents
	//GET zstack/v1/consoles/agents/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/consoles/agents")
	} else {
		url = fmt.Sprintf("zstack/v1/consoles/agents/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryConsoleProxyAgentFailed, err.Error())
	}
	var respResult model.QueryConsoleProxyAgentResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//重连控制台代理
func ReconnectConsoleProxyAgent(params model.ReconnectConsoleProxyAgentRequest) models.Result {
	//PUT zstack/v1/consoles/agents
	url := fmt.Sprintf("zstack/v1/consoles/agents")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ReconnectConsoleProxyAgentFailed, err.Error())
	}
	var respResult model.ReconnectConsoleProxyAgentResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//更新控制台代理
func UpdateConsoleProxyAgent(params model.UpdateConsoleProxyAgentRequest) models.Result {
	//PUT zstack/v1/consoles/agents/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/consoles/agents/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateConsoleProxyAgentFailed, err.Error())
	}
	var respResult model.UpdateConsoleProxyAgentResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//增加IP访问控制规则
func AddAccessControlRule(params model.AddAccessControlRuleRequest) models.Result {
	//POST zstack/v1/login-control/access-control/rules
	url := fmt.Sprintf("zstack/v1/login-control/access-control/rules")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddAccessControlRuleFailed, err.Error())
	}
	var respResult model.AddAccessControlRuleResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//删除IP访问控制规则
func DeleteAccessControlRule(params model.DeleteAccessControlRuleRequest) models.Result {
	//DELETE zstack/v1/login-control/access-control/rules/{uuid}
	url := fmt.Sprintf("zstack/v1/login-control/access-control/rules/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteAccessControlRuleFailed, err.Error())
	}
	var respResult model.DeleteAccessControlRuleResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//更新IP访问控制规则
func UpdateAccessControlRule(params model.UpdateAccessControlRuleRequest) models.Result {
	//PUT zstack/v1/login-control/access-control/rules/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/login-control/access-control/rules/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateAccessControlRuleFailed, err.Error())
	}
	var respResult model.UpdateAccessControlRuleResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询IP访问控制规则
func QueryAccessControlRule(params model.QueryAccessControlRuleRequest) models.Result {
	//GET zstack/v1/login-control/access-control/rules
	//GET zstack/v1/login-control/access-control/rules/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/login-control/access-control/rules")
	} else {
		url = fmt.Sprintf("zstack/v1/login-control/access-control/rules/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAccessControlRuleFailed, err.Error())
	}
	var respResult model.QueryAccessControlRuleResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//创建AccessKey
func CreateAccessKey(params model.CreateAccessKeyRequest) models.Result {
	//POST zstack/v1/accesskeys
	url := fmt.Sprintf("zstack/v1/accesskeys")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateAccessKeyFailed, err.Error())
	}
	var respResult model.CreateAccessKeyResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//删除AccessKey
func DeleteAccessKey(params model.DeleteAccessKeyRequest) models.Result {
	//DELETE zstack/v1/accesskeys/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/accesskeys/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteAccessKeyFailed, err.Error())
	}
	var respResult model.DeleteAccessKeyResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询AccessKey
func QueryAccessKey(params model.QueryAccessKeyRequest) models.Result {
	//GET zstack/v1/accesskeys
	//GET zstack/v1/accesskeys/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/accesskeys")
	} else {
		url = fmt.Sprintf("zstack/v1/accesskeys/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAccessKeyFailed, err.Error())
	}
	var respResult model.QueryAccessKeyResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//开启或关闭AccessKey
func ChangeAccessKeyState(params model.ChangeAccessKeyStateRequest) models.Result {
	//PUT zstack/v1/accesskeys/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/accesskeys/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeAccessKeyStateFailed, err.Error())
	}
	var respResult model.ChangeAccessKeyStateResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//添加日志服务器配置
func AddLogConfiguration(params model.AddLogConfigurationRequest) models.Result {
	//POST zstack/v1/log/configurations
	url := fmt.Sprintf("zstack/v1/log/configurations")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddLogConfigurationFailed, err.Error())
	}
	var respResult model.AddLogConfigurationResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//删除日志服务器配置
func DeleteLogConfiguration(params model.DeleteLogConfigurationRequest) models.Result {
	//DELETE zstack/v1/log/configurations/log4j2
	url := fmt.Sprintf("zstack/v1/log/configurations/log4j2")

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteLogConfigurationFailed, err.Error())
	}
	var respResult model.DeleteLogConfigurationResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取日志服务器配置
func GetLogConfiguration(params model.GetLogConfigurationRequest) models.Result {
	//GET zstack/v1/log/configurations
	//GET zstack/v1/log/configurations/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/log/configurations")
	} else {
		url = fmt.Sprintf("zstack/v1/log/configurations/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetLogConfigurationFailed, err.Error())
	}
	var respResult model.GetLogConfigurationResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//修改日志服务器配置
func UpdateLogConfiguration(params model.UpdateLogConfigurationRequest) models.Result {
	//PUT zstack/v1/log/configurations
	url := fmt.Sprintf("zstack/v1/log/configurations")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateLogConfigurationFailed, err.Error())
	}
	var respResult model.UpdateLogConfigurationResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}
