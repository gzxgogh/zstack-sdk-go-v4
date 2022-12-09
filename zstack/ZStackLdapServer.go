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

//添加AD/LDAP服务器
func AddLdapServer(params model.AddLdapServerRequest) models.Result {
	//POST zstack/v1/ldap/servers
	url := fmt.Sprintf("zstack/v1/ldap/servers")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddLdapServerFailed, err.Error())
	}
	var respResult model.AddLdapServerResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//删除AD/LDAP服务器
func DeleteLdapServer(params model.DeleteLdapServerRequest) models.Result {
	//DELETE zstack/v1/ldap/servers/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/ldap/servers/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteLdapServerFailed, err.Error())
	}
	var respResult model.DeleteLdapServerResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询AD/LDAP服务器
func QueryLdapServer(params model.QueryLdapServerRequest) models.Result {
	//GET zstack/v1/ldap/servers
	//GET zstack/v1/ldap/servers/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/ldap/servers")
	} else {
		url = fmt.Sprintf("zstack/v1/ldap/servers/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryLdapServerFailed, err.Error())
	}
	var respResult model.QueryLdapServerResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//更新AD/LDAP服务器
func UpdateLdapServer(params model.UpdateLdapServerRequest) models.Result {
	//PUT zstack/v1/ldap/servers/{ldapServerUuid}
	url := fmt.Sprintf("zstack/v1/ldap/servers/%s", params.LdapServerUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateLdapServerFailed, err.Error())
	}
	var respResult model.UpdateLdapServerResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//创建AD/LDAP绑定
func CreateLdapBinding(params model.CreateLdapBindingRequest) models.Result {
	//POST zstack/v1/ldap/bindings
	url := fmt.Sprintf("zstack/v1/ldap/bindings")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateLdapBindingFailed, err.Error())
	}
	var respResult model.CreateLdapBindingResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//删除AD/LDAP绑定
func DeleteLdapBinding(params model.DeleteLdapBindingRequest) models.Result {
	//DELETE zstack/v1/ldap/bindings/{uuid}
	url := fmt.Sprintf("zstack/v1/ldap/bindings/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteLdapBindingFailed, err.Error())
	}
	var respResult model.DeleteLdapBindingResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询AD/LDAP绑定
func QueryLdapBinding(params model.QueryLdapBindingRequest) models.Result {
	//GET zstack/v1/ldap/bindings
	//GET zstack/v1/ldap/bindings/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/ldap/bindings")
	} else {
		url = fmt.Sprintf("zstack/v1/ldap/bindings/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryLdapBindingFailed, err.Error())
	}
	var respResult model.QueryLdapBindingResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//清理无效的AD/LDAP绑定
func CleanInvalidLdapBinding(params model.CleanInvalidLdapBindingRequest) models.Result {
	//PUT zstack/v1/ldap/bindings/actions
	url := fmt.Sprintf("zstack/v1/ldap/bindings/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.CleanInvalidLdapBindingFailed, err.Error())
	}
	var respResult model.CleanInvalidLdapBindingResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//使用AD/LDAP身份登录
func LogInByLdap(params model.LogInByLdapRequest) models.Result {
	//PUT zstack/v1/ldap/login
	url := fmt.Sprintf("zstack/v1/ldap/login")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.LogInByLdapFailed, err.Error())
	}
	var respResult model.LogInByLdapResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取AD/LDAP条目
func GetLdapEntry(params model.GetLdapEntryRequest) models.Result {
	//GET zstack/v1/ldap/entry
	url := fmt.Sprintf("zstack/v1/ldap/entry")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetLdapEntryFailed, err.Error())
	}
	var respResult model.GetLdapEntryResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取可绑定的AD/LDAP条目
func GetCandidateLdapEntryForBinding(params model.GetCandidateLdapEntryForBindingRequest) models.Result {
	//GET zstack/v1/ldap/entries/candidates
	url := fmt.Sprintf("zstack/v1/ldap/entries/candidates")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetCandidateLdapEntryForBindingFailed, err.Error())
	}
	var respResult model.GetCandidateLdapEntryForBindingResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}
