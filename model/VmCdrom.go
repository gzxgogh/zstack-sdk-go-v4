package model

type VmCdrom struct {
	VMInstanceUUID string `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //	云主机UUID
	DeviceId       int    `json:"deviceId" bson:"deviceId"`
	IsoUuid        string `json:"isoUuid" bson:"isoUuid"` //光盘镜像uuid
	IsoInstallPath string `json:"isoInstallPath" bson:"isoInstallPath"`
	Name           string `json:"name" bson:"name"`
	Description    string `json:"description" bson:"description"`
	CreateDate     string `json:"createDate" bson:"createDate"`
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"`
	AccountUuid    string `json:"accountUuid" bson:"accountUuid"`
}
