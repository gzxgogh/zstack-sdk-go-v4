package model

//获取网络服务类型
type GetNetworkServiceTypesRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetNetworkServiceTypesResponse struct {
	Types map[string]interface{} `json:"types" bson:"types"`
	Error ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询网络服务模块
type QueryNetworkServiceProviderRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryNetworkServiceProviderResponse struct {
	Inventories []NetworkServiceProviderInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询网络服务与三层网络引用
type QueryNetworkServiceL3NetworkRefRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryNetworkServiceL3NetworkRefResponse struct {
	Inventories []NetworkServiceL3NetworkRefInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//挂载网络服务到三层网络
type AttachNetworkServiceToL3NetworkRequest struct {
	ReqConfig
	L3NetworkUUID string                                `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //	三层网络UUID
	Params        AttachNetworkServiceToL3NetworkParams `json:"params" bson:"params"`
	SystemTags    []string                              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string                              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachNetworkServiceToL3NetworkParams struct {
	NetworkServices map[string]interface{} `json:"networkServiceProviderUuid" bson:"networkServiceProviderUuid"` //网络服务提供模块UUID
}

type AttachNetworkServiceToL3NetworkResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从三层网络卸载网络服务
type DetachNetworkServiceFromL3NetworkRequest struct {
	ReqConfig
	L3NetworkUUID   string                 `json:"l3NetworkUuid" bson:"l3NetworkUuid"`                           //	三层网络UUID
	NetworkServices map[string]interface{} `json:"networkServiceProviderUuid" bson:"networkServiceProviderUuid"` //网络服务提供模块UUID
	SystemTags      []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"`             //云主机系统标签
	UserTags        []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachNetworkServiceFromL3NetworkResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type NetworkServiceProviderInventory struct {
	UUID                   string   `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name                   string   `json:"name" bson:"name"`                                   //资源名称
	Description            string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Type                   string   `json:"type,omitempty" bson:"type,omitempty"`               //保留域, 请不要使用,zstack
	CreateDate             string   `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate             string   `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
	NetworkServiceTypes    []string `json:"networkServiceTypes" bson:"networkServiceTypes"`
	AttachedL2NetworkUuids []string `json:"attachedL2NetworkUuids" bson:"attachedL2NetworkUuids"`
}

type NetworkServiceL3NetworkRefInventory struct {
	L3NetworkUUID              string `json:"l3NetworkUuid" bson:"l3NetworkUuid"`                           //	三层网络UUID
	NetworkServiceProviderUuid string `json:"networkServiceProviderUuid" bson:"networkServiceProviderUuid"` //网络服务提供模块UUID
	NetworkServiceType         string `json:"networkServiceType" bson:"networkServiceType"`
}
