package model

//创建资源价格
type CreateResourcePriceRequest struct {
	ReqConfig
	Params     CreateResourcePriceParams `json:"params" bson:"params"`
	SystemTags []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateResourcePriceParams struct {
	ResourceName string  `json:"resourceName" bson:"resourceName"`                     //cpu,memory,rootvolume,datavolume,snapshot,gpu,pubIpVmNicBandwidthOut,pubIpVmNicBandwidthIn,pubIpVipBandwidthOut,pubIpVipBandwidthIn
	ResourceUnit string  `json:"resourceUnit,omitempty" bson:"resourceUnit,omitempty"` //可选值根据resourceName而定
	TimeUnit     string  `json:"timeUnit,omitempty" bson:"timeUnit,omitempty"`         //计费时间单元
	Price        float64 `json:"price,omitempty" bson:"price,omitempty"`               //单位价格
	AccountUuid  string  `json:"accountUuid,omitempty" bson:"accountUuid,omitempty"`   //账户UUID
	DateInLong   int64   `json:"dateInLong,omitempty" bson:"dateInLong,omitempty"`     //长整型时刻
	TableUuid    string  `json:"tableUuid,omitempty" bson:"tableUuid,omitempty"`       //价目表UUID
	ResourceUuid string  `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateResourcePriceResponse struct {
	Inventory PriceInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除资源价格
type DeleteResourcePriceRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteResourcePriceResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改资源价格
type UpdateResourcePriceRequest struct {
	ReqConfig
	UUID                string                    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateResourcePrice UpdateResourcePriceParams `json:"updateResourcePrice" bson:"updateResourcePrice"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateResourcePriceParams struct {
	EndDateInLong                     int64 `json:"endDateInLong ,omitempty" bson:"endDateInLong ,omitempty"`                                         //资源价格生效截止时间
	SetEndDateInLongBaseOnCurrentTime bool  `json:"setEndDateInLongBaseOnCurrentTime ,omitempty" bson:"setEndDateInLongBaseOnCurrentTime ,omitempty"` //设置资源价格生效截止时间为当前时间
}

type UpdateResourcePriceResponse struct {
	Inventory PriceInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询资源价格
type QueryResourcePriceRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryResourcePriceResponse struct {
	Inventories []PriceInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode        `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//计算账户花费
type CalculateAccountSpendingRequest struct {
	ReqConfig
	AccountUuid              string                         `json:"accountUuid" bson:"accountUuid"`
	CalculateAccountSpending CalculateAccountSpendingParams `json:"calculateAccountSpending" bson:"calculateAccountSpending"`
	SystemTags               []string                       `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                 []string                       `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CalculateAccountSpendingParams struct {
	DateStart int64 `json:"dateStart " bson:"dateStart "`
	DateEnd   int64 `json:"dateEnd " bson:"dateEnd "`
}

type CalculateAccountSpendingResponse struct {
	Total    float64   `json:"total" bson:"total"` //总额
	Success  bool      `json:"success" bson:"success"`
	Spending Spending  `json:"spending" bson:"spending"`
	Error    ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询账户账单
type QueryAccountBillingRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAccountBillingResponse struct {
	Inventories []AccountBillingInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                 `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建计费价目
type CreatePriceTableRequest struct {
	ReqConfig
	Params     CreatePriceTableParams `json:"params" bson:"params"`
	SystemTags []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreatePriceTableParams struct {
	Name         string           `json:"name,omitempty" bson:"name,omitempty"` //资源名称
	Description  string           `json:"description" bson:"description"`       //详细描述
	Prices       []PriceInventory `json:"prices" bson:"prices"`
	ResourceUuid string           `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string         `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreatePriceTableResponse struct {
	Inventory PriceTableInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除计费价目
type DeletePriceTableRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeletePriceTableResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改计费价目
type UpdatePriceTableRequest struct {
	ReqConfig
	UUID             string                 `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdatePriceTable UpdatePriceTableParams `json:"updatePriceTable" bson:"updatePriceTable"`
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdatePriceTableParams struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"` //资源名称
	Description string `json:"description" bson:"description"`       //详细描述
}

type UpdatePriceTableResponse struct {
	Inventory PriceTableInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询计费价目
type QueryPriceTableRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryPriceTableResponse struct {
	Inventories []PriceTableInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//为账户/项目指定计费价目
type AttachPriceTableToAccountRequest struct {
	ReqConfig
	AccountUuid string   `json:"accountUuid" bson:"accountUuid"`                   //账户UUID
	TableUuid   string   `json:"tableUuid" bson:"tableUuid"`                       //价目表UUID
	TagUuids    []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`     //标签UUID列表
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachPriceTableToAccountResponse struct {
	Inventory PriceTableInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//取消账户/项目关联的计费价目
type DetachPriceTableFromAccountRequest struct {
	ReqConfig
	AccountUuid string   `json:"accountUuid" bson:"accountUuid"`                   //账户UUID
	TableUuid   string   `json:"tableUuid" bson:"tableUuid"`                       //价目表UUID
	TagUuids    []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`     //标签UUID列表
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachPriceTableFromAccountResponse struct {
	Inventory []PriceTableInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改账户/项目计费价目
type ChangeAccountPriceTableBindingRequest struct {
	ReqConfig
	AccountUuid                    string                               `json:"accountUuid" bson:"accountUuid"` //账户UUID
	TableUuid                      string                               `json:"tableUuid" bson:"tableUuid"`     //价目表UUID
	ChangeAccountPriceTableBinding ChangeAccountPriceTableBindingParams `json:"changeAccountPriceTableBinding" bson:"changeAccountPriceTableBinding"`
	SystemTags                     []string                             `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                       []string                             `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeAccountPriceTableBindingParams struct {
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`         //标签UUID列表
}

type ChangeAccountPriceTableBindingResponse struct {
	Inventory PriceTableInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode           `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获得项目账户使用的计费价目
type GetAccountPriceTableRefRequest struct {
	ReqConfig
	AccountUuid string   `json:"accountUuid,omitempty" bson:"accountUuid,omitempty"` //账户UUID
	TableUuid   string   `json:"tableUuid,omitempty" bson:"tableUuid,omitempty"`     //价目表UUID
	SystemTags  []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`   //云主机系统标签
	UserTags    []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetAccountPriceTableRefResponse struct {
	AccountUuids []string  `json:"accountUuids" bson:"accountUuids"`
	TableUuid    string    `json:"tableUuid" bson:"tableUuid"` //价目表UUID
	Success      bool      `json:"success" bson:"success"`
	Error        ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询账户项目与计费价目的关联关系
type QueryAccountPriceTableRefRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAccountPriceTableRefResponse struct {
	Inventories []PriceTableInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//生成账单
type GenerateAccountBillingRequest struct {
	ReqConfig
	AccountUuid            string                       `json:"accountUuid,omitempty" bson:"accountUuid,omitempty"` //账户UUID
	GenerateAccountBilling GenerateAccountBillingParams `json:"generateAccountBilling" bson:"generateAccountBilling"`
	SystemTags             []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags               []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GenerateAccountBillingParams struct {
}

type GenerateAccountBillingResponse struct {
	Inventory AccountBillingInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type PriceInventory struct {
	UUID               string             `json:"uuid" bson:"uuid"`                 //资源的UUID，唯一标示该资源
	ResourceName       string             `json:"resourceName" bson:"resourceName"` //cpu,memory,rootvolume,datavolume,snapshot,gpu,pubIpVmNicBandwidthOut,pubIpVmNicBandwidthIn,pubIpVipBandwidthOut,pubIpVipBandwidthIn
	ResourceUnit       string             `json:"resourceUnit" bson:"resourceUnit"` //可选值根据resourceName而定
	TimeUnit           string             `json:"timeUnit" bson:"timeUnit"`         //计费时间单元
	Price              float64            `json:"price" bson:"price"`               //单位价格
	DateInLong         int64              `json:"dateInLong" bson:"dateInLong"`     //长整型时刻
	CreateDate         string             `json:"createDate" bson:"createDate"`     //创建时间
	LastOpDate         string             `json:"lastOpDate" bson:"lastOpDate"`     //最后一次修改时间
	TableUuid          string             `json:"tableUuid" bson:"tableUuid"`       //价目表UUID
	PciDeviceOfferings PciDeviceOfferings `json:"pciDeviceOfferings" bson:"pciDeviceOfferings"`
}

type PciDeviceOfferings struct {
	PriceUuid             string `json:"priceUuid" bson:"priceUuid"`
	PciDeviceOfferingUuid string `json:"pciDeviceOfferingUuid" bson:"pciDeviceOfferingUuid"`
	CreateDate            string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate            string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type Spending struct {
	SpendingType string  `json:"spendingType" bson:"spendingType"` //花费类型
	Spending     float64 `json:"spending" bson:"spending"`         //花费总额
	DateStart    int64   `json:"dateStart " bson:"dateStart "`     //计费开始日期
	DateEnd      int64   `json:"dateEnd " bson:"dateEnd "`         //计费结束日期
	Details      Details `json:"details" bson:"details"`
}

type Details struct {
	ResourceUuid string  `json:"resourceUuid" bson:"resourceUuid"` //资源UUID
	ResourceName string  `json:"resourceName" bson:"resourceName"` //资源名称
	Spending     float64 `json:"spending" bson:"spending"`         //花费
	Type         string  `json:"type" bson:"type"`                 //类型
}

type AccountBillingInventory struct {
	Id             string  `json:"id" bson:"id"`
	BillingType    string  `json:"billingType" bson:"billingType"`       ///账单类型
	AccountUuid    string  `json:"accountUuid" bson:"accountUuid"`       //账户UUID
	ResourceUuid   string  `json:"resourceUuid" bson:"resourceUuid"`     //资源UUID
	ResourceName   string  `json:"resourceName" bson:"resourceName"`     //资源名称
	Spending       float64 `json:"spending" bson:"spending"`             //花费
	StartTime      int64   `json:"startTime" bson:"startTime"`           //资源计费开始时间
	EndTime        int64   `json:"endTime" bson:"endTime"`               //资源计费结束时间
	HypervisorType string  `json:"hypervisorType" bson:"hypervisorType"` //虚拟机管理程序类型,KVM Simulator
	CreateDate     string  `json:"createDate" bson:"createDate"`         //创建时间
	LastOpDate     string  `json:"lastOpDate" bson:"lastOpDate"`         //最后一次修改时间
}

type PriceTableInventory struct {
	UUID        string `json:"uuid" bson:"uuid"`                     //资源的UUID，唯一标示该资源
	Name        string `json:"name,omitempty" bson:"name,omitempty"` //资源名称
	Description string `json:"description" bson:"description"`       //详细描述
	CreateDate  string `json:"createDate" bson:"createDate"`         //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"`         //最后一次修改时间
}
