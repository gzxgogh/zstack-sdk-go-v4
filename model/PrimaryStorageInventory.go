package model

type PrimaryStorageInventory struct {
	UUID                      string   `json:"uuid" bson:"uuid"`         //资源的UUID，唯一标示该资源
	ZoneUuid                  string   `json:"zoneUuid" bson:"zoneUuid"` //区域UUID 若指定，云主机会在指定区域创建。
	Name                      string   `json:"name" bson:"name"`         //资源名称
	Url                       string   `json:"url" bson:"url"`
	Description               string   `json:"description" bson:"description"` //资源的详细描述
	TotalCapacity             int64    `json:"totalCapacity" bson:"totalCapacity"`
	AvailableCapacity         int64    `json:"availableCapacity" bson:"availableCapacity"`
	TotalPhysicalCapacity     int64    `json:"totalPhysicalCapacity" bson:"totalPhysicalCapacity"`
	AvailablePhysicalCapacity int64    `json:"availablePhysicalCapacity" bson:"availablePhysicalCapacity"`
	SystemUsedCapacity        int64    `json:"systemUsedCapacity" bson:"systemUsedCapacity"`
	Type                      string   `json:"type" bson:"type"` //云主机类型 保留字段，无需指定。UserVm/ApplianceVm
	State                     string   `json:"state" bson:"state"`
	Status                    string   `json:"status" bson:"status"`
	MountPath                 string   `json:"mountPath" bson:"mountPath"`
	CreateDate                string   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate                string   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	AttachedClusterUuids      []string `json:"attachedClusterUuids" bson:"attachedClusterUuids"`
}

type LocalStorageResourceRefInventory struct {
	ResourceUuid              string `json:"resourceUuid" bson:"resourceUuid"`
	HostUuid                  string `json:"hostUuid" bson:"hostUuid"` //	物理机UUID
	TotalCapacity             int64  `json:"totalCapacity" bson:"totalCapacity"`
	AvailableCapacity         int64  `json:"availableCapacity" bson:"availableCapacity"`
	TotalPhysicalCapacity     int64  `json:"totalPhysicalCapacity" bson:"totalPhysicalCapacity"`
	AvailablePhysicalCapacity int64  `json:"availablePhysicalCapacity" bson:"availablePhysicalCapacity"`
	CreateDate                string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate                string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type CephPrimaryStorageInventory struct {
	Fsid                      string   `json:"fsid" bson:"fsid"`
	RootVolumePoolName        string   `json:"rootVolumePoolName" bson:"rootVolumePoolName"`
	DataVolumePoolName        string   `json:"dataVolumePoolName" bson:"dataVolumePoolName"`
	ImageCachePoolName        string   `json:"imageCachePoolName" bson:"imageCachePoolName"`
	UUID                      string   `json:"uuid" bson:"uuid"`         //资源的UUID，唯一标示该资源
	ZoneUuid                  string   `json:"zoneUuid" bson:"zoneUuid"` //区域UUID 若指定，云主机会在指定区域创建。
	Name                      string   `json:"name" bson:"name"`         //资源名称
	Url                       string   `json:"url" bson:"url"`
	Description               string   `json:"description" bson:"description"` //资源的详细描述
	TotalCapacity             int64    `json:"totalCapacity" bson:"totalCapacity"`
	AvailableCapacity         int64    `json:"availableCapacity" bson:"availableCapacity"`
	TotalPhysicalCapacity     int64    `json:"totalPhysicalCapacity" bson:"totalPhysicalCapacity"`
	AvailablePhysicalCapacity int64    `json:"availablePhysicalCapacity" bson:"availablePhysicalCapacity"`
	SystemUsedCapacity        int64    `json:"systemUsedCapacity" bson:"systemUsedCapacity"`
	Type                      string   `json:"type,omitempty" bson:"type,omitempty"` //云主机类型 保留字段，无需指定。UserVm/ApplianceVm
	State                     string   `json:"state" bson:"state"`
	Status                    string   `json:"status" bson:"status"`
	MountPath                 string   `json:"mountPath" bson:"mountPath"`
	CreateDate                string   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate                string   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	AttachedClusterUuids      []string `json:"attachedClusterUuids" bson:"attachedClusterUuids"`
	Mons                      Mons     `json:"mons" bson:"mons"`
}

type Mons struct {
	Hostname           string `json:"hostname" bson:"hostname"`                     //mon节点新主机IP地址
	MonPort            int    `json:"monPort" bson:"monPort"`                       //mon的新端口
	CreateDate         string `json:"createDate" bson:"createDate"`                 //创建时间
	LastOpDate         string `json:"lastOpDate" bson:"lastOpDate"`                 //最后一次修改时间
	PrimaryStorageUuid string `json:"primaryStorageUuid" bson:"primaryStorageUuid"` //主存储UUID
	MonAddr            string `json:"monAddr" bson:"monAddr"`
	SshUsername        string `json:"sshUsername" bson:"sshUsername"` //mon节点主机 ssh 用户名
	SshPassword        string `json:"sshPassword" bson:"sshPassword"` //mon节点主机 ssh 用户密码
	SshPort            string `json:"sshPort" bson:"sshPort"`         //mon节点主机 ssh 端口
	Status             string `json:"status" bson:"status"`
	MonUuid            string `json:"monUuid" bson:"monUuid"`
}

type CephPrimaryStoragePoolInventory struct {
	UUID               string `json:"uuid" bson:"uuid"`                             //资源的UUID，唯一标示该资源
	PrimaryStorageUuid string `json:"primaryStorageUuid" bson:"primaryStorageUuid"` //主存储UUID
	PoolName           string `json:"poolName" bson:"poolName"`
	AliasName          string `json:"aliasName,omitempty" bson:"aliasName,omitempty"`
	Description        string `json:"description" bson:"description"`       //资源的详细描述
	CreateDate         string `json:"createDate" bson:"createDate"`         //创建时间
	LastOpDate         string `json:"lastOpDate" bson:"lastOpDate"`         //最后一次修改时间
	Type               string `json:"type,omitempty" bson:"type,omitempty"` //云主机类型 保留字段，无需指定。UserVm/ApplianceVm
	AvailableCapacity  int64  `json:"availableCapacity" bson:"availableCapacity"`
	UsedCapacity       int64  `json:"usedCapacity" bson:"usedCapacity"`
	ReplicatedSize     int    `json:"replicatedSize" bson:"replicatedSize"`
}

type SharedBlockGroupPrimaryStorageHostRefInventory struct {
	PrimaryStorageUuid string                   `json:"primaryStorageUuid" bson:"primaryStorageUuid"` //主存储UUID
	HostUuid           string                   `json:"hostUuid" bson:"hostUuid"`                     //	物理机UUID
	HostId             string                   `json:"hostId" bson:"hostId"`
	CreateDate         string                   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate         string                   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	Status             PrimaryStorageHostStatus `json:"status" bson:"status"`
}

type PrimaryStorageHostStatus struct {
	Name    string `json:"name" bson:"name"` //资源名称
	Ordinal int    `json:"ordinal" bson:"ordinal"`
}

type SharedBlockGroupPrimaryStorageInventory struct {
	UUID                      string               `json:"uuid" bson:"uuid"`         //资源的UUID，唯一标示该资源
	ZoneUuid                  string               `json:"zoneUuid" bson:"zoneUuid"` //区域UUID 若指定，云主机会在指定区域创建。
	Name                      string               `json:"name" bson:"name"`         //资源名称
	Url                       string               `json:"url" bson:"url"`
	Description               string               `json:"description" bson:"description"` //资源的详细描述
	TotalCapacity             int64                `json:"totalCapacity" bson:"totalCapacity"`
	AvailableCapacity         int64                `json:"availableCapacity" bson:"availableCapacity"`
	TotalPhysicalCapacity     int64                `json:"totalPhysicalCapacity" bson:"totalPhysicalCapacity"`
	AvailablePhysicalCapacity int64                `json:"availablePhysicalCapacity" bson:"availablePhysicalCapacity"`
	SystemUsedCapacity        int64                `json:"systemUsedCapacity" bson:"systemUsedCapacity"`
	Type                      string               `json:"type" bson:"type"` //云主机类型 保留字段，无需指定。UserVm/ApplianceVm
	State                     string               `json:"state" bson:"state"`
	Status                    string               `json:"status" bson:"status"`
	MountPath                 string               `json:"mountPath" bson:"mountPath"`
	CreateDate                string               `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate                string               `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	AttachedClusterUuids      []string             `json:"attachedClusterUuids" bson:"attachedClusterUuids"`
	SharedBlocks              SharedBlocks         `json:"sharedBlocks" bson:"sharedBlocks"`
	sharedBlockGroupType      SharedBlockGroupType `json:"sharedBlockGroupType" bson:"sharedBlockGroupType"`
}

type SharedBlocks struct {
	UUID                 string            `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	SharedBlockGroupUuid string            `json:"sharedBlockGroupUuid" bson:"sharedBlockGroupUuid"`
	DiskUuid             string            `json:"diskUuid" bson:"diskUuid"`
	Name                 string            `json:"name" bson:"name"`               //资源名称
	Description          string            `json:"description" bson:"description"` //资源的详细描述
	CreateDate           string            `json:"createDate" bson:"createDate"`   //创建时间
	LastOpDate           string            `json:"lastOpDate" bson:"lastOpDate"`   //最后一次修改时间
	Type                 SharedBlockType   `json:"type" bson:"type"`
	State                SharedBlockState  `json:"state" bson:"state"`
	Status               SharedBlockStatus `json:"status" bson:"status"`
}

type SharedBlockType struct {
	Name    string `json:"name" bson:"name"`
	Ordinal int    `json:"ordinal" bson:"ordinal"`
}

type SharedBlockState struct {
	Name    string `json:"name" bson:"name"`
	Ordinal int    `json:"ordinal" bson:"ordinal"`
}

type SharedBlockStatus struct {
	Name    string `json:"name" bson:"name"`
	Ordinal int    `json:"ordinal" bson:"ordinal"`
}

type SharedBlockGroupType struct {
	Name    string `json:"name" bson:"name"`
	Ordinal int    `json:"ordinal" bson:"ordinal"`
}
