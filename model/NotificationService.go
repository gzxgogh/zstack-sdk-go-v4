package model

//创建主题
type CreateSNSTopicRequest struct {
	ReqConfig
	Params     CreateSNSTopicParams `json:"params" bson:"params"`
	SystemTags []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSNSTopicParams struct {
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateSNSTopicResponse struct {
	Inventory SNSTopicInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除主题
type DeleteSNSTopicRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteSNSTopicResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更改主题状态
type ChangeSNSTopicStateRequest struct {
	ReqConfig
	UUID                string                    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeSNSTopicState ChangeSNSTopicStateParams `json:"changeSNSTopicState" bson:"changeSNSTopicState"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeSNSTopicStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"`
}

type ChangeSNSTopicStateResponse struct {
	Inventory SNSTopicInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新主题
type UpdateSNSTopicRequest struct {
	ReqConfig
	UUID           string               `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateSNSTopic UpdateSNSTopicParams `json:"updateSNSTopic" bson:"updateSNSTopic"`
	SystemTags     []string             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateSNSTopicParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateSNSTopicResponse struct {
	Inventory SNSTopicInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询主题
type QuerySNSTopicRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySNSTopicResponse struct {
	Inventories []SNSTopicInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建邮件服务器
type CreateSNSEmailPlatformRequest struct {
	ReqConfig
	Params     CreateSNSEmailPlatformParams `json:"params" bson:"params"`
	SystemTags []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSNSEmailPlatformParams struct {
	SsmtpServer  string `json:"smtpServer" bson:"smtpServer"`                         //SMTP服务器地址
	SmtpPort     string `json:"smtpPort" bson:"smtpPort"`                             //SMTP端口
	Username     string `json:"username" bson:"username"`                             //用户名
	Password     string `json:"password,omitempty" bson:"password,omitempty"`         //用户密码
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	EncryptType  string `json:"encryptType,omitempty" bson:"encryptType,omitempty"`   //SSL,STARTTLS,NONE
}

type CreateSNSEmailPlatformResponse struct {
	Inventory SNSApplicationPlatformInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//测试邮件服务器
type ValidateSNSEmailPlatformRequest struct {
	ReqConfig
	UUID                     string                         `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ValidateSNSEmailPlatform ValidateSNSEmailPlatformParams `json:"validateSNSEmailPlatform" bson:"validateSNSEmailPlatform"`
	SystemTags               []string                       `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                 []string                       `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ValidateSNSEmailPlatformParams struct {
}

type ValidateSNSEmailPlatformResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除SNS应用平台
type DeleteSNSApplicationPlatformRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteSNSApplicationPlatformResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询邮件服务器
type QuerySNSEmailPlatformRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySNSEmailPlatformResponse struct {
	Inventories []SNSApplicationPlatformInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新SNS应用平台
type UpdateSNSApplicationPlatformRequest struct {
	ReqConfig
	UUID                         string                             `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateSNSApplicationPlatform UpdateSNSApplicationPlatformParams `json:"updateSNSApplicationPlatform" bson:"updateSNSApplicationPlatform"`
	SystemTags                   []string                           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                     []string                           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateSNSApplicationPlatformParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateSNSApplicationPlatformResponse struct {
	Inventory SNSApplicationPlatformInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询SNS应用平台
type QuerySNSApplicationPlatformRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySNSApplicationPlatformResponse struct {
	Inventories []SNSApplicationPlatformInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更改SNS应用平台状态
type ChangeSNSApplicationPlatformStateRequest struct {
	ReqConfig
	UUID                              string                                  `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeSNSApplicationPlatformState ChangeSNSApplicationPlatformStateParams `json:"changeSNSApplicationPlatformState" bson:"changeSNSApplicationPlatformState"`
	SystemTags                        []string                                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                          []string                                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeSNSApplicationPlatformStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"`
}

type ChangeSNSApplicationPlatformStateResponse struct {
	Inventory SNSApplicationPlatformInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建邮件接收端
type CreateSNSEmailEndpointRequest struct {
	ReqConfig
	UUID                   string                       `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	CreateSNSEmailEndpoint CreateSNSEmailEndpointParams `json:"createSNSEmailEndpoint" bson:"createSNSEmailEndpoint"`
	SystemTags             []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags               []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSNSEmailEndpointParams struct {
	Email        string   `json:"email" bson:"email"`
	Emails       []string `json:"emails,omitempty" bson:"emails,omitempty"`
	Name         string   `json:"name" bson:"name"`                                   //资源名称
	Description  string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	PlatformUuid string   `json:"platformUuid,omitempty" bson:"platformUuid,omitempty"`
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type CreateSNSEmailEndpointResponse struct {
	Inventory SNSApplicationEndpointInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询邮件接收端
type QuerySNSEmailEndpointRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySNSEmailEndpointResponse struct {
	Inventories []SNSApplicationEndpointInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建HTTP接收端
type CreateSNSHttpEndpointRequest struct {
	ReqConfig
	Params     CreateSNSHttpEndpointParams `json:"params" bson:"params"`
	SystemTags []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSNSHttpEndpointParams struct {
	Url          string `json:"url" bson:"url"`
	Name         string `json:"name" bson:"name"`                                   //资源名称
	Username     string `json:"username,omitempty" bson:"username,omitempty"`       //用户名
	Password     string `json:"password,omitempty" bson:"password,omitempty"`       //用户密码
	Description  string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	PlatformUuid string `json:"platformUuid,omitempty" bson:"platformUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateSNSHttpEndpointResponse struct {
	Inventory SNSApplicationEndpointInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询HTTP接收端
type QuerySNSHttpEndpointRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySNSHttpEndpointResponse struct {
	Inventories []SNSApplicationEndpointInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建钉钉接收端
type CreateSNSDingTalkEndpointRequest struct {
	ReqConfig
	Params     CreateSNSDingTalkEndpointParams `json:"params" bson:"params"`
	SystemTags []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSNSDingTalkEndpointParams struct {
	Url                  string   `json:"url" bson:"url"`
	AtAll                bool     `json:"atAll,omitempty" bson:"atAll,omitempty"`                               //是否消息@所有人
	AtPersonPhoneNumbers []string `json:"atPersonPhoneNumbers,omitempty" bson:"atPersonPhoneNumbers,omitempty"` //要@用户的电话号码
	Name                 string   `json:"name" bson:"name"`                                                     //资源名称
	Description          string   `json:"description,omitempty" bson:"description,omitempty"`                   //详细描述
	PlatformUuid         string   `json:"platformUuid,omitempty" bson:"platformUuid,omitempty"`
	ResourceUuid         string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateSNSDingTalkEndpointResponse struct {
	Inventory SNSDingTalkEndpointInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加钉钉用户
type AddSNSDingTalkAtPersonRequest struct {
	ReqConfig
	Params     AddSNSDingTalkAtPersonParams `json:"params" bson:"params"`
	SystemTags []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddSNSDingTalkAtPersonParams struct {
	PhoneNumber  string `json:"phoneNumber" bson:"phoneNumber"`                       //用户电话号码（钉钉用户以电话注册）
	EndpointUuid string `json:"endpointUuid" bson:"endpointUuid"`                     //钉钉终端UUID
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type AddSNSDingTalkAtPersonResponse struct {
	Inventory SNSDingTalkAtPersonInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除钉钉用户
type RemoveSNSDingTalkAtPersonRequest struct {
	ReqConfig
	PhoneNumber  string   `json:"phoneNumber" bson:"phoneNumber"`                   //用户电话号码（钉钉用户以电话注册）
	EndpointUuid string   `json:"endpointUuid" bson:"endpointUuid"`                 //钉钉终端UUID
	DeleteMode   string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveSNSDingTalkAtPersonResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询钉钉终端
type QuerySNSDingTalkEndpointRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySNSDingTalkEndpointResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除接收端
type DeleteSNSApplicationEndpointRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteSNSApplicationEndpointResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新接收端
type UpdateSNSApplicationEndpointRequest struct {
	ReqConfig
	UUID                         string                             `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateSNSApplicationEndpoint UpdateSNSApplicationEndpointParams `json:"updateSNSApplicationEndpoint" bson:"updateSNSApplicationEndpoint"`
	SystemTags                   []string                           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                     []string                           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateSNSApplicationEndpointParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateSNSApplicationEndpointResponse struct {
	Inventory SNSApplicationEndpointInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询接收端
type QuerySNSApplicationEndpointRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySNSApplicationEndpointResponse struct {
	Inventories []SNSApplicationEndpointInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                         `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更改接收端
type ChangeSNSApplicationEndpointStateRequest struct {
	ReqConfig
	UUID                              string                                  `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeSNSApplicationEndpointState ChangeSNSApplicationEndpointStateParams `json:"changeSNSApplicationEndpointState" bson:"changeSNSApplicationEndpointState"`
	SystemTags                        []string                                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                          []string                                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeSNSApplicationEndpointStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"`
}

type ChangeSNSApplicationEndpointStateResponse struct {
	Inventory SNSApplicationEndpointInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建阿里云短信接收端
type CreateSNSAliyunSmsEndpointRequest struct {
	ReqConfig
	Params     CreateSNSAliyunSmsEndpointParams `json:"params" bson:"params"`
	SystemTags []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSNSAliyunSmsEndpointParams struct {
	AccessKeyUuid string   `json:"accessKeyUuid" bson:"accessKeyUuid"`                 //阿里云AccessKey信息Uuid
	Receivers     []string `json:"receivers" bson:"receivers"`                         //短信接收者
	Name          string   `json:"name" bson:"name"`                                   //资源名称
	Description   string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	PlatformUuid  string   `json:"platformUuid,omitempty" bson:"platformUuid,omitempty"`
	ResourceUuid  string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids      []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type CreateSNSAliyunSmsEndpointResponse struct {
	Inventory SNSSmsEndpointInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//验证阿里云短信接收端
type ValidateSNSAliyunSmsEndpointRequest struct {
	ReqConfig
	UUID                         string                             `json:"uuid" bson:"uuid"`
	ValidateSNSAliyunSmsEndpoint ValidateSNSAliyunSmsEndpointParams `json:"validateSNSAliyunSmsEndpoint" bson:"validateSNSAliyunSmsEndpoint"`
	SystemTags                   []string                           `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                     []string                           `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ValidateSNSAliyunSmsEndpointParams struct {
	PhoneNumbers []string `json:"phoneNumbers" bson:"phoneNumbers"` //短信接收者
}

type ValidateSNSAliyunSmsEndpointResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加短信接收者
type AddSNSSmsReceiverRequest struct {
	ReqConfig
	Params     AddSNSSmsReceiverParams `json:"params" bson:"params"`
	SystemTags []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddSNSSmsReceiverParams struct {
	PhoneNumber  string   `json:"phoneNumber" bson:"phoneNumber"`                       //
	EndpointUuid string   `json:"endpointUuid" bson:"endpointUuid"`                     //短信接收端Uuid
	Type         string   `json:"type" bson:"type"`                                     //短信接收端类型:AliyunSms
	Description  string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type AddSNSSmsReceiverResponse struct {
	Inventory SNSSmsReceiverInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除短信接收者
type RemoveSNSSmsReceiverRequest struct {
	ReqConfig
	PhoneNumber  string   `json:"phoneNumber" bson:"phoneNumber"`                   //
	EndpointUuid string   `json:"endpointUuid" bson:"endpointUuid"`                 //短信接收端Uuid
	DeleteMode   string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveSNSSmsReceiverResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询短信接收端
type QuerySNSSmsEndpointRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySNSSmsEndpointResponse struct {
	Inventories []SNSSmsReceiverInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加邮箱地址到邮箱接收端
type AddEmailAddressToSNSEmailEndpointRequest struct {
	ReqConfig
	Params     AddSNSSmsReceiverParams `json:"params" bson:"params"`
	SystemTags []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddEmailAddressToSNSEmailEndpointParams struct {
	EmailAddress string   `json:"emailAddress" bson:"emailAddress"`                     //邮箱地址
	EndpointUuid string   `json:"endpointUuid" bson:"endpointUuid"`                     //接收端Uuid
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type AddEmailAddressToSNSEmailEndpointResponse struct {
	Inventory SNSEmailAddressInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新接收端邮箱地址
type UpdateEmailAddressOfSNSEmailEndpointRequest struct {
	ReqConfig
	UpdateEmailAddressOfSNSEmailEndpoint UpdateEmailAddressOfSNSEmailEndpointParams `json:"updateEmailAddressOfSNSEmailEndpoint" bson:"updateEmailAddressOfSNSEmailEndpoint"`
	SystemTags                           []string                                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                             []string                                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateEmailAddressOfSNSEmailEndpointParams struct {
	EmailAddressUuid string `json:"emailAddressUuid" bson:"emailAddressUuid"`             //邮箱地址uuid
	EmailAddress     string `json:"emailAddress,omitempty" bson:"emailAddress,omitempty"` //邮箱地址
	EndpointUuid     string `json:"endpointUuid" bson:"endpointUuid"`                     //接收端Uuid
}

type UpdateEmailAddressOfSNSEmailEndpointResponse struct {
	Inventory SNSEmailAddressInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除邮箱接收端的地址
type DeleteEmailAddressOfSNSEmailEndpointRequest struct {
	ReqConfig
	EmailAddressUuid string   `json:"emailAddressUuid" bson:"emailAddressUuid"`         //邮箱地址uuid
	EndpointUuid     string   `json:"endpointUuid" bson:"endpointUuid"`                 //接收端Uuid
	DeleteMode       string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags       []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteEmailAddressOfSNSEmailEndpointResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询接收端邮箱地址
type QuerySNSEmailAddressRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySNSEmailAddressResponse struct {
	Inventories []SNSEmailAddressInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建微软Teams接收端
type CreateSNSMicrosoftTeamsEndpointRequest struct {
	ReqConfig
	Params     AddSNSSmsReceiverParams `json:"params" bson:"params"`
	SystemTags []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateSNSMicrosoftTeamsEndpointParams struct {
	Url          string   `json:"Url" bson:"Url"`
	Name         string   `json:"name" bson:"name"`
	Description  string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	PlatformUuid string   `json:"platformUuid" bson:"platformUuid"`
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type CreateSNSMicrosoftTeamsEndpointResponse struct {
	Inventory SNSSmsEndpointInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询微软Teams接收端
type QuerySNSMicrosoftTeamsEndpointRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySNSMicrosoftTeamsEndpointResponse struct {
	Inventories []SNSSmsEndpointInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//订阅主题
type SubscribeSNSTopicRequest struct {
	ReqConfig
	TopicUuid    string   `json:"topicUuid" bson:"topicUuid"`                       //应用主题UUID
	EndpointUuid string   `json:"endpointUuid" bson:"endpointUuid"`                 //钉钉终端UUID
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type SubscribeSNSTopicParams struct {
}

type SubscribeSNSTopicResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询订阅主题的终端
type QuerySNSTopicSubscriberRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QuerySNSTopicSubscriberResponse struct {
	Inventories []SNSTopicSubscriberInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//退订主题
type UnsubscribeSNSTopicRequest struct {
	ReqConfig
	TopicUuid    string   `json:"topicUuid" bson:"topicUuid"`                       //主题UUID
	EndpointUuid string   `json:"endpointUuid" bson:"endpointUuid"`                 //终端UUID
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UnsubscribeSNSTopicResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取管理节点目录容量
type GetManagementNodeDirCapacityRequest struct {
	ReqConfig
	ManagementNodeUuids []string `json:"managementNodeUuids,omitempty" bson:"managementNodeUuids,omitempty"` //管理节点的uuid
	SystemTags          []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`                   //云主机系统标签
	UserTags            []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetManagementNodeDirCapacityResponse struct {
	Result  map[string]interface{} `json:"result" bson:"result"`
	Success bool                   `json:"success" bson:"success"`
	Error   ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type SNSTopicInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	State       string `json:"state" bson:"state"`
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type SNSApplicationPlatformInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	State       string `json:"state" bson:"state"`
	Type        string `json:"type" bson:"type"`             //类型
	CreateDate  string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type SNSApplicationEndpointInventory struct {
	UUID         string `json:"uuid" bson:"uuid"`
	Name         string `json:"name" bson:"name"`                                   //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	State        string `json:"state" bson:"state"`
	Type         string `json:"type" bson:"type"` //类型
	PlatformUuid string `json:"platformUuid" bson:"platformUuid"`
	CreateDate   string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type SNSDingTalkEndpointInventory struct {
	Url                  string   `json:"url" bson:"url"`
	AtAll                bool     `json:"atAll,omitempty" bson:"atAll,omitempty"`                               //是否消息@所有人
	AtPersonPhoneNumbers []string `json:"atPersonPhoneNumbers,omitempty" bson:"atPersonPhoneNumbers,omitempty"` //要@用户的电话号码
	Name                 string   `json:"name" bson:"name"`                                                     //资源名称
	UUID                 string   `json:"uuid" bson:"uuid"`
	Description          string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Type                 string   `json:"type" bson:"type"`                                   //类型
	PlatformUuid         string   `json:"platformUuid,omitempty" bson:"platformUuid,omitempty"`
	CreateDate           string   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate           string   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type SNSDingTalkAtPersonInventory struct {
	UUID         string `json:"uuid" bson:"uuid"`
	PhoneNumber  string `json:"phoneNumber" bson:"phoneNumber"`   //用户电话号码（钉钉用户以电话注册）
	EndpointUuid string `json:"endpointUuid" bson:"endpointUuid"` //钉钉终端UUID
}

type SNSSmsEndpointInventory struct {
	Receivers    []string `json:"receivers" bson:"receivers"` //短信接收者
	Url          string   `json:"url" bson:"url"`
	Name         string   `json:"name" bson:"name"` //资源名称
	UUID         string   `json:"uuid" bson:"uuid"`
	Description  string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	Type         string   `json:"type" bson:"type"`                                   //类型
	State        string   `json:"state" bson:"state"`
	PlatformUuid string   `json:"platformUuid,omitempty" bson:"platformUuid,omitempty"`
	CreateDate   string   `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate   string   `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type SNSSmsReceiverInventory struct {
	UUID         string `json:"uuid" bson:"uuid"`
	PhoneNumber  string `json:"phoneNumber" bson:"phoneNumber"`                       //
	EndpointUuid string `json:"endpointUuid,omitempty" bson:"endpointUuid,omitempty"` //短信接收端Uuid
	Type         string `json:"type" bson:"type"`                                     //短信接收端类型:AliyunSms
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	CreateDate   string `json:"createDate" bson:"createDate"`                         //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"`                         //最后一次修改时间
}

type SNSEmailAddressInventory struct {
	UUID         string `json:"uuid" bson:"uuid"`
	EmailAddress string `json:"emailAddress" bson:"emailAddress"`
	EndpointUuid string `json:"endpointUuid" bson:"endpointUuid"` //钉钉终端UUID
}

type SNSTopicSubscriberInventory struct {
	TopicUuid    string `json:"topicUuid" bson:"topicUuid"`       //主题UUID
	EndpointUuid string `json:"endpointUuid" bson:"endpointUuid"` //终端UUID
	CreateDate   string `json:"createDate" bson:"createDate"`     //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"`     //最后一次修改时间
}
