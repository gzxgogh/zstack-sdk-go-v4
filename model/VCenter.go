package model

//添加vCenter资源
type AddVCenterRequest struct {
	ReqConfig
	FlowMeterUuid string          `json:"flowMeterUuid" bson:"flowMeterUuid"` //流量监控资源的uuid
	Params        CreateVipParams `json:"params" bson:"params"`
	SystemTags    []string        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddVCenterParams struct {
	Username     string `json:"username" bson:"username"`                             //登录 vCenter 用户名
	Password     string `json:"password" bson:"password"`                             //登录 vCenter 用户密码
	ZoneUuid     string `json:"zoneUuid" bson:"zoneUuid"`                             //vCenter 将要添加至的目标区域UUID
	Name         string `json:"name,omitempty" bson:"name,omitempty"`                 //vCenter 的名称
	Https        bool   `json:"https,omitempty" bson:"https,omitempty"`               //是否启用 HTTPS 登录（默认开启）
	DomainName   string `json:"domainName" bson:"domainName"`                         //vCenter 域名
	Port         int    `json:"port,omitempty" bson:"port,omitempty"`                 //vCenter端口号
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type AddVCenterResponse struct {
	Inventory VCenterInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询vCenter资源
type QueryVCenterRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVCenterResponse struct {
	Inventories []VCenterInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除vCenter资源
type DeleteVCenterRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteVCenterResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新vCenter资源
type UpdateVCenterRequest struct {
	ReqConfig
	UUID          string              `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateVCenter UpdateVCenterParams `json:"updateVCenter" bson:"updateVCenter"`
	SystemTags    []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateVCenterParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Username    string `json:"username,omitempty" bson:"username,omitempty"`       //登录 vCenter 用户名
	Password    string `json:"password,omitempty" bson:"password,omitempty"`       //登录 vCenter 用户密码
	DomainName  string `json:"domainName,omitempty" bson:"domainName,omitempty"`   //vCenter 域名
	Port        int    `json:"port,omitempty" bson:"port,omitempty"`               //vCenter端口号
}

type UpdateVCenterResponse struct {
	Inventory VCenterInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//同步vCenter资源
type SyncVCenterRequest struct {
	ReqConfig
	UUID        string            `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	SyncVCenter SyncVCenterParams `json:"updateVCenter" bson:"updateVCenter"`
	SystemTags  []string          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type SyncVCenterParams struct {
}

type SyncVCenterResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询vCenter集群
type QueryVCenterClusterRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVCenterClusterResponse struct {
	Inventories []VCenterClusterInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询vCenter主存储
type QueryVCenterPrimaryStorageRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVCenterPrimaryStorageResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询vCenter镜像服务器
type QueryVCenterBackupStorageRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVCenterBackupStorageResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询vCenter资源池
type QueryVCenterResourcePoolRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVCenterResourcePoolResponse struct {
	Inventories []VCenterResourcePool `json:"inventories" bson:"inventories"`
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type VCenterInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`                       //资源的UUID，唯一标示该资源
	Name        string `json:"name" bson:"name"`                       //资源名称
	Description string `json:"description" bson:"description"`         //资源的详细描述
	DomainName  string `json:"domainName" bson:"domainName"`           //vCenter 域名
	Port        int    `json:"port,omitempty" bson:"port,omitempty"`   //vCenter端口号
	Username    string `json:"username" bson:"username"`               //登录 vCenter 用户名
	ZoneUuid    string `json:"zoneUuid" bson:"zoneUuid"`               //vCenter 将要添加至的目标区域UUID
	Https       bool   `json:"https,omitempty" bson:"https,omitempty"` //是否启用 HTTPS 登录（默认开启）
	State       string `json:"state" bson:"state"`
	Status      string `json:"status" bson:"status"`
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type VCenterClusterInventory struct {
	UUID           string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name           string `json:"name" bson:"name"`               //资源名称
	Description    string `json:"description" bson:"description"` //资源的详细描述
	VCenterUuid    string `json:"vCenterUuid" bson:"vCenterUuid"`
	DataCenterUuid string `json:"dataCenterUuid" bson:"dataCenterUuid"`
	Morval         string `json:"morval" bson:"morval"`
	State          string `json:"state" bson:"state"`
	HypervisorType string `json:"hypervisorType" bson:"hypervisorType"` //虚拟化类型
	ZoneUuid       string `json:"zoneUuid" bson:"zoneUuid"`             //vCenter 将要添加至的目标区域UUID
	Type           string `json:"type" bson:"type"`
	CreateDate     string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type VCenterResourcePool struct {
	UUID                string       `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name                string       `json:"name" bson:"name"` //资源名称
	VCenterClusterUuid  string       `json:"vCenterClusterUuid" bson:"vCenterClusterUuid"`
	Morval              string       `json:"morval" bson:"morval"`
	ParentUuid          string       `json:"parentUuid" bson:"parentUuid"`
	CPULimit            int64        `json:"CPULimit" bson:"CPULimit"`
	CPUOverheadLimit    int64        `json:"CPUOverheadLimit" bson:"CPUOverheadLimit"`
	CPUReservation      int64        `json:"CPUReservation" bson:"CPUReservation"`
	CPUShares           int64        `json:"CPUShares" bson:"CPUShares"`
	CPULevel            string       `json:"CPULevel" bson:"CPULevel"`
	MemoryLimit         int64        `json:"memoryLimit" bson:"memoryLimit"`
	MemoryOverheadLimit int64        `json:"memoryOverheadLimit" bson:"memoryOverheadLimit"`
	MemoryReservation   int64        `json:"memoryReservation" bson:"memoryReservation"`
	MemoryShares        int64        `json:"memoryShares" bson:"memoryShares"`
	MemoryLevel         string       `json:"memoryLevel" bson:"memoryLevel"`
	CreateDate          string       `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate          string       `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	SubResources        SubResources `json:"subResources" bson:"subResources"`
}

type SubResources struct {
	UUID                    string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	VCenterResourcePoolUuid string `json:"vCenterResourcePoolUuid" bson:"vCenterResourcePoolUuid"`
	ResourceUuid            string `json:"resourceUuid" bson:"resourceUuid"`
	ResourceType            string `json:"resourceType" bson:"resourceType"`
	CreateDate              string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate              string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}
