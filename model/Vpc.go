package model

//创建VPC路由器
type CreateVpcVRouterRequest struct {
	ReqConfig
	Params     CreateVpcVRouterParams `json:"params" bson:"params"`
	SystemTags []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateVpcVRouterParams struct {
	Name                      string `json:"name" bson:"name"` //资源名称
	VirtualRouterOfferingUuid string `json:"virtualRouterOfferingUuid" bson:"virtualRouterOfferingUuid"`
	Description               string `json:"description,omitempty" bson:"description,omitempty"`   //资源的详细描述
	ResourceUuid              string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateVpcVRouterResponse struct {
	Inventory VpcRouterVmInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询VPC路由器
type QueryVpcRouterRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVpcRouterResponse struct {
	Inventories []VpcRouterVmInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取VPC路由器可加载的三层网络
type GetAttachableVpcL3NetworkRequest struct {
	ReqConfig
	UUID       string                          `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Params     GetAttachableVpcL3NetworkParams `json:"params" bson:"params"`
	SystemTags []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetAttachableVpcL3NetworkParams struct {
}

type GetAttachableVpcL3NetworkResponse struct {
	Inventories []VpcRouterVmInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取实时流量状态
type GetVpcVRouterDistributedRoutingConnectionsRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源,要开启分布式路由，才能获取到VPC路由器实时流量状态。
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetVpcVRouterDistributedRoutingConnectionsResponse struct {
	Inventories map[string]interface{} `json:"inventories" bson:"inventories"`
	Error       ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取分布式路由是否打开
type GetVpcVRouterDistributedRoutingEnabledRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetVpcVRouterDistributedRoutingEnabledResponse struct {
	Enable bool      `json:"enable" bson:"enable"`
	Error  ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置分布式路由开关
type SetVpcVRouterDistributedRoutingEnabledRequest struct {
	ReqConfig
	UUID       string                                       `json:"uuid" bson:"uuid"` //资源的UUID
	Params     SetVpcVRouterDistributedRoutingEnabledParams `json:"params" bson:"params"`
	SystemTags []string                                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type SetVpcVRouterDistributedRoutingEnabledParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"`
}

type SetVpcVRouterDistributedRoutingEnabledResponse struct {
	Enable bool      `json:"enable" bson:"enable"`
	Error  ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//向VPC路由器添加DNS
type AddDnsToVpcRouterRequest struct {
	ReqConfig
	UUID       string                  `json:"uuid" bson:"uuid"` //资源的UUID
	Params     AddDnsToVpcRouterParams `json:"params" bson:"params"`
	SystemTags []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddDnsToVpcRouterParams struct {
	Dns          string `json:"dns" bson:"dns"`
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //用户指定的资源uuid
}

type AddDnsToVpcRouterResponse struct {
	Inventory VpcRouterVmInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从VPC路由器移除DNS
type RemoveDnsFromVpcRouterRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"` //资源的UUID
	Dns        string   `json:"dns" bson:"dns"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveDnsFromVpcRouterResponse struct {
	Inventory VpcRouterVmInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取VPC路由器的网络服务状态
type GetVpcVRouterNetworkServiceStateRequest struct {
	ReqConfig
	UUID           string   `json:"uuid" bson:"uuid"` //资源的UUID
	NetworkService string   `json:"networkService" bson:"networkService"`
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetVpcVRouterNetworkServiceStateResponse struct {
	State string    `json:"state" bson:"state"`
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置VPC路由器的网络服务
type SetVpcVRouterNetworkServiceStateRequest struct {
	ReqConfig
	UUID       string                                 `json:"uuid" bson:"uuid"` //资源的UUID
	Params     SetVpcVRouterNetworkServiceStateParams `json:"params" bson:"params"`
	SystemTags []string                               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type SetVpcVRouterNetworkServiceStateParams struct {
	NetworkService string `json:"networkService" bson:"networkService"` //路由器提供的网络服务:SNAT
	State          string `json:"state" bson:"state"`                   //服务状态:enable,disable
}

type SetVpcVRouterNetworkServiceStateResponse struct {
	State string    `json:"state" bson:"state"`
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新高可用组仲裁ip
type ChangeVpcHaGroupMonitorIpsRequest struct {
	ReqConfig
	UUID                       string                           `json:"uuid" bson:"uuid"` //资源的UUID
	ChangeVpcHaGroupMonitorIps ChangeVpcHaGroupMonitorIpsParams `json:"changeVpcHaGroupMonitorIps" bson:"changeVpcHaGroupMonitorIps"`
	SystemTags                 []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeVpcHaGroupMonitorIpsParams struct {
	MonitorIps []string `json:"monitorIps " bson:"monitorIps "` //
}

type ChangeVpcHaGroupMonitorIpsResponse struct {
	Inventory VpcHaGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建高可用组
type CreateVpcHaGroupRequest struct {
	ReqConfig
	UUID       string                           `json:"uuid" bson:"uuid"` //资源的UUID
	Params     ChangeVpcHaGroupMonitorIpsParams `json:"params" bson:"params"`
	SystemTags []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateVpcHaGroupParams struct {
	Name         string   `json:"name" bson:"name"`                                   //资源名称
	Description  string   `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
	MonitorIp    string   `json:"monitorIp,omitempty" bson:"monitorIp,omitempty"`
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreateVpcHaGroupResponse struct {
	Inventory VpcHaGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除高可用组
type DeleteVpcHaGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteVpcHaGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新高可用组
type UpdateVpcHaGroupRequest struct {
	ReqConfig
	UUID             string                 `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateVpcHaGroup UpdateVpcHaGroupParams `json:"updateVpcHaGroup" bson:"updateVpcHaGroup"`
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateVpcHaGroupParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
}

type UpdateVpcHaGroupResponse struct {
	Inventory VpcHaGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询高可用
type QueryVpcHaGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVpcHaGroupResponse struct {
	Inventories []VpcHaGroupInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新虚拟路由器
type UpdateVirtualRouterRequest struct {
	ReqConfig
	VmInstanceUuid      string                    `json:"vm_instance_uuid" bson:"vm_instance_uuid"`
	UpdateVirtualRouter UpdateVirtualRouterParams `json:"updateVirtualRouter" bson:"updateVirtualRouter"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateVirtualRouterParams struct {
	DefaultRouteL3NetworkUuid string `json:"defaultRouteL3NetworkUuid " bson:"defaultRouteL3NetworkUuid "`
}

type UpdateVirtualRouterResponse struct {
	Inventory VirtualRouterVmInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type VpcRouterVmInventory struct {
	PublicNetworkUuid         string            `json:"publicNetworkUuid" bson:"publicNetworkUuid"`
	VirtualRouterVips         []string          `json:"virtualRouterVips" bson:"virtualRouterVips"`
	ApplianceVmType           string            `json:"applianceVmType" bson:"applianceVmType"`
	ManagementNetworkUuid     string            `json:"managementNetworkUuid" bson:"managementNetworkUuid"`
	DefaultRouteL3NetworkUuid string            `json:"defaultRouteL3NetworkUuid" bson:"defaultRouteL3NetworkUuid"`
	Status                    string            `json:"status" bson:"status"`
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
	VMNics                    []VmNicInventory  `json:"vmNics" bson:"vmNics"`         //所有网卡信息
	AllVolumes                []VolumeInventory `json:"allVolumes" bson:"allVolumes"` //所有卷
	Dns                       []Dns             `json:"dns" bson:"dns"`
}

type Dns struct {
	VpcRouterUuid string `json:"vpcRouterUuid" bson:"vpcRouterUuid"`
	Dns           string `json:"dns" bson:"dns"`
	CreateDate    string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate    string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type VpcHaGroupInventory struct {
	UUID        string     `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name        string     `json:"name" bson:"name"`               //资源名称
	Description string     `json:"description" bson:"description"` //资源的详细描述
	CreateDate  string     `json:"createDate" bson:"createDate"`   //创建时间
	LastOpDate  string     `json:"lastOpDate" bson:"lastOpDate"`   //最后一次修改时间
	Monitors    []Monitors `json:"monitors" bson:"monitors"`
	VrRefs      []VrRefs   `json:"VrRefs" bson:"VrRefs"`
	Services    []Services `json:"services" bson:"services"`
	UsedIps     []UsedIps  `json:"usedIps" bson:"usedIps"`
}

type Monitors struct {
	Id              int64  `json:"id" bson:"id"`
	VpcHaRouterUuid string `json:"vpcHaRouterUuid" bson:"vpcHaRouterUuid"`
	MonitorIp       string `json:"monitorIp" bson:"monitorIp"`
	CreateDate      string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type VrRefs struct {
	UUID            string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	VpcHaRouterUuid string `json:"VpcHaRouterUuid" bson:"VpcHaRouterUuid"`
}

type Services struct {
	Id                 int64  `json:"id" bson:"id"`
	VpcHaRouterUuid    string `json:"VpcHaRouterUuid" bson:"VpcHaRouterUuid"`
	NetworkServiceName string `json:"networkServiceName" bson:"networkServiceName"`
	NetworkServiceUuid string `json:"networkServiceUuid" bson:"networkServiceUuid"`
	CreateDate         string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate         string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type UsedIps struct {
	Id              int64  `json:"id" bson:"id"`
	VpcHaRouterUuid string `json:"VpcHaRouterUuid" bson:"VpcHaRouterUuid"`
	VipUuid         string `json:"vipUuid" bson:"vipUuid"`             //VIP UUID
	L3NetworkUuid   string `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //三层网络UUID
	Ip              string `json:"ip" bson:"ip"`
	Netmask         string `json:"netmask" bson:"netmask"`
	CreateDate      string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type VirtualRouterVmInventory struct {
	PublicNetworkUuid         string            `json:"publicNetworkUuid" bson:"publicNetworkUuid"`
	VirtualRouterVips         []string          `json:"virtualRouterVips" bson:"virtualRouterVips"`
	ApplianceVmType           string            `json:"applianceVmType" bson:"applianceVmType"`
	ManagementNetworkUuid     string            `json:"managementNetworkUuid" bson:"managementNetworkUuid"`
	DefaultRouteL3NetworkUuid string            `json:"defaultRouteL3NetworkUuid" bson:"defaultRouteL3NetworkUuid"`
	Status                    string            `json:"status" bson:"status"`
	AgentPort                 int               `json:"agentPort" bson:"agentPort"`
	HaStatus                  string            `json:"haStatus" bson:"haStatus"`
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
	VmCdRoms                  []VmCdrom         `json:"vmCdRoms" bson:"vmCdRoms"`     //所有虚拟光驱
}
