package model

//删除主存储
type DeletePrimaryStorageRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeletePrimaryStorageResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询主存储
type QueryPrimaryStorageRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPrimaryStorageResponse struct {
	Inventories []PrimaryStorageInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//向集群添加主存储
type AttachPrimaryStorageToClusterRequest struct {
	ReqConfig
	ClusterUuid        string   `json:"clusterUuid" bson:"clusterUuid"`                   //集群UUID
	PrimaryStorageUuid string   `json:"primaryStorageUuid" bson:"primaryStorageUuid"`     //主存储UUID
	SystemTags         []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachPrimaryStorageToClusterResponse struct {
	Inventory PrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从集群卸载主存储
type DetachPrimaryStorageFromClusterRequest struct {
	ReqConfig
	ClusterUuid        string   `json:"clusterUuid" bson:"clusterUuid"`                   //集群UUID
	PrimaryStorageUuid string   `json:"primaryStorageUuid" bson:"primaryStorageUuid"`     //主存储UUID
	SystemTags         []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachPrimaryStorageFromClusterResponse struct {
	Inventory PrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//重连主存储
type ReconnectPrimaryStorageRequest struct {
	ReqConfig
	Uuid                    string                        `json:"uuid" bson:"uuid"`                                                           //集群UUID
	ReconnectPrimaryStorage ReconnectPrimaryStorageParams `json:"reconnectPrimaryStorage,omitempty" bson:"reconnectPrimaryStorage,omitempty"` //主存储UUID
	SystemTags              []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"`                           //云主机系统标签
	UserTags                []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ReconnectPrimaryStorageParams struct {
}

type ReconnectPrimaryStorageResponse struct {
	Inventory PrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取主存储容量
type GetPrimaryStorageCapacityRequest struct {
	ReqConfig
	ZoneUuids           []string `json:"zoneUuids,omitempty" bson:"zoneUuids,omitempty"`                     //区域UUID列表
	ClusterUuids        []string `json:"clusterUuids,omitempty" bson:"clusterUuids,omitempty"`               //集群UUID列表
	PrimaryStorageUuids []string `json:"primaryStorageUuids,omitempty" bson:"primaryStorageUuids,omitempty"` //主存储UUID列表
	All                 bool     `json:"all,omitempty" bson:"all,omitempty"`                                 //当主存储UUID列表为空时，该项为真表示查询系统中所有的主存储。
	SystemTags          []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`                   //云主机系统标签
	UserTags            []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetPrimaryStorageCapacityResponse struct {
	TotalCapacity             int64     `json:"totalCapacity" bson:"totalCapacity"`
	AvailableCapacity         int64     `json:"availableCapacity" bson:"availableCapacity"`
	TotalPhysicalCapacity     int64     `json:"totalPhysicalCapacity" bson:"totalPhysicalCapacity"`
	AvailablePhysicalCapacity int64     `json:"availablePhysicalCapacity" bson:"availablePhysicalCapacity"`
	Success                   bool      `json:"success" bson:"success"`
	Error                     ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//刷新主存储容量
type SyncPrimaryStorageCapacityRequest struct {
	ReqConfig
	PrimaryStorageUuid         string                           `json:"primaryStorageUuid" bson:"primaryStorageUuid"` //主存储UUID列表
	SyncPrimaryStorageCapacity SyncPrimaryStorageCapacityParams `json:"syncPrimaryStorageCapacity" bson:"syncPrimaryStorageCapacity"`
	SystemTags                 []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type SyncPrimaryStorageCapacityParams struct {
}

type SyncPrimaryStorageCapacityResponse struct {
	Inventory PrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更改主存储状态
type ChangePrimaryStorageStateRequest struct {
	ReqConfig
	Uuid                      string                          `json:"uuid" bson:"uuid"` //集群UUID
	ChangePrimaryStorageState ChangePrimaryStorageStateParams `json:"changePrimaryStorageState" bson:"changePrimaryStorageState"`
	SystemTags                []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                  []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangePrimaryStorageStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //主存储的目标状态:enable,disable,maintain,deleting
}

type ChangePrimaryStorageStateResponse struct {
	Inventory PrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新主存储信息
type UpdatePrimaryStorageRequest struct {
	ReqConfig
	Uuid                 string                     `json:"uuid" bson:"uuid"` //集群UUID
	UpdatePrimaryStorage UpdatePrimaryStorageParams `json:"updatePrimaryStorage" bson:"updatePrimaryStorage"`
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdatePrimaryStorageParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"` //资源名称
	Url         string `json:"url,omitempty" bson:"url,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
}

type UpdatePrimaryStorageResponse struct {
	Inventory PrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//清除主存储镜像缓存
type CleanUpImageCacheOnPrimaryStorageRequest struct {
	ReqConfig
	Uuid                              string                                  `json:"uuid" bson:"uuid"` //集群UUID
	CleanUpImageCacheOnPrimaryStorage CleanUpImageCacheOnPrimaryStorageParams `json:"cleanUpImageCacheOnPrimaryStorage" bson:"cleanUpImageCacheOnPrimaryStorage"`
	SystemTags                        []string                                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                          []string                                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CleanUpImageCacheOnPrimaryStorageParams struct {
}

type CleanUpImageCacheOnPrimaryStorageResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取主存储分配策略清单
type GetPrimaryStorageAllocatorStrategiesRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetPrimaryStorageAllocatorStrategiesResponse struct {
	Strategies []string  `bson:"strategies" json:"strategies"`
	Error      ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取主存储类型列表
type GetPrimaryStorageTypesRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetPrimaryStorageTypesResponse struct {
	Types []string  `bson:"types" json:"types"`
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取候选列表
type GetPrimaryStorageCandidatesForVolumeMigrationRequest struct {
	ReqConfig
	VolumeUuid string   `json:"volumeUuid" bson:"volumeUuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetPrimaryStorageCandidatesForVolumeMigrationResponse struct {
	Inventories []PrimaryStorageInventory `bson:"inventories" json:"inventories"`
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取跨存储迁移可选物理机列表
type GetHostCandidatesForVmMigrationRequest struct {
	ReqConfig
	VmInstanceUuid        string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`               //云主机UUID
	DstPrimaryStorageUuid string   `json:"dstPrimaryStorageUuid" bson:"dstPrimaryStorageUuid"` //目的主存储UUID
	Limit                 int      `json:"limit,omitempty" bson:"limit,omitempty"`             //数量限制
	SystemTags            []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`   //云主机系统标签
	UserTags              []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetHostCandidatesForVmMigrationResponse struct {
	Inventories []HostInventory `bson:"inventories" json:"inventories"`
	Error       ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//跨主存储迁移云盘
type PrimaryStorageMigrateVolumeRequest struct {
	ReqConfig
	VolumeUuid                  string                            `json:"volumeUuid" bson:"volumeUuid"`
	PrimaryStorageMigrateVolume PrimaryStorageMigrateVolumeParams `json:"primaryStorageMigrateVolume" bson:"primaryStorageMigrateVolume"`
	SystemTags                  []string                          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                    []string                          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type PrimaryStorageMigrateVolumeParams struct {
	DstPrimaryStorageUuid string `json:"dstPrimaryStorageUuid" bson:"dstPrimaryStorageUuid"` //目标主存储UUID
}

type PrimaryStorageMigrateVolumeResponse struct {
	Inventory VolumeInventory `bson:"inventory" json:"inventory"`
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//跨主存储迁移云主机
type PrimaryStorageMigrateVmRequest struct {
	ReqConfig
	VmInstanceUuid          string                        `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	PrimaryStorageMigrateVm PrimaryStorageMigrateVmParams `json:"primaryStorageMigrateVm" bson:"primaryStorageMigrateVm"`
	SystemTags              []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type PrimaryStorageMigrateVmParams struct {
	DstPrimaryStorageUuid string `json:"dstPrimaryStorageUuid" bson:"dstPrimaryStorageUuid"`         //目标主存储UUID
	WithDataVolumes       bool   `json:"withDataVolumes,omitempty" bson:"withDataVolumes,omitempty"` //迁移包含云盘
	WithSnapshots         bool   `json:"withSnapshots,omitempty" bson:"withSnapshots,omitempty"`     //迁移包含快照
	DstHostUuid           string `json:"dstHostUuid" bson:"dstHostUuid"`
}

type PrimaryStorageMigrateVmResponse struct {
	Inventory VmInstanceInventory `bson:"inventory" json:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取存储迁移候选列表
type GetPrimaryStorageCandidatesForVmMigrationRequest struct {
	ReqConfig
	VmInstanceUuid                            string                                          `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	GetPrimaryStorageCandidatesForVmMigration GetPrimaryStorageCandidatesForVmMigrationParams `json:"getPrimaryStorageCandidatesForVmMigration" bson:"getPrimaryStorageCandidatesForVmMigration"`
	SystemTags                                []string                                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                                  []string                                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetPrimaryStorageCandidatesForVmMigrationParams struct {
	WithDataVolumes bool `json:"withDataVolumes,omitempty" bson:"withDataVolumes,omitempty"` //迁移包含云盘
}

type GetPrimaryStorageCandidatesForVmMigrationResponse struct {
	Inventory VmInstanceInventory `bson:"inventory" json:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取主存储的License信息
type GetPrimaryStorageLicenseInfoRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetPrimaryStorageLicenseInfoResponse struct {
	Uuid       string    `json:"uuid" bson:"uuid"`
	Name       string    `json:"name" bson:"name"` //资源名称
	ExpireTime string    `json:"expireTime" bson:"expireTime"`
	Success    bool      `json:"success" bson:"success"`
	Error      ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加本地存储为主存储
type AddLocalPrimaryStorageRequest struct {
	ReqConfig
	Params     AddLocalPrimaryStorageParams `json:"params" bson:"params"`
	SystemTags []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddLocalPrimaryStorageParams struct {
	Url          string `json:"url" bson:"url"`
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //资源的详细描述
	Type         string `json:"type,omitempty" bson:"type,omitempty"`                 //云主机类型 保留字段，无需指定。UserVm/ApplianceVm
	ZoneUuid     string `json:"zoneUuid" bson:"zoneUuid"`                             //区域UUID 若指定，云主机会在指定区域创建。
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID
}

type AddLocalPrimaryStorageResponse struct {
	Inventory PrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询本地存储资源引用
type QueryLocalStorageResourceRefRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryLocalStorageResourceRefResponse struct {
	Inventories []PrimaryStorageInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//迁移本地存储上存放的云盘
type LocalStorageMigrateVolumeRequest struct {
	ReqConfig
	VolumeUuid                string                          `json:"volumeUuid" bson:"volumeUuid"` //云盘UUID
	LocalStorageMigrateVolume LocalStorageMigrateVolumeParams `json:"localStorageMigrateVolume" bson:"localStorageMigrateVolume"`
	SystemTags                []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                  []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type LocalStorageMigrateVolumeParams struct {
	DestHostUuid string `json:"destHostUuid" bson:"destHostUuid"`
}

type LocalStorageMigrateVolumeResponse struct {
	Inventory LocalStorageResourceRefInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取主机本地存储容量
type GetLocalStorageHostDiskCapacityRequest struct {
	ReqConfig
	HostUuid           string   `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`     //	物理机UUID
	PrimaryStorageUuid string   `json:"primaryStorageUuid" bson:"primaryStorageUuid"`     //主存储UUID
	SystemTags         []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetLocalStorageHostDiskCapacityResponse struct {
	Inventories []LocalStorageResourceRefInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取迁移本地存储物理机
type LocalStorageGetVolumeMigratableHostsRequest struct {
	ReqConfig
	VolumeUuid string   `json:"volumeUuid" bson:"volumeUuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type LocalStorageGetVolumeMigratableHostsResponse struct {
	Inventories []HostInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加NFS主存储
type AddNfsPrimaryStorageRequest struct {
	ReqConfig
	Params     AddNfsPrimaryStorageParams `json:"params" bson:"params"`
	SystemTags []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddNfsPrimaryStorageParams struct {
	Url          string `json:"url" bson:"url"`
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //资源的详细描述
	Type         string `json:"type,omitempty" bson:"type,omitempty"`                 //云主机类型 保留字段，无需指定。UserVm/ApplianceVm
	ZoneUuid     string `json:"zoneUuid" bson:"zoneUuid"`                             //区域UUID 若指定，云主机会在指定区域创建。
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID
}

type AddNfsPrimaryStorageResponse struct {
	Inventory PrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加一个共享挂载点的主存储
type AddSharedMountPointPrimaryStorageRequest struct {
	ReqConfig
	Params     AddSharedMountPointPrimaryStorageParams `json:"params" bson:"params"`
	SystemTags []string                                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddSharedMountPointPrimaryStorageParams struct {
	Url          string `json:"url" bson:"url"`
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //资源的详细描述
	Type         string `json:"type,omitempty" bson:"type,omitempty"`                 //云主机类型 保留字段，无需指定。UserVm/ApplianceVm
	ZoneUuid     string `json:"zoneUuid" bson:"zoneUuid"`                             //区域UUID 若指定，云主机会在指定区域创建。
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID
}

type AddSharedMountPointPrimaryStorageResponse struct {
	Inventory PrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加Ceph主存储
type AddCephPrimaryStorageRequest struct {
	ReqConfig
	Params     AddNfsPrimaryStorageParams `json:"params" bson:"params"`
	SystemTags []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddCephPrimaryStorageParams struct {
	MonUrls            []string `json:"monUrls" bson:"monUrls"`
	RootVolumePoolName string   `json:"rootVolumePoolName,omitempty" bson:"rootVolumePoolName,omitempty"`
	DataVolumePoolName string   `json:"dataVolumePoolName,omitempty" bson:"dataVolumePoolName,omitempty"` //资源名称
	ImageCachePoolName string   `json:"imageCachePoolName,omitempty" bson:"imageCachePoolName,omitempty"` //资源的详细描述
	Url                string   `json:"url" bson:"url"`
	Name               string   `json:"name" bson:"name"`                                     //资源名称
	Description        string   `json:"description,omitempty" bson:"description,omitempty"`   //资源的详细描述
	Type               string   `json:"type,omitempty" bson:"type,omitempty"`                 //云主机类型 保留字段，无需指定。UserVm/ApplianceVm
	ZoneUuid           string   `json:"zoneUuid" bson:"zoneUuid"`                             //区域UUID 若指定，云主机会在指定区域创建。
	ResourceUuid       string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID
}

type AddCephPrimaryStorageResponse struct {
	Inventory PrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询Ceph主存储
type QueryCephPrimaryStorageRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryCephPrimaryStorageResponse struct {
	Inventories []PrimaryStorageInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//为Ceph主存储添加mon节点
type AddMonToCephPrimaryStorageRequest struct {
	ReqConfig
	Uuid       string                           `json:"uuid" bson:"uuid"`
	Params     AddMonToCephPrimaryStorageParams `json:"params" bson:"params"`
	SystemTags []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddMonToCephPrimaryStorageParams struct {
	MonUrls []string `json:"monUrls" bson:"monUrls"`
}

type AddMonToCephPrimaryStorageResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从Ceph主存储删除mon节点
type RemoveMonFromCephPrimaryStorageRequest struct {
	ReqConfig
	Uuid         string   `json:"uuid" bson:"uuid"`
	MonHostnames []string `json:"monHostnames" bson:"monHostnames"`
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveMonFromCephPrimaryStorageResponse struct {
	Inventory CephPrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新Ceph主存储mon节点
type UpdateCephPrimaryStorageMonRequest struct {
	ReqConfig
	MonUuid                     string                            `json:"monUuid" bson:"monUuid"` //mon节点UUID
	UpdateCephPrimaryStorageMon UpdateCephPrimaryStorageMonParams `json:"updateCephPrimaryStorageMon" bson:"updateCephPrimaryStorageMon"`
	SystemTags                  []string                          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                    []string                          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateCephPrimaryStorageMonParams struct {
	Hostname    string `json:"hostname" bson:"hostname"`       //mon节点新主机IP地址
	SshUsername string `json:"sshUsername" bson:"sshUsername"` //mon节点主机 ssh 用户名
	SshPassword string `json:"sshPassword" bson:"sshPassword"` //mon节点主机 ssh 用户密码
	SshPort     string `json:"sshPort" bson:"sshPort"`         //mon节点主机 ssh 端口
	MonPort     int    `json:"monPort" bson:"monPort"`         //mon的新端口
}

type UpdateCephPrimaryStorageMonResponse struct {
	Inventory CephPrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加Ceph主存储池
type AddCephPrimaryStoragePoolRequest struct {
	ReqConfig
	PrimaryStorageUuid string                          `json:"primaryStorageUuid" bson:"primaryStorageUuid"` //mon节点UUID
	Params             AddCephPrimaryStoragePoolParams `json:"params" bson:"params"`
	SystemTags         []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddCephPrimaryStoragePoolParams struct {
	PoolName     string `json:"poolName" bson:"poolName"`
	AliasName    string `json:"aliasName,omitempty" bson:"aliasName,omitempty"`
	Description  string `json:"description,omitempty" bson:"description,omitempty"` //描述
	IsCreate     bool   `json:"isCreate,omitempty" bson:"isCreate,omitempty"`       //mon节点主机 ssh 端口
	Type         string `json:"type" bson:"type"`
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`
}

type AddCephPrimaryStoragePoolResponse struct {
	Inventory CephPrimaryStoragePoolInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除Ceph主存储池
type DeleteCephPrimaryStoragePoolRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`                                 //mon节点UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteCephPrimaryStoragePoolResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询Ceph主存储池
type QueryCephPrimaryStoragePoolRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //mon节点UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryCephPrimaryStoragePoolResponse struct {
	Inventories []CephPrimaryStoragePoolInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新Ceph主存储池
type UpdateCephPrimaryStoragePoolRequest struct {
	ReqConfig
	Uuid                         string                             `json:"uuid" bson:"uuid"` //mon节点UUID
	UpdateCephPrimaryStoragePool UpdateCephPrimaryStoragePoolParams `json:"updateCephPrimaryStoragePool" bson:"updateCephPrimaryStoragePool"`
	SystemTags                   []string                           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                     []string                           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateCephPrimaryStoragePoolParams struct {
	AliasName   string `json:"aliasName,omitempty" bson:"aliasName,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"` //描述
}

type UpdateCephPrimaryStoragePoolResponse struct {
	Inventory CephPrimaryStoragePoolInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加Shared Block主存储
type AddSharedBlockGroupPrimaryStorageRequest struct {
	ReqConfig
	Params     UpdateCephPrimaryStoragePoolParams `json:"params" bson:"params"`
	SystemTags []string                           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddSharedBlockGroupPrimaryStorageParams struct {
	DiskUuids    []string `json:"diskUuids" bson:"diskUuids"`                         //磁盘唯一表示（例如UUID, WWN, WWID）
	Name         string   `json:"name" bson:"name"`                                   //资源名称
	Description  string   `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
	Type         string   `json:"type,omitempty" bson:"type,omitempty"`               //主存储类型，此处为 SharedBlock
	Url          string   `json:"url" bson:"url"`
	ZoneUuid     string   `json:"zoneUuid" bson:"zoneUuid"`                             //区域UUID 若指定，云主机会在指定区域创建。
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID
}

type AddSharedBlockGroupPrimaryStorageResponse struct {
	Inventory PrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询连接状态
type QuerySharedBlockGroupPrimaryStorageHostRefRequest struct {
	ReqConfig
	PrimaryStorageUuid string   `json:"primaryStorageUuid" bson:"primaryStorageUuid"`
	SystemTags         []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySharedBlockGroupPrimaryStorageHostRefResponse struct {
	Inventories []SharedBlockGroupPrimaryStorageHostRefInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                                        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询Shared Block主存储
type QuerySharedBlockGroupPrimaryStorageRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySharedBlockGroupPrimaryStorageResponse struct {
	Inventories []SharedBlockGroupPrimaryStorageInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加LUN到Shared Block主存储
type AddSharedBlockToSharedBlockGroupRequest struct {
	ReqConfig
	Uuid       string                                 `json:"uuid" bson:"uuid"`
	Params     AddSharedBlockToSharedBlockGroupParams `json:"params" bson:"params"`
	SystemTags []string                               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddSharedBlockToSharedBlockGroupParams struct {
	DiskUuid string `json:"diskUuid" bson:"diskUuid"` //磁盘维一标示（例如UUID, WWN, WWID）
}

type AddSharedBlockToSharedBlockGroupResponse struct {
	Inventory SharedBlockGroupPrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取共享块设备候选清单
type GetSharedBlockCandidateRequest struct {
	ReqConfig
	ClusterUuid string   `json:"clusterUuid" bson:"clusterUuid"`
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetSharedBlockCandidateResponse struct {
	Inventory SharedBlockGroupPrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//刷新共享块设备容量
type RefreshSharedblockDeviceCapacityRequest struct {
	ReqConfig
	SharedBlockGroupUuid string                                 `json:"sharedBlockGroupUuid" bson:"sharedBlockGroupUuid"`
	Uuid                 string                                 `json:"uuid" bson:"uuid"`
	Params               RefreshSharedblockDeviceCapacityParams `json:"params" bson:"params"`
	SystemTags           []string                               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RefreshSharedblockDeviceCapacityParams struct {
}

type RefreshSharedblockDeviceCapacityResponse struct {
	Inventory SharedBlockGroupPrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新共享块存储中共享块的diskUuid
type UpdateSharedBlockDiskUuidRequest struct {
	ReqConfig
	SharedBlockGroupUuid string                                 `json:"sharedBlockGroupUuid" bson:"sharedBlockGroupUuid"`
	Uuid                 string                                 `json:"uuid" bson:"uuid"`
	Params               RefreshSharedblockDeviceCapacityParams `json:"params" bson:"params"`
	SystemTags           []string                               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateSharedBlockDiskUuidParams struct {
	DiskUuid string `json:"diskUuid" bson:"diskUuid"` //磁盘维一标示（例如UUID, WWN, WWID）
}

type UpdateSharedBlockDiskUuidResponse struct {
	Inventory SharedBlockGroupPrimaryStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询共享块设备
type QuerySharedBlockRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySharedBlockResponse struct {
	Inventories []SharedBlocks `json:"inventories" bson:"inventories"`
	Error       ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}
