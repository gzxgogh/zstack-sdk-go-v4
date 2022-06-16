package model

//添加镜像
type AddImageRequest struct {
	ReqConfig
	Params     AddImageParams `json:"params" bson:"params"`
	SystemTags []string       `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string       `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type AddImageParams struct {
	Name               string   `json:"name" bson:"name"`                                     //镜像名称
	Description        string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	Url                string   `json:"url" bson:"url"`                                       //被添加镜像的URL地址
	MediaType          string   `json:"mediaType,omitempty" bson:"mediaType,omitempty"`       //镜像的类型,RootVolumeTemplate,ISO,DataVolumeTemplate
	GuestOsType        string   `json:"guestOsType,omitempty" bson:"guestOsType,omitempty"`   //镜像对应客户机操作系统的类型
	System             string   `json:"system,omitempty" bson:"system,omitempty"`             //是否系统镜像（如，云路由镜像）
	Format             string   `json:"format" bson:"format"`                                 //镜像的格式，比如：raw
	Platform           string   `json:"platform,omitempty" bson:"platform,omitempty"`         //镜像的系统平台,Linux,Windows,WindowsVirtio,Other,Paravirtualization
	BackupStorageUuids []string `json:"backupStorageUuids" bson:"backupStorageUuids"`         //指定添加镜像的镜像服务器UUID列表
	Type               string   `json:"type,omitempty" bson:"type,omitempty"`                 //内部使用字段
	ResourceUuid       string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	Architecture       string   `json:"architecture,omitempty" bson:"architecture,omitempty"` //x86_64,aarch64,mips64el
	TagUuids           []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type AddImageResponse struct {
	Inventory ImageInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除镜像
type DeleteImageRequest struct {
	ReqConfig
	Uuid               string   `json:"uuid" bson:"uuid"`                                                 //镜像的UUID，唯一标示该子资源
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty" bson:"backupStorageUuids,omitempty"` //指定添加镜像的镜像服务器UUID列表
	DeleteMode         string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"`                 //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags         []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`                 //云主机系统标签
	UserTags           []string `json:"userTags,omitempty" bson:"userTags,omitempty"`                     //云主机用户标签
}

type DeleteImageResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//彻底删除镜像
type ExpungeImageRequest struct {
	ReqConfig
	ImageUuid    string             `json:"imageUuid" bson:"imageUuid"` //镜像UUID
	ExpungeImage ExpungeImageParams `json:"expungeImage" bson:"expungeImage"`
	SystemTags   []string           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string           `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type ExpungeImageParams struct {
	Uuid               string   `json:"uuid" bson:"uuid"`                               //镜像的UUID，唯一标示该子资源
	BackupStorageUuids []string `json:"backupStorageUuids " bson:"backupStorageUuids "` //指定添加镜像的镜像服务器UUID列表
}

type ExpungeImageResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询镜像
type QueryImageRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //镜像的UUID，唯一标示该子资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type QueryImageResponse struct {
	Success     bool             `json:"success" bson:"success"`
	Inventories []ImageInventory `json:"inventories" bson:"inventories"`         //云主机详细清单
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//恢复镜像
type RecoverImageRequest struct {
	ReqConfig
	ImageUuid    string             `json:"imageUuid" bson:"imageUuid"`                       //镜像UUID
	RecoverImage RecoverImageParams `json:"recoverImage" bson:"recoverImage"`                 //放backupStorageUuids
	SystemTags   []string           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string           `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type RecoverImageParams struct {
	BackupStorageUuids []string `json:"backupStorageUuids " bson:"backupStorageUuids "` //指定添加镜像的镜像服务器UUID列表
}

type RecoverImageResponse struct {
	Inventory ImageInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改镜像状态
type ChangeImageStateRequest struct {
	ReqConfig
	Uuid             string                 `json:"uuid" bson:"uuid"`                                 //镜像的UUID，唯一标示该镜像
	ChangeImageState ChangeImageStateParams `json:"recoverImage" bson:"recoverImage"`                 //放stateEvent
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type ChangeImageStateParams struct {
	StateEvent string `json:"stateEvent " bson:"stateEvent "` //镜像的状态,enable,disable
}

type ChangeImageStateResponse struct {
	Inventory ImageInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新镜像信息
type UpdateImageRequest struct {
	ReqConfig
	Uuid        string            `json:"uuid" bson:"uuid"` //镜像的UUID，唯一标示该镜像
	UpdateImage UpdateImageParams `json:"updateImage" bson:"updateImage"`
	SystemTags  []string          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string          `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type UpdateImageParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //镜像名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	GuestOsType string `json:"guestOsType,omitempty" bson:"guestOsType,omitempty"` //镜像对应的客户机操作系统类型
	MediaType   string `json:"mediaType,omitempty" bson:"mediaType,omitempty"`     //镜像的类型,RootVolumeTemplate,DataVolumeTemplate,ISO
	Format      string `json:"format,omitempty" bson:"format,omitempty"`           //云盘格式,raw,qcow2,iso
	System      string `json:"system,omitempty" bson:"system,omitempty"`           //标识是否为系统镜像
	Platform    string `json:"platform,omitempty" bson:"platform,omitempty"`       //镜像的系统平台,Linux,Windows,WindowsVirtio,Other,Paravirtualization
}

type UpdateImageResponse struct {
	Inventory ImageInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//刷新镜像大小信息
type SyncImageSizeRequest struct {
	ReqConfig
	Uuid          string                 `json:"uuid" bson:"uuid"`                                       //镜像的UUID，唯一标示该镜像
	SyncImageSize map[string]interface{} `json:"syncImageSize,omitempty" bson:"syncImageSize,omitempty"` //放空
	SystemTags    []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"`       //云主机系统标签
	UserTags      []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`           //云主机用户标签
}

type SyncImageSizeResponse struct {
	Inventory ImageInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取镜像服务器候选
type GetCandidateBackupStorageForCreatingImageRequest struct {
	ReqConfig
	VolumeUuid         string   `json:"volumeUuid,omitempty" bson:"volumeUuid,omitempty"`                 //云盘UUID，注意：volumeUuid 和 volumeSnapshotUuid 二选一
	VolumeSnapshotUuid string   `json:"volumeSnapshotUuid,omitempty" bson:"volumeSnapshotUuid,omitempty"` //云盘快照UUID，注意：volumeUuid 和 volumeSnapshotUuid 二选一
	SystemTags         []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`                 //云主机系统标签
	UserTags           []string `json:"userTags,omitempty" bson:"userTags,omitempty"`                     //云主机用户标签
}

type GetCandidateBackupStorageForCreatingImageResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从根云盘创建根云盘镜像
type CreateRootVolumeTemplateFromRootVolumeRequest struct {
	ReqConfig
	RootVolumeUuid string                                       `json:"rootVolumeUuid" bson:"rootVolumeUuid"`             //根云盘UUID
	Params         CreateRootVolumeTemplateFromRootVolumeParams `json:"params" bson:"params"`                             //结构体中的其他参数
	SystemTags     []string                                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string                                     `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateRootVolumeTemplateFromRootVolumeParams struct {
	Name               string   `json:"name" bson:"name"`                                                 //名称
	RootVolumeUuid     string   `json:"rootVolumeUuid" bson:"rootVolumeUuid"`                             //根云盘UUID
	Description        string   `json:"description,omitempty" bson:"description,omitempty"`               //详细描述
	GuestOsType        string   `json:"guestOsType,omitempty" bson:"guestOsType,omitempty"`               //根云盘镜像对应客户机操作系统类型
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty" bson:"backupStorageUuids,omitempty"` //镜像服务器UUID列表
	Platform           string   `json:"platform,omitempty" bson:"platform,omitempty"`                     //镜像的系统平台,Linux,Windows,WindowsVirtio,Other,Paravirtualization
	System             bool     `json:"system,omitempty" bson:"system,omitempty"`                         //是否系统根云盘镜像
	ResourceUuid       string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`             //根云盘镜像UUID。若指定，根云盘镜像会使用该字段值作为UUID。
	Architecture       string   `json:"architecture,omitempty" bson:"architecture,omitempty"`             //x86_64,aarch64,mips64el
	TagUuids           []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`                     //标签UUID列表
}

type CreateRootVolumeTemplateFromRootVolumeResponse struct {
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory ImageInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
}

//从云盘快照创建根云盘镜像
type CreateRootVolumeTemplateFromVolumeSnapshotRequest struct {
	ReqConfig
	SnapshotUuid string                                           `json:"snapshotUuid" bson:"snapshotUuid"`                 //快照UUID
	Params       CreateRootVolumeTemplateFromVolumeSnapshotParams `json:"params" bson:"params"`                             //结构体中的其他参数
	SystemTags   []string                                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string                                         `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateRootVolumeTemplateFromVolumeSnapshotParams struct {
	Name               string   `json:"name" bson:"name"`                                                 //名称
	Description        string   `json:"description,omitempty" bson:"description,omitempty"`               //详细描述
	GuestOsType        string   `json:"guestOsType,omitempty" bson:"guestOsType,omitempty"`               //根云盘镜像对应客户机操作系统类型
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty" bson:"backupStorageUuids,omitempty"` //镜像服务器UUID列表
	Platform           string   `json:"platform,omitempty" bson:"platform,omitempty"`                     //镜像的系统平台,Linux,Windows,WindowsVirtio,Other,Paravirtualization
	System             bool     `json:"system" bson:"system"`                                             //是否系统根云盘镜像
	ResourceUuid       string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`             //根云盘镜像UUID。若指定，根云盘镜像会使用该字段值作为UUID。
}

type CreateRootVolumeTemplateFromVolumeSnapshotResponse struct {
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory ImageInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Failures  Failures       `json:"failures" bson:"failures"`
}

type Failures struct {
	BackupStorageUuid string    `json:"backupStorageUuid" bson:"backupStorageUuid"`
	Error             ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从云盘创建数据云盘镜像
type CreateDataVolumeTemplateFromVolumeRequest struct {
	ReqConfig
	VolumeUuid string                                   `json:"volumeUuid" bson:"volumeUuid"`                     //快照UUID
	Params     CreateDataVolumeTemplateFromVolumeParams `json:"params" bson:"params"`                             //结构体中的其他参数
	SystemTags []string                                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                                 `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateDataVolumeTemplateFromVolumeParams struct {
	Name               string   `json:"name" bson:"name"`                                                 //名称
	Description        string   `json:"description,omitempty" bson:"description,omitempty"`               //详细描述
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty" bson:"backupStorageUuids,omitempty"` //镜像服务器UUID列表
	ResourceUuid       string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`             //根云盘镜像UUID。若指定，根云盘镜像会使用该字段值作为UUID。
}

type CreateDataVolumeTemplateFromVolumeResponse struct {
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory ImageInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
}

//从云盘快照创建数据云盘镜像
type CreateDataVolumeTemplateFromVolumeSnapshotRequest struct {
	ReqConfig
	SnapshotUuid string                                           `json:"snapshotUuid" bson:"snapshotUuid"`                 //快照UUID
	Params       CreateDataVolumeTemplateFromVolumeSnapshotParams `json:"params" bson:"params"`                             //结构体中的其他参数
	SystemTags   []string                                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string                                         `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateDataVolumeTemplateFromVolumeSnapshotParams struct {
	Name               string   `json:"name" bson:"name"`                                                 //名称
	Description        string   `json:"description,omitempty" bson:"description,omitempty"`               //详细描述
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty" bson:"backupStorageUuids,omitempty"` //镜像服务器UUID列表
	ResourceUuid       string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`             //根云盘镜像UUID。若指定，根云盘镜像会使用该字段值作为UUID。
	TagUuids           []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`                     //标签UUID列表
}

type CreateDataVolumeTemplateFromVolumeSnapshotResponse struct {
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory ImageInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Failures  Failures       `json:"failures" bson:"failures"`
}

//获取镜像Qga
type GetImageQgaRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetImageQgaResponse struct {
	Error  ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Enable bool      `json:"enable" bson:"enable"`
}

//设置镜像Qga
type SetImageQgaRequest struct {
	ReqConfig
	Uuid        string            `json:"uuid" bson:"uuid"`
	SetImageQga SetImageQgaParams `json:"setImageQga" bson:"setImageQga"`                   //放enable
	SystemTags  []string          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string          `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetImageQgaParams struct {
	Enable bool `json:"enable" bson:"enable"`
}

type SetImageQgaResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置镜像启动模式
type SetImageBootModeRequest struct {
	ReqConfig
	Uuid             string                 `json:"uuid" bson:"uuid"`
	SetImageBootMode SetImageBootModeParams `json:"setImageBootMode" bson:"setImageBootMode"`         //放bootMode
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetImageBootModeParams struct {
	BootMode string `json:"bootMode" bson:"bootMode"` //镜像启动模式,Legacy,UEFI,UEFI_WITH_CSM
}

type SetImageBootModeResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取上传镜像任务详情
type GetUploadImageJobDetailsRequest struct {
	ReqConfig
	ImageId    string   `json:"imageId" bson:"imageId"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetUploadImageJobDetailsResponse struct {
	Success            bool               `json:"success" bson:"success"`
	ExistingJobDetails ExistingJobDetails `json:"existingJobDetails" bson:"existingJobDetails"`
	Error              ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}
