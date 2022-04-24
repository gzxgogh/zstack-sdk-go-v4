package model

//创建弹性IP
type CreateEipRequest struct {
	ReqConfig
	Params     CreateEipParams `json:"params" bson:"params"`
	SystemTags []string        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateEipParams struct {
	Name         string `json:"name" bson:"name"`                                   //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	VipUuid      string `json:"vipUuid" bson:"vipUuid"`
	VmNicUuid    string `json:"vmNicUuid,omitempty" bson:"vmNicUuid,omitempty"`
	UsedIpUuid   int    `json:"usedIpUuid,omitempty" bson:"usedIpUuid,omitempty"`     //亲和组策略
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateEipResponse struct {
	Inventory EipInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除弹性IP
type DeleteEipRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteEipResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询弹性IP
type QueryEipRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryEipResponse struct {
	Inventories []EipInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新弹性IP
type UpdateEipRequest struct {
	ReqConfig
	UUID       string          `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateEip  UpdateEipParams `json:"updateSecurityGroup" bson:"updateSecurityGroup"`
	SystemTags []string        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateEipParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateEipResponse struct {
	Inventory SecurityGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更改弹性IP启用状态
type ChangeEipStateRequest struct {
	ReqConfig
	UUID           string               `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeEipState ChangeEipStateParams `json:"changeVipState" bson:"changeVipState"`
	SystemTags     []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeEipStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"`
}

type ChangeEipStateResponse struct {
	Inventory VipInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取可绑定指定弹性IP的云主机网卡
type GetEipAttachableVmNicsRequest struct {
	ReqConfig
	EipUuid    string   `json:"eipUuid,omitempty" bson:"eipUuid,omitempty"`       //弹性IP UUID
	VipUuid    string   `json:"vipUuid,omitempty" bson:"vipUuid,omitempty"`       //VIP UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetEipAttachableVmNicsResponse struct {
	Inventories []VmNicInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//绑定弹性IP
type AttachEipRequest struct {
	ReqConfig
	EipUuid    string               `json:"eipUuid" bson:"eipUuid"` //弹性IP UUID
	VmNicUuid  string               `json:"vmNicUuid" bson:"vmNicUuid"`
	Params     ChangeEipStateParams `json:"params" bson:"params"`
	SystemTags []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachEipParams struct {
	UsedIpUuid string `json:"usedIpUuid,omitempty" bson:"usedIpUuid,omitempty"`
}

type AttachEipResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//解绑弹性IP
type DetachEipRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachEipResponse struct {
	Inventory EipInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type EipInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name        string `json:"name" bson:"name"`               //资源名称
	Description string `json:"description" bson:"description"` //详细描述
	VmNicUuid   string `json:"vmNicUuid" bson:"vmNicUuid"`     //云主机网卡UUID
	VipUuid     string `json:"vipUuid" bson:"vipUuid"`
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	State       string `json:"state" bson:"state"`
	VipIp       string `json:"VipIp" bson:"VipIp"`
	GuestIp     string `json:"guestIp" bson:"guestIp"`
}
