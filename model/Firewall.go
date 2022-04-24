package model

//创建防火墙
type CreateVpcFirewallRequest struct {
	ReqConfig
	Params     CreateVpcFirewallParams `json:"params" bson:"params"`                             //资源的UUID，唯一标示该资源
	SystemTags []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateVpcFirewallParams struct {
	VpcUuid      string   `json:"vpcUuid" bson:"vpcUuid"`
	Description  string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Name         string   `json:"name" bson:"name"`
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreateVpcFirewallResponse struct {
	Inventory VpcFirewallInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询防火墙
type QueryVpcFirewallRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVpcFirewallResponse struct {
	Inventories []VpcFirewallInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新防火墙
type UpdateVpcFirewallRequest struct {
	ReqConfig
	UUID              string                  `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateVpcFirewall CreateVpcFirewallParams `json:"updateVpcFirewall" bson:"updateVpcFirewall"`
	SystemTags        []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateVpcFirewallParams struct {
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
}

type UpdateVpcFirewallResponse struct {
	Inventory VpcFirewallInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//刷新防火墙配置
type RefreshFirewallRequest struct {
	ReqConfig
	UUID            string                  `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	RefreshFirewall CreateVpcFirewallParams `json:"refreshFirewall" bson:"refreshFirewall"`
	SystemTags      []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RefreshFirewallParams struct {
}

type RefreshFirewallResponse struct {
	Inventory VpcFirewallInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除防火墙
type DeleteFirewallRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteFirewallResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建防火墙规则集
type CreateFirewallRuleSetRequest struct {
	ReqConfig
	Params     CreateFirewallRuleSetParams `json:"params" bson:"params"`                             //资源的UUID，唯一标示该资源
	SystemTags []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateFirewallRuleSetParams struct {
	Name            string   `json:"name" bson:"name"` //资源名称
	VpcFirewallUuid string   `json:"vpcFirewallUuid" bson:"vpcFirewallUuid"`
	ActionType      string   `json:"actionType" bson:"actionType"`                         //drop,accept,reject
	Description     string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid    string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids        []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreateFirewallRuleSetResponse struct {
	Inventory VpcFirewallRuleSetInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询防火墙规则集
type QueryFirewallRuleSetRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryFirewallRuleSetResponse struct {
	Inventories []VpcFirewallRuleSetInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新防火墙规则集
type UpdateFirewallRuleSetRequest struct {
	ReqConfig
	UUID                  string                      `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateFirewallRuleSet UpdateFirewallRuleSetParams `json:"updateFirewallRuleSet" bson:"updateFirewallRuleSet"`
	SystemTags            []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags              []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateFirewallRuleSetParams struct {
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	ActionType  string `json:"actionType,omitempty" bson:"actionType,omitempty"` //drop,accept,reject
}

type UpdateFirewallRuleSetResponse struct {
	Inventory VpcFirewallRuleSetInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除防火墙规则集
type DeleteFirewallRuleSetRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteFirewallRuleSetResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//网络加载防火墙规则集
type AttachFirewallRuleSetToL3Request struct {
	ReqConfig
	L3Uuid      string                      `json:"l3Uuid" bson:"l3Uuid"`           //资源的UUID，唯一标示该资源
	RuleSetUuid string                      `json:"ruleSetUuid" bson:"ruleSetUuid"` //策略路由规则集uuid
	Params      UpdateFirewallRuleSetParams `json:"params" bson:"params"`
	SystemTags  []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachFirewallRuleSetToL3Params struct {
	VpcFirewallUuid string `json:"vpcFirewallUuid" bson:"vpcFirewallUuid"`
	Forward         string `json:"forward" bson:"forward"` //in,out
}

type AttachFirewallRuleSetToL3Response struct {
	Inventory VpcFirewallRuleSetL3RefInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//网络卸载防火墙规则集
type DetachFirewallRuleSetFromL3Request struct {
	ReqConfig
	L3Uuid     string                      `json:"l3Uuid" bson:"l3Uuid"` //资源的UUID，唯一标示该资源
	Params     UpdateFirewallRuleSetParams `json:"params" bson:"params"`
	SystemTags []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachFirewallRuleSetFromL3Params struct {
	VpcFirewallUuid string `json:"vpcFirewallUuid" bson:"vpcFirewallUuid"`
	Forward         string `json:"forward" bson:"forward"` //in,out
}

type DetachFirewallRuleSetFromL3Response struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询防火墙规则集与三层网络关联关系
type QueryFirewallRuleSetL3RefRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryFirewallRuleSetL3RefResponse struct {
	Inventories []VpcFirewallRuleSetL3RefInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建防火墙规则
type CreateFirewallRuleRequest struct {
	ReqConfig
	Params     CreateFirewallRuleSetParams `json:"params" bson:"params"`                             //资源的UUID，唯一标示该资源
	SystemTags []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateFirewallRuleParams struct {
	VpcFirewallUuid string   `json:"vpcFirewallUuid" bson:"vpcFirewallUuid"`
	RuleSetUuid     string   `json:"ruleSetUuid" bson:"ruleSetUuid"`
	Action          string   `json:"action" bson:"action"` //drop,accept
	Protocol        string   `json:"protocol,omitempty" bson:"protocol,omitempty"`
	DestPort        string   `json:"destPort,omitempty" bson:"destPort,omitempty"`
	SourcePort      string   `json:"sourcePort,omitempty" bson:"sourcePort,omitempty"` //源端口
	SourceIp        string   `json:"sourceIp,omitempty" bson:"sourceIp,omitempty"`     //源ip
	DestIp          string   `json:"destIp,omitempty" bson:"destIp,omitempty"`         //目标ip
	AllowStates     string   `json:"allowStates,omitempty" bson:"allowStates,omitempty"`
	TcpFlag         string   `json:"tcpFlag,omitempty" bson:"tcpFlag,omitempty"`
	IcmpTypeName    string   `json:"icmpTypeName,omitempty" bson:"icmpTypeName,omitempty"`
	RuleNumber      int      `json:"ruleNumber" bson:"ruleNumber"`
	EnableLog       bool     `json:"enableLog,omitempty" bson:"enableLog,omitempty"`
	State           string   `json:"state" bson:"state"`                                   //enable,disable
	Description     string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid    string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids        []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreateFirewallRuleResponse struct {
	Inventory VpcFirewallRuleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询防火墙规则
type QueryFirewallRuleRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryFirewallRuleResponse struct {
	Inventories []VpcFirewallRuleInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新防火墙规则
type UpdateFirewallRuleRequest struct {
	ReqConfig
	UUID               string                   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateFirewallRule UpdateFirewallRuleParams `json:"updateFirewallRule" bson:"updateFirewallRule"`
	SystemTags         []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateFirewallRuleParams struct {
	VpcFirewallUuid string `json:"vpcFirewallUuid" bson:"vpcFirewallUuid"`
	RuleSetUuid     string `json:"ruleSetUuid" bson:"ruleSetUuid"`
	Action          string `json:"action" bson:"action"` //drop,accept
	Protocol        string `json:"protocol,omitempty" bson:"protocol,omitempty"`
	DestPort        string `json:"destPort,omitempty" bson:"destPort,omitempty"`
	SourcePort      string `json:"sourcePort,omitempty" bson:"sourcePort,omitempty"` //源端口
	SourceIp        string `json:"sourceIp,omitempty" bson:"sourceIp,omitempty"`     //源ip
	DestIp          string `json:"destIp,omitempty" bson:"destIp,omitempty"`         //目标ip
	AllowStates     string `json:"allowStates,omitempty" bson:"allowStates,omitempty"`
	TcpFlag         string `json:"tcpFlag,omitempty" bson:"tcpFlag,omitempty"`
	IcmpTypeName    string `json:"icmpTypeName,omitempty" bson:"icmpTypeName,omitempty"`
	RuleNumber      int    `json:"rule_number" bson:"rule_number"`
	EnableLog       bool   `json:"enableLog,omitempty" bson:"enableLog,omitempty"`
	State           string `json:"state" bson:"state"`                                 //enable,disable
	Description     string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateFirewallRuleResponse struct {
	Inventory VpcFirewallRuleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除防火墙规则
type DeleteFirewallRuleRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteFirewallRuleResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更改防火墙规则状态
type ChangeFirewallRuleStateRequest struct {
	ReqConfig
	UUID                    string                        `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeFirewallRuleState ChangeFirewallRuleStateParams `json:"changeFirewallRuleState" bson:"changeFirewallRuleState"`
	SystemTags              []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeFirewallRuleStateParams struct {
	State string `json:"state" bson:"state"`
}

type ChangeFirewallRuleStateResponse struct {
	Inventory VpcFirewallRuleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建防火墙规则模板
type CreateFirewallRuleTemplateRequest struct {
	ReqConfig
	UUID                       string                           `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	CreateFirewallRuleTemplate CreateFirewallRuleTemplateParams `json:"createFirewallRuleTemplate" bson:"createFirewallRuleTemplate"`
	SystemTags                 []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateFirewallRuleTemplateParams struct {
	Action       string   `json:"action" bson:"action"` //drop,accept,reject
	Protocol     string   `json:"protocol,omitempty" bson:"protocol,omitempty"`
	Name         string   `json:"name" bson:"name"`
	DestPort     string   `json:"destPort,omitempty" bson:"destPort,omitempty"`
	SourcePort   string   `json:"sourcePort,omitempty" bson:"sourcePort,omitempty"` //源端口
	SourceIp     string   `json:"sourceIp,omitempty" bson:"sourceIp,omitempty"`     //源ip
	DestIp       string   `json:"destIp,omitempty" bson:"destIp,omitempty"`         //目标ip
	AllowStates  string   `json:"allowStates,omitempty" bson:"allowStates,omitempty"`
	TcpFlag      string   `json:"tcpFlag,omitempty" bson:"tcpFlag,omitempty"`
	IcmpTypeName string   `json:"icmpTypeName,omitempty" bson:"icmpTypeName,omitempty"`
	RuleNumber   int      `json:"ruleNumber" bson:"ruleNumber"`
	EnableLog    bool     `json:"enableLog,omitempty" bson:"enableLog,omitempty"`
	State        string   `json:"state" bson:"state"`                                   //enable,disable
	Description  string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreateFirewallRuleTemplateResponse struct {
	Inventory VpcFirewallRuleTemplateInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除防火墙规则模板
type DeleteFirewallRuleTemplateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteFirewallRuleTemplateResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新防火墙规则模板
type UpdateFirewallRuleTemplateRequest struct {
	ReqConfig
	UUID                       string                           `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateFirewallRuleTemplate UpdateFirewallRuleTemplateParams `json:"updateFirewallRule" bson:"updateFirewallRule"`
	SystemTags                 []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateFirewallRuleTemplateParams struct {
	Name         string   `json:"name" bson:"name"`
	Action       string   `json:"action" bson:"action"` //drop,accept,reject
	Protocol     string   `json:"protocol,omitempty" bson:"protocol,omitempty"`
	DestPort     string   `json:"destPort,omitempty" bson:"destPort,omitempty"`
	SourcePort   string   `json:"sourcePort,omitempty" bson:"sourcePort,omitempty"` //源端口
	SourceIp     string   `json:"sourceIp,omitempty" bson:"sourceIp,omitempty"`     //源ip
	DestIp       string   `json:"destIp,omitempty" bson:"destIp,omitempty"`         //目标ip
	AllowStates  string   `json:"allowStates,omitempty" bson:"allowStates,omitempty"`
	TcpFlag      string   `json:"tcpFlag,omitempty" bson:"tcpFlag,omitempty"`
	IcmpTypeName string   `json:"icmpTypeName,omitempty" bson:"icmpTypeName,omitempty"`
	RuleNumber   int      `json:"ruleNumber" bson:"ruleNumber"`
	EnableLog    bool     `json:"enableLog,omitempty" bson:"enableLog,omitempty"`
	State        string   `json:"state" bson:"state"`                                   //enable,disable
	Description  string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type UpdateFirewallRuleTemplateResponse struct {
	Inventory VpcFirewallRuleTemplateInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询防火墙规则模板
type QueryFirewallRuleTemplateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryFirewallRuleTemplateResponse struct {
	Inventories []VpcFirewallRuleInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询防火墙与单节点路由器关联关系
type QueryVpcFirewallVRouterRefRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVpcFirewallVRouterRefResponse struct {
	Inventories []VpcFirewallVRouterRefInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//下发规则集更改内容到路由器
type ApplyRuleSetChangesRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ApplyRuleSetChangesResponse struct {
	Inventories VpcFirewallRuleSetInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建防火墙IP或端口模板
type CreateFirewallIpSetTemplateRequest struct {
	ReqConfig
	UUID                        string                            `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	CreateFirewallIpSetTemplate CreateFirewallIpSetTemplateParams `json:"changeFirewallRuleState" bson:"changeFirewallRuleState"`
	SystemTags                  []string                          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                    []string                          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateFirewallIpSetTemplateParams struct {
	Name         string   `json:"name" bson:"name"`
	SourceValue  string   `json:"sourceValue,omitempty" bson:"sourceValue,omitempty"`
	DestValue    string   `json:"destValue,omitempty" bson:"destValue,omitempty"`
	Type         string   `json:"type" bson:"type"`                                     //ip,port
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreateFirewallIpSetTemplateResponse struct {
	Inventory VpcFirewallIpSetTemplateInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除防火墙IP或端口模板
type DeleteFirewallIpSetTemplateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteFirewallIpSetTemplateResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新防火墙IP或端口模板
type UpdateFirewallIpSetTemplateRequest struct {
	ReqConfig
	UUID                        string                            `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateFirewallIpSetTemplate UpdateFirewallIpSetTemplateParams `json:"updateFirewallRule" bson:"updateFirewallRule"`
	SystemTags                  []string                          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                    []string                          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateFirewallIpSetTemplateParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	SourceValue string `json:"sourceValue,omitempty" bson:"sourceValue,omitempty"`
	DestValue   string `json:"destValue,omitempty" bson:"destValue,omitempty"`
	Type        string `json:"type" bson:"type"` //ip,port
}

type UpdateFirewallIpSetTemplateResponse struct {
	Inventory VpcFirewallRuleTemplateInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询防火墙IP或端口模板
type QueryFirewallIpSetTemplateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryFirewallIpSetTemplateResponse struct {
	Inventories []VpcFirewallIpSetTemplateInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type VpcFirewallInventory struct {
	UUID        string                           `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name        string                           `json:"name" bson:"name"`
	Description string                           `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	CreateDate  string                           `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate  string                           `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
	RuleSets    []VpcFirewallRuleSetInventory    `json:"ruleSets" bson:"ruleSets"`
	Refs        VpcFirewallRuleSetL3RefInventory `json:"refs" bson:"refs"`
}

type VpcFirewallRuleSetInventory struct {
	UUID            string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name            string `json:"name" bson:"name"`
	VpcFirewallUuid string `json:"vpcFirewallUuid" bson:"vpcFirewallUuid"`
	Description     string `json:"description" bson:"description"` //详细描述
	IsDefault       bool   `json:"isDefault" bson:"isDefault"`     //默认
	CreateDate      string `json:"createDate" bson:"createDate"`   //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"`   //最后一次修改时间
	ActionType      string `json:"actionType" bson:"actionType"`
}

type VpcFirewallRuleSetL3RefInventory struct {
	Id                 string `json:"id" bson:"id"`
	RuleSetUuid        string `json:"ruleSetUuid" bson:"ruleSetUuid"`
	L3Uuid             string `json:"l3Uuid" bson:"l3Uuid"`
	CreateDate         string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate         string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	PacketsForwardType string `json:"packetsForwardType" bson:"packetsForwardType"`
}

type VpcFirewallVRouterRefInventory struct {
	Id              string `json:"id" bson:"id"`
	VpcFirewallUuid string `json:"vpcFirewallUuid" bson:"vpcFirewallUuid"`
	VRouterUuid     string `json:"VRouterUuid" bson:"VRouterUuid"`
	CreateDate      string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间

}

type VpcFirewallRuleInventory struct {
	UUID            string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	VpcFirewallUuid string `json:"vpcFirewallUuid" bson:"vpcFirewallUuid"`
	RuleSetUuid     string `json:"ruleSetUuid" bson:"ruleSetUuid"`
	RuleSetName     string `json:"ruleSetName" bson:"ruleSetName"`
	DestPort        string `json:"destPort" bson:"destPort"`
	SourcePort      string `json:"sourcePort" bson:"sourcePort"` //源端口
	SourceIp        string `json:"sourceIp" bson:"sourceIp"`     //源ip
	DestIp          string `json:"destIp" bson:"destIp"`         //目标ip
	RuleNumber      int    `json:"rule_number" bson:"rule_number"`
	AllowStates     string `json:"allowStates" bson:"allowStates"`
	TcpFlag         string `json:"tcpFlag" bson:"tcpFlag"`
	IcmpTypeName    string `json:"icmpTypeName" bson:"icmpTypeName"`
	IsDefault       bool   `json:"isDefault" bson:"isDefault"`
	Description     string `json:"description" bson:"description"` //详细描述
	Action          string `json:"action" bson:"action"`           //drop,accept
	Protocol        string `json:"protocol" bson:"protocol"`
	State           string `json:"state" bson:"state"`           //enable,disable
	CreateDate      string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type VpcFirewallRuleTemplateInventory struct {
	UUID         string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	RuleSetUuid  string `json:"ruleSetUuid" bson:"ruleSetUuid"`
	DestPort     string `json:"destPort" bson:"destPort"`
	SourcePort   string `json:"sourcePort" bson:"sourcePort"` //源端口
	SourceIp     string `json:"sourceIp" bson:"sourceIp"`     //源ip
	DestIp       string `json:"destIp" bson:"destIp"`         //目标ip
	RuleNumber   int    `json:"rule_number" bson:"rule_number"`
	AllowStates  string `json:"allowStates" bson:"allowStates"`
	TcpFlag      string `json:"tcpFlag" bson:"tcpFlag"`
	IcmpTypeName string `json:"icmpTypeName" bson:"icmpTypeName"`
	IsDefault    bool   `json:"isDefault" bson:"isDefault"`
	IsApplied    bool   `json:"isApplied" bson:"isApplied"`
	Expired      bool   `json:"expired" bson:"expired"`
	Description  string `json:"description" bson:"description"` //详细描述
	Action       string `json:"action" bson:"action"`           //drop,accept
	Protocol     string `json:"protocol" bson:"protocol"`
	State        string `json:"state" bson:"state"`           //enable,disable
	CreateDate   string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type VpcFirewallIpSetTemplateInventory struct {
	UUID        string                             `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name        string                             `json:"name" bson:"name"`
	Description string                             `json:"description" bson:"description"` //详细描述
	IsDefault   bool                               `json:"isDefault" bson:"isDefault"`
	IsApplied   bool                               `json:"isApplied" bson:"isApplied"`
	ActionType  string                             `json:"actionType" bson:"actionType"`
	Rules       []VpcFirewallRuleTemplateInventory `json:"rules" bson:"rules"`
}
