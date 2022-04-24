package model

//创建IPsec连接
type CreateIPsecConnectionRequest struct {
	ReqConfig
	Params     CreateIPsecConnectionParams `json:"params" bson:"params"`
	SystemTags []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateIPsecConnectionParams struct {
	Name                      string   `json:"name" bson:"name"`                                       //资源名称
	Description               string   `json:"description,omitempty" bson:"description,omitempty"`     //详细描述
	L3NetworkUuid             string   `json:"l3NetworkUuid,omitempty" bson:"l3NetworkUuid,omitempty"` //三层网络UUID
	PeerAddress               string   `json:"peerAddress" bson:"peerAddress"`                         //对端地址
	AuthMode                  string   `json:"authMode,omitempty" bson:"authMode,omitempty"`           //认证模式:psk,certs
	AuthKey                   string   `json:"authKey" bson:"authKey"`                                 //认证密钥
	VipUuid                   string   `json:"vipUuid" bson:"vipUuid"`
	PeerCidrs                 []string `json:"peerCidrs,omitempty" bson:"peerCidrs,omitempty"`                                 //对端CIDR
	IkeAuthAlgorithm          string   `json:"ikeAuthAlgorithm,omitempty" bson:"ikeAuthAlgorithm,omitempty"`                   //IKE验证算法:md5,sha1,sha256,sha384,sha512
	IkeEncryptionAlgorithm    string   `json:"ikeEncryptionAlgorithm,omitempty" bson:"ikeEncryptionAlgorithm,omitempty"`       //IKE加密算法:3des,aes-128,aes-192,aes-256
	IkeDhGroup                string   `json:"ikeDhGroup,omitempty" bson:"ikeDhGroup,omitempty"`                               //IKE完整前向保密
	PolicyAuthAlgorithm       string   `json:"policyAuthAlgorithm,omitempty" bson:"policyAuthAlgorithm,omitempty"`             //ESP认证算法:md5,sha1,sha256,sha384,sha512
	PolicyEncryptionAlgorithm string   `json:"policyEncryptionAlgorithm,omitempty" bson:"policyEncryptionAlgorithm,omitempty"` //ESP加密算法:3des,aes-128,aes-192,aes-256
	Pfs                       string   `json:"pfs,omitempty" bson:"pfs,omitempty"`                                             //完全正向保密:dh-group2,dh-group5,dh-group14,dh-group15,dh-group16,dh-group17,dh-group18,dh-group19,dh-group20,dh-group21,dh-group22,dh-group23,dh-group24,dh-group25,dh-group26
	PolicyMode                string   `json:"policyMode,omitempty" bson:"policyMode,omitempty"`                               //工作模式:tunnel,transport
	TransformProtocol         string   `json:"transformProtocol,omitempty" bson:"transformProtocol,omitempty"`                 //传输安全协议:esp,ah,ah-esp
	ResourceUuid              string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`                           //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateIPsecConnectionResponse struct {
	Inventory IPsecConnectionInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除IPsec连接
type DeleteIPsecConnectionRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteIPsecConnectionResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新IPsec连接
type UpdateIPsecConnectionRequest struct {
	ReqConfig
	UUID                  string                      `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateIPsecConnection UpdateIPsecConnectionParams `json:"updateIPsecConnection" bson:"updateIPsecConnection"`
	SystemTags            []string                    `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags              []string                    `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateIPsecConnectionParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateIPsecConnectionResponse struct {
	Inventory IPsecConnectionInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询IPsec连接
type QueryIPSecConnectionRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryIPSecConnectionResponse struct {
	Inventories []IPsecConnectionInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更改IPsec连接状态
type ChangeIPSecConnectionStateRequest struct {
	ReqConfig
	UUID                       string                           `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	ChangeIPSecConnectionState ChangeIPSecConnectionStateParams `json:"changeIPSecConnectionState" bson:"changeIPSecConnectionState"`
	SystemTags                 []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeIPSecConnectionStateParams struct {
	StateEvent string `json:"stateEvent" bson:"stateEvent"`
}

type ChangeIPSecConnectionStateResponse struct {
	Inventory IPsecConnectionInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加三层网络到IPsec连接
type AttachL3NetworksToIPsecConnectionRequest struct {
	ReqConfig
	UUID       string                                  `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Params     AttachL3NetworksToIPsecConnectionParams `json:"params" bson:"params"`
	SystemTags []string                                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AttachL3NetworksToIPsecConnectionParams struct {
	L3NetworkUuids []string `json:"l3NetworkUuids" bson:"l3NetworkUuids"`
	ResourceUuid   string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type AttachL3NetworksToIPsecConnectionResponse struct {
	Inventory IPsecConnectionInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从IPsec连接删除三层网络
type DetachL3NetworksFromIPsecConnectionRequest struct {
	ReqConfig
	UUID           string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	L3NetworkUuids []string `json:"l3NetworkUuids" bson:"l3NetworkUuids"`
	SystemTags     []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags       []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DetachL3NetworksFromIPsecConnectionResponse struct {
	Inventory IPsecConnectionInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加远端CIDR到IPsec连接
type AddRemoteCidrsToIPsecConnectionRequest struct {
	ReqConfig
	UUID       string                                `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Params     AddRemoteCidrsToIPsecConnectionParams `json:"params" bson:"params"`
	SystemTags []string                              `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddRemoteCidrsToIPsecConnectionParams struct {
	PeerCidrs    []string `json:"peerCidrs" bson:"peerCidrs"`
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type AddRemoteCidrsToIPsecConnectionResponse struct {
	Inventory IPsecConnectionInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除远端CIDR
type RemoveRemoteCidrsFromIPsecConnectionRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	PeerCidrs  []string `json:"peerCidrs" bson:"peerCidrs"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveRemoteCidrsFromIPsecConnectionResponse struct {
	Inventory IPsecConnectionInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type IPsecConnectionInventory struct {
	UUID                      string          `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name                      string          `json:"name" bson:"name"`                                   //资源名称
	Description               string          `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	PeerAddress               string          `json:"peerAddress" bson:"peerAddress"`                     //对端地址
	AuthMode                  string          `json:"authMode,omitempty" bson:"authMode,omitempty"`       //认证模式:psk,certs
	AuthKey                   string          `json:"authKey" bson:"authKey"`                             //认证密钥
	VipUuid                   string          `json:"vipUuid" bson:"vipUuid"`
	IkeAuthAlgorithm          string          `json:"ikeAuthAlgorithm,omitempty" bson:"ikeAuthAlgorithm,omitempty"`                   //IKE验证算法:md5,sha1,sha256,sha384,sha512
	IkeEncryptionAlgorithm    string          `json:"ikeEncryptionAlgorithm,omitempty" bson:"ikeEncryptionAlgorithm,omitempty"`       //IKE加密算法:3des,aes-128,aes-192,aes-256
	IkeDhGroup                string          `json:"ikeDhGroup,omitempty" bson:"ikeDhGroup,omitempty"`                               //IKE完整前向保密
	PolicyAuthAlgorithm       string          `json:"policyAuthAlgorithm,omitempty" bson:"policyAuthAlgorithm,omitempty"`             //ESP认证算法:md5,sha1,sha256,sha384,sha512
	PolicyEncryptionAlgorithm string          `json:"policyEncryptionAlgorithm,omitempty" bson:"policyEncryptionAlgorithm,omitempty"` //ESP加密算法:3des,aes-128,aes-192,aes-256
	Pfs                       string          `json:"pfs,omitempty" bson:"pfs,omitempty"`                                             //完全正向保密:dh-group2,dh-group5,dh-group14,dh-group15,dh-group16,dh-group17,dh-group18,dh-group19,dh-group20,dh-group21,dh-group22,dh-group23,dh-group24,dh-group25,dh-group26
	PolicyMode                string          `json:"policyMode,omitempty" bson:"policyMode,omitempty"`                               //工作模式:tunnel,transport
	TransformProtocol         string          `json:"transformProtocol,omitempty" bson:"transformProtocol,omitempty"`                 //传输安全协议:esp,ah,ah-esp
	State                     string          `json:"state" bson:"state"`
	Status                    string          `json:"status" bson:"status"`
	CreateDate                string          `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate                string          `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	PeerCidrs                 []PeerCidrs     `json:"peerCidrs" bson:"peerCidrs"`
	L3NetworkRefs             []L3NetworkRefs `json:"l3NetworkUuid,omitempty" bson:"l3NetworkUuid,omitempty"` //三层网络UUID
}

type PeerCidrs struct {
	UUID           string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Cidr           string `json:"cidr" bson:"cidr"`
	ConnectionUuid string `json:"connectionUuid" bson:"connectionUuid"`
	CreateDate     string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type L3NetworkRefs struct {
	UUID           string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	L3NetworkUuid  string `json:"l3NetworkUuid" bson:"l3NetworkUuid"`
	ConnectionUuid string `json:"connectionUuid" bson:"connectionUuid"`
	CreateDate     string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate     string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}
