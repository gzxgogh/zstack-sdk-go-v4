package model

//删除镜像服务器
type DeleteBackupStorageRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteBackupStorageResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询镜像服务器
type QueryBackupStorageRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryBackupStorageResponse struct {
	Inventories []BackupStorageInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//重连镜像服务器
type ReconnectBackupStorageRequest struct {
	ReqConfig
	UUID                   string                       `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ReconnectBackupStorage ReconnectBackupStorageParams `json:"reconnectBackupStorage" bson:"reconnectBackupStorage"`
	SystemTags             []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags               []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ReconnectBackupStorageParams struct {
}

type ReconnectBackupStorageResponse struct {
	Inventory BackupStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更改镜像服务器可用状态
type ChangeBackupStorageStateRequest struct {
	ReqConfig
	UUID                     string                         `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeBackupStorageState ChangeBackupStorageStateParams `json:"changeBackupStorageState" bson:"changeBackupStorageState"`
	SystemTags               []string                       `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                 []string                       `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeBackupStorageStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //镜像服务器的目标状态:enable,disable
}

type ChangeBackupStorageStateResponse struct {
	Inventory BackupStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取镜像服务器存储容量
type GetBackupStorageCapacityRequest struct {
	ReqConfig
	ZoneUuids          []string `json:"zoneUuids,omitempty" bson:"zoneUuids,omitempty"`                   //区域UUID列表
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty" bson:"backupStorageUuids,omitempty"` //镜像服务器UUID列表:zoneUuids, backupStorageUuids 至少有一个不为空,若zoneUuids, backupStorageUuids均为空，all为真
	All                bool     `json:"all,omitempty" bson:"all,omitempty"`                               //当镜像服务器UUID列表为空时，该项为真表示查询系统中所有的镜像服务器
	SystemTags         []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`                 //云主机系统标签
	UserTags           []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetBackupStorageCapacityResponse struct {
	TotalCapacity     int64     `json:"totalCapacity" bson:"totalCapacity"`
	AvailableCapacity int64     `json:"availableCapacity" bson:"availableCapacity"`
	Success           bool      `json:"success" bson:"success"`
	Error             ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取镜像服务器类型列表
type GetBackupStorageTypesRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetBackupStorageTypesResponse struct {
	Types []string  `json:"types" bson:"types"`
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新镜像服务器信息
type UpdateBackupStorageRequest struct {
	ReqConfig
	UUID                string                    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateBackupStorage UpdateBackupStorageParams `json:"updateBackupStorage" bson:"updateBackupStorage"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateBackupStorageParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
}

type UpdateBackupStorageResponse struct {
	Inventory BackupStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从镜像服务器导出镜像
type ExportImageFromBackupStorageRequest struct {
	ReqConfig
	BackupStorageUuid            string                             `json:"backupStorageUuid" bson:"backupStorageUuid"` //资源的UUID，唯一标示该资源
	ExportImageFromBackupStorage ExportImageFromBackupStorageParams `json:"exportImageFromBackupStorage" bson:"exportImageFromBackupStorage"`
	SystemTags                   []string                           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                     []string                           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ExportImageFromBackupStorageParams struct {
	ImageUuid    string `json:"imageUuid" bson:"imageUuid"`                           //镜像UUID
	ExportFormat string `json:"exportFormat,omitempty" bson:"exportFormat,omitempty"` //导出镜像的保存格式
}

type ExportImageFromBackupStorageResponse struct {
	ImageUrl     string    `json:"imageUrl" bson:"imageUrl"`               //被导出镜像的访问地址
	ExportMd5Sum string    `json:"exportMd5Sum" bson:"exportMd5Sum"`       //被导出镜像的md5值
	Success      bool      `json:"success" bson:"success"`                 //是否成功
	Error        ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从镜像服务器删除导出的镜像
type DeleteExportedImageFromBackupStorageRequest struct {
	ReqConfig
	BackupStorageUuid string   `json:"backupStorageUuid" bson:"backupStorageUuid"` //资源的UUID，唯一标示该资源
	ImageUuid         string   `json:"imageUuid" bson:"imageUuid"`
	SystemTags        []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteExportedImageFromBackupStorageResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//挂载镜像服务器至区域
type AttachBackupStorageToZoneRequest struct {
	ReqConfig
	ZoneUUID          string   `json:"zoneUuid" bson:"zoneUuid"` //区域UUID
	BackupStorageUuid string   `json:"backupStorageUuid" bson:"backupStorageUuid"`
	SystemTags        []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachBackupStorageToZoneResponse struct {
	Inventory BackupStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从区域中卸载已经挂载的镜像服务器
type DetachBackupStorageFromZoneRequest struct {
	ReqConfig
	ZoneUUID          string   `json:"zoneUuid" bson:"zoneUuid"` //区域UUID
	BackupStorageUuid string   `json:"backupStorageUuid" bson:"backupStorageUuid"`
	SystemTags        []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachBackupStorageFromZoneResponse struct {
	Inventory BackupStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//跨镜像服务器迁移镜像
type BackupStorageMigrateImageRequest struct {
	ReqConfig
	ImageUuid                 string                          `json:"imageUuid" bson:"imageUuid"`
	BackupStorageMigrateImage BackupStorageMigrateImageParams `json:"backupStorageMigrateImage" bson:"backupStorageMigrateImage"`
	SystemTags                []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                  []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type BackupStorageMigrateImageParams struct {
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" bson:"srcBackupStorageUuid"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" bson:"dstBackupStorageUuid"`
}

type BackupStorageMigrateImageResponse struct {
	Inventory ImageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取候选列表
type GetBackupStorageCandidatesForImageMigrationRequest struct {
	ReqConfig
	SrcBackupStorageUuid string   `json:"srcBackupStorageUuid" bson:"srcBackupStorageUuid"`
	SystemTags           []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetBackupStorageCandidatesForImageMigrationResponse struct {
	Inventories []BackupStorageInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加镜像仓库服务器
type AddImageStoreBackupStorageRequest struct {
	ReqConfig
	Params     AddImageStoreBackupStorageParams `json:"params" bson:"params"`
	SystemTags []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddImageStoreBackupStorageParams struct {
	Hostname     string `json:"hostname" bson:"hostname"`                             //服务器主机地址
	Username     string `json:"username" bson:"username"`                             //服务器 SSH 用户名 (用于 Ansible 部署)
	Password     string `json:"password" bson:"password"`                             //服务器 SSH 用户密码
	SshPort      string `json:"sshPort,omitempty" bson:"sshPort,omitempty"`           //服务器 SSH 端口
	Name         string `json:"name" bson:"name"`                                     //镜像仓库服务器名称
	Url          string `json:"url" bson:"url"`                                       //镜像仓库本地数据存放路径
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //镜像仓库的详细描述
	Type         string `json:"type,omitempty" bson:"type,omitempty"`                 //这里是 ImageStoreBackupStorage
	ImportImages bool   `json:"importImages,omitempty" json:"importImages,omitempty"` //是否导入镜像
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type AddImageStoreBackupStorageResponse struct {
	Inventory BackupStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询镜像仓库服务器
type QueryImageStoreBackupStorageRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryImageStoreBackupStorageResponse struct {
	Inventories []BackupStorageInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//重连镜像仓库服务器
type ReconnectImageStoreBackupStorageRequest struct {
	ReqConfig
	UUID                             string                                 `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ReconnectImageStoreBackupStorage ReconnectImageStoreBackupStorageParams `json:"reconnectImageStoreBackupStorage" bson:"reconnectImageStoreBackupStorage"`
	SystemTags                       []string                               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                         []string                               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ReconnectImageStoreBackupStorageParams struct {
}

type ReconnectImageStoreBackupStorageResponse struct {
	Inventory []ImageStoreBackupStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新镜像仓库服务器信息
type UpdateImageStoreBackupStorageRequest struct {
	ReqConfig
	UUID                          string                              `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateImageStoreBackupStorage UpdateImageStoreBackupStorageParams `json:"updateImageStoreBackupStorage" bson:"updateImageStoreBackupStorage"`
	SystemTags                    []string                            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                      []string                            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateImageStoreBackupStorageParams struct {
	Username    string `json:"username,omitempty" bson:"username,omitempty"`       //服务器 SSH 用户名 (用于 Ansible 部署)
	Password    string `json:"password,omitempty" bson:"password,omitempty"`       //服务器 SSH 用户密码
	Hostname    string `json:"hostname,omitempty" bson:"hostname,omitempty"`       //服务器主机地址
	SshPort     string `json:"sshPort,omitempty" bson:"sshPort,omitempty"`         //服务器 SSH 端口
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
}

type UpdateImageStoreBackupStorageResponse struct {
	Inventory BackupStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从镜像仓库回收磁盘空间
type ReclaimSpaceFromImageStoreRequest struct {
	ReqConfig
	UUID                       string                           `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ReclaimSpaceFromImageStore ReclaimSpaceFromImageStoreParams `json:"reclaimSpaceFromImageStore" bson:"reclaimSpaceFromImageStore"`
	SystemTags                 []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ReclaimSpaceFromImageStoreParams struct {
}

type ReclaimSpaceFromImageStoreResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加Ceph镜像服务器
type AddCephBackupStorageRequest struct {
	ReqConfig
	Params     AddCephBackupStorageParams `json:"params" bson:"params"`
	SystemTags []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddCephBackupStorageParams struct {
	MonUrls      []string `json:"monUrls" bson:"monUrls"`                               //Ceph mon 的地址列表
	PoolName     string   `json:"poolName,omitempty" bson:"poolName,omitempty"`         //用于存放镜像的 Ceph pool 的名字
	Name         string   `json:"name" bson:"name"`                                     //镜像仓库服务器名称
	Url          string   `json:"url" bson:"url"`                                       //镜像仓库本地数据存放路径
	Description  string   `json:"description,omitempty" bson:"description,omitempty"`   //镜像仓库的详细描述
	Type         string   `json:"type,omitempty" bson:"type,omitempty"`                 //镜像服务器的类型，此处为 Ceph
	ImportImages bool     `json:"importImages,omitempty" json:"importImages,omitempty"` //是否导入镜像
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type AddCephBackupStorageResponse struct {
	Inventory BackupStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询Ceph镜像服务器
type QueryCephBackupStorageRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryCephBackupStorageResponse struct {
	Inventories []BackupStorageInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新Ceph镜像服务器mon节点
type UpdateCephBackupStorageMonRequest struct {
	ReqConfig
	MonUuid                    string                           `json:"monUuid" bson:"monUuid"` //资源的UUID，唯一标示该资源
	UpdateCephBackupStorageMon UpdateCephBackupStorageMonParams `json:"updateCephBackupStorageMon" bson:"updateCephBackupStorageMon"`
	SystemTags                 []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateCephBackupStorageMonParams struct {
	Hostname    string `json:"hostname,omitempty" bson:"hostname,omitempty"`       //mon节点新主机IP地址
	SshUsername string `json:"sshUsername,omitempty" bson:"sshUsername,omitempty"` //mon节点主机 ssh 用户名
	SshPassword string `json:"sshPassword,omitempty" bson:"sshPassword,omitempty"` //mon节点主机 ssh 用户密码
	SshPort     string `json:"sshPort,omitempty" bson:"sshPort,omitempty"`         //mon节点主机 ssh 端口
	MonPort     string `json:"monPort ,omitempty" bson:"monPort ,omitempty"`       //mon节点的端口
}

type UpdateCephBackupStorageMonResponse struct {
	Inventory CephBackupStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//Ceph镜像服务器添加mon节点
type AddMonToCephBackupStorageRequest struct {
	ReqConfig
	UUID       string                          `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Params     AddMonToCephBackupStorageParams `json:"params" bson:"params"`
	SystemTags []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddMonToCephBackupStorageParams struct {
	MonUrls []string `json:"monUrls" bson:"monUrls"`
}

type AddMonToCephBackupStorageResponse struct {
	Inventory CephBackupStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//Ceph镜像服务器删除mon节点
type RemoveMonFromCephBackupStorageRequest struct {
	ReqConfig
	UUID         string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	MonHostnames []string `json:"monHostnames" bson:"monHostnames"`
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveMonFromCephBackupStorageResponse struct {
	Inventory CephBackupStorageInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type BackupStorageInventory struct {
	UUID              string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name              string   `json:"name" bson:"name"` //资源名称
	Url               string   `json:"url" bson:"url"`
	Description       string   `json:"description" bson:"description"` //资源的详细描述
	TotalCapacity     int64    `json:"totalCapacity" bson:"totalCapacity"`
	AvailableCapacity int64    `json:"availableCapacity" bson:"availableCapacity"`
	Type              string   `json:"type" bson:"type"` //云主机类型 保留字段，无需指定。UserVm/ApplianceVm
	State             string   `json:"state" bson:"state"`
	Status            string   `json:"status" bson:"status"`
	CreateDate        string   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate        string   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	AttachedZoneUuids []string `json:"attachedZoneUuids" bson:"attachedZoneUuids"`
}

type ImageStoreBackupStorageInventory struct {
	Hostname          string   `json:"hostname" bson:"hostname"`                           //服务器主机地址
	Username          string   `json:"username" bson:"username"`                           //服务器 SSH 用户名 (用于 Ansible 部署)
	SshPort           string   `json:"sshPort,omitempty" bson:"sshPort,omitempty"`         //服务器 SSH 端口
	UUID              string   `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name              string   `json:"name" bson:"name"`                                   //镜像仓库服务器名称
	Url               string   `json:"url" bson:"url"`                                     //镜像仓库本地数据存放路径
	Description       string   `json:"description,omitempty" bson:"description,omitempty"` //镜像仓库的详细描述
	TotalCapacity     int64    `json:"totalCapacity" bson:"totalCapacity"`
	AvailableCapacity int64    `json:"availableCapacity" bson:"availableCapacity"`
	Type              string   `json:"type,omitempty" bson:"type,omitempty"` //这里是 ImageStoreBackupStorage
	State             string   `json:"state" bson:"state"`
	Status            string   `json:"status" bson:"status"`
	CreateDate        string   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate        string   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	AttachedZoneUuids []string `json:"attachedZoneUuids" bson:"attachedZoneUuids"`
}

type CephBackupStorageInventory struct {
	Fsid                  string   `json:"fsid" bson:"fsid"`
	PoolName              string   `json:"poolName" bson:"poolName"`
	UUID                  string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name                  string   `json:"name" bson:"name"` //资源名称
	Url                   string   `json:"url" bson:"url"`
	Description           string   `json:"description" bson:"description"` //资源的详细描述
	TotalCapacity         int64    `json:"totalCapacity" bson:"totalCapacity"`
	AvailableCapacity     int64    `json:"availableCapacity" bson:"availableCapacity"`
	Type                  string   `json:"type,omitempty" bson:"type,omitempty"` //云主机类型 保留字段，无需指定。UserVm/ApplianceVm
	State                 string   `json:"state" bson:"state"`
	Status                string   `json:"status" bson:"status"`
	CreateDate            string   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate            string   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	AttachedClusterUuids  []string `json:"attachedClusterUuids" bson:"attachedClusterUuids"`
	PoolUsedCapacity      int64    `json:"poolUsedCapacity" bson:"poolUsedCapacity"`
	PoolReplicatedSize    int      `json:"poolReplicatedSize" bson:"poolReplicatedSize"`
	PoolAvailableCapacity int64    `json:"poolAvailableCapacity" bson:"poolAvailableCapacity"`
	Mons                  Mons     `json:"mons" bson:"mons"`
}
