package model

//创建路由区域资源
type CreateVRouterOspfArea struct {
	Params     CreateVRouterOspfAreaParams `json:"params" bson:"params"`
	SystemTags []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateVRouterOspfAreaParams struct {
	AreaId       string   `json:"areaId" bson:"areaId"`                                 //区域Id，区域标识
	AreaAuth     string   `json:"areaAuth,omitempty" bson:"areaAuth,omitempty"`         //OSPF区域的认证方式
	AreaType     string   `json:"areaType,omitempty" bson:"areaType,omitempty"`         //区域类型
	Password     string   `json:"password,omitempty" bson:"password,omitempty"`         //认证方式为plaintext时的密码
	KeyId        string   `json:"keyId,omitempty" bson:"keyId,omitempty"`               //认证方式为MD5时用到的keyid
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreateVRouterOspfAreaResponse struct {
	Inventory RouterAreaInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除OSPF路由区域
type DeleteVRouterOspfAreaRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteVRouterOspfAreaResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取OSPF的邻居信息
type GetVRouterOspfNeighborRequest struct {
	ReqConfig
	VRouterUuid string   `json:"vRouterUuid" bson:"uuid"`                          //资源的UUID，唯一标示该资源
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetVRouterOspfNeighborResponse struct {
	Neighbors []string  `json:"neighbors" bson:"neighbors"`
	Error     ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询OSPF路由区域信息
type QueryVRouterOspfAreaRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVRouterOspfAreaResponse struct {
	Inventories []RouterAreaInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取路由器的ID
type GetVRouterRouterIdRequest struct {
	ReqConfig
	VRouterUuid string   `json:"vRouterUuid" bson:"uuid"`                          //资源的UUID，唯一标示该资源
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetVRouterRouterIdResponse struct {
	RouterId string    `json:"routerId" bson:"routerId"`
	Error    ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置路由器的ID
type SetVRouterRouterIdRequest struct {
	ReqConfig
	VRouterUuid string                   `json:"vRouterUuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Params      SetVRouterRouterIdParams `json:"params" bson:"params"`
	SystemTags  []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type SetVRouterRouterIdParams struct {
	RouterId string `json:"routerId" bson:"routerId"`
}

type SetVRouterRouterIdResponse struct {
	RouterId string    `json:"routerId" bson:"routerId"`
	Error    ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加网络到OSPF的区域
type AddVRouterNetworksToOspfAreaRequest struct {
	ReqConfig
	VRouterUuid    string                             `json:"vRouterUuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	RouterAreaUuid string                             `json:"routerAreaUuid" bson:"routerAreaUuid"`
	Params         AddVRouterNetworksToOspfAreaParams `json:"params" bson:"params"`
	SystemTags     []string                           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string                           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddVRouterNetworksToOspfAreaParams struct {
	L3NetworkUuids []string `json:"l3NetworkUuids" bson:"routerId"`
	ResourceUuid   string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids       []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type AddVRouterNetworksToOspfAreaResponse struct {
	RouterId string    `json:"routerId" bson:"routerId"`
	Error    ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从路由区域中移除网络
type RemoveVRouterNetworksFromOspfAreaRequest struct {
	ReqConfig
	UUIDS      []string `json:"uuids" bson:"uuids"`                               //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveVRouterNetworksFromOspfAreaResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更改OSPF的路由区域属性
type UpdateVRouterOspfAreaRequest struct {
	ReqConfig
	UUID                  string                      `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateVRouterOspfArea UpdateVRouterOspfAreaParams `json:"updateVRouterOspfArea" bson:"updateVRouterOspfArea"`
	SystemTags            []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags              []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateVRouterOspfAreaParams struct {
	AreaAuth string `json:"areaAuth,omitempty" bson:"areaAuth,omitempty"` //OSPF区域的认证方式
	AreaType string `json:"areaType,omitempty" bson:"areaType,omitempty"` //区域类型
	Password string `json:"password,omitempty" bson:"password,omitempty"` //认证方式为plaintext时的密码
	KeyId    string `json:"keyId,omitempty" bson:"keyId,omitempty"`       //认证方式为MD5时用到的keyid
}

type UpdateVRouterOspfAreaResponse struct {
	Inventory RouterAreaInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询路由加入到区域的网络信息
type QueryVRouterOspfNetworkRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVRouterOspfNetworkResponse struct {
	Inventories []VRouterOspfNetworkInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type RouterAreaInventory struct {
	UUID           string `json:"uuid" bson:"uuid"`     //资源的UUID，唯一标示该资源
	AreaId         string `json:"areaId" bson:"areaId"` //区域Id，区域标识
	Type           string `json:"type" bson:"type"`
	Authentication string `json:"authentication" bson:"authentication"`
	Password       string `json:"password,omitempty" bson:"password,omitempty"` //认证方式为plaintext时的密码
	KeyId          string `json:"keyId,omitempty" bson:"keyId,omitempty"`       //认证方式为MD5时用到的keyid
	CreateDate     string `json:"createDate" bson:"createDate"`                 //创建时间
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"`                 //最后一次修改时间
}

type VRouterOspfNetworkInventory struct {
	UUID           string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	VRouterUuid    string `json:"vRouterUuid" bson:"vRouterUuid"`
	RouterAreaUuid string `json:"routerAreaUuid" bson:"routerAreaUuid"`
	L3NetworkUuid  string `json:"l3NetworkUuid" bson:"l3NetworkUuid"`
	CreateDate     string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}
