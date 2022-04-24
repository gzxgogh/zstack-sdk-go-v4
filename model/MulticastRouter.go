package model

//创建组播路由器
type CreateMulticastRouterRequest struct {
	ReqConfig
	Params     CreateMulticastRouterParams `json:"params" bson:"params"`
	SystemTags []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateMulticastRouterParams struct {
	VpcRouterVmUuid string   `json:"vpcRouterVmUuid" bson:"vpcRouterVmUuid"`
	Description     string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid    string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids        []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreateMulticastRouterResponse struct {
	Inventory MulticastRouterInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除组播路由器
type DeleteMulticastRouterRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteMulticastRouterResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询组播路由器
type QueryMulticastRouterRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryMulticastRouterResponse struct {
	Inventories []MulticastRouterInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//改变组播路由器状态
type ChangeMulticastRouterStateRequest struct {
	ReqConfig
	UUID                       string                           `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeMulticastRouterState ChangeMulticastRouterStateParams `json:"changeMulticastRouterState" bson:"changeMulticastRouterState"`
	SystemTags                 []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeMulticastRouterStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"`
}

type ChangeMulticastRouterStateResponse struct {
	Inventory MulticastRouterInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取组播路由表
type GetVpcMulticastRouteRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetVpcMulticastRouteResponse struct {
	Inventories []VpcMulticastRouteInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加组播路由静态RP地址
type AddRendezvousPointToMulticastRouterRequest struct {
	ReqConfig
	UUID       string                                    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Params     AddRendezvousPointToMulticastRouterParams `json:"params" bson:"params"`
	SystemTags []string                                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddRendezvousPointToMulticastRouterParams struct {
	RpAddress    string   `json:"rpAddress" bson:"rpAddress"`
	GroupAddress string   `json:"groupAddress" bson:"groupAddress"`
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type AddRendezvousPointToMulticastRouterResponse struct {
	Inventories []VpcMulticastRouteInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除组播路由器静态RP地址
type RemoveRendezvousPointFromMulticastRouterRequest struct {
	ReqConfig
	UUID         string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode   string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	RpAddress    string   `json:"rpAddress" bson:"rpAddress"`
	GroupAddress string   `json:"groupAddress" bson:"groupAddress"`
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveRendezvousPointFromMulticastRouterResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type MulticastRouterInventory struct {
	UUID            string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	VpcRouterVmUuid string   `json:"vpcRouterVmUuid" bson:"vpcRouterVmUuid"`
	Description     string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	State           string   `json:"state" bson:"state"`
	CreateDate      string   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate      string   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	RpGroups        RpGroups `json:"rpGroups" bson:"rpGroups"`
	VpcVrs          VpcVrs   `json:"vpcVrs" bson:"vpcVrs"`
}

type RpGroups struct {
	UUID                string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	MulticastRouterUuid string `json:"multicastRouterUuid" bson:"multicastRouterUuid"`
	RpAddress           string `json:"rpAddress" bson:"rpAddress"`
	GroupAddress        string `json:"groupAddress" bson:"groupAddress"`
	CreateDate          string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate          string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type VpcVrs struct {
	UUID          string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	VpcRouterUuid string `json:"vpcRouterUuid" bson:"vpcRouterUuid"`
}

type VpcMulticastRouteInventory struct {
	SourceAddress     string `json:"sourceAddress" bson:"sourceAddress"`
	GroupAddress      string `json:"groupAddress" bson:"groupAddress"`
	IngressInterfaces string `json:"ingressInterfaces" bson:"ingressInterfaces"`
	EgressInterfaces  string `json:"egressInterfaces" bson:"egressInterfaces"`
}
