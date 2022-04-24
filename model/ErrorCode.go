package model

type ErrorCode struct {
	Code        string                 `json:"code,omitempty" bson:"code,omitempty"`               //错误码号，错误的全局唯一标识，例如SYS.1000, HOST.1001
	Description string                 `json:"description,omitempty" bson:"description,omitempty"` //错误的概要描述
	Details     string                 `json:"details,omitempty" bson:"details,omitempty"`         //错误的详细信息
	Elaboration string                 `json:"elaboration,omitempty" bson:"elaboration,omitempty"` //保留字段，默认为null
	Opaque      map[string]interface{} `json:"opaque,omitempty" bson:"opaque,omitempty"`           //保留字段
	Cause       *ErrorCode             `json:"cause,omitempty" bson:"cause,omitempty"`             //根错误，引发当前错误的源错误，若无原错误，该字段为null
}
