package model

type Formats struct {
	Format                    string   `json:"format" bson:"format"`                                       //云盘格式
	MasterHypervisorType      string   `json:"masterHypervisorType" bson:"masterHypervisorType"`           //默认的Hypervisor类型
	SupportingHypervisorTypes []string `json:"supportingHypervisorTypes" bson:"supportingHypervisorTypes"` //支持的Hypervisor类型列表
}
