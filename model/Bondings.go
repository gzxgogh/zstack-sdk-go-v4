package model

type Bondings struct {
	UUID            string   `json:"uuid" bson:"uuid"`                     //资源的UUID，唯一标示该资源
	HostUuid        string   `json:"hostUuid" bson:"hostUuid"`             //物理机UUID
	BondingName     string   `json:"bondingName" bson:"bondingName"`       //Bond名称
	Mode            string   `json:"mode" bson:"mode"`                     //Bond模式
	XmitHashPolicy  string   `json:"xmitHashPolicy" bson:"xmitHashPolicy"` //哈希策略
	MiiStatus       string   `json:"miiStatus" bson:"miiStatus"`           //mii状态
	Mac             string   `json:"mac" bson:"mac"`                       //MAC地址
	IpAddresses     []string `json:"ipAddresses" bson:"ipAddresses"`       //IP地址
	Miimon          int64    `json:"miimon" bson:"miimon"`                 //mii监控间隔
	AllSlavesActive bool     `json:"allSlavesActive" bson:"allSlavesActive"`
	CreateDate      string   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate      string   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	Slaves          []Slaves `json:"slaves" bson:"slaves"`
}

type Slaves struct {
	UUID             string   `json:"uuid" bson:"uuid"`                         //资源的UUID，唯一标示该资源
	HostUuid         string   `json:"hostUuid" bson:"hostUuid"`                 //物理机UUID
	BondingUuid      string   `json:"bondingUuid" bson:"bondingUuid"`           //Bond UUID
	InterfaceName    string   `json:"interfaceName" bson:"interfaceName"`       //网卡名称
	InterfaceType    string   `json:"interfaceType" bson:"interfaceType"`       //网卡应用状态，有nomaster、bridgeSlave、bondSlave
	Speed            int64    `json:"speed" bson:"speed"`                       //网卡速率
	SlaveActive      bool     `json:"slaveActive" bson:"slaveActive"`           //Bond链路状态
	CarrierActive    bool     `json:"carrierActive" bson:"carrierActive"`       //物理链路状态
	IpAddresses      []string `json:"ipAddresses" bson:"ipAddresses"`           //IP地址
	Mac              string   `json:"mac" bson:"mac"`                           //MAC地址
	PciDeviceAddress string   `json:"pciDeviceAddress" bson:"pciDeviceAddress"` //网卡PCI地址
	CreateDate       string   `json:"createDate" bson:"createDate"`             //创建时间
	LastOpDate       string   `json:"lastOpDate" bson:"lastOpDate"`
}
