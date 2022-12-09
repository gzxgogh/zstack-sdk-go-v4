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

//创建账户
func CreateAccount(params model.CreateAccountRequest) models.Result {
	//POST zstack/v1/accounts
	url := fmt.Sprintf("zstack/v1/accounts")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateAccountFailed, err.Error())
	}
	var result model.CreateAccountResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除账户
func DeleteAccount(params model.DeleteAccountRequest) models.Result {
	//DELETE zstack/v1/accounts/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/accounts/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteAccountFailed, err.Error())
	}
	var result model.DeleteAccountResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询账户
func QueryAccount(params model.QueryAccountRequest) models.Result {
	//GET zstack/v1/accounts
	//GET zstack/v1/accounts/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/accounts")
	} else {
		url = fmt.Sprintf("zstack/v1/accounts/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAccountFailed, err.Error())
	}
	var result model.QueryAccountResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新账户
func UpdateAccount(params model.UpdateAccountRequest) models.Result {
	//PUT zstack/v1/accounts/{uuid}
	url := fmt.Sprintf("zstack/v1/accounts/%s", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateAccountFailed, err.Error())
	}
	var result model.UpdateAccountResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//使用账户身份登录
func LogInByAccount(params model.LogInByAccountRequest) models.Result {
	//PUT zstack/v1/accounts/login
	url := fmt.Sprintf("zstack/v1/accounts/login")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.LogInByAccountFailed, err.Error())
	}
	var result model.LogInByAccountResponse

	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取登录验证码
func GetLoginCaptcha(params model.GetLoginCaptchaRequest) models.Result {
	//GET zstack/v1/login/control/captcha
	url := fmt.Sprintf("zstack/v1/login/control/captcha")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetLoginCaptchaFailed, err.Error())
	}
	var result model.GetLoginCaptchaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//刷新验证码
func RefreshCaptcha(params model.RefreshCaptchaRequest) models.Result {
	//GET zstack/v1/captcha/refresh
	url := fmt.Sprintf("zstack/v1/captcha/refresh")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.RefreshCaptchaFailed, err.Error())
	}
	var result model.RefreshCaptchaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取双因子认证密匙
func GetTwoFactorAuthenticationSecret(params model.GetTwoFactorAuthenticationSecretRequest) models.Result {
	//GET zstack/v1/twofactorauthentication/secret
	url := fmt.Sprintf("zstack/v1/twofactorauthentication/secret")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetTwoFactorAuthenticationSecretFailed, err.Error())
	}
	var result model.GetTwoFactorAuthenticationSecretResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询双因子认证密匙
func QueryTwoFactorAuthentication(params model.QueryTwoFactorAuthenticationRequest) models.Result {
	//GET zstack/v1/twofactorauthentication/secrets
	//GET zstack/v1/twofactorauthentication/secrets/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/twofactorauthentication/secrets")
	} else {
		url = fmt.Sprintf("zstack/v1/twofactorauthentication/secrets/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryTwoFactorAuthenticationFailed, err.Error())
	}
	var result model.QueryTwoFactorAuthenticationResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取双因子认证状态
func GetTwoFactorAuthenticationState(params model.GetTwoFactorAuthenticationStateRequest) models.Result {
	//GET zstack/v1/twofactorauthentication/state
	url := fmt.Sprintf("zstack/v1/twofactorauthentication/state")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetTwoFactorAuthenticationStateFailed, err.Error())
	}
	var result model.GetTwoFactorAuthenticationStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取账户配额使用情况
func GetAccountQuotaUsage(params model.GetAccountQuotaUsageRequest) models.Result {
	//GET zstack/v1/accounts/quota/{uuid}/usages
	url := fmt.Sprintf("zstack/v1/accounts/quota/%s/usages", params.UUID)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetAccountQuotaUsageFailed, err.Error())
	}
	var result model.GetAccountQuotaUsageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询账户资源引用
func QueryAccountResourceRef(params model.QueryAccountResourceRefRequest) models.Result {
	//GET zstack/v1/accounts/resources/refs
	//GET zstack/v1/accounts/resources/refs/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/accounts/resources/refs")
	} else {
		url = fmt.Sprintf("zstack/v1/accounts/resources/refs/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAccountResourceRefFailed, err.Error())
	}
	var result model.QueryAccountResourceRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//共享资源给账户
func ShareResource(params model.ShareResourceRequest) models.Result {
	//PUT zstack/v1/accounts/resources/actions
	url := fmt.Sprintf("zstack/v1/accounts/resources/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ShareResourceFailed, err.Error())
	}
	var result model.ShareResourceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建用户组
func CreateUserGroup(params model.CreateUserGroupRequest) models.Result {
	//POST zstack/v1/accounts/groups
	url := fmt.Sprintf("zstack/v1/accounts/groups")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateUserGroupFailed, err.Error())
	}
	var result model.CreateUserGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除用户组
func DeleteUserGroup(params model.DeleteUserGroupRequest) models.Result {
	//DELETE zstack/v1/accounts/groups/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/accounts/groups/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteUserGroupFailed, err.Error())
	}
	var result model.DeleteUserGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询用户组
func QueryUserGroup(params model.QueryUserGroupRequest) models.Result {
	//GET zstack/v1/accounts/groups
	//GET zstack/v1/accounts/groups/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/accounts/groups")
	} else {
		url = fmt.Sprintf("zstack/v1/accounts/groups/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryUserGroupFailed, err.Error())
	}
	var result model.QueryUserGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新用户组
func UpdateUserGroup(params model.UpdateUserGroupRequest) models.Result {
	//PUT zstack/v1/accounts/groups/actions
	url := fmt.Sprintf("zstack/v1/accounts/groups/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateUserGroupFailed, err.Error())
	}
	var result model.UpdateUserGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加到用户组
func AddUserToGroup(params model.AddUserToGroupRequest) models.Result {
	//POST zstack/v1/accounts/groups/{groupUuid}/users
	url := fmt.Sprintf("zstack/v1/accounts/groups/%s/users", params.GroupUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddUserToGroupFailed, err.Error())
	}
	var result model.AddUserToGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//绑定策略到用户组
func AttachPolicyToUserGroup(params model.AttachPolicyToUserGroupRequest) models.Result {
	//POST zstack/v1/accounts/groups/{groupUuid}/policies
	url := fmt.Sprintf("zstack/v1/accounts/groups/%s/policies", params.GroupUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachPolicyToUserGroupFailed, err.Error())
	}
	var result model.AttachPolicyToUserGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//将策略从用户组解绑
func DetachPolicyFromUserGroup(params model.DetachPolicyFromUserGroupRequest) models.Result {
	//DELETE zstack/v1/accounts/groups/{groupUuid}/policies/{policyUuid}
	url := fmt.Sprintf("zstack/v1/accounts/groups/%s/policies/%s", params.GroupUuid, params.PolicyUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachPolicyFromUserGroupFailed, err.Error())
	}
	var result model.DetachPolicyFromUserGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从用户组中移除用户
func RemoveUserFromGroup(params model.RemoveUserFromGroupRequest) models.Result {
	//DELETE zstack/v1/accounts/groups/{groupUuid}/users/{userUuid}
	url := fmt.Sprintf("zstack/v1/accounts/groups/%s/users/%s", params.GroupUuid, params.UserUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveUserFromGroupFailed, err.Error())
	}
	var result model.RemoveUserFromGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建用户
func CreateUser(params model.CreateUserRequest) models.Result {
	//POST zstack/v1/accounts/users
	url := fmt.Sprintf("zstack/v1/accounts/users")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateUserFailed, err.Error())
	}
	var result model.CreateUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除用户
func DeleteUser(params model.DeleteUserRequest) models.Result {
	//DELETE zstack/v1/accounts/users/{uuid}
	url := fmt.Sprintf("zstack/v1/accounts/users/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteUserFailed, err.Error())
	}
	var result model.DeleteAccountResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询用户
func QueryUser(params model.QueryUserRequest) models.Result {
	//GET zstack/v1/accounts/users
	//GET zstack/v1/accounts/users/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/accounts/users")
	} else {
		url = fmt.Sprintf("zstack/v1/accounts/users/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryUserFailed, err.Error())
	}
	var result model.QueryUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新用户
func UpdateUser(params model.UpdateUserRequest) models.Result {
	//PUT zstack/v1/accounts/users/actions
	url := fmt.Sprintf("zstack/v1/accounts/users/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateUserFailed, err.Error())
	}
	var result model.UpdateUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//使用用户身份登录
func LogInByUser(params model.LogInByUserRequest) models.Result {
	//PUT zstack/v1/accounts/users/login
	url := fmt.Sprintf("zstack/v1/accounts/users/login")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.LogInByUserFailed, err.Error())
	}
	var result model.LogInByUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//绑定一条策略到用户
func AttachPolicyToUser(params model.AttachPolicyToUserRequest) models.Result {
	//POST zstack/v1/accounts/users/{userUuid}/policies
	url := fmt.Sprintf("zstack/v1/accounts/users/%s/policies", params.UserUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachPolicyToUserFailed, err.Error())
	}
	var result model.AttachPolicyToUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//将一条策略从用户解绑
func DetachPolicyFromUser(params model.DetachPolicyFromUserRequest) models.Result {
	//DELETE zstack/v1/accounts/groups/{groupUuid}/policies/{policyUuid}
	url := fmt.Sprintf("zstack/v1/accounts/groups/%s/policies/%s", params.GroupUuid, params.PolicyUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachPolicyFromUserFailed, err.Error())
	}
	var result model.DetachPolicyFromUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//绑定多条策略到用户
func AttachPoliciesToUser(params model.AttachPoliciesToUserRequest) models.Result {
	//POST zstack/v1/accounts/users/{userUuid}/policy-collection
	url := fmt.Sprintf("zstack/v1/accounts/users/%s/policy-collection", params.UserUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachPoliciesToUserFailed, err.Error())
	}
	var result model.AttachPoliciesToUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//将多条策略从用户解绑
func DetachPoliciesFromUser(params model.DetachPoliciesFromUserRequest) models.Result {
	//DELETE zstack/v1/accounts/users/{userUuid}/policies?policyUuids={policyUuids}
	url := fmt.Sprintf("zstack/v1/accounts/users/%s/policies?policyUuids=%s", params.UserUuid, params.PolicyUuids)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachPoliciesFromUserFailed, err.Error())
	}
	var result model.DetachPoliciesFromUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建策略
func CreatePolicy(params model.CreatePolicyRequest) models.Result {
	//POST zstack/v1/accounts/policies
	url := fmt.Sprintf("zstack/v1/accounts/policies")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreatePolicyFailed, err.Error())
	}
	var result model.CreatePolicyResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除策略
func DeletePolicy(params model.DeletePolicyRequest) models.Result {
	//DELETE zstack/v1/accounts/policies/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/accounts/policies/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeletePolicyFailed, err.Error())
	}
	var result model.DeletePolicyResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询策略
func QueryPolicy(params model.QueryPolicyRequest) models.Result {
	//GET zstack/v1/accounts/policies
	//GET zstack/v1/accounts/policies/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/accounts/policies")
	} else {
		url = fmt.Sprintf("zstack/v1/accounts/policies/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPolicyFailed, err.Error())
	}
	var result model.QueryPolicyResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询配额
func QueryQuota(params model.QueryQuotaRequest) models.Result {
	//GET zstack/v1/accounts/quotas
	//GET zstack/v1/accounts/quotas/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/accounts/quotas")
	} else {
		url = fmt.Sprintf("zstack/v1/accounts/quotas/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryQuotaFFailed, err.Error())
	}
	var result model.QueryQuotaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新配额
func UpdateQuota(params model.UpdateQuotaRequest) models.Result {
	//PUT zstack/v1/accounts/quotas/actions
	url := fmt.Sprintf("zstack/v1/accounts/quotas/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateQuotaFailed, err.Error())
	}
	var result model.UpdateQuotaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取资源名称
func GetResourceNames(params model.GetResourceNamesRequest) models.Result {
	//GET zstack/v1/resources/names
	url := fmt.Sprintf("zstack/v1/resources/names")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetResourceNamesFailed, err.Error())
	}
	var result model.GetResourceNamesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询共享资源
func QuerySharedResource(params model.QuerySharedResourceRequest) models.Result {
	//GET zstack/v1/accounts/resources
	url := fmt.Sprintf("zstack/v1/accounts/resources")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QuerySharedResourceFailed, err.Error())
	}
	var result model.QuerySharedResourceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//解除资源共享
func RevokeResourceSharingRequest(params model.RevokeResourceSharingRequest) models.Result {
	//PUT zstack/v1/accounts/resources/actions
	url := fmt.Sprintf("zstack/v1/accounts/resources/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.RevokeResourceSharingRequestFailed, err.Error())
	}
	var result model.RevokeResourceSharingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//变更资源所有者
func ChangeResourceOwner(params model.ChangeResourceOwnerRequest) models.Result {
	//POST zstack/v1/account/{accountUuid}/resources
	url := fmt.Sprintf("zstack/v1/account/%s/resources", params.AccountUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.ChangeResourceOwnerFailed, err.Error())
	}
	var result model.ChangeResourceOwnerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//检查API权限
func CheckApiPermission(params model.CheckApiPermissionRequest) models.Result {
	//PUT zstack/v1/accounts/permissions/actions
	url := fmt.Sprintf("zstack/v1/accounts/permissions/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.CheckApiPermissionFailed, err.Error())
	}
	var result model.CheckApiPermissionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//验证会话的有效性
func ValidateSession(params model.ValidateSessionRequest) models.Result {
	//GET zstack/v1/accounts/sessions/{sessionUuid}/valid
	url := fmt.Sprintf("zstack/v1/accounts/sessions/%s/valid", params.SessionUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.ValidateSessionFailed, err.Error())
	}
	var result model.ValidateSessionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新会话
func RenewSession(params model.RenewSessionRequest) models.Result {
	//PUT zstack/v1/accounts/sessions/{sessionUuid}/renew
	url := fmt.Sprintf("zstack/v1/accounts/sessions/%s/renew", params.SessionUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.RenewSessionFailed, err.Error())
	}
	var result model.RenewSessionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//退出当前登录状态
func LogOut(params model.LogOutRequest) models.Result {
	//DELETE zstack/v1/accounts/sessions/{sessionUuid}
	url := fmt.Sprintf("stack/v1/accounts/sessions/%s", params.SessionUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.LogOutFailed, err.Error())
	}
	var result model.LogOutResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
