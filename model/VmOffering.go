package model

//创建云主机规格
type CreateInstanceOfferingRequest struct {
	ReqConfig
	Params     CreateInstanceOfferingParams `json:"params" bson:"params"`
	SystemTags []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateInstanceOfferingParams struct {
	Name              string `json:"name" bson:"name"`                                   //资源名称
	Description       string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	CPUNum            int    `json:"cpuNum" bson:"cpuNum"`                               //cpu数量
	MemorySize        int64  `json:"memorySize" bson:"memorySize"`                       //内存大小
	AllocatorStrategy string `json:"allocatorStrategy,omitempty" bson:"allocatorStrategy,omitempty"`
	//分配策略:DefaultHostAllocatorStrategy,LastHostPreferredAllocatorStrategy,LeastVmPreferredHostAllocatorStrategy,MinimumCPUUsageHostAllocatorStrategy,MinimumMemoryUsageHostAllocatorStrategy,MaxInstancePerHostHostAllocatorStrategy
	SortKey      int    `json:"sortKey,omitempty" bson:"sortKey,omitempty"`           //排序键
	Type         string `json:"type,omitempty" bson:"type,omitempty"`                 //类型
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID
}

type CreateInstanceOfferingResponse struct {
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory InstanceOfferingInventory `json:"inventory" bson:"inventory"`
}

//删除云主机规格
type DeleteInstanceOfferingRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteInstanceOfferingResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询云主机规格
type QueryInstanceOfferingRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryInstanceOfferingResponse struct {
	Error       ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []InstanceOfferingInventory `json:"inventories" bson:"inventories"`
}

//更改云主机规格
type ChangeInstanceOfferingRequest struct {
	ReqConfig
	VmInstanceUuid               string                       `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //云主机UUID
	ChangeInstanceOfferingParams ChangeInstanceOfferingParams `json:"changeInstanceOffering" bson:"changeInstanceOffering"`
	SystemTags                   []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                     []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeInstanceOfferingParams struct {
	InstanceOfferingUuid string `json:"instanceOfferingUuid" bson:"instanceOfferingUuid"`
}

type ChangeInstanceOfferingResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新云主机规格
type UpdateInstanceOfferingRequest struct {
	ReqConfig
	Uuid                   string                       `json:"uuid" bson:"uuid"` //云主机UUID
	UpdateInstanceOffering UpdateInstanceOfferingParams `json:"updateInstanceOffering" bson:"updateInstanceOffering"`
	SystemTags             []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags               []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateInstanceOfferingParams struct {
	Name              string `json:"name" bson:"name"`                                   //资源名称
	Description       string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	AllocatorStrategy string `json:"allocatorStrategy,omitempty" bson:"allocatorStrategy,omitempty"`
}

type UpdateInstanceOfferingResponse struct {
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory InstanceOfferingInventory `json:"inventory" bson:"inventory"`
}

//更改云主机规格的启用状态
type ChangeInstanceOfferingStateRequest struct {
	ReqConfig
	Uuid                        string                            `json:"uuid" bson:"uuid"` //云主机UUID
	ChangeInstanceOfferingState ChangeInstanceOfferingStateParams `json:"changeInstanceOfferingState" bson:"changeInstanceOfferingState"`
	SystemTags                  []string                          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                    []string                          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeInstanceOfferingStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //状态事件 :enable,disable
}

type ChangeInstanceOfferingStateResponse struct {
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory InstanceOfferingInventory `json:"inventory" bson:"inventory"`
}

type InstanceOfferingInventory struct {
	UUID              string `json:"uuid" bson:"uuid"`                           //资源的UUID，唯一标示该资源
	Name              string `json:"name" bson:"name"`                           //资源名称
	Description       string `json:"description" bson:"description"`             //资源的详细描述
	CPUNum            int    `json:"cpuNum" bson:"cpuNum"`                       //cpu数量
	CpuSpeed          int    `json:"cpuSpeed" bson:"cpuSpeed"`                   //CPU速度
	MemorySize        int64  `json:"memorySize" bson:"memorySize"`               //内存大小
	Type              string `json:"type,omitempty" bson:"type,omitempty"`       //类型
	AllocatorStrategy string `json:"allocatorStrategy" bson:"allocatorStrategy"` //分配策略
	SortKey           int    `json:"sortKey,omitempty" bson:"sortKey,omitempty"` //排序键
	CreateDate        string `json:"createDate" bson:"createDate"`               //创建时间
	LastOpDate        string `json:"lastOpDate" bson:"lastOpDate"`               //最后一次修改时间
	State             string `json:"state" bson:"state"`                         //状态（启用，禁用）
}
