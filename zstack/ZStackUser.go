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

//创建账户
func CreateAccount(params model.CreateAccountRequest) mgresult.Result {
	//POST zstack/v1/accounts
	url := fmt.Sprintf("zstack/v1/accounts")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateAccountFailed, err)
	}
	var result model.CreateAccountResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除账户
func DeleteAccount(params model.DeleteAccountRequest) mgresult.Result {
	//DELETE zstack/v1/accounts/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/accounts/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteAccountFailed, err)
	}
	var result model.DeleteAccountResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询账户
func QueryAccount(params model.QueryAccountRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryAccountFailed, err)
	}
	var result model.QueryAccountResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新账户
func UpdateAccount(params model.UpdateAccountRequest) mgresult.Result {
	//PUT zstack/v1/accounts/{uuid}
	url := fmt.Sprintf("zstack/v1/accounts/%s", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateAccountFailed, err)
	}
	var result model.UpdateAccountResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//使用账户身份登录
func LogInByAccount(params model.LogInByAccountRequest) mgresult.Result {
	//PUT zstack/v1/accounts/login
	url := fmt.Sprintf("zstack/v1/accounts/login")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.LogInByAccountFailed, err)
	}
	var result model.LogInByAccountResponse

	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取登录验证码
func GetLoginCaptcha(params model.GetLoginCaptchaRequest) mgresult.Result {
	//GET zstack/v1/login/control/captcha
	url := fmt.Sprintf("zstack/v1/login/control/captcha")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetLoginCaptchaFailed, err)
	}
	var result model.GetLoginCaptchaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//刷新验证码
func RefreshCaptcha(params model.RefreshCaptchaRequest) mgresult.Result {
	//GET zstack/v1/captcha/refresh
	url := fmt.Sprintf("zstack/v1/captcha/refresh")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RefreshCaptchaFailed, err)
	}
	var result model.RefreshCaptchaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取双因子认证密匙
func GetTwoFactorAuthenticationSecret(params model.GetTwoFactorAuthenticationSecretRequest) mgresult.Result {
	//GET zstack/v1/twofactorauthentication/secret
	url := fmt.Sprintf("zstack/v1/twofactorauthentication/secret")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetTwoFactorAuthenticationSecretFailed, err)
	}
	var result model.GetTwoFactorAuthenticationSecretResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询双因子认证密匙
func QueryTwoFactorAuthentication(params model.QueryTwoFactorAuthenticationRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryTwoFactorAuthenticationFailed, err)
	}
	var result model.QueryTwoFactorAuthenticationResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取双因子认证状态
func GetTwoFactorAuthenticationState(params model.GetTwoFactorAuthenticationStateRequest) mgresult.Result {
	//GET zstack/v1/twofactorauthentication/state
	url := fmt.Sprintf("zstack/v1/twofactorauthentication/state")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetTwoFactorAuthenticationStateFailed, err)
	}
	var result model.GetTwoFactorAuthenticationStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取账户配额使用情况
func GetAccountQuotaUsage(params model.GetAccountQuotaUsageRequest) mgresult.Result {
	//GET zstack/v1/accounts/quota/{uuid}/usages
	url := fmt.Sprintf("zstack/v1/accounts/quota/%s/usages", params.UUID)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetAccountQuotaUsageFailed, err)
	}
	var result model.GetAccountQuotaUsageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询账户资源引用
func QueryAccountResourceRef(params model.QueryAccountResourceRefRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryAccountResourceRefFailed, err)
	}
	var result model.QueryAccountResourceRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//共享资源给账户
func ShareResource(params model.ShareResourceRequest) mgresult.Result {
	//PUT zstack/v1/accounts/resources/actions
	url := fmt.Sprintf("zstack/v1/accounts/resources/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ShareResourceFailed, err)
	}
	var result model.ShareResourceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建用户组
func CreateUserGroup(params model.CreateUserGroupRequest) mgresult.Result {
	//POST zstack/v1/accounts/groups
	url := fmt.Sprintf("zstack/v1/accounts/groups")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateUserGroupFailed, err)
	}
	var result model.CreateUserGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除用户组
func DeleteUserGroup(params model.DeleteUserGroupRequest) mgresult.Result {
	//DELETE zstack/v1/accounts/groups/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/accounts/groups/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteUserGroupFailed, err)
	}
	var result model.DeleteUserGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询用户组
func QueryUserGroup(params model.QueryUserGroupRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryUserGroupFailed, err)
	}
	var result model.QueryUserGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新用户组
func UpdateUserGroup(params model.UpdateUserGroupRequest) mgresult.Result {
	//PUT zstack/v1/accounts/groups/actions
	url := fmt.Sprintf("zstack/v1/accounts/groups/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateUserGroupFailed, err)
	}
	var result model.UpdateUserGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加到用户组
func AddUserToGroup(params model.AddUserToGroupRequest) mgresult.Result {
	//POST zstack/v1/accounts/groups/{groupUuid}/users
	url := fmt.Sprintf("zstack/v1/accounts/groups/%s/users", params.GroupUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddUserToGroupFailed, err)
	}
	var result model.AddUserToGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//绑定策略到用户组
func AttachPolicyToUserGroup(params model.AttachPolicyToUserGroupRequest) mgresult.Result {
	//POST zstack/v1/accounts/groups/{groupUuid}/policies
	url := fmt.Sprintf("zstack/v1/accounts/groups/%s/policies", params.GroupUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachPolicyToUserGroupFailed, err)
	}
	var result model.AttachPolicyToUserGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将策略从用户组解绑
func DetachPolicyFromUserGroup(params model.DetachPolicyFromUserGroupRequest) mgresult.Result {
	//DELETE zstack/v1/accounts/groups/{groupUuid}/policies/{policyUuid}
	url := fmt.Sprintf("zstack/v1/accounts/groups/%s/policies/%s", params.GroupUuid, params.PolicyUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachPolicyFromUserGroupFailed, err)
	}
	var result model.DetachPolicyFromUserGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从用户组中移除用户
func RemoveUserFromGroup(params model.RemoveUserFromGroupRequest) mgresult.Result {
	//DELETE zstack/v1/accounts/groups/{groupUuid}/users/{userUuid}
	url := fmt.Sprintf("zstack/v1/accounts/groups/%s/users/%s", params.GroupUuid, params.UserUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveUserFromGroupFailed, err)
	}
	var result model.RemoveUserFromGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建用户
func CreateUser(params model.CreateUserRequest) mgresult.Result {
	//POST zstack/v1/accounts/users
	url := fmt.Sprintf("zstack/v1/accounts/users")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateUserFailed, err)
	}
	var result model.CreateUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除用户
func DeleteUser(params model.DeleteUserRequest) mgresult.Result {
	//DELETE zstack/v1/accounts/users/{uuid}
	url := fmt.Sprintf("zstack/v1/accounts/users/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteUserFailed, err)
	}
	var result model.DeleteAccountResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询用户
func QueryUser(params model.QueryUserRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryUserFailed, err)
	}
	var result model.QueryUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新用户
func UpdateUser(params model.UpdateUserRequest) mgresult.Result {
	//PUT zstack/v1/accounts/users/actions
	url := fmt.Sprintf("zstack/v1/accounts/users/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateUserFailed, err)
	}
	var result model.UpdateUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//使用用户身份登录
func LogInByUser(params model.LogInByUserRequest) mgresult.Result {
	//PUT zstack/v1/accounts/users/login
	url := fmt.Sprintf("zstack/v1/accounts/users/login")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.LogInByUserFailed, err)
	}
	var result model.LogInByUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//绑定一条策略到用户
func AttachPolicyToUser(params model.AttachPolicyToUserRequest) mgresult.Result {
	//POST zstack/v1/accounts/users/{userUuid}/policies
	url := fmt.Sprintf("zstack/v1/accounts/users/%s/policies", params.UserUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachPolicyToUserFailed, err)
	}
	var result model.AttachPolicyToUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将一条策略从用户解绑
func DetachPolicyFromUser(params model.DetachPolicyFromUserRequest) mgresult.Result {
	//DELETE zstack/v1/accounts/groups/{groupUuid}/policies/{policyUuid}
	url := fmt.Sprintf("zstack/v1/accounts/groups/%s/policies/%s", params.GroupUuid, params.PolicyUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachPolicyFromUserFailed, err)
	}
	var result model.DetachPolicyFromUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//绑定多条策略到用户
func AttachPoliciesToUser(params model.AttachPoliciesToUserRequest) mgresult.Result {
	//POST zstack/v1/accounts/users/{userUuid}/policy-collection
	url := fmt.Sprintf("zstack/v1/accounts/users/%s/policy-collection", params.UserUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachPoliciesToUserFailed, err)
	}
	var result model.AttachPoliciesToUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将多条策略从用户解绑
func DetachPoliciesFromUser(params model.DetachPoliciesFromUserRequest) mgresult.Result {
	//DELETE zstack/v1/accounts/users/{userUuid}/policies?policyUuids={policyUuids}
	url := fmt.Sprintf("zstack/v1/accounts/users/%s/policies?policyUuids=%s", params.UserUuid, params.PolicyUuids)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachPoliciesFromUserFailed, err)
	}
	var result model.DetachPoliciesFromUserResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建策略
func CreatePolicy(params model.CreatePolicyRequest) mgresult.Result {
	//POST zstack/v1/accounts/policies
	url := fmt.Sprintf("zstack/v1/accounts/policies")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreatePolicyFailed, err)
	}
	var result model.CreatePolicyResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除策略
func DeletePolicy(params model.DeletePolicyRequest) mgresult.Result {
	//DELETE zstack/v1/accounts/policies/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/accounts/policies/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeletePolicyFailed, err)
	}
	var result model.DeletePolicyResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询策略
func QueryPolicy(params model.QueryPolicyRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryPolicyFailed, err)
	}
	var result model.QueryPolicyResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询配额
func QueryQuota(params model.QueryQuotaRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryQuotaFFailed, err)
	}
	var result model.QueryQuotaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新配额
func UpdateQuota(params model.UpdateQuotaRequest) mgresult.Result {
	//PUT zstack/v1/accounts/quotas/actions
	url := fmt.Sprintf("zstack/v1/accounts/quotas/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateQuotaFailed, err)
	}
	var result model.UpdateQuotaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取资源名称
func GetResourceNames(params model.GetResourceNamesRequest) mgresult.Result {
	//GET zstack/v1/resources/names
	url := fmt.Sprintf("zstack/v1/resources/names")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetResourceNamesFailed, err)
	}
	var result model.GetResourceNamesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询共享资源
func QuerySharedResource(params model.QuerySharedResourceRequest) mgresult.Result {
	//GET zstack/v1/accounts/resources
	url := fmt.Sprintf("zstack/v1/accounts/resources")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QuerySharedResourceFailed, err)
	}
	var result model.QuerySharedResourceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//解除资源共享
func RevokeResourceSharingRequest(params model.RevokeResourceSharingRequest) mgresult.Result {
	//PUT zstack/v1/accounts/resources/actions
	url := fmt.Sprintf("zstack/v1/accounts/resources/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RevokeResourceSharingRequestFailed, err)
	}
	var result model.RevokeResourceSharingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//变更资源所有者
func ChangeResourceOwner(params model.ChangeResourceOwnerRequest) mgresult.Result {
	//POST zstack/v1/account/{accountUuid}/resources
	url := fmt.Sprintf("zstack/v1/account/%s/resources", params.AccountUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeResourceOwnerFailed, err)
	}
	var result model.ChangeResourceOwnerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//检查API权限
func CheckApiPermission(params model.CheckApiPermissionRequest) mgresult.Result {
	//PUT zstack/v1/accounts/permissions/actions
	url := fmt.Sprintf("zstack/v1/accounts/permissions/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CheckApiPermissionFailed, err)
	}
	var result model.CheckApiPermissionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//验证会话的有效性
func ValidateSession(params model.ValidateSessionRequest) mgresult.Result {
	//GET zstack/v1/accounts/sessions/{sessionUuid}/valid
	url := fmt.Sprintf("zstack/v1/accounts/sessions/%s/valid", params.SessionUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ValidateSessionFailed, err)
	}
	var result model.ValidateSessionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新会话
func RenewSession(params model.RenewSessionRequest) mgresult.Result {
	//PUT zstack/v1/accounts/sessions/{sessionUuid}/renew
	url := fmt.Sprintf("zstack/v1/accounts/sessions/%s/renew", params.SessionUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RenewSessionFailed, err)
	}
	var result model.RenewSessionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//退出当前登录状态
func LogOut(params model.LogOutRequest) mgresult.Result {
	//DELETE zstack/v1/accounts/sessions/{sessionUuid}
	url := fmt.Sprintf("stack/v1/accounts/sessions/%s", params.SessionUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.LogOutFailed, err)
	}
	var result model.LogOutResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
