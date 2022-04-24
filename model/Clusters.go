package model

//创建一个集群
type CreateClusterRequest struct {
	ReqConfig
	Params     CreateClusterParams `json:"params" bson:"params"`
	SystemTags []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateClusterParams struct {
	ZoneUuid       string   `json:"zoneUuid" bson:"zoneUuid"`                             //区域UUID
	Name           string   `json:"name" bson:"name"`                                     //资源名称
	Description    string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	HypervisorType string   `json:"hypervisorType" bson:"hypervisorType"`                 //虚拟机管理程序类型,KVM Simulator
	Type           string   `json:"type,omitempty" bson:"type,omitempty"`                 //保留域, 请不要使用,zstack
	ResourceUuid   string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //用户指定的资源uuid
	Architecture   string   `json:"architecture,omitempty" bson:"architecture,omitempty"` //x86_64,aarch64,mips64el
	TagUuids       []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type CreateClusterResponse struct {
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory ClusterInventory `json:"inventory" bson:"inventory"`
}

//删除一个集群
type DeleteClusterRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteClusterResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询集群
type QueryClusterRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryClusterResponse struct {
	Error       ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []ClusterInventory `json:"inventories" bson:"inventories"`
}

//更新集群
type UpdateClusterRequest struct {
	ReqConfig
	UUID          string              `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateCluster UpdateClusterParams `json:"updateCluster" bson:"updateCluster"`
	SystemTags    []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateClusterParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateClusterResponse struct {
	Error     ErrorCode     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory ZoneInventory `json:"inventory" bson:"inventory"`
}

//改变一个集群的可用状态
type ChangeClusterStateRequest struct {
	ReqConfig
	UUID               string                   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeClusterState ChangeClusterStateParams `json:"changeClusterState" bson:"changeClusterState"`
	SystemTags         []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeClusterStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //enable，disable
}

type ChangeClusterStateResponse struct {
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory ClusterInventory `json:"inventory" bson:"inventory"`
}

//升级集群内物理机的操作系统
type UpdateClusterOSRequest struct {
	ReqConfig
	UUID            string                `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateClusterOS UpdateClusterOSParams `json:"updateClusterOS" bson:"updateClusterOS"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateClusterOSParams struct {
	ExcludePackages []string `json:"excludePackages,omitempty" bson:"excludePackages,omitempty"` //不升级的包列表
	UpdatePackages  []string `json:"updatePackages,omitempty" bson:"updatePackages,omitempty"`   //要升级的包列表
	ReleaseVersion  string   `json:"releaseVersion,omitempty" bson:"releaseVersion,omitempty"`   //升级所用源的版
	ResourceUuid    string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`       //用户指定的资源uuid
	TagUuids        []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`               //标签UUID列表
}

type UpdateClusterOSResponse struct {
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory LongJobInventory `json:"inventory" bson:"inventory"`
}

type ClusterInventory struct {
	Uuid           string `json:"uuid" bson:"uuid"`
	Name           string `json:"name" bson:"name"`                                   //云主机名称
	Description    string `json:"description,omitempty" bson:"description,omitempty"` //云主机的详细描述
	State          string `json:"state" bson:"state"`                                 //集群状态
	HypervisorType string `json:"hypervisorType" bson:"hypervisorType"`               //虚拟机管理程序类型
	Type           string `json:"type,omitempty" bson:"type,omitempty"`               //云主机类型 保留字段，无需指定。UserVm/ApplianceVm
	CreateDate     string `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
	ZoneUuid       string `json:"zoneUuid" bson:"zoneUuid"`                           //区域UUID
}

type LongJobInventory struct {
	Uuid               string `json:"uuid" bson:"uuid"`
	Name               string `json:"name" bson:"name"`                                   //云主机名称
	Description        string `json:"description,omitempty" bson:"description,omitempty"` //云主机的详细描述
	ApiId              string `json:"apiId" bson:"apiId"`                                 //用于关联TaskProgress的APIID
	JobName            string `json:"jobName" bson:"jobName"`                             //任务名称
	JobData            string `json:"jobData" bson:"jobData"`                             //任务数据
	JobResult          string `json:"jobResult" bson:"jobResult"`                         //	任务结果
	TargetResourceUuid string `json:"targetResourceUuid" bson:"targetResourceUuid"`       //目标资源UUID
	ManagementNodeUuid string `json:"managementNodeUuid" bson:"managementNodeUuid"`       //管理节点UUID
	CreateDate         string `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate         string `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
	ExecuteTime        int64  `json:"executeTime" bson:"executeTime"`                     //最后一次修改时间
	State              string `json:"state" bson:"state"`                                 //Waiting,Suspended,Running,Succeeded,Canceling,Canceled,Failed
}
