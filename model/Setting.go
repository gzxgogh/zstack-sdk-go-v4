package model

//查询全局设置
type QueryGlobalConfigRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryGlobalConfigResponse struct {
	Inventories []GlobalConfigInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新全局设置
type UpdateGlobalConfigRequest struct {
	ReqConfig
	Name               string                   `json:"name" bson:"name"`         //资源名称
	Category           string                   `json:"category" bson:"category"` //设置类型
	UpdateGlobalConfig UpdateGlobalConfigParams `json:"updateSecurityGroup" bson:"updateSecurityGroup"`
	SystemTags         []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateGlobalConfigParams struct {
	Value string `json:"value" bson:"value"` //默认配额值
}

type UpdateGlobalConfigResponse struct {
	Inventory SecurityGroupInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//重置全局配置
type ResetGlobalConfigRequest struct {
	ReqConfig
	ResetGlobalConfig ResetGlobalConfigParams `json:"resetGlobalConfig" bson:"resetGlobalConfig"`
	SystemTags        []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags          []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ResetGlobalConfigParams struct {
}

type ResetGlobalConfigResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type GlobalConfigInventory struct {
	Name         string `json:"name" bson:"name"`                 //资源名称
	Category     string `json:"category" bson:"category"`         //设置类型
	Description  string `json:"description" bson:"description"`   //详细描述
	DefaultValue string `json:"defaultValue" bson:"defaultValue"` //默认值
	Value        string `json:"value" bson:"value"`               //默认配额值
}

//获取资源的资源高级设置
type GetResourceConfigRequest struct {
	ReqConfig
	Name         string   `json:"name" bson:"name"`         //资源名称
	Category     string   `json:"category" bson:"category"` //设置类型
	ResourceUuid string   `json:"resourceUuid" bson:"resourceUuid"`
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetResourceConfigResponse struct {
	Value            string                    `json:"value" bson:"value"` //默认配额值
	Success          string                    `json:"success" bson:"success"`
	EffectiveConfigs []ResourceConfigInventory `json:"effectiveConfigs" bson:"effectiveConfigs"`
	Error            ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取可配置的资源高级设置
type GetResourceBindableConfigRequest struct {
	ReqConfig
	Category   string   `json:"category" bson:"category"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetResourceBindableConfigResponse struct {
	BindableConfigs []BindableConfigs `json:"bindableConfigs" bson:"bindableConfigs"`
	Error           ErrorCode         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除资源的资源高级设置
type DeleteResourceConfigRequest struct {
	ReqConfig
	Category     string   `json:"category" bson:"category"`                             //设置类型
	Name         string   `json:"name" bson:"name"`                                     //资源名称
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	DeleteMode   string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"`     //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`     //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteResourceConfigResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询资源高级设置
type QueryResourceConfigRequest struct {
	ReqConfig
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryResourceConfigResponse struct {
	Inventories []ResourceConfigInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新资源高级设置
type UpdateResourceConfigRequest struct {
	ReqConfig
	Category             string                     `json:"category" bson:"category"`                             //设置类型
	Name                 string                     `json:"name" bson:"name"`                                     //资源名称
	ResourceUuid         string                     `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	UpdateResourceConfig UpdateResourceConfigParams `json:"updateResourceConfig" bson:"updateResourceConfig"`
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateResourceConfigParams struct {
	Value string `json:"value" bson:"value"` //默认配额值
}

type UpdateResourceConfigResponse struct {
	Inventory ResourceConfigInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type ResourceConfigInventory struct {
	UUID         string `json:"uuid" bson:"uuid"`                                     //资源的UUID，唯一标示该资源
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	ResourceType string `json:"resourceType" bson:"resourceType"`                     //所查询资源的类型（地址范围、三层网络、区域）
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description" bson:"description"`                       //详细描述
	Category     string `json:"category" bson:"category"`                             //设置类型
	Value        string `json:"value" bson:"value"`                                   //默认配额值
	CreateDate   string `json:"createDate" bson:"createDate"`                         //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"`                         //最后一次修改时间
}

type BindableConfigs struct {
	Name              string `json:"name" bson:"name"`               //资源名称
	Description       string `json:"description" bson:"description"` //详细描述
	Category          string `json:"category" bson:"category"`       //设置类型
	BindResourceTypes string `json:"bindResourceTypes" bson:"bindResourceTypes"`
}

//应用模板
type ApplyTemplateConfigRequest struct {
	ReqConfig
	ApplyTemplateConfig ApplyTemplateConfigParams `json:"applyTemplateConfig" bson:"applyTemplateConfig"`
	TemplateUuid        string                    `json:"templateUuid" bson:"templateUuid"`                 //模板Uuid
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ApplyTemplateConfigParams struct {
}

type ApplyTemplateConfigResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询所有模板信息
type QueryGlobalConfigTemplateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryGlobalConfigTemplateResponse struct {
	Inventories []GlobalConfigTemplateInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询模板具体配置
type QueryTemplateConfigRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryTemplateConfigResponse struct {
	Inventories []TemplateConfigInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新模板值
type UpdateTemplateConfigRequest struct {
	ReqConfig
	TemplateUuid         string                     `json:"templateUuid" bson:"templateUuid"` //模板Uuid
	UpdateTemplateConfig UpdateTemplateConfigParams `json:"updateTemplateConfig" bson:"updateTemplateConfig"`
	SystemTags           []string                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags             []string                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateTemplateConfigParams struct {
	Category string `json:"category" bson:"category"` //对应的GlobalConfig配置类型
	Name     string `json:"name" bson:"name"`         //资源名称
	Value    string `json:"value" bson:"value"`       //模板值
}

type UpdateTemplateConfigResponse struct {
	Inventory TemplateConfigInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//重置模板到初始状态
type ResetTemplateConfigRequest struct {
	ReqConfig
	TemplateUuid        string                    `json:"templateUuid" bson:"templateUuid"` //模板Uuid
	ResetTemplateConfig ResetTemplateConfigParams `json:"resetTemplateConfig" bson:"resetTemplateConfig"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ResetTemplateConfigParams struct {
}

type ResetTemplateConfigResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type GlobalConfigTemplateInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`               //资源的UUID，唯一标示该资源
	Type        string `json:"type" bson:"type"`               //模板类型
	Name        string `json:"name" bson:"name"`               //资源名称
	Description string `json:"description" bson:"description"` //详细描述
}

type TemplateConfigInventory struct {
	TemplateUuid string `json:"templateUuid" bson:"templateUuid"`
	Category     string `json:"category" bson:"category"`         //对应的GlobalConfig配置类型
	Name         string `json:"name" bson:"name"`                 //资源名称
	DefaultValue string `json:"defaultValue" bson:"defaultValue"` //默认值
	Value        string `json:"value" bson:"value"`               //模板值
}
