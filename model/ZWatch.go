package model

//时序数据
//获取所有metric元数据
type GetAllMetricMetadataRequest struct {
	ReqConfig
	Name       string   `json:"name,omitempty" bson:"name,omitempty"`             //资源名称
	Namespace  string   `json:"namespace,omitempty" bson:"namespace,omitempty"`   //详细描述
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetAllMetricMetadataResponse struct {
	Metrics []Metrics `json:"metrics" bson:"metrics"`
	Error   ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type Metrics struct {
	Namespace   string   `json:"namespace" bson:"namespace"`
	Name        string   `json:"name" bson:"name"`                                   //资源名称
	Description string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	LabelNames  []string `json:"labelNames" bson:"labelNames"`
}

//获取metric的标签值
type GetMetricLabelValueRequest struct {
	ReqConfig
	Namespace    string   `json:"namespace" bson:"namespace"`                           //名字空间名称
	MetricName   string   `json:"metricName" bson:"metricName"`                         //监控指标名称
	LabelNames   []string `json:"labelNames" bson:"labelNames"`                         //要获取值得标签名列表
	FilterLabels []string `json:"filterLabels,omitempty" bson:"filterLabels,omitempty"` //标签过滤器列表，例如可以指定标签HostUuid=e47f7145f4cd4fca8e2856038ecdf3e1来选择特定物理机的，labelNames中指定标签的值
	StartTime    int64    `json:"startTime,omitempty" bson:"startTime,omitempty"`       //开始时间
	EndTime      int64    `json:"endTime,omitempty" bson:"endTime,omitempty"`           //结束时间
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`     //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetMetricLabelValueResponse struct {
	Labels []MetricLabelValueLabels `json:"labels" bson:"labels"`
	Error  ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type MetricLabelValueLabels struct {
	VMUuid string `json:"VMUuid" bson:"VMUuid"`
	CPUNum string `json:"CPUNum" bson:"CPUNum"`
}

//获取metric数据
type GetMetricDataRequest struct {
	ReqConfig
	Namespace                string   `json:"namespace" bson:"namespace"`                     //名字空间名称
	MetricName               string   `json:"metricName" bson:"metricName"`                   //监控指标名称
	StartTime                string   `json:"startTime,omitempty" bson:"startTime,omitempty"` //开始时间
	EndTime                  string   `json:"endTime,omitempty" bson:"endTime,omitempty"`     //结束时间
	Period                   int      `json:"period,omitempty" bson:"period,omitempty"`       //数据精度
	Labels                   []string `json:"labels,omitempty" bson:"labels,omitempty"`       //过滤标签
	Functions                []string `json:"functions,omitempty" bson:"functions,omitempty"` //函数列表
	OffsetAheadOfCurrentTime int64    `json:"offsetAheadOfCurrentTime,omitempty" bson:"offsetAheadOfCurrentTime,omitempty"`
	SystemTags               []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                 []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetMetricDataResponse struct {
	MetricData []interface{} `json:"data" bson:"data"`
	Error      ErrorCode     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type MetricData struct {
	Value  string                   `json:"value" bson:"value"`
	Time   int64                    `json:"time" bson:"time"`
	Labels []MetricLabelValueLabels `json:"labels" bson:"labels"`
}

//存入自定义metric数据
type PutMetricDataRequest struct {
	ReqConfig
	Params     PutMetricDataParams `json:"params" bson:"params"`
	SystemTags []string            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type PutMetricDataParams struct {
	Namespace string       `json:"namespace" bson:"namespace"` //名字空间名称
	Data      []MetricData `json:"data" bson:"data"`
}

type PutMetricDataResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//事件
//获取所有event元数据
type GetAllEventMetadataRequest struct {
	ReqConfig
	Namespace  string   `json:"namespace,omitempty" bson:"namespace,omitempty"`   //事件名称
	Name       string   `json:"name,omitempty" bson:"name,omitempty"`             //事件命名空间
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetAllEventMetadataResponse struct {
	Events []EventsMetadataInventory `json:"events" bson:"events"`
	Error  ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type EventsMetadataInventory struct {
	Namespace   string   `json:"namespace" bson:"namespace"`     //事件名称
	Name        string   `json:"name" bson:"name"`               //事件命名空间
	Description string   `json:"description" bson:"description"` //资源的详细描述
	LabelNames  []string `json:"labelNames" bson:"labelNames"`
}

//获取事件
type GetEventDataRequest struct {
	ReqConfig
	StartTime                int64    `json:"startTime,omitempty" bson:"startTime,omitempty"`                               //开始时间
	EndTime                  int64    `json:"endTime,omitempty" bson:"endTime,omitempty"`                                   //结束时间
	Limit                    int      `json:"limit,omitempty" bson:"limit,omitempty"`                                       //最大返回条数
	OffsetAheadOfCurrentTime int64    `json:"offsetAheadOfCurrentTime,omitempty" bson:"offsetAheadOfCurrentTime,omitempty"` //早于当前时间的毫秒数，例如：查询最近一小时的消息，则传入3600000
	Conditions               []string `json:"conditions,omitempty" bson:"conditions,omitempty"`                             //查询条件
	Count                    bool     `json:"count,omitempty" bson:"count,omitempty"`                                       //是否查询事件数量
	EndpointUuid             string   `json:"endpointUuid,omitempty" bson:"endpointUuid,omitempty"`                         //接收端Uuid
	SystemTags               []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`                             //云主机系统标签
	UserTags                 []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetEventDataResponse struct {
	Total   int64                `json:"total" bson:"total"`
	Success bool                 `json:"success" bson:"success"`
	Events  []EventDataInventory `json:"events" bson:"events"`
	Error   ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type EventDataInventory struct {
	Namespace      string                 `json:"namespace" bson:"namespace"`             //事件名称
	Name           string                 `json:"name" bson:"name"`                       //事件命名空间
	Labels         map[string]interface{} `json:"labels" bson:"labels"`                   //标签
	ResourceId     string                 `json:"resourceId" bson:"resourceId"`           //产生事件资源的ID（如果为ZStack资源则为资源UUID）
	ResourceName   string                 `json:"resourceName" bson:"resourceName"`       //资源名称
	Error          string                 `json:"error,omitempty" bson:"error,omitempty"` //如果事件代表错误，该字段为错误详情
	Time           int64                  `json:"time" bson:"time"`                       //事件产生时间
	DataUuid       string                 `json:"dataUuid" bson:"dataUuid"`               //消息UUID
	AccountUuid    string                 `json:"accountUuid" bson:"accountUuid"`         //账户UUID
	EmergencyLevel string                 `json:"emergencyLevel" bson:"emergencyLevel"`
}

//更新事件消息
type UpdateEventDataRequest struct {
	ReqConfig
	UpdateEventData UpdateEventDataParams `json:"updateEventData" bson:"updateEventData"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateEventDataParams struct {
	DataUuid      string `json:"dataUuid,omitempty" bson:"dataUuid,omitempty"`           //消息UUID
	DataStartTime int64  `json:"dataStartTime,omitempty" bson:"dataStartTime,omitempty"` //选择目标消息的起始时间，updateMode需要传入InRange
	DataEndTime   int64  `json:"dataEndTime,omitempty" bson:"dataEndTime,omitempty"`     //选择目标消息的结束时间，updateMode需要传入InRange
	UpdateMode    string `json:"updateMode" bson:"updateMode"`                           //选择更新目标消息的范围，OnlyOne表示只更新指定的消息，InRange表示更新某个时间范围产生的消息， All表示更新所有消息
	ReadStatus    string `json:"readStatus,omitempty" bson:"readStatus,omitempty"`       //更新消息已读状态:Read,unRead
}

type UpdateEventDataResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询事件报警器报警记录
type QueryEventRecordRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryEventRecordResponse struct {
	Inventories []EventRecordInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type EventRecordInventory struct {
	Time             int64  `json:"time" bson:"time"`   //时间
	Count            int64  `json:"count" bson:"count"` //数量
	Id               int64  `json:"id" bson:"id"`
	CreateTime       int64  `json:"createTime" bson:"createTime"`             //创建时间
	Namespace        string `json:"namespace" bson:"namespace"`               //事件名称
	Name             string `json:"name" bson:"name"`                         //事件命名空间
	EmergencyLevel   string `json:"emergencyLevel" bson:"emergencyLevel"`     //报警等级
	ResourceId       string `json:"resourceId" bson:"resourceId"`             //资源UUID
	Error            string `json:"error,omitempty" bson:"error,omitempty"`   //错误信息
	DataUuid         string `json:"dataUuid" bson:"dataUuid"`                 //消息UUID
	AccountUuid      string `json:"accountUuid" bson:"accountUuid"`           //账户UUID
	SubscriptionUuid string `json:"subscriptionUuid" bson:"subscriptionUuid"` //事件报警器UUID
	Labels           string `json:"labels" bson:"labels"`                     //标签
	ReadStatus       bool   `json:"Boolean" bson:"Boolean"`                   //已读状态
}

//报警器
//创建报警器
type CreateAlarmRequest struct {
	ReqConfig
	Params     CreateAlarmParams `json:"params" bson:"params"`
	SystemTags []string          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateAlarmParams struct {
	Name               string                `json:"name" bson:"name"`                                         //资源名称
	Description        string                `json:"description,omitempty" bson:"description,omitempty"`       //详细描述
	ComparisonOperator string                `json:"comparisonOperator" bson:"comparisonOperator"`             //阈值比较符:GreaterThanOrEqualTo,GreaterThan,LessThan,LessThanOrEqualTo
	Period             int                   `json:"period,omitempty" bson:"period,omitempty"`                 //阈值持续时间
	Namespace          string                `json:"namespace" bson:"namespace"`                               //名字空间
	MetricName         string                `json:"metricName" bson:"metricName"`                             //监控指标名称
	Threshold          float64               `json:"threshold" bson:"threshold"`                               //阈值
	RepeatInterval     int                   `json:"repeatInterval,omitempty" bson:"repeatInterval,omitempty"` //报警重复时间
	Labels             []AlarmLabelInventory `json:"labels,omitempty" bson:"labels,omitempty"`                 //标签列表
	Actions            []AlarmActions        `json:"actions,omitempty" bson:"actions,omitempty"`               //报警动作列表
	RepeatCount        int                   `json:"repeatCount,omitempty" bson:"repeatCount,omitempty"`       //报警重复次数
	Type               string                `json:"type,omitempty" bson:"type,omitempty"`                     //报警器类型
	EnableRecovery     bool                  `json:"enableRecovery" bson:"enableRecovery"`                     //开启恢复通知
	EmergencyLevel     string                `json:"emergencyLevel,omitempty" bson:"emergencyLevel,omitempty"` //报警等级
	ResourceUuid       string                `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`     //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateAlarmResponse struct {
	Inventory AlarmInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除报警器
type DeleteAlarmRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteAlarmResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改报警器
type UpdateAlarmRequest struct {
	ReqConfig
	UUID        string            `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateAlarm UpdateAlarmParams `json:"updateAlarm" bson:"updateAlarm"`
	SystemTags  []string          `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string          `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateAlarmParams struct {
	Name               string  `json:"name,omitempty" bson:"name,omitempty"`                             //资源名称
	Description        string  `json:"description,omitempty" bson:"description,omitempty"`               //详细描述
	ComparisonOperator string  `json:"comparisonOperator,omitempty" bson:"comparisonOperator,omitempty"` //阈值比较符:GreaterThanOrEqualTo,GreaterThan,LessThan,LessThanOrEqualTo
	Period             int     `json:"period,omitempty" bson:"period,omitempty"`                         //阈值持续时间
	Threshold          float64 `json:"threshold,omitempty" bson:"threshold,omitempty"`                   //阈值
	RepeatInterval     int     `json:"repeatInterval,omitempty" bson:"repeatInterval,omitempty"`         //报警重复时间
	RepeatCount        int     `json:"repeatCount,omitempty" bson:"repeatCount,omitempty"`               //报警重复次数
	EnableRecovery     bool    `json:"enableRecovery,omitempty" bson:"enableRecovery,omitempty"`         //开启恢复通知
	EmergencyLevel     string  `json:"emergencyLevel,omitempty" bson:"emergencyLevel,omitempty"`         //报警等级
}

type UpdateAlarmResponse struct {
	Inventory SecurityGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改事件报警器
type UpdateSubscribeEventRequest struct {
	ReqConfig
	UUID                 string                     `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateSubscribeEvent UpdateSubscribeEventParams `json:"updateSubscribeEvent" bson:"updateSubscribeEvent"`
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateSubscribeEventParams struct {
	EmergencyLevel string `json:"emergencyLevel,omitempty" bson:"emergencyLevel,omitempty"` //报警等级:Emergent, Important, Normal
}

type UpdateSubscribeEventResponse struct {
	Inventory EventSubscriptionInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//增加报警动作
type AddActionToAlarmRequest struct {
	ReqConfig
	AlarmUuid  string                 `json:"alarmUuid" bson:"alarmUuid"` //资源的UUID，唯一标示该资源
	Params     AddActionToAlarmParams `json:"params" bson:"params"`
	SystemTags []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddActionToAlarmParams struct {
	ActionType string `json:"actionType" bson:"actionType"` //报警动作类型
	ActionUuid string `json:"actionUuid" bson:"actionUuid"` //报警动作UUID
}

type AddActionToAlarmResponse struct {
	Inventory AlarmInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除报警动作
type RemoveActionFromAlarmRequest struct {
	ReqConfig
	AlarmUuid  string   `json:"alarmUuid" bson:"alarmUuid"`
	ActionUuid string   `json:"actionUuid" bson:"actionUuid"`                     //报警动作UUID
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveActionFromAlarmResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//增加标签到报警器
type AddLabelToAlarmRequest struct {
	ReqConfig
	AlarmUuid  string                `json:"alarmUuid" bson:"alarmUuid"` //资源的UUID，唯一标示该资源
	Params     AddLabelToAlarmParams `json:"params" bson:"params"`
	SystemTags []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddLabelToAlarmParams struct {
	Key          string `json:"key" bson:"key"`                                       //标签键
	Value        string `json:"value" bson:"value"`                                   //标签值
	Operator     string `json:"operator" bson:"operator"`                             //操作符:Regex,Equal
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //用户指定的资源uuid
}

type AddLabelToAlarmResponse struct {
	Inventory AlarmInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改报警器标签
type UpdateAlarmLabelRequest struct {
	ReqConfig
	UUID       string                `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Params     AddLabelToAlarmParams `json:"params" bson:"params"`
	SystemTags []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateAlarmLabelParams struct {
	Key      string `json:"key" bson:"key"`           //标签键
	Value    string `json:"value" bson:"value"`       //标签值
	Operator string `json:"operator" bson:"operator"` //操作符:Regex,Equal
}

type UpdateAlarmLabelResponse struct {
	Inventory AlarmLabelInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从报警器上删除标签
type RemoveLabelFromAlarmRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveLabelFromAlarmResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更改报警器状态
type ChangeAlarmStateRequest struct {
	ReqConfig
	UUID       string                `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Params     AddLabelToAlarmParams `json:"params" bson:"params"`
	SystemTags []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeAlarmStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"` //状态事件:disable,enable
}

type ChangeAlarmStateResponse struct {
	Inventory AlarmInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询报警器
type QueryAlarmRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAlarmResponse struct {
	Inventories []AlarmInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取报警器历史
type GetAlarmDataRequest struct {
	ReqConfig
	StartTime           int64    `json:"startTime,omitempty" bson:"startTime,omitempty"`       //开始时间
	EndTime             int64    `json:"endTime,omitempty" bson:"endTime,omitempty"`           //结束时间
	Limit               int      `json:"limit,omitempty" bson:"limit,omitempty"`               //最大返回条数
	Conditions          []string `json:"conditions,omitempty" bson:"conditions,omitempty"`     //查询条件
	Count               bool     `json:"count,omitempty" bson:"count,omitempty"`               //是否查询事件数量
	ExcludeOtherAccount bool     `json:"excludeOtherAccount " bson:"excludeOtherAccount "`     //是否排除非当前账号的消息（只对SystemAdmin生效）
	EndpointUuid        string   `json:"endpointUuid,omitempty" bson:"endpointUuid,omitempty"` //接收端Uuid
	SystemTags          []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`     //云主机系统标签
	UserTags            []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetAlarmDataResponse struct {
	Histories []Histories `json:"histories" bson:"histories"`
	Error     ErrorCode   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新报警器历史消息
type UpdateAlarmDataRequest struct {
	ReqConfig
	DataUuid      string   `json:"dataUuid" bson:"dataUuid"`
	DataStartTime int64    `json:"dataStartTime,omitempty" bson:"dataStartTime,omitempty"` //选择目标消息的起始时间，updateMode需要传入InRange
	DataEndTime   int64    `json:"dataEndTime,omitempty" bson:"dataEndTime,omitempty"`     //选择目标消息的结束时间，updateMode需要传入InRange
	UpdateMode    string   `json:"updateMode" bson:"updateMode"`                           //选择更新目标消息的范围，OnlyOne表示只更新指定的消息，InRange表示更新某个时间范围产生的消息， All表示更新所有消息
	ReadStatus    string   `json:"readStatus,omitempty" bson:"readStatus,omitempty"`       //更新消息已读状态:Read,unRead
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`       //云主机系统标签
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateAlarmDataResponse struct {
	Histories []Histories `json:"histories" bson:"histories"`
	Error     ErrorCode   `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//统计报警消息数量
type GetZWatchAlertHistogramRequest struct {
	ReqConfig
	TableName     string   `json:"tableName" bson:"tableName"`                           //消息类型
	StartTime     int64    `json:"startTime" bson:"startTime"`                           //开始时间
	EndTime       int64    `json:"endTime" bson:"endTime"`                               //结束时间
	IntervalHours string   `json:"intervalHours" bson:"intervalHours"`                   //间隔小时
	GroupColumns  []string `json:"groupColumns,omitempty" bson:"groupColumns,omitempty"` //分组属性
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`     //云主机系统标签
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetZWatchAlertHistogramResponse struct {
	Success    bool                   `json:"success" bson:"success"`
	Histograms []ZWatchAlertHistogram `json:"histograms" bson:"histograms"`
	Error      ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type ZWatchAlertHistogram struct {
	Time  int64 `json:"time" bson:"time"`
	Count int64 `json:"count" bson:"count"`
	Tags  Tags  `json:"tags" bson:"tags"`
}

type Tags struct {
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
}

//查询资源报警器报警记录
type QueryAlarmRecordRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAlarmRecordResponse struct {
	Success    bool                    `json:"success" bson:"success"`
	Histograms []AlarmRecordHistograms `json:"histograms" bson:"histograms"`
	Error      ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type AlarmRecordHistograms struct {
	Id                 string  `json:"id" bson:"id"`
	CreateTime         int64   `json:"createTime" bson:"createTime"`                     //创建时间
	AccountUuid        string  `json:"accountUuid" bson:"accountUuid"`                   //账户UUID
	AlarmName          string  `json:"alarmName" bson:"alarmName"`                       //报警器名称
	AlarmStatus        string  `json:"alarmStatus" bson:"alarmStatus"`                   //报警器状态
	AlarmUuid          string  `json:"alarmUuid" bson:"alarmUuid"`                       //报警器UUID
	ComparisonOperator string  `json:"comparisonOperator" bson:"comparisonOperator"`     //阈值比较符:GreaterThanOrEqualTo,GreaterThan,LessThan,LessThanOrEqualTo
	Context            string  `json:"context" bson:"context"`                           //内容
	DataUuid           string  `json:"dataUuid" bson:"dataUuid"`                         //消息UUID
	EmergencyLevel     string  `json:"emergencyLevel" bson:"emergencyLevel"`             //报警等级
	Labels             string  `json:"labels" bson:"labels"`                             //标签
	MetricName         string  `json:"metricName" bson:"metricName"`                     //监控指标名称
	MetricValue        float64 `json:"metricValue" bson:"metricValue"`                   //监控项当时值
	Namespace          string  `json:"namespace,omitempty" bson:"namespace,omitempty"`   //详细描述
	Period             int     `json:"period,omitempty" bson:"period",omitempty`         //数据精度
	ReadStatus         string  `json:"readStatus,omitempty" bson:"readStatus,omitempty"` //更新消息已读状态:Read,unRead
	ResourceUuid       string  `json:"resourceUuid" bson:"resourceUuid"`                 //资源UUID
	Threshold          float64 `json:"threshold" bson:"threshold"`                       //阈值
}

//订阅事件
type SubscribeEventRequest struct {
	ReqConfig
	Params     SubscribeEventParams `json:"params" bson:"params"`
	SystemTags []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type SubscribeEventParams struct {
	Namespace      string                `json:"namespace" bson:"namespace"`                 //事件名称
	EventName      string                `json:"eventName" bson:"eventName"`                 //事件名
	Actions        []AlarmActions        `json:"actions,omitempty" bson:"actions,omitempty"` //报警动作列表
	Labels         []AlarmLabelInventory `json:"labels,omitempty" bson:"labels,omitempty"`   //标签列表
	Name           string                `json:"name,omitempty" bson:"name,omitempty"`
	EmergencyLevel string                `json:"emergencyLevel,omitempty" bson:"emergencyLevel,omitempty"` //报警等级
	ResourceUuid   string                `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`     //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids       []string              `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type SubscribeEventResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//退订事件
type UnsubscribeEventRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UnsubscribeEventResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询事件订阅
type QueryEventSubscriptionRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryEventSubscriptionResponse struct {
	Inventories []EventSubscriptionInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新事件订阅的标签
type UpdateEventSubscriptionLabelRequest struct {
	ReqConfig
	UUID                         string                             `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateEventSubscriptionLabel UpdateEventSubscriptionLabelParams `json:"updateEventSubscriptionLabel" bson:"updateEventSubscriptionLabel"`
	SystemTags                   []string                           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                     []string                           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateEventSubscriptionLabelParams struct {
	Key      string `json:"key" bson:"key"`
	Operator string `json:"operator" bson:"operator"`
	Value    string `json:"value" bson:"value"`
}

type UpdateEventSubscriptionLabelResponse struct {
	Inventory AlarmLabelInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建报警消息模板
type CreateSNSTextTemplateRequest struct {
	ReqConfig
	Params     CreateSecurityGroupParams `json:"params" bson:"params"`
	SystemTags []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSNSTextTemplateParams struct {
	Name                    string `json:"name" bson:"name"`                                           //资源名称
	Description             string `json:"description,omitempty" bson:"description,omitempty"`         //详细描述
	ApplicationPlatformType string `json:"applicationPlatformType" bson:"applicationPlatformType"`     //SNS应用平台类型
	Template                string `json:"template" bson:"template"`                                   //报警消息模板文本
	RecoveryTemplate        string `json:"recoveryTemplate" bson:"recoveryTemplate"`                   //恢复报警模板文本
	DefaultTemplate         string `json:"defaultTemplate,omitempty" bson:"defaultTemplate,omitempty"` //是否作为默认模板
	ResourceUuid            string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`       //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateSNSTextTemplateResponse struct {
	Inventory SNSTextTemplateInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除报警消息模板
type DeleteSNSTextTemplateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteSNSTextTemplateResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新报警消息模板
type UpdateSNSTextTemplateRequest struct {
	ReqConfig
	UUID                  string                      `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateSNSTextTemplate UpdateSNSTextTemplateParams `json:"updateSNSTextTemplate" bson:"updateSNSTextTemplate"`
	SystemTags            []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags              []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateSNSTextTemplateParams struct {
	Name             string `json:"name,omitempty" bson:"name,omitempty"`                         //资源名称
	Description      string `json:"description,omitempty" bson:"description,omitempty"`           //详细描述
	Template         string `json:"template,omitempty" bson:"template,omitempty"`                 //报警消息模板文本
	RecoveryTemplate string `json:"recoveryTemplate,omitempty" bson:"recoveryTemplate,omitempty"` //恢复报警模板文本
	DefaultTemplate  string `json:"defaultTemplate,omitempty" bson:"defaultTemplate,omitempty"`   //是否作为默认模板
}

type UpdateSNSTextTemplateResponse struct {
	Inventory SNSTextTemplateInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询报警消息模板
type QuerySNSTextTemplateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySNSTextTemplateResponse struct {
	Inventories []SNSTextTemplateInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建SNS监控阿里云短信模板
type CreateAliyunSmsSNSTextTemplateRequest struct {
	ReqConfig
	Params     CreateAliyunSmsSNSTextTemplateParams `json:"params" bson:"params"`
	SystemTags []string                             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateAliyunSmsSNSTextTemplateParams struct {
	Sign                    string   `json:"sign" bson:"sign"`                                             //短信签名名称
	AlarmTemplateCode       string   `json:"alarmTemplateCode" bson:"alarmTemplateCode"`                   //资源报警器模板Code
	EventTemplateCode       string   `json:"eventTemplateCode" bson:"eventTemplateCode"`                   //事件报警器模板Code
	EventTemplate           string   `json:"eventTemplate,omitempty" bson:"eventTemplate,omitempty"`       //事件报警器模板文本
	Name                    string   `json:"name" bson:"name"`                                             //资源名称
	Description             string   `json:"description,omitempty" bson:"description,omitempty"`           //详细描述
	ApplicationPlatformType string   `json:"applicationPlatformType" bson:"applicationPlatformType"`       //SNS应用平台类型
	Template                string   `json:"template" bson:"template"`                                     //报警消息模板文本
	RecoveryTemplate        string   `json:"recoveryTemplate,omitempty" bson:"recoveryTemplate,omitempty"` //恢复报警模板文本
	DefaultTemplate         string   `json:"defaultTemplate,omitempty" bson:"defaultTemplate,omitempty"`   //是否作为默认模板
	ResourceUuid            string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`         //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids                []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`                 //标签UUID列表
}

type CreateAliyunSmsSNSTextTemplateResponse struct {
	Inventory AliyunSmsSNSTextTemplateInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新SNS阿里云短信文本模板
type UpdateAliyunSmsSNSTextTemplateRequest struct {
	ReqConfig
	UUID                           string                               `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateAliyunSmsSNSTextTemplate UpdateAliyunSmsSNSTextTemplateParams `json:"updateAliyunSmsSNSTextTemplate" bson:"updateAliyunSmsSNSTextTemplate"`
	SystemTags                     []string                             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                       []string                             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateAliyunSmsSNSTextTemplateParams struct {
	Sign              string `json:"sign" bson:"sign"`                                             //短信签名名称
	AlarmTemplateCode string `json:"alarmTemplateCode" bson:"alarmTemplateCode"`                   //资源报警器模板Code
	EventTemplateCode string `json:"eventTemplateCode" bson:"eventTemplateCode"`                   //事件报警器模板Code
	EventTemplate     string `json:"eventTemplate,omitempty" bson:"eventTemplate,omitempty"`       //事件报警器模板文本
	Name              string `json:"name" bson:"name"`                                             //资源名称
	Description       string `json:"description,omitempty" bson:"description,omitempty"`           //详细描述
	Template          string `json:"template" bson:"template"`                                     //报警消息模板文本
	RecoveryTemplate  string `json:"recoveryTemplate,omitempty" bson:"recoveryTemplate,omitempty"` //恢复报警模板文本
	DefaultTemplate   string `json:"defaultTemplate,omitempty" bson:"defaultTemplate,omitempty"`   //是否作为默认模板
}

type UpdateAliyunSmsSNSTextTemplateResponse struct {
	Inventory AliyunSmsSNSTextTemplateInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询SNS监控阿里云短信模板
type QueryAliyunSmsSNSTextTemplateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAliyunSmsSNSTextTemplateResponse struct {
	Inventories []AliyunSmsSNSTextTemplateInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取审计数据
type GetAuditDataRequest struct {
	ReqConfig
	StartTime  int64    `json:"startTime,omitempty" bson:"startTime,omitempty"`   //起始时间
	EndTime    int64    `json:"endTime,omitempty" bson:"endTime,omitempty"`       //结束时间
	Limit      int      `json:"limit,omitempty" bson:"limit,omitempty"`           //数量限制
	Labels     []string `json:"labels,omitempty" bson:"labels,omitempty"`         //过滤标签列表
	AuditType  string   `json:"auditType,omitempty" bson:"auditType,omitempty"`   //审计类型，默认为Resource:Login ,Resource
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetAuditDataResponse struct {
	Audits []AuditDataInventory `json:"audits" bson:"audits"`
	Error  ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type AlarmInventory struct {
	UUID               string                `json:"uuid" bson:"uuid"`                                         //资源的UUID，唯一标示该资源
	Name               string                `json:"name" bson:"name"`                                         //资源名称
	Description        string                `json:"description,omitempty" bson:"description,omitempty"`       //详细描述
	Period             int                   `json:"period" bson:"period"`                                     //阈值持续时间
	Namespace          string                `json:"namespace" bson:"namespace"`                               //名字空间
	MetricName         string                `json:"metricName" bson:"metricName"`                             //监控指标名称
	Threshold          float64               `json:"threshold" bson:"threshold"`                               //阈值
	RepeatInterval     int                   `json:"repeatInterval,omitempty" bson:"repeatInterval,omitempty"` //报警重复时间
	EnableRecovery     bool                  `json:"enableRecovery" bson:"enableRecovery"`                     //开启恢复通知
	CreateDate         string                `json:"createDate" bson:"createDate"`                             //创建时间
	LastOpDate         string                `json:"lastOpDate" bson:"lastOpDate"`                             //最后一次修改时间
	ComparisonOperator string                `json:"comparisonOperator" bson:"comparisonOperator"`             //阈值比较符:GreaterThanOrEqualTo,GreaterThan,LessThan,LessThanOrEqualTo
	Status             string                `json:"status" bson:"status"`
	Labels             []AlarmLabelInventory `json:"labels,omitempty" bson:"labels,omitempty"`   //标签列表
	Actions            []AlarmActions        `json:"actions,omitempty" bson:"actions,omitempty"` //报警动作列表
}

type AlarmLabelInventory struct {
	UUID     string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Key      string   `json:"key" bson:"key"`
	Operator Operator `json:"operator" bson:"operator"`
	Value    string   `json:"value" bson:"value"`
}

type Operator struct {
	Op      string `json:"op" bson:"op"`
	Name    string `json:"name" bson:"name"` //资源名称
	Ordinal string `json:"ordinal" bson:"ordinal"`
}

type AlarmActions struct {
	AlarmUuid  string `json:"alarmUuid" bson:"alarmUuid"`
	ActionType string `json:"actionType" bson:"actionType"` //报警动作类型
	ActionUuid string `json:"actionUuid" bson:"actionUuid"` //报警动作UUID
}

type EventSubscriptionInventory struct {
	UUID           string                `json:"uuid" bson:"uuid"`                                         //资源的UUID，唯一标示该资源
	Name           string                `json:"name" bson:"name"`                                         //资源名称
	Namespace      string                `json:"namespace" bson:"namespace"`                               //名字空间
	EventName      string                `json:"eventName" bson:"eventName"`                               //事件名
	CreateDate     string                `json:"createDate" bson:"createDate"`                             //创建时间
	LastOpDate     string                `json:"lastOpDate" bson:"lastOpDate"`                             //最后一次修改时间
	EmergencyLevel string                `json:"emergencyLevel,omitempty" bson:"emergencyLevel,omitempty"` //报警等级:Emergent, Important, Normal
	State          string                `json:"state" bson:"state"`
	Labels         []AlarmLabelInventory `json:"labels,omitempty" bson:"labels,omitempty"`   //标签列表
	Actions        []AlarmActions        `json:"actions,omitempty" bson:"actions,omitempty"` //报警动作列表
}

type Histories struct {
	AlarmUuid    string  `json:"alarmUuid" bson:"alarmUuid"`                           //报警器UUID
	Namespace    string  `json:"namespace" bson:"namespace"`                           //名字空间
	MetricName   string  `json:"metricName" bson:"metricName"`                         //监控指标名称
	AccountUuid  string  `json:"accountUuid" bson:"accountUuid"`                       //账户UUID
	ResourceUuid string  `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	ResourceType string  `json:"resourceType" bson:"resourceType"`                     //资源类型
	AlarmStatus  string  `json:"alarmStatus" bson:"alarmStatus"`                       //报警器状态
	AlarmName    string  `json:"alarmName" bson:"alarmName"`                           //报警器名称
	Threshold    float64 `json:"threshold" bson:"threshold"`                           //报警阈值
	Period       int     `json:"period" bson:"period"`                                 ///阈值持续周期
	Labels       string  `json:"labels" bson:"labels"`                                 //标签列表
	MetricValue  float64 `json:"metricValue" bson:"metricValue"`                       //监控项当时值
	Time         int64   `json:"time" bson:"time"`                                     //报警纪录生成时间
}

type SNSTextTemplateInventory struct {
	UUID                    string `json:"uuid" bson:"uuid"`                                           //资源的UUID，唯一标示该资源
	Name                    string `json:"name" bson:"name"`                                           //资源名称
	Namespace               string `json:"namespace" bson:"namespace"`                                 //名字空间
	ApplicationPlatformType string `json:"applicationPlatformType" bson:"applicationPlatformType"`     //SNS应用平台类型
	Template                string `json:"template" bson:"template"`                                   //报警消息模板文本
	DefaultTemplate         string `json:"defaultTemplate,omitempty" bson:"defaultTemplate,omitempty"` //是否作为默认模板
	CreateDate              string `json:"createDate" bson:"createDate"`                               //创建时间
	LastOpDate              string `json:"lastOpDate" bson:"lastOpDate"`                               //最后一次修改时间
}

type AliyunSmsSNSTextTemplateInventory struct {
	AlarmTemplateCode       string `json:"alarmTemplateCode" bson:"alarmTemplateCode"`                   //资源报警器模板Code
	Sign                    string `json:"sign" bson:"sign"`                                             //短信签名名称
	EventTemplateCode       string `json:"eventTemplateCode" bson:"eventTemplateCode"`                   //事件报警器模板Code
	EventTemplate           string `json:"eventTemplate,omitempty" bson:"eventTemplate,omitempty"`       //事件报警器模板文本
	UUID                    string `json:"uuid" bson:"uuid"`                                             //资源的UUID，唯一标示该资源
	Name                    string `json:"name" bson:"name"`                                             //资源名称
	Description             string `json:"description,omitempty" bson:"description,omitempty"`           //详细描述
	ApplicationPlatformType string `json:"applicationPlatformType" bson:"applicationPlatformType"`       //SNS应用平台类型
	Template                string `json:"template" bson:"template"`                                     //报警消息模板文本
	RecoveryTemplate        string `json:"recoveryTemplate,omitempty" bson:"recoveryTemplate,omitempty"` //恢复报警模板文本
	DefaultTemplate         string `json:"defaultTemplate,omitempty" bson:"defaultTemplate,omitempty"`   //是否作为默认模板
	CreateDate              string `json:"createDate" bson:"createDate"`                                 //创建时间
	LastOpDate              string `json:"lastOpDate" bson:"lastOpDate"`                                 //最后一次修改时间
}

type AuditDataInventory struct {
	ResourceUuid        string `json:"resourceUuid" bson:"resourceUuid"`               //资源UUID
	ResourceType        string `json:"resourceType" bson:"resourceType"`               //资源类型
	ApiName             string `json:"apiName" bson:"apiName"`                         //API名称
	Error               string `json:"error,omitempty" bson:"error,omitempty"`         //错误详情
	OperatorAccountUuid string `json:"operatorAccountUuid" bson:"operatorAccountUuid"` //发起API账号UUID
	Duration            int64  `json:"duration" bson:"duration"`                       //API耗时
	RequestUuid         string `json:"requestUuid" bson:"requestUuid"`                 //API请求UUID
	ResponseUuid        string `json:"responseUuid" bson:"responseUuid"`               //API返回UUID
	SessionUuid         string `json:"sessionUuid" bson:"sessionUuid"`                 //会话UUID
	RequestDump         string `json:"requestDump" bson:"requestDump"`                 //API请求JSON存档
	ResponseDump        string `json:"responseDump" bson:"responseDump"`               //API返回JSON存档
	Time                int64  `json:"time" bson:"time"`                               //记录生成时间
}
