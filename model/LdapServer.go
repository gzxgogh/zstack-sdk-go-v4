package model

//添加AD/LDAP服务器
type AddLdapServerRequest struct {
	ReqConfig
	Params     CreatePriceTableParams `json:"params" bson:"params"`
	SystemTags []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddLdapServerParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Url         string `json:"url" bson:"url"`
	Base        string `json:"base" bson:"base"`             //LDAP服务器访问地址
	UserName    string `json:"username" bson:"username"`     //用户名称
	Password    string `json:"password" bson:"password"`     //密码
	Encryption  string `json:"encryption" bson:"encryption"` //None,tls
}

type AddLdapServerResponse struct {
	Inventory LdapServerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除AD/LDAP服务器
type DeleteLdapServerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteLdapServerResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询AD/LDAP服务器
type QueryLdapServerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryLdapServerResponse struct {
	Inventories []LdapServerInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新AD/LDAP服务器
type UpdateLdapServerRequest struct {
	ReqConfig
	LdapServerUuid   string                 `json:"ldapServerUuid" bson:"ldapServerUuid"` //资源的UUID，唯一标示该资源
	UpdateLdapServer UpdateLdapServerParams `json:"updateLdapServer" bson:"updateLdapServer"`
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateLdapServerParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Url         string `json:"url,omitempty" bson:"url,omitempty"`
	Base        string `json:"base,omitempty" bson:"base,omitempty"`             //LDAP服务器访问地址
	UserName    string `json:"username,omitempty" bson:"username,omitempty"`     //用户名称
	Password    string `json:"password,omitempty" bson:"password,omitempty"`     //密码
	Encryption  string `json:"encryption,omitempty" bson:"encryption,omitempty"` //None,tls
}

type UpdateLdapServerResponse struct {
	Inventory LdapServerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建AD/LDAP绑定
type CreateLdapBindingRequest struct {
	ReqConfig
	Params     CreateLdapBindingParams `json:"params" bson:"params"`
	SystemTags []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateLdapBindingParams struct {
	LdapUid     string `json:"ldapUid" bson:"ldapUid"`         //LDAP UID
	AccountUuid string `json:"accountUuid" bson:"accountUuid"` //账户UUID
}

type CreateLdapBindingResponse struct {
	Inventory LdapAccountRefInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除AD/LDAP绑定
type DeleteLdapBindingRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteLdapBindingResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询AD/LDAP绑定
type QueryLdapBindingRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryLdapBindingResponse struct {
	Inventories []LdapAccountRefInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//清理无效的AD/LDAP绑定
type CleanInvalidLdapBindingRequest struct {
	ReqConfig
	CleanInvalidLdapBinding CleanInvalidLdapBindingParams `json:"cleanInvalidLdapBinding" bson:"cleanInvalidLdapBinding"`
	SystemTags              []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CleanInvalidLdapBindingParams struct {
}

type CleanInvalidLdapBindingResponse struct {
	Inventories []LdapBindingInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//使用AD/LDAP身份登录
type LogInByLdapRequest struct {
	ReqConfig
	LogInByLdap LogInByLdapParams `json:"logInByLdap" bson:"logInByLdap"`
	SystemTags  []string          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type LogInByLdapParams struct {
	Uid         string                 `json:"uid" bson:"uid"`
	Password    string                 `json:"password" bson:"password"`                           //密码
	CaptchaUuid string                 `json:"captchaUuid,omitempty" bson:"captchaUuid,omitempty"` //验证码UUID
	VerifyCode  string                 `json:"verifyCode,omitempty" bson:"verifyCode,omitempty"`   //验证码
	ClientInfo  map[string]interface{} `json:"clientInfo,omitempty" bson:"clientInfo,omitempty"`   //客户端信息
}

type LogInByLdapResponse struct {
	Success          bool             `json:"success" bson:"success"`
	Inventory        SessionInventory `json:"inventory" bson:"inventory"`
	AccountInventory AccountInventory `json:"AccountInventory" bson:"AccountInventory"`
	Error            ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取AD/LDAP条目
type GetLdapEntryRequest struct {
	ReqConfig
	LdapFilter string   `json:"ldapFilter" bson:"ldapFilter"` //资源的UUID，唯一标示该资源
	Limit      int      `json:"limit" bson:"limit"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetLdapEntryResponse struct {
	Inventories []map[string]interface{} `json:"inventories" bson:"inventories"`
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取可绑定的AD/LDAP条目
type GetCandidateLdapEntryForBindingRequest struct {
	ReqConfig
	LdapFilter string   `json:"ldapFilter" bson:"ldapFilter"` //资源的UUID，唯一标示该资源
	Limit      int      `json:"limit" bson:"limit"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetCandidateLdapEntryForBindingResponse struct {
	Inventories []map[string]interface{} `json:"inventories" bson:"inventories"`
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type LdapServerInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Url         string `json:"url" bson:"url"`
	Base        string `json:"base" bson:"base"`             //LDAP服务器访问地址
	UserName    string `json:"username" bson:"username"`     //用户名称
	Password    string `json:"password" bson:"password"`     //密码
	Encryption  string `json:"encryption" bson:"encryption"` //None,tls
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type LdapAccountRefInventory struct {
	UUID           string `json:"uuid" bson:"uuid"`                     //资源的UUID，唯一标示该资源
	LdapUid        string `json:"ldapUid" bson:"ldapUid"`               //LDAP UID
	LdapServerUuid string `json:"ldapServerUuid" bson:"ldapServerUuid"` //LDAP服务器UUID
	AccountUuid    string `json:"accountUuid" bson:"accountUuid"`       //账户UUID
	CreateDate     string `json:"createDate" bson:"createDate"`         //创建时间
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"`         //最后一次修改时间
}

type LdapBindingInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Type        string `json:"type" bson:"type"`
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}
