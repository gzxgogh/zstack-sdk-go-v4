package model

//创建云盘规格
type CreateDiskOfferingRequest struct {
	ReqConfig
	Params     CreateDiskOfferingParams `json:"params" bson:"params"`
	SystemTags []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateDiskOfferingParams struct {
	Name               string `json:"name" bson:"name"`                                                 //资源名称
	Description        string `json:"description,omitempty" bson:"description,omitempty"`               //详细描述
	DiskSize           int64  `json:"diskSize" bson:"diskSize"`                                         //云盘大小
	SortKey            int    `json:"sortKey,omitempty" bson:"sortKey,omitempty"`                       //排序键
	AllocationStrategy string `json:"allocationStrategy,omitempty" bson:"allocationStrategy,omitempty"` //分配策略:DefaultHostAllocatorStrategy,LastHostPreferredAllocatorStrategy,
	//LeastVmPreferredHostAllocatorStrategy,MinimumCPUUsageHostAllocatorStrategy,MinimumMemoryUsageHostAllocatorStrategy,MaxInstancePerHostHostAllocatorStrategy
	Type         string `json:"type,omitempty" bson:"type,omitempty"`                 //类型
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateDiskOfferingResponse struct {
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory DiskOfferingInventory `json:"inventory" bson:"inventory"`
}

//删除云盘规格
type DeleteDiskOfferingRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteDiskOfferingResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询云盘规格
type QueryDiskOfferingRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"` //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryDiskOfferingResponse struct {
	Error       ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []DiskOfferingInventory `json:"inventories" bson:"inventories"`
}

//更改云盘规格的启用状态
type ChangeDiskOfferingStateRequest struct {
	ReqConfig
	UUID                    string                        `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeDiskOfferingState ChangeDiskOfferingStateParams `json:"changeDiskOfferingState" bson:"ChangeDiskOfferingState"`
	SystemTags              []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags                []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeDiskOfferingStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //enable,disable
}

type ChangeDiskOfferingStateResponse struct {
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory DiskOfferingInventory `json:"inventory" bson:"inventory"`
}

//更新云盘规格
type UpdateDiskOfferingRequest struct {
	ReqConfig
	UUID               string                   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateDiskOffering UpdateDiskOfferingParams `json:"updateDiskOffering" bson:"updateDiskOffering"`
	SystemTags         []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags           []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateDiskOfferingParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateDiskOfferingResponse struct {
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AffinityGroupInventory `json:"inventory" bson:"inventory"`
}

type DiskOfferingInventory struct {
	UUID               string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name               string `json:"name" bson:"name"`               //资源名称
	Description        string `json:"description" bson:"description"` //资源的详细描述
	DiskSize           int64  `json:"diskSize" bson:"diskSize"`       //云盘大小
	SortKey            int    `json:"sortKey" bson:"sortKey"`         //排序键
	State              string `json:"state" bson:"state"`
	Type               string `json:"type,omitempty" bson:"type,omitempty"`         //亲和组类型, 当前为物理机亲和,未来将增加网络亲和、路由器亲和、数据中心或机架亲和等多种类型
	CreateDate         string `json:"createDate" bson:"createDate"`                 //创建时间
	LastOpDate         string `json:"lastOpDate" bson:"lastOpDate"`                 //最后一次修改时间
	AllocationStrategy string `json:"allocationStrategy" bson:"allocationStrategy"` //分配策略
}
