package model

type UsedIp struct {
	UUID          string `json:"uuid" bson:"uuid"`                   //资源的UUID，唯一标示该资源
	IpRangeUUID   string `json:"ipRangeUuid" bson:"ipRangeUuid"`     //IP段UUID
	L3NetworkUUID string `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //	三层网络UUID
	IpVersion     int    `json:"ipVersion" bson:"ipVersion"`         //IP地址版本
	IP            string `json:"ip" bson:"ip"`                       //ip地址
	Netmask       string `json:"netmask" bson:"netmask"`             //子网掩码
	Gateway       string `json:"gateway" bson:"gateway"`             //网关
	UsedFor       string `json:"usedFor" bson:"usedFor"`             //
	//MetaData      string `json:"metaData" bson:"metaData"`
	IpInLong   int64  `json:"ipInLong" bson:"ipInLong"`     //
	VmNicUUID  string `json:"vmNicUuid" bson:"vmNicUuid"`   //云主机网卡UUID
	CreateDate string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}
