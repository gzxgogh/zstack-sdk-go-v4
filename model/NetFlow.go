package model

//创建流量监控搜集器
type CreateFlowCollectorRequest struct {
	ReqConfig
	FlowMeterUuid string                    `json:"flowMeterUuid" bson:"flowMeterUuid"` //流量监控资源的uuid
	Params        CreateFlowCollectorParams `json:"params" bson:"params"`
	SystemTags    []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateFlowCollectorParams struct {
	Name         string   `json:"name,omitempty" bson:"name,omitempty"`                 //资源名称
	Description  string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	Server       string   `json:"server,omitempty" bson:"server,omitempty"`             //流量搜集器的ip地址
	Port         int64    `json:"port ,omitempty" bson:"port ,omitempty"`               //流量搜集器服务的port
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type CreateFlowCollectorResponse struct {
	Inventory FlowCollectorInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询流量监控的搜集器
type QueryFlowCollectorRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryFlowCollectorResponse struct {
	Inventories []FlowCollectorInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新搜集器信息
type UpdateFlowCollectorRequest struct {
	ReqConfig
	UUID                string                    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateFlowCollector UpdateFlowCollectorParams `json:"updateFlowCollector" bson:"updateFlowCollector"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateFlowCollectorParams struct {
	Server string `json:"server,omitempty" bson:"server,omitempty"` //流量搜集器的ip地址
	Port   int64  `json:"port ,omitempty" bson:"port ,omitempty"`   //流量搜集器服务的port
}

type UpdateFlowCollectorResponse struct {
	Inventory FlowCollectorInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除流量监控搜集器
type DeleteFlowCollectorRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteFlowCollectorResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建流量监控资源
type CreateFlowMeterRequest struct {
	ReqConfig
	Params     CreateFlowMeterParams `json:"params" bson:"params"`
	SystemTags []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateFlowMeterParams struct {
	Version          string   `json:"version,omitempty" bson:"version,omitempty"`                   //流量监控协议的版本号:V5,V9
	Type             string   `json:"type" bson:"type"`                                             //流量监控使用的协议：NetFlow，sFlow
	Sample           int      `json:"sample,omitempty" bson:"sample,omitempty "`                    //流量监控的采样率
	GenerateInterval int      `json:"generateInterval,omitempty" bson:"generateInterval,omitempty"` //流量监控向搜集器发送的时间间隔
	Name             string   `json:"name,omitempty" bson:"name,omitempty"`                         //资源名称
	Description      string   `json:"description,omitempty" bson:"description,omitempty"`           //详细描述
	Server           string   `json:"server,omitempty" bson:"server,omitempty"`                     //流量搜集器的ip地址
	Port             int64    `json:"port ,omitempty" bson:"port ,omitempty"`                       //流量搜集器服务的port
	ResourceUuid     string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`         //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids         []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`                 //标签UUID列表
}

type CreateFlowMeterResponse struct {
	Inventory FlowMeterInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询流量监控资源
type QueryFlowMeterRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryFlowMeterResponse struct {
	Inventories []FlowMeterInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新流量监控的信息
type UpdateFlowMeterRequest struct {
	ReqConfig
	UUID                string                    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateFlowCollector UpdateFlowCollectorParams `json:"updateFlowCollector" bson:"updateFlowCollector"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateFlowMeterParams struct {
	Version        string `json:"version,omitempty" bson:"version,omitempty"`               //流量监控协议的版本号:V5,V9
	Sample         int64  `json:"sample,omitempty" bson:"sample,omitempty "`                //流量监控的采样率
	Name           string `json:"name,omitempty" bson:"name,omitempty"`                     //资源名称
	ExpireInterval string `json:"expireInterval,omitempty" bson:"expireInterval,omitempty"` //流量发送间隔
	Description    string `json:"description,omitempty" bson:"description,omitempty"`       //详细描述
}

type UpdateFlowMeterResponse struct {
	Inventory FlowMeterInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除流量监控资源
type DeleteFlowMeterRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteFlowMeterResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加VPC路由器网络到FlowMeter中
type AddVRouterNetworksToFlowMeterRequest struct {
	ReqConfig
	FlowMeterUuid       string                    `json:"flowMeterUuid" bson:"flowMeterUuid"` //流量监控资源的uuid
	VRouterUuid         string                    `json:"vRouterUuid" json:"vRouterUuid"`
	UpdateFlowCollector UpdateFlowCollectorParams `json:"updateFlowCollector" bson:"updateFlowCollector"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddVRouterNetworksToFlowMeterParams struct {
	L3NetworkUuids []string `json:"l3NetworkUuids" bson:"l3NetworkUuids"`                 //VPC路由器网络的uuid
	ResourceUuid   string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids       []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type AddVRouterNetworksToFlowMeterResponse struct {
	Inventory NetworkRefs `json:"inventory" bson:"inventory"`
	Error     ErrorCode   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询开启流量监控的VPC网络
type QueryVRouterFlowMeterNetworkRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVRouterFlowMeterNetworkResponse struct {
	Inventories []NetworkRefs `json:"inventories" bson:"inventories"`
	Error       ErrorCode     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//流量监控的统计信息
type GetVRouterFlowCounterRequest struct {
	ReqConfig
	VRouterUuid string   `json:"vRouterUuid" bson:"vRouterUuid"`
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetVRouterFlowCounterResponse struct {
	Counters []Counters `json:"counters" bson:"counters"`
	Error    ErrorCode  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//清除VPC路由器网络的流量监控
type RemoveVRouterNetworksFromFlowMeterRequest struct {
	ReqConfig
	UUIDS      []string `json:"uuids" bson:"uuids"`                               //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveVRouterNetworksFromFlowMeterResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取流量监控协议中使用的系统标识
type GetFlowMeterRouterIdRequest struct {
	ReqConfig
	VRouterUuid string   `json:"vRouterUuid" bson:"vRouterUuid"`
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetFlowMeterRouterIdResponse struct {
	RouterId int64     `json:"routerId" bson:"routerId"`
	Error    ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置流量监控协议使用的VPC标识
type SetFlowMeterRouterIdRequest struct {
	ReqConfig
	VRouterUuid string                     `json:"vRouterUuid" bson:"vRouterUuid"`
	Params      SetFlowMeterRouterIdParams `json:"params" bson:"params"`
	SystemTags  []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type SetFlowMeterRouterIdParams struct {
	RouterId int64 `json:"routerId" bson:"routerId"`
}

type SetFlowMeterRouterIdResponse struct {
	RouterId int64     `json:"routerId" bson:"routerId"`
	Error    ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type FlowCollectorInventory struct {
	UUID          string `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name          string `json:"name" bson:"name"`                                   //资源名称
	Description   string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	FlowMeterUuid string `json:"flowMeterUuid" bson:"flowMeterUuid"`                 //流量监控资源的uuid
	Server        string `json:"server,omitempty" bson:"server,omitempty"`           //流量搜集器的ip地址
	Port          int64  `json:"port ,omitempty" bson:"port ,omitempty"`             //流量搜集器服务的port
	CreateDate    string `json:"createDate" bson:"createDate"`
	LastOpDate    string `json:"lastOpDate" bson:"lastOpDate"`
}

type FlowMeterInventory struct {
	UUID           string        `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name           string        `json:"name" bson:"name"`                                   //资源名称
	Description    string        `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Sample         int           `json:"sample,omitempty" bson:"sample,omitempty "`          //流量监控的采样率
	ExpireInterval string        `json:"expireInterval" bson:"expireInterval"`
	Version        string        `json:"version" bson:"version"`
	Type           string        `json:"type" bson:"type"`
	CreateDate     string        `json:"createDate" bson:"createDate"`
	LastOpDate     string        `json:"lastOpDate" bson:"lastOpDate"`
	Collectors     []Collectors  `json:"collectors" bson:"collectors"`
	NetworkRefs    []NetworkRefs `json:"networkRefs" bson:"networkRefs"`
}

type Collectors struct {
	UUID          string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name          string `json:"name" bson:"name"`               //资源名称
	Description   string `json:"description" bson:"description"` //详细描述
	FlowMeterUuid string `json:"flowMeterUuid" bson:"flowMeterUuid"`
	Server        string `json:"server" bson:"server"` //流量搜集器的ip地址
	Port          int64  `json:"port" bson:"port"`     //流量搜集器服务的port
	CreateDate    string `json:"createDate" bson:"createDate"`
	LastOpDate    string `json:"lastOpDate" bson:"lastOpDate"`
}

type NetworkRefs struct {
	UUID          string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	VRouterUuid   string `json:"vRouterUuid" bson:"vRouterUuid"`
	FlowMeterUuid string `json:"flowMeterUuid" bson:"flowMeterUuid"`
	L3NetworkUuid string `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //	三层网络UUID
	CreateDate    string `json:"createDate" bson:"createDate"`
	LastOpDate    string `json:"lastOpDate" bson:"lastOpDate"`
}

type Counters struct {
	Device       string `json:"device" bson:"device"`
	TotalEntries string `json:"totalEntries" bson:"totalEntries"`
	TotalPkts    string `json:"totalPkts" bson:"totalPkts"`
	TotalBytes   string `json:"totalBytes" bson:"totalBytes"`
}
