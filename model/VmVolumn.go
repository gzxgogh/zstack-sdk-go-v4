package model

//创建云盘
type CreateDataVolumeRequest struct {
	ReqConfig
	Params     CreateDataVolumeParams `json:"params" bson:"params"`
	SystemTags []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateDataVolumeParams struct {
	Name               string   `json:"name" bson:"name"`                                                 //云盘名称
	Description        string   `json:"description,omitempty" bson:"description,omitempty"`               //资源的详细描述
	DiskOfferingUuid   string   `json:"diskOfferingUuid" bson:"diskOfferingUuid"`                         //云盘规格UUID
	PrimaryStorageUuid string   `json:"primaryStorageUuid,omitempty" bson:"primaryStorageUuid,omitempty"` //主存储UUID
	ResourceUuid       string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`             //资源UUID
	TagUuids           []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`                     //标签UUID列表
}

type CreateDataVolumeResponse struct {
	Inventory VolumeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除云盘
type DeleteDataVolumeRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode" bson:"deleteMode"`                     //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DeleteDataVolumeResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//彻底删除云盘
type ExpungeDataVolumeRequest struct {
	ReqConfig
	Uuid              string                  `json:"Uuid" bson:"Uuid"`                                               //资源的UUID，唯一标示该资源
	ExpungeDataVolume ExpungeDataVolumeParams `json:"expungeDataVolume,omitempty" bson:"expungeDataVolume,omitempty"` //填空
	SystemTags        []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"`               //云主机系统标签
	UserTags          []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`                   //云主机用户标签
}

type ExpungeDataVolumeParams struct {
}

type ExpungeDataVolumeResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//恢复云盘
type RecoverDataVolumeRequest struct {
	ReqConfig
	Uuid              string            `json:"Uuid" bson:"Uuid"` //资源的UUID，唯一标示该资源
	RecoverDataVolume RecoverDataVolume `json:"recoverDataVolume" bson:"recoverDataVolume"`
	SystemTags        []string          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string          `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type RecoverDataVolume struct {
}

type RecoverDataVolumeResponse struct {
	Inventory VolumeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//开启或关闭云盘
type ChangeVolumeStateRequest struct {
	ReqConfig
	Uuid              string                  `json:"Uuid" bson:"Uuid"`                                 //资源的UUID，唯一标示该资源
	ChangeVolumeState ChangeVolumeStateParams `json:"changeVolumeState" bson:"changeVolumeState"`       //填StateEvent
	SystemTags        []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type ChangeVolumeStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //开启或关闭,enable,disable
}

type ChangeVolumeStateResponse struct {
	Inventory VolumeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从镜像创建云盘
type CreateDataVolumeFromVolumeTemplateRequest struct {
	ReqConfig
	ImageUUID  string                                   `json:"imageUuid" bson:"imageUuid"`                       //镜像UUID 云主机的根云盘会从该字段指定的镜像创建。
	Params     CreateDataVolumeFromVolumeTemplateParams `json:"params" bson:"params"`                             //放name，description,primaryStorageUuid,hostUuid
	SystemTags []string                                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                                 `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateDataVolumeFromVolumeTemplateParams struct {
	Name               string `json:"name" bson:"name"`                                     //云盘名称
	Description        string `json:"description,omitempty" bson:"description,omitempty"`   //资源的详细描述
	PrimaryStorageUuid string `json:"primaryStorageUuid" bson:"primaryStorageUuid"`         //主存储UUID
	HostUUID           string `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`         //物理机UUID
	ResourceUuid       string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID
}

type CreateDataVolumeFromVolumeTemplateResponse struct {
	Inventory VolumeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从快照创建云盘
type CreateDataVolumeFromVolumeSnapshotRequest struct {
	ReqConfig
	VolumeSnapshotUuid string                                   `json:"volumeSnapshotUuid" bson:"volumeSnapshotUuid"`     //云盘快照UUID
	Params             CreateDataVolumeFromVolumeSnapshotParams `json:"params" bson:"params"`                             //放name，description,primaryStorageUuid
	SystemTags         []string                                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string                                 `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateDataVolumeFromVolumeSnapshotParams struct {
	Name               string `json:"name" bson:"name"`                                     //云盘名称
	Description        string `json:"description,omitempty" bson:"description,omitempty"`   //资源的详细描述
	PrimaryStorageUuid string `json:"primaryStorageUuid" bson:"primaryStorageUuid"`         //主存储UUID
	ResourceUuid       string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID
}

type CreateDataVolumeFromVolumeSnapshotResponse struct {
	Inventory VolumeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云盘清单
type QueryVolumeRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type QueryVolumeResponse struct {
	Inventories []VolumeInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云盘格式
type GetVolumeFormatRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVolumeFormatResponse struct {
	Formats []Formats `json:"formats" bson:"formats"`
	Error   ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云盘支持的类型的能力
type GetVolumeCapabilitiesRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`                                 //UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVolumeCapabilitiesResponse struct {
	Capabilities Capabilities `json:"capabilities" bson:"capabilities"`
	Error        ErrorCode    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type Capabilities struct {
	MigrationToOtherPrimaryStorage   bool `json:"MigrationToOtherPrimaryStorage" bson:"MigrationToOtherPrimaryStorage"`
	MigrationInCurrentPrimaryStorage bool `json:"MigrationInCurrentPrimaryStorage" bson:"MigrationInCurrentPrimaryStorage"`
}

//同步云盘大小
type SyncVolumeSizeRequest struct {
	ReqConfig
	Uuid           string               `json:"uuid" bson:"uuid"` //UUID
	SyncVolumeSize SyncVolumeSizeParams `json:"syncVolumeSize" bson:"syncVolumeSize"`
	SystemTags     []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SyncVolumeSizeParams struct {
}

type SyncVolumeSizeResponse struct {
	Inventory VolumeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//扩展根云盘
type ResizeRootVolumeRequest struct {
	ReqConfig
	Uuid             string                 `json:"uuid" bson:"uuid"`                                 //UUID
	ResizeRootVolume ResizeRootVolumeParams `json:"resizeRootVolume" bson:"resizeRootVolume"`         //放size
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type ResizeRootVolumeParams struct {
	Size int64 `json:"size" bson:"size"`
}

type ResizeRootVolumeResponse struct {
	Inventory VolumeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//扩展数据云盘
type ResizeDataVolumeRequest struct {
	ReqConfig
	Uuid             string                 `json:"uuid" bson:"uuid"`                                 //UUID
	ResizeDataVolume ResizeDataVolumeParams `json:"resizeDataVolume" bson:"resizeDataVolume"`         //放size
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type ResizeDataVolumeParams struct {
	Size int64 `json:"size" bson:"size"`
}

type ResizeDataVolumeResponse struct {
	Inventory VolumeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改云盘属性
type UpdateVolumeRequest struct {
	ReqConfig
	Uuid         string            `json:"uuid" bson:"uuid"`                                 //UUID
	UpdateVolume UpdateVolumeParas `json:"updateVolume" bson:"updateVolume"`                 //放name,description
	SystemTags   []string          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string          `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type UpdateVolumeParas struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //云盘名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
}

type UpdateVolumeResponse struct {
	Inventory VolumeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云盘限速
type SetVolumeQoSRequest struct {
	ReqConfig
	Uuid       string             `json:"uuid" bson:"uuid"`                                 //UUID
	Params     SetVolumeQoSParams `json:"setVolumeQos" bson:"setVolumeQos"`                 //放volumeBandwidth,mode
	SystemTags []string           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string           `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVolumeQoSParams struct {
	VolumeBandwidth int64  `json:"volumeBandwidth" bson:"volumeBandwidth"` //云盘限速带宽
	Mode            string `json:"mode,omitempty" bson:"mode,omitempty"`   //可选total,read,write
}

type SetVolumeQoSResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云盘限速
type GetVolumeQoSRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`                                 //UUID
	ForceSync  bool     `json:"forceSync,omitempty" bson:"forceSync,omitempty"`   //是否到物理机上去同步数据
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVolumeQoSResponse struct {
	VolumeUuid                      string    `json:"volumeUuid" bson:"volumeUuid"`                                           //云盘UUID
	VolumeBandwidth                 int64     `json:"volumeBandwidth" bson:"volumeBandwidth"`                                 //云盘带宽，默认-1
	VolumeBandwidthRead             int64     `json:"volumeBandwidthRead" bson:"volumeBandwidthRead"`                         //云盘读带宽，默认-1
	VolumeBandwidthWrite            int64     `json:"volumeBandwidthWrite" bson:"volumeBandwidthWrite"`                       //云盘写带宽，默认-1
	VolumeBandwidthUpthreshold      int64     `json:"volumeBandwidthUpthreshold" bson:"volumeBandwidthUpthreshold"`           //	云盘带宽上限，默认-1
	VolumeBandwidthReadUpthreshold  int64     `json:"volumeBandwidthReadUpthreshold" bson:"volumeBandwidthReadUpthreshold"`   //云盘读带宽上限，默认-1
	VolumeBandwidthWriteUpthreshold int64     `json:"volumeBandwidthWriteUpthreshold" bson:"volumeBandwidthWriteUpthreshold"` //云盘写带宽上限，默认-1
	Error                           ErrorCode `json:"error,omitempty" bson:"error,omitempty"`                                 //错误信息
}

//取消云盘网卡限速
type DeleteVolumeQoSRequest struct {
	ReqConfig
	Uuid         string             `json:"uuid" bson:"uuid"`                                 //UUID
	SetVolumeQoS SetVolumeQoSParams `json:"setVolumeQoS" bson:"setVolumeQoS"`                 //mode
	SystemTags   []string           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string           `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DeleteVolumeQoSResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云盘可被加载的云主机列表
type GetDataVolumeAttachableVmRequest struct {
	ReqConfig
	VolumeUuid string   `json:"volumeUuid" bson:"volumeUuid"`                     //UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetDataVolumeAttachableVmResponse struct {
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []VmInstanceInventory `json:"inventories" bson:"inventories"`
}

//挂载云盘到云主机上
type AttachDataVolumeToVmRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	VolumeUuid     string   `json:"volumeUuid" bson:"volumeUuid"`                     //云盘UUID
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type AttachDataVolumeToVmResponse struct {
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VolumeInventory `json:"inventory" bson:"inventory"`
}

//从云主机上卸载云盘
type DetachDataVolumeFromVmRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`                                 //云盘的UUID，唯一标示该资源
	VmUuid     string   `json:"vmUuid,omitempty" bson:"vmUuid,omitempty"`         //云主机的UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DetachDataVolumeFromVmResponse struct {
	Error     ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VolumeInventory `json:"inventory" bson:"inventory"`
}

//从云盘创建快照
type CreateVolumeSnapshotRequest struct {
	ReqConfig
	VolumeUuid string                     `json:"volumeUuid" bson:"volumeUuid"` //云盘UUID
	Params     CreateVolumeSnapshotParams `json:"params" bson:"params"`
	SystemTags []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateVolumeSnapshotParams struct {
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //资源的详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID 若指定，云主机会使用该字段值作为UUID。
}

type CreateVolumeSnapshotResponse struct {
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VolumeSnapshotInventory `json:"inventory" bson:"inventory"`
}

//查询云盘快照
type QueryVolumeSnapshotRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //云盘UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type QueryVolumeSnapshotResponse struct {
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []VolumeSnapshotInventory `json:"inventories" bson:"inventories"`
}

//查询快照树
type QueryVolumeSnapshotTreeRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //云盘UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type QueryVolumeSnapshotTreeResponse struct {
	Inventories []SnapshotTree `json:"inventories" bson:"inventories"`
	Error       ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新云盘快照信息
type UpdateVolumeSnapshotRequest struct {
	ReqConfig
	Uuid                 string                     `json:"uuid" bson:"uuid"`                                 //云盘UUID
	UpdateVolumeSnapshot UpdateVolumeSnapshotParams `json:"updateVolumeSnapshot" bson:"updateVolumeSnapshot"` //放name
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type UpdateVolumeSnapshotParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
}

type UpdateVolumeSnapshotResponse struct {
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VolumeSnapshotInventory `json:"inventory" bson:"inventory"`
}

//删除云盘快照
type DeleteVolumeSnapshotRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`                                 //云盘UUID
	DeleteMode string   `json:"deleteMode" bson:"deleteMode"`                     //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DeleteVolumeSnapshotResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//将云盘回滚至指定快照
type RevertVolumeFromSnapshotRequest struct {
	ReqConfig
	Uuid                     string                         `json:"uuid" bson:"uuid"`                                         //云盘UUID
	RevertVolumeFromSnapshot RevertVolumeFromSnapshotParams `json:"revertVolumeFromSnapshot" bson:"revertVolumeFromSnapshot"` //放空
	SystemTags               []string                       `json:"systemTags,omitempty" bson:"systemTags,omitempty"`         //云主机系统标签
	UserTags                 []string                       `json:"userTags,omitempty" bson:"userTags,omitempty"`             //云主机用户标签
}

type RevertVolumeFromSnapshotParams struct {
}

type RevertVolumeFromSnapshotResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取快照容量
type GetVolumeSnapshotSizeRequest struct {
	ReqConfig
	Uuid                  string                      `json:"uuid" bson:"uuid"`                                   //云盘UUID
	GetVolumeSnapshotSize GetVolumeSnapshotSizeParams `json:"getVolumeSnapshotSize" bson:"getVolumeSnapshotSize"` //放空
	SystemTags            []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"`   //云主机系统标签
	UserTags              []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`       //云主机用户标签
}

type GetVolumeSnapshotSizeParams struct {
}

type GetVolumeSnapshotSizeResponse struct {
	Size       int64     `json:"size" bson:"size"`             //	快照容量
	ActualSize int64     `json:"actualSize" bson:"actualSize"` //快照实际容量
	Success    bool      `json:"success" bson:"success"`
	Error      ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//快照瘦身
type ShrinkVolumeSnapshotRequest struct {
	ReqConfig
	Uuid                 string                     `json:"uuid" bson:"uuid"`                                 //云盘UUID
	ShrinkVolumeSnapshot ShrinkVolumeSnapshotParams `json:"shrinkVolumeSnapshot" bson:"shrinkVolumeSnapshot"` //空
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type ShrinkVolumeSnapshotParams struct {
}

type ShrinkVolumeSnapshotResponse struct {
	ShrinkResult ShrinkResult `json:"shrinkResult" bson:"shrinkResult"`
	Error        ErrorCode    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type ShrinkResult struct {
	OldSize   int64 `json:"oldSize" bson:"oldSize"`
	Size      int64 `json:"size" bson:"size"`
	DeltaSize int64 `json:"deltaSize" bson:"deltaSize"`
}
