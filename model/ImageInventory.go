package model

type ImageInventory struct {
	UUID              string                   `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name              string                   `json:"name" bson:"name"`                                   //资源名称
	Description       string                   `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
	State             string                   `json:"state,omitempty" bson:"state,omitempty"`
	Status            string                   `json:"status" bson:"status"`
	Size              int64                    `json:"size" bson:"size"`
	ActualSize        int64                    `json:"actualSize" bson:"actualSize"`
	Md5Sum            string                   `json:"md5Sum" bson:"md5Sum"`
	Url               string                   `json:"url" bson:"url"` //被添加镜像的URL地址
	MediaType         string                   `json:"mediaType" bson:"mediaType"`
	GuestOsType       string                   `json:"guestOsType,omitempty" bson:"guestOsType,omitempty"` //镜像对应客户机操作系统的类型
	Type              string                   `json:"type" bson:"type"`
	Platform          string                   `json:"platform,omitempty" bson:"platform,omitempty"` //镜像的系统平台,Linux,Windows,WindowsVirtio,Other,Paravirtualization
	Format            string                   `json:"format" bson:"format"`
	System            string                   `json:"system,omitempty" bson:"system,omitempty"` //是否系统镜像（如，云路由镜像）
	CreateDate        string                   `json:"createDate" bson:"createDate"`             //创建时间
	LastOpDate        string                   `json:"lastOpDate" bson:"lastOpDate"`             //最后一次修改时间
	BackupStorageRefs []ImageBackupStorageRefs `json:"backupStorageRefs" bson:"backupStorageRefs"`
}

type ImageBackupStorageRefs struct {
	ImageUuid         string `json:"imageUuid" bson:"imageUuid"`                 //镜像UUID
	BackupStorageUuid string `json:"backupStorageUuid" bson:"backupStorageUuid"` //镜像存储UUID
	InstallPath       string `json:"installPath" bson:"installPath"`
	ExportUrl         string `json:"exportUrl" bson:"exportUrl"`
	ExportMd5Sum      string `json:"exportMd5Sum" bson:"exportMd5Sum"`
	State             string `json:"state,omitempty" bson:"state,omitempty"`
	CreateDate        string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate        string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type ExistingJobDetails struct {
	LongJobUuid    string `json:"longJobUuid" bson:"longJobUuid"`
	LongJobState   string `json:"longJobState" bson:"longJobState"`
	ImageUuid      string `json:"imageUuid" bson:"imageUuid"`
	ImageUploadUrl string `json:"imageUploadUrl" bson:"imageUploadUrl"`
	Offset         int64  `json:"offset" bson:"offset"`
}
