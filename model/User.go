package model

//创建账户
type CreateAccountRequest struct {
	ReqConfig
	Params     CreateAccountParams `json:"params" bson:"params"`
	SystemTags []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateAccountParams struct {
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Password     string `json:"password" bson:"password"`                             //密码
	Type         string `json:"type,omitempty" bson:"type,omitempty"`                 //账户类型
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateAccountResponse struct {
	Inventory AccountInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除账户
type DeleteAccountRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteAccountResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询账户
type QueryAccountRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAccountResponse struct {
	Inventories []AccountInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新账户
type UpdateAccountRequest struct {
	ReqConfig
	UUID          string              `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateAccount UpdateAccountParams `json:"updateAccount" bson:"updateAccount"`
	SystemTags    []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateAccountParams struct {
	Password    string `json:"password,omitempty" bson:"password,omitempty"`       //密码
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateAccountResponse struct {
	Inventory AccountInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//使用账户身份登录
type LogInByAccountRequest struct {
	ReqConfig
	LogInByAccount LogInByAccountParams `json:"logInByAccount" bson:"logInByAccount"`             //资源的UUID，唯一标示该资源
	SystemTags     []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type LogInByAccountParams struct {
	AccountName string                 `json:"accountName" bson:"accountName"`                     //账户名称,账户名称和账户UUID不能同时为空
	Password    string                 `json:"password,omitempty" bson:"password,omitempty"`       //密码
	CaptchaUuid string                 `json:"captchaUuid,omitempty" bson:"captchaUuid,omitempty"` //验证码UUID
	VerifyCode  string                 `json:"verifyCode,omitempty" bson:"verifyCode,omitempty"`   //验证码
	ClientInfo  map[string]interface{} `json:"clientInfo,omitempty" bson:"clientInfo,omitempty"`   //客户端信息
}

type LogInByAccountResponse struct {
	Inventory SessionInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取登录验证码
type GetLoginCaptchaRequest struct {
	ReqConfig
	ResourceName string   `json:"resourceName" bson:"resourceName"`                 //资源名称
	LoginType    string   `json:"loginType" bson:"loginType"`                       //验证码UUID
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetLoginCaptchaResponse struct {
	CaptchaUuid string    `json:"captchaUuid" bson:"captchaUuid"` //验证码UUID
	Captcha     string    `json:"captcha" bson:"captcha"`         //验证码图片的base64形式
	Success     bool      `json:"success" bson:"success"`
	Error       ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//刷新验证码
type RefreshCaptchaRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RefreshCaptchaResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取双因子认证密匙
type GetTwoFactorAuthenticationSecretRequest struct {
	ReqConfig
	Name        string   `json:"name" bson:"name"`                                   //资源名称
	Password    string   `json:"password" bson:"password"`                           //密码
	Type        string   `json:"type" bson:"type"`                                   //账户类型
	CaptchaUuid string   `json:"captchaUuid,omitempty" bson:"captchaUuid,omitempty"` //验证码UUID
	VerifyCode  string   `json:"verifyCode,omitempty" bson:"verifyCode,omitempty"`   //验证码
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`   //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetTwoFactorAuthenticationSecretResponse struct {
	Inventory TwoFactorAuthenticationSecretInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询双因子认证密匙
type QueryTwoFactorAuthenticationRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryTwoFactorAuthenticationResponse struct {
	Inventories []TwoFactorAuthenticationSecretInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取双因子认证状态
type GetTwoFactorAuthenticationStateRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetTwoFactorAuthenticationStateResponse struct {
	State string    `json:"state" bson:"state"`
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取账户配额使用情况
type GetAccountQuotaUsageRequest struct {
	ReqConfig
	UUID       string                     `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Params     GetAccountQuotaUsageParams `json:"params" bson:"params"`
	SystemTags []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetAccountQuotaUsageParams struct {
}

type GetAccountQuotaUsageResponse struct {
	Usages []AccountQuotaUsage `json:"usages" bson:"usages"`
	Error  ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询账户资源引用
type QueryAccountResourceRefRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAccountResourceRefResponse struct {
	Usages []AccountResourceRef `json:"usages" bson:"usages"`
	Error  ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//共享资源给账户
type ShareResourceRequest struct {
	ReqConfig
	ShareResource ShareResourceParams `json:"shareResource" bson:"shareResource"`
	SystemTags    []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ShareResourceParams struct {
	ResourceUuids []string `json:"resourceUuids" bson:"resourceUuids"`
	AccountUuids  []string `json:"accountUuids,omitempty" bson:"accountUuids,omitempty"`
	ToPublic      bool     `json:"toPublic,omitempty" bson:"toPublic,omitempty"`
}

type ShareResourceResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建用户组
type CreateUserGroupRequest struct {
	ReqConfig
	Params     CreateAccountParams `json:"params" bson:"params"`
	SystemTags []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateUserGroupParams struct {
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateUserGroupResponse struct {
	Inventory UserGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除用户组
type DeleteUserGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteUserGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询用户组
type QueryUserGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryUserGroupResponse struct {
	Inventories []UserGroupInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新用户组
type UpdateUserGroupRequest struct {
	ReqConfig
	UUID            string                `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateUserGroup UpdateUserGroupParams `json:"updateUserGroup" bson:"updateUserGroup"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateUserGroupParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateUserGroupResponse struct {
	Inventory UserGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加到用户组
type AddUserToGroupRequest struct {
	ReqConfig
	GroupUuid  string               `json:"groupUuid" bson:"groupUuid"` //用户组UUID
	Params     AddUserToGroupParams `json:"params" bson:"params"`
	SystemTags []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddUserToGroupParams struct {
	UserUuid string `json:"userUuid" bson:"userUuid"` //用户UUID
}

type AddUserToGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//绑定策略到用户组
type AttachPolicyToUserGroupRequest struct {
	ReqConfig
	GroupUuid  string                        `json:"groupUuid" bson:"groupUuid"` //用户组UUID
	Params     AttachPolicyToUserGroupParams `json:"params" bson:"params"`
	SystemTags []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachPolicyToUserGroupParams struct {
	PolicyUuid string `json:"policyUuid" bson:"policyUuid"` //权限策略UUID
}

type AttachPolicyToUserGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//将策略从用户组解绑
type DetachPolicyFromUserGroupRequest struct {
	ReqConfig
	PolicyUuid string   `json:"policyUuid" bson:"policyUuid"`                     //权限策略UUID
	GroupUuid  string   `json:"groupUuid" bson:"groupUuid"`                       //用户组UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachPolicyFromUserGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从用户组中移除用户
type RemoveUserFromGroupRequest struct {
	ReqConfig
	GroupUuid  string   `json:"groupUuid" bson:"groupUuid"`                       //用户组UUID
	UserUuid   string   `json:"userUuid" bson:"userUuid"`                         //用户UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveUserFromGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建用户
type CreateUserRequest struct {
	ReqConfig
	Params     CreateUserParams `json:"params" bson:"params"`
	SystemTags []string         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateUserParams struct {
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Password     string `json:"password" bson:"password"`                             //密码
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateUserResponse struct {
	Inventory UserGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除用户
type DeleteUserRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteUserResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询用户
type QueryUserRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryUserResponse struct {
	Inventories []UserGroupInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新用户
type UpdateUserRequest struct {
	ReqConfig
	UpdateUser UpdateUserParams `json:"updateAccount" bson:"updateAccount"`
	SystemTags []string         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateUserParams struct {
	UUID        string `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Password    string `json:"password,omitempty" bson:"password,omitempty"`       //密码
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateUserResponse struct {
	Inventory UserGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//使用用户身份登录
type LogInByUserRequest struct {
	ReqConfig
	LogInByUser LogInByUserParams `json:"logInByUser" bson:"logInByUser"`
	SystemTags  []string          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type LogInByUserParams struct {
	AccountUuid string                 `json:"accountUuid " bson:"accountUuid "`                 //账户UUID
	AccountName string                 `json:"accountName" bson:"accountName"`                   //账户名称,账户名称和账户UUID不能同时为空
	UserName    string                 `json:"userName" bson:"userName"`                         //用户名称
	Password    string                 `json:"password" bson:"password"`                         //密码
	ClientInfo  map[string]interface{} `json:"clientInfo,omitempty" bson:"clientInfo,omitempty"` //客户端信息
}

type LogInByUserResponse struct {
	Inventory SessionInventory `json:"inventory" bson:"inventories"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//绑定一条策略到用户
type AttachPolicyToUserRequest struct {
	ReqConfig
	UserUuid   string   `json:"userUuid" bson:"userUuid"`                         //用户UUID
	PolicyUuid string   `json:"policyUuid" bson:"policyUuid"`                     //权限策略UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachPolicyToUserResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//将一条策略从用户解绑
type DetachPolicyFromUserRequest struct {
	ReqConfig
	GroupUuid  string   `json:"groupUuid" bson:"groupUuid"`                       //用户组UUID
	PolicyUuid string   `json:"policyUuid" bson:"policyUuid"`                     //权限策略UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachPolicyFromUserResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//绑定多条策略到用户
type AttachPoliciesToUserRequest struct {
	ReqConfig
	UserUuid    string   `json:"userUuid" bson:"userUuid"`                         //用户UUID
	PolicyUuids []string `json:"policyUuids" bson:"policyUuids"`                   //权限策略UUID
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachPoliciesToUserResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//将多条策略从用户解绑
type DetachPoliciesFromUserRequest struct {
	ReqConfig
	UserUuid    string   `json:"userUuid" bson:"userUuid"`                         //用户UUID
	PolicyUuids []string `json:"policyUuids" bson:"policyUuids"`                   //权限策略UUID
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachPoliciesFromUserResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建策略
type CreatePolicyRequest struct {
	ReqConfig
	Params     CreatePolicyParams `json:"params" bson:"params"`
	SystemTags []string           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreatePolicyParams struct {
	Name         string   `json:"name" bson:"name"`                                   //资源名称
	Description  string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Statements   []string `json:"statements" bson:"statements"`
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreatePolicyResponse struct {
	Inventory PolicyInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除策略
type DeletePolicyRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeletePolicyResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询策略
type QueryPolicyRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPolicyResponse struct {
	Inventories []PolicyInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询配额
type QueryQuotaRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryQuotaResponse struct {
	Inventories []QuotaInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新配额
type UpdateQuotaRequest struct {
	ReqConfig
	UpdateQuota UpdateQuotaParams `json:"updateQuota" bson:"updateQuota"`
	SystemTags  []string          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateQuotaParams struct {
	Name         string `json:"name" bson:"name"`                 //资源名称
	IdentityUuid string `json:"identityUuid" bson:"identityUuid"` //身份UUID（账户UUID，用户UUID）
	Value        int64  `json:"value" bson:"value"`               //默认配额值
}

type UpdateQuotaResponse struct {
	Inventory PolicyInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取资源名称
type GetResourceNamesRequest struct {
	ReqConfig
	UUIDS      []string `json:"uuids" bson:"uuids"`                               //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetResourceNamesResponse struct {
	Inventories []ResourceInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询共享资源
type QuerySharedResourceRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySharedResourceResponse struct {
	Inventories []SharedResourceInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//解除资源共享
type RevokeResourceSharingRequest struct {
	ReqConfig
	RevokeResourceSharing RevokeResourceSharingParams `json:"revokeResourceSharing" bson:"revokeResourceSharing"`
	SystemTags            []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags              []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RevokeResourceSharingParams struct {
	ResourceUuids []string `json:"resourceUuids" bson:"resourceUuids"`
	ToPublic      bool     `json:"toPublic,omitempty" bson:"toPublic,omitempty"`
	AccountUuids  []string `json:"accountUuids,omitempty" bson:"accountUuids,omitempty"`
	All           bool     `json:"all,omitempty" bson:"all,omitempty"`
}

type RevokeResourceSharingResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//变更资源所有者
type ChangeResourceOwnerRequest struct {
	ReqConfig
	AccountUuid string                    `json:"accountUuid" bson:"accountUuid"` //账户UUID
	Params      ChangeResourceOwnerParams `json:"params" bson:"params"`
	SystemTags  []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeResourceOwnerParams struct {
	ResourceUuid string `json:"resourceUuid" bson:"resourceUuid"`
}

type ChangeResourceOwnerResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//检查API权限
type CheckApiPermissionRequest struct {
	ReqConfig
	CheckApiPermission CheckApiPermissionParams `json:"checkApiPermission" bson:"checkApiPermission"`
	SystemTags         []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CheckApiPermissionParams struct {
	UserUuid string   `json:"userUuid,omitempty" bson:"userUuid,omitempty"` //用户UUID
	ApiNames []string `json:"apiNames" bson:"apiNames"`
}

type CheckApiPermissionResponse struct {
	Inventory map[string]interface{} `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//验证会话的有效性
type ValidateSessionRequest struct {
	ReqConfig
	SessionUuid string   `json:"sessionUuid" bson:"sessionUuid"`                   //会话uuid
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ValidateSessionResponse struct {
	ValidSession bool      `json:"validSession" bson:"validSession"`
	Error        ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新会话
type RenewSessionRequest struct {
	ReqConfig
	SessionUuid  string             `json:"sessionUuid" bson:"sessionUuid"` //会话uuid
	RenewSession RenewSessionParams `json:"renewSession" bson:"renewSession"`
	SystemTags   []string           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RenewSessionParams struct {
	Duration int64 `json:"duration" bson:"duration"` //用户UUID
}

type RenewSessionResponse struct {
	Inventory SessionInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//退出当前登录状态
type LogOutRequest struct {
	ReqConfig
	SessionUuid string                 `json:"sessionUuid" bson:"sessionUuid"` //会话uuid
	ClientInfo  map[string]interface{} `json:"clientInfo" bson:"clientInfo"`
	SystemTags  []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type LogOutResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type AccountInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Type        string `json:"type,omitempty" bson:"type,omitempty"`               //账户类型
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	CreateDate  string `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
}

type SessionInventory struct {
	UUID                string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	AccountUuid         string `json:"accountUuid" bson:"accountUuid"` //账户UUID
	UserUuid            string `json:"userUuid" bson:"userUuid"`       //用户UUID
	ExpiredDate         string `json:"expiredDate" bson:"expiredDate"` //会话过期日期
	CreateDate          string `json:"createDate" bson:"createDate"`   //创建时间
	NoSessionEvaluation bool   `json:"noSessionEvaluation" bson:"noSessionEvaluation"`
}

type TwoFactorAuthenticationSecretInventory struct {
	UUID     string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Secret   string `json:"secret" bson:"secret"`
	UserUuid string `json:"userUuid" bson:"userUuid"`
	UserType string `json:"userType" bson:"userType"`
	Status   string `json:"status" bson:"status"`
}

type AccountQuotaUsage struct {
	Name  string `json:"name" bson:"name"` //资源名称
	Total int64  `json:"total" bson:"total"`
	Used  int64  `json:"used" bson:"used"`
}

type AccountResourceRef struct {
	AccountUuid  string `json:"accountUuid" bson:"accountUuid"`
	ResourceUuid string `json:"resourceUuid" bson:"resourceUuid"`
	ResourceType string `json:"resourceType" bson:"resourceType"`
	CreateDate   string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type UserGroupInventory struct {
	UUID        string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	AccountUuid string `json:"accountUuid" bson:"accountUuid"`
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	CreateDate  string `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
}

type PolicyInventory struct {
	UUID        string     `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name        string     `json:"name" bson:"name"` //资源名称
	AccountUuid string     `json:"accountUuid" bson:"accountUuid"`
	Statements  Statements `json:"statements" bson:"statements"`
}

type Statements struct {
	Name       string   `json:"name" bson:"name"` //资源名称
	Principals []string `json:"principals" bson:"principals"`
	Actions    []string `json:"actions" bson:"actions"`     //一个匹配API的字符串列表
	Resources  []string `json:"resources" bson:"resources"` //资源列表
	Effect     string   `json:"effect" bson:"effect"`
}

type QuotaInventory struct {
	Name         string `json:"name" bson:"name"`                 //资源名称
	IdentityUuid string `json:"identityUuid" bson:"identityUuid"` //身份UUID（账户UUID，用户UUID）
	IdentityType string `json:"identityType" bson:"identityType"` //身份类型（账户，用户
	Value        int64  `json:"value" bson:"value"`               //默认配额值
	CreateDate   string `json:"createDate" bson:"createDate"`     //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"`     //最后一次修改时间
}

type ResourceInventory struct {
	UUID         string `json:"uuid" bson:"uuid"`                 //资源的UUID，唯一标示该资源
	ResourceName string `json:"resourceName" bson:"resourceName"` //资源名称
	ResourceType string `json:"resourceType" bson:"resourceType"`
}

type SharedResourceInventory struct {
	OwnerAccountUuid    string `json:"ownerAccountUuid" bson:"ownerAccountUuid"`       //所有者账户UUID
	ReceiverAccountUuid string `json:"receiverAccountUuid" bson:"receiverAccountUuid"` //接受者账户UUID
	ToPublic            string `json:"toPublic" bson:"toPublic"`                       //是否全局共享
	ResourceType        string `json:"resourceType" bson:"resourceType"`               //资源类型
	ResourceUuid        string `json:"resourceUuid" bson:"resourceUuid"`               //资源UUID
	CreateDate          string `json:"createDate" bson:"createDate"`                   //创建时间
	LastOpDate          string `json:"lastOpDate" bson:"lastOpDate"`                   //最后一次修改时间
}
