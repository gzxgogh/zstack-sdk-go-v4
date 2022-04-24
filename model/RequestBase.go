package model

type RequestBase struct {
	SessionId       string `json:"sessionId" bson:"sessionId"`
	AccessKeyId     string `json:"accessKeyId" bson:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" bson:"accessKeySecret"`
	RequestIp       string `json:"requestIp" bson:"requestIp"`
}
