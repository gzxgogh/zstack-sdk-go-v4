package model

//创建三层网络
type CreateL3NetworkRequest struct {
	ReqConfig
	Params     CreateL3NetworkParams `json:"params" bson:"params"`
	SystemTags []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateL3NetworkParams struct {
	Name          string `json:"name" bson:"name"`                                     //资源名称
	Description   string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	Type          string `json:"type,omitempty" bson:"type,omitempty"`                 //三层网络类型:L3BasicNetwork,L3VpcNetwork
	L2NetworkUuid string `json:"l2NetworkUuid" bson:"l2NetworkUuid"`                   //二层网络UUID
	IpVersion     string `json:"ipVersion,omitempty" bson:"ipVersion,omitempty"`       //ip协议号:4,6
	System        bool   `json:"system,omitempty" bson:"system,omitempty"`             //是否用于系统云主机
	DnsDomain     string `json:"dnsDomain,omitempty" bson:"dnsDomain,omitempty"`       //DNS域
	ResourceUuid  string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	Category      string `json:"category,omitempty" bson:"category,omitempty"`         //网络类型，需要与system标签搭配使用，system为true时可设置为Public、Private
}

type CreateL3NetworkResponse struct {
	Inventory L3NetworkInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除三层网络
type DeleteL3NetworkRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteL3NetworkResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询三层网络
type QueryL3NetworkRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryL3NetworkResponse struct {
	Inventories []L3NetworkInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新三层网络
type UpdateL3NetworkRequest struct {
	ReqConfig
	UUID            string                `json:"uuid" bson:"uuid"`
	UpdateL3Network UpdateL3NetworkParams `json:"updateL3Network" bson:"updateL3Network"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateL3NetworkParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	System      bool   `json:"system,omitempty" bson:"system,omitempty"`           //是否用于系统云主机
	DnsDomain   string `json:"dnsDomain,omitempty" bson:"dnsDomain,omitempty"`     //DNS域
	Category    string `json:"category,omitempty" bson:"category,omitempty"`       //网络类型，需要与system标签搭配使用，system为true时可设置为Public、Private
}

type UpdateL3NetworkResponse struct {
	Inventory L3NetworkInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取三层网络类型
type GetL3NetworkTypesRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetL3NetworkTypesResponse struct {
	Types []string  `json:"types" bson:"types"`
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//改变三层网络状态
type ChangeL3NetworkStateRequest struct {
	ReqConfig
	UUID                 string                     `json:"uuid" bson:"uuid"`
	ChangeL3NetworkState ChangeL3NetworkStateParams `json:"changeL3NetworkState" bson:"changeL3NetworkState"`
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeL3NetworkStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //enable,disable
}

type ChangeL3NetworkStateResponse struct {
	Inventory L3NetworkInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取网络DHCP服务所用地址
type GetL3NetworkDhcpIpAddressRequest struct {
	ReqConfig
	L3NetworkUuid string   `json:"l3NetworkUuid" bson:"l3NetworkUuid"`
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetL3NetworkDhcpIpAddressResponse struct {
	Ip      string    `json:"ip" bson:"ip"`
	Ip6     string    `json:"ip6" bson:"ip6"`
	Success bool      `json:"success" bson:"success"`
	Error   ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//向三层网络添加DNS
type AddDnsToL3NetworkRequest struct {
	ReqConfig
	L3NetworkUuid string                  `json:"l3NetworkUuid" bson:"l3NetworkUuid"`
	Params        AddDnsToL3NetworkParams `json:"params" bson:"params"`
	SystemTags    []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddDnsToL3NetworkParams struct {
	Dns string `json:"dns" bson:"dns"`
}

type AddDnsToL3NetworkResponse struct {
	Inventory L3NetworkInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从三层网络移除DNS
type RemoveDnsFromL3NetworkRequest struct {
	ReqConfig
	L3NetworkUuid string   `json:"l3NetworkUuid" bson:"l3NetworkUuid"`
	Dns           string   `json:"dns" bson:"dns"`
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveDnsFromL3NetworkResponse struct {
	Inventory L3NetworkInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//向三层网络添加主机路由
type AddHostRouteToL3NetworkRequest struct {
	ReqConfig
	L3NetworkUuid string                        `json:"l3NetworkUuid" bson:"l3NetworkUuid"`
	Params        AddHostRouteToL3NetworkParams `json:"params" bson:"params"`
	SystemTags    []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddHostRouteToL3NetworkParams struct {
	Prefix  string `json:"prefix" bson:"prefix"`
	Nexthop string `json:"nexthop" bson:"nexthop"`
}

type AddHostRouteToL3NetworkResponse struct {
	Inventory L3NetworkInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从三层网络移除主机路由
type RemoveHostRouteFromL3NetworkRequest struct {
	ReqConfig
	L3NetworkUuid string   `json:"l3NetworkUuid" bson:"l3NetworkUuid"`
	Prefix        string   `json:"prefix" bson:"prefix"`
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveHostRouteFromL3NetworkResponse struct {
	Inventory L3NetworkInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取空闲IP
type GetFreeIpRequest struct {
	ReqConfig
	L3NetworkUuid string   `json:"l3NetworkUuid,omitempty" bson:"l3NetworkUuid,omitempty"` //三层网络UUID,l3NetworkUuid和ipRangeUuid二选一
	IpRangeUuid   string   `json:"ipRangeUuid,omitempty" bson:"ipRangeUuid,omitempty"`     //IP段UUID
	Start         string   `json:"start,omitempty" bson:"start,omitempty"`                 //起始值
	IpRangeType   string   `json:"ipRangeType,omitempty" bson:"ipRangeType,omitempty"`     //地址类型:Normal,AddressPool
	IpVersion     int      `json:"ipVersion,omitempty" bson:"ipVersion,omitempty"`         //IP地址版本号:4,6
	Limit         int      `json:"limit,omitempty" bson:"limit,omitempty"`                 //数量限制
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`       //云主机系统标签
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetFreeIpResponse struct {
	Inventory L3NetworkInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//检查IP可用性
type CheckIpAvailabilityRequest struct {
	ReqConfig
	L3NetworkUuid string   `json:"l3NetworkUuid" bson:"l3NetworkUuid"`               //三层网络UUID,l3NetworkUuid和ipRangeUuid二选一
	Ip            string   `json:"ip" bson:"ip"`                                     //IP
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CheckIpAvailabilityResponse struct {
	Available bool      `json:"available" bson:"available"`
	Success   bool      `json:"success" bson:"success"`
	Error     ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取IP网络地址容量
type GetIpAddressCapacityRequest struct {
	ReqConfig
	ZoneUuids      []string `json:"zoneUuids,omitempty" bson:"zoneUuids,omitempty"`           //区域UUID
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty" bson:"l3NetworkUuids,omitempty"` //三层网络UUID,l3NetworkUuid和ipRangeUuid二选一
	IpRangeUuids   []string `json:"ipRangeUuids,omitempty" bson:"ipRangeUuids,omitempty"`
	All            bool     `json:"all,omitempty" bson:"all,omitempty"`               //	系统全局
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetIpAddressCapacityResponse struct {
	TotalCapacity           int64            `json:"totalCapacity" bson:"totalCapacity"`                     //IP地址容量
	AvailableCapacity       int64            `json:"availableCapacity" bson:"availableCapacity"`             //可用IP地址容量
	UsedIpAddressNumber     int64            `json:"usedIpAddressNumber" bson:"usedIpAddressNumber"`         //已使用IP数量
	Ipv4TotalCapacity       int64            `json:"ipv4TotalCapacity" bson:"ipv4TotalCapacity"`             //IPv4地址容量
	Ipv4AvailableCapacity   int64            `json:"ipv4AvailableCapacity" bson:"ipv4AvailableCapacity"`     //可用IPv4地址容量
	Ipv4UsedIpAddressNumber int64            `json:"ipv4UsedIpAddressNumber" bson:"ipv4UsedIpAddressNumber"` //已使用IPv4数量
	Ipv6TotalCapacity       int64            `json:"ipv6TotalCapacity" bson:"ipv6TotalCapacity"`             //IPv6地址容量
	Ipv6AvailableCapacity   int64            `json:"ipv6AvailableCapacity" bson:"ipv6AvailableCapacity"`     //可用IPv6地址容量
	Ipv6UsedIpAddressNumber int64            `json:"ipv6UsedIpAddressNumber" bson:"ipv6UsedIpAddressNumber"` //已使用IPv6数量
	ResourceType            string           `json:"resourceType" bson:"resourceType"`                       //所查询资源的类型（地址范围、三层网络、区域）
	Success                 bool             `json:"success" bson:"success"`                                 //成功
	CapacityData            []IpCapacityData `json:"capacityData" bson:"capacityData"`
	Error                   ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加IP地址范围
type AddIpRangeRequest struct {
	ReqConfig
	L3NetworkUuid string           `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //三层网络UUID,l3NetworkUuid和ipRangeUuid二选一
	Params        AddIpRangeParams `json:"params" bson:"params"`
	SystemTags    []string         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddIpRangeParams struct {
	Name         string `json:"name" bson:"name"`               //资源名称
	Description  string `json:"description" bson:"description"` //资源的详细描述
	StartIp      string `json:"startIp" bson:"startIp"`
	EndIp        string `json:"endIp" bson:"endIp"`
	Netmask      string `json:"netmask" bson:"netmask"`
	Gateway      string `json:"gateway" bson:"gateway"`
	IpRangeType  string `json:"ipRangeType" bson:"ipRangeType"`
	ResourceUuid string `json:"resourceUuid" bson:"resourceUuid"`
}

type AddIpRangeResponse struct {
	Inventory IpRangeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除IP地址范围
type DeleteIpRangeRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteIpRangeResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询IP地址范围
type QueryIpRangeRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryIpRangeResponse struct {
	Inventories []IpRangeInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新IP地址范围
type UpdateIpRangeRequest struct {
	ReqConfig
	UUID          string              `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateIpRange UpdateIpRangeParams `json:"updateIpRange" bson:"updateIpRange"`
	SystemTags    []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateIpRangeParams struct {
	Name        string `json:"name" bson:"name"`               //资源名称
	Description string `json:"description" bson:"description"` //资源的详细描述
}

type UpdateIpRangeResponse struct {
	Inventory IpRangeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//通过网络CIDR添加IP地址范围
type AddIpRangeByNetworkCidrRequest struct {
	ReqConfig
	L3NetworkUuid string                        `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //三层网络UUID
	Params        AddIpRangeByNetworkCidrParams `json:"params" bson:"params"`
	SystemTags    []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddIpRangeByNetworkCidrParams struct {
	Name         string `json:"name" bson:"name"`                                     //三层网络的名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //三层网络的详细描述
	NetworkCidr  string `json:"networkCidr" bson:"networkCidr"`                       //网络CIDR
	Gateway      string `json:"gateway,omitempty" bson:"gateway,omitempty"`           //网关
	IpRangeType  string `json:"ipRangeType,omitempty" bson:"ipRangeType,omitempty"`   //地址类型:Normal,AddressPool
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，三层网络会使用该字段值作为UUID
}

type AddIpRangeByNetworkCidrResponse struct {
	Inventory IpRangeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取三层网络Mtu值
type GetL3NetworkMtuRequest struct {
	ReqConfig
	L3NetworkUuid string   `json:"l3NetworkUuid" bson:"l3NetworkUuid"`               //三层网络UUID
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetL3NetworkMtuResponse struct {
	Mtu     int64     `json:"mtu" bson:"mtu"`
	Success bool      `json:"success" bson:"success"`
	Error   ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置三层网络Mtu值
type SetL3NetworkMtuRequest struct {
	ReqConfig
	L3NetworkUuid string                `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //三层网络UUID
	Params        SetL3NetworkMtuParams `json:"params" bson:"params"`
	SystemTags    []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type SetL3NetworkMtuParams struct {
	Mtu int64 `json:"mtu" bson:"mtu"`
}

type SetL3NetworkMtuResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取三层网络上路由器的接口地址
type GetL3NetworkRouterInterfaceIpRequest struct {
	ReqConfig
	L3NetworkUuid string   `json:"l3NetworkUuid" bson:"l3NetworkUuid"`               //三层网络UUID
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetL3NetworkRouterInterfaceIpResponse struct {
	RouterInterfaceIp string    `json:"routerInterfaceIp" bson:"routerInterfaceIp"` //三层网络上路由器的接口地址，仅当会在普通三层网络上创建云路由器或在VPC网络上加载VPC路由器时有效
	Success           bool      `json:"success" bson:"success"`
	Error             ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//设置三层网络路由器接口IP
type SetL3NetworkRouterInterfaceIpRequest struct {
	ReqConfig
	L3NetworkUuid string                              `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //三层网络UUID
	Params        SetL3NetworkRouterInterfaceIpParams `json:"params" bson:"params"`
	SystemTags    []string                            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string                            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type SetL3NetworkRouterInterfaceIpParams struct {
	RouterInterfaceIp string `json:"routerInterfaceIp" bson:"routerInterfaceIp"` //三层网络上路由器的接口地址，仅当会在普通三层网络上创建云路由器或在VPC网络上加载VPC路由器时有效
}

type SetL3NetworkRouterInterfaceIpResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加IPv6地址范围
type AddIpv6RangeRequest struct {
	ReqConfig
	L3NetworkUuid string             `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //三层网络UUID
	Params        AddIpv6RangeParams `json:"params" bson:"params"`
	SystemTags    []string           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddIpv6RangeParams struct {
	Name         string `json:"name" bson:"name"`               //资源名称
	Description  string `json:"description" bson:"description"` //资源的详细描述
	StartIp      string `json:"StartIp" bson:"StartIp"`
	EndIp        string `json:"EndIp" bson:"EndIp"`
	Gateway      string `json:"gateway" bson:"gateway"`
	PrefixLen    string `json:"prefixLen" bson:"prefixLen"`                           //掩码长度
	AddressMode  string `json:"addressMode" bson:"addressMode"`                       //IPv6地址分配模式
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type AddIpv6RangeResponse struct {
	Inventory IpRangeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//通过网络CIDR添加IPv6地址范围
type AddIpv6RangeByNetworkCidrRequest struct {
	ReqConfig
	L3NetworkUuid string                          `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //三层网络UUID
	Params        AddIpv6RangeByNetworkCidrParams `json:"params" bson:"params"`
	SystemTags    []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddIpv6RangeByNetworkCidrParams struct {
	Name         string `json:"name" bson:"name"`               //资源名称
	Description  string `json:"description" bson:"description"` //资源的详细描述
	NetworkCidr  string `json:"networkCidr" bson:"networkCidr"`
	AddressMode  string `json:"addressMode" bson:"addressMode"`                       //IPv6地址分配模式
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type AddIpv6RangeByNetworkCidrResponse struct {
	Inventory IpRangeInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询IP地址
type QueryIpAddressRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryIpAddressResponse struct {
	Inventories []IpAddressInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取三层网络IP地址使用情况统计
type GetL3NetworkIpStatisticRequest struct {
	ReqConfig
	L3NetworkUuid string                        `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //三层网络UUID
	Params        GetL3NetworkIpStatisticParams `json:"params" bson:"params"`
	SystemTags    []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetL3NetworkIpStatisticParams struct {
	ResourceType   string `json:"resourceType,omitempty" bson:"resourceType,omitempty"`
	Ip             string `json:"ip,omitempty" bson:"ip,omitempty"`
	SortBy         string `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortDirection  string `json:"sortDirection,omitempty" bson:"sortDirection,omitempty"`
	Start          int    `json:"start,omitempty" bson:"start,omitempty"`
	Limit          int    `json:"limit,omitempty" bson:"limit,omitempty"`
	ReplyWithCount bool   `json:"replyWithCount,omitempty" bson:"replyWithCount,omitempty"`
}

type GetL3NetworkIpStatisticResponse struct {
	Total        int          `json:"total" bson:"total"`
	IpStatistics IpStatistics `json:"ipStatistics" bson:"ipStatistics"`
	Success      bool         `json:"success" bson:"success"`
	Error        ErrorCode    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type IpStatistics struct {
	Ip             string   `json:"ip,omitempty" bson:"ip,omitempty"`
	VipUuid        string   `json:"vipUuid" bson:"vipUuid"`
	VipName        string   `json:"vipName" bson:"vipName"`
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	VmInstanceName string   `json:"vmInstanceName" bson:"vmInstanceName"`
	VmInstanceType string   `json:"vmInstanceType" bson:"vmInstanceType"`
	VmDefaultIp    string   `json:"vmDefaultIp" bson:"vmDefaultIp"`
	ResourceTypes  []string `json:"resourceTypes" bson:"resourceTypes"`
	State          string   `json:"state" bson:"state"`
	UseFor         string   `json:"useFor" bson:"useFor"`
	CreateDate     string   `json:"createDate" bson:"createDate"` //创建时间
	OwnerName      string   `json:"ownerName" bson:"ownerName"`
}

//查询IP地址池
type QueryAddressPoolRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //三层网络UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAddressPoolResponse struct {
	Inventories []IpRangeInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type L3NetworkInventory struct {
	UUID            string             `json:"uuid" bson:"uuid"`                               //资源的UUID，唯一标示该资源
	Name            string             `json:"name" bson:"name"`                               //资源名称
	Description     string             `json:"description" bson:"description"`                 //资源的详细描述
	Type            string             `json:"type" bson:"type"`                               //三层网络类型
	ZoneUuid        string             `json:"zoneUuid" bson:"zoneUuid"`                       //区域UUID 若指定，云主机会在指定区域创建。
	L2NetworkUuid   string             `json:"l2NetworkUuid" bson:"l2NetworkUuid"`             //二层网络UUID
	State           string             `json:"state" bson:"state"`                             //三层网络的可用状态
	DnsDomain       string             `json:"dnsDomain" bson:"dnsDomain"`                     //DNS域
	System          bool               `json:"system" bson:"system"`                           //是否用于系统云主机
	Category        string             `json:"category" bson:"category"`                       //网络类型，需要与system标签搭配使用，system为true时可设置为Public、Private
	IpVersion       int                `json:"ipVersion,omitempty" bson:"ipVersion,omitempty"` //ip协议号:4,6
	CreateDate      string             `json:"createDate" bson:"createDate"`                   //创建时间
	LastOpDate      string             `json:"lastOpDate" bson:"lastOpDate"`                   //最后一次修改时间
	Dns             []string           `json:"dns" bson:"dns"`
	IpRanges        []IpRangeInventory `json:"ipRanges" bson:"ipRanges"`
	NetworkServices []NetworkServices  `json:"networkServices" bson:"networkServices"`
	HostRoute       []HostRoute        `json:"hostRoute" bson:"hostRoute"`
}

type IpRangeInventory struct {
	UUID          string `json:"uuid" bson:"uuid"`                   //资源的UUID，唯一标示该资源
	L3NetworkUuid string `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //三层网络UUID
	Name          string `json:"name" bson:"name"`                   //资源名称
	Description   string `json:"description" bson:"description"`     //资源的详细描述
	StartIp       string `json:"StartIp" bson:"StartIp"`
	EndIp         string `json:"EndIp" bson:"EndIp"`
	Netmask       string `json:"netmask" bson:"netmask"`     //网络掩码
	PrefixLen     string `json:"prefixLen" bson:"prefixLen"` //掩码长度
	Gateway       string `json:"gateway" bson:"gateway"`     //网关地址
	NetworkCidr   string `json:"networkCidr" bson:"networkCidr"`
	IpVersion     string `json:"ipVersion" bson:"ipVersion"`     //ip协议号:4,6
	AddressMode   string `json:"addressMode" bson:"addressMode"` //IPv6地址分配模式
	CreateDate    string `json:"createDate" bson:"createDate"`   //创建时间
	LastOpDate    string `json:"lastOpDate" bson:"lastOpDate"`   //最后一次修改时间
	IpRangeType   string `json:"ipRangeType" bson:"ipRangeType"`
}

type NetworkServices struct {
	L3NetworkUuid              string `json:"l3NetworkUuid" bson:"l3NetworkUuid"`                           //三层网络UUID
	NetworkServiceProviderUuid string `json:"networkServiceProviderUuid" bson:"networkServiceProviderUuid"` //网络服务提供模块UUID
	NetworkServiceType         string `json:"networkServiceType" bson:"networkServiceType"`
}

type HostRoute struct {
	Id            string `json:"id" bson:"id"`
	L3NetworkUuid string `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //三层网络UUID
	Prefix        string `json:"prefix" bson:"prefix"`
	Nexthop       string `json:"nexthop" bson:"nexthop"`
	CreateDate    string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate    string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type IpCapacityData struct {
	ResourceUuid            string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`   //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TotalCapacity           int64  `json:"totalCapacity" bson:"totalCapacity"`                     //IP地址总容量
	AvailableCapacity       int64  `json:"availableCapacity" bson:"availableCapacity"`             //可用IP地址容量
	UsedIpAddressNumber     int64  `json:"usedIpAddressNumber" bson:"usedIpAddressNumber"`         //已用IP地址容量
	Ipv4TotalCapacity       int64  `json:"ipv4TotalCapacity" bson:"ipv4TotalCapacity"`             //IPv4地址总容量
	Ipv4AvailableCapacity   int64  `json:"ipv4AvailableCapacity" bson:"ipv4AvailableCapacity"`     //可用IPv4地址容量
	Ipv4UsedIpAddressNumber int64  `json:"ipv4UsedIpAddressNumber" bson:"ipv4UsedIpAddressNumber"` //已用IPv4地址容量
	Ipv6TotalCapacity       int64  `json:"ipv6TotalCapacity" bson:"ipv6TotalCapacity"`             //IPv6地址总容量
	Ipv6AvailableCapacity   int64  `json:"ipv6AvailableCapacity" bson:"ipv6AvailableCapacity"`     //	可用IPv6地址容量
	Ipv6UsedIpAddressNumber int64  `json:"ipv6UsedIpAddressNumber" bson:"ipv6UsedIpAddressNumber"` //已用IPv6地址容量
}

type IpAddressInventory struct {
	UUID          string `json:"uuid" bson:"uuid"`                   //资源的UUID，唯一标示该资源
	L3NetworkUuid string `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //三层网络UUID
	IpRangeUuid   string `json:"ipRangeUuid" bson:"ipRangeUuid"`     //资源名称
	IpVersion     string `json:"ipVersion" bson:"ipVersion"`         //ip协议号:4,6
	Ip            string `json:"ip" bson:"ip"`
	Netmask       string `json:"netmask" bson:"netmask"` //网络掩码
	Gateway       string `json:"gateway" bson:"gateway"` //网关地址
	UsedFor       string `json:"usedFor" bson:"usedFor"`
	IpInLong      int64  `json:"ipInLong" bson:"ipInLong"`
	VmNicUuid     string `json:"vmNicUuid" bson:"vmNicUuid"`   //云主机网卡UUID
	CreateDate    string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate    string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}
