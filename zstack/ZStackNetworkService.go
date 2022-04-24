package zstack

import (
	"fmt"
	"github.com/gzxgogh/zstack-sdk-go-v4/errcode"
	"github.com/gzxgogh/zstack-sdk-go-v4/model"
	"github.com/gzxgogh/zstack-sdk-go-v4/request"
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/logs"
	"github.com/maczh/mgerr"
	"github.com/maczh/utils"
)

//获取网络服务类型
func GetNetworkServiceTypes(params model.GetNetworkServiceTypesRequest) mgresult.Result {
	//GET zstack/v1/network-services/types
	url := fmt.Sprintf("zstack/v1/network-services/types")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetNetworkServiceTypesFailed, err)
	}
	var result model.GetNetworkServiceTypesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询网络服务模块
func QueryNetworkServiceProvider(params model.QueryNetworkServiceProviderRequest) mgresult.Result {
	//GET zstack/v1/network-services/providers
	url := fmt.Sprintf("zstack/v1/network-services/providers")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryNetworkServiceProviderFailed, err)
	}
	var result model.QueryNetworkServiceProviderResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询网络服务与三层网络引用
func QueryNetworkServiceL3NetworkRef(params model.QueryNetworkServiceL3NetworkRefRequest) mgresult.Result {
	//GET zstack/v1/l3-networks/network-services/refs
	url := fmt.Sprintf("zstack/v1/l3-networks/network-services/refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryNetworkServiceL3NetworkRefFailed, err)
	}
	var result model.QueryNetworkServiceL3NetworkRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//挂载网络服务到三层网络
func AttachNetworkServiceToL3Network(params model.AttachNetworkServiceToL3NetworkRequest) mgresult.Result {
	//POST zstack/v1/l3-networks/{l3NetworkUuid}/network-services
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/network-services", params.L3NetworkUUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachNetworkServiceToL3NetworkFailed, err)
	}
	var result model.AttachNetworkServiceToL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从三层网络卸载网络服务
func DetachNetworkServiceFromL3Network(params model.DetachNetworkServiceFromL3NetworkRequest) mgresult.Result {
	//DELETE zstack/v1/l3-networks/{l3NetworkUuid}/network-services?networkServices={networkServices}
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/network-services?networkServices=%s", params.L3NetworkUUID, params.NetworkServices)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachNetworkServiceFromL3NetworkFailed, err)
	}
	var result model.DetachNetworkServiceFromL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
