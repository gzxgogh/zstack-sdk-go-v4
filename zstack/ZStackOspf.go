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

//创建路由区域资源
func CreateVRouterOspfArea(params model.CreateVRouterOspfArea) models.Result {
	//POST zstack/v1/routerArea
	url := fmt.Sprintf("zstack/v1/routerArea")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateVRouterOspfAreaFailed, err.Error())
	}
	var result model.CreateVRouterOspfAreaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除OSPF路由区域
func DeleteVRouterOspfArea(params model.DeleteVRouterOspfAreaRequest) models.Result {
	//DELETE zstack/v1/routerArea/{uuid}
	url := fmt.Sprintf("zstack/v1/routerArea/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteVRouterOspfAreaFailed, err.Error())
	}
	var result model.DeleteVRouterOspfAreaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取OSPF的邻居信息
func GetVRouterOspfNeighbor(params model.GetVRouterOspfNeighborRequest) models.Result {
	//GET zstack/v1/routerArea/{vRouterUuid}/neighbor
	url := fmt.Sprintf("zstack/v1/routerArea/%s/neighbor", params.VRouterUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetVRouterOspfNeighborFailed, err.Error())
	}
	var result model.GetVRouterOspfNeighborResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询OSPF路由区域信息
func QueryVRouterOspfArea(params model.QueryVRouterOspfAreaRequest) models.Result {
	//GET zstack/v1/routerArea
	//GET zstack/v1/routerArea/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/routerArea")
	} else {
		url = fmt.Sprintf("zstack/v1/routerArea/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVRouterOspfAreaFailed, err.Error())
	}
	var result model.QueryVRouterOspfAreaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取路由器的ID
func GetVRouterRouterId(params model.GetVRouterRouterIdRequest) models.Result {
	//GET zstack/v1/routerArea/{vRouterUuid}/routerid
	url := fmt.Sprintf("zstack/v1/routerArea/%s/routerid", params.VRouterUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetVRouterRouterIdFailed, err.Error())
	}
	var result model.GetVRouterRouterIdResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//设置路由器的ID
func SetVRouterRouterId(params model.SetVRouterRouterIdRequest) models.Result {
	//POST zstack/v1/routerArea/{vRouterUuid}/routerid
	url := fmt.Sprintf("zstack/v1/routerArea/%s/routerid", params.VRouterUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.SetVRouterRouterIdFailed, err.Error())
	}
	var result model.SetVRouterRouterIdResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加网络到OSPF的区域
func AddVRouterNetworksToOspfArea(params model.AddVRouterNetworksToOspfAreaRequest) models.Result {
	//POST zstack/v1/routerArea/{routerAreaUuid}/router/{vRouterUuid}/addnetworks
	url := fmt.Sprintf("zstack/v1/routerArea/%s/router/%s/addnetworks", params.VRouterUuid, params.VRouterUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddVRouterNetworksToOspfAreaFailed, err.Error())
	}
	var result model.AddVRouterNetworksToOspfAreaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从路由区域中移除网络
func RemoveVRouterNetworksFromOspfArea(params model.RemoveVRouterNetworksFromOspfAreaRequest) models.Result {
	//DELETE zstack/v1/routerArea/networks
	url := fmt.Sprintf("zstack/v1/routerArea/networks")

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveVRouterNetworksFromOspfAreaFailed, err.Error())
	}
	var result model.RemoveVRouterNetworksFromOspfAreaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更改OSPF的路由区域属性
func UpdateVRouterOspfArea(params model.UpdateVRouterOspfAreaRequest) models.Result {
	//PUT zstack/v1/routerArea/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/routerArea/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateVRouterOspfAreaFailed, err.Error())
	}
	var result model.UpdateVRouterOspfAreaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询路由加入到区域的网络信息
func QueryVRouterOspfNetwork(params model.QueryVRouterOspfNetworkRequest) models.Result {
	//GET zstack/v1/routerArea/network
	//GET zstack/v1/routerArea/networkR/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/routerArea/network")
	} else {
		url = fmt.Sprintf("zstack/v1/routerArea/network/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVRouterOspfNetworkFailed, err.Error())
	}
	var result model.QueryVRouterOspfNetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
