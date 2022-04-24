package model

//创建策略路由规则集
type CreatePolicyRouteRuleSetRequest struct {
	ReqConfig
	Params     CreateMulticastRouterParams `json:"params" bson:"params"`
	SystemTags []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreatePolicyRouteRuleSetParams struct {
	Name         string   `json:"name" bson:"name"`                                   //资源名称
	Description  string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	VRouterUuid  string   `json:"vRouterUuid" bson:"vRouterUuid"`
	Type         string   `json:"type,omitempty" bson:"type,omitempty"`
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreatePolicyRouteRuleSetResponse struct {
	Inventory PolicyRouteRuleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除策略路由规则集
type DeletePolicyRouteRuleSetRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeletePolicyRouteRuleSetResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新策略路由规则集属性
type UpdatePolicyRouteRuleSetRequest struct {
	ReqConfig
	UUID                     string                         `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdatePolicyRouteRuleSet UpdatePolicyRouteRuleSetParams `json:"updatePolicyRouteRuleSet" bson:"updatePolicyRouteRuleSet"`
	SystemTags               []string                       `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                 []string                       `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdatePolicyRouteRuleSetParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdatePolicyRouteRuleSetResponse struct {
	Inventory PolicyRouteRuleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询策略路由规则集
type QueryPolicyRouteRuleSetRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPolicyRouteRuleSetResponse struct {
	Inventories []PolicyRouteRuleInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询策略路由规则集网络关联
type QueryPolicyRouteRuleSetL3RefRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPolicyRouteRuleSetL3RefResponse struct {
	Inventories []L3Refs  `json:"inventories" bson:"inventories"`
	Error       ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询策略路由规则集与单节点云路由关联
type QueryPolicyRouteRuleSetVRouterRefRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPolicyRouteRuleSetVRouterRefResponse struct {
	Inventories []PolicyRouteRuleSetVRouterRef `json:"inventories" bson:"inventories"`
	Error       ErrorCode                      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建策略路由规则集规则
type CreatePolicyRouteRuleRequest struct {
	ReqConfig
	Params     CreateMulticastRouterParams `json:"params" bson:"params"`
	SystemTags []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreatePolicyRouteRuleParams struct {
	RuleSetUuid  string   `json:"ruleSetUuid" bson:"ruleSetUuid"`                   //策略路由规则集uuid
	TableUuid    string   `json:"tableUuid" bson:"tableUuid"`                       //策略路由表uuid
	RuleNumber   int      `json:"ruleNumber" bson:"ruleNumber"`                     //规则优先级
	DestIp       string   `json:"destIp,omitempty" bson:"destIp,omitempty"`         //目标ip
	SourceIp     string   `json:"sourceIp,omitempty" bson:"sourceIp,omitempty"`     //源ip
	DestPort     string   `json:"destPort,omitempty" bson:"destPort,omitempty"`     //目标端口
	SourcePort   string   `json:"sourcePort,omitempty" bson:"sourcePort,omitempty"` //源端口
	Protocol     string   `json:"protocol,omitempty" bson:"protocol,omitempty"`
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreatePolicyRouteRuleResponse struct {
	Inventory PolicyRouteRuleInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除策略路由规则集规则
type DeletePolicyRouteRuleRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeletePolicyRouteRuleResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建策略路由表
type CreatePolicyRouteTableRequest struct {
	ReqConfig
	Params     CreatePolicyRouteTableParams `json:"params" bson:"params"`
	SystemTags []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreatePolicyRouteTableParams struct {
	VRouterUuid  string   `json:"vRouterUuid" bson:"vRouterUuid"`                       //云路由uuid
	Number       string   `json:"number" bson:"number"`                                 //表名
	Description  string   `json:"description " bson:"description "`                     //资源的详细描述
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreatePolicyRouteTableResponse struct {
	Inventory PolicyRouteTableInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除策略路由表
type DeletePolicyRouteTableRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeletePolicyRouteTableResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询策略路由表
type QueryPolicyRouteTableRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPolicyRouteTableResponse struct {
	Inventories []PolicyRouteTableInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询策略路由表与单节点云路由关联
type QueryPolicyRouteTableVRouterRefeRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPolicyRouteTableVRouterRefResponse struct {
	Inventories []PolicyRouteTableVRouterRef `json:"inventories" bson:"inventories"`
	Error       ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建策略路由
type CreatePolicyRouteTableRouteEntryRequest struct {
	ReqConfig
	Params     CreatePolicyRouteTableParams `json:"params" bson:"params"`
	SystemTags []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreatePolicyRouteTableRouteEntryParams struct {
	TableUuid       string   `json:"tableUuid" bson:"tableUuid"`                           //策略路由表uuid
	DestinationCidr string   `json:"destinationCidr" bson:"destinationCidr"`               //目标ip cidr
	NextHopIp       string   `json:"nextHopIp" bson:"nextHopIp"`                           //下一跳ip
	Distance        int      `json:"distance,omitempty" bson:"distance,omitempty"`         //优先级
	ResourceUuid    string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids        []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreatePolicyRouteTableRouteEntryResponse struct {
	Inventory PolicyRouteTableRouteEntryInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除策略路由
type DeletePolicyRouteTableRouteEntryRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeletePolicyRouteTableRouteEntryResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询策略路由
type QueryPolicyRouteTableRouteEntryRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPolicyRouteTableRouteEntryResponse struct {
	Inventories []PolicyRouteTableRouteEntryInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询策略路由规则
type QueryPolicyRouteRuleRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPolicyRouteRuleResponse struct {
	Inventories []Rules   `json:"inventories" bson:"inventories"`
	Error       ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//网络加载策略路由规则集
type AttachPolicyRouteRuleSetToL3Request struct {
	ReqConfig
	L3Uuid      string                             `json:"l3Uuid" bson:"l3Uuid"`           //资源的UUID，唯一标示该资源
	RuleSetUuid string                             `json:"ruleSetUuid" bson:"ruleSetUuid"` //策略路由规则集uuid
	Params      AttachPolicyRouteRuleSetToL3Params `json:"params" bson:"params"`
	SystemTags  []string                           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string                           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachPolicyRouteRuleSetToL3Params struct {
}

type AttachPolicyRouteRuleSetToL3Response struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//网络解绑策略路由
type DetachPolicyRouteRuleSetFromL3Request struct {
	ReqConfig
	L3Uuid      string   `json:"l3Uuid" bson:"l3Uuid"`                             //资源的UUID，唯一标示该资源
	RuleSetUuid string   `json:"ruleSetUuid" bson:"ruleSetUuid"`                   //策略路由规则集uuid
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachPolicyRouteRuleSetFromL3Response struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取虚拟路由器的策略路由集合
type GetPolicyRouteRuleSetFromVirtualRouterRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetPolicyRouteRuleSetFromVirtualRouterResponse struct {
	Inventories []PolicyRouteRuleInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type PolicyRouteRuleInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Type        string `json:"type" bson:"type"`
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	Rules       Rules  `json:"rules" bson:"rules"`
	L3Refs      L3Refs `json:"L3Refs" bson:"L3Refs"`
}

type Rules struct {
	UUID        string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	RuleNumber  int    `json:"ruleNumber" bson:"ruleNumber"`   //规则优先级
	RuleSetUuid string `json:"ruleSetUuid" bson:"ruleSetUuid"` //策略路由规则集uuid
	TableUuid   string `json:"tableUuid" bson:"tableUuid"`     //策略路由表uuid
	DestIp      string `json:"destIp" bson:"destIp"`           //目标ip
	SourceIp    string `json:"sourceIp" bson:"sourceIp"`       //源ip
	DestPort    string `json:"destPort" bson:"destPort"`       //目标端口
	SourcePort  string `json:"source_port" bson:"source_port"` //源端口
	CreateDate  string `json:"createDate" bson:"createDate"`   //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"`   //最后一次修改时间
	Protocol    string `json:"protocol" bson:"protocol"`
	State       string `json:"state" bson:"state"`
}

type L3Refs struct {
	Id            int64  `json:"id" bson:"id"`
	L3NetworkUuid string `json:"l3NetworkUuid" bson:"l3NetworkUuid"`
	RuleSetUuid   string `json:"ruleSetUuid" bson:"ruleSetUuid"`
	CreateDate    string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate    string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type PolicyRouteRuleSetVRouterRef struct {
	Id          int64  `json:"id" bson:"id"`
	VRouterUuid string `json:"vRouterUuid" bson:"vRouterUuid"`
	RuleSetUuid string `json:"ruleSetUuid" bson:"ruleSetUuid"`
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type PolicyRouteTableVRouterRef struct {
	Id          int64  `json:"id" bson:"id"`
	VRouterUuid string `json:"vRouterUuid" bson:"vRouterUuid"`
	TableUuid   string `json:"tableUuid" bson:"tableUuid"`
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type PolicyRouteTableInventory struct {
	UUID        string                              `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	TableNumber string                              `json:"table_number" bson:"table_number"`
	Description string                              `json:"description " bson:"description "` //资源的详细描述
	CreateDate  string                              `json:"createDate" bson:"createDate"`     //创建时间
	LastOpDate  string                              `json:"lastOpDate" bson:"lastOpDate"`     //最后一次修改时间
	Routes      PolicyRouteTableRouteEntryInventory `json:"routes" bson:"routes"`
}

type PolicyRouteTableRouteEntryInventory struct {
	UUID            string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	TableUuid       string `json:"tableUuid" bson:"tableUuid	"`
	DestinationCidr string `json:"destinationCidr" bson:"destinationCidr"`
	NextHopIp       string `json:"nextHopIp" bson:"nextHopIp"`
	Distance        int    `json:"distance" bson:"distance"`
	CreateDate      string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}
