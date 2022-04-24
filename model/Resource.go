package model

//添加资源栈模板
type AddStackTemplateRequest struct {
	ReqConfig
	Params     AddStackTemplateParams `json:"params" bson:"params"`
	SystemTags []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddStackTemplateParams struct {
	Name            string `json:"name" bson:"name"`                                     //资源名称
	Description     string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	Type            string `json:"type,omitempty" bson:"type,omitempty"`                 //模板类型，默认为zstack
	TemplateContent string `json:"templateContent" bson:"templateContent"`               //模板内容，json字符串
	ResourceUuid    string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type AddStackTemplateResponse struct {
	Inventory StackTemplateInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除资源栈模板
type DeleteStackTemplateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteStackTemplateResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询资源栈模板
type QueryStackTemplateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryStackTemplateResponse struct {
	Inventories []StackTemplateInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改资源栈模板
type UpdateStackTemplateRequest struct {
	ReqConfig
	UUID                string                    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateStackTemplate UpdateStackTemplateParams `json:"updateSecurityGroup" bson:"updateSecurityGroup"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateStackTemplateParams struct {
	Name            string `json:"name,omitempty" bson:"name,omitempty"`                       //资源名称
	Description     string `json:"description,omitempty" bson:"description,omitempty"`         //详细描述
	State           bool   `json:"state,omitempty" bson:"state,omitempty"`                     //模板是否可用
	TemplateContent string `json:"templateContent,omitempty" bson:"templateContent,omitempty"` //模板内容，json字符串
}

type UpdateStackTemplateResponse struct {
	Inventory SecurityGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取支持的资源列表
type GetSupportedCloudFormationResourcesRequest struct {
	ReqConfig
	Type       string   `json:"type,omitempty" bson:"type,omitempty"`             //模板类型，默认zstack
	Version    string   `json:"version,omitempty" bson:"version,omitempty"`       //模板版本号
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetSupportedCloudFormationResourcesResponse struct {
	Resources []SupportedCloudFormationResources `json:"resources" bson:"resources"`
	Error     ErrorCode                          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type SupportedCloudFormationResources struct {
	Name       string   `json:"name" bson:"name"`
	Type       string   `json:"type" bson:"type"`
	ActionName string   `json:"actionName" bson:"actionName"`
	Resources  []string `json:"resources" bson:"resources"`
}

//创建资源栈
type CreateResourceStackRequest struct {
	ReqConfig
	Params     CreateSecurityGroupParams `json:"params" bson:"params"`
	SystemTags []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateResourceStackParams struct {
	Name            string `json:"name" bson:"name"`                                           //资源名称
	Description     string `json:"description,omitempty" bson:"description,omitempty"`         //详细描述
	Type            string `json:"type,omitempty" bson:"type,omitempty"`                       //模板类型，默认zstack
	Rollback        bool   `json:"rollback,omitempty" bson:"rollback,omitempty"`               //堆栈创建失败是否回滚，默认回滚
	TemplateContent string `json:"templateContent,omitempty" bson:"templateContent,omitempty"` //堆栈内容，json字符串
	TemplateUuid    string `json:"templateUuid,omitempty" bson:"templateUuid,omitempty"`       //模板UUID 与参数templateContent二选一
	Parameters      string `json:"parameters" bson:"parameters"`                               //参数列表，json字符串
	ResourceUuid    string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`       //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateResourceStackResponse struct {
	Inventory ResourceStackInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//预览资源栈
type PreviewResourceStackRequest struct {
	ReqConfig
	Params     PreviewResourceStackParams `json:"params" bson:"params"`
	SystemTags []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type PreviewResourceStackParams struct {
	Type            string `json:"type,omitempty" bson:"type,omitempty"`                       //模板类型，默认zstack
	UUID            string `json:"uuid" bson:"uuid"`                                           //资源的UUID，唯一标示该资源
	TemplateContent string `json:"templateContent,omitempty" bson:"templateContent,omitempty"` //堆栈内容，json字符串
	Parameters      string `json:"parameters" bson:"parameters"`                               //参数列表，json字符串
}

type PreviewResourceStackResponse struct {
	Inventory PreviewResourceStruct `json:"inventory" bson:"inventory"`
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除资源栈
type DeleteResourceStackRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteResourceStackResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改资源栈
type UpdateResourceStackRequest struct {
	ReqConfig
	UUID                string                    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateResourceStack UpdateResourceStackParams `json:"updateSecurityGroup" bson:"updateSecurityGroup"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateResourceStackParams struct {
	Name            string `json:"name,omitempty" bson:"name,omitempty"`                       //资源名称
	Description     string `json:"description,omitempty" bson:"description,omitempty"`         //详细描述
	Rollback        bool   `json:"rollback,omitempty" bson:"rollback,omitempty"`               //堆栈创建失败是否回滚，默认回滚
	TemplateContent string `json:"templateContent,omitempty" bson:"templateContent,omitempty"` //堆栈内容，json字符串
	Parameters      string `json:"parameters,omitempty" bson:"parameters,omitempty"`           //参数列表，json字符串
}

type UpdateResourceStackResponse struct {
	Inventory ResourceStackInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询资源栈
type QueryResourceStackRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryResourceStackResponse struct {
	Inventories []ResourceStackInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取资源栈内资源列表
type GetResourceFromResourceStackRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetResourceFromResourceStackResponse struct {
	Resources []map[string]interface{} `json:"resources" bson:"resources"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询资源栈内事件列表
type QueryEventFromResourceStackRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryEventFromResourceStackResponse struct {
	Inventories []EventFromResourceStackInventory `json:"Inventories" bson:"Inventories"`
	Error       ErrorCode                         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//重启资源栈
type RestartResourceStackRequest struct {
	ReqConfig
	UUID                 string                     `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	RestartResourceStack RestartResourceStackParams `json:"restartResourceStack" bson:"restartResourceStack"`
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RestartResourceStackParams struct {
}

type RestartResourceStackResponse struct {
	Inventory ResourceStackInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//检查资源栈模板参数
type CheckStackTemplateParametersRequest struct {
	ReqConfig                                     //资源的UUID，唯一标示该资源
	Params     CheckStackTemplateParametersParams `json:"params" bson:"params"`
	SystemTags []string                           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CheckStackTemplateParametersParams struct {
	UUID            string `json:"uuid" bson:"uuid"`                                           //资源的UUID，唯一标示该资源
	Type            string `json:"type,omitempty" bson:"type,omitempty"`                       //模板类型，默认zstack
	TemplateContent string `json:"templateContent,omitempty" bson:"templateContent,omitempty"` //堆栈内容，json字符串
}

type CheckStackTemplateParametersResponse struct {
	Inventory CheckParametesInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从资源编排模板解析成资源关系图
type DecodeStackTemplateRequest struct {
	ReqConfig                            //资源的UUID，唯一标示该资源
	Params     DecodeStackTemplateParams `json:"params" bson:"params"`
	SystemTags []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DecodeStackTemplateParams struct {
	UUID            string `json:"uuid,omitempty" bson:"uuid,omitempty"`                       //资源的UUID，唯一标示该资源
	Type            string `json:"type,omitempty" bson:"type,omitempty"`                       //模板类型，默认zstack
	TemplateContent string `json:"templateContent,omitempty" bson:"templateContent,omitempty"` //堆栈内容，json字符串
	Parameters      string `json:"parameters,omitempty" bson:"parameters,omitempty"`           //参数列表(Json)
	Preparameters   string `json:"preparameters,omitempty" bson:"preparameters,omitempty"`     //预渲染参数列表(Json)
}

type DecodeStackTemplateResponse struct {
	Resources DecodeStackTemplateResourcesInventory `json:"resources" bson:"resources"`
	Error     ErrorCode                             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取资源对应的资源栈
type GetResourceStackFromResourceRequest struct {
	ReqConfig
	ResourceUuid string   `json:"resourceUuid" bson:"resourceUuid"`                 //资源的UUID，唯一标示该资源
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetResourceStackFromResourceResponse struct {
	Stack map[string]interface{} `json:"stack" bson:"stack"`
	Error ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取资源栈中云主机端口监控状态(
type GetResourceStackVmStatusRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetResourceStackVmStatusResponse struct {
	PortStatus map[string]interface{} `json:"portStatus" bson:"portStatus"`
	Error      ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加资源栈中云主机端口监控
type AddResourceStackVmPortMonitorRequest struct {
	ReqConfig
	Params     AddResourceStackVmPortMonitorParams `json:"params" bson:"params"`
	SystemTags []string                            `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                            `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddResourceStackVmPortMonitorParams struct {
	StackUuid      string `json:"stackUuid,omitempty" bson:"stackUuid,omitempty"` //资源栈UUID
	VmInstanceUuid string `json:"vmInstanceUuid" bson:"vmInstanceUuid"`           //云主机UUID
	Port           int    `json:"port" bson:"port"`                               //端口号
}

type AddResourceStackVmPortMonitorResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除资源栈中云主机端口监控
type DeleteResourceStackVmPortMonitorRequest struct {
	ReqConfig
	StackUuid      string   `json:"stackUuid,omitempty" bson:"stackUuid,omitempty"`   //资源栈UUID
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`             //云主机UUID
	Port           int      `json:"port,omitempty" bson:"port,omitempty"`             //端口号
	DeleteMode     string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteResourceStackVmPortMonitorResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type StackTemplateInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Name        string `json:"name" bson:"name"`               //资源名称
	Description string `json:"description" bson:"description"` //资源的详细描述
	Type        string `json:"type" bson:"type"`               //模板类型，默认zstack
	Version     string `json:"version" bson:"version"`         //模板版本号
	State       bool   `json:"state" bson:"state"`             //启用状态
	Content     string `json:"content" bson:"content"`         //模板内容，json字符串
	Md5sum      string `json:"md5sum" bson:"md5sum"`           //content字段内容的md5校验值
	CreateDate  string `json:"createDate" bson:"createDate"`   //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"`   //最后一次修改时间
}

type ResourceStackInventory struct {
	UUID            string `json:"uuid" bson:"uuid"`                       //资源的UUID，唯一标示该资源
	Name            string `json:"name" bson:"name"`                       //资源名称
	Description     string `json:"description" bson:"description"`         //资源的详细描述
	Type            string `json:"type" bson:"type"`                       //模板类型，默认zstack
	Version         string `json:"version" bson:"version"`                 //模板版本号
	TemplateContent string `json:"templateContent" bson:"templateContent"` //堆栈内容，json字符串
	ParamContent    string `json:"paramContent" bson:"paramContent"`       //模板UUID 与参数templateContent二选一
	Status          string `json:"status" bson:"status"`
	Reason          string `json:"reason" bson:"reason"` //不可恢复的理由，如可恢复则为空
	EnableRollback  bool   `json:"enableRollback" bson:"enableRollback"`
	CreateDate      string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type PreviewResourceStruct struct {
	Conditions map[string]interface{} `json:"conditions" bson:"conditions"`
	Actions    PreviewResourceActions `json:"actions" bson:"actions"`
}

type PreviewResourceActions struct {
	ResourceName string                 `json:"resourceName" bson:"resourceName"`
	ActionName   string                 `json:"actionName" bson:"actionName"`
	Round        int                    `json:"round" bson:"round"`
	InDegree     []interface{}          `json:"inDegree" bson:"inDegree"`
	Actions      map[string]interface{} `json:"actions" bson:"actions"`
}

type EventFromResourceStackInventory struct {
	Id           string `json:"id" bson:"id"`
	Description  string `json:"description" bson:"description"`
	Actions      string `json:"actions" bson:"actions"`           //事件名称
	Content      string `json:"content" bson:"content"`           //事件参数列表
	ResourceName string `json:"resourceName" bson:"resourceName"` //资源名称
	ActionStatus string `json:"actionStatus" bson:"actionStatus"` //执行状态
	StackUuid    string `json:"stackUuid" bson:"stackUuid"`       //堆栈UUID
	Duration     string `json:"duration" bson:"duration"`         //事件持续时间
	CreateDate   string `json:"createDate" bson:"createDate"`     //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"`     //最后一次修改时间
}

type CheckParametesInventory struct {
	ParamName             string `json:"paramName" bson:"paramName"`                         //参数名称
	Type                  string `json:"type" bson:"type"`                                   //模板类型，默认zstack
	DefaultValue          string `json:"defaultValue" bson:"defaultValue"`                   //默认值
	Description           string `json:"description" bson:"description"`                     //资源的详细描述
	NoEcho                string `json:"noEcho" bson:"noEcho"`                               //是否在输出中显示
	Label                 string `json:"label" bson:"label"`                                 //前端显示名称
	ConstraintDescription string `json:"constraintDescription" bson:"constraintDescription"` //若校验失败，返回内容
	ResourceType          string `json:"resourceType" bson:"resourceType"`                   //若参数为ZStack资源，返回资源类型，否则返回null
}

type DecodeStackTemplateResourcesInventory struct {
	ResourceName string                 `json:"resourceName" bson:"resourceName"` //资源编排模板中的资源名称
	ResourceType string                 `json:"resourceType" bson:"resourceType"` //ZStack中的资源类型
	DeletePolicy string                 `json:"deletePolicy" bson:"deletePolicy"` //删除策略
	Description  string                 `json:"description" bson:"description"`   //资源的详细描述
	InDegree     []interface{}          `json:"inDegree" bson:"inDegree"`         //依赖的资源列表
	Action       string                 `json:"action" bson:"action"`             //后续的操作行为
	Properties   map[string]interface{} `json:"properties" bson:"properties"`     //操作参数
	Results      map[string]interface{} `json:"results" bson:"results"`           //操作完成后的结果，若还未执行操作，则为空
	Created      bool                   `json:"created" bson:"created"`           //是否己创建
	MockFailed   bool                   `json:"mockFailed" bson:"mockFailed"`     //测试用于mock失败
	Type         string                 `json:"type" bson:"type"`
}
