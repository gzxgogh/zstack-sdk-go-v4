package model

//创建弹性伸缩组
type CreateAutoScalingGroupRequest struct {
	ReqConfig
	Params     CreateAutoScalingGroupParams `json:"params" bson:"params"`
	SystemTags []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateAutoScalingGroupParams struct {
	Name                string   `json:"name" bson:"name"`                                     //资源名称
	Description         string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ScalingResourceType string   `json:"scalingResourceType" bson:"scalingResourceType"`       //伸缩组伸缩资源类型，目前只支持云主机 	VmInstance
	MinResourceSize     int      `json:"minResourceSize" bson:"minResourceSize"`               //伸缩组内云主机最少数量
	MaxResourceSize     int      `json:"maxResourceSize" bson:"maxResourceSize"`               //伸缩组内云主机最多数量
	DefaultCooldown     int64    `json:"defaultCooldown" bson:"defaultCooldown"`               //伸缩组规则默认冷却时间
	RemovalPolicy       string   `json:"removalPolicy" bson:"removalPolicy"`                   //删除云主机规则:OldestInstance,NewestInstance,OldestScalingConfiguration,MinimumCPUUsageInstance,MinimumMemoryUsageInstance
	DefaultEnable       bool     `json:"defaultEnable" bson:"defaultEnable"`                   //创建完成后，是否默认启用
	ResourceUuid        string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids            []string `json:"tagUuids " bson:"tagUuids "`                           //标签UUID列表
}

type CreateAutoScalingGroupResponse struct {
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AutoScalingGroupInventory `json:"inventory" bson:"inventory"`
}

//删除弹性伸缩组
type DeleteAutoScalingGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteAutoScalingGroupResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改弹性伸缩组
type UpdateAutoScalingGroupRequest struct {
	ReqConfig
	UUID                   string                       `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateAutoScalingGroup UpdateAutoScalingGroupParams `json:"updateAutoScalingGroup" bson:"updateAutoScalingGroup"`
	SystemTags             []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags               []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateAutoScalingGroupParams struct {
	Name            string `json:"name" bson:"name"`                                           //资源名称
	Description     string `json:"description,omitempty" bson:"description,omitempty"`         //详细描述
	MinResourceSize int    `json:"minResourceSize,omitempty" bson:"minResourceSize,omitempty"` //伸缩组内云主机最少数量
	MaxResourceSize int    `json:"maxResourceSize,omitempty" bson:"maxResourceSize,omitempty"` //伸缩组内云主机最多数量
	RemovalPolicy   string `json:"removalPolicy,omitempty" bson:"removalPolicy,omitempty"`     //删除云主机规则:OldestInstance,NewestInstance,OldestScalingConfiguration,MinimumCPUUsageInstance,MinimumMemoryUsageInstance
}

type UpdateAutoScalingGroupResponse struct {
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AutoScalingGroupInventory `json:"inventory" bson:"inventory"`
}

//查询弹性伸缩组
type QueryAutoScalingGroupRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"` //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAutoScalingGroupResponse struct {
	Error       ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []AutoScalingGroupInventory `json:"inventories" bson:"inventories"`
}

//查询弹性伸缩组活动列表
type QueryAutoScalingGroupActivityRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"` //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAutoScalingGroupActivityResponse struct {
	Error       ErrorCode                           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []AutoScalingGroupActivityInventory `json:"inventories" bson:"inventories"`
}

//创建伸缩组云主机扩容规则
type CreateAutoScalingGroupAddingNewInstanceRuleRequest struct {
	ReqConfig
	Params     CreateAutoScalingGroupAddingNewInstanceRuleParams `json:"params" bson:"params"` //资源的UUID，唯一标示该资源
	SystemTags []string                                          `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string                                          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateAutoScalingGroupAddingNewInstanceRuleParams struct {
	AdjustmentType       string   `json:"adjustmentType" bson:"adjustmentType"`               //资源的UUID，唯一标示该资源
	AdjustmentValue      int      `json:"adjustmentValue" bson:"adjustmentValue"`             //增加大小
	Name                 string   `json:"name" bson:"name"`                                   //资源名称
	Description          string   `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
	AutoScalingGroupUuid string   `json:"autoScalingGroupUuid" bson:"autoScalingGroupUuid"`   //伸缩组UUID
	Type                 string   `json:"type,omitempty" bson:"type,omitempty"`               //伸缩规则类型
	Cooldown             string   `json:"cooldown,omitempty" bson:"cooldown,omitempty"`       //伸缩规则触发后的冷却时间
	ResourceUuid         string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`
	TagUuids             []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"` //标签UUID列表
}

type CreateAutoScalingGroupAddingNewInstanceRuleResponse struct {
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AutoScalingRuleInventory `json:"inventory" bson:"inventory"`
}

//创建伸缩组云主机缩容规则
type CreateAutoScalingGroupRemovalInstanceRuleRequest struct {
	ReqConfig
	Params     CreateAutoScalingGroupRemovalInstanceRuleParams `json:"params" bson:"params"`
	SystemTags []string                                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string                                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateAutoScalingGroupRemovalInstanceRuleParams struct {
	AdjustmentType       string   `json:"adjustmentType" bson:"adjustmentType"`               //资源的UUID，唯一标示该资源
	AdjustmentValue      int      `json:"adjustmentValue" bson:"adjustmentValue"`             //减少数值
	RemovalPolicy        string   `json:"removalPolicy" bson:"removalPolicy"`                 //删除云主机规则:OldestInstance,NewestInstance,OldestScalingConfiguration,MinimumCPUUsageInstance,MinimumMemoryUsageInstance
	Name                 string   `json:"name" bson:"name"`                                   //资源名称
	Description          string   `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
	AutoScalingGroupUuid string   `json:"autoScalingGroupUuid" bson:"autoScalingGroupUuid"`   //伸缩组UUID
	Type                 string   `json:"type,omitempty" bson:"type,omitempty"`               //伸缩规则类型
	Cooldown             string   `json:"cooldown,omitempty" bson:"cooldown,omitempty"`       //伸缩规则触发后的冷却时间
	ResourceUuid         string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`
	TagUuids             []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreateAutoScalingGroupRemovalInstanceRuleResponse struct {
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AutoScalingRuleInventory `json:"inventory" bson:"inventory"`
}

//删除伸缩规则
type DeleteAutoScalingRuleRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteAutoScalingRuleResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改弹性伸缩组规则
type UpdateAutoScalingRuleRequest struct {
	ReqConfig
	UUID                  string                      `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateAutoScalingRule UpdateAutoScalingRuleParams `json:"updateAutoScalingRule" bson:"updateAutoScalingRule"`
	SystemTags            []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags              []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateAutoScalingRuleParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
	Cooldown    string `json:"cooldown,omitempty" bson:"cooldown,omitempty"`       //伸缩规则触发后的冷却时间
}

type UpdateAutoScalingRuleResponse struct {
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AutoScalingRuleInventory `json:"inventory" bson:"inventory"`
}

//修改伸缩组扩容规则
type UpdateAutoScalingGroupAddingNewInstanceRuleRequest struct {
	ReqConfig
	UUID                                        string                                            `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateAutoScalingGroupAddingNewInstanceRule UpdateAutoScalingGroupAddingNewInstanceRuleParams `json:"updateAutoScalingGroupAddingNewInstanceRule" bson:"updateAutoScalingGroupAddingNewInstanceRule"`
	SystemTags                                  []string                                          `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags                                    []string                                          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateAutoScalingGroupAddingNewInstanceRuleParams struct {
	AdjustmentType  string `json:"adjustmentType" bson:"adjustmentType"`               //资源的UUID，唯一标示该资源
	AdjustmentValue int    `json:"adjustmentValue" bson:"adjustmentValue"`             //减少数值
	Name            string `json:"name" bson:"name"`                                   //资源名称
	Description     string `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
	Cooldown        string `json:"cooldown,omitempty" bson:"cooldown,omitempty"`       //伸缩规则触发后的冷却时间
}

type UpdateAutoScalingGroupAddingNewInstanceRuleResponse struct {
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AutoScalingRuleInventory `json:"inventory" bson:"inventory"`
}

//修改伸缩组缩容规则
type UpdateAutoScalingGroupRemovalInstanceRuleRequest struct {
	ReqConfig
	UUID                                      string                                          `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateAutoScalingGroupRemovalInstanceRule UpdateAutoScalingGroupRemovalInstanceRuleParams `json:"updateAutoScalingGroupRemovalInstanceRule" bson:"updateAutoScalingGroupRemovalInstanceRule"`
	SystemTags                                []string                                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags                                  []string                                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateAutoScalingGroupRemovalInstanceRuleParams struct {
	AdjustmentType  string `json:"adjustmentType" bson:"adjustmentType"`               //资源的UUID，唯一标示该资源
	AdjustmentValue int    `json:"adjustmentValue" bson:"adjustmentValue"`             //减少数值
	RemovalPolicy   string `json:"removalPolicy" bson:"removalPolicy"`                 //删除云主机规则:OldestInstance,NewestInstance,OldestScalingConfiguration,MinimumCPUUsageInstance,MinimumMemoryUsageInstance
	Name            string `json:"name" bson:"name"`                                   //资源名称
	Description     string `json:"description,omitempty" bson:"description,omitempty"` //资源的详细描述
	Cooldown        string `json:"cooldown,omitempty" bson:"cooldown,omitempty"`       //伸缩规则触发后的冷却时间
}

type UpdateAutoScalingGroupRemovalInstanceRuleResponse struct {
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AutoScalingRuleInventory `json:"inventory" bson:"inventory"`
}

//查询伸缩规则
type QueryAutoScalingRuleRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"` //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAutoScalingRuleResponse struct {
	Error       ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []AutoScalingRuleInventory `json:"inventories" bson:"inventories"`
}

//创建伸缩规则触发器
type CreateAutoScalingRuleAlarmTriggerRequest struct {
	ReqConfig
	AlarmUuid  string                                  `json:"alarmUuid" bson:"alarmUuid"` //报警UUID
	RuleUuid   string                                  `json:"ruleUuid" bson:"ruleUuid"`   //触发器类型
	Params     CreateAutoScalingRuleAlarmTriggerParams `json:"params" bson:"params"`
	SystemTags []string                                `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string                                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateAutoScalingRuleAlarmTriggerParams struct {
	Name         string   `json:"name" bson:"name"`                                     //资源名称
	Description  string   `json:"description,omitempty" bson:"description,omitempty"`   //资源的详细描述
	TriggerType  string   `bson:"triggerType,omitempty" json:"triggerType,omitempty"`   //触发器类型
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type CreateAutoScalingRuleAlarmTriggerResponse struct {
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AutoScalingRuleTriggerInventory `json:"inventory" bson:"inventory"`
}

//创建伸缩规则定时任务触发器
type CreateAutoScalingRuleSchedulerJobTriggeRequest struct {
	ReqConfig
	SchedulerJobUuid string                                         `json:"schedulerJobUuid" bson:"schedulerJobUuid"` //定时任务Uuid
	RuleUuid         string                                         `json:"ruleUuid" bson:"ruleUuid"`
	Params           CreateAutoScalingRuleSchedulerJobTriggerParams `json:"params" bson:"params"`
	SystemTags       []string                                       `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags         []string                                       `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateAutoScalingRuleSchedulerJobTriggerParams struct {
	Name         string   `json:"name" bson:"name"`                                     //资源名称
	Description  string   `json:"description,omitempty" bson:"description,omitempty"`   //资源的详细描述
	TriggerType  string   `bson:"triggerType,omitempty" json:"triggerType,omitempty"`   //触发器类型
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type CreateAutoScalingRuleSchedulerJobTriggerResponse struct {
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AutoScalingRuleTriggerInventory `json:"inventory" bson:"inventory"`
}

//删除伸缩规则触发器
type DeleteAutoScalingRuleTriggerRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid" bson:"uuid"` //UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteAutoScalingRuleTriggerResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询伸缩规则触发器列表
type QueryAutoScalingRuleTriggerRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"` //UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAutoScalingRuleTriggerResponse struct {
	Error       ErrorCode                         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []AutoScalingRuleTriggerInventory `json:"inventories" bson:"inventories"`
}

//创建伸缩组云主机模块
type CreateAutoScalingVmTemplateRequest struct {
	ReqConfig
	Params     CreateAutoScalingVmTemplateParams `json:"params" bson:"params"`
	SystemTags []string                          `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string                          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateAutoScalingVmTemplateParams struct {
	Name                            string   `json:"name" bson:"name"`                                                                           //资源名称
	Description                     string   `json:"description,omitempty" bson:"description,omitempty"`                                         //活动详细描述
	Type                            string   `json:"type,omitempty" bson:"type,omitempty"`                                                       //模板类型
	VmInstanceName                  string   `json:"vmInstanceName" bson:"vmInstanceName"`                                                       //云主机名称
	VmInstanceOfferingUuid          string   `json:"vmInstanceOfferingUuid" bson:"vmInstanceOfferingUuid"`                                       //云主机实例规则
	ImageUuid                       string   `json:"imageUuid" bson:"imageUuid"`                                                                 //云主机镜像UUID
	L3NetworkUuids                  []string `json:"l3NetworkUuids" bson:"l3NetworkUuids"`                                                       //云主机三层网络列表
	VmInstanceDescription           string   `json:"vmInstanceDescription,omitempty" bson:"vmInstanceDescription,omitempty"`                     //云主机描述
	RootDiskOfferingUuid            string   `json:"rootDiskOfferingUuid,omitempty" bson:"rootDiskOfferingUuid,omitempty"`                       //云主机根云盘规格
	DataDiskOfferingUuids           []string `json:"dataDiskOfferingUuids,omitempty" bson:"dataDiskOfferingUuids,omitempty"`                     //数据盘规格列表
	VmInstanceZoneUuid              string   `json:"vmInstanceZoneUuid,omitempty" bson:"vmInstanceZoneUuid,omitempty"`                           //云主机所属地区
	VmInstanceClusterUuid           string   `json:"vmInstanceClusterUuid,omitempty" bson:"vmInstanceClusterUuid,omitempty"`                     //云主机所属集群
	HostUuid                        string   `json:"hostUuid,omitempty" bson:"hostUuid,omitempty"`                                               //物理机UUID
	PrimaryStorageUuidForRootVolume string   `json:"primaryStorageUuidForRootVolume,omitempty" bson:"primaryStorageUuidForRootVolume,omitempty"` //根云盘主存储UUID
	DefaultL3NetworkUuid            string   `json:"defaultL3NetworkUuid" bson:"defaultL3NetworkUuid"`                                           //云主机默认三层网络
	ResourceUuid                    string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`
	TagUuids                        []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"` //标签UUID列表
}

type CreateAutoScalingVmTemplateResponse struct {
	Error     ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AutoScalingTemplateInventory `json:"inventory" bson:"inventories"`
}

//将云主机模块添加到弹性伸缩组
type AttachAutoScalingTemplateToGroupRequest struct {
	ReqConfig
	UUID       string                 `json:"uuid" bson:"uuid"`
	GroupUuid  string                 `json:"groupUuid" bson:"groupUuid"` //伸缩组UUID
	Params     map[string]interface{} `json:"params" bson:"params"`       //空
	SystemTags []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachAutoScalingTemplateToGroupResponse struct {
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory AutoScalingGroupInventory `json:"inventory" bson:"inventories"`
}

//删除伸缩组模板
type DeleteAutoScalingTemplateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteAutoScalingTemplateResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//卸载伸缩组模板
type DetachAutoScalingTemplateFromGroupRequest struct {
	ReqConfig
	TemplateUuid string   `json:"templateUuid" bson:"templateUuid"` //UUID
	GroupUuid    string   `json:"groupUuid" bson:"groupUuid"`       //伸缩组UUID
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachAutoScalingTemplateFromGroupResponse struct {
	Inventory AutoScalingGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询伸缩组云主机模板
type QueryAutoScalingVmTemplateRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"` //UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAutoScalingVmTemplateResponse struct {
	Error       ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []VmInstanceTemplateInventory `json:"inventories" bson:"inventories"`
}

//手动删除弹性伸缩组内云主机
type DeleteAutoScalingGroupInstanceRequest struct {
	ReqConfig
	InstanceUuid string   `json:"instanceUuid" bson:"instanceUuid"` //模板UUID
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteAutoScalingGroupInstanceResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询弹性伸缩组内云主机列表
type QueryAutoScalingGroupInstanceRequest struct {
	ReqConfig
	Uuid       string   `json:"uuid,omitempty" bson:"uuid,omitempty"` //UUID
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAutoScalingGroupInstanceResponse struct {
	Error       ErrorCode                           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []AutoScalingGroupInstanceInventory `json:"inventories" bson:"inventories"`
}

//修改弹性伸缩组启用状态
type ChangeAutoScalingGroupStateRequest struct {
	ReqConfig
	Uuid                        string                            `json:"uuid" bson:"uuid"` //UUID
	ChangeAutoScalingGroupState ChangeAutoScalingGroupStateParams `json:"changeAutoScalingGroupState" bson:"changeAutoScalingGroupState"`
	SystemTags                  []string                          `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags                    []string                          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeAutoScalingGroupStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"`
}

type ChangeAutoScalingGroupStateResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//手动执行伸缩组规则
type ExecuteAutoScalingRuleRequest struct {
	ReqConfig
	Uuid                   string                       `json:"uuid" bson:"uuid"`                                                         //UUID
	ExecuteAutoScalingRule ExecuteAutoScalingRuleParams `json:"executeAutoScalingRule,omitempty" bson:"executeAutoScalingRule,omitempty"` //空
	SystemTags             []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags               []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ExecuteAutoScalingRuleParams struct {
}

type ExecuteAutoScalingRuleResponse struct {
	Error               ErrorCode `json:"error,omitempty" bson:"error,omitempty"`         //错误信息
	Success             bool      `json:"success" bson:"success"`                         //是否成功
	ScalingActivityUuid string    `json:"scalingActivityUuid" bson:"scalingActivityUuid"` //伸缩活动UUID
}

//更新伸缩组实例信息
type UpdateAutoScalingGroupInstanceRequest struct {
	ReqConfig
	InstanceUuid                   string                               `json:"instanceUuid" bson:"instanceUuid"` //	伸缩组内实例UUID
	UpdateAutoScalingGroupInstance UpdateAutoScalingGroupInstanceParams `json:"updateAutoScalingGroupInstance" bson:"updateAutoScalingGroupInstance"`
	SystemTags                     []string                             `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags                       []string                             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateAutoScalingGroupInstanceParams struct {
	GroupUuid          string `json:"groupUuid" bson:"groupUuid"`                   //伸缩组UUID
	ProtectionStrategy string `json:"protectionStrategy" bson:"protectionStrategy"` //伸缩组内实例保护策略
}

type UpdateAutoScalingGroupInstanceResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新伸缩组云主机模板
type UpdateAutoScalingVmTemplateRequest struct {
	ReqConfig
	Uuid                        string                            `json:"uuid" bson:"uuid"` //UUID
	UpdateAutoScalingVmTemplate UpdateAutoScalingVmTemplateParams `json:"updateAutoScalingVmTemplate" bson:"updateAutoScalingVmTemplate"`
	SystemTags                  []string                          `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags                    []string                          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateAutoScalingVmTemplateParams struct {
	Name                   string `json:"name" bson:"name"`                                     //资源名称
	Description            string `json:"description" bson:"description"`                       //活动详细描述
	VmInstanceName         string `json:"vmInstanceName" bson:"vmInstanceName"`                 //云主机名称
	VmInstanceDescription  string `json:"vmInstanceDescription" bson:"vmInstanceDescription"`   //云主机描述
	VmInstanceOfferingUuid string `json:"vmInstanceOfferingUuid" bson:"vmInstanceOfferingUuid"` //云主机实例规格
	ImageUuid              string `json:"imageUuid" bson:"imageUuid"`                           //云主机镜像UUID
}

type UpdateAutoScalingVmTemplateResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}
