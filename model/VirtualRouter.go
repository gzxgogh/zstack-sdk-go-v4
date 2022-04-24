package model

//重连云路由器
type ReconnectVirtualRouterRequest struct {
	ReqConfig
	VmInstanceUuid string                       `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	Params         ReconnectVirtualRouterParams `json:"params" bson:"params"`
	SystemTags     []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ReconnectVirtualRouterParams struct {
}

type ReconnectVirtualRouterResponse struct {
	Inventory ApplianceVmInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询云路由器
type QueryVirtualRouterVmRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVirtualRouterVmResponse struct {
	Inventory ApplianceVmInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询服务云主机
type QueryApplianceVmRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryApplianceVmResponse struct {
	Inventories []ApplianceVmInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建云路由规格
type CreateVirtualRouterOfferingRequest struct {
	ReqConfig
	Uuid       string                            `json:"uuid" bson:"uuid"`
	Params     CreateVirtualRouterOfferingParams `json:"params" bson:"params"`
	SystemTags []string                          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateVirtualRouterOfferingParams struct {
	ZoneUuid              string `json:"zoneUuid" bson:"zoneUuid"`                                       //区域UUID 若指定，云主机会在指定区域创建。
	ManagementNetworkUuid string `json:"managementNetworkUuid" bson:"managementNetworkUuid"`             //管理L3网络UUID
	ImageUUID             string `json:"imageUuid" bson:"imageUuid"`                                     //镜像UUID
	PublicNetworkUuid     string `json:"publicNetworkUuid,omitempty" bson:"publicNetworkUuid,omitempty"` //公有L3网络UUID
	IsDefault             bool   `json:"isDefault,omitempty" bson:"isDefault,omitempty"`                 //默认
	Name                  string `json:"name" bson:"name"`                                               //资源名称
	Description           string `json:"description,omitempty" bson:"description,omitempty"`             //资源的详细描述
	MemorySize            int64  `json:"memorySize" bson:"memorySize"`                                   //内存大小
	CPUNum                int    `json:"cpuNum" bson:"cpuNum"`                                           //cpu数量
	AllocatorStrategy     string `json:"allocatorStrategy,omitempty" bson:"allocatorStrategy,omitempty"` //分配策略
	SortKey               int    `json:"sortKey,omitempty" bson:"sortKey,omitempty"`                     //排序主键
	Type                  string `json:"type,omitempty" bson:"type,omitempty"`                           //类型
	ResourceUuid          string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`           //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateVirtualRouterOfferingResponse struct {
	Inventory InstanceOfferingInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询云路由规格
type QueryVirtualRouterOfferingRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVirtualRouterOfferingResponse struct {
	Inventories []InstanceOfferingInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新云路由规格
type UpdateVirtualRouterOfferingRequest struct {
	ReqConfig
	Uuid                        string                            `json:"uuid" bson:"uuid"`
	UpdateVirtualRouterOffering UpdateVirtualRouterOfferingParams `json:"updateVirtualRouterOffering" bson:"updateVirtualRouterOffering"`
	SystemTags                  []string                          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                    []string                          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateVirtualRouterOfferingParams struct {
	IsDefault         bool   `json:"isDefault,omitempty" bson:"isDefault,omitempty"`                 //默认
	ImageUUID         string `json:"imageUuid,omitempty" bson:"imageUuid,omitempty"`                 //镜像UUID
	Name              string `json:"name,omitempty" bson:"name,omitempty"`                           //资源名称
	Description       string `json:"description,omitempty" bson:"description,omitempty"`             //资源的详细描述
	AllocatorStrategy string `json:"allocatorStrategy,omitempty" bson:"allocatorStrategy,omitempty"` //分配策略
}

type UpdateVirtualRouterOfferingResponse struct {
	Inventory InstanceOfferingInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云路由器可加载外部网络
type GetAttachablePublicL3ForVRouterRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetAttachablePublicL3ForVRouterResponse struct {
	Inventories L3NetworkInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建云路由路由表
type CreateVRouterRouteTableRequest struct {
	ReqConfig
	Params     CreateVRouterRouteTableParams `json:"params" bson:"params"`
	SystemTags []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateVRouterRouteTableParams struct {
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description" bson:"description"`                       //资源的详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateVRouterRouteTableResponse struct {
	Inventory VRouterRouteTableInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除云路由路由表
type DeleteVRouterRouteTableRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteVRouterRouteTableResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询云路由路由表
type QueryVRouterRouteTableRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVRouterRouteTableResponse struct {
	Inventories []VRouterRouteTableInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取路由器实时路由表
type GetVRouterRouteTableRequest struct {
	ReqConfig
	VirtualRouterVmUuid string   `json:"virtualRouterVmUuid" bson:"virtualRouterVmUuid"`   //资源的UUID，唯一标示该资源
	SystemTags          []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetVRouterRouteTableResponse struct {
	Inventories []VRouterRouteTableInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加云路由路由条目
type AddVRouterRouteEntryRequest struct {
	ReqConfig
	RouteTableUuid string                     `json:"routeTableUuid" bson:"routeTableUuid"` //资源的UUID，唯一标示该资源
	Params         AddVRouterRouteEntryParams `json:"params" bson:"params"`
	SystemTags     []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddVRouterRouteEntryParams struct {
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //资源的详细描述
	Type         string `json:"type,omitempty" bson:"type,omitempty"`                 //允许用户添加"静态路由"、"黑洞路由"两种类型，系统会根据是否填下一跳自动判断类型
	Destination  string `json:"destination,omitempty" bson:"destination,omitempty"`   //目标网络地址，使用网络地址CIDR格式，如果用户填写的不是标准CIDR格式，系统会自动转换
	Target       string `json:"target,omitempty" bson:"target,omitempty"`             //为一个云路由设备目前可以直接到达的IP地址，如果不可以直接到达，将会进行递归路由
	Distance     string `json:"distance,omitempty" bson:"distance,omitempty"`         //路由优先级，在最小匹配下如果有多条路由规则匹配，优先级数字小的规则将会被匹配
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type AddVRouterRouteEntryResponse struct {
	Inventory VRouterRouteEntryInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除云路由路由条目
type DeleteVRouterRouteEntryRequest struct {
	ReqConfig
	UUID           string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	RouteTableUuid string   `json:"routeTableUuid" bson:"routeTableUuid"`
	DeleteMode     string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteVRouterRouteEntryResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询云路由路由条目
type QueryVRouterRouteEntryRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVRouterRouteEntryResponse struct {
	Inventories []VRouterRouteEntryInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//绑定云路由路由表到云路由器
type AttachVRouterRouteTableToVRouterRequest struct {
	ReqConfig
	RouteTableUuid string                                 `json:"routeTableUuid" bson:"routeTableUuid"` //资源的UUID，唯一标示该资源
	Params         AttachVRouterRouteTableToVRouterParams `json:"params" bson:"params"`
	SystemTags     []string                               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string                               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachVRouterRouteTableToVRouterParams struct {
	VirtualRouterVmUuid string `json:"virtualRouterVmUuid" bson:"virtualRouterVmUuid"`
}

type AttachVRouterRouteTableToVRouterResponse struct {
	Inventory VRouterRouteTableInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//解绑云路由路由表
type DetachVRouterRouteTableFromVRouterRequest struct {
	ReqConfig
	RouteTableUuid      string   `json:"routeTableUuid" bson:"routeTableUuid"` //资源的UUID，唯一标示该资源
	VirtualRouterVmUuid string   `json:"virtualRouterVmUuid" bson:"virtualRouterVmUuid"`
	SystemTags          []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachVRouterRouteTableFromVRouterResponse struct {
	Inventory VRouterRouteTableInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询绑定关系
type QueryVirtualRouterVRouterRouteTableRefRequest struct {
	ReqConfig
	//UUID       string   `json:"uuid" bson:"uuid"`      //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVirtualRouterVRouterRouteTableRefResponse struct {
	Inventories []AttachedRouterRefs `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type ApplianceVmInventory struct {
	ApplianceVmType           string            `json:"applianceVmType" bson:"applianceVmType"`
	ManagementNetworkUuid     string            `json:"managementNetworkUuid" bson:"managementNetworkUuid"`
	DefaultRouteL3NetworkUuid string            `json:"defaultRouteL3NetworkUuid" bson:"defaultRouteL3NetworkUuid"`
	Status                    string            `json:"status" bson:"status"` //云盘的连接状态
	AgentPort                 int               `json:"agentPort" bson:"agentPort"`
	UUID                      string            `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name                      string            `json:"name" bson:"name"`                                   //资源名称
	Description               string            `json:"description" bson:"description"`                     //资源的详细描述
	ZoneUuid                  string            `json:"zoneUuid,omitempty" bson:"zoneUuid,omitempty"`       //区域UUID 若指定，云主机会在指定区域创建。
	ClusterUUID               string            `json:"clusterUuid,omitempty" bson:"clusterUuid,omitempty"` //集群UUID 若指定，云主机会在指定集群创建，该字段优先级高于zoneUuid。
	ImageUUID                 string            `json:"imageUuid" bson:"imageUuid"`                         //镜像UUID 云主机的根云盘会从该字段指定的镜像创建。
	HostUuid                  string            `json:"hostUuid" bson:"hostUuid"`                           //物理机UUID
	LastHostUUID              string            `json:"lastHostUuid" bson:"lastHostUuid"`                   //上一次运行云主机的物理机UUID
	InstanceOfferingUUID      string            `json:"instanceOfferingUuid" bson:"instanceOfferingUuid"`   //计算规格UUID 指定云主机的CPU、内存等参数。
	RootVolumeUuid            string            `json:"rootVolumeUuid" bson:"rootVolumeUuid"`
	Platform                  string            `json:"platform" bson:"platform"`
	DefaultL3NetworkUuid      string            `json:"defaultL3NetworkUuid" bson:"defaultL3NetworkUuid"`
	Type                      string            `json:"type" bson:"type"`
	HypervisorType            string            `json:"hypervisorType" bson:"hypervisorType"` //虚拟机管理程序类型,KVM Simulator
	MemorySize                int64             `json:"memorySize" bson:"memorySize"`         //内存大小
	CPUNum                    int               `json:"cpuNum" bson:"cpuNum"`                 //cpu数量
	CPUSpeed                  int64             `json:"cpuSpeed" bson:"cpuSpeed"`             //cpu主频
	AllocatorStrategy         string            `json:"allocatorStrategy,omitempty" bson:"allocatorStrategy,omitempty"`
	CreateDate                string            `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate                string            `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	State                     string            `json:"state" bson:"state"`
	VMNics                    []VmNicInventory  `json:"vmNics" bson:"vmNics"`         //所有网卡信息
	AllVolumes                []VolumeInventory `json:"allVolumes" bson:"allVolumes"` //所有卷
}

type VRouterRouteTableInventory struct {
	UUID               string                     `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name               string                     `json:"name" bson:"name"`               //资源名称
	Description        string                     `json:"description" bson:"description"` //资源的详细描述
	CreateDate         string                     `json:"createDate" bson:"createDate"`   //创建时间
	LastOpDate         string                     `json:"lastOpDate" bson:"lastOpDate"`   //最后一次修改时间
	AttachedRouterRefs AttachedRouterRefs         `json:"attachedRouterRefs" bson:"attachedRouterRefs"`
	RouteEntries       VRouterRouteEntryInventory `json:"routeEntries" bson:"routeEntries"`
}

type AttachedRouterRefs struct {
	VirtualRouterVmUuid string `json:"virtualRouterVmUuid" bson:"virtualRouterVmUuid"`
	RouteTableUuid      string `json:"routeTableUuid" bson:"routeTableUuid"` //云路由路由表UUID
}

type VRouterRouteEntryInventory struct {
	UUID           string `json:"uuid" bson:"uuid"`                     //资源的UUID，唯一标示该资源
	Description    string `json:"description" bson:"description"`       //资源的详细描述
	Type           string `json:"type" bson:"type"`                     //允许用户添加"静态路由"、"黑洞路由"两种类型，系统会根据是否填下一跳自动判断类型
	RouteTableUuid string `json:"routeTableUuid" bson:"routeTableUuid"` //云路由路由表UUID
	Destination    string `json:"destination" bson:"destination"`       //目标网络地址
	Target         string `json:"target" bson:"target"`                 //为一个云路由设备目前可以直接到达的IP地址，如果不可以直接到达，将会进行递归路由
	Distance       string `json:"distance" bson:"distance"`             //路由优先级，在最小匹配下如果有多条路由规则匹配，优先级数字小的规则将会被匹配
	CreateDate     string `json:"createDate" bson:"createDate"`         //创建时间
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"`         //最后一次修改时间

}
