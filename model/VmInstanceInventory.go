package model

type VmInstanceInventory struct {
	UUID                 string            `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	Name                 string            `json:"name" bson:"name"`                                 //资源名称
	Description          string            `json:"description" bson:"description"`                   //资源的详细描述
	ZoneUUID             string            `json:"zoneUuid" bson:"zoneUuid"`                         //区域UUID
	ClusterUUID          string            `json:"clusterUuid" bson:"clusterUuid"`                   //集群UUID
	ImageUUID            string            `json:"imageUuid" bson:"imageUuid"`                       //镜像UUID
	HostUUID             string            `json:"hostUuid" bson:"hostUuid"`                         //物理机UUID
	LastHostUUID         string            `json:"lastHostUuid" bson:"lastHostUuid"`                 //上一次运行云主机的物理机UUID
	InstanceOfferingUUID string            `json:"instanceOfferingUuid" bson:"instanceOfferingUuid"` //计算规格UUID
	RootVolumeUUID       string            `json:"rootVolumeUuid" bson:"rootVolumeUuid"`             //根云盘UUID
	Platform             string            `json:"platform" bson:"platform"`                         //云主机运行平台
	DefaultL3NetworkUUID string            `json:"defaultL3NetworkUuid" bson:"defaultL3NetworkUuid"` //默认三层网络UUID
	Type                 string            `json:"type" bson:"type"`                                 //云主机类型
	HypervisorType       string            `json:"hypervisorType" bson:"hypervisorType"`             //云主机的hypervisor类型
	MemorySize           int64             `json:"memorySize" bson:"memorySize"`                     //内存大小
	CPUNum               int               `json:"cpuNum" bson:"cpuNum"`                             //cpu数量
	CPUSpeed             int64             `json:"cpuSpeed" bson:"cpuSpeed"`                         //cpu主频
	AllocatorStrategy    string            `json:"allocatorStrategy" bson:"allocatorStrategy"`       //分配策略
	CreateDate           string            `json:"createDate" bson:"createDate"`
	LastOpDate           string            `json:"lastOpDate" bson:"lastOpDate"`
	State                string            `json:"state" bson:"state"`           //云主机的可用状态
	VMNics               []VmNicInventory  `json:"vmNics" bson:"vmNics"`         //所有网卡信息
	AllVolumes           []VolumeInventory `json:"allVolumes" bson:"allVolumes"` //所有卷
}

type VmInventories struct {
	UUID              string              `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name              string              `json:"name" bson:"name"`                                   //云主机名称
	Description       string              `json:"description,omitempty" bson:"description,omitempty"` //云主机的详细描述
	State             string              `json:"state" bson:"state"`
	Status            string              `json:"status" bson:"status"`
	Size              int64               `json:"size" bson:"size"`             //云盘虚拟容量
	ActualSize        int64               `json:"actualSize" bson:"actualSize"` //云盘实际容量
	Md5Sum            string              `json:"md5Sum" bson:"md5Sum"`
	Url               string              `json:"url" bson:"url"` //对应路径
	MediaType         string              `json:"mediaType" bson:"mediaType"`
	GuestOsType       string              `json:"guestOsType" bson:"guestOsType"`
	Type              string              `json:"type,omitempty" bson:"type,omitempty"`         //云主机类型 保留字段，无需指定。UserVm/ApplianceVm
	Platform          string              `json:"platform,omitempty" bson:"platform,omitempty"` //云盘系统平台
	Format            string              `json:"format" bson:"format"`
	System            string              `json:"system" bson:"system"`
	CreateDate        string              `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate        string              `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	BackupStorageRefs []BackupStorageRefs `json:"backupStorageRefs" bson:"backupStorageRefs"`
}

type CloneVmInstanceResults struct {
	NumberOfClonedVm int                   `json:"numberOfClonedVm" bson:"numberOfClonedVm"`
	Inventories      []VmInstanceInventory `json:"inventories" bson:"inventories"`
}

type VmCdRomInventory struct {
	UUID           string `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	VmInstanceUuid string `json:"vmInstanceUuid" bson:"vmInstanceUuid"`               //云主机UUID
	DeviceId       int    `json:"deviceId" bson:"deviceId"`                           //光驱顺序号
	IsoUuid        string `json:"isoUuid" bson:"isoUuid"`                             //ISO镜像UUID
	IsoInstallPath string `json:"isoInstallPath" bson:"isoInstallPath"`               //ISO镜像挂载路径
	Name           string `json:"name" bson:"name"`                                   //名称
	Description    string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	CreateDate     string `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
}

type GuestToolsInventory struct {
	UUID               string `json:"uuid" bson:"uuid"`                             //资源的UUID，唯一标示该资源
	Name               string `json:"name" bson:"name"`                             //名称
	Description        string `json:"description" bson:"description"`               //详细描述
	ManagementNodeUuid string `json:"managementNodeUuid" bson:"managementNodeUuid"` //管理节点UUID
	Architecture       string `json:"architecture" bson:"architecture"`             //架构
	HypervisorType     string `json:"hypervisorType" bson:"hypervisorType"`         //虚拟化类型
	Version            string `json:"version" bson:"version"`                       //版本
	CreateDate         string `json:"createDate" bson:"createDate"`                 //创建时间
	LastOpDate         string `json:"lastOpDate" bson:"lastOpDate"`                 //最后一次修改时间
}

type VmSchedulerInventory struct {
	Uuid               string `json:"uuid" bson:"uuid"`
	TargetResourceUuid string `json:"targetResourceUuid" bson:"targetResourceUuid"`
	SchedulerName      string `json:"schedulerName" bson:"schedulerName"`
	SchedulerJob       string `json:"schedulerJob" bson:"schedulerJob"`
	SchedulerType      string `json:"schedulerType" bson:"schedulerType"`
	SchedulerInterval  int    `json:"schedulerInterval" bson:"schedulerInterval"`
	RepeatCount        int    `json:"repeatCount" bson:"repeatCount"`
	CronScheduler      string `json:"cronScheduler" bson:"cronScheduler"`
	StartTime          int    `json:"startTime" bson:"startTime"`
	StopTime           int    `json:"stopTime" bson:"stopTime"`
	CreateDate         string `json:"createDate" bson:"createDate"`
	LastOpDate         string `json:"lastOpDate" bson:"lastOpDate"`
	State              string `json:"state" bson:"state"`
}
