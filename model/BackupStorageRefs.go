package model

type BackupStorageRefs struct {
	ImageUUID         string `json:"imageUuid" bson:"imageUuid"`                 //镜像UUID 云主机的根云盘会从该字段指定的镜像创建。
	BackupStorageUuid string `json:"backupStorageUuid" bson:"backupStorageUuid"` //镜像存储UUID
	InstallPath       string `json:"installPath" bson:"installPath"`
	ExportUrl         string `json:"exportUrl" bson:"exportUrl"`
	ExportMd5Sum      string `json:"exportMd5Sum" bson:"exportMd5Sum"`
	Status            string `json:"status" bson:"status"`
	CreateDate        string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate        string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}
