package model

//创建负载均衡器
type CreateLoadBalancerRequest struct {
	ReqConfig
	Params     CreateLoadBalancerParams `json:"params" bson:"params"`
	SystemTags []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateLoadBalancerParams struct {
	VipUuid      string `json:"vipUuid" bson:"vipUuid"`
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateLoadBalancerResponse struct {
	Inventory LoadBalancerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除负载均衡器
type DeleteLoadBalancerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteLoadBalancerResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询负载均衡器
type QueryLoadBalancerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryLoadBalancerResponse struct {
	Inventories []LoadBalancerInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//刷新负载均衡器
type RefreshLoadBalancerRequest struct {
	ReqConfig
	UUID                string                    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	RefreshLoadBalancer RefreshLoadBalancerParams `json:"updatePortForwardingRule" bson:"updatePortForwardingRule"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RefreshLoadBalancerParams struct {
}

type RefreshLoadBalancerResponse struct {
	Inventory LoadBalancerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建负载均衡监听器
type CreateLoadBalancerListenerRequest struct {
	ReqConfig
	LoadBalancerUuid string                           `json:"loadBalancerUuid" bson:"loadBalancerUuid"` //负载均衡器UUID
	Params           CreateLoadBalancerListenerParams `json:"params" bson:"params"`
	SystemTags       []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateLoadBalancerListenerParams struct {
	Name                string   `json:"name" bson:"name"`                                     //资源名称
	Description         string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	InstancePort        int      `json:"instancePort,omitempty" bson:"instancePort,omitempty"` //云主机端口
	LoadBalancerPort    int      `json:"loadBalancerPort" bson:"loadBalancerPort"`             //负载均衡器端口
	Protocol            string   `json:"protocol,omitempty" bson:"protocol,omitempty"`         //协议:tcp,HTTP,https,dcp
	CertificateUuid     string   `json:"certificateUuid,omitempty" bson:"certificateUuid,omitempty"`
	ResourceUuid        string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`               //资源UUID。若指定，镜像会使用该字段值作为UUID。
	HealthCheckProtocol string   `json:"healthCheckProtocol,omitempty" bson:"healthCheckProtocol,omitempty"` //健康检查协议
	HealthCheckMethod   string   `json:"healthCheckMethod,omitempty" bson:"healthCheckMethod,omitempty"`     //健康检查方法
	HealthCheckURI      string   `json:"healthCheckURI,omitempty" bson:"healthCheckURI,omitempty"`           //健康检查的URI
	HealthCheckHttpCode string   `json:"healthCheckHttpCode,omitempty" bson:"healthCheckHttpCode,omitempty"` //健康检查期望的返回码
	AclStatus           string   `json:"aclStatus,omitempty" bson:"aclStatus,omitempty"`                     //访问控制策略状态
	AclUuids            []string `json:"aclUuids,omitempty" bson:"aclUuids,omitempty"`                       //访问控制策略组
	AclType             string   `json:"aclType,omitempty" bson:"aclType,omitempty"`                         //访问控制策略类型:white,black
	TagUuids            []string `json:"tagUuids ,omitempty" bson:"tagUuids ,omitempty"`                     //源CIDR; 端口转发规则只作用于源CIDR的流量; 如果忽略不设置, 会默认设置为to 0.0.0.0/0
}

type CreateLoadBalancerListenerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除负载均衡监听器
type DeleteLoadBalancerListenerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteLoadBalancerListenerResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询负载均衡监听器
type QueryLoadBalancerListenerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryLoadBalancerListenerResponse struct {
	Inventories []LoadBalancerListenerInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新负载均衡监听器
type UpdateLoadBalancerListenerRequest struct {
	ReqConfig
	UUID                       string                           `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateLoadBalancerListener UpdateLoadBalancerListenerParams `json:"updateLoadBalancerListener" bson:"updateLoadBalancerListener"`
	SystemTags                 []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateLoadBalancerListenerParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateLoadBalancerListenerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改负载均衡监听器参数
type ChangeLoadBalancerListenerRequest struct {
	ReqConfig
	UUID                       string                           `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateLoadBalancerListener ChangeLoadBalancerListenerParams `json:"updateLoadBalancerListener" bson:"updateLoadBalancerListener"`
	SystemTags                 []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeLoadBalancerListenerParams struct {
	ConnectionIdleTimeout int    `json:"connectionIdleTimeout,omitempty" bson:"connectionIdleTimeout,omitempty"`
	MaxConnection         int    `json:"maxConnection,omitempty" bson:"maxConnection,omitempty"`
	BalancerAlgorithm     string `json:"balancer_algorithm,omitempty" bson:"balancer_algorithm,omitempty"` //roundrobin,leastconn,source
	HealthCheckTarget     string `json:"healthCheckTarget,omitempty" bson:"healthCheckTarget,omitempty"`
	HealthyThreshold      int    `json:"healthyThreshold,omitempty" bson:"healthyThreshold,omitempty"`
	UnhealthyThreshold    int    `json:"unhealthyThreshold,omitempty" bson:"unhealthyThreshold,omitempty"`
	HealthCheckInterval   int    `json:"healthCheckInterval,omitempty" bson:"healthCheckInterval,omitempty"`
	HealthCheckProtocol   string `json:"healthCheckProtocol,omitempty" bson:"healthCheckProtocol,omitempty"` //健康检查协议:tcp,udp,HTTP
	HealthCheckMethod     string `json:"healthCheckMethod,omitempty" bson:"healthCheckMethod,omitempty"`     //健康检查方法:GET,HEAD
	HealthCheckURI        string `json:"healthCheckURI,omitempty" bson:"healthCheckURI,omitempty"`           //健康检查的URI
	HealthCheckHttpCode   string `json:"healthCheckHttpCode,omitempty" bson:"healthCheckHttpCode,omitempty"` //健康检查期望的返回码
	AclStatus             string `json:"aclStatus,omitempty" bson:"aclStatus,omitempty"`                     //访问控制策略状态
}

type ChangeLoadBalancerListenerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机网卡
type GetCandidateVmNicsForLoadBalancerRequest struct {
	ReqConfig
	ListenerUuid string   `json:"listenerUuid" bson:"listenerUuid"`                 //负载均衡监听器UUID
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetCandidateVmNicsForLoadBalancerResponse struct {
	Inventories []VmNicInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加云主机网卡到负载均衡器
type AddVmNicToLoadBalancerRequest struct {
	ReqConfig
	ListenerUuid string                       `json:"listenerUuid" bson:"listenerUuid"` //负载均衡监听器UUID
	Params       AddVmNicToLoadBalancerParams `json:"params" bson:"params"`
	SystemTags   []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddVmNicToLoadBalancerParams struct {
	VmNicUuids []string `json:"vmNicUuids" bson:"vmNicUuids"`
}

type AddVmNicToLoadBalancerResponse struct {
	Inventory LoadBalancerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从负载均衡器移除云主机网卡
type RemoveVmNicFromLoadBalancerRequest struct {
	ReqConfig
	ListenerUuid string   `json:"listenerUuid" bson:"listenerUuid"` //负载均衡监听器UUID
	VmNicUuid    string   `json:"vmNicUuid" bson:"vmNicUuid"`
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveVmNicFromLoadBalancerResponse struct {
	Inventory LoadBalancerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新负载均衡器
type UpdateLoadBalancerRequest struct {
	ReqConfig
	UUID               string                   `json:"uuid" bson:"uuid"`
	UpdateLoadBalancer UpdateLoadBalancerParams `json:"updateLoadBalancer" bson:"updateLoadBalancer"`
	SystemTags         []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateLoadBalancerParams struct {
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type UpdateLoadBalancerResponse struct {
	Inventory LoadBalancerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建证书
type CreateCertificateRequest struct {
	ReqConfig
	Params     CreateCertificateParams `json:"params" bson:"params"`
	SystemTags []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateCertificateParams struct {
	Certificate  string `json:"certificate" bson:"certificate"`
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateCertificateResponse struct {
	Inventory CertificateInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除证书
type DeleteCertificateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteCertificateResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询证书
type QueryCertificateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryCertificateResponse struct {
	Inventories []CertificateInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加证书到负载均衡
type AddCertificateToLoadBalancerListenerRequest struct {
	ReqConfig
	ListenerUuid string                                     `json:"listenerUuid" bson:"listenerUuid"` //负载均衡监听器UUID
	Params       AddCertificateToLoadBalancerListenerParams `json:"params" bson:"params"`
	SystemTags   []string                                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string                                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddCertificateToLoadBalancerListenerParams struct {
	CertificateUuid string `json:"certificateUuid" bson:"certificateUuid"`
}

type AddCertificateToLoadBalancerListenerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从负载均衡移除证书
type RemoveCertificateFromLoadBalancerListenerRequest struct {
	ReqConfig
	ListenerUuid    string   `json:"listenerUuid" bson:"listenerUuid"` //负载均衡监听器UUID
	CertificateUuid string   `json:"certificateUuid" bson:"certificateUuid"`
	SystemTags      []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveCertificateFromLoadBalancerListenerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新证书信息
type UpdateCertificateRequest struct {
	ReqConfig
	UUID              string                  `json:"uuid" bson:"uuid"` //
	UpdateCertificate UpdateCertificateParams `json:"updateCertificate" bson:"updateCertificate"`
	SystemTags        []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateCertificateParams struct {
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type UpdateCertificateResponse struct {
	Inventory CertificateInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加转发规则到访问控制策略组
type AddAccessControlListRedirectRuleRequest struct {
	ReqConfig
	AclUuid    string                                 `json:"aclUuid" bson:"aclUuid"`
	Params     AddAccessControlListRedirectRuleParams `json:"params" bson:"params"`
	SystemTags []string                               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddAccessControlListRedirectRuleParams struct {
	Name         string   `json:"name,omitempty" bson:"name,omitempty"`                 //资源名称
	Description  string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	Domain       string   `json:"domain,omitempty" bson:"domain,omitempty"`             //域名
	Url          string   `json:"url,omitempty" bson:"url,omitempty"`                   //url
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type AddAccessControlListRedirectRuleResponse struct {
	Inventory AccessControlListEntryInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改控制策略组转发规则名称
type ChangeAccessControlListRedirectRuleRequest struct {
	ReqConfig
	Uuid                                string                                    `json:"uuid" bson:"uuid"`
	ChangeAccessControlListRedirectRule ChangeAccessControlListRedirectRuleParams `json:"changeAccessControlListRedirectRule" bson:"changeAccessControlListRedirectRule"`
	SystemTags                          []string                                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                            []string                                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeAccessControlListRedirectRuleParams struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"` //资源名称
}

type ChangeAccessControlListRedirectRuleResponse struct {
	Inventory AccessControlListEntryInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改访问控制组绑定的后端服务器组
type ChangeAccessControlListServerGroupRequest struct {
	ReqConfig
	AclUuid                            string                                   `json:"aclUuid" bson:"aclUuid"`
	ChangeAccessControlListServerGroup ChangeAccessControlListServerGroupParams `json:"changeAccessControlListServerGroup" bson:"changeAccessControlListServerGroup"`
	SystemTags                         []string                                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                           []string                                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeAccessControlListServerGroupParams struct {
	ServerGroupUuids []string `json:"serverGroupUuids" bson:"serverGroupUuids"`
	ListenerUuid     string   `json:"listenerUuid" bson:"listenerUuid"`
}

type ChangeAccessControlListServerGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取负载均衡监听器访问控制策略条目
type GetLoadBalancerListenerACLEntriesRequest struct {
	ReqConfig
	ListenerUuids []string `json:"listenerUuids,omitempty" bson:"listenerUuids,omitempty"`
	Type          string   `json:"type,omitempty" bson:"type,omitempty"`
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetLoadBalancerListenerACLEntriesResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建访问控制策略组
type CreateAccessControlListRequest struct {
	ReqConfig
	Params     CreateAccessControlListParams `json:"params" bson:"params"`
	SystemTags []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateAccessControlListParams struct {
	Name         string   `json:"name" bson:"name"`                                   //资源名称
	Description  string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	IpVersion    string   `json:"ipVersion,omitempty" bson:"ipVersion,omitempty"`
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreateAccessControlListResponse struct {
	Inventory AccessControlListInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除访问控制策略组
type DeleteAccessControlListRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteAccessControlListResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询访问控制策略组
type QueryAccessControlListRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAccessControlListResponse struct {
	Inventories []AccessControlListInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//向访问控制策略组添加IP组
type AddAccessControlListEntryRequest struct {
	ReqConfig
	AclUuid    string                          `json:"aclUuid" bson:"aclUuid"` //资源的UUID，唯一标示该资源
	Params     AddAccessControlListEntryParams `json:"params" bson:"params"`
	SystemTags []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddAccessControlListEntryParams struct {
	Entries      string   `json:"entries" bson:"entries"`
	Description  string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type AddAccessControlListEntryResponse struct {
	Inventory AccessControlListEntryInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除访问控制策略的IP组
type RemoveAccessControlListEntryRequest struct {
	ReqConfig
	AclUuid    string   `json:"aclUuid" bson:"aclUuid"` //资源的UUID，唯一标示该资源
	UUID       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveAccessControlListEntryResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加监听器的访问控制策略
type AddAccessControlListToLoadBalancerRequest struct {
	ReqConfig
	ListenerUuid string                                   `json:"listenerUuid" bson:"listenerUuid"` //负载均衡监听器UUID
	Params       AddAccessControlListToLoadBalancerParams `json:"params" bson:"params"`
	SystemTags   []string                                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string                                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddAccessControlListToLoadBalancerParams struct {
	AclUuids         []string `json:"aclUuids" bson:"aclUuids"`
	AclType          string   `json:"aclType" bson:"aclType"`                   //访问控制策略类型:white,black,redirect
	ServerGroupUuids []string `json:"serverGroupUuids" bson:"serverGroupUuids"` //负载均衡器服务器组uuid
}

type AddAccessControlListToLoadBalancerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除监听器访问控制策略
type RemoveAccessControlListFromLoadBalancerRequest struct {
	ReqConfig
	AclUuids         []string `json:"aclUuids" bson:"aclUuids"`
	ListenerUuid     string   `json:"listenerUuid" bson:"listenerUuid"`                 //负载均衡监听器UUID
	ServerGroupUuids []string `json:"serverGroupUuids" bson:"serverGroupUuids"`         //负载均衡器服务器组uuid
	SystemTags       []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveAccessControlListFromLoadBalancerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取监听器可加载L3网络
type GetCandidateL3NetworksForLoadBalancerRequest struct {
	ReqConfig
	ListenerUuid string   `json:"listenerUuid" bson:"listenerUuid"` //负载均衡监听器UUID
	Limit        int      `json:"limit" bson:"limit"`
	Start        int      `json:"start" bson:"start"`
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetCandidateL3NetworksForLoadBalancerResponse struct {
	Inventories []L3NetworkInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建负载均衡器服务器组
type CreateLoadBalancerServerGroupRequest struct {
	ReqConfig
	LoadBalancerUuid string                              `json:"loadBalancerUuid" bson:"loadBalancerUuid"` //负载均衡器端口
	Params           CreateLoadBalancerServerGroupParams `json:"params" bson:"params"`
	SystemTags       []string                            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string                            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateLoadBalancerServerGroupParams struct {
	Name         string   `json:"name" bson:"name"`                                     //资源名称
	Description  string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreateLoadBalancerServerGroupResponse struct {
	Inventory AccessControlListEntryInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除负载均衡器服务器组
type DeleteLoadBalancerServerGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteLoadBalancerServerGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新负载均衡器服务器组
type UpdateLoadBalancerServerGroupRequest struct {
	ReqConfig
	UUID                          string                              `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateLoadBalancerServerGroup UpdateLoadBalancerServerGroupParams `json:"updateLoadBalancerServerGroup" bson:"updateLoadBalancerServerGroup"`
	SystemTags                    []string                            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                      []string                            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateLoadBalancerServerGroupParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateLoadBalancerServerGroupResponse struct {
	Inventory AccessControlListEntryInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询负载均衡器服务器组
type QueryLoadBalancerServerGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryLoadBalancerServerGroupResponse struct {
	Inventories []AccessControlListEntryInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加后端服务器到服务器组
type AddBackendServerToServerGroupRequest struct {
	ReqConfig
	ServerGroupUuid string                              `json:"serverGroupUuid" bson:"serverGroupUuid"` //负载均衡器服务器组uuid
	Params          AddBackendServerToServerGroupParams `json:"params" bson:"params"`
	SystemTags      []string                            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string                            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddBackendServerToServerGroupParams struct {
	VmNics  []map[string]string `json:"vmNics,omitempty" bson:"vmNics,omitempty"`   //后端服务器网卡
	Servers []map[string]string `json:"servers,omitempty" bson:"servers,omitempty"` //后端服务器IP
}

type AddBackendServerToServerGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从服务器组移除后端服务器
type RemoveBackendServerFromServerGroupRequest struct {
	ReqConfig
	ServerGroupUuid                    string                                   `json:"serverGroupUuid" bson:"serverGroupUuid"` //负载均衡器服务器组uuid
	RemoveBackendServerFromServerGroup RemoveBackendServerFromServerGroupParams `json:"removeBackendServerFromServerGroup" bson:"removeBackendServerFromServerGroup"`
	SystemTags                         []string                                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                           []string                                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveBackendServerFromServerGroupParams struct {
	VmNics  []map[string]string `json:"vmNics,omitempty" bson:"vmNics,omitempty"`   //后端服务器网卡
	Servers []map[string]string `json:"servers,omitempty" bson:"servers,omitempty"` //后端服务器IP
}

type RemoveBackendServerFromServerGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加服务器组到负载均衡监听器
type AddServerGroupToLoadBalancerListenerRequest struct {
	ReqConfig
	ListenerUuid string                              `json:"listenerUuid" bson:"listenerUuid"` //负载均衡器监听器uuid
	Params       AddBackendServerToServerGroupParams `json:"params" bson:"params"`
	SystemTags   []string                            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string                            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddServerGroupToLoadBalancerListenerParams struct {
	ServerGroupUuid string `json:"serverGroupUuid" bson:"serverGroupUuid"` //负载均衡器服务器组uuid
}

type AddServerGroupToLoadBalancerListenerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从负载均衡监听器移除服务器组
type RemoveServerGroupFromLoadBalancerListenerRequest struct {
	ReqConfig
	ListenerUuid    string   `json:"listenerUuid" bson:"listenerUuid"`                 //负载均衡器监听器uuid
	ServerGroupUuid string   `json:"serverGroupUuid" bson:"serverGroupUuid"`           //负载均衡器服务器组uuid
	SystemTags      []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveServerGroupFromLoadBalancerListenerResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新负载均衡后端服务器参数
type ChangeLoadBalancerBackendServerRequest struct {
	ReqConfig
	ServerGroupUuid                 string                                `json:"serverGroupUuid" bson:"serverGroupUuid"` //负载均衡器服务器组uuid
	ChangeLoadBalancerBackendServer ChangeLoadBalancerBackendServerParams `json:"changeLoadBalancerBackendServer" bson:"changeLoadBalancerBackendServer"`
	SystemTags                      []string                              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                        []string                              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeLoadBalancerBackendServerParams struct {
	VmNics  []map[string]string `json:"vmNics,omitempty" bson:"vmNics,omitempty"`   //后端服务器网卡
	Servers []map[string]string `json:"servers,omitempty" bson:"servers,omitempty"` //后端服务器IP
}

type ChangeLoadBalancerBackendServerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取可供负载均衡器服务器组添加的云主机网卡
type GetCandidateVmNicsForLoadBalancerServerGroupRequest struct {
	ReqConfig
	LoadBalancerUuid string   `json:"loadBalancerUuid " bson:"loadBalancerUuid "`
	ServerGroupUuid  string   `json:"serverGroupUuid" bson:"serverGroupUuid"`           //负载均衡器服务器组uuid
	SystemTags       []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetCandidateVmNicsForLoadBalancerServerGroupResponse struct {
	Inventories []VmNicInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建性能独享型负载均衡器组
type CreateSlbGroupRequest struct {
	ReqConfig
	LoadBalancerUuid string               `json:"loadBalancerUuid" bson:"loadBalancerUuid"` //负载均衡器UUID
	Params           CreateSlbGroupParams `json:"params" bson:"params"`
	SystemTags       []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSlbGroupParams struct {
	Name                  string   `json:"name" bson:"name"`                                                       //资源名称
	SlbOfferingUuid       string   `json:"slbOfferingUuid" bson:"slbOfferingUuid"`                                 //高性能实例规格uuid
	FrontEndL3NetworkUuid string   `json:"frontEndL3NetworkUuid" bson:"frontEndL3NetworkUuid"`                     //前端网络uuid
	BackendL3NetworkUuids []string `json:"backendL3NetworkUuids,omitempty" bson:"backendL3NetworkUuids,omitempty"` //后端所属网络uuid
	BackendType           string   `json:"backendType,omitempty" bson:"backendType,omitempty"`                     //后端类型:VYOS
	DeployType            string   `json:"deployType,omitempty" bson:"deployType,omitempty"`                       //部署类型:Direct
	Description           string   `json:"description,omitempty" bson:"description,omitempty"`                     //详细描述
	ResourceUuid          string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`                   //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids              []string `json:"tagUuids ,omitempty" bson:"tagUuids ,omitempty"`                         //标签UUID列表
}

type CreateSlbGroupResponse struct {
	Inventory SlbGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除性能独享型负载均衡器组
type DeleteSlbGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteSlbGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新性能独享型负载均衡器组
type UpdateSlbGroupRequest struct {
	ReqConfig
	UUID           string               `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateSlbGroup UpdateSlbGroupParams `json:"updateLoadBalancerListener" bson:"updateLoadBalancerListener"`
	SystemTags     []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateSlbGroupParams struct {
	Name            string `json:"name,omitempty" bson:"name,omitempty"`                       //资源名称
	SlbOfferingUuid string `json:"slbOfferingUuid,omitempty" bson:"slbOfferingUuid,omitempty"` //高性能实例规格uuid
	Description     string `json:"description,omitempty" bson:"description,omitempty"`         //详细描述
}

type UpdateSlbGroupResponse struct {
	Inventory SlbGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询性能独享型负载均衡器组
type QuerySlbGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySlbGroupResponse struct {
	Inventories []SlbGroupInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建高性能实例性负载均衡器规格
type CreateSlbOfferingRequest struct {
	ReqConfig
	Params     CreateSlbGroupParams `json:"params" bson:"params"`
	SystemTags []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSlbOfferingParams struct {
	ZoneUuid              string   `json:"zoneUuid" bson:"zoneUuid"`                           //区域UUID 若指定，云主机会在指定区域创建。
	ManagementNetworkUuid string   `json:"managementNetworkUuid" bson:"managementNetworkUuid"` //管理L3网络UUID
	ImageUuid             string   `json:"imageUuid" bson:"imageUuid"`                         //云主机镜像UUID
	Name                  string   `json:"name" bson:"name"`                                   //资源名称
	Description           string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	CPUNum                int      `json:"cpuNum" bson:"cpuNum"`                               //cpu数量
	MemorySize            int64    `json:"memorySize" bson:"memorySize"`                       //内存大小
	AllocatorStrategy     string   `json:"allocatorStrategy,omitempty" bson:"allocatorStrategy,omitempty"`
	SortKey               int      `json:"sortKey,omitempty" bson:"sortKey,omitempty"`           //排序键
	Type                  string   `json:"type,omitempty" bson:"type,omitempty"`                 //类型
	ResourceUuid          string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids              []string `json:"tagUuids ,omitempty" bson:"tagUuids ,omitempty"`       //标签UUID列表
}

type CreateSlbOfferingResponse struct {
	Inventory InstanceOfferingInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询性能独享型负载均衡器规格
type QuerySlbOfferingRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySlbOfferingResponse struct {
	Inventories []SlbOfferingInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建高性能负载均衡器实例
type CreateSlbInstanceRequest struct {
	ReqConfig
	Params     CreateSlbGroupParams `json:"params" bson:"params"`
	SystemTags []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSlbInstanceParams struct {
	Name         string   `json:"name" bson:"name"` //资源名称
	SlbGroupUuid string   `json:"slbGroupUuid" bson:"slbGroupUuid"`
	Description  string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	ZoneUuid     string   `json:"zoneUuid,omitempty" bson:"zoneUuid,omitempty"`       //区域UUID 若指定，云主机会在指定区域创建。
	ClusterUUID  string   `json:"clusterUuid,omitempty" bson:"clusterUuid,omitempty"` //集群UUID 若指定，云主机会在指定集群创建，该字段优先级高于zoneUuid。
	HostUuid     string   `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids ,omitempty" bson:"tagUuids ,omitempty"`       //标签UUID列表
}

type CreateSlbInstanceResponse struct {
	Inventory SlbVmInstanceInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type LoadBalancerInventory struct {
	UUID        string                          `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	VipUuid     string                          `json:"vipUuid" bson:"vipUuid"`
	Name        string                          `json:"name" bson:"name"`                                   //资源名称
	Description string                          `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	State       string                          `json:"state" bson:"state"`
	CreateDate  string                          `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string                          `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	Listeners   []LoadBalancerListenerInventory `json:"listeners" bson:"listeners"`
}

type LoadBalancerListenerInventory struct {
	UUID             string            `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name             string            `json:"name" bson:"name"`                                   //资源名称
	Description      string            `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	LoadBalancerUuid string            `json:"loadBalancerUuid" bson:"loadBalancerUuid"`           //负载均衡器UUID
	InstancePort     string            `json:"instancePort" bson:"instancePort"`
	LoadBalancerPort string            `json:"loadBalancerPort" bson:"loadBalancerPort"`
	Protocol         string            `json:"protocol" bson:"protocol"`
	ServerGroupUuid  string            `json:"serverGroupUuid" bson:"serverGroupUuid"`
	CreateDate       string            `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate       string            `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	VmNicRefs        []VmNicRefs       `json:"vmNicRefs" bson:"vmNicRefs"`
	AclRefs          []AclRefs         `json:"aclRefs" bson:"aclRefs"`
	CertificateRefs  []CertificateRefs `json:"certificateRefs" bson:"certificateRefs"`
	ServerGroupRefs  []ServerGroupRefs `json:"serverGroupRefs" bson:"serverGroupRefs"`
}

type VmNicRefs struct {
	Id           int64  `json:"id" bson:"id"` //资源的UUID，唯一标示该资源
	ListenerUuid string `json:"listenerUuid" bson:"listenerUuid"`
	VmNicUuid    string `json:"vmNicUuid" bson:"vmNicUuid"` //云主机网卡UUID
	Status       string `json:"status" bson:"status"`
	CreateDate   string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type AclRefs struct {
	Id           int64  `json:"id" bson:"id"`                     //资源的UUID，唯一标示该资源
	ListenerUuid string `json:"listenerUuid" bson:"listenerUuid"` //监听器唯一标识
	AclUuid      string `json:"aclUuid" bson:"aclUuid"`           //访问策略组唯一标识
	Type         string `json:"type" bson:"type"`                 //访问策略类型
	CreateDate   string `json:"createDate" bson:"createDate"`     //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"`     //最后一次修改时间
}

type CertificateInventory struct {
	UUID        string            `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name        string            `json:"name" bson:"name"` //资源名称
	Certificate string            `json:"certificate" bson:"certificate"`
	Description string            `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	CreateDate  string            `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate  string            `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
	Listeners   []CertificateRefs `json:"listeners" bson:"listeners"`
}

type CertificateRefs struct {
	Id              int64  `json:"id" bson:"id"` //资源的UUID，唯一标示该资源
	ListenerUuid    string `json:"listenerUuid" bson:"listenerUuid"`
	CertificateUuid string `json:"certificateUuid" bson:"certificateUuid"`
	CreateDate      string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type ServerGroupRefs struct {
	Id              int64  `json:"id" bson:"id"` //资源的UUID，唯一标示该资源
	ListenerUuid    string `json:"listenerUuid" bson:"listenerUuid"`
	ServerGroupUuid string `json:"serverGroupUuid" bson:"serverGroupUuid"`
	CreateDate      string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type AccessControlListInventory struct {
	UUID        string                            `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name        string                            `json:"name" bson:"name"` //资源名称
	IpVersion   string                            `json:"ipVersion" bson:"ipVersion"`
	Description string                            `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	CreateDate  string                            `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate  string                            `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
	Entries     []AccessControlListEntryInventory `json:"entries" bson:"entries"`
}

type AccessControlListEntryInventory struct {
	UUID        string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	AclUuid     string `json:"aclUuid" bson:"aclUuid"`
	Type        string `json:"type" bson:"type"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`     //资源名称
	Domain      string `json:"domain,omitempty" bson:"domain,omitempty"` //域名
	Url         string `json:"url,omitempty" bson:"url,omitempty"`       //url
	IpEntries   string `json:"ipEntries" bson:"ipEntries"`
	Description string `json:"description" bson:"description"` //详细描述
	CreateDate  string `json:"createDate" bson:"createDate"`   //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"`   //最后一次修改时间
}

type SlbGroupInventory struct {
	UUID            string                   `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name            string                   `json:"name" bson:"name"`                                   //资源名称
	BackendType     string                   `json:"backendType,omitempty" bson:"backendType,omitempty"` //后端类型:VYOS
	DeployType      string                   `json:"deployType,omitempty" bson:"deployType,omitempty"`   //部署类型:Direct
	SlbOfferingUuid string                   `json:"slbOfferingUuid" bson:"slbOfferingUuid"`             //高性能实例规格uuid
	Description     string                   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	CreateDate      string                   `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate      string                   `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
	SlbVms          []SlbVmInstanceInventory `json:"slbVms" bson:"slbVms"`
	Lbs             []Lbs                    `json:"lbs" bson:"lbs"`
	Networks        []Networks               `json:"networks" bson:"networks"`
}

type SlbVmInstanceInventory struct {
	SlbGroupUuid              string            `json:"slbGroupUuid" bson:"slbGroupUuid"`
	PublicNetworkUuid         string            `json:"publicNetworkUuid" bson:"publicNetworkUuid"`
	VirtualRouterVips         []string          `json:"virtualRouterVips" bson:"virtualRouterVips"`
	ApplianceVmType           string            `json:"applianceVmType" bson:"applianceVmType"`
	ManagementNetworkUuid     string            `json:"managementNetworkUuid" bson:"managementNetworkUuid"`
	DefaultRouteL3NetworkUuid string            `json:"defaultRouteL3NetworkUuid" bson:"defaultRouteL3NetworkUuid"`
	Status                    string            `json:"status" bson:"status"`
	AgentPort                 int               `json:"agentPort" bson:"agentPort"`
	HaStatus                  string            `json:"haStatus" bson:"haStatus"`
	UUID                      string            `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	Name                      string            `json:"name" bson:"name"`                                 //资源名称
	Description               string            `json:"description" bson:"description"`                   //资源的详细描述
	ZoneUUID                  string            `json:"zoneUuid" bson:"zoneUuid"`                         //区域UUID
	ClusterUUID               string            `json:"clusterUuid" bson:"clusterUuid"`                   //集群UUID
	ImageUUID                 string            `json:"imageUuid" bson:"imageUuid"`                       //镜像UUID
	HostUUID                  string            `json:"hostUuid" bson:"hostUuid"`                         //物理机UUID
	LastHostUUID              string            `json:"lastHostUuid" bson:"lastHostUuid"`                 //上一次运行云主机的物理机UUID
	InstanceOfferingUUID      string            `json:"instanceOfferingUuid" bson:"instanceOfferingUuid"` //计算规格UUID
	RootVolumeUUID            string            `json:"rootVolumeUuid" bson:"rootVolumeUuid"`             //根云盘UUID
	Platform                  string            `json:"platform" bson:"platform"`                         //云主机运行平台
	DefaultL3NetworkUUID      string            `json:"defaultL3NetworkUuid" bson:"defaultL3NetworkUuid"` //默认三层网络UUID
	Type                      string            `json:"type" bson:"type"`                                 //云主机类型
	HypervisorType            string            `json:"hypervisorType" bson:"hypervisorType"`             //云主机的hypervisor类型
	MemorySize                int64             `json:"memorySize" bson:"memorySize"`                     //内存大小
	CPUNum                    int               `json:"cpuNum" bson:"cpuNum"`                             //cpu数量
	CPUSpeed                  int64             `json:"cpuSpeed" bson:"cpuSpeed"`                         //cpu主频
	AllocatorStrategy         string            `json:"allocatorStrategy" bson:"allocatorStrategy"`       //分配策略
	CreateDate                string            `json:"createDate" bson:"createDate"`
	LastOpDate                string            `json:"lastOpDate" bson:"lastOpDate"`
	State                     string            `json:"state" bson:"state"`           //云主机的可用状态
	VMNics                    []VmNicInventory  `json:"vmNics" bson:"vmNics"`         //所有网卡信息
	AllVolumes                []VolumeInventory `json:"allVolumes" bson:"allVolumes"` //所有卷
	VmCdRoms                  []VmCdrom         `json:"vmCdRoms" bson:"vmCdRoms"`
}

type Lbs struct {
	SlbGroupUuid string                          `json:"slbGroupUuid" bson:"slbGroupUuid"`
	UUID         string                          `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name         string                          `json:"name" bson:"name"`               //资源名称
	Description  string                          `json:"description" bson:"description"` //详细描述
	State        string                          `json:"state" bson:"state"`
	Type         string                          `json:"type" bson:"type"`
	VipUuid      string                          `json:"vipUuid" bson:"vipUuid"`
	CreateDate   string                          `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate   string                          `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	Listeners    []LoadBalancerListenerInventory `json:"listeners" bson:"listeners"`
}

type Networks struct {
	SlbGroupUuid      string `json:"slbGroupUuid" bson:"slbGroupUuid"`
	L3NetworkUuid     string `json:"l3NetworkUuid" bson:"l3NetworkUuid"`
	L3NetworkCategory string `json:"l3NetworkCategory" bson:"l3NetworkCategory"`
	L3NetworkType     string `json:"l3NetworkType" bson:"l3NetworkType"`
	Type              string `json:"type" bson:"type"`
	CreateDate        string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate        string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type SlbOfferingInventory struct {
	ZoneUuid              string `json:"zoneUuid" bson:"zoneUuid"`                           //区域UUID 若指定，云主机会在指定区域创建。
	ManagementNetworkUuid string `json:"managementNetworkUuid" bson:"managementNetworkUuid"` //管理L3网络UUID
	ImageUuid             string `json:"imageUuid" bson:"imageUuid"`                         //云主机镜像UUID
	UUID                  string `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name                  string `json:"name" bson:"name"`                                   //资源名称
	Description           string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	CPUNum                int    `json:"cpuNum" bson:"cpuNum"`                               //cpu数量
	CpuSpeed              int    `json:"cpuSpeed" bson:"cpuSpeed"`                           //CPU速度
	MemorySize            int64  `json:"memorySize" bson:"memorySize"`                       //内存大小
	Type                  string `json:"type,omitempty" bson:"type,omitempty"`               //类型
	AllocatorStrategy     string `json:"allocatorStrategy,omitempty" bson:"allocatorStrategy,omitempty"`
	SortKey               int    `json:"sortKey,omitempty" bson:"sortKey,omitempty"` //排序键
	State                 string `json:"state" bson:"state"`
	CreateDate            string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate            string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}
