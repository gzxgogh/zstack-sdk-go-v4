package model

type VmNicInventory struct {
	UUID           string   `json:"uuid" bson:"uuid"`                     //资源的UUID，唯一标示该资源
	VMInstanceUUID string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //	云主机UUID
	L3NetworkUUID  string   `json:"l3NetworkUuid" bson:"l3NetworkUuid"`   //	三层网络UUID
	IP             string   `json:"ip" bson:"ip"`                         //ip地址
	Mac            string   `json:"mac" bson:"mac"`                       //mac地址
	HypervisorType string   `json:"hypervisorType" bson:"hypervisorType"` //虚拟化类型
	Netmask        string   `json:"netmask" bson:"netmask"`               //子网掩码
	Gateway        string   `json:"gateway" bson:"gateway"`               //网关
	MetaData       string   `json:"metaData" bson:"metaData"`             //内部使用的保留域，元数据
	IpVersion      int      `json:"ipVersion" bson:"ipVersion"`           //IP地址版本
	DeviceID       int      `json:"deviceId" bson:"deviceId"`             //设备ID 标识网卡在客户操作系统（guest operating system）以太网设备中顺序的整形数字。例如， 0通常代表eth0，1通常代表eth1。
	InternalName   string   `json:"internalName" bson:"internalName"`     //名称
	Type           string   `json:"type" bson:"type"`                     //
	CreateDate     string   `json:"createDate" bson:"createDate"`         //创建时间
	LastOpDate     string   `json:"lastOpDate" bson:"lastOpDate"`         //最后一次修改时间
	UsedIps        []UsedIp `json:"usedIps" bson:"usedIps"`
}
