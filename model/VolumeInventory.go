package model

type VolumeInventory struct {
	UUID               string `json:"uuid" bson:"uuid"`                                                 //资源的UUID，唯一标示该资源
	Name               string `json:"name" bson:"name"`                                                 //资源名称
	Description        string `json:"description,omitempty" bson:"description,omitempty"`               //资源的详细描述
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty" bson:"primaryStorageUuid,omitempty"` //主存储UUID
	VmInstanceUuid     string `json:"vmInstanceUuid" bson:"vmInstanceUuid"`                             //云主机UUID
	DiskOfferingUuid   string `json:"diskOfferingUuid" bson:"diskOfferingUuid"`                         //云盘规格UUID
	RootImageUuid      string `json:"rootImageUuid" bson:"rootImageUuid"`
	InstallPath        string `json:"installPath" bson:"installPath"`
	Type               string `json:"type" bson:"type"`
	Format             string `json:"format" bson:"format"`
	Size               int64  `json:"size" bson:"size"`
	ActualSize         int64  `json:"actualSize" bson:"actualSize"`
	DeviceId           int64  `json:"deviceId" bson:"deviceId"`
	State              string `json:"state,omitempty" bson:"state,omitempty"`
	Status             string `json:"status" bson:"status"`
	CreateDate         string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate         string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	IsShareable        bool   `json:"isShareable" bson:"isShareable"`
}

type VolumeSnapshotInventory struct {
	UUID                      string                  `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name                      string                  `json:"name" bson:"name"`                                   //资源名称
	Description               string                  `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
	Type                      string                  `json:"type" bson:"type"`
	VolumeUuid                string                  `json:"volumeUuid" bson:"volumeUuid"` //云盘UUID
	TreeUuid                  string                  `json:"treeUuid" bson:"treeUuid"`
	ParentUuid                string                  `json:"parentUuid" bson:"parentUuid"`
	PrimaryStorageUuid        string                  `json:"primaryStorageUuid" bson:"primaryStorageUuid"` //主存储UUID
	PrimaryStorageInstallPath string                  `json:"primaryStorageInstallPath" bson:"primaryStorageInstallPath"`
	VolumeType                string                  `json:"volumeType" bson:"volumeType"`
	Format                    string                  `json:"format" bson:"format"`
	Latest                    bool                    `json:"latest" bson:"latest"`
	Size                      int64                   `json:"size" bson:"size"`
	State                     string                  `json:"state,omitempty" bson:"state,omitempty"`
	Status                    string                  `json:"status" bson:"status"`
	CreateDate                string                  `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate                string                  `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	BackupStorageRefs         volumeBackupStorageRefs `json:"backupStorageRefs" bson:"backupStorageRefs"`
}

type volumeBackupStorageRefs struct {
	VolumeSnapshotUuid string `json:"volumeBackupStorageRefs" bson:"volumeBackupStorageRefs"` //云盘快照UUID
	BackupStorageUuid  string `json:"backupStorageUuid" bson:"backupStorageUuid"`             //镜像存储UUID
	InstallPath        string `json:"installPath" bson:"installPath"`
}

type SnapshotTree struct {
	UUID       string                  `json:"uuid" bson:"uuid"`             //资源的UUID，唯一标示该资源
	VolumeUuid string                  `json:"volumeUuid" bson:"volumeUuid"` //云盘UUID
	Current    bool                    `json:"current" bson:"current"`
	State      string                  `json:"state,omitempty" bson:"state,omitempty"`
	Tree       VolumeSnapshotInventory `json:"tree" bson:"tree"`
	CreateDate string                  `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate string                  `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}
