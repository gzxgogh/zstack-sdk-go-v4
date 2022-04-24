package model

//查询PCI设备
type QueryPciDeviceRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPciDeviceResponse struct {
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []PciDeviceInventory `json:"inventories" bson:"inventories"`
}

//更新PCI设备
type UpdatePciDeviceRequest struct {
	ReqConfig
	UUID            string                `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdatePciDevice UpdatePciDeviceParams `json:"updatePciDevice" bson:"updatePciDevice"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdatePciDeviceParams struct {
	State       string `json:"state" bson:"state"`             //Enabled,Disable
	Description string `json:"description" bson:"description"` //详细描述
	MetaData    string `json:"metaData" bson:"metaData"`
}

type UpdatePciDeviceResponse struct {
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory PciDeviceInventory `json:"inventory" bson:"inventory"`
}

//删除PCI设备
type DeletePciDeviceRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeletePciDeviceResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取PCI设备列表
type GetPciDeviceCandidatesForAttachingVmRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //云主机UUID
	Types          []string `json:"types" bson:"types"`
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetPciDeviceCandidatesForAttachingVmResponse struct {
	Success     bool                 `json:"success" bson:"success"`
	Inventories []PciDeviceInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取可加载PCI设备
type GetPciDeviceCandidatesForNewCreateVmRequest struct {
	ReqConfig
	HostUuid     string   `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"` //物理机UUID
	ClusterUuids []string `json:"clusterUuids,omitempty" bson:"clusterUuids,omitempty"`
	Types        []string `json:"types,omitempty" bson:"types,omitempty"`
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetPciDeviceCandidatesForNewCreateVmResponse struct {
	Success     bool                 `json:"success" bson:"success"`
	Inventories []PciDeviceInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//绑定PCI设备到云主机
type AttachPciDeviceToVmRequest struct {
	ReqConfig
	PciDeviceUuid string                    `json:"pciDeviceUuid" bson:"pciDeviceUuid"` //物理机UUID
	Params        AttachPciDeviceToVmParams `json:"params " bson:"params "`
	SystemTags    []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachPciDeviceToVmParams struct {
	VmInstanceUuid string `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
}

type AttachPciDeviceToVmResponse struct {
	Inventory PciDeviceInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//卸载PCI设备
type DetachPciDeviceFromVmRequest struct {
	ReqConfig
	PciDeviceUuid string                    `json:"pciDeviceUuid" bson:"pciDeviceUuid"` //物理机UUID
	Params        AttachPciDeviceToVmParams `json:"params " bson:"params "`
	SystemTags    []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachPciDeviceFromVmResponse struct {
	Inventory PciDeviceInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建PCI设备规格
type CreatePciDeviceOfferingRequest struct {
	ReqConfig
	Params     CreatePciDeviceOfferingParams `json:"params " bson:"params "`
	SystemTags []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreatePciDeviceOfferingParams struct {
	Name         string `json:"name,omitempty" bson:"name,omitempty"`               //名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Type         string `json:"type,omitempty" bson:"type,omitempty"`
	VendorId     string `json:"vendorId" bson:"vendorId"`                             //供应商ID
	DeviceId     string `json:"deviceId" bson:"deviceId"`                             //设备ID
	SubvendorId  string `json:"subvendorId,omitempty" bson:"subvendorId,omitempty"`   //子供应商ID
	SubdeviceId  string `json:"subdeviceId,omitempty" bson:"subdeviceId,omitempty"`   //子设备ID
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreatePciDeviceOfferingResponse struct {
	Inventory PciDeviceOfferingInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除PCI设备规格
type DeletePciDeviceOfferingRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeletePciDeviceOfferingResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询PCI设备规格信息
type QueryPciDeviceOfferingRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPciDeviceOfferingResponse struct {
	Error       ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []PciDeviceOfferingInventory `json:"inventories" bson:"inventories"`
}

//查询PCI设备规格匹配
type QueryPciDevicePciDeviceOfferingRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPciDevicePciDeviceOfferingResponse struct {
	Error       ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []MatchedPciDeviceOfferingRef `json:"inventories" bson:"inventories"`
}

//获取物理机lommu启用状态
type GetHostIommuStatusRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetHostIommuStatusResponse struct {
	Success bool      `json:"success" bson:"success"`
	Error   ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Status  string    `json:"status" bson:"status"`
}

//更新物理机Iommu启用状态
type UpdateHostIommuStateRequest struct {
	ReqConfig
	Uuid                 string                     `json:"uuid" bson:"uuid"`
	UpdateHostIommuState UpdateHostIommuStateParams `json:"updateHostIommuState" bson:"updateHostIommuState"`
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateHostIommuStateParams struct {
	State string `json:"state" bson:"state"`
}

type UpdateHostIommuStateResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	State string    `json:"state" bson:"state"`
}

//获取物理机lommu就绪状态
type GetHostIommuStateRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetHostIommuStateResponse struct {
	Success bool      `json:"success" bson:"success"`
	Error   ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	State   string    `json:"state" bson:"state"`
}

//将PCI设备规格加载到云主机
type AddPciDeviceSpecToVmInstanceRequest struct {
	ReqConfig
	PciSpecUuid    string                             `json:"pciSpecUuid" bson:"pciSpecUuid"`
	VmInstanceUuid string                             `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	Params         AddPciDeviceSpecToVmInstanceParams `json:"params" bson:"params"`
	SystemTags     []string                           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string                           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddPciDeviceSpecToVmInstanceParams struct {
	PciDeviceNumber int `json:"pciDeviceNumber " bson:"pciDeviceNumber "`
}

type AddPciDeviceSpecToVmInstanceResponse struct {
	Error     ErrorCode                           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VmInstancePciDeviceSpecRefInventory `json:"inventory" bson:"inventory"`
}

//将PCI设备规格从云主机卸载
type RemovePciDeviceSpecFromVmInstanceRequest struct {
	ReqConfig
	PciSpecUuid    string   `json:"pciSpecUuid" bson:"pciSpecUuid"`
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemovePciDeviceSpecFromVmInstanceResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新PCI设备规格
type UpdatePciDeviceSpecRequest struct {
	ReqConfig
	Uuid                string                    `json:"uuid" bson:"uuid"`
	UpdatePciDeviceSpec UpdatePciDeviceSpecParams `json:"updatePciDeviceSpec" bson:"updatePciDeviceSpec"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdatePciDeviceSpecParams struct {
	Name           string `json:"name,omitempty" bson:"name,omitempty"`                     //名称
	Description    string `json:"description,omitempty" bson:"description,omitempty"`       //详细描述
	RomContent     string `json:"romContent,omitempty" bson:"romContent,omitempty"`         //BASE64编码后的固件内容
	RomVersion     string `json:"romVersion,omitempty" bson:"romVersion,omitempty"`         //固件版本
	AbandonSpecRom string `json:"abandonSpecRom,omitempty" bson:"abandonSpecRom,omitempty"` //删除已有固件
	State          string `json:"state,omitempty" bson:"state,omitempty"`                   //PCI设备规格启用状态
}

type UpdatePciDeviceSpecResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取PCI设备规格候选列表
type GetPciDeviceSpecCandidatesRequest struct {
	ReqConfig
	ClusterUuids    string   `json:"clusterUuids,omitempty" bson:"clusterUuids,omitempty"`
	HostUuid        string   `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`
	VmInstanceUuid  string   `json:"vmInstanceUuid,omitempty" bson:"vmInstanceUuid,omitempty"`
	VmInstanceUuids string   `json:"vmInstanceUuids,omitempty" bson:"vmInstanceUuids,omitempty"`
	Types           []string `json:"types,omitempty" bson:"types,omitempty"`
	SystemTags      []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetPciDeviceSpecCandidatesResponse struct {
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []SpecCandidatesInventory `json:"inventories" bson:"inventories"`
}

//查询PCI设备规格
type QueryPciDeviceSpecRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPciDeviceSpecResponse struct {
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []SpecCandidatesInventory `json:"inventories" bson:"inventories"`
}

//查询云主机与PCI设备规格的关联关系
type QueryVmInstancePciDeviceSpecRefRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //云主机uuid
	PciSpecUuid    string   `json:"pciSpecUuid" bson:"pciSpecUuid"`
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVmInstancePciDeviceSpecRefResponse struct {
	Error       ErrorCode                             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []VmInstancePciDeviceSpecRefInventory `json:"inventories" bson:"inventories"`
}

//虚拟化切分支持SR-IOV的PCI设备
type GenerateSriovPciDevicesRequest struct {
	ReqConfig
	PciDeviceUuid           string                        `json:"pciDeviceUuid" bson:"pciDeviceUuid"` //PCI UUID
	GenerateSriovPciDevices GenerateSriovPciDevicesParams `json:"generateSriovPciDevices" bson:"generateSriovPciDevices"`
	SystemTags              []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GenerateSriovPciDevicesParams struct {
	VirtPartNum int `json:"virtPartNum" bson:"virtPartNum"`
}

type GenerateSriovPciDevicesResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//虚拟化还原支持SR-IOV的PCI设备
type UngenerateSriovPciDevicesRequest struct {
	ReqConfig
	PciDeviceUuid             string                          `json:"pciDeviceUuid" bson:"pciDeviceUuid"` //PCI UUID
	UngenerateSriovPciDevices UngenerateSriovPciDevicesParams `json:"ungenerateSriovPciDevices" bson:"ungenerateSriovPciDevices"`
	SystemTags                []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                  []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UngenerateSriovPciDevicesParams struct {
}

type UngenerateSriovPciDevicesResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//虚拟化切分支持VFIO_MDEV的PCI设备
type GenerateMdevDevicesRequest struct {
	ReqConfig
	PciDeviceUuid       string                    `json:"pciDeviceUuid" bson:"pciDeviceUuid"` //PCI UUID
	GenerateMdevDevices generateMdevDevicesParams `json:"generateMdevDevices" bson:"generateMdevDevices"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type generateMdevDevicesParams struct {
	MdevSpecUuid string `json:"mdevSpecUuid" bson:"mdevSpecUuid"`
}

type GenerateMdevDevicesResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//虚拟化还原支持VFIO_MDEV的PCI设备
type UngenerateMdevDevicesRequest struct {
	ReqConfig
	PciDeviceUuid         string                      `json:"pciDeviceUuid" bson:"pciDeviceUuid"` //PCI UUID
	UngenerateMdevDevices UngenerateMdevDevicesParams `json:"ungenerateMdevDevices" bson:"ungenerateMdevDevices"`
	SystemTags            []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags              []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UngenerateMdevDevicesParams struct {
}

type UngenerateMdevDevicesResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询PCI设备切分出的MDEV设备
type QueryMdevDeviceSpecRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //PCI UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryMdevDeviceSpecResponse struct {
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []MdevDeviceSpecInventory `json:"inventories" bson:"inventories"`
}

//将PCI设备切分出的MDEV设备绑定到云主机
type AttachMdevDeviceToVmRequest struct {
	ReqConfig
	MdevDeviceUuid string   `json:"mdevDeviceUuid" bson:"mdevDeviceUuid"`             //MDEV设备UUID
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachMdevDeviceToVmResponse struct {
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory MdevDeviceInventory `json:"inventory" bson:"inventory"`
}

//将PCI设备切分出的MDEV设备从云主机卸载
type DetachMdevDeviceFromVmRequest struct {
	ReqConfig
	MdevDeviceUuid string   `json:"mdevDeviceUuid" bson:"mdevDeviceUuid"`             //MDEV设备UUID
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachMdevDeviceFromVmResponse struct {
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory MdevDeviceInventory `json:"inventory" bson:"inventory"`
}

//更新PCI设备切分出的MDEV设备
type UpdateMdevDeviceRequest struct {
	ReqConfig
	Uuid             string                 `json:"Uuid" bson:"Uuid"` //MDEV设备UUID
	UpdateMdevDevice UpdateMdevDeviceParams `json:"updateMdevDevice" bson:"updateMdevDevice"`
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateMdevDeviceParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	State       string `json:"state,omitempty" bson:"state,omitempty"`             //Enabled,Disable
}

type UpdateMdevDeviceResponse struct {
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory MdevDeviceInventory `json:"inventory" bson:"inventory"`
}

//获取可用的MDEV设备
type GetMdevDeviceCandidatesRequest struct {
	ReqConfig
	ClusterUuids   []string `json:"clusterUuids,omitempty" bson:"clusterUuids,omitempty"`     //集群UUID
	HostUuid       string   `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`             //物理机UUID
	VmInstanceUuid string   `json:"vmInstanceUuid,omitempty" bson:"vmInstanceUuid,omitempty"` //	云主机UUID
	Types          []string `json:"types,omitempty" bson:"types,omitempty"`                   //设备类型
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`         //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`             //用户标签
}

type GetMdevDeviceCandidatesResponse struct {
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []MdevDeviceInventory `json:"inventories" bson:"inventories"`
}

//查询PCI设备切分出的MDEV设备
type QueryMdevDeviceRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //集群UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //用户标签
}

type QueryMdevDeviceResponse struct {
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []MdevDeviceInventory `json:"inventories" bson:"inventories"`
}

//将MDEV设备规格加载到云主机
type AddMdevDeviceSpecToVmInstanceRequest struct {
	ReqConfig
	VMInstanceUUID string                              `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //云主机UUID
	MdevSpecUuid   string                              `json:"mdevSpecUuid" bson:"mdevSpecUuid"`     //MDEV设备规格UUID
	Params         AddMdevDeviceSpecToVmInstanceParams `json:"params" bson:"params"`
	SystemTags     []string                            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string                            `json:"userTags,omitempty" bson:"userTags,omitempty"`     //用户标签
}

type AddMdevDeviceSpecToVmInstanceParams struct {
	MdevDeviceNumber int `json:"mdevDeviceNumber " bson:"mdevDeviceNumber "`
}

type AddMdevDeviceSpecToVmInstanceResponse struct {
	Error     ErrorCode                            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VmInstanceMdevDeviceSpecRefInventory `json:"inventory" bson:"inventory"`
}

//将MDEV设备规格从云主机卸载
type RemoveMdevDeviceSpecFromVmInstanceRequest struct {
	ReqConfig
	VMInstanceUUID string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //UUID
	MdevSpecUuid   string   `json:"mdevSpecUuid" bson:"mdevSpecUuid"`                 //MDEV设备规格UUID
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //用户标签
}

type RemoveMdevDeviceSpecFromVmInstanceResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新MDEV设备规格
type UpdateMdevDeviceSpecRequest struct {
	ReqConfig
	UUID                 string                     `json:"uuid" bson:"uuid"` //UUID
	UpdateMdevDeviceSpec UpdateMdevDeviceSpecParams `json:"updateMdevDeviceSpec" bson:"updateMdevDeviceSpec"`
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`     //用户标签
}

type UpdateMdevDeviceSpecParams struct {
	Name        string `json:"name" bson:"name"`               //名称
	Description string `json:"description" bson:"description"` //详细描述
	State       string `json:"state" bson:"state"`             //Enabled,Disable
}

type UpdateMdevDeviceSpecResponse struct {
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory MdevDeviceSpecInventory `json:"inventory" bson:"inventory"`
}

//获取可用的MDEV设备规格
type GetMdevDeviceSpecCandidatesRequest struct {
	ReqConfig
	HostUuid        string   `json:"hostUuid,omitempty" bson:"hostUuid",omitempty` //物理机UUID
	ClusterUuids    []string `json:"clusterUuids,omitempty" bson:"clusterUuids,omitempty"`
	VmInstanceUuid  string   `json:"vmInstanceUuid,omitempty" bson:"vmInstanceUuid,omitempty"`
	VmInstanceUuids string   `json:"vmInstanceUuids,omitempty" bson:"vmInstanceUuids,omitempty"`
	Types           []string `json:"types,omitempty" bson:"types,omitempty"`
	SystemTags      []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //用户标签
}

type GetMdevDeviceSpecCandidatesResponse struct {
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []MdevDeviceSpecInventory `json:"inventories" bson:"inventories"`
}

//查询云主机与MDEV设备规格的关联关系
type QueryVmInstanceMdevDeviceSpecRefRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	MdevSpecUuid   string   `json:"mdevSpecUuid" bson:"mdevSpecUuid"`
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`     //用户标签
}

type QueryVmInstanceMdevDeviceSpecRefResponse struct {
	Error       ErrorCode                              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []VmInstanceMdevDeviceSpecRefInventory `json:"inventories" bson:"inventories"`
}
