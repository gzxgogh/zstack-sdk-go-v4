package model

//添加SDN控制器
type AddSdnControllerRequest struct {
	ReqConfig
	Params     AddSdnControllerParams `json:"params" bson:"params"`
	SystemTags []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddSdnControllerParams struct {
	VendorType   string   `json:"vendorType" bson:"vendorType"`
	Name         string   `json:"name" bson:"name"`                                   //资源名称
	Description  string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Ip           string   `json:"ip" bson:"ip"`
	Username     string   `json:"userName" bson:"userName"`                             //用户名
	Password     string   `json:"password" bson:"password"`                             //用户密码
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type AddSdnControllerResponse struct {
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory SdnControllerInventory `json:"inventory" bson:"inventory"`
}

//查询SDN控制器
type QuerySdnControllerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySdnControllerResponse struct {
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []SdnControllerInventory `json:"inventories" bson:"inventories"`
}

//删除SDN控制器
type RemoveSdnControllerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveSdnControllerResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新SDN控制器
type UpdateSdnControllerRequest struct {
	ReqConfig
	UUID                string                    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateSdnController UpdateSdnControllerParams `json:"updateSdnController" bson:"updateSdnController"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateSdnControllerParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateSdnControllerResponse struct {
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory SdnControllerInventory `json:"inventory" bson:"inventory"`
}

type SdnControllerInventory struct {
	UUID        string     `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	VendorType  string     `json:"vendorType" bson:"vendorType"`
	Name        string     `json:"name" bson:"name"`               //资源名称
	Description string     `json:"description" bson:"description"` //资源的详细描述
	Ip          string     `json:"ip" bson:"ip"`
	Username    string     `json:"username" bson:"username"`     //用户名
	Password    string     `json:"password" bson:"password"`     //用户密码
	CreateDate  string     `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string     `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	VniRanges   VniRanges  `json:"vniRanges" bson:"vniRanges"`
	VxlanPools  VxlanPools `json:"vxlanPools" bson:"vxlanPools"`
}

type VniRanges struct {
	StartVni int `json:"startVni" bson:"startVni"`
	EndVni   int `json:"endVni" bson:"endVni"`
}

type VxlanPools struct {
	SdnControllerUuid        string                   `json:"sdnControllerUuid" bson:"sdnControllerUuid"`
	AttachedCidrs            map[string]interface{}   `json:"attachedCidrs" bson:"attachedCidrs"`
	UUID                     string                   `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name                     string                   `json:"name" bson:"name"`               //资源名称
	Description              string                   `json:"description" bson:"description"` //资源的详细描述
	ZoneUuid                 string                   `json:"zoneUuid" bson:"zoneUuid"`       //区域UUID 若指定，云主机会在指定区域创建。
	PhysicalInterface        string                   `json:"physicalInterface" bson:"physicalInterface"`
	Type                     string                   `json:"type" bson:"type"`
	CreateDate               string                   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate               string                   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	AttachedClusterUuids     []string                 `json:"attachedClusterUuids" bson:"attachedClusterUuids"`
	AttachedVtepRefs         AttachedVtepRefs         `json:"attachedVtepRefs" bson:"attachedVtepRefs"`
	AttachedVxlanNetworkRefs AttachedVxlanNetworkRefs `json:"attachedVxlanNetworkRefs" bson:"attachedVxlanNetworkRefs"`
	AttachedVniRanges        VniRangeInventory        `json:"attachedVniRanges" bson:"attachedVniRanges"`
}

type AttachedVtepRefs struct {
	UUID       string `json:"uuid" bson:"uuid"`         //资源的UUID，唯一标示该资源
	HostUuid   string `json:"hostUuid" bson:"hostUuid"` //物理机UUID
	VtepIp     string `json:"vtepIp" bson:"vtepIp"`
	Port       int    `json:"port" bson:"port"`
	Type       string `json:"type" bson:"type"`
	CreateDate string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	PoolUuid   string `json:"poolUuid" bson:"poolUuid"`
}

type AttachedVxlanNetworkRefs struct {
	Vni                  int      `json:"vni" bson:"vni"`
	PoolUuid             string   `json:"poolUuid" bson:"poolUuid"`
	UUID                 string   `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name                 string   `json:"name" bson:"name"`               //资源名称
	Description          string   `json:"description" bson:"description"` //资源的详细描述
	ZoneUuid             string   `json:"zoneUuid" bson:"zoneUuid"`       //区域UUID 若指定，云主机会在指定区域创建。
	PhysicalInterface    string   `json:"physicalInterface" bson:"physicalInterface"`
	Type                 string   `json:"type" bson:"type"`
	CreateDate           string   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate           string   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	AttachedClusterUuids []string `json:"attachedClusterUuids" bson:"attachedClusterUuids"`
}
