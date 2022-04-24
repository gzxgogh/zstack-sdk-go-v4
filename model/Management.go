package model

//请求控制台访问地址
type RequestConsoleAccessRequest struct {
	ReqConfig
	Params     RequestConsoleAccessParams `json:"params" bson:"params"`
	SystemTags []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RequestConsoleAccessParams struct {
	VMInstanceUUID string `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //云主机UUID
}

type RequestConsoleAccessResponse struct {
	Inventory ConsoleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询控制台代理
type QueryConsoleProxyAgentRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryConsoleProxyAgentResponse struct {
	Inventories []ConsoleProxyAgentInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//重连控制台代理
type ReconnectConsoleProxyAgentRequest struct {
	ReqConfig
	ReconnectConsoleProxyAgent ReconnectConsoleProxyAgentParams `json:"reconnectConsoleProxyAgent" bson:"reconnectConsoleProxyAgent"`
	SystemTags                 []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ReconnectConsoleProxyAgentParams struct {
	AgentUuids []string `json:"agentUuids" bson:"agentUuids"`
}

type ReconnectConsoleProxyAgentResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新控制台代理
type UpdateConsoleProxyAgentRequest struct {
	ReqConfig
	UUID                    string                        `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateConsoleProxyAgent UpdateConsoleProxyAgentParams `json:"updateConsoleProxyAgent" bson:"updateConsoleProxyAgent"`
	SystemTags              []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateConsoleProxyAgentParams struct {
	ConsoleProxyOverriddenIp string `json:"consoleProxyOverriddenIp" bson:"consoleProxyOverriddenIp"` //新的控制台代理IP
}

type UpdateConsoleProxyAgentResponse struct {
	Inventory ConsoleProxyAgentInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type ConsoleInventory struct {
	Scheme   string `json:"scheme" bson:"scheme"`     //访问协议类型
	Hostname string `json:"hostname" bson:"hostname"` //宿主机名称
	Port     int    `json:"port" bson:"port"`         //端口
	Token    string `json:"token" bson:"token"`       //口令
}

type ConsoleProxyAgentInventory struct {
	UUID         string `json:"uuid" bson:"uuid"`                 //资源的UUID，唯一标示该资源
	Description  string `json:"description" bson:"description"`   //详细描述
	ManagementIp string `json:"managementIp" bson:"managementIp"` //管理节点IP
	Type         string `json:"type" bson:"type"`                 //类型
	Status       string `json:"status" bson:"status"`             //状态（连接，断开）
	State        string `json:"state" bson:"state"`               //启用状态:Enabled,Disabled
	CreateDate   string `json:"createDate" bson:"createDate"`     //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"`     //最后一次修改时间
}

//增加IP访问控制规则
type AddAccessControlRuleRequest struct {
	ReqConfig
	Params     AddAccessControlRuleParams `json:"params" bson:"params"`
	SystemTags []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddAccessControlRuleParams struct {
	Name            string   `json:"name" bson:"name"`                                   //资源名称
	Description     string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Rule            string   `json:"rule,omitempty" bson:"rule,omitempty"`
	ControlStrategy string   `json:"controlStrategy" bson:"controlStrategy "`
	ResourceUuid    string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids        []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type AddAccessControlRuleResponse struct {
	Inventory AccessControlRuleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除IP访问控制规则
type DeleteAccessControlRuleRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteAccessControlRuleResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新IP访问控制规则
type UpdateAccessControlRuleRequest struct {
	ReqConfig
	UUID                    string                        `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateAccessControlRule UpdateAccessControlRuleParams `json:"updateSecurityGroup" bson:"updateSecurityGroup"`
	SystemTags              []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateAccessControlRuleParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateAccessControlRuleResponse struct {
	Inventory AccessControlRuleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询IP访问控制规则
type QueryAccessControlRuleRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAccessControlRuleResponse struct {
	Inventories []AccessControlRuleInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type AccessControlRuleInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name        string `json:"name" bson:"name"`               //资源名称
	Description string `json:"description" bson:"description"` //详细描述
	Rule        string `json:"rule,omitempty" bson:"rule,omitempty"`
	Strategy    string `json:"strategy" bson:"strategy "`
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

//创建AccessKey
type CreateAccessKeyRequest struct {
	ReqConfig
	Params     CreateAccessKeyParams `json:"params" bson:"params"`
	SystemTags []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateAccessKeyParams struct {
	AccountUuid     string `json:"accountUuid" bson:"accountUuid"`
	UserUuid        string `json:"userUuid" bson:"userUuid"`                                   //用户UUID
	Description     string `json:"description,omitempty" bson:"description,omitempty"`         //详细描述
	AccessKeyID     string `json:"AccessKeyID,omitempty" bson:"AccessKeyID,omitempty"`         //生成的AccessKeyID
	AccessKeySecret string `json:"AccessKeySecret,omitempty" bson:"AccessKeySecret,omitempty"` //生成的AccessKeySecret
	ResourceUuid    string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`       //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateAccessKeyResponse struct {
	Inventory AccessKeyInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除AccessKey
type DeleteAccessKeyRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteAccessKeyResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询AccessKey
type QueryAccessKeyRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAccessKeyResponse struct {
	Inventories []AccessKeyInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//开启或关闭AccessKey
type ChangeAccessKeyStateRequest struct {
	ReqConfig
	UUID                 string                     `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeAccessKeyState ChangeAccessKeyStateParams `json:"updateSecurityGroup" bson:"updateSecurityGroup"`
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeAccessKeyStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //enable，disable
}

type ChangeAccessKeyStateResponse struct {
	Inventory AccessKeyInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type AccessKeyInventory struct {
	UUID            string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Description     string `json:"description" bson:"description"` //详细描述
	AccountUuid     string `json:"accountUuid" bson:"accountUuid"`
	UserUuid        string `json:"userUuid" bson:"userUuid"`               //用户UUID
	AccessKeyID     string `json:"AccessKeyID" bson:"AccessKeyID"`         //生成的AccessKeyID
	AccessKeySecret string `json:"AccessKeySecret" bson:"AccessKeySecret"` //生成的AccessKeySecret
	State           string `json:"state" bson:"state"`
	CreateDate      string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

//添加日志服务器配置
type AddLogConfigurationRequest struct {
	ReqConfig
	Params     AddLogConfigurationParams `json:"params" bson:"params"`
	SystemTags []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddLogConfigurationParams struct {
	Name          string   `json:"name" bson:"name"`                                   //资源名称
	Description   string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Type          string   `json:"type" bson:"type"`
	Configuration string   `json:"configuration" bson:"configuration"`
	ResourceUuid  string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids      []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type AddLogConfigurationResponse struct {
	Inventory LogConfigurationInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除日志服务器配置
type DeleteLogConfigurationRequest struct {
	ReqConfig
	ConfigId   int64    `json:"configId" bson:"configId"`                         //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteLogConfigurationResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取日志服务器配置
type GetLogConfigurationRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetLogConfigurationResponse struct {
	Inventories []LogConfigurationInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改日志服务器配置
type UpdateLogConfigurationRequest struct {
	ReqConfig
	UpdateLogConfiguration UpdateLogConfigurationParams `json:"updateLogConfiguration" bson:"updateLogConfiguration"`
	SystemTags             []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags               []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateLogConfigurationParams struct {
	ConfigId    int64  `json:"configId" bson:"configId"`                           //资源的UUID，唯一标示该资源
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateLogConfigurationResponse struct {
	Inventory LogConfigurationInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type LogConfigurationInventory struct {
	Id           int64  `json:"id" bson:"id"`
	LabelKey     string `json:"labelKey" bson:"labelKey"`
	LabelValue   string `json:"labelValue" bson:"labelValue"`
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	CreateDate   string `json:"createDate" bson:"createDate"`                         //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"`                         //最后一次修改时间
}
