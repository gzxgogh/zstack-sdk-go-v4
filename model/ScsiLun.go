package model

//查询SCSI Lun
type QueryScsiLunRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryScsiLunResponse struct {
	Inventories []ScsiLunInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//将LUN设备从物理机卸载
type DetachScsiLunFromHostRequest struct {
	ReqConfig
	UUID                  string                      `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	DetachScsiLunFromHost DetachScsiLunFromHostParams `json:"detachScsiLunFromHost" bson:"detachScsiLunFromHost"`
	SystemTags            []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags              []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachScsiLunFromHostParams struct {
	HostUuid string `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`
}

type DetachScsiLunFromHostResponse struct {
	Inventories []ScsiLunInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加iSCSI服务器
type AddIscsiServerRequest struct {
	ReqConfig
	Params     AddIscsiServerParams `json:"params" bson:"params"`                             //资源的UUID，唯一标示该资源
	SystemTags []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddIscsiServerParams struct {
	Name             string `json:"name,omitempty" bson:"name,omitempty"` //资源名称
	Ip               string `json:"ip" bson:"ip"`
	Port             string `json:"port,omitempty" bson:"port,omitempty"`
	ChapUserName     string `json:"chapUserName,omitempty" bson:"chapUserName,omitempty"`
	ChapUserPassword string `json:"chapUserPassword,omitempty" bson:"chapUserPassword,omitempty"`
	ResourceUuid     string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type AddIscsiServerResponse struct {
	Inventory IscsiServerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除iSCSI服务器
type DeleteIscsiServerRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteIscsiServerResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询iSCSI服务器
type QueryIscsiServerRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryIscsiServerResponse struct {
	Inventories []IscsiServerInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//刷新iSCSI服务器
type RefreshIscsiServerRequest struct {
	ReqConfig
	Uuid       string                   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Params     RefreshIscsiServerParams `json:"params" bson:"params"`
	SystemTags []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RefreshIscsiServerParams struct {
}

type RefreshIscsiServerResponse struct {
	Inventory IscsiServerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新iSCSI服务器配置
type UpdateIscsiServerRequest struct {
	ReqConfig
	Uuid              string                  `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateIscsiServer UpdateIscsiServerParams `json:"updateIscsiServer" bson:"updateIscsiServer"`
	SystemTags        []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateIscsiServerParams struct {
	Name             string `json:"name,omitempty" bson:"name,omitempty"`                         //资源名称
	ChapUserName     string `json:"chapUserName,omitempty" bson:"chapUserName,omitempty"`         //CHAP用户名
	ChapUserPassword string `json:"chapUserPassword,omitempty" bson:"chapUserPassword,omitempty"` //CHAP密码
	State            string `json:"state,omitempty" bson:"state,omitempty"`                       //启用状态:Enabled,Disabled
	ResourceUuid     string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`         //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type UpdateIscsiServerResponse struct {
	Inventory IscsiServerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//将iSCSI服务器加载到集群
type AttachIscsiServerToClusterRequest struct {
	ReqConfig
	Uuid        string                           `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ClusterUuid string                           `json:"clusterUuid" bson:"clusterUuid"`
	Params      AttachIscsiServerToClusterParams `json:"params" bson:"params"`
	SystemTags  []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachIscsiServerToClusterParams struct {
}

type AttachIscsiServerToClusterResponse struct {
	Inventory IscsiServerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//将iSCSI服务器从集群卸载
type DetachIscsiServerFromClusterRequest struct {
	ReqConfig
	Uuid        string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ClusterUuid string   `json:"clusterUuid" bson:"clusterUuid"`
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachIscsiServerFromClusterResponse struct {
	Inventory IscsiServerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询iSCSI磁盘
type QueryIscsiLunRequest struct {
	ReqConfig
	Uuid        string   `json:"uuid,omitempty" bson:"uuid,omitempty"` //资源的UUID，唯一标示该资源
	ClusterUuid string   `json:"clusterUuid" bson:"clusterUuid"`
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryIscsiLunResponse struct {
	Inventories []IscsiLuns `json:"inventories" bson:"inventories"`
	Error       ErrorCode   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询FC SAN存储
type QueryFiberChannelStorageRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryFiberChannelStorageResponse struct {
	Inventories []FiberInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//刷新FC SAN存储
type RefreshFiberChannelStorageRequest struct {
	ReqConfig
	Params     RefreshFiberChannelStorageParams `json:"params" bson:"params"`
	SystemTags []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RefreshFiberChannelStorageParams struct {
	Uuid     string `json:"uuid,omitempty" bson:"uuid,omitempty"` //资源的UUID，唯一标示该资源
	ZoneUuid string `json:"zoneUuid" bson:"zoneUuid"`
}

type RefreshFiberChannelStorageResponse struct {
	Inventories []FiberInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新SCSI LUN配置
type UpdateScsiLunRequest struct {
	ReqConfig
	Uuid          string              `json:"uuid,omitempty" bson:"uuid,omitempty"` //资源的UUID，唯一标示该资源
	UpdateScsiLun UpdateScsiLunParams `json:"updateScsiLun" bson:"updateScsiLun"`
	SystemTags    []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateScsiLunParams struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`   //资源名称
	State string `json:"state,omitempty" bson:"state,omitempty"` //启用状态
}

type UpdateScsiLunResponse struct {
	Inventory ScsiLunInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//将SCSI LUN加载到云主机
type AttachScsiLunToVmInstanceRequest struct {
	ReqConfig
	Uuid           string                          `json:"uuid,omitempty" bson:"uuid,omitempty"` //资源的UUID，唯一标示该资源
	VmInstanceUuid string                          `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	Params         AttachScsiLunToVmInstanceParams `json:"params" bson:"params"`
	SystemTags     []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachScsiLunToVmInstanceParams struct {
	DisableMultiPathAttach string `json:"disableMultiPathAttach " bson:"disableMultiPathAttach "`
}

type AttachScsiLunToVmInstanceResponse struct {
	Inventory ScsiLunInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//将SCSI LUN从云主机卸载
type DetachScsiLunFromVmInstanceRequest struct {
	ReqConfig
	Uuid           string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachScsiLunFromVmInstanceResponse struct {
	Inventory ScsiLunInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//检查SCSI Lun与集群连接关系
type CheckScsiLunClusterStatusRequest struct {
	ReqConfig
	Uuid                      string                          `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	CusterUuid                string                          `json:"clusterUuid" bson:"clusterUuid"`
	CheckScsiLunClusterStatus CheckScsiLunClusterStatusParams `json:"checkScsiLunClusterStatus" bson:"checkScsiLunClusterStatus"`
	SystemTags                []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                  []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CheckScsiLunClusterStatusParams struct {
}

type CheckScsiLunClusterStatusResponse struct {
	Inventory ScsiLunInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取可加载的SCSI Lun
type GetScsiLunCandidatesForAttachingVmRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetScsiLunCandidatesForAttachingVmResponse struct {
	Success     bool               `json:"success" bson:"success"`
	Inventories []ScsiLunInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type ScsiLunInventory struct {
	FiberChannelStorageUuid string                `json:"fiberChannelStorageUuid" bson:"fiberChannelStorageUuid"`
	UUID                    string                `json:"uuid" bson:"uuid"`             //资源的UUID，唯一标示该资源
	Name                    string                `json:"name" bson:"name"`             //资源名称
	Wwid                    string                `json:"wwid" bson:"wwid"`             //磁盘全局唯一表示
	Vendor                  string                `json:"vendor" bson:"vendor"`         //磁盘供应商
	Model                   string                `json:"model" bson:"model"`           //磁盘型号
	Wwn                     string                `json:"wwn" bson:"wwn"`               //磁盘WWN
	Serial                  string                `json:"serial" bson:"serial"`         //磁盘序列号
	Type                    string                `json:"type" bson:"type"`             //磁盘类型
	Path                    string                `json:"path" bson:"path"`             //磁盘路径
	State                   string                `json:"state" bson:"state"`           //磁盘启用状态
	Size                    int64                 `json:"size" bson:"size"`             //磁盘大小
	CreateDate              string                `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate              string                `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	ScsiLunHostRefs         ScsiLunHostRefs       `json:"scsiLunHostRefs" bson:"scsiLunHostRefs"`
	ScsiLunVmInstanceRefs   ScsiLunVmInstanceRefs `json:"ScsiLunVmInstanceRefs" bson:"ScsiLunVmInstanceRefs"`
}

type ScsiLunHostRefs struct {
	ScsiLunUuid string `json:"scsiLunUuid" bson:"scsiLunUuid"`
	HostUuid    string `json:"hostUuid" bson:"hostUuid"`
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type ScsiLunVmInstanceRefs struct {
	ScsiLunUuid    string `json:"scsiLunUuid" bson:"scsiLunUuid"`
	VmInstanceUuid string `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	CreateDate     string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type IscsiServerInventory struct {
	UUID             string           `json:"uuid" bson:"uuid"`                         //资源的UUID，唯一标示该资源
	Name             string           `json:"name" bson:"name"`                         //资源名称
	Ip               string           `json:"ip" bson:"ip"`                             //IP地址
	Port             string           `json:"port" bson:"port"`                         //端口
	ChapUserName     string           `json:"chapUserName" bson:"chapUserName"`         //CHAP用户名
	ChapUserPassword string           `json:"chapUserPassword" bson:"chapUserPassword"` //CHAP密码
	State            string           `json:"state" bson:"state"`                       //启用状态
	CreateDate       string           `json:"createDate" bson:"createDate"`             //创建时间
	LastOpDate       string           `json:"lastOpDate" bson:"lastOpDate"`             //最后一次修改时间
	IscsiTargets     IscsiTargets     `json:"iscsiTargets" bson:"iscsiTargets"`
	IscsiClusterRefs IscsiClusterRefs `json:"iscsiClusterRefs" bson:"iscsiClusterRefs"`
}

type IscsiTargets struct {
	IscsiServerUuid string    `json:"iscsiServerUuid" bson:"iscsiServerUuid"` //iSCSI服务器UUID
	UUID            string    `json:"uuid" bson:"uuid"`                       //资源的UUID，唯一标示该资源
	Iqn             string    `json:"iqn" bson:"iqn"`                         //iSCSI IQN
	CreateDate      string    `json:"createDate" bson:"createDate"`           //创建时间
	LastOpDate      string    `json:"lastOpDate" bson:"lastOpDate"`           //最后一次修改时间
	IscsiLuns       IscsiLuns `json:"iscsiLuns" bson:"iscsiLuns"`
}

type IscsiLuns struct {
	UUID            string `json:"uuid" bson:"uuid"`                       //资源的UUID，唯一标示该资源
	IscsiTargetUuid string `json:"iscsiTargetUuid" bson:"iscsiTargetUuid"` //iSCSI目标UUID
	Wwid            string `json:"wwid" bson:"wwid"`                       //唯一识别ID
	Vendor          string `json:"vendor" bson:"vendor"`                   //磁盘供应商
	Model           string `json:"model" bson:"model"`                     //磁盘型号
	Wwn             string `json:"wwn" bson:"wwn"`                         //磁盘WWN
	Serial          string `json:"serial" bson:"serial"`                   //磁盘序列号
	Hctl            string `json:"hctl" bson:"hctl"`                       //SCSI设备HCTL
	Type            string `json:"type" bson:"type"`                       //磁盘类型
	Path            string `json:"path" bson:"path"`                       //磁盘路径
	Size            int64  `json:"size" bson:"size"`                       //磁盘大小
	CreateDate      string `json:"createDate" bson:"createDate"`           //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"`           //最后一次修改时间
}

type IscsiClusterRefs struct {
	IscsiServerUuid string `json:"iscsiServerUuid" bson:"iscsiServerUuid"`
	ClusterUuid     string `json:"clusterUuid" bson:"clusterUuid"` //集群UUID
	CreateDate      string `json:"createDate" bson:"createDate"`   //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"`   //最后一次修改时间
}

type FiberInventory struct {
	UUID             string             `json:"uuid" bson:"uuid"`             //资源的UUID，唯一标示该资源
	Name             string             `json:"name" bson:"name"`             //资源名称
	Wwnn             string             `json:"wwnn" bson:"wwnn"`             //iSCSI目标UUID
	State            string             `json:"state" bson:"state"`           //启用状态
	CreateDate       string             `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate       string             `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	FiberChannelLuns []ScsiLunInventory `json:"fiberChannelLuns" bson:"fiberChannelLuns"`
}
