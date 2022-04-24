package model

type RootVolumePrimaryStorages struct {
	UUID                      string   `json:"uuid" bson:"uuid"`                             //资源的UUID，唯一标示该资源
	ZoneUuid                  string   `json:"zoneUuid,omitempty" bson:"zoneUuid,omitempty"` //区域UUID 若指定，云主机会在指定区域创建。
	Name                      string   `json:"name" bson:"name"`                             //资源名称
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
}
