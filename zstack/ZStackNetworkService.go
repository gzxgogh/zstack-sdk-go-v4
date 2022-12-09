package zstack

import (
	"fmt"
	"github.com/gzxgogh/zstack-sdk-go-v4/errcode"
	"github.com/gzxgogh/zstack-sdk-go-v4/model"
	"github.com/gzxgogh/zstack-sdk-go-v4/request"
	"github.com/maczh/mgin/logs"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
)

//获取网络服务类型
func GetNetworkServiceTypes(params model.GetNetworkServiceTypesRequest) models.Result {
	//GET zstack/v1/network-services/types
	url := fmt.Sprintf("zstack/v1/network-services/types")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetNetworkServiceTypesFailed, err.Error())
	}
	var result model.GetNetworkServiceTypesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询网络服务模块
func QueryNetworkServiceProvider(params model.QueryNetworkServiceProviderRequest) models.Result {
	//GET zstack/v1/network-services/providers
	url := fmt.Sprintf("zstack/v1/network-services/providers")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryNetworkServiceProviderFailed, err.Error())
	}
	var result model.QueryNetworkServiceProviderResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询网络服务与三层网络引用
func QueryNetworkServiceL3NetworkRef(params model.QueryNetworkServiceL3NetworkRefRequest) models.Result {
	//GET zstack/v1/l3-networks/network-services/refs
	url := fmt.Sprintf("zstack/v1/l3-networks/network-services/refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryNetworkServiceL3NetworkRefFailed, err.Error())
	}
	var result model.QueryNetworkServiceL3NetworkRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//挂载网络服务到三层网络
func AttachNetworkServiceToL3Network(params model.AttachNetworkServiceToL3NetworkRequest) models.Result {
	//POST zstack/v1/l3-networks/{l3NetworkUuid}/network-services
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/network-services", params.L3NetworkUUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachNetworkServiceToL3NetworkFailed, err.Error())
	}
	var result model.AttachNetworkServiceToL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从三层网络卸载网络服务
func DetachNetworkServiceFromL3Network(params model.DetachNetworkServiceFromL3NetworkRequest) models.Result {
	//DELETE zstack/v1/l3-networks/{l3NetworkUuid}/network-services?networkServices={networkServices}
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/network-services?networkServices=%s", params.L3NetworkUUID, params.NetworkServices)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachNetworkServiceFromL3NetworkFailed, err.Error())
	}
	var result model.DetachNetworkServiceFromL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
