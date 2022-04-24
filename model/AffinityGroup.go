package model

//创建亲和组
type CreateAffinityGroupRequest struct {
	ReqConfig
	Params     CreateAffinityGroupParams `json:"params" bson:"params"`
	SystemTags []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateAffinityGroupParams struct {
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	Policy       string `json:"policy" bson:"policy"`                                 //亲和组策略
	Type         string `json:"type,omitempty" bson:"type,omitempty"`                 //亲和组类型,目前支持物理机亲和,未来将增加网络亲和、路由器亲和、数据中心或机架亲和等多种类型
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateAffinityGroupResponse struct {
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AffinityGroupInventory `json:"inventory" bson:"inventory"`
}

//删除亲和组
type DeleteAffinityGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteAffinityGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询亲和组
type QueryAffinityGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAffinityGroupResponse struct {
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []AffinityGroupInventory `json:"inventories" bson:"inventories"`
}

//更新亲和组
type UpdateAffinityGroupRequest struct {
	ReqConfig
	UUID                string                    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateAffinityGroup UpdateAffinityGroupParams `json:"updateAffinityGroup" bson:"updateAffinityGroup"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateAffinityGroupParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateAffinityGroupResponse struct {
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AffinityGroupInventory `json:"inventory" bson:"inventory"`
}

//添加云主机到亲和组
type AddVmToAffinityGroupRequest struct {
	ReqConfig
	UUID              string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	AffinityGroupUuid string   `json:"affinityGroupUuid" bson:"affinityGroupUuid"`       //亲和组UUID
	SystemTags        []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddVmToAffinityGroupResponse struct {
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AffinityGroupInventory `json:"inventory" bson:"inventory"`
}

//从亲和组移除云主机
type RemoveVmFromAffinityGroupRequest struct {
	ReqConfig
	UUID              string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	AffinityGroupUuid string   `json:"affinityGroupUuid" bson:"affinityGroupUuid"`       //亲和组UUID
	SystemTags        []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveVmFromAffinityGroupResponse struct {
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AffinityGroupInventory `json:"inventory" bson:"inventory"`
}

//改变亲和组的使用状态
type ChangeAffinityGroupStateRequest struct {
	ReqConfig
	UUID                     string                         `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeAffinityGroupState ChangeAffinityGroupStateParams `json:"changeAffinityGroupState" bson:"changeAffinityGroupState"`
	SystemTags               []string                       `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                 []string                       `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeAffinityGroupStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //enable,disable
}

type ChangeAffinityGroupStateResponse struct {
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AffinityGroupInventory `json:"inventory" bson:"inventory"`
}

//获取可绑定云主机的亲和组
type GetCandidateAffinityGroupForAttachingVmRequest struct {
	ReqConfig
	VmUuid     string   `json:"vmUuid" bson:"vmUuid"`                             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetCandidateAffinityGroupForAttachingVmResponse struct {
	Success     bool                     `json:"success" bson:"success"`
	Inventories []AffinityGroupInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取可绑定亲和组的云主机
type GetCandidateVMForAttachingAffinityGroupRequest struct {
	ReqConfig
	AffinityGroupUuid string   `json:"affinityGroupUuid" bson:"affinityGroupUuid"`       //资源的UUID，唯一标示该资源
	SystemTags        []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetCandidateVMForAttachingAffinityGroupResponse struct {
	Success     bool                  `json:"success" bson:"success"`
	Inventories []VmInstanceInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建云主机时获取可加入非亲和组
type GetCandidateAffinityGroupForCreatingVmRequest struct {
	ReqConfig
	ZoneUuid    string   `json:"zoneUuid" bson:"zoneUuid"`                           //区域UUID
	ClusterUuid string   `json:"clusterUuid,omitempty" bson:"clusterUuid,omitempty"` //集群UUID
	HostUuid    string   `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`       //物理机UUID
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`   //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetCandidateAffinityGroupForCreatingVmResponse struct {
	Success     bool                     `json:"success" bson:"success"`
	Inventories []AffinityGroupInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type AffinityGroupInventory struct {
	UUID        string                `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name        string                `json:"name" bson:"name"`               //资源名称
	Description string                `json:"description" bson:"description"` //资源的详细描述
	Policy      string                `json:"policy" bson:"policy"`           //亲和组策略
	Version     string                `json:"version" bson:"version"`         //亲和组分配算法的版本
	Type        string                `json:"type" bson:"type"`               //亲和组类型, 当前为物理机亲和,未来将增加网络亲和、路由器亲和、数据中心或机架亲和等多种类型
	Appliance   string                `json:"appliance" bson:"appliance"`     //亲和组使用者标识
	State       string                `json:"state" bson:"state"`
	CreateDate  string                `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string                `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	Usages      []AffinityGroupUsages `json:"usages" bson:"usages"`
}

type AffinityGroupUsages struct {
	UUID              string `json:"uuid" bson:"uuid"`                           //资源的UUID，唯一标示该资源
	AffinityGroupUuid string `json:"affinityGroupUuid" bson:"affinityGroupUuid"` //亲和组UUID
	ResourceUuid      string `json:"resourceUuid" bson:"resourceUuid"`           //加入亲和组的资源UUID
	ResourceType      string `json:"resourceType" bson:"resourceType"`           //加入亲和组的资源类型
	CreateDate        string `json:"createDate" bson:"createDate"`               //创建时间
	LastOpDate        string `json:"lastOpDate" bson:"lastOpDate"`               //最后一次修改时间
}
