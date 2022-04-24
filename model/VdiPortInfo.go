package model

type VdiPortInfo struct {
	VncPort      int `json:"vncPort" bson:"vncPort"`           //	vnc端口号
	SpicePort    int `json:"spicePort" bson:"spicePort"`       //spice端口号
	SpiceTlsPort int `json:"spiceTlsPort" bson:"spiceTlsPort"` //spice开启Tls加密，会存在spiceTlsPort和spicePort两个端口号，通过spice客户端连接云主机需要使用spiceTlsPort端口号
}
