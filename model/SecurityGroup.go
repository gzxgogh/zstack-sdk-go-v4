package model

//创建安全组
type CreateSecurityGroupRequest struct {
	ReqConfig
	Params     CreateSecurityGroupParams `json:"params" bson:"params"`
	SystemTags []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSecurityGroupParams struct {
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	IpVersion    int    `json:"ipVersion,omitempty" bson:"ipVersion,omitempty"`       //亲和组策略
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateSecurityGroupResponse struct {
	Inventory SecurityGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除安全组
type DeleteSecurityGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteSecurityGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询安全组
type QuerySecurityGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySecurityGroupResponse struct {
	Inventories []SecurityGroupInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新安全组
type UpdateSecurityGroupRequest struct {
	ReqConfig
	UUID                string                    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateSecurityGroup UpdateSecurityGroupParams `json:"updateSecurityGroup" bson:"updateSecurityGroup"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateSecurityGroupParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateSecurityGroupResponse struct {
	Inventory SecurityGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//改变安全组状态
type ChangeSecurityGroupStateRequest struct {
	ReqConfig
	UUID                     string                         `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeSecurityGroupState ChangeSecurityGroupStateParams `json:"changeSecurityGroupState" bson:"changeSecurityGroupState"`
	SystemTags               []string                       `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                 []string                       `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeSecurityGroupStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //enable,disable
}

type ChangeSecurityGroupStateResponse struct {
	Inventory SecurityGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//挂载安全组到L3网络
type AttachSecurityGroupToL3NetworkRequest struct {
	ReqConfig
	SecurityGroupUuid string   `json:"securityGroupUuid" bson:"securityGroupUuid"` //安全组UUID
	L3NetworkUuid     string   `json:"l3NetworkUuid" bson:"l3NetworkUuid"`
	SystemTags        []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachSecurityGroupToL3NetworkResponse struct {
	Inventory SecurityGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从L3网络卸载安全组
type DetachSecurityGroupFromL3NetworkRequest struct {
	ReqConfig
	SecurityGroupUuid string   `json:"securityGroupUuid" bson:"securityGroupUuid"` //安全组UUID
	L3NetworkUuid     string   `json:"l3NetworkUuid" bson:"l3NetworkUuid"`
	SystemTags        []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachSecurityGroupFromL3NetworkResponse struct {
	Inventory SecurityGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取网卡列表清单
type GetCandidateVmNicForSecurityGroupRequest struct {
	ReqConfig
	SecurityGroupUuid string   `json:"securityGroupUuid" bson:"securityGroupUuid"`       //安全组UUID
	SystemTags        []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetCandidateVmNicForSecurityGroupResponse struct {
	Inventories []VmNicInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加虚拟机网卡到安全组
type AddVmNicToSecurityGroupRequest struct {
	ReqConfig
	SecurityGroupUuid string                        `json:"securityGroupUuid" bson:"securityGroupUuid"` //安全组UUID
	Params            AddVmNicToSecurityGroupParams `json:"params" bson:"params"`
	SystemTags        []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddVmNicToSecurityGroupParams struct {
	VmNicUuids []string `json:"vmNicUuids" bson:"vmNicUuids"` //网卡的uuid列表
}

type AddVmNicToSecurityGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从安全组删除虚拟机网卡
type DeleteVmNicFromSecurityGroupRequest struct {
	ReqConfig
	SecurityGroupUuid string   `json:"securityGroupUuid" bson:"securityGroupUuid"`       //安全组UUID
	VmNicUuids        []string `json:"vmNicUuids" bson:"vmNicUuids"`                     //网卡的uuid列表
	SystemTags        []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteVmNicFromSecurityGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询应用了安全组的网卡列表
type QueryVmNicInSecurityGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVmNicInSecurityGroupResponse struct {
	Inventories []VmNicInSecurityGroupInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加规则到安全组
type AddSecurityGroupRuleRequest struct {
	ReqConfig
	SecurityGroupUuid string                     `json:"securityGroupUuid" bson:"securityGroupUuid"` //安全组UUID
	Params            AddSecurityGroupRuleParams `json:"params" bson:"params"`
	SystemTags        []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddSecurityGroupRuleParams struct {
	Rules []map[string]interface{} `json:"rules" bson:"rules"`
	//[]map[string]interface{}是结构体SecurityGroupRules，但是Protocol的值为TCP,UDP时必修填写StartPort，EndPort，为ALL,ICMP时不能有StartPort，EndPort字段，也不能为0，所以不用结构体。
	RemoteSecurityGroupUuids []string `json:"remoteSecurityGroupUuids" bson:"remoteSecurityGroupUuids"` //应用组间策略的远端安全组UUID
}

//type AddSecurityGroupRuleParams struct {
//	Rules                    []SecurityGroupRules `json:"rules" bson:"rules"`                                       //安全组中的规则
//	RemoteSecurityGroupUuids []string             `json:"remoteSecurityGroupUuids" bson:"remoteSecurityGroupUuids"` //应用组间策略的远端安全组UUID
//}

type AddSecurityGroupRuleResponse struct {
	Inventory SecurityGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询安全组规则
type QuerySecurityGroupRuleRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySecurityGroupRuleResponse struct {
	Inventories []SecurityGroupRules `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除安全组规则
type DeleteSecurityGroupRuleRequest struct {
	ReqConfig
	RuleUuids  []string `json:"ruleUuids" bson:"ruleUuids"`                       //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteSecurityGroupRuleResponse struct {
	Inventory SecurityGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type SecurityGroupInventory struct {
	UUID                   string               `json:"uuid" bson:"uuid"`
	Name                   string               `json:"name" bson:"name"`               //资源名称
	Description            string               `json:"description" bson:"description"` //详细描述
	State                  string               `json:"state" bson:"state"`
	IpVersion              int                  `json:"ipVersion" bson:"ipVersion"`   //ip协议号
	CreateDate             string               `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate             string               `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	AttachedL3NetworkUuids []string             `json:"attachedL3NetworkUuids" bson:"attachedL3NetworkUuids"`
	Rules                  []SecurityGroupRules `json:"rules" bson:"rules"`
}

type SecurityGroupRules struct {
	UUID                    string `json:"uuid" bson:"uuid"`
	SecurityGroupUuid       string `json:"securityGroupUuid" bson:"securityGroupUuid"` //安全组UUID
	Type                    string `json:"type" bson:"type"`                           //流量类型
	IpVersion               int    `json:"ipVersion" bson:"ipVersion"`                 //ip协议号
	StartPort               int    `json:"startPort" bson:"startPort"`                 //如果协议是TCP/UDP, 它是端口范围（port range）的起始端口号; 如果协议是ICMP, 它是ICMP类型（type）
	EndPort                 int    `json:"endPort" bson:"endPort"`                     //如果协议是TCP/UDP, 它是端口范围（port range）的起始端口号; 如果协议是ICMP, 它是ICMP类型（type）
	Protocol                string `json:"protocol" bson:"protocol"`                   //流量协议类型
	State                   string `json:"state" bson:"state"`                         //规则的可用状态, 当前版本未实现
	AllowedCidr             string `json:"allowedCidr" bson:"allowedCidr"`             //允许的CIDR,根据流量类型的不同, 允许的CIDR有不同的含义,如果流量类型是Ingress, 允许的CIDR是允许访问虚拟机网卡的源CIDR,如果流量类型是Egress, 允许的CIDR是允许从虚拟机网卡离开并到达的目的地CIDR
	RemoteSecurityGroupUuid string `json:"remoteSecurityGroupUuid" bson:"remoteSecurityGroupUuid"`
	CreateDate              string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate              string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type VmNicInSecurityGroupInventory struct {
	UUID              string `json:"uuid" bson:"uuid"`
	VmNicUuid         string `json:"vmNicUuid" bson:"vmNicUuid"`
	SecurityGroupUuid string `json:"securityGroupUuid" bson:"securityGroupUuid"` //安全组UUID
	VMInstanceUUID    string `json:"vmInstanceUuid" bson:"vmInstanceUuid"`       //	云主机UUID
	CreateDate        string `json:"createDate" bson:"createDate"`               //创建时间
	LastOpDate        string `json:"lastOpDate" bson:"lastOpDate"`               //最后一次修改时间
}
