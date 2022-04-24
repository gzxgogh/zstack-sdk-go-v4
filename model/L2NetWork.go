package model

//创建VXLAN网络池
type CreateL2VxlanNetworkPoolRequest struct {
	ReqConfig
	Params     CreateL2VxlanNetworkPoolParams `json:"params" bson:"params"`
	SystemTags []string                       `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                       `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateL2VxlanNetworkPoolParams struct {
	Name              string `json:"name" bson:"name"`                                   //资源名称
	Description       string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	ZoneUuid          string `json:"zoneUuid" bson:"zoneUuid"`                           //区域UUID 若指定，云主机会在指定区域创建。
	PhysicalInterface string `json:"physicalInterface" bson:"physicalInterface"`
	Type              string `json:"type,omitempty" bson:"type,omitempty"`
	ResourceUuid      string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateL2VxlanNetworkPoolResponse struct {
	Error     ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory L2VxlanNetworkPoolInventory `json:"inventory" bson:"inventory"`
}

//查询VXLAN网络池
type QueryL2VxlanNetworkPoolRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}
type QueryL2VxlanNetworkPoolResponse struct {
	Error       ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []L2VxlanNetworkPoolInventory `json:"inventories" bson:"inventories"`
}

//创建VXLAN网络
type CreateL2VxlanNetworkRequest struct {
	ReqConfig
	Params     CreateL2VxlanNetworkParams `json:"params" bson:"params"`
	SystemTags []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateL2VxlanNetworkParams struct {
	Vni               int    `json:"vni,omitempty" bson:"vni,omitempty"`
	PoolUuid          string `json:"poolUuid" bson:"poolUuid"`
	Name              string `json:"name" bson:"name"`                                   //资源名称
	Description       string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	ZoneUuid          string `json:"zoneUuid" bson:"zoneUuid"`                           //区域UUID
	PhysicalInterface string `json:"physicalInterface" bson:"physicalInterface"`
	Type              string `json:"type,omitempty" bson:"type,omitempty"`
	ResourceUuid      string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateL2VxlanNetworkResponse struct {
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AttachedVxlanNetworkRefs `json:"inventory" bson:"inventory"`
}

//查询VXLAN网络
type QueryL2VxlanNetworkRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}
type QueryL2VxlanNetworkResponse struct {
	Error       ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []AttachedVxlanNetworkRefs `json:"inventories" bson:"inventories"`
}

//创建普通二层网络
type CreateL2NoVlanNetworkRequest struct {
	ReqConfig
	Params     CreateL2NoVlanNetworkParams `json:"params" bson:"params"`
	SystemTags []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateL2NoVlanNetworkParams struct {
	Name              string `json:"name" bson:"name"`                                   //资源名称
	Description       string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	ZoneUuid          string `json:"zoneUuid" bson:"zoneUuid"`                           //区域UUID 若指定，云主机会在指定区域创建。
	PhysicalInterface string `json:"physicalInterface" bson:"physicalInterface"`
	Type              string `json:"type,omitempty" bson:"type,omitempty"`
	ResourceUuid      string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateL2NoVlanNetworkResponse struct {
	Error     ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory L2VxlanNetworkPoolInventory `json:"inventory" bson:"inventory"`
}

//创建二层VLAN网络
type CreateL2VlanNetworkRequest struct {
	ReqConfig
	Params     CreateL2VlanNetworkParams `json:"params" bson:"params"`
	SystemTags []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateL2VlanNetworkParams struct {
	Vlan              int    `json:"vlan" bson:"vlan"`
	Name              string `json:"name" bson:"name"`                                   //资源名称
	Description       string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	ZoneUuid          string `json:"zoneUuid" bson:"zoneUuid"`                           //区域UUID
	PhysicalInterface string `json:"physicalInterface" bson:"physicalInterface"`
	Type              string `json:"type,omitempty" bson:"type,omitempty"`
	ResourceUuid      string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateL2VlanNetworkResponse struct {
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory L2VlanNetworkInventory `json:"inventory" bson:"inventory"`
}

//查询二层VLAN网络
type QueryL2VlanNetworkRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryL2VlanNetworkResponse struct {
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []L2VlanNetworkInventory `json:"inventories" bson:"inventories"`
}

//删除二层网络
type DeleteL2NetworkRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteL2NetworkResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询二层网络
type QueryL2NetworkRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryL2NetworkResponse struct {
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []L2VlanNetworkInventory `json:"inventories" bson:"inventories"`
}

//更新二层网络
type UpdateL2NetworkRequest struct {
	ReqConfig
	UUID            string                `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateL2Network UpdateL2NetworkParams `json:"updateL2Network" bson:"updateL2Network"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateL2NetworkParams struct {
	Name        string `json:"name" bson:"name"`               //资源名称
	Description string `json:"description" bson:"description"` //资源的详细描述
}

type UpdateL2NetworkResponse struct {
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory L2VlanNetworkInventory `json:"inventory" bson:"inventory"`
}

//获取二层网络类型
type GetL2NetworkTypesRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetL2NetworkTypesResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Types []string  `json:"types" bson:"types"`
}

//挂载二层网络到集群
type AttachL2NetworkToClusterRequest struct {
	ReqConfig
	L2NetworkUuid string                   `json:"l2NetworkUuid" bson:"l2NetworkUuid"`
	ClusterUuid   string                   `json:"clusterUuid" bson:"clusterUuid"`                   //集群UUID
	SystemTags    []map[string]interface{} `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachL2NetworkToClusterResponse struct {
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory L2VlanNetworkInventory `json:"inventory" bson:"inventory"`
}

//从集群上卸载二层网络
type DetachL2NetworkFromClusterRequest struct {
	ReqConfig
	L2NetworkUuid string   `json:"l2NetworkUuid" bson:"l2NetworkUuid"`
	ClusterUuid   string   `json:"clusterUuid" bson:"clusterUuid"`                   //集群UUID
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachL2NetworkFromClusterResponse struct {
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory L2VlanNetworkInventory `json:"inventory" bson:"inventory"`
}

//创建Vni Range
type CreateVniRangeRequest struct {
	ReqConfig
	L2NetworkUuid string               `json:"l2NetworkUuid" bson:"l2NetworkUuid"`
	Params        CreateVniRangeParams `json:"params" bson:"params"`
	SystemTags    []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateVniRangeParams struct {
	Name         string `json:"name" bson:"name"`               //资源名称
	Description  string `json:"description" bson:"description"` //资源的详细描述
	StartVni     int    `json:"startVni" bson:"startVni"`
	EndVni       int    `json:"endVni" bson:"endVni"`
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateVniRangeResponse struct {
	Error     ErrorCode         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VniRangeInventory `json:"inventory" bson:"inventory"`
}

//查询Vni Range
type QueryVniRangeRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryVniRangeResponse struct {
	Error       ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []VniRangeInventory `json:"inventories" bson:"inventories"`
}

//删除Vni Range
type DeleteVniRangeRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"`
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteVniRangeResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改Vni Range
type UpdateVniRangeRequest struct {
	ReqConfig
	Uuid           string               `json:"uuid" bson:"uuid"`
	UpdateVniRange UpdateVniRangeParams `json:"updateVniRange" bson:"updateVniRange"`
	SystemTags     []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateVniRangeParams struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"` //资源名称
}

type UpdateVniRangeResponse struct {
	Error     ErrorCode         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VniRangeInventory `json:"inventory" bson:"inventory"`
}

type VniRangeInventory struct {
	UUID          string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name          string `json:"name" bson:"name"`               //资源名称
	Description   string `json:"description" bson:"description"` //资源的详细描述
	StartVni      int    `json:"startVni" bson:"startVni"`
	EndVni        int    `json:"endVni" bson:"endVni"`
	CreateDate    string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate    string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	L2NetworkUuid string `json:"l2NetworkUuid" bson:"l2NetworkUuid"`
}

type L2VxlanNetworkPoolInventory struct {
	AttachedCidrs            map[string]interface{}     `json:"attachedCidrs,omitempty" bson:"attachedCidrs,omitempty"`
	UUID                     string                     `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name                     string                     `json:"name" bson:"name"`               //资源名称
	Description              string                     `json:"description" bson:"description"` //资源的详细描述
	ZoneUuid                 string                     `json:"zoneUuid" bson:"zoneUuid"`       //区域UUID 若指定，云主机会在指定区域创建。
	PhysicalInterface        string                     `json:"physicalInterface" bson:"physicalInterface"`
	Type                     string                     `json:"type" bson:"type"`
	CreateDate               string                     `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate               string                     `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	AttachedClusterUuids     []string                   `json:"attachedClusterUuids" bson:"attachedClusterUuids"`
	AttachedVtepRefs         []AttachedVtepRefs         `json:"attachedVtepRefs" bson:"attachedVtepRefs"`
	AttachedVxlanNetworkRefs []AttachedVxlanNetworkRefs `json:"attachedVxlanNetworkRefs" bson:"attachedVxlanNetworkRefs"`
	AttachedVniRanges        []VniRangeInventory        `json:"attachedVniRanges" bson:"attachedVniRanges"`
}

type L2VlanNetworkInventory struct {
	Vlan                 int      `json:"vlan" bson:"vlan"`
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
