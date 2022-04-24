package model

//创建虚拟IP
type CreateVipRequest struct {
	ReqConfig
	Params     CreateVipParams `json:"params" bson:"params"`
	SystemTags []string        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateVipParams struct {
	Name              string `json:"name" bson:"name"`                                               //资源名称
	Description       string `json:"description,omitempty" bson:"description,omitempty"`             //详细描述
	L3NetworkUUID     string `json:"l3NetworkUuid" bson:"l3NetworkUuid"`                             //	三层网络UUID
	IpRangeUUID       string `json:"ipRangeUuid,omitempty" bson:"ipRangeUuid,omitempty"`             //IP段UUID
	AllocatorStrategy string `json:"allocatorStrategy,omitempty" bson:"allocatorStrategy,omitempty"` //分配策略
	RequiredIp        string `json:"requiredIp,omitempty" bson:"requiredIp,omitempty"`               //请求的IP
	ResourceUuid      string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`           //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateVipResponse struct {
	Inventory VipInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除虚拟IP
type DeleteVipRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteVipResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询虚拟IP
type QueryVipRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVipResponse struct {
	Inventories []VipInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新虚拟IP
type UpdateVipRequest struct {
	ReqConfig
	UUID       string          `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateVip  UpdateVipParams `json:"updateVip" bson:"updateVip"`
	SystemTags []string        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateVipParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateVipResponse struct {
	Inventory VipInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更改虚拟IP启用状态
type ChangeVipStateRequest struct {
	ReqConfig
	UUID           string               `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeVipState ChangeVipStateParams `json:"changeVipState" bson:"changeVipState"`
	SystemTags     []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeVipStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"`
}

type ChangeVipStateResponse struct {
	Inventory VipInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取虚拟IP所有业务端口列表
type GetVipUsedPortsRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	Protocol   string   `json:"protocol" bson:"protocol"`                         //协议:tcp,dcp
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetVipUsedPortsResponse struct {
	Inventories []VipUsedPortsInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置虚拟IP限速
type SetVipQosRequest struct {
	ReqConfig
	UUID       string          `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	SetVipQos  SetVipQosParams `json:"setVipQos" bson:"setVipQos"`
	SystemTags []string        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type SetVipQosParams struct {
	Port              int   `json:"port,omitempty" bson:"port,omitempty"`                           //端口
	OutboundBandwidth int64 `json:"outboundBandwidth,omitempty" bson:"outboundBandwidth,omitempty"` //出流量带宽限制
	InboundBandwidth  int64 `json:"inboundBandwidth,omitempty" bson:"inboundBandwidth,omitempty"`   //入流量带宽限制
}

type SetVipQosResponse struct {
	Inventory VipQosInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取虚拟IP限速
type GetVipQosRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetVipQosResponse struct {
	Inventories []VipQosInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//取消虚拟IP限速
type DeleteVipQosRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Port       int      `json:"port,omitempty" bson:"port,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteVipQosResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type VipInventory struct {
	UUID               string         `json:"uuid" bson:"uuid"`
	Name               string         `json:"name" bson:"name"`                                   //资源名称
	Description        string         `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	L3NetworkUUID      string         `json:"l3NetworkUuid" bson:"l3NetworkUuid"`                 //	三层网络UUID
	Ip                 string         `json:"ip" bson:"ip"`
	State              string         `json:"state" bson:"state"`
	Gateway            string         `json:"gateway" bson:"gateway"`                       //网关
	Netmask            string         `json:"netmask" bson:"netmask"`                       //子网掩码
	PrefixLen          string         `json:"prefixLen" bson:"prefixLen"`                   //掩码长度
	ServiceProvider    string         `json:"serviceProvider" bson:"serviceProvider"`       //提供VIP服务的服务提供者
	PeerL3NetworkUuids string         `json:"peerL3NetworkUuids" bson:"peerL3NetworkUuids"` //提供VIP服务的L3网络UUID
	UseFor             string         `json:"useFor" bson:"useFor"`                         //用途，例如：端口转发
	System             bool           `json:"system" bson:"system"`                         //是否系统创建
	CreateDate         string         `json:"createDate" bson:"createDate"`                 //创建时间
	LastOpDate         string         `json:"lastOpDate" bson:"lastOpDate"`                 //最后一次修改时间
	ServicesRefs       []ServicesRefs `json:"servicesRefs" bson:"servicesRefs"`
}

type ServicesRefs struct {
	UUID        string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	ServiceType string `json:"serviceType" bson:"serviceType"` //服务类型
	VipUuid     string `json:"vipUuid" bson:"vipUuid"`
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type VipUsedPortsInventory struct {
	UUID      string   `json:"uuid" bson:"uuid"`         //资源的UUID，唯一标示该资源
	Protocol  string   `json:"protocol" bson:"protocol"` //协议:tcp,dcp
	UsedPorts []string `json:"usedPorts" bson:"usedPorts"`
}

type VipQosInventory struct {
	UUID              string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	VipUuid           string `json:"vipUuid" bson:"vipUuid"`
	Port              int    `json:"port,omitempty" bson:"port,omitempty"`                           //端口
	OutboundBandwidth int64  `json:"outboundBandwidth,omitempty" bson:"outboundBandwidth,omitempty"` //出流量带宽限制
	InboundBandwidth  int64  `json:"inboundBandwidth,omitempty" bson:"inboundBandwidth,omitempty"`   //入流量带宽限制
	Type              string `json:"type" bson:"type"`
	CreateDate        string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate        string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}
