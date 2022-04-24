package model

//获取云主机usb重定向开关状态
type GetVmUsbRedirectRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetVmUsbRedirectResponse struct {
	Error  ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Enable bool      `json:"enable" bson:"enable"`
}

//设置云主机usb重定向开关
type SetVmUsbRedirectRequest struct {
	ReqConfig
	UUID             string                 `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	SetVmUsbRedirect SetVmUsbRedirectParams `json:"setVmUsbRedirect" bson:"setVmUsbRedirect"`
	SystemTags       []string               `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags         []string               `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type SetVmUsbRedirectParams struct {
	Enable bool `json:"enable" bson:"enable"`
}

type SetVmUsbRedirectResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//云主机加载物理机USB设备
type AttachUsbDeviceToVmRequest struct {
	ReqConfig
	UsbDeviceUuid string                    `json:"usbDeviceUuid" bson:"usbDeviceUuid"` //资源的UUID，唯一标示该资源
	Params        AttachUsbDeviceToVmParams `json:"params" bson:"params"`
	SystemTags    []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags      []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachUsbDeviceToVmParams struct {
	VmInstanceUuid string `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	AttachType     string `json:"attachType,omitempty" bson:"attachType,omitempty"` //加载类型:PassThrough，Redirect
}

type AttachUsbDeviceToVmResponse struct {
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory UsbDeviceInventory `json:"inventory" bson:"inventory"`
}

//将云主机挂载的USB设备卸载
type DetachUsbDeviceFromVmRequest struct {
	ReqConfig
	UsbDeviceUuid string                      `json:"usbDeviceUuid" bson:"usbDeviceUuid"` //资源的UUID，唯一标示该资源
	Params        DetachUsbDeviceFromVmParams `json:"params" bson:"params"`
	SystemTags    []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags      []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachUsbDeviceFromVmParams struct {
}

type DetachUsbDeviceFromVmResponse struct {
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory UsbDeviceInventory `json:"inventory" bson:"inventory"`
}

//获取USB透传候选列表
type GetUsbDeviceCandidatesForAttachingVmRequest struct {
	ReqConfig
	VmInstanceUuid string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"`
	AttachType     string   `json:"attachType,omitempty" bson:"attachType,omitempty"`
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetUsbDeviceCandidatesForAttachingVmResponse struct {
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []UsbDeviceInventory `json:"inventories" bson:"inventories"`
}

//查询USB设备
type QueryUsbDeviceRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"` //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryUsbDeviceResponse struct {
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventories []UsbDeviceInventory `json:"inventories" bson:"inventories"`
}

//更新USB设备
type UpdateUsbDeviceRequest struct {
	ReqConfig
	UUID            string                `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateUsbDevice UpdateUsbDeviceParams `json:"updateUsbDevice" bson:"updateUsbDevice"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateUsbDeviceParams struct {
	Name        string `json:"name" bson:"name"`               //资源名称
	Description string `json:"description" bson:"description"` //资源的详细描述
	State       string `json:"state" bson:"state"`             //USB设备状态
}

type UpdateUsbDeviceResponse struct {
	Error     ErrorCode          `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory UsbDeviceInventory `json:"inventory" bson:"inventory"`
}

type UsbDeviceInventory struct {
	UUID           string `json:"uuid" bson:"uuid"`                     //资源的UUID，唯一标示该资源
	Name           string `json:"name" bson:"name"`                     //资源名称
	Description    string `json:"description" bson:"description"`       //资源的详细描述
	HostUuid       string `json:"hostUuid" bson:"hostUuid"`             //物理机UUID
	VmInstanceUuid string `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //云主机uuid
	BusNum         string `json:"busNum" bson:"busNum"`                 //总线号
	DevNum         string `json:"devNum" bson:"devNum"`                 //设备号
	IdVendor       string `json:"idVendor" bson:"idVendor"`             //VendorID
	IdProduct      string `json:"idProduct" bson:"idProduct"`           //ProductID
	IManufacturer  string `json:"iManufacturer" bson:"iManufacturer"`   //生产商
	IProduct       string `json:"iProduct" bson:"iProduct"`             //设备类型
	ISerial        string `json:"iSerial" bson:"iSerial"`               //序列号
	UsbVersion     string `json:"usbVersion" bson:"usbVersion"`         //USB版本
	AttachType     string `json:"attachType" bson:"attachType"`         //加载方式
	CreateDate     string `json:"createDate" bson:"createDate"`         //创建时间
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"`         //最后一次修改时间
	State          string `json:"state" bson:"state"`
}
