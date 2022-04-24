package model

//L3网络中是否存在可用VF网卡
type IsVfNicAvailableInL3NetworkRequest struct {
	ReqConfig
	L3NetworkUUID string   `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //	三层网络UUID
	HostUuid      string   `json:"hostUuid" bson:"hostUuid"`           //物理机UUID
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type IsVfNicAvailableInL3NetworkResponse struct {
	Error          ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	VfNicAvailable bool      `json:"vfNicAvailable" bson:"vfNicAvailable"`
}

//修改云主机网卡类型
type ChangeVmNicTypeRequest struct {
	ReqConfig
	VmNicUuid       string                `json:"vmNicUuid" bson:"vmNicUuid"` //云主机网卡UUID
	ChangeVmNicType ChangeVmNicTypeParams `json:"changeVmNicType" bson:"changeVmNicType"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeVmNicTypeParams struct {
	VmNicType string `json:"vmNicType" bson:"vmNicType"` //云主机网卡类型:VNIC
}

type ChangeVmNicTypeResponse struct {
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VmNicInventory `json:"inventory" bson:"inventory"`
}
