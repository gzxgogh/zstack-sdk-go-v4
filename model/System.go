package model

//查询管理节点
type QueryManagementNodeRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryManagementNodeResponse struct {
	Inventories []ManagementNodeInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type ManagementNodeInventory struct {
	UUID      string `json:"uuid" bson:"uuid"`           //资源的UUID，唯一标示该资源
	Hostname  string `json:"hostname" bson:"hostname"`   //宿主机名称
	JoinDate  string `json:"joinDate" bson:"joinDate"`   //加入时间
	HeartBeat string `json:"heartBeat" bson:"heartBeat"` //心跳时间
}

//获取当前版本
type GetVersionRequest struct {
	ReqConfig
	GetVersion GetVersionParams `json:"getVersion" bson:"getVersion"`
	SystemTags []string         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetVersionParams struct {
}

type GetVersionResponse struct {
	Version string    `json:"version" bson:"version"`                 //管理节点当前版本
	Success bool      `json:"success" bson:"success"`                 //是否成功
	Error   ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取当前时间
type GetCurrentTimeRequest struct {
	ReqConfig
	GetCurrentTime GetCurrentTimeParams `json:"getCurrentTime" bson:"getCurrentTime"`
	SystemTags     []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetCurrentTimeParams struct {
}

type GetCurrentTimeResponse struct {
	CurrentTime map[string]interface{} `json:"currentTime" bson:"currentTime"`         //管理节点当前时间
	Success     bool                   `json:"success" bson:"success"`                 //是否成功
	Error       ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取管理节点当前的时区信息
type GetPlatformTimeZoneRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetPlatformTimeZoneResponse struct {
	Timezone string    `json:"timezone" bson:"timezone"`               //时区名称
	Offset   string    `json:"offset" bson:"offset"`                   //与GMT时间的时差
	Success  bool      `json:"success" bson:"success"`                 //是否成功
	Error    ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//检查管理节点是否能正常开始工作
type IsReadyToGoRequest struct {
	ReqConfig
	ManagementNodeId string   `json:"managementNodeId,omitempty" bson:"managementNodeId,omitempty"`
	SystemTags       []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type IsReadyToGoResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取管理节点操作系统信息
type GetManagementNodeOSRequest struct {
	ReqConfig
	GetManagementNodeOS GetManagementNodeOSParams `json:"getManagementNodeOS" bson:"getManagementNodeOS"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetManagementNodeOSParams struct {
}

type GetManagementNodeOSResponse struct {
	Name    string    `json:"name" bson:"name"`       //资源名称
	Version string    `json:"version" bson:"version"` //版本
	Success bool      `json:"success" bson:"success"`
	Error   ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取管理节点CPU架构类型
type GetManagementNodeArchRequest struct {
	ReqConfig
	GetManagementNodeArch GetManagementNodeArchParams `json:"getManagementNodeArch" bson:"getManagementNodeArch"`
	SystemTags            []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags              []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetManagementNodeArchParams struct {
}

type GetManagementNodeArchResponse struct {
	Architecture string    `json:"architecture" bson:"architecture"` //资源名称
	Success      bool      `json:"success" bson:"success"`
	Error        ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建系统标签
type CreateSystemTagRequest struct {
	ReqConfig
	Params     CreateSystemTagParams `json:"params" bson:"params"`
	SystemTags []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSystemTagParams struct {
	ResourceUuid string `json:"resourceUuid" bson:"resourceUuid"` //当创建一个标签时, 用户必须指定标签所关联的资源类型
	ResourceType string `json:"resourceType" bson:"resourceType"` //用户指定的资源UUID，若指定，系统不会为该资源随机分配UUID
	Tag          string `json:"tag" bson:"tag"`                   //标签字符串
}

type CreateSystemTagResponse struct {
	Inventory SystemTagInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询系统标签
type QuerySystemTagRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySystemTagResponse struct {
	Inventories []SystemTagInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新系统标签
type UpdateSystemTagRequest struct {
	ReqConfig
	UUID            string                `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateSystemTag UpdateSystemTagParams `json:"updateSecurityGroup" bson:"updateSecurityGroup"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateSystemTagParams struct {
	Tag string `json:"tag" bson:"tag"`
}

type UpdateSystemTagResponse struct {
	Inventory SystemTagInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建用户标签
type CreateUserTagRequest struct {
	ReqConfig
	Params     CreateUserTagParams `json:"params" bson:"params"`
	SystemTags []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateUserTagParams struct {
	ResourceUuid string `json:"resourceUuid" bson:"resourceUuid"` //当创建一个标签时, 用户必须指定标签所关联的资源类型
	ResourceType string `json:"resourceType" bson:"resourceType"` //用户指定的资源UUID，若指定，系统不会为该资源随机分配UUID
	Tag          string `json:"tag" bson:"tag"`                   //标签字符串
}

type CreateUserTagResponse struct {
	Inventory UserTagInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询用户标签
type QueryUserTagRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryUserTagResponse struct {
	Inventories []UserTagInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除标签
type DeleteTagRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteTagResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建资源标签
type CreateTagRequest struct {
	ReqConfig
	Params     CreateTagParams `json:"params" bson:"params"`
	SystemTags []string        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateTagParams struct {
	Name         string `json:"name" bson:"name"`                                   //资源名称
	Value        string `json:"value" bson:"value"`                                 //标签的值
	Description  string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Color        string `json:"color,omitempty" bson:"color,omitempty"`             //标签的颜色
	Type         string `json:"type,omitempty" bson:"type,omitempty"`               //标签的类型:simple,withToken
	ResourceUuid string `json:"resourceUuid" bson:"resourceUuid"`                   //当创建一个标签时, 用户必须指定标签所关联的资源类型
}

type CreateTagResponse struct {
	Inventory TagInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询资源标签
type QueryTagRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryTagResponse struct {
	Inventories []TagInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新资源标签
type UpdateTagRequest struct {
	ReqConfig
	UUID       string          `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateTag  UpdateTagParams `json:"updateSecurityGroup" bson:"updateSecurityGroup"`
	SystemTags []string        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateTagParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Value       string `json:"value" bson:"value"`                                 //标签的值
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Color       string `json:"color,omitempty" bson:"color,omitempty"`             //标签的颜色
}

type UpdateTagResponse struct {
	Inventory TagInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//加载资源标签
type AttachTagToResourcesRequest struct {
	ReqConfig
	TagUuid    string                     `json:"tagUuid" bson:"tagUuid"` //资源的UUID，唯一标示该资源
	Params     AttachTagToResourcesParams `json:"params" bson:"params"`
	SystemTags []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachTagToResourcesParams struct {
	ResourceUuids []string               `json:"resourceUuids" bson:"resourceUuids"` //资源UUID
	Tokens        map[string]interface{} `json:"tokens " bson:"tokens "`             //通过标签存放的键和值
}

type AttachTagToResourcesResponse struct {
	Success bool                       `json:"success" bson:"success"`
	Results AttachTagToResourcesResult `json:"results" bson:"results"`
	Error   ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//卸载资源标签
type DetachTagFromResourcesRequest struct {
	ReqConfig
	TagUuid       string   `json:"tagUuid" bson:"tagUuid"`                           //资源的UUID，唯一标示该资源
	ResourceUuids []string `json:"resourceUuids" bson:"resourceUuids"`               //资源UUID
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachTagFromResourcesResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type SystemTagInventory struct {
	Inherent     bool   `json:"inherent" bson:"inherent"`         //内部系统标签
	UUID         string `json:"uuid" bson:"uuid"`                 //资源的UUID，唯一标示该资源
	ResourceUuid string `json:"resourceUuid" bson:"resourceUuid"` //当创建一个标签时, 用户必须指定标签所关联的资源类型
	ResourceType string `json:"resourceType" bson:"resourceType"` //用户指定的资源UUID，若指定，系统不会为该资源随机分配UUID
	Tag          string `json:"tag" bson:"tag"`                   //标签字符串
	Type         string `json:"type" bson:"type"`                 //保留域, 请不要使用它
	CreateDate   string `json:"createDate" bson:"createDate"`     //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"`     //最后一次修改时间

}

type UserTagInventory struct {
	UUID         string `json:"uuid" bson:"uuid"`                 //资源的UUID，唯一标示该资源
	ResourceUuid string `json:"resourceUuid" bson:"resourceUuid"` //当创建一个标签时, 用户必须指定标签所关联的资源类型
	ResourceType string `json:"resourceType" bson:"resourceType"` //用户指定的资源UUID，若指定，系统不会为该资源随机分配UUID
	Tag          string `json:"tag" bson:"tag"`                   //标签字符串
	Type         string `json:"type" bson:"type"`                 //保留域, 请不要使用它
	CreateDate   string `json:"createDate" bson:"createDate"`     //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"`     //最后一次修改时间
}

type TagInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Value       string `json:"value" bson:"value"`                                 //标签的值
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Color       string `json:"color,omitempty" bson:"color,omitempty"`             //标签的颜色
	Type        string `json:"type" bson:"type"`                                   //保留域, 请不要使用它
	CreateDate  string `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
}

type AttachTagToResourcesResult struct {
	Success   bool             `json:"success" bson:"success"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory UserTagInventory `json:"inventory" bson:"inventory"`
}

//获取任务进度
type GetTaskProgressRequest struct {
	ReqConfig
	ApiId      string   `json:"apiId,omitempty " bson:"apiId,omitempty"` //资源的UUID，唯一标示该资源
	All        bool     `json:"all,omitempty" bson:"all,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetTaskProgressResponse struct {
	Inventories TaskProgressInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type TaskProgressInventory struct {
	TaskUuid    string                   `json:"taskUuid" bson:"taskUuid"`
	TaskName    string                   `json:"taskName" bson:"taskName"`
	ParentUuid  string                   `json:"parentUuid" bson:"parentUuid"`
	Type        string                   `json:"type" bson:"type"`
	Content     string                   `json:"content" bson:"content"`
	Opaque      []map[string]interface{} `json:"opaque" bson:"opaque"`
	Time        int64                    `json:"time" bson:"time"`
	TotalSteps  int                      `json:"totalSteps" bson:"totalSteps"`
	CurrentStep int                      `json:"currentStep" bson:"currentStep"`
	SubTasks    []map[string]interface{} `json:"subTasks" bson:"subTasks"`
}

//获取CPU和内存容量
type GetCpuMemoryCapacityRequest struct {
	ReqConfig
	ZoneUuids      []string `json:"zoneUuids,omitempty" bson:"zoneUuids,omitempty"`           //区域的UUID
	ClusterUuids   []string `json:"clusterUuids,omitempty" bson:"clusterUuids,omitempty"`     //集群的UUID。用于挂载网络、存储等
	HostUuids      []string `json:"hostUuids,omitempty" bson:"hostUuids,omitempty"`           //物理机的UUID。用于添加、删除host等
	All            bool     `json:"all,omitempty" bson:"all,omitempty"`                       //all默认为false，此时区域UUID，集群UUID,物理机UUID必须有一个不为空，或者全部都填写,all设置为true时，区域UUID、集群UUID、物理机UUID均可不填
	HypervisorType string   `json:"hypervisorType,omitempty" bson:"hypervisorType,omitempty"` //KVM,ESC
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`         //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetCpuMemoryCapacityResponse struct {
	TotalCpu        int64          `json:"totalCpu" bson:"totalCpu"`               //CPU总数
	AvailableCpu    int64          `json:"availableCpu" bson:"availableCpu"`       //可用CPU数量
	TotalMemory     int64          `json:"totalMemory" bson:"totalMemory"`         //内存总量
	AvailableMemory int64          `json:"availableMemory" bson:"availableMemory"` //可用内存
	ManagedCpuNum   int64          `json:"managedCpuNum" bson:"managedCpuNum"`     //受管理的物理CPU数量
	ResourceType    string         `json:"resourceType" bson:"resourceType"`       //所查资源的类型（物理机、集群、区域）
	Success         bool           `json:"success" bson:"success"`
	CapacityData    []CapacityData `json:"capacityData" bson:"capacityData"`
	Error           ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type CapacityData struct {
	ResourceUuid    string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TotalCpu        int64  `json:"totalCpu" bson:"totalCpu"`                             //CPU总数
	AvailableCpu    int64  `json:"availableCpu" bson:"availableCpu"`                     //可用CPU数量
	TotalMemory     int64  `json:"totalMemory" bson:"totalMemory"`                       //内存总量
	AvailableMemory int64  `json:"availableMemory" bson:"availableMemory"`               //可用内存
	ManagedCpuNum   int64  `json:"managedCpuNum" bson:"managedCpuNum"`                   //受管理的物理CPU数量
}

//触发垃圾回收任务
type TriggerGCJobRequest struct {
	ReqConfig
	TriggerGCJob TriggerGCJobParams `json:"triggerGCJob" bson:"triggerGCJob"`
	UUID         string             `json:"uuid " bson:"uuid"`                                //资源的UUID，唯一标示该资源
	SystemTags   []string           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type TriggerGCJobParams struct {
}

type TriggerGCJobResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除垃圾回收任务
type DeleteGCJobRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteGCJobResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询垃圾回收任务
type QueryGCJobRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryGCJobResponse struct {
	Inventories []GCJobInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type GCJobInventory struct {
	UUID               string `json:"uuid " bson:"uuid"` //资源的UUID，唯一标示该资源
	Name               string `json:"name" bson:"name"`  //资源名称
	RunnerClass        string `json:"runnerClass" bson:"runnerClass"`
	Context            string `json:"context" bson:"context"`
	Status             string `json:"status" bson:"status"`
	ManagementNodeUuid string `json:"managementNodeUuid" bson:"managementNodeUuid"`
	Type               string `json:"type" bson:"type"`
	CreateDate         string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate         string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

//获取许可证信息
type GetLicenseInfoRequest struct {
	ReqConfig
	AdditionSession string   `json:"additionSession,omitempty" bson:"additionSession,omitempty"`
	SystemTags      []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetLicenseInfoResponse struct {
	Inventory LicenseInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取许可证历史授权信息
type GetLicenseRecordsRequest struct {
	ReqConfig
	Limit          int      `json:"limit,omitempty" bson:"limit,omitempty"`
	Start          int      `json:"start,omitempty" bson:"start,omitempty"`
	ReplyWithCount bool     `json:"replyWithCount,omitempty" bson:"replyWithCount,omitempty"`
	Count          bool     `json:"count,omitempty" bson:"count,omitempty"`
	SortBy         string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortDirection  string   `json:"sortDirection,omitempty" bson:"sortDirection,omitempty"`
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetLicenseRecordsResponse struct {
	Inventory LicenseInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取许可证容量
type GetLicenseCapabilitiesRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetLicenseCapabilitiesResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除许可证文件
type DeleteLicenseRequest struct {
	ReqConfig
	UUID               string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	ManagementNodeUuid string   `json:"managementNodeUuid" bson:"managementNodeUuid"`     //管理节点的UUID
	SystemTags         []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteLicenseResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//重新加载许可证
type ReloadLicenseRequest struct {
	ReqConfig
	ReloadLicense ReloadLicenseParams `json:"reloadLicense" bson:"reloadLicense"`
	SystemTags    []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ReloadLicenseParams struct {
	ManagementNodeUuids []string `json:"managementNodeUuids" bson:"managementNodeUuids"`
}

type ReloadLicenseResponse struct {
	Inventory LicenseInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新许可证信息
type UpdateLicenseRequest struct {
	ReqConfig
	ManagementNodeUuid string              `json:"managementNodeUuid" bson:"managementNodeUuid"`
	UpdateLicense      UpdateLicenseParams `json:"updateLicense" bson:"updateLicense"`
	SystemTags         []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateLicenseParams struct {
	License         string `json:"license" bson:"license"`
	AdditionSession string `json:"additionSession,omitempty" bson:"additionSession,omitempty"`
}

type UpdateLicenseResponse struct {
	Inventory LicenseInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type LicenseInventory struct {
	LicenseType      string `json:"licenseType" bson:"licenseType"`
	LicenseRequest   string `json:"licenseRequest" bson:"licenseRequest"`
	ExpiredDate      string `json:"expiredDate" bson:"expiredDate"`
	IssuedDate       string `json:"issuedDate" bson:"issuedDate"`
	User             string `json:"user" bson:"user"`
	HostNum          int    `json:"hostNum" bson:"hostNum"`
	CpuNum           int    `json:"cpuNum" bson:"cpuNum"`
	AvailableHostNum int    `json:"availableHostNum" bson:"availableHostNum"`
	AvailableCpuNum  int    `json:"availableCpuNum" bson:"availableCpuNum"`
	Expired          bool   `json:"expired" bson:"expired"`
}

//提交长时任务
type SubmitLongJobRequest struct {
	ReqConfig
	Params     SubmitLongJobParams `json:"params" bson:"params"`
	SystemTags []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type SubmitLongJobParams struct {
	Name               string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description        string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	JobName            string `json:"jobName" bson:"jobName"`                             //任务名称
	JobData            string `json:"jobData,omitempty" bson:"jobData,omitempty"`         //任务数据
	TargetResourceUuid string `json:"targetResourceUuid,omitempty" bson:"targetResourceUuid,omitempty"`
	ResourceUuid       string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type SubmitLongJobResponse struct {
	Inventory LongJobInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新长时任务
type UpdateLongJobRequest struct {
	ReqConfig
	UUID          string              `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateLongJob UpdateLongJobParams `json:"updateLongJob" bson:"updateLongJob"`
	SystemTags    []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateLongJobParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateLongJobResponse struct {
	Inventory LongJobInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除长时任务
type DeleteLongJobRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteLongJobResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询长时任务
type QueryLongJobRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryLongJobResponse struct {
	Inventories []LongJobInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查看系统错误码
type GetElaborationsRequest struct {
	ReqConfig
	Category   string   `json:"category,omitempty" bson:"category,omitempty"`     //错误码目录
	Regex      string   `json:"regex,omitempty" bson:"regex,omitempty"`           //错误码关键字
	Code       string   `json:"code,omitempty" bson:"code,omitempty"`             //错误代码，与category一起使用
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetElaborationsResponse struct {
	Contents []Contents `json:"contents" bson:"contents"`
	Error    ErrorCode  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type Contents struct {
	Category   string  `json:"category,omitempty" bson:"category,omitempty"` //错误码目录
	Regex      string  `json:"regex" bson:"regex"`                           //错误码关键字
	Code       string  `json:"code" bson:"code"`                             //错误代码，与category一起使用
	Message_cn string  `json:"message_cn" bson:"message_cn"`                 //错误码内容中文
	Message_en string  `json:"message_en" bson:"message_en"`                 //错误码内容英文
	Source     string  `json:"source" bson:"source"`                         //错误来源
	Method     string  `json:"method" bson:"method"`                         //匹配方法，distance(字符串比较)或regex(正则)
	Distance   float64 `json:"distance" bson:"distance"`                     //若使用distance匹配，此处为精确度(1为最精确)
}

//重新生成索引
type RefreshSearchIndexesRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RefreshSearchIndexesResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}
