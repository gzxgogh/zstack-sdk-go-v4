package model

//创建端口镜像
type CreatePortMirrorRequest struct {
	ReqConfig
	Params     CreateVipParams `json:"params" bson:"params"`
	SystemTags []string        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreatePortMirrorParams struct {
	MirrorNetworkUuid string   `json:"mirrorNetworkUuid" bson:"mirrorNetworkUuid"`           //镜像网络资源的UUID
	Name              string   `json:"name,omitempty" bson:"name,omitempty"`                 //资源名称
	Description       string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	StateEvent        string   `json:"stateEvent ,omitempty" bson:"stateEvent ,omitempty"`   //流量镜像服务的状态:enable,disable
	ResourceUuid      string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids          []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type CreatePortMirrorResponse struct {
	Inventory PortMirrorInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询端口镜像
type QueryPortMirrorRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPortMirrorResponse struct {
	Inventories []PortMirrorInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新端口镜像
type UpdatePortMirrorRequest struct {
	ReqConfig
	UUID             string                 `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdatePortMirror UpdatePortMirrorParams `json:"updatePortMirror" bson:"updatePortMirror"`
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdatePortMirrorParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdatePortMirrorResponse struct {
	Inventory PortMirrorInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除端口镜像
type DeletePortMirrorRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeletePortMirrorResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建端口镜像会话
type CreatePortMirrorSessionRequest struct {
	ReqConfig
	Params     CreatePortMirrorSessionParams `json:"params" bson:"params"`
	SystemTags []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreatePortMirrorSessionParams struct {
	PortMirrorUuid string   `json:"portMirrorUuid" bson:"portMirrorUuid"`               //镜像网络资源的UUID
	Name           string   `json:"name" bson:"name"`                                   //资源名称
	Description    string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Type           string   `json:"type" bson:"type"`
	SrcEndPoint    string   `json:"srcEndPoint" bson:"srcEndPoint"`
	DstEndPoint    string   `json:"dstEndPoint" bson:"dstEndPoint"`
	ResourceUuid   string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids       []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type CreatePortMirrorSessionResponse struct {
	Inventory PortMirrorSessionInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询端口镜像会话
type QueryPortMirrorSessionRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPortMirrorSessionResponse struct {
	Inventories []PortMirrorInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新端口镜像服务状态
type ChangePortMirrorStateRequest struct {
	ReqConfig
	UUID                  string                      `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangePortMirrorState ChangePortMirrorStateParams `json:"changePortMirrorState" bson:"changePortMirrorState"`
	SystemTags            []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags              []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangePortMirrorStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //资源名称
}

type ChangePortMirrorStateResponse struct {
	Inventory PortMirrorInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除端口镜像会话
type DeletePortMirrorSessionRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeletePortMirrorSessionResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取系统中可以使用端口镜像服务的网卡
type GetCandidateVmNicsForPortMirrorRequest struct {
	ReqConfig
	PortMirrorUuid string   `json:"portMirrorUuid" bson:"portMirrorUuid"`
	Type           string   `json:"type" bson:"type"`
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetCandidateVmNicsForPortMirrorResponse struct {
	Inventories []VmNicInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询端口镜像网络分配的IP
type QueryPortMirrorNetworkUsedIpRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPortMirrorNetworkUsedIpResponse struct {
	Inventories []PortMirrorNetworkUsedIpInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type PortMirrorInventory struct {
	UUID              string                     `json:"uuid" bson:"uuid"`                           //资源的UUID，唯一标示该资源
	Name              string                     `json:"name" bson:"name"`                           //资源名称
	Description       string                     `json:"description" bson:"description"`             //详细描述
	MirrorNetworkUuid string                     `json:"mirrorNetworkUuid" bson:"mirrorNetworkUuid"` //资源名称
	CreateDate        string                     `json:"createDate" bson:"createDate"`
	LastOpDate        string                     `json:"lastOpDate" bson:"lastOpDate"`
	State             string                     `json:"state" bson:"state"`
	Sessions          PortMirrorSessionInventory `json:"sessions" bson:"sessions"`
}

type PortMirrorSessionInventory struct {
	UUID           string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name           string `json:"name" bson:"name"`               //资源名称
	Description    string `json:"description" bson:"description"` //详细描述
	InternalId     int64  `json:"internalId" bson:"internalId"`
	SrcEndPoint    string `json:"srcEndPoint" bson:"srcEndPoint"`
	DstEndPoint    string `json:"dstEndPoint" bson:"dstEndPoint"`
	PortMirrorUuid string `json:"portMirrorUuid" bson:"portMirrorUuid"`
	CreateDate     string `json:"createDate" bson:"createDate"`
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"`
	Status         string `json:"status" bson:"status"`
	Type           string `json:"type" bson:"type"`
}

type PortMirrorNetworkUsedIpInventory struct {
	UUID            string `json:"uuid" bson:"uuid"`                             //资源的UUID，唯一标示该资源
	HostUuid        string `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"` //物理机UUID 若指定，云主机会在指定物理机创建，该字段优先级高于zoneUuid和clusterUuid。
	L3NetworkUUID   string `json:"l3NetworkUuid" bson:"l3NetworkUuid"`           //	三层网络UUID
	Description     string `json:"description" bson:"description"`               //详细描述
	ClusterUuid     string `json:"clusterUuid" bson:"clusterUuid"`
	UsedIpInventory UsedIp `json:"usedIpInventory" bson:"usedIpInventory"`
}
