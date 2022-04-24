package model

//创建快照组
type CreateVolumeSnapshotGroupRequest struct {
	ReqConfig
	Params     CreateVolumeSnapshotGroupParams `json:"params" bson:"params"`
	SystemTags []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateVolumeSnapshotGroupParams struct {
	RootVolumeUuid string   `json:"rootVolumeUuid" bson:"rootVolumeUuid"`
	Name           string   `json:"name" bson:"name"`                                     //资源名称
	Description    string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid   string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids       []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreateVolumeSnapshotGroupResponse struct {
	Error     ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VolumeSnapshotGroupInventory `json:"inventory" bson:"inventory"`
}

//删除快照组
type DeleteVolumeSnapshotGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteVolumeSnapshotGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新快照组信息
type UpdateVolumeSnapshotGroupRequest struct {
	ReqConfig
	UUID                      string                          `json:"uuid" bson:"uuid"`                                           //资源的UUID，唯一标示该资源
	UpdateVolumeSnapshotGroup UpdateVolumeSnapshotGroupParams `json:"updateVolumeSnapshotGroup" bson:"updateVolumeSnapshotGroup"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags                []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags                  []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateVolumeSnapshotGroupParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateVolumeSnapshotGroupResponse struct {
	Error     ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VolumeSnapshotGroupInventory `json:"inventory" bson:"inventory"`
}

//查询快照组
type QueryVolumeSnapshotGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"` //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVolumeSnapshotGroupResponse struct {
	Error       ErrorCode                      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []VolumeSnapshotGroupInventory `json:"inventories" bson:"inventories"`
}

//检查快照组可用性
type CheckVolumeSnapshotGroupAvailabilityRequest struct {
	ReqConfig
	Uuids      []string `json:"uuids" bson:"uuids"` //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CheckVolumeSnapshotGroupAvailabilityResponse struct {
	Error   ErrorCode                                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Results []CheckVolumeSnapshotGroupAvailabilityResult `json:"results" bson:"results"`
}

type CheckVolumeSnapshotGroupAvailabilityResult struct {
	UUID      string `json:"uuid" bson:"uuid"`           //资源的UUID，唯一标示该资源
	Available bool   `json:"available" bson:"available"` //是否可以恢复
	Reason    string `json:"reason" bson:"reason"`       //不可恢复的理由，如可恢复则为空
}

//恢复快照组
type RevertVmFromSnapshotGroupRequest struct {
	ReqConfig
	Uuid                      string                          `json:"uuid" bson:"uuid"`                                                               //资源的UUID，唯一标示该资源
	RevertVmFromSnapshotGroup RevertVmFromSnapshotGroupParams `json:"revertVmFromSnapshotGroup,omitempty" bson:"revertVmFromSnapshotGroup,omitempty"` //空
	SystemTags                []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags                  []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RevertVmFromSnapshotGroupParams struct {
}

type RevertVmFromSnapshotGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//解绑快照组
type UngroupVolumeSnapshotGroupRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UngroupVolumeSnapshotGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type VolumeSnapshotGroupInventory struct {
	UUID               string               `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	SnapshotCount      int                  `json:"snapshotCount" bson:"snapshotCount"`                 //组内快照数量
	Name               string               `json:"name" bson:"name"`                                   //资源名称
	Description        string               `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	VmInstanceUuid     string               `json:"vmInstanceUuid" bson:"vmInstanceUuid"`               //云主机UUID
	CreateDate         string               `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate         string               `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
	VolumeSnapshotRefs []VolumeSnapshotRefs `json:"volumeSnapshotRefs" bson:"volumeSnapshotRefs"`
}

type VolumeSnapshotRefs struct {
	VolumeSnapshotUuid        string `json:"volumeSnapshotUuid" bson:"volumeSnapshotUuid"`               //云盘快照UUID
	VolumeSnapshotGroupUuid   string `json:"volumeSnapshotGroupUuid" bson:"volumeSnapshotGroupUuid"`     //快照组UUID
	DeviceId                  int    `json:"deviceId" bson:"deviceId"`                                   //打快照时云盘的加载序号
	SnapshotDeleted           bool   `json:"snapshotDeleted" bson:"snapshotDeleted"`                     //快照是否已经被删除
	VolumeUuid                string `json:"volumeUuid" bson:"volumeUuid"`                               //云盘UUID
	VolumeName                string `json:"volumeName" bson:"volumeName"`                               //云盘的名字
	VolumeType                string `json:"volumeType" bson:"volumeType"`                               //	云盘的类型
	VolumeSnapshotInstallPath string `json:"volumeSnapshotInstallPath" bson:"volumeSnapshotInstallPath"` //快照的安装路径
	VolumeSnapshotName        string `json:"volumeSnapshotName" bson:"volumeSnapshotName"`               //快照的名字
	CreateDate                string `json:"createDate" bson:"createDate"`                               //创建时间
	LastOpDate                string `json:"lastOpDate" bson:"lastOpDate"`                               //最后一次修改时间
}
