package model

type PciDeviceInventory struct {
	Uuid                        string                        `json:"uuid" bson:"uuid"`
	Name                        string                        `json:"name" bson:"name"`                         //名称
	Description                 string                        `json:"description" bson:"description"`           //详细描述
	HostUuid                    string                        `json:"hostUuid" bson:"hostUuid"`                 //物理机UUID
	ParentUuid                  string                        `json:"parentUuid" bson:"parentUuid"`             //物理PCI设备UUID
	VMInstanceUUID              string                        `json:"vmInstanceUuid" bson:"vmInstanceUuid"`     //云主机UUID
	PciSpecUuid                 string                        `json:"pciSpecUuid" bson:"pciSpecUuid"`           //PCI设备规格UUID
	VendorId                    string                        `json:"vendorId" bson:"vendorId"`                 //供应商ID
	DeviceId                    string                        `json:"deviceId" bson:"deviceId"`                 //设备ID
	SubvendorId                 string                        `json:"subvendorId" bson:"subvendorId"`           //子供应商ID
	SubdeviceId                 string                        `json:"subdeviceId" bson:"subdeviceId"`           //子设备ID
	PciDeviceAddress            string                        `json:"pciDeviceAddress" bson:"pciDeviceAddress"` //PCI设备地址
	CreateDate                  string                        `json:"createDate" bson:"createDate"`             //创建时间
	LastOpDate                  string                        `json:"lastOpDate" bson:"lastOpDate"`             //最后一次修改时间
	Type                        string                        `json:"type" bson:"type"`                         //类型,GPU_Video_Controller,GPU_Audio_Controller,
	State                       string                        `json:"state" bson:"state"`                       //Enabled,Disable
	Status                      string                        `json:"status" bson:"status"`                     //Active,Attached,System
	VirtStatus                  string                        `json:"virtStatus" bson:"virtStatus"`
	MetaData                    string                        `json:"metaData" bson:"metaData"`
	MatchedPciDeviceOfferingRef []MatchedPciDeviceOfferingRef `json:"matchedPciDeviceOfferingRef" bson:"matchedPciDeviceOfferingRef"`
	MdevSpecRefs                []mdevSpecRefs                `json:"mdevSpecRefs" bson:"mdevSpecRefs"`
}

type metaDataEntries struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
	Op    string `json:"op" bson:"op"`
}

type MatchedPciDeviceOfferingRef struct {
	PciDeviceUuid         string `json:"pciDeviceUuid" bson:"pciDeviceUuid"` //PCI设备UUID
	PciDeviceOfferingUuid string `json:"pciDeviceOfferingUuid" bson:"pciDeviceOfferingUuid"`
}

type mdevSpecRefs struct {
	PciDeviceUuid string `json:"pciDeviceUuid" bson:"pciDeviceUuid"` //PCI设备UUID
	MdevSpecUuid  string `json:"mdevSpecUuid" bson:"mdevSpecUuid"`   //MDEV设备规格UUID
	Effective     string `json:"effective" bson:"effective"`         //当前MDEV规格是否被用于切分该PCI设备
	CreateDate    string `json:"createDate" bson:"createDate"`       //创建时间
	LastOpDate    string `json:"lastOpDate" bson:"lastOpDate"`       //最后一次修改时间
}

type PciDeviceOfferingInventory struct {
	Uuid                      string                        `json:"uuid" bson:"uuid"`
	Name                      string                        `json:"name" bson:"name"`               //名称
	Description               string                        `json:"description" bson:"description"` //详细描述
	VendorId                  string                        `json:"vendorId" bson:"vendorId"`       //供应商ID
	DeviceId                  string                        `json:"deviceId" bson:"deviceId"`       //设备ID
	SubvendorId               string                        `json:"subvendorId" bson:"subvendorId"` //子供应商ID
	SubdeviceId               string                        `json:"subdeviceId" bson:"subdeviceId"` //子设备ID
	CreateDate                string                        `json:"createDate" bson:"createDate"`   //创建时间
	LastOpDate                string                        `json:"lastOpDate" bson:"lastOpDate"`   //最后一次修改时间
	Type                      string                        `json:"type" bson:"type"`               //GPU_Video,GPU_Audio,Generic
	AttachedInstanceOfferings []AttachedInstanceOfferings   `json:"AttachedInstanceOfferings" bson:"AttachedInstanceOfferings"`
	MatchedPciDevices         []MatchedPciDeviceOfferingRef `json:"matchedPciDevices" bson:"matchedPciDevices"`
}

type AttachedInstanceOfferings struct {
	Id                    int64    `json:"id" bson:"id"`
	InstanceOfferingUuid  string   `json:"instanceOfferingUuid" bson:"instanceOfferingUuid"`
	PciDeviceOfferingUuid string   `json:"pciDeviceOfferingUuid" bson:"pciDeviceOfferingUuid"`
	PciDeviceCount        string   `json:"pciDeviceCount" bson:"pciDeviceCount"`
	Metadata              []string `json:"metadata" bson:"metadata"`
}

type VmInstancePciDeviceSpecRefInventory struct {
	VMInstanceUUID  string `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //云主机UUID
	PciSpecUuid     string `json:"pciSpecUuid" bson:"pciSpecUuid"`       //PCI设备规格UUID
	PciDeviceNumber int    `json:"pciDeviceNumber " bson:"pciDeviceNumber "`
	CreateDate      string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type SpecCandidatesInventory struct {
	Uuid        string `json:"uuid" bson:"uuid"`
	Name        string `json:"name" bson:"name"`               //名称
	Description string `json:"description" bson:"description"` //详细描述
	VendorId    string `json:"vendorId" bson:"vendorId"`       //供应商ID
	DeviceId    string `json:"deviceId" bson:"deviceId"`       //设备ID
	SubvendorId string `json:"subvendorId" bson:"subvendorId"` //子供应商ID
	SubdeviceId string `json:"subdeviceId" bson:"subdeviceId"` //子设备ID
	RamSize     string `json:"ramSize" bson:"ramSize"`         //显存容量
	MaxPartNum  int    `json:"maxPartNum" bson:"maxPartNum"`   //最大切分数量
	IsVirtual   bool   `json:"isVirtual" bson:"isVirtual"`     //是否虚拟设备
	RomVersion  string `json:"romVersion" bson:"romVersion"`   //固件版本
	RomMd5sum   string `json:"romMd5sum" bson:"romMd5sum"`     //固件MD5
	CreateDate  string `json:"createDate" bson:"createDate"`   //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"`
	Type        string `json:"type" bson:"type"`   //GPU_Video,GPU_Audio,Generic
	State       string `json:"state" bson:"state"` //Enabled,Disable
}

type MdevDeviceSpecInventory struct {
	Uuid          string `json:"uuid" bson:"uuid"`
	Name          string `json:"name" bson:"name"`               //名称
	Description   string `json:"description" bson:"description"` //详细描述
	Specification string `json:"specification" bson:"specification"`
	CreateDate    string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate    string `json:"lastOpDate" bson:"lastOpDate"`
	Type          string `json:"type" bson:"type"`   //GPU_Video,GPU_Audio,Generic
	State         string `json:"state" bson:"state"` //Enabled,Disable
}

type MdevDeviceInventory struct {
	Uuid           string `json:"uuid" bson:"uuid"`
	Name           string `json:"name" bson:"name"`                     //名称
	Description    string `json:"description" bson:"description"`       //详细描述
	HostUuid       string `json:"hostUuid" bson:"hostUuid"`             //物理机UUID
	ParentUuid     string `json:"parentUuid" bson:"parentUuid"`         //物理PCI设备UUID
	VMInstanceUUID string `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //云主机UUID
	MdevSpecUuid   string `json:"mdevSpecUuid" bson:"mdevSpecUuid"`     //MDEV设备规格UUID
	CreateDate     string `json:"createDate" bson:"createDate"`         //创建时间
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"`         //最后一次修改时间
	Type           string `json:"type" bson:"type"`                     //类型,GPU_Video_Controller,GPU_Audio_Controller,
	State          string `json:"state" bson:"state"`                   //Enabled,Disable
	Status         string `json:"status" bson:"status"`                 //Active,Attached,System
}

type VmInstanceMdevDeviceSpecRefInventory struct {
	VMInstanceUUID   string `json:"vmInstanceUuid" bson:"vmInstanceUuid"`     //云主机UUID
	MdevSpecUuid     string `json:"mdevSpecUuid" bson:"mdevSpecUuid"`         //MDEV设备规格UUID
	MdevDeviceNumber int    `json:"mdevDeviceNumber" bson:"mdevDeviceNumber"` //需要为云主机挂载的符合设备规格的设备个数，默认为1
	CreateDate       string `json:"createDate" bson:"createDate"`             //创建时间
	LastOpDate       string `json:"lastOpDate" bson:"lastOpDate"`             //最后一次修改时间
}
