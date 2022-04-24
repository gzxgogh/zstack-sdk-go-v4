package model

//查询物理机
type QueryHostRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryHostResponse struct {
	Error       ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []HostInventory `json:"inventories" bson:"inventories"`
}

//删除物理机
type DeleteHostRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteHostResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新物理机信息
type UpdateHostRequest struct {
	ReqConfig
	UUID       string           `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateHost UpdateHostParams `json:"updateHost" bson:"updateHost"`
	SystemTags []string         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateHostParams struct {
	Name         string `json:"name" bson:"name"`                                     //云主机名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //云主机的详细描述
	ManagementIp string `json:"managementIp,omitempty" bson:"managementIp,omitempty"` //
}

type UpdateHostResponse struct {
	Error     ErrorCode     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory HostInventory `json:"inventory" bson:"inventory"`
}

//更新物理机启用状态
type ChangeHostStateRequest struct {
	ReqConfig
	UUID            string                `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeHostState ChangeHostStateParams `json:"changeHostState" bson:"changeHostState"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeHostStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //enable，disable
}

type ChangeHostStateResponse struct {
	Error     ErrorCode     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory HostInventory `json:"inventory" bson:"inventory"`
}

//重连物理机
type ReconnectHostRequest struct {
	ReqConfig
	UUID          string              `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ReconnectHost ReconnectHostParams `json:"reconnectHost" bson:"reconnectHost"`
	SystemTags    []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ReconnectHostParams struct {
}

type ReconnectHostResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取物理机分配策略
type GetHostAllocatorStrategiesRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetHostAllocatorStrategiesResponse struct {
	Strategies []string  `json:"strategies" bson:"strategies"`
	Error      ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机虚拟化技术类型
type GetHypervisorTypesRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetHypervisorTypesResponse struct {
	HypervisorTypes []string  `json:"hypervisorTypes" bson:"hypervisorTypes"`
	Error           ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新KVM主机信息
type UpdateKVMHostRequest struct {
	ReqConfig
	UUID          string              `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateKVMHost UpdateKVMHostParams `json:"updateKVMHost" bson:"updateKVMHost"`
	SystemTags    []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateKVMHostParams struct {
	Username     string `json:"username" bson:"username"`                           //用户名
	Password     string `json:"password" bson:"password"`                           //密码
	SshPort      string `json:"sshPort " bson:"sshPort "`                           //ssh端口号
	Name         string `json:"name" bson:"name"`                                   //云主机名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"` //云主机的详细描述
	ManagementIp string `json:"managementIp " bson:"managementIp "`                 //管理节点IP
}

type UpdateKVMHostResponse struct {
	Error     ErrorCode     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory HostInventory `json:"inventory" bson:"inventory"`
}

//添加KVM主机
type AddKVMHostRequest struct {
	ReqConfig
	Params     AddKVMHostParams `json:"updateKVMHost" bson:"updateKVMHost"`
	SystemTags []string         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddKVMHostParams struct {
	Username     string `json:"username" bson:"username"`                             //用户名
	Password     string `json:"password" bson:"password"`                             //密码
	SshPort      string `json:"sshPort,omitempty" bson:"sshPort,omitempty"`           //ssh端口号
	Name         string `json:"name" bson:"name"`                                     //云主机名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //云主机的详细描述
	ManagementIp string `json:"managementIp " bson:"managementIp "`                   //管理节点IP
	ClusterUuid  string `json:"clusterUuid" bson:"clusterUuid"`                       //集群UUID
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID
}

type AddKVMHostResponse struct {
	Error     ErrorCode     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory HostInventory `json:"inventory" bson:"inventory"`
}

//KVM运行Shell命令
type KvmRunShellRequest struct {
	ReqConfig
	KvmRunShell KvmRunShellParams `json:"kvmRunShell" bson:"kvmRunShell"`
	SystemTags  []string          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type KvmRunShellParams struct {
	HostUuids []string `json:"hostUuids" bson:"hostUuids"` //目标机器UUID
	Script    string   `json:"script" bson:"script"`       //脚本
}

type KvmRunShellResponse struct {
	Inventory KvmRunShellInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type KvmRunShellInventory struct {
	ReturnCode int       `json:"returnCode" bson:"returnCode"`
	Stdout     string    `json:"stdout" bson:"stdout"`
	Stderr     string    `json:"stderr" bson:"stderr"`
	Error      ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//通过文件导入方式添加物理机
type AddKVMHostFromConfigFileRequest struct {
	ReqConfig
	Params     AddKVMHostFromConfigFileParams `json:"params" bson:"params"`
	SystemTags []string                       `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                       `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddKVMHostFromConfigFileParams struct {
	HostInfo     string `json:"hostInfo" bson:"hostInfo"` //经过base64编码的物理机信息
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`
}

type AddKVMHostFromConfigFileResponse struct {
	Result AddKVMHostFromConfigFileResult `json:"result" bson:"result"`
	Error  ErrorCode                      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type AddKVMHostFromConfigFileResult struct {
	Ip      string    `json:"ip" bson:"ip"`
	Success bool      `json:"success" bson:"success"`                 //是否成功
	Error   ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加物理机文件语法检查
type CheckKVMHostConfigFileRequest struct {
	ReqConfig
	Params     CheckKVMHostConfigFileParams `json:"params" bson:"params"`
	SystemTags []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CheckKVMHostConfigFileParams struct {
	HostInfo string `json:"hostInfo" bson:"hostInfo"` //经过base64编码的物理机信息
}

type CheckKVMHostConfigFileResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取物理机物理网络信息
type GetHostNetworkFactsRequest struct {
	ReqConfig
	HostUuid   string   `json:"hostUuid" bson:"hostUuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetHostNetworkFactsResponse struct {
	Success  bool       `json:"success" bson:"success"` //是否成功
	Bondings []Bondings `json:"bondings" bson:"bondings"`
	Nics     []Slaves   `json:"nics" bson:"nics"`
	Error    ErrorCode  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询物理机Bond信息
type QueryHostNetworkBondingRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryHostNetworkBondingResponse struct {
	Inventories []Bondings `json:"inventories" bson:"inventories"`
	Error       ErrorCode  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询物理机网卡信息
type QueryHostNetworkInterfaceRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryHostNetworkInterfaceResponse struct {
	Inventories []Slaves  `json:"inventories" bson:"inventories"`
	Error       ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询物理机的NUMA拓扑信息
type GetHostNUMATopologyRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetHostNUMATopologyResponse struct {
	Uuid     string    `json:"uuid" bson:"uuid"`
	Name     string    `json:"name" bson:"name"` //云主机名称
	Topology Topology  `json:"topology" bson:"topology"`
	Error    ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//根据云主机需求获取物理机资源的分配信息
type GetHostResourceAllocationRequest struct {
	ReqConfig
	Uuid       string                          `json:"uuid" bson:"uuid"`
	Params     GetHostResourceAllocationParams `json:"params" bson:"params"`
	SystemTags []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetHostResourceAllocationParams struct {
	Strategy string `json:"strategy" bson:"strategy"` //资源分配策略
	Scene    string `json:"scene" bson:"scene"`       //场景类型
	VCpu     string `json:"vcpu" bson:"vcpu"`         //云主机需要分配的vCPU数量
}

type GetHostResourceAllocationResponse struct {
	Uuid    string              `json:"uuid" bson:"uuid"`
	Name    string              `json:"name" bson:"name"` //云主机名称
	VCPUPin []map[string]string `json:"vCPUPin" bson:"vCPUPin"`
	Error   ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type HostInventory struct {
	ZoneUuid                string `json:"zoneUuid" bson:"zoneUuid"` //区域UUID
	Uuid                    string `json:"uuid" bson:"uuid"`
	Name                    string `json:"name" bson:"name"`                                   //云主机名称
	ClusterUuid             string `json:"clusterUuid" bson:"clusterUuid"`                     //集群UUID
	Description             string `json:"description,omitempty" bson:"description,omitempty"` //云主机的详细描述
	ManagementIp            string `json:"managementIp" bson:"managementIp"`                   //
	HypervisorType          string `json:"hypervisorType" bson:"hypervisorType"`
	State                   string `json:"state" bson:"state"`   //物理机状态，包括Enabled,Disabled,PreMaintenance,Maintenance
	Status                  string `json:"status" bson:"status"` //Connecting,Connected，Disconnected
	TotalCpuCapacity        int64  `json:"totalCpuCapacity" bson:"totalCpuCapacity"`
	AvailableCpuCapacity    int64  `json:"availableCpuCapacity" bson:"availableCpuCapacity"`
	CpuSockets              int    `json:"cpuSockets" bson:"cpuSockets"`
	TotalMemoryCapacity     int64  `json:"totalMemoryCapacity" bson:"totalMemoryCapacity"`
	AvailableMemoryCapacity int64  `json:"availableMemoryCapacity" bson:"availableMemoryCapacity"`
	CpuNum                  int    `json:"cpuNum" bson:"cpuNum"`
	CreateDate              string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate              string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type Topology struct {
	Cpus     []string `json:"cpus" bson:"cpus"`         //NUMA node包含的CPUs id列表
	Size     int64    `json:"size" bson:"size"`         //NUMA node总内存大小(单位为B)
	Distance []string `json:"distance" bson:"distance"` //该NUMA node与其他node的距离
	VMsUuid  []string `json:"VMsUuid" bson:"VMsUuid"`   //该NUMA node关联VM的Uuid
	Free     int64    `json:"free" bson:"free"`         //NUMA node可用内存大小(单位为B)
}
