package model

type ReqConfig struct {
	Host            string `json:"host,omitempty" bson:"host,omitempty"`
	AccessKeyId     string `json:"accessKeyId,omitempty" bson:"accessKeyId,omitempty"`
	AccessKeySecret string `json:"accessKeySecret,omitempty" bson:"accessKeySecret,omitempty"`
}
