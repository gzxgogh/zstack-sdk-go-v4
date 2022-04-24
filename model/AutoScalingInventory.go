package model

type AutoScalingGroupInventory struct {
	UUID                string   `json:"uuid" bson:"uuid"`                               //资源的UUID，唯一标示该资源
	Name                string   `json:"name" bson:"name"`                               //资源名称
	Description         string   `json:"description" bson:"description"`                 //资源的详细描述
	ScalingResourceType string   `json:"scalingResourceType" bson:"scalingResourceType"` //伸缩组伸缩资源类型，目前只支持云主机 	VmInstance
	State               string   `json:"state" bson:"state"`                             //伸缩组启用状态
	DefaultCooldown     int64    `json:"defaultCooldown" bson:"defaultCooldown"`         //伸缩规则默认冷却时间
	MinResourceSize     int      `json:"minResourceSize" bson:"minResourceSize"`         //伸缩组内云主机最少数量
	MaxResourceSize     int      `json:"maxResourceSize" bson:"maxResourceSize"`         //伸缩组内云主机最多数量
	RemovalPolicy       string   `json:"removalPolicy" bson:"removalPolicy"`             //删除云主机规则:OldestInstance,NewestInstance,OldestScalingConfiguration,MinimumCPUUsageInstance,MinimumMemoryUsageInstance
	CreateDate          string   `json:"createDate" bson:"createDate"`                   //创建时间
	LastOpDate          string   `json:"lastOpDate" bson:"lastOpDate"`                   //最后一次修改时间
	AttachedTemplates   []string `json:"attachedTemplates" bson:"attachedTemplates"`     //挂载的伸缩组云主机模板列表
	SystemTags          []string `json:"systemTags" bson:"systemTags"`
}

type AutoScalingGroupActivityInventory struct {
	UUID                        string `json:"uuid" bson:"uuid"`                                               //资源的UUID，唯一标示该资源
	Name                        string `json:"name" bson:"name"`                                               //资源名称
	ScalingGroupUuid            string `json:"scalingGroupUuid" bson:"scalingGroupUuid"`                       //伸缩组UUID
	ActivityAction              string `json:"activityAction" bson:"activityAction"`                           //伸缩活动类型
	ScalingGroupRuleUuid        string `json:"scalingGroupRuleUuid" bson:"scalingGroupRuleUuid"`               //伸缩规则UUID
	Cause                       string `json:"cause" bson:"cause"`                                             //触发伸缩活动的原因
	Status                      string `json:"status" bson:"status"`                                           //伸缩活动状态
	Description                 string `json:"description" bson:"description"`                                 //活动详细描述
	ActivityActionResultMessage string `json:"activityActionResultMessage" bson:"activityActionResultMessage"` //伸缩活动执行结果
	EndDate                     string `json:"endDate" bson:"endDate"`                                         //伸缩活动执行结束时间
	CreateDate                  string `json:"createDate" bson:"createDate"`                                   //创建时间
	LastOpDate                  string `json:"lastOpDate" bson:"lastOpDate"`                                   //最后一次修改时间
}

type AutoScalingRuleInventory struct {
	UUID             string       `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name             string       `json:"name" bson:"name"` //资源名称
	Type             string       `json:"type,omitempty" bson:"type,omitempty"`
	Description      string       `json:"description" bson:"description"` //活动详细描述
	Cooldown         string       `json:"cooldown,omitempty" bson:"cooldown,omitempty"`
	SystemTags       []string     `json:"systemTags" bson:"systemTags"`
	CreateDate       string       `json:"createDate" bson:"createDate"`             //创建时间
	LastOpDate       string       `json:"lastOpDate" bson:"lastOpDate"`             //最后一次修改时间
	ScalingGroupUuid string       `json:"scalingGroupUuid" bson:"scalingGroupUuid"` //伸缩组UUID
	State            string       `json:"state" bson:"state"`
	Status           string       `json:"status" bson:"status"` //伸缩活动状态
	RuleTriggers     RuleTriggers `json:"ruleTriggers" bson:"ruleTriggers"`
}

type RuleTriggers struct {
	UUID        string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name        string `json:"name" bson:"name"` //资源名称
	Type        string `json:"type,omitempty" bson:"type,omitempty"`
	Description string `json:"description" bson:"description"` //活动详细描述
	RuleUuid    string `json:"RuleUuid" bson:"RuleUuid"`
	State       string `json:"state" bson:"state"`
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间

}

type AutoScalingRuleTriggerInventory struct {
	UUID        string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name        string `json:"name" bson:"name"` //资源名称
	Type        string `json:"type,omitempty" bson:"type,omitempty"`
	Description string `json:"description" bson:"description"` //活动详细描述
	RuleUuid    string `json:"RuleUuid" bson:"RuleUuid"`
	State       string `json:"state" bson:"state"`
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type AutoScalingTemplateInventory struct {
	UUID        string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name        string   `json:"name" bson:"name"` //资源名称
	Type        string   `json:"type,omitempty" bson:"type,omitempty"`
	Description string   `json:"description" bson:"description"` //活动详细描述
	RuleUuid    string   `json:"RuleUuid" bson:"RuleUuid"`
	State       string   `json:"state" bson:"state"`
	CreateDate  string   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	SystemTags  []string `json:"systemTags" bson:"systemTags"`
}

type VmInstanceTemplateInventory struct {
	VmInstanceName                  string   `json:"vmInstanceName" bson:"vmInstanceName"`                                   //云主机名称
	VmInstanceType                  string   `json:"vmInstanceType" bson:"vmInstanceType"`                                   //云主机类型
	VmInstanceDescription           string   `json:"vmInstanceDescription" bson:"vmInstanceDescription"`                     //云主机描述
	VmInstanceOfferingUuid          string   `json:"vmInstanceOfferingUuid" bson:"vmInstanceOfferingUuid"`                   //云主机实例规格
	ImageUuid                       string   `json:"imageUuid" bson:"imageUuid"`                                             //云主机镜像UUID
	L3NetworkUuids                  []string `json:"l3NetworkUuids" bson:"l3NetworkUuids"`                                   //云主机三层网络列表
	RootDiskOfferingUuid            string   `json:"rootDiskOfferingUuid" bson:"rootDiskOfferingUuid"`                       //根云盘规格UUID 如果imageUuid字段指定的镜像类型是ISO，该字段必须指定以确定需要创建的根云盘大小。如果镜像类型是非ISO，该字段无需指定。
	DataDiskOfferingUuids           []string `json:"dataDiskOfferingUuids" bson:"dataDiskOfferingUuids"`                     //云盘规格UUID列表 可指定一个或多个云盘规格UUID（UUID可以重复）为云主机创建一个或多个数据云盘。
	VmInstanceZoneUuid              string   `json:"vmInstanceZoneUuid" bson:"vmInstanceZoneUuid"`                           //云主机所属地区
	VmInstanceClusterUuid           string   `json:"vmInstanceClusterUuid" bson:"vmInstanceClusterUuid"`                     //云主机所属集群
	HostUuid                        string   `json:"hostUuid" bson:"hostUuid"`                                               //物理机UUID 若指定，云主机会在指定物理机创建，该字段优先级高于zoneUuid和clusterUuid。
	PrimaryStorageUuidForRootVolume string   `json:"primaryStorageUuidForRootVolume" bson:"primaryStorageUuidForRootVolume"` //主存储UUID 若指定，云主机的根云盘会在指定主存储创建。
	DefaultL3NetworkUUID            string   `json:"defaultL3NetworkUuid" bson:"defaultL3NetworkUuid"`                       //默认三层网络UUID
	UUID                            string   `json:"uuid" bson:"uuid"`                                                       //资源的UUID，唯一标示该资源
	Name                            string   `json:"name" bson:"name"`                                                       //资源名称
	Type                            string   `json:"type" bson:"type"`                                                       //模板类型
	Description                     string   `json:"description" bson:"description"`                                         //活动详细描述
	State                           string   `json:"state" bson:"state"`                                                     //模板启用状态
	CreateDate                      string   `json:"createDate" bson:"createDate"`                                           //创建时间
	LastOpDate                      string   `json:"lastOpDate" bson:"lastOpDate"`                                           //最后一次修改时间
	SystemTags                      []string `json:"systemTags" bson:"systemTags"`
}

type AutoScalingGroupInstanceInventory struct {
	UUID                     string `json:"uuid" bson:"uuid"`                                         //资源的UUID，唯一标示该资源
	VmInstanceUuid           string `json:"vmInstanceUuid" bson:"vmInstanceUuid"`                     //云主机UUID
	ScalingGroupUuid         string `json:"scalingGroupUuid" bson:"scalingGroupUuid"`                 //伸缩组UUID
	TemplateUuid             string `json:"templateUuid" bson:"templateUuid"`                         //伸缩组云主机模块UUID
	ScalingGroupActivityUuid string `json:"scalingGroupActivityUuid" bson:"scalingGroupActivityUuid"` //伸缩活动UUID
	Status                   string `json:"status" bson:"status"`                                     //云主机在伸缩组里的状态
	HealthStatus             string `json:"healthStatus" bson:"healthStatus"`                         //云主机在伸缩组里的健康状态
	Description              string `json:"description" bson:"description"`                           //资源的详细描述
	CreateDate               string `json:"createDate" bson:"createDate"`                             //创建时间
	LastOpDate               string `json:"lastOpDate" bson:"lastOpDate"`                             //最后一次修改时间
	ProtectionStrategy       string `json:"protectionStrategy" bson:"protectionStrategy"`             //实例保护策略
}
