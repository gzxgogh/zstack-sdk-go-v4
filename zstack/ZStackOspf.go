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

//创建路由区域资源
func CreateVRouterOspfArea(params model.CreateVRouterOspfArea) mgresult.Result {
	//POST zstack/v1/routerArea
	url := fmt.Sprintf("zstack/v1/routerArea")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateVRouterOspfAreaFailed, err)
	}
	var result model.CreateVRouterOspfAreaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除OSPF路由区域
func DeleteVRouterOspfArea(params model.DeleteVRouterOspfAreaRequest) mgresult.Result {
	//DELETE zstack/v1/routerArea/{uuid}
	url := fmt.Sprintf("zstack/v1/routerArea/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVRouterOspfAreaFailed, err)
	}
	var result model.DeleteVRouterOspfAreaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取OSPF的邻居信息
func GetVRouterOspfNeighbor(params model.GetVRouterOspfNeighborRequest) mgresult.Result {
	//GET zstack/v1/routerArea/{vRouterUuid}/neighbor
	url := fmt.Sprintf("zstack/v1/routerArea/%s/neighbor", params.VRouterUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVRouterOspfNeighborFailed, err)
	}
	var result model.GetVRouterOspfNeighborResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询OSPF路由区域信息
func QueryVRouterOspfArea(params model.QueryVRouterOspfAreaRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryVRouterOspfAreaFailed, err)
	}
	var result model.QueryVRouterOspfAreaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取路由器的ID
func GetVRouterRouterId(params model.GetVRouterRouterIdRequest) mgresult.Result {
	//GET zstack/v1/routerArea/{vRouterUuid}/routerid
	url := fmt.Sprintf("zstack/v1/routerArea/%s/routerid", params.VRouterUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVRouterRouterIdFailed, err)
	}
	var result model.GetVRouterRouterIdResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//设置路由器的ID
func SetVRouterRouterId(params model.SetVRouterRouterIdRequest) mgresult.Result {
	//POST zstack/v1/routerArea/{vRouterUuid}/routerid
	url := fmt.Sprintf("zstack/v1/routerArea/%s/routerid", params.VRouterUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVRouterRouterIdFailed, err)
	}
	var result model.SetVRouterRouterIdResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加网络到OSPF的区域
func AddVRouterNetworksToOspfArea(params model.AddVRouterNetworksToOspfAreaRequest) mgresult.Result {
	//POST zstack/v1/routerArea/{routerAreaUuid}/router/{vRouterUuid}/addnetworks
	url := fmt.Sprintf("zstack/v1/routerArea/%s/router/%s/addnetworks", params.VRouterUuid, params.VRouterUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddVRouterNetworksToOspfAreaFailed, err)
	}
	var result model.AddVRouterNetworksToOspfAreaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从路由区域中移除网络
func RemoveVRouterNetworksFromOspfArea(params model.RemoveVRouterNetworksFromOspfAreaRequest) mgresult.Result {
	//DELETE zstack/v1/routerArea/networks
	url := fmt.Sprintf("zstack/v1/routerArea/networks")

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveVRouterNetworksFromOspfAreaFailed, err)
	}
	var result model.RemoveVRouterNetworksFromOspfAreaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更改OSPF的路由区域属性
func UpdateVRouterOspfArea(params model.UpdateVRouterOspfAreaRequest) mgresult.Result {
	//PUT zstack/v1/routerArea/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/routerArea/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateVRouterOspfAreaFailed, err)
	}
	var result model.UpdateVRouterOspfAreaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询路由加入到区域的网络信息
func QueryVRouterOspfNetwork(params model.QueryVRouterOspfNetworkRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryVRouterOspfNetworkFailed, err)
	}
	var result model.QueryVRouterOspfNetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
