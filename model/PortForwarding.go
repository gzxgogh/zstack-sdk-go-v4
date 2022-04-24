package model

//创建端口转发规则
type CreatePortForwardingRuleRequest struct {
	ReqConfig
	Params     CreatePortForwardingRuleParams `json:"params" bson:"params"`
	SystemTags []string                       `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                       `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreatePortForwardingRuleParams struct {
	VipUuid          string `json:"vipUuid" bson:"vipUuid"`
	VipPortStart     int    `json:"vipPortStart" bson:"vipPortStart"`                             //VIP的起始端口号
	VipPortEnd       int    `json:"vipPortEnd,omitempty" bson:"vipPortEnd,omitempty"`             //VIP的结束端口号; 如果忽略不设置, 会默认设置为vipPortStart
	PrivatePortStart int    `json:"privatePortStart,omitempty" bson:"privatePortStart,omitempty"` //客户IP（虚拟机网卡的IP地址）的起始端口号; 如果忽略不设置, 会默认设置为vipPortStart
	PrivatePortEnd   int    `json:"privatePortEnd,omitempty" bson:"privatePortEnd,omitempty"`     //客户IP（虚拟机网卡的IP地址）的结束端口号; 如果忽略不设置, 会默认设置为vipPortEnd
	ProtocolType     string `json:"protocolType" bson:"protocolType"`                             //网络流量协议类型:tpc,udp
	VmNicUuid        string `json:"vmNicUuid,omitempty" bson:"vmNicUuid,omitempty"`               //云主机网卡UUID
	AllowedCidr      string `json:"allowedCidr,omitempty" bson:"allowedCidr,omitempty"`           //源CIDR; 端口转发规则只作用于源CIDR的流量; 如果忽略不设置, 会默认设置为to 0.0.0.0/0
	Name             string `json:"name" bson:"name"`                                             //资源名称
	Description      string `json:"description,omitempty" bson:"description,omitempty"`           //详细描述
	ResourceUuid     string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`         //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreatePortForwardingRuleResponse struct {
	Inventory PortForwardingRuleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除端口转发规则
type DeletePortForwardingRuleRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeletePortForwardingRuleResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询端口转发规则
type QueryPortForwardingRuleRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPortForwardingRuleResponse struct {
	Inventories []PortForwardingRuleInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新端口转发规则
type UpdatePortForwardingRuleRequest struct {
	ReqConfig
	UUID                     string                         `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdatePortForwardingRule UpdatePortForwardingRuleParams `json:"updatePortForwardingRule" bson:"updatePortForwardingRule"`
	SystemTags               []string                       `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                 []string                       `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdatePortForwardingRuleParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdatePortForwardingRuleResponse struct {
	Inventory PortForwardingRuleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//改变端口转发规则的状态
type ChangePortForwardingRuleStateRequest struct {
	ReqConfig
	UUID                          string                              `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangePortForwardingRuleState ChangePortForwardingRuleStateParams `json:"changePortForwardingRuleState" bson:"changePortForwardingRuleState"`
	SystemTags                    []string                            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                      []string                            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangePortForwardingRuleStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"`
}

type ChangePortForwardingRuleStateResponse struct {
	Inventory PortForwardingRuleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机网卡列
type GetPortForwardingAttachableVmNicsRequest struct {
	ReqConfig
	RuleUuid   string   `json:"ruleUuid" bson:"ruleUuid"`                         //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetPortForwardingAttachableVmNicsResponse struct {
	Inventories []VmNicInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//挂载规则到虚拟机网卡上
type AttachPortForwardingRuleRequest struct {
	ReqConfig
	RuleUuid   string   `json:"ruleUuid" bson:"ruleUuid"`
	VmNicUuid  string   `json:"vmNicUuid" bson:"vmNicUuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachPortForwardingRuleResponse struct {
	Inventory PortForwardingRuleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从虚拟机网卡卸载规则
type DetachPortForwardingRuleRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachPortForwardingRuleResponse struct {
	Inventory PortForwardingRuleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type PortForwardingRuleInventory struct {
	UUID             string `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name             string `json:"name" bson:"name"`                                   //资源名称
	Description      string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	VipIp            string `json:"VipIp" bson:"VipIp"`                                 //VIP的IP地址
	GuestIp          string `json:"guestIp" bson:"guestIp"`                             //虚拟机网卡的IP地址
	VipUuid          string `json:"vipUuid" bson:"vipUuid"`
	VipPortStart     int    `json:"vipPortStart" bson:"vipPortStart"`         //VIP的起始端口号
	VipPortEnd       int    `json:"vipPortEnd" bson:"vipPortEnd"`             //VIP的结束端口号; 如果忽略不设置, 会默认设置为vipPortStart
	PrivatePortStart int    `json:"privatePortStart" bson:"privatePortStart"` //客户IP（虚拟机网卡的IP地址）的起始端口号; 如果忽略不设置, 会默认设置为vipPortStart
	PrivatePortEnd   int    `json:"privatePortEnd" bson:"privatePortEnd"`     //客户IP（虚拟机网卡的IP地址）的结束端口号; 如果忽略不设置, 会默认设置为vipPortEnd
	VmNicUuid        string `json:"vmNicUuid" bson:"vmNicUuid"`               //云主机网卡UUID，该网卡所在网络会从云主机卸载掉
	ProtocolType     string `json:"protocolType" bson:"protocolType"`         //网络流量协议类型:tpc,udp
	State            string `json:"state,omitempty" bson:"state,omitempty"`
	AllowedCidr      string `json:"allowedCidr" bson:"allowedCidr"` //允许的CIDR,根据流量类型的不同, 允许的CIDR有不同的含义,如果流量类型是Ingress, 允许的CIDR是允许访问虚拟机网卡的源CIDR,如果流量类型是Egress, 允许的CIDR是允许从虚拟机网卡离开并到达的目的地CIDR
	CreateDate       string `json:"createDate" bson:"createDate"`   //创建时间
	LastOpDate       string `json:"lastOpDate" bson:"lastOpDate"`   //最后一次修改时间
}
