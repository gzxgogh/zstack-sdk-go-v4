package model

//创建定时器
type CreateSchedulerTriggerRequest struct {
	ReqConfig
	Params     CreateSchedulerTriggerParams `json:"params" bson:"params"`
	SystemTags []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSchedulerTriggerParams struct {
	Name              string `json:"name,omitempty" bson:"name,omitempty"`                           //资源名称
	Description       string `json:"description,omitempty" bson:"description,omitempty"`             //详细描述
	SchedulerInterval int    `json:"schedulerInterval,omitempty" bson:"schedulerInterval,omitempty"` //间隔时间:当简单定时任务执行超过一次时，必须设置间隔时间,简单定时任务永远重复时，必须设置间隔时间
	RepeatCount       int    `json:"repeatCount,omitempty" bson:"repeatCount,omitempty"`
	StartTime         int64  `json:"startTime,omitempty" bson:"startTime,omitempty"` //开始时间
	SchedulerType     string `json:"schedulerType" bson:"schedulerType"`             //simple,cron
	Cron              string `json:"cron,omitempty" bson:"cron,omitempty"`
	ResourceUuid      string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateSchedulerTriggerResponse struct {
	Inventory SchedulerTriggerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除定时器
type DeleteSchedulerTriggerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteSchedulerTriggerResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询定时器
type QuerySchedulerTriggerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySchedulerTriggerResponse struct {
	Inventories []SchedulerTriggerInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新定时器
type UpdateSchedulerTriggerRequest struct {
	ReqConfig
	UUID                   string                       `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateSchedulerTrigger UpdateSchedulerTriggerParams `json:"updateSchedulerTrigger" bson:"updateSchedulerTrigger"`
	SystemTags             []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags               []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateSchedulerTriggerParams struct {
	Name              string `json:"name,omitempty" bson:"name,omitempty"`                           //资源名称
	Description       string `json:"description,omitempty" bson:"description,omitempty"`             //详细描述
	SchedulerInterval int    `json:"schedulerInterval,omitempty" bson:"schedulerInterval,omitempty"` //间隔时间:当简单定时任务执行超过一次时，必须设置间隔时间,简单定时任务永远重复时，必须设置间隔时间
	RepeatCount       int    `json:"repeatCount,omitempty" bson:"repeatCount,omitempty"`
	StartTime         int64  `json:"startTime,omitempty" bson:"startTime,omitempty"` //开始时间
	Cron              string `json:"cron,omitempty" bson:"cron,omitempty"`
}

type UpdateSchedulerTriggerResponse struct {
	Inventory SchedulerTriggerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加定时任务到定时器
type AddSchedulerJobToSchedulerTriggerRequest struct {
	ReqConfig
	SchedulerJobUuid     string                       `json:"schedulerJobUuid" bson:"schedulerJobUuid"`         //定时任务UUID
	SchedulerTriggerUuid string                       `json:"schedulerTriggerUuid" bson:"schedulerTriggerUuid"` //定时器UUID
	Params               UpdateSchedulerTriggerParams `json:"params" bson:"params"`
	SystemTags           []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddSchedulerJobToSchedulerTriggerParams struct {
	TriggerNow bool `json:"triggerNow,omitempty" bson:"triggerNow,omitempty"`
}

type AddSchedulerJobToSchedulerTriggerResponse struct {
	Inventory SchedulerJobSchedulerTriggerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从定时器移除定时任务
type RemoveSchedulerJobFromSchedulerTriggerRequest struct {
	ReqConfig
	SchedulerJobUuid     string   `json:"schedulerJobUuid" bson:"schedulerJobUuid"`         //定时任务UUID
	SchedulerTriggerUuid string   `json:"schedulerTriggerUuid" bson:"schedulerTriggerUuid"` //定时器UUID
	SystemTags           []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveSchedulerJobFromSchedulerTriggerResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建定时任务
type CreateSchedulerJobRequest struct {
	ReqConfig
	Params     CreateSchedulerJobParams `json:"params" bson:"params"`
	SystemTags []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSchedulerJobParams struct {
	Name               string                 `json:"name,omitempty" bson:"name,omitempty"`                 //资源名称
	Description        string                 `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	TargetResourceUuid string                 `json:"targetResourceUuid" bson:"targetResourceUuid"`         //目标资源UUID
	Type               string                 `json:"type" bson:"type"`                                     //startVm,stopVm,rebootVm,volumeSnapshot
	Parameters         map[string]interface{} `json:"parameters,omitempty" bson:"parameters,omitempty"`     //参数列表，json字符串
	ResourceUuid       string                 `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateSchedulerJobResponse struct {
	Inventory SchedulerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除定时任务
type DeleteSchedulerJobRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteSchedulerJobResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询定时任务
type QuerySchedulerJobRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySchedulerJobResponse struct {
	Inventories []SchedulerTriggerInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新定时任务
type UpdateSchedulerJobRequest struct {
	ReqConfig
	UUID               string                   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateSchedulerJob UpdateSchedulerJobParams `json:"updateSchedulerJob" bson:"updateSchedulerJob"`
	SystemTags         []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateSchedulerJobParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`               //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateSchedulerJobResponse struct {
	Inventory SchedulerTriggerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取未挂载定时器的任务
type GetNoTriggerSchedulerJobsRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetNoTriggerSchedulerJobsResponse struct {
	Inventories []SchedulerTriggerInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//改变定时任务状态
type ChangeSchedulerStateRequest struct {
	ReqConfig
	UUID                 string                     `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeSchedulerState ChangeSchedulerStateParams `json:"changeSchedulerState" bson:"changeSchedulerState"`
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeSchedulerStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"`
}

type ChangeSchedulerStateResponse struct {
	Inventory SchedulerTriggerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type SchedulerTriggerInventory struct {
	UUID              string   `json:"uuid" bson:"uuid"`                           //资源的UUID，唯一标示该资源
	Name              string   `json:"name" bson:"name"`                           //资源名称
	Description       string   `json:"description" bson:"description"`             //详细描述
	SchedulerType     string   `json:"schedulerType" bson:"schedulerType"`         //simple,cron
	SchedulerInterval int      `json:"schedulerInterval" bson:"schedulerInterval"` //间隔时间:当简单定时任务执行超过一次时，必须设置间隔时间,简单定时任务永远重复时，必须设置间隔时间
	RepeatCount       int      `json:"repeatCount" bson:"repeatCount"`
	StartTime         int64    `json:"startTime" bson:"startTime"` //开始时间
	StopTime          int64    `json:"stopTime" bson:"stopTime"`
	CreateDate        string   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate        string   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	JobsUuid          []string `json:"jobsUuid" bson:"jobsUuid"`
}

type SchedulerJobSchedulerTriggerInventory struct {
	UUID                 string `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SchedulerJobUuid     string `json:"schedulerJobUuid" bson:"schedulerJobUuid"`         //定时任务UUID
	SchedulerTriggerUuid string `json:"schedulerTriggerUuid" bson:"schedulerTriggerUuid"` //定时器UUID
	JobGroup             string `json:"jobGroup" bson:"jobGroup"`
	TriggerGroup         string `json:"triggerGroup" bson:"triggerGroup"`
	CreateDate           string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate           string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type SchedulerInventory struct {
	UUID               string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	TargetResourceUuid string   `json:"targetResourceUuid" bson:"targetResourceUuid"`
	Name               string   `json:"name" bson:"name"`               //资源名称
	Description        string   `json:"description" bson:"description"` //详细描述
	State              string   `json:"state" bson:"state"`
	CreateDate         string   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate         string   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	TriggersUuid       []string `json:"triggersUuid" bson:"triggersUuid"`
}
