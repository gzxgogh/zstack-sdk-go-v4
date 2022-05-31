package model

//创建云主机
type CreateVmInstanceRequest struct {
	ReqConfig
	Params     CreateVmInstanceRequestParams `json:"params" bson:"params"`
	SystemTags []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateVmInstanceRequestParams struct {
	Name                            string   `json:"name" bson:"name"`                                                                           //云主机名称
	InstanceOfferingUUID            string   `json:"instanceOfferingUuid" bson:"instanceOfferingUuid"`                                           //计算规格UUID 指定云主机的CPU、内存等参数。
	ImageUUID                       string   `json:"imageUuid" bson:"imageUuid"`                                                                 //镜像UUID 云主机的根云盘会从该字段指定的镜像创建。
	L3NetworkUuids                  []string `json:"l3NetworkUuids" bson:"l3NetworkUuids"`                                                       //三层网络UUID列表 可指定一个或多个三层网络，云主机会在每个三层网络上创建一个网卡。
	Type                            string   `json:"type,omitempty" bson:"type,omitempty"`                                                       //云主机类型 保留字段，无需指定。UserVm/ApplianceVm
	RootDiskOfferingUuid            string   `json:"rootDiskOfferingUuid,omitempty" bson:"rootDiskOfferingUuid,omitempty"`                       //根云盘规格UUID 如果imageUuid字段指定的镜像类型是ISO，该字段必须指定以确定需要创建的根云盘大小。如果镜像类型是非ISO，该字段无需指定。
	DataDiskOfferingUuids           []string `json:"dataDiskOfferingUuids,omitempty" bson:"dataDiskOfferingUuids,omitempty"`                     //云盘规格UUID列表 可指定一个或多个云盘规格UUID（UUID可以重复）为云主机创建一个或多个数据云盘。
	ZoneUuid                        string   `json:"zoneUuid,omitempty" bson:"zoneUuid,omitempty"`                                               //区域UUID 若指定，云主机会在指定区域创建。
	ClusterUUID                     string   `json:"clusterUuid,omitempty" bson:"clusterUuid,omitempty"`                                         //集群UUID 若指定，云主机会在指定集群创建，该字段优先级高于zoneUuid。
	HostUuid                        string   `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`                                               //物理机UUID 若指定，云主机会在指定物理机创建，该字段优先级高于zoneUuid和clusterUuid。
	PrimaryStorageUuidForRootVolume string   `json:"primaryStorageUuidForRootVolume,omitempty" bson:"primaryStorageUuidForRootVolume,omitempty"` //主存储UUID 若指定，云主机的根云盘会在指定主存储创建。
	Description                     string   `json:"description,omitempty" bson:"description,omitempty"`                                         //云主机的详细描述
	DefaultL3NetworkUuid            string   `json:"defaultL3NetworkUuid,omitempty" bson:"defaultL3NetworkUuid,omitempty"`                       //默认三层网络UUID 当l3NetworkUuids指定了多个三层网络时，该字段指定提供默认路由的三层网络。若不指定，l3NetworkUuids的第一个网络被选为默认网络。
	ResourceUuid                    string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`                                       //资源UUID 若指定，云主机会使用该字段值作为UUID。
	TagUuids                        []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`                                               //标签UUID列表
	Strategy                        string   `json:"strategy,omitempty" bson:"strategy,omitempty"`                                               //云主机创建策略 创建后立刻启动InstantStart 创建后不启动CreateStopped
}

type CreateVmInstanceResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从云盘创建云主机
type CreateVmFromVolumeRequest struct {
	ReqConfig
	Params     CreateVmFromVolumeRequestParams `json:"params" bson:"params"`
	SystemTags []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateVmFromVolumeRequestParams struct {
	Name                 string   `json:"name" bson:"name"`                                                     //云主机名称
	Description          string   `json:"description,omitempty" bson:"description,omitempty"`                   //资源的详细描述
	InstanceOfferingUuid string   `bson:"instanceOfferingUuid,omitempty" bson:"instanceOfferingUuid,omitempty"` //计算规格UUID，注意：该参数与CPU数量、内存大小二选一
	CpuNum               int      `json:"cpuNum,omitempty" bson:"cpuNum,omitempty"`                             //CPU数量/内存大小，注意：该参数与instanceOfferingUuid二选一
	MemorySize           int64    `json:"memorySize,omitempty" bson:"memorySize,omitempty"`                     //CPU数量/内存大小，注意：该参数与instanceOfferingUuid二选一
	L3NetworkUuids       []string `json:"l3NetworkUuids" bson:"l3NetworkUuids"`                                 //三层网络UUID列表 可指定一个或多个三层网络，云主机会在每个三层网络上创建一个网卡。
	Type                 string   `json:"type,omitempty" bson:"type,omitempty"`                                 //云主机类型保留字段，无需指定。
	VolumeUuid           string   `json:"volumeUuid,omitempty" bson:"volumeUuid,omitempty"`                     //云盘UUID
	Platform             string   `json:"platform,omitempty" bson:"platform,omitempty"`                         //云盘系统平台
	ZoneUuid             string   `json:"zoneUuid,omitempty" bson:"zoneUuid,omitempty"`                         //区域UUID 若指定，云主机会在指定区域创建。
	ClusterUuid          string   `json:"clusterUuid,omitempty" bson:"clusterUuid,omitempty"`                   //集群UUID 若指定，云主机会在指定集群创建，该字段优先级高于zoneUuid
	HostUuid             string   `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`                         //物理机UUID 若指定，云主机会在指定物理机创建，该字段优先级高于zoneUuid和clusterUuid
	PrimaryStorageUuid   string   `json:"primaryStorageUuid,omitempty" bson:"primaryStorageUuid,omitempty"`     //主存储UUID 若指定，云主机的根云盘会在指定主存储创建。
	DefaultL3NetworkUuid string   `json:"defaultL3NetworkUuid,omitempty" bson:"defaultL3NetworkUuid,omitempty"` //默认三层网络UUID 当l3NetworkUuids指定了多个三层网络时，该字段指定提供默认路由的三层网络。若不指定，l3NetworkUuids的第一个网络被选为默认网络。
	Strategy             string   `json:"strategy,omitempty" bson:"strategy,omitempty"`                         //云主机创建策略 1.创建后立刻启动 2.创建后不启动
	ResourceUuid         string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`                 //资源UUID 若指定，云主机会使用该字段值作为UUID。
	TagUuids             []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`                         //标签UUID列表
}

type CreateVmFromVolumeResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从快照创建云主机
type CreateVmInstanceFromVolumeSnapshotRequest struct {
	ReqConfig
	VolumeSnapshotUuid string                                   `json:"volumeSnapshotUuid" bson:"volumeSnapshotUuid"`
	Params             CreateVmInstanceFromVolumeSnapshotParams `json:"params" bson:"params"`
	SystemTags         []string                                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string                                 `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateVmInstanceFromVolumeSnapshotParams struct {
	Name                 string   `json:"name" bson:"name"`                                                     //云主机名称
	Description          string   `json:"description,omitempty" bson:"description,omitempty"`                   //资源的详细描述
	InstanceOfferingUuid string   `bson:"instanceOfferingUuid,omitempty" bson:"instanceOfferingUuid,omitempty"` //计算规格UUID，注意：该参数与CPU数量、内存大小二选一
	CpuNum               int      `json:"cpuNum,omitempty" bson:"cpuNum,omitempty"`                             //CPU数量/内存大小，注意：该参数与instanceOfferingUuid二选一
	MemorySize           int64    `json:"memorySize,omitempty" bson:"memorySize,omitempty"`                     //CPU数量/内存大小，注意：该参数与instanceOfferingUuid二选一
	L3NetworkUuids       []string `json:"l3NetworkUuids" bson:"l3NetworkUuids"`                                 //三层网络UUID列表 可指定一个或多个三层网络，云主机会在每个三层网络上创建一个网卡。
	Type                 string   `json:"type,omitempty" bson:"type,omitempty"`                                 //云主机类型保留字段，无需指定。
	Platform             string   `json:"platform,omitempty" bson:"platform,omitempty"`                         //云盘系统平台
	ZoneUuid             string   `json:"zoneUuid,omitempty" bson:"zoneUuid,omitempty"`                         //区域UUID 若指定，云主机会在指定区域创建。
	ClusterUuid          string   `json:"clusterUuid,omitempty" bson:"clusterUuid,omitempty"`                   //集群UUID 若指定，云主机会在指定集群创建，该字段优先级高于zoneUuid
	HostUuid             string   `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`                         //物理机UUID 若指定，云主机会在指定物理机创建，该字段优先级高于zoneUuid和clusterUuid
	PrimaryStorageUuid   string   `json:"primaryStorageUuid,omitempty" bson:"primaryStorageUuid,omitempty"`     //主存储UUID 若指定，云主机的根云盘会在指定主存储创建。
	DefaultL3NetworkUuid string   `json:"defaultL3NetworkUuid,omitempty" bson:"defaultL3NetworkUuid,omitempty"` //默认三层网络UUID 当l3NetworkUuids指定了多个三层网络时，该字段指定提供默认路由的三层网络。若不指定，l3NetworkUuids的第一个网络被选为默认网络。
	Strategy             string   `json:"strategy,omitempty" bson:"strategy,omitempty"`                         //云主机创建策略 1.创建后立刻启动 2.创建后不启动
	ResourceUuid         string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`                 //资源UUID 若指定，云主机会使用该字段值作为UUID。
	TagUuids             []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`                         //标签UUID列表
}

type CreateVmInstanceFromVolumeSnapshotResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从快照创建云主机
type CreateVmInstanceFromVolumeSnapshotGroupRequest struct {
	ReqConfig
	VolumeSnapshotGroupUuid string                                        `json:"volumeSnapshotGroupUuid" bson:"volumeSnapshotGroupUuid"`
	Params                  CreateVmInstanceFromVolumeSnapshotGroupParams `json:"params" bson:"params"`
	SystemTags              []string                                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                []string                                      `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateVmInstanceFromVolumeSnapshotGroupParams struct {
	Name                 string   `json:"name" bson:"name"`                                                     //云主机名称
	Description          string   `json:"description,omitempty" bson:"description,omitempty"`                   //资源的详细描述
	InstanceOfferingUuid string   `bson:"instanceOfferingUuid,omitempty" bson:"instanceOfferingUuid,omitempty"` //计算规格UUID，注意：该参数与CPU数量、内存大小二选一
	CpuNum               int      `json:"cpuNum,omitempty" bson:"cpuNum,omitempty"`                             //CPU数量/内存大小，注意：该参数与instanceOfferingUuid二选一
	MemorySize           int64    `json:"memorySize,omitempty" bson:"memorySize,omitempty"`                     //CPU数量/内存大小，注意：该参数与instanceOfferingUuid二选一
	L3NetworkUuids       []string `json:"l3NetworkUuids" bson:"l3NetworkUuids"`                                 //三层网络UUID列表 可指定一个或多个三层网络，云主机会在每个三层网络上创建一个网卡。
	Type                 string   `json:"type,omitempty" bson:"type,omitempty"`                                 //云主机类型保留字段，无需指定。
	Platform             string   `json:"platform,omitempty" bson:"platform,omitempty"`                         //云盘系统平台
	ZoneUuid             string   `json:"zoneUuid,omitempty" bson:"zoneUuid,omitempty"`                         //区域UUID 若指定，云主机会在指定区域创建。
	ClusterUuid          string   `json:"clusterUuid,omitempty" bson:"clusterUuid,omitempty"`                   //集群UUID 若指定，云主机会在指定集群创建，该字段优先级高于zoneUuid
	HostUuid             string   `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`                         //物理机UUID 若指定，云主机会在指定物理机创建，该字段优先级高于zoneUuid和clusterUuid
	PrimaryStorageUuid   string   `json:"primaryStorageUuid,omitempty" bson:"primaryStorageUuid,omitempty"`     //主存储UUID 若指定，云主机的根云盘会在指定主存储创建。
	DefaultL3NetworkUuid string   `json:"defaultL3NetworkUuid,omitempty" bson:"defaultL3NetworkUuid,omitempty"` //默认三层网络UUID 当l3NetworkUuids指定了多个三层网络时，该字段指定提供默认路由的三层网络。若不指定，l3NetworkUuids的第一个网络被选为默认网络。
	Strategy             string   `json:"strategy,omitempty" bson:"strategy,omitempty"`                         //云主机创建策略 1.创建后立刻启动 2.创建后不启动
	ResourceUuid         string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`                 //资源UUID 若指定，云主机会使用该字段值作为UUID。
	TagUuids             []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`                         //标签UUID列表
}

type CreateVmInstanceFromVolumeSnapshotGroupResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除云主机
type DestroyVmInstanceRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode" bson:"deleteMode"`                     //删除方式，默认填 Permissive
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DestroyVmInstanceResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//恢复已删除云主机
type RecoverVmInstanceRequest struct {
	ReqConfig
	UUID              string                  `json:"uuid" bson:"uuid"`                                               //资源的UUID，唯一标示该资源
	RecoverVmInstance RecoverVmInstanceParams `json:"recoverVmInstance,omitempty" bson:"recoverVmInstance,omitempty"` //里面存uuid
	SystemTags        []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"`               //云主机系统标签
	UserTags          []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`                   //云主机用户标签
}

type RecoverVmInstanceParams struct {
}

type RecoverVmInstanceResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//彻底删除云主机
type ExpungeVmInstanceRequest struct {
	ReqConfig
	UUID              string                  `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ExpungeVmInstance ExpungeVmInstanceParams `json:"expungeVmInstance" bson:"expungeVmInstance"`
	SystemTags        []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type ExpungeVmInstanceParams struct {
}

type ExpungeVmInstanceResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询云主机
type QueryVmInstanceRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type QueryVmInstanceResponse struct {
	Inventories []VmInstanceInventory `json:"inventories" bson:"inventories"`         //云主机详细清单
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//通过ip查询云主机信息
type QueryVmByIpRequest struct {
	ReqConfig
	Ip         string   `json:"ip" bson:"ip"`                                     //ip
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type QueryVmByIpResponse struct {
	Inventory VmInstanceInventory `json:"inventories" bson:"inventories"`         //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//启动云主机
type StartVmInstanceRequest struct {
	ReqConfig
	UUID            string                `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	StartVmInstance StartVmInstanceParams `json:"startVmInstance" bson:"startVmInstance"`           //里面存uuid
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type StartVmInstanceParams struct { //资源的UUID，唯一标示该资源
	ClusterUuid string `json:"clusterUuid,omitempty" bson:"clusterUuid,omitempty"` //集群UUID。若指定，云主机将在该集群启动
	HostUuid    string `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`       //物理机UUID。若指定，云主机将在该物理机启动。该字段覆盖clusterUuid字段
}

type StartVmInstanceResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//停止云主机
type StopVmInstanceRequest struct {
	ReqConfig
	UUID           string               `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	StopVmInstance StopVmInstanceParams `json:"stopVmInstance" bson:"stopVmInstance"`             //需要存uuid和 type
	SystemTags     []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type StopVmInstanceParams struct {
	Type   string `json:"type,omitempty" bson:"type,omitempty"`     //云主机类型保留字段，无需指定。
	StopHA string `json:"stopHa,omitempty" bson:"stopHa,omitempty"` //彻底关闭HA云主机
}

type StopVmInstanceResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//重启云主机
type RebootVmInstanceRequest struct {
	ReqConfig
	UUID             string                 `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	RebootVmInstance RebootVmInstanceParams `json:"rebootVmInstance" bson:"rebootVmInstance"`         //填uuid
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type RebootVmInstanceParams struct {
}

type RebootVmInstanceResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//暂停云主机
type PauseVmInstanceRequest struct {
	ReqConfig
	UUID            string                `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	PauseVmInstance PauseVmInstanceParams `json:"pauseVmInstance" bson:"pauseVmInstance"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type PauseVmInstanceParams struct {
}

type PauseVmInstanceResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//恢复暂停的云主机
type ResumeVmInstanceRequest struct {
	ReqConfig
	UUID             string                 `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	ResumeVmInstance ResumeVmInstanceParams `json:"resumeVmInstance" bson:"resumeVmInstance"`         //填uuid
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type ResumeVmInstanceParams struct {
}

type ResumeVmInstanceResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//重置云主机
type ReimageVmInstanceRequest struct {
	ReqConfig
	VmInstanceUuid    string                  `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //云主机UUID
	ReimageVmInstance ReimageVmInstanceParams `json:"reimageVmInstance" bson:"reimageVmInstance"`
	SystemTags        []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type ReimageVmInstanceParams struct {
}

type ReimageVmInstanceResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//热迁移云主机
type MigrateVmRequest struct {
	ReqConfig
	VmInstanceUuid string          `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	MigrateVm      MigrateVmParams `json:"migrateVm" bson:"migrateVm"`                       //放HostUuid和MigrateFromDestination
	SystemTags     []string        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string        `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type MigrateVmParams struct {
	HostUuid               string `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`                             //物理机UUID		                                    //物理机UUID 若指定，云主机会在指定物理机创建，该字段优先级高于zoneUuid和clusterUuid。
	MigrateFromDestination bool   `json:"migrateFromDestination,omitempty" bson:"migrateFromDestination,omitempty"` //从迁移目的物理机器发起迁移命令
}

type MigrateVmResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询共享磁盘所挂载的云主机
type QueryShareableVolumeVmInstanceRefRequest struct {
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type QueryShareableVolumeVmInstanceRefResponse struct {
	Inventories []VmInstanceInventory `json:"inventories" bson:"inventories"` //云主机详细清单
	Schema      interface{}           `json:"schema" bson:"schema"`
}

//获取可热迁移的物理机列表
type GetVmMigrationCandidateHostsRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmMigrationCandidateHostsResponse struct {
	Inventories []HostInventory `json:"inventories" bson:"inventories"`         //云主机详细清单
	Error       ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取可选择的主存储
type GetCandidatePrimaryStoragesForCreatingVmRequest struct {
	ReqConfig
	ImageUuid             string   `json:"imageUuid" bson:"imageUuid"`                                             //镜像UUID
	L3NetworkUuids        []string `json:"l3NetworkUuids" bson:"l3NetworkUuids"`                                   //三层网络UUID列表 可指定一个或多个三层网络，云主机会在每个三层网络上创建一个网卡。
	RootDiskOfferingUuid  string   `json:"rootDiskOfferingUuid,omitempty" bson:"rootDiskOfferingUuid,omitempty"`   //根云盘规格UUID 如果imageUuid字段指定的镜像类型是ISO，该字段必须指定以确定需要创建的根云盘大小。如果镜像类型是非ISO，该字段无需指定。
	RootDiskSize          int64    `json:"rootDiskSize" bson:"rootDiskSize"`                                       //自定义根云盘大小。仅在imageUuid指定的镜像是ISO时且
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty" bson:"dataDiskOfferingUuids,omitempty"` //云盘规格UUID列表 可指定一个或多个云盘规格UUID（UUID可以重复）为云主机创建一个或多个数据云盘。
	DataDiskSizes         int64    `json:"dataDiskSizes" bson:"dataDiskSizes"`                                     //自定义数据云盘大小
	ZoneUuid              string   `json:"zoneUuid,omitempty" bson:"zoneUuid,omitempty"`                           //区域UUID 若指定，云主机会在指定区域创建。
	ClusterUUID           string   `json:"clusterUuid,omitempty" bson:"clusterUuid,omitempty"`                     //集群UUID 若指定，云主机会在指定集群创建，该字段优先级高于zoneUuid。
	DefaultL3NetworkUuid  string   `json:"defaultL3NetworkUuid,omitempty" bson:"defaultL3NetworkUuid,omitempty"`   //默认三层网络UUID 当l3NetworkUuids指定了多个三层网络时，该字段指定提供默认路由的三层网络。若不指定，l3NetworkUuids的第一个网络被选为默认网络。
	SystemTags            []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`                       //云主机系统标签
	UserTags              []string `json:"userTags,omitempty" bson:"userTags,omitempty"`                           //云主机用户标签
}

type GetCandidatePrimaryStoragesForCreatingVmResponse struct {
	DataVolumePrimaryStorages map[string]RootVolumePrimaryStorages `json:"dataVolumePrimaryStorages" bson:"dataVolumePrimaryStorages"` //一个键值对，值为RootVolumePrimaryStorages
	Success                   bool                                 `json:"success" bson:"success"`                                     //操作是否成功
	RootVolumePrimaryStorages RootVolumePrimaryStorages            `json:"rootVolumePrimaryStorages" bson:"rootVolumePrimaryStorages"` //
	Error                     ErrorCode                            `json:"error,omitempty" bson:"error,omitempty"`                     //错误信息
}

//获取云主机可加载ISO列表
type GetCandidateIsoForAttachingVmRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetCandidateIsoForAttachingVmResponse struct {
	Inventories []VmInventories `json:"inventories" bson:"inventories"`         //
	Error       ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取ISO可加载云主机列表
type GetCandidateVmForAttachingIsoRequest struct {
	ReqConfig
	IsoUuid    string   `json:"isoUuid" bson:"isoUuid"`                           //ISO UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetCandidateVmForAttachingIsoResponse struct {
	Inventories []VmInstanceInventory `json:"inventories" bson:"inventories"`         //
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//加载ISO到云主机
type AttachIsoToVmInstanceRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	IsoUuid        string   `json:"isoUuid" bson:"isoUuid"`                           //ISO UUID
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type AttachIsoToVmInstanceResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//卸载云主机上的ISO
type DetachIsoFromVmInstanceRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	IsoUuid        string   `json:"isoUuid,omitempty" bson:"isoUuid,omitempty"`       //ISO UUID
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DetachIsoFromVmInstanceResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机可加载云盘列表
type GetVmAttachableDataVolumeRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmAttachableDataVolumeResponse struct {
	Inventories []VolumeInventory `json:"inventories" bson:"inventories"` //
	Success     bool              `json:"success" bson:"success"`         //结果信息
}

//获取云主机可加载L3网络列表
type GetVmAttachableL3NetworkRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmAttachableL3NetworkResponse struct {
	Inventories []L3NetworkInventory `json:"inventories" bson:"inventories"`         //
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//加载L3网络到云主机
type AttachL3NetworkToVmRequest struct {
	ReqConfig
	VmInstanceUuid      string              `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //云主机UUID
	L3NetworkUuid       string              `json:"l3NetworkUuid" bson:"l3NetworkUuid"`   //三层网络UUID
	AttachL3NetworkToVm AttachL3NetworkToVm `json:"attachL3NetworkToVm" bson:"attachL3NetworkToVm"`
	SystemTags          []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type AttachL3NetworkToVm struct {
	DriverType string `json:"driverType" bson:"driverType"`                 //网卡驱动类型,vcenter侧的VPC可选的网卡驱动:1.E1000E,2.E1000,3.Vmxnet3,4.Sriov
	StaticIp   string `json:"staticIp,omitempty" bson:"staticIp,omitempty"` //指定分配给云主机的IP地址
}

type AttachL3NetworkToVmResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从云主机卸载网络
type DetachL3NetworkFromVmRequest struct {
	ReqConfig
	VmNicUuid  string   `json:"vmNicUuid" bson:"vmNicUuid"`                       //云主机网卡UUID，该网卡所在网络会从云主机卸载掉
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DetachL3NetworkFromVmResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建云主机网卡
type CreateVmNicRequest struct {
	ReqConfig
	Params       CreateVmNicParams `json:"params" bson:"params"`
	ResourceUuid string            `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //云主机网卡UUID，该网卡所在网络会从云主机卸载掉
	SystemTags   []string          `json:"systemTags,omitempty" bson:"systemTags,omitempty"`     //云主机系统标签
	UserTags     []string          `json:"userTags,omitempty" bson:"userTags,omitempty"`         //云主机用户标签
}

type CreateVmNicParams struct {
	L3NetworkUUID string `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //	三层网络UUID
	Ip            string `json:"ip" bson:"ip"`
	ResourceUuid  string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateVmNicResponse struct {
	Inventory VmNicInventory `json:"inventory" bson:"inventory"`             //
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//加载网卡到云主机
type AttachVmNicToVmRequest struct {
	ReqConfig
	VmNicUuid      string   `json:"vmNicUuid" bson:"vmNicUuid"`                       //云主机网卡UUID，该网卡所在网络会从云主机卸载掉
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type AttachVmNicToVmResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`             //
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除云主机网卡
type DeleteVmNicRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"` //	资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DeleteVmNicResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询云主机网卡
type QueryVmNicRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type QueryVmNicResponse struct {
	Inventories []VmNicInventory `json:"inventories" bson:"inventories"`         //
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取网卡加载的网络服务名称
type GetVmNicAttachedNetworkServiceRequest struct {
	ReqConfig
	VmNicUuid  string   `json:"vmNicUuid" bson:"vmNicUuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmNicAttachedNetworkServiceResponse struct {
	NetworkServices []string  `json:"networkServices" bson:"networkServices"`
	Error           ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改云主机网卡三层网络
type ChangeVmNicNetworkRequest struct {
	ReqConfig
	VmNicUuid         string `json:"vmNicUuid" bson:"vmNicUuid"`
	DestL3NetworkUuid string `json:"destL3NetworkUuid,omitempty" bson:"destL3NetworkUuid,omitempty"`

	Params     ChangeVmNicNetworkParam `json:"params" bson:"params"`
	SystemTags []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type ChangeVmNicNetworkParam struct {
}

type ChangeVmNicNetworkResponse struct {
	NetworkServices []string  `json:"networkServices" bson:"networkServices"`
	Error           ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机网卡可挂载的三层网络
type GetCandidateL3NetworksForChangeVmNicNetworkRequest struct {
	ReqConfig
	VmNicUuid  string   `json:"vmNicUuid" bson:"vmNicUuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetCandidateL3NetworksForChangeVmNicNetworkResponse struct {
	Inventories []L3NetworkInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机网卡限速
type SetNicQoSRequest struct {
	ReqConfig
	Uuid       string          `json:"uuid" bson:"uuid"`
	SetNicQos  SetNicQoSParams `json:"setNicQos" bson:"setNicQos"`                       //存出流量带宽限制，入流量带宽限制
	SystemTags []string        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string        `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetNicQoSParams struct {
	OutboundBandwidth string `json:"outboundBandwidth" bson:"outboundBandwidth"` //出流量带宽限制
	InboundBandwidth  string `json:"inboundBandwidth" bson:"inboundBandwidth"`   //入流量带宽限制
}

type SetNicQoSResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机网卡限速
type GetNicQoSRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetNicQoSResponse struct {
	OutboundBandwidth            int64     `json:"outboundBandwidth" bson:"outboundBandwidth"` //出流量带宽限制
	InboundBandwidth             int64     `json:"inboundBandwidth" bson:"inboundBandwidth"`   //入流量带宽限制
	OutboundBandwidthUpthreshold int64     `json:"outboundBandwidthUpthreshold" bson:"outboundBandwidthUpthreshold"`
	InboundBandwidthUpthreshold  int64     `json:"inboundBandwidthUpthreshold" bson:"inboundBandwidthUpthreshold"`
	Success                      bool      `json:"success" bson:"success"`
	Error                        ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//取消云主机网卡限速
type DeleteNicQoSRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	Direction  string   `json:"direction" bson:"direction"`                       //入方向还是出方向(in or out)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DeleteNicQoSResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取相互依赖的镜像和L3网络
type GetInterdependentL3NetworksImagesRequest struct {
	ReqConfig
	ZoneUuid       string   `json:"zoneUuid,omitempty" bson:"zoneUuid,omitempty"`     //区域UUID 若指定，云主机会在指定区域创建。
	L3NetworkUuids []string `json:"l3NetworkUuids" bson:"l3NetworkUuids"`             //三层网络UUID列表 可指定一个或多个三层网络，云主机会在每个三层网络上创建一个网卡。
	ImageUUID      string   `json:"imageUuid" bson:"imageUuid"`                       //镜像UUID 云主机的根云盘会从该字段指定的镜像创建。
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetInterdependentL3NetworksImagesResponse struct {
	Inventories []L3NetworkInventory `json:"inventories" bson:"inventories"`         //云主机详细清单
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机SSH Key
type SetVmSshKeyRequest struct {
	ReqConfig
	Uuid        string            `json:"uuid" bson:"uuid"`
	SetVmSshKey SetVmSshKeyParams `json:"setVmSshKey" bson:"setVmSshKey"`                   //放SshKey
	SystemTags  []string          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string          `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmSshKeyParams struct {
	SshKey string `json:"sshKey" bson:"sshKey"` //云主机SSH Key
}

type SetVmSshKeyResponse struct {
	Inventory VmInstanceInventory `json:"inventories" bson:"inventories"`         //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机SSH Key
type GetVmSshKeyRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmSshKeyResponse struct {
	SshKey string    `json:"sshKey" bson:"sshKey"`                   //云主机SSH Key
	Error  ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除云主机SSH Key
type DeleteVmSshKeyRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DeleteVmSshKeyResponse struct {
	Inventory VmInstanceInventory `json:"inventories" bson:"inventories"`         //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//变更云主机密码
type ChangeVmPasswordRequest struct {
	ReqConfig
	Uuid             string                 `json:"uuid" bson:"uuid"`
	ChangeVmPassword ChangeVmPasswordParams `json:"changeVmPassword" bson:"changeVmPassword"`         //password，account
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type ChangeVmPasswordParams struct {
	Password string `json:"password" bson:"password"` //密码
	Account  string `json:"account" bson:"account"`   //账户
}

type ChangeVmPasswordResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机控制台密码
type SetVmConsolePasswordRequest struct {
	ReqConfig
	Uuid                 string                     `json:"uuid" bson:"uuid"`
	SetVmConsolePassword SetVmConsolePasswordParams `json:"setVmConsolePassword" bson:"setVmConsolePassword"` //consolePassword
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmConsolePasswordParams struct {
	ConsolePassword string `json:"consolePassword" bson:"consolePassword"` //密码
}

type SetVmConsolePasswordResponse struct {
	Inventory VmInstanceInventory `json:"inventories" bson:"inventories"`         //云主机详细清单
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机控制台密码
type GetVmConsolePasswordRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmConsolePasswordResponse struct {
	ConsolePassword string    `json:"consolePassword" bson:"consolePassword"` //密码
	Error           ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除云主机控制台密码
type DeleteVmConsolePasswordRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DeleteVmConsolePasswordResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机控制台地址和访问协议
type GetVmConsoleAddressRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmConsoleAddressResponse struct {
	HostIp      string      `json:"hostIp" bson:"hostIp"`                   //云主机运行物理机IP
	Port        string      `json:"port" bson:"port"`                       //云主机控制台端口
	Protocol    string      `json:"protocol" bson:"protocol"`               //云主机控制台协议，vnc或spice或vncAndSpice
	Success     bool        `json:"success" bson:"success"`                 //操作是否成功
	VdiPortInfo VdiPortInfo `json:"vdiPortInfo" bson:"vdiPortInfo"`         //端口组
	Error       ErrorCode   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机Hostname
type SetVmHostnameRequest struct {
	ReqConfig
	Uuid          string              `json:"uuid" bson:"uuid"`
	SetVmHostname SetVmHostnameParams `json:"setVmHostname" bson:"setVmHostname"`               //放Hostname
	SystemTags    []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmHostnameParams struct {
	Hostname string `json:"hostname" bson:"hostname"` //hostname，必须符合RFC 1123标准
}

type SetVmHostnameResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机Hostname
type GetVmHostnameRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmHostnameResponse struct {
	Hostname string    `json:"hostname" bson:"hostname"`               //hostname，必须符合RFC 1123标准
	Error    ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除云主机Hostname
type DeleteVmHostnameRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DeleteVmHostnameResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建启动云主机的定时任务
type CreateStartVmInstanceSchedulerRequest struct {
	ReqConfig
	VmUuid     string                               `json:"vmUuid" bson:"vmUuid"` //云主机uuid
	Params     CreateStartVmInstanceSchedulerParams `json:"params" bson:"params"`
	SystemTags []string                             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                             `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateStartVmInstanceSchedulerParams struct {
	ClusterUuid          string `json:"clusterUuid,omitempty" bson:"clusterUuid,omitempty"`                   //集群UUID
	HostUuid             string `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`                         //物理机UUID
	SchedulerName        string `json:"schedulerName" bson:"schedulerName"`                                   //定时任务名称
	SchedulerDescription string `json:"schedulerDescription,omitempty" bson:"schedulerDescription,omitempty"` //定时任务描述
	Type                 string `json:"type" bson:"type"`                                                     //定时任务类型，simple或者cron
	Interval             int    `json:"interval,omitempty" bson:"interval,omitempty"`                         //定时任务间隔，单位秒
	RepeatCount          int    `json:"repeatCount,omitempty" bson:"repeatCount,omitempty"`                   //定时任务重复次数，仅针对simple类型的定时任务生效
	StartTime            int64  `json:"startTime,omitempty" bson:"startTime,omitempty"`                       //定时任务启动时间，必须遵循unix timestamp格式，0为从立刻开始
	Cron                 string `json:"cron,omitempty" bson:"cron,omitempty"`                                 //cron表达式，需遵循Java Quartz组件cron格式标准
	ResourceUuid         string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`                 //用户可指定创建Scheduler所使用的uuid
}

type CreateStartVmInstanceSchedulerResponse struct {
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VmSchedulerInventory `json:"inventory" bson:"inventory"`
}

//创建停止云主机的定时任务
type CreateStopVmInstanceSchedulerRequest struct {
	ReqConfig
	VmUuid     string                              `json:"vmUuid" bson:"vmUuid"` //云主机uuid
	Params     CreateStopVmInstanceSchedulerParams `json:"params" bson:"params"`
	SystemTags []string                            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                            `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateStopVmInstanceSchedulerParams struct {
	SchedulerName        string `json:"schedulerName" bson:"schedulerName"`                                   //定时任务名称
	SchedulerDescription string `json:"schedulerDescription,omitempty" bson:"schedulerDescription,omitempty"` //定时任务描述
	Type                 string `json:"type" bson:"type"`                                                     //定时任务类型，simple或者cron
	Interval             int    `json:"interval,omitempty" bson:"interval,omitempty"`                         //定时任务间隔，单位秒
	RepeatCount          int    `json:"repeatCount,omitempty" bson:"repeatCount,omitempty"`                   //定时任务重复次数，仅针对simple类型的定时任务生效
	StartTime            int64  `json:"startTime,omitempty" bson:"startTime,omitempty"`                       //定时任务启动时间，必须遵循unix timestamp格式，0为从立刻开始
	Cron                 string `json:"cron,omitempty" bson:"cron,omitempty"`                                 //cron表达式，需遵循Java Quartz组件cron格式标准
	ResourceUuid         string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`                 //用户可指定创建Scheduler所使用的uuid
}

type CreateStopVmInstanceSchedulerResponse struct {
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VmSchedulerInventory `json:"inventory" bson:"inventory"`
}

//创建重启云主机的定时任务
type CreateRebootVmInstanceSchedulerRequest struct {
	ReqConfig
	VmUuid     string                                `json:"vmUuid" bson:"vmUuid"` //云主机uuid
	Params     CreateRebootVmInstanceSchedulerParams `json:"params" bson:"params"`
	SystemTags []string                              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                              `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateRebootVmInstanceSchedulerParams struct {
	SchedulerName        string `json:"schedulerName" bson:"schedulerName"`                                   //定时任务名称
	SchedulerDescription string `json:"schedulerDescription,omitempty" bson:"schedulerDescription,omitempty"` //定时任务描述
	Type                 string `json:"type" bson:"type"`                                                     //定时任务类型，simple或者cron
	Interval             int    `json:"interval,omitempty" bson:"interval,omitempty"`                         //定时任务间隔，单位秒
	RepeatCount          int    `json:"repeatCount,omitempty" bson:"repeatCount,omitempty"`                   //定时任务重复次数，仅针对simple类型的定时任务生效
	StartTime            int64  `json:"startTime,omitempty" bson:"startTime,omitempty"`                       //定时任务启动时间，必须遵循unix timestamp格式，0为从立刻开始
	Cron                 string `json:"cron,omitempty" bson:"cron,omitempty"`                                 //cron表达式，需遵循Java Quartz组件cron格式标准
	ResourceUuid         string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`                 //用户可指定创建Scheduler所使用的uuid
}

type CreateRebootVmInstanceSchedulerResponse struct {
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VmSchedulerInventory `json:"inventory" bson:"inventory"`
}

//获取云主机启动设备列表
type GetVmBootOrderRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`                                 //云主机uuid
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmBootOrderResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Order []string  `json:"order" bson:"order"`
}

//指定云主机启动设备
type SetVmBootOrderRequest struct {
	ReqConfig
	Uuid           string               `json:"uuid" bson:"uuid"`                                 //云主机uuid
	SetVmBootOrder SetVmBootOrderParams `json:"setVmBootOrder" bson:"setVmBootOrder"`             //放bootOrder
	SystemTags     []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmBootOrderParams struct {
	BootOrder []string `json:"bootOrder,omitempty" bson:"bootOrder,omitempty"` //启动设备。CdRom：光驱，HardDisk：云盘，Network：网络。若该字段不指定，则表示使用系统默认启动设备顺序(HardDisk, CdRom, Network)
}

type SetVmBootOrderResponse struct {
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`
}

//获取目的地列表
type GetCandidateZonesClustersHostsForCreatingVmRequest struct {
	ReqConfig
	InstanceOfferingUuid  string   `json:"instanceOfferingUuid,omitempty" bson:"instanceOfferingUuid,omitempty"`   //计算规格UUID
	ImageUuid             string   `json:"imageUuid" bson:"imageUuid"`                                             //镜像UUID
	L3NetworkUuids        []string `json:"l3NetworkUuids" bson:"l3NetworkUuids"`                                   //三层网络列表
	RootDiskOfferingUuid  string   `json:"rootDiskOfferingUuid,omitempty" bson:"rootDiskOfferingUuid,omitempty"`   //根云盘规格。仅在imageUuid指定的镜像是ISO时需要指定
	DataDiskOfferingUuids string   `json:"dataDiskOfferingUuids,omitempty" bson:"dataDiskOfferingUuids,omitempty"` //云盘规格列表
	ZoneUuid              string   `json:"zoneUuid,omitempty" bson:"zoneUuid,omitempty"`                           //区域UUID
	ClusterUuid           string   `json:"clusterUuid ,omitempty" bson:"clusterUuid,omitempty"`                    //集群UUID
	DefaultL3NetworkUuid  string   `json:"defaultL3NetworkUuid,omitempty" bson:"defaultL3NetworkUuid ,omitempty"`  //默认三层网络UUID
	CpuNum                int      `json:"cpuNum,omitempty" bson:"cpuNum,omitempty"`                               //CPU数目
	MemorySize            int64    `json:"memorySize,omitempty" bson:"memorySize,omitempty"`                       //内存大小, 单位Byte
	SystemTags            []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`                       //云主机系统标签
	UserTags              []string `json:"userTags,omitempty" bson:"userTags,omitempty"`                           //云主机用户标签
}

type GetCandidateZonesClustersHostsForCreatingVmResponse struct {
	Error    ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Success  bool             `json:"success" bson:"success"`                 //操作是否成功
	Zones    ZoneInventory    `json:"zones" bson:"zones"`                     //区域组
	Clusters ClusterInventory `bson:"clusters" bson:"clusters"`               //集群组
	Hosts    HostInventory    `json:"hosts" bson:"hosts"`                     //主机组
}

//获取云主机可启动目的地列表
type GetVmStartingCandidateClustersHostsRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmStartingCandidateClustersHostsResponse struct {
	Error              ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	ClusterInventories ClusterInventory `bson:"clusterInventories" bson:"clusters"`     //集群组
	HostInventories    HostInventory    `json:"hostInventories" bson:"hostInventories"` //主机组
}

//指定云主机IP
type SetVmStaticIpRequest struct {
	ReqConfig
	VmInstanceUuid string              `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //云主机UUID
	SetVmStaticIp  SetVmStaticIpParams `json:"setVmStaticIp" bson:"setVmStaticIp"`
	SystemTags     []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmStaticIpParams struct {
	L3NetworkUuid string `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //三层网络UUID
	Ip            string `json:"ip" bson:"ip"`                       //指定IP地址
	Ip6           string `json:"ip6,omitempty" bson:"ip6,omitempty"` //指定IPv6地址
}

type SetVmStaticIpResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除云主机指定IP
type DeleteVmStaticIpRequest struct {
	ReqConfig
	VmInstanceUuid string                 `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //云主机UUID
	DeleteMode     string                 `json:"deleteMode" bson:"deleteMode"`         //删除方式，默认填 Permissive
	Params         DeleteVmStaticIpParams `json:"params" bson:"params"`
	SystemTags     []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DeleteVmStaticIpParams struct {
	L3NetworkUuid string `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //三层网络UUID
}

type DeleteVmStaticIpResponse struct {
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`
}

//获取云主机能力
type GetVmCapabilitiesRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmCapabilitiesResponse struct {
	Error        ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Capabilities map[string]bool `json:"capabilities" bson:"capabilities"`       //云主机能力结果
}

//更新云主机信息
type UpdateVmInstanceRequest struct {
	ReqConfig
	UUID             string                 `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateVmInstance UpdateVmInstanceParams `json:"updateVmInstance" bson:"updateVmInstance"`
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type UpdateVmInstanceParams struct {
	Name                 string `json:"name,omitempty" bson:"name,omitempty"`               //云主机名称
	Description          string `json:"description,omitempty" bson:"description,omitempty"` //云主机的详细描述
	State                string `json:"state,omitempty" bson:"state,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty" bson:"defaultL3NetworkUuid,omitempty"` //默认三层网络UUID 当l3NetworkUuids指定了多个三层网络时，该字段指定提供默认路由的三层网络。若不指定，l3NetworkUuids的第一个网络被选为默认网络。
	Platform             string `json:"platform,omitempty" bson:"platform,omitempty"`                         //云盘系统平台
	CpuNum               int    `json:"cpuNum,omitempty" bson:"cpuNum,omitempty"`                             //CPU数目
	MemorySize           int64  `json:"memorySize,omitempty" bson:"memorySize,omitempty"`                     //CPU数量/内存大小，注意：该参数与instanceOfferingUuid二选一
	GuestOsType          string `json:"guestOsType" bson:"guestOsType"`
}

type UpdateVmInstanceResponse struct {
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`
}

//克隆云主机到指定物理机
type CloneVmInstanceRequest struct {
	ReqConfig
	VmInstanceUuid  string                `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	CloneVmInstance CloneVmInstanceParams `json:"cloneVmInstance" bson:"cloneVmInstance"`           //放strategy，Names，Full
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CloneVmInstanceParams struct {
	Strategy string   `json:"strategy,omitempty" bson:"strategy,omitempty"` //云主机创建策略 创建后立刻启动InstantStart 创建后不启动CreateStopped
	Names    []string `json:"names" bson:"names"`                           //云主机的名字清单
	Full     bool     `json:"full" bson:"full"`                             //是否克隆已挂载数据盘
}

type CloneVmInstanceResponse struct {
	Error   ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Success bool                   `json:"success" bson:"success"`                 //操作是否成功
	Result  CloneVmInstanceResults `json:"result" bson:"result"`
}

//设置云主机时钟同步
type SetVmClockTrackRequest struct {
	ReqConfig
	UUID            string                `json:"uuid" bson:"uuid"` //云主机UUID
	SetVmClockTrack SetVmClockTrackParams `json:"setVmClockTrack" bson:"setVmClockTrack"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmClockTrackParams struct {
	Track string `json:"track" bson:"track"` //guest,host
}

type SetVmClockTrackResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机高可用级别
type SetVmInstanceHaLevelRequest struct {
	ReqConfig
	UUID       string                     `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	Params     SetVmInstanceHaLevelParams `json:"params" bson:"params"`                             //放level
	SystemTags []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmInstanceHaLevelParams struct {
	Level string `json:"level" bson:"level"` //云主机高可用级别： NeverStop 永不停机
}

type SetVmInstanceHaLevelResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机高可用级别
type GetVmInstanceHaLevelRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmInstanceHaLevelResponse struct {
	Level   string    `json:"level" bson:"level"` //云主机高可用级别： NeverStop 永不停机
	Success bool      `json:"success" bson:"success"`
	Error   ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//取消云主机高可用
type DeleteVmInstanceHaLevelRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DeleteVmInstanceHaLevelResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机Qga
type GetVmQgaRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmQgaResponse struct {
	UUID   string    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Enable bool      `json:"enable" bson:"enable"`
	Error  ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机Qga
type SetVmQgaRequest struct {
	ReqConfig
	UUID       string         `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SetVmQga   SetVmQgaParams `json:"setVmQga" bson:"setVmQga"`                         //放Enable
	SystemTags []string       `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string       `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmQgaParams struct {
	Enable bool `json:"enable" bson:"enable"` //资源的UUID，唯一标示该资源
}

type SetVmQgaResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机RDP开关状态
type GetVmRDPRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmRDPResponse struct {
	Enable bool      `json:"enable" bson:"enable"`
	Error  ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机RDP开关状态
type SetVmRDPRequest struct {
	ReqConfig
	UUID       string         `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SetVmRDP   SetVmRDPParams `json:"setVmRDP" bson:"setVmRDP"`                         //放Enable
	SystemTags []string       `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string       `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmRDPParams struct {
	Enable bool `json:"enable" bson:"enable"`
}

type SetVmRDPResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机支持的屏幕数
type GetVmMonitorNumberRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmMonitorNumberResponse struct {
	MonitorNumber int       `json:"monitorNumber" bson:"monitorNumber"`
	Error         ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机支持的屏幕数
type SetVmMonitorNumberRequest struct {
	ReqConfig
	UUID               string                   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SetVmMonitorNumber SetVmMonitorNumberParams `json:"setVmMonitorNumber" bson:"setVmMonitorNumber"`     //放MonitorNumber
	SystemTags         []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmMonitorNumberParams struct {
	MonitorNumber int `json:"monitorNumber" bson:"monitorNumber"`
}

type SetVmMonitorNumberResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改云主机根云盘
type ChangeVmImageRequest struct {
	ReqConfig
	VmInstanceUuid string              `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //资源的UUID，唯一标示该资源
	ChangeVmImage  ChangeVmImageParams `json:"changeVmImage" bson:"changeVmImage"`               //放ImageUuid
	SystemTags     []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type ChangeVmImageParams struct {
	ImageUuid    string `json:"imageUuid" bson:"imageUuid"`       //镜像UUID
	ResourceUuid string `json:"resourceUuid" bson:"resourceUuid"` //
}

type ChangeVmImageResponse struct {
	Inventory VmInstanceInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取候选镜像列表
type GetImageCandidatesForVmToChangeRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //资源的UUID，唯一标示该资源
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetImageCandidatesForVmToChangeResponse struct {
	Inventories []ImageInventory `bson:"inventories" bson:"inventories"`
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新云主机mac地址
type UpdateVmNicMacRequest struct {
	ReqConfig
	VmNicUuid      string               `json:"vmNicUuid" bson:"vmNicUuid"`                       //资源的UUID，唯一标示该资源
	UpdateVmNicMac UpdateVmNicMacParams `json:"updateVmNicMac" bson:"updateVmNicMac"`             //放mac
	SystemTags     []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type UpdateVmNicMacParams struct {
	Mac string `json:"mac" bson:"mac"` //mac地址
}

type UpdateVmNicMacResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机防IP欺骗启用状态
type SetVmCleanTrafficRequest struct {
	ReqConfig
	Uuid              string                  `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SetVmCleanTraffic SetVmCleanTrafficParams `json:"setVmCleanTraffic" bson:"setVmCleanTraffic"`       //放Enable
	SystemTags        []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmCleanTrafficParams struct {
	Enable bool `json:"enable" bson:"enable"` //
}

type SetVmCleanTrafficResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//为云主机创建CDROM
type CreateVmCdRomRequest struct {
	ReqConfig
	Params     CreateVmCdRomParams `json:"params" bson:"params"`                             //放Name,VmInstanceUuid,Description
	SystemTags []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type CreateVmCdRomParams struct {
	VmInstanceUuid string `json:"vmInstanceUuid" bson:"vmInstanceUuid"`               //云主机UUID
	Name           string `json:"name" bson:"name"`                                   //资源名称
	IsoUuid        string `json:"isoUuid,omitempty" bson:"isoUuid,omitempty"`         //ISO镜像UUID
	Description    string `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
	ResourceUuid   string `json:"resourceUuid " bson:"resourceUuid "`                 //资源UUID
}

type CreateVmCdRomResponse struct {
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VmCdRomInventory `json:"inventory" bson:"inventory"`
}

//删除CDROM
type DeleteVmCdRomRequest struct {
	ReqConfig
	Uuid       string   `json:"Uuid" bson:"Uuid"` //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"delete_mode" bson:"delete_mode"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type DeleteVmCdRomResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改CDROM
type UpdateVmCdRomRequest struct {
	ReqConfig
	Uuid          string              `json:"uuid" bson:"uuid"` //UUID
	UpdateVmCdRom UpdateVmCdRomParams `json:"updateVmCdRom" bson:"updateVmCdRom"`
	SystemTags    []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type UpdateVmCdRomParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //云盘名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
}

type UpdateVmCdRomResponse struct {
	Inventory VmCdRomInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机默认CDROM
type SetVmInstanceDefaultCdRomRequest struct {
	ReqConfig
	Uuid                      string                          `json:"Uuid" bson:"Uuid"`                                           //资源的UUID，唯一标示该资源
	VmInstanceUuid            string                          `json:"vmInstanceUuid" bson:"vmInstanceUuid"`                       //云主机UUID
	SetVmInstanceDefaultCdRom SetVmInstanceDefaultCdRomParams `json:"setVmInstanceDefaultCdRom" bson:"setVmInstanceDefaultCdRom"` //放空
	SystemTags                []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"`           //云主机系统标签
	UserTags                  []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`               //云主机用户标签
}

type SetVmInstanceDefaultCdRomParams struct {
}

type SetVmInstanceDefaultCdRomResponse struct {
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VmCdRomInventory `json:"inventory" bson:"inventory"`
}

//查询CDROM清单
type QueryVmCdRomRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type QueryVmCdRomResponse struct {
	Error       ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []VmCdRomInventory `json:"inventories" bson:"inventories"`
}

//更改云主机优先级级别
type UpdateVmPriorityRequest struct {
	ReqConfig
	Uuid             string                 `json:"Uuid" bson:"Uuid"` //资源的UUID，唯一标示该资源`
	UpdateVmPriority UpdateVmPriorityParams `json:"updateVmPriority" bson:"updateVmPriority"`
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type UpdateVmPriorityParams struct {
	Priority string `json:"priority" bson:"priority"`
}

type UpdateVmPriorityResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机显存
type SetVmQxlMemoryRequest struct {
	ReqConfig
	Uuid           string               `json:"Uuid" bson:"Uuid"`                                 //资源的UUID，唯一标示该资源
	SetVmQxlMemory SetVmQxlMemoryParams `json:"setVmQxlMemory" bson:"setVmQxlMemory"`             //放ram,vram,vgamem
	SystemTags     []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmQxlMemoryParams struct {
	Ram    int `json:"ram,omitempty" bson:"ram,omitempty"`       //默认值为65536
	Vram   int `json:"vram,omitempty" bson:"vram,omitempty"`     //默认值为32768
	Vgamem int `json:"vgamem,omitempty" bson:"vgamem,omitempty"` //默认值为16384
}

type SetVmQxlMemoryResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机虚拟声卡类型
type SetVmSoundTypeRequest struct {
	ReqConfig
	Uuid           string               `json:"Uuid" bson:"Uuid"`                                 //资源的UUID，唯一标示该资源
	SetVmSoundType SetVmSoundTypeParams `json:"setVmSoundType" bson:"setVmSoundType"`             //放SoundType
	SystemTags     []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmSoundTypeParams struct {
	SoundType string `json:"soundType" bson:"soundType"` //虚拟声卡类型，默认为ich6,可选ac97，默认为ich6
}

type SetVmSoundTypeResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取spice的CA证书
type GetSpiceCertificatesRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetSpiceCertificatesResponse struct {
	CertificateStr string    `json:"certificateStr" bson:"certificateStr"`   //spice CA证书
	Error          ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//为云主机加载增强工具镜像
type AttachGuestToolsIsoToVmRequest struct {
	ReqConfig
	Uuid                    string                        `json:"Uuid" bson:"Uuid"` //资源的UUID，唯一标示该资源
	AttachGuestToolsIsoToVm AttachGuestToolsIsoToVmParams `json:"attachGuestToolsIsoToVm" bson:"attachGuestToolsIsoToVm"`
	SystemTags              []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type AttachGuestToolsIsoToVmParams struct {
}

type AttachGuestToolsIsoToVmResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机可用的最新增强工具
type GetLatestGuestToolsForVmRequest struct {
	ReqConfig
	Uuid       string   `json:"Uuid" bson:"Uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetLatestGuestToolsForVmResponse struct {
	Success   bool                `json:"success" bson:"success"`                 //是否成功
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory GuestToolsInventory `json:"inventory" bson:"inventory"`
}

//获取云主机内部增强工具的信息
type GetVmGuestToolsInfoRequest struct {
	ReqConfig
	Uuid       string   `json:"Uuid" bson:"Uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmGuestToolsInfoResponse struct {
	Success  bool                   `json:"success" bson:"success"` //是否成功
	Features map[string]interface{} `json:"features" bson:"features"`
	Error    ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Version  string                 `json:"version" bson:"version"`                 //增强工具版本
	Status   string                 `json:"status" bson:"status"`                   //增强工具运行状态
}

//获取云主机第一启动项
type GetVmInstanceFirstBootDeviceRequest struct {
	ReqConfig
	Uuid       string   `json:"Uuid" bson:"Uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmInstanceFirstBootDeviceResponse struct {
	Error           ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	FirstBootDevice string    `json:"firstBootDevice" bson:"firstBootDevice"` //云主机第一启动项
}

//设置网卡型号
type UpdateVmNicDriverRequest struct {
	ReqConfig
	VmInstanceUuid    string                  `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //云主机UUID
	VmNicUuid         string                  `json:"vmNicUuid" bson:"vmNicUuid"`           //云主机网卡UUID，该网卡所在网络会从云主机卸载掉
	UpdateVmNicDriver UpdateVmNicDriverParams `json:"updateVmNicDriver" bson:"updateVmNicDriver"`
	SystemTags        []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type UpdateVmNicDriverParams struct {
	DriverType string `json:"driverType" bson:"driverType"` //网卡驱动类型,vcenter侧的VPC可选的网卡驱动:1.E1000E,2.E1000,3.Vmxnet3,4.Sriov
}

type UpdateVmNicDriverResponse struct {
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VmNicInventory `json:"inventory" bson:"inventory"`             //云主机第一启动项
}

//获取云主机设备地址
type GetVmDeviceAddressRequest struct {
	ReqConfig
	Uuid          string   `json:"Uuid" bson:"Uuid"`                                 //资源的UUID，唯一标示该资源
	ResourceTypes []string `json:"resourceTypes" bson:"resourceTypes"`               //资源类型
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmDeviceAddressResponse struct {
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Addresses map[string]interface{} `json:"addresses" bson:"addresses"`             //云主机第一启动项
	Success   bool                   `json:"success" bson:"success"`
}

//批量获取云主机能力
type GetVmsCapabilitiesRequest struct {
	ReqConfig
	Params     GetVmsCapabilitiesParams `json:"params" bson:"params"`
	SystemTags []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmsCapabilitiesParams struct {
	VmUuids []string `json:"vmUuids" bson:"vmUuids"` //待查询的云主机uuid集合
}

type GetVmsCapabilitiesResponse struct {
	VmsCaps map[string]interface{} `json:"vmsCaps" bson:"vmsCaps"`                 //云主机第一启动项
	Error   ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机vNUMA
type SetVmNumaRequest struct {
	ReqConfig
	UUID       string          `json:"uuid" bson:"uuid"`
	SetVmNuma  SetVmNumaParams `json:"setVmNuma" bson:"setVmNuma"`
	SystemTags []string        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string        `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmNumaParams struct {
	Enable bool `json:"enable" bson:"enable"` //云主机是否开启vm Numa功能
}

type SetVmNumaResponse struct {
	VmsCaps map[string]interface{} `json:"vmsCaps" bson:"vmsCaps"`                 //云主机第一启动项
	Error   ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取VM Numa开关状态
type GetVmNumaRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmNumaResponse struct {
	Enable bool      `json:"enable" bson:"enable"`                   //云主机是否开启vm Numa功能
	Error  ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询云主机的vNUMA拓扑信息
type GetVmvNUMATopologyRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmvNUMATopologyResponse struct {
	UUID     string           `json:"uuid" bson:"uuid"`         //资源的UUID，唯一标示该资源
	Name     string           `json:"name" bson:"name"`         //资源名称
	HostUuid string           `json:"hostUuid" bson:"hostUuid"` //物理机UUID
	Topology []VmNicInventory `json:"topology" bson:"topology"`
	Error    ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置云主机Emulator Pinning功能
type SetVmEmulatorPinningRequest struct {
	ReqConfig
	UUID                 string                     `json:"uuid" bson:"uuid"`
	SetVmEmulatorPinning SetVmEmulatorPinningParams `json:"setVmEmulatorPinning" bson:"setVmEmulatorPinning"`
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type SetVmEmulatorPinningParams struct {
	EmulatorPinning string `json:"emulatorPinning" bson:"emulatorPinning"` //云主机Emulator 绑定在宿主机的某些cpu上,格式为(\^?(\d+)(-\d+)?,)+(例如0-1,16)
}

type SetVmEmulatorPinningResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取VM Emulator Pin在那些物理机的cpu上
type GetVmEmulatorPinningRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //云主机用户标签
}

type GetVmEmulatorPinningResponse struct {
	EmulatorPinning string    `json:"emulatorPinning" bson:"emulatorPinning"` //云主机Emulator 绑定在宿主机的某些cpu上,格式为(\^?(\d+)(-\d+)?,)+(例如0-1,16)
	Error           ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}
