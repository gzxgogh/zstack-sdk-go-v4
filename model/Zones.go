package model

//创建区域
type CreateZoneRequest struct {
	ReqConfig
	Params     CreateZoneParams `json:"params" bson:"params"`
	SystemTags []string         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateZoneParams struct {
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //用户指定的资源uuid
}

type CreateZoneResponse struct {
	Error     ErrorCode     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory ZoneInventory `json:"inventory" bson:"inventory"`
}

//删除区域
type DeleteZoneRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteZoneResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询区域
type QueryZoneRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryZoneResponse struct {
	Error       ErrorCode       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []ZoneInventory `json:"inventories" bson:"inventories"`
}

//更新区域
type UpdateZoneRequest struct {
	ReqConfig
	UUID       string           `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateZone UpdateZoneParams `json:"updateZone" bson:"updateZone"`
	SystemTags []string         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateZoneParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateZoneResponse struct {
	Error     ErrorCode     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory ZoneInventory `json:"inventory" bson:"inventory"`
}

//改变区域的可用状态
type ChangeZoneStateRequest struct {
	ReqConfig
	UUID            string                `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeZoneState ChangeZoneStateParams `json:"changeZoneState" bson:"changeZoneState"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeZoneStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //enable，disable
}

type ChangeZoneStateResponse struct {
	Error     ErrorCode     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory ZoneInventory `json:"inventory" bson:"inventory"`
}

type ZoneInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name        string `json:"name" bson:"name"`               //资源名称
	Description string `json:"description" bson:"description"` //详细描述
	State       string `json:"state" bson:"state"`
	Type        string `json:"type" bson:"type"`             //
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}
