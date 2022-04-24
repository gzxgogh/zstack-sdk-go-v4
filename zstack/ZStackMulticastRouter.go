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

//创建组播路由器
func CreateMulticastRouter(params model.CreateMulticastRouterRequest) mgresult.Result {
	//POST zstack/v1/multicast/virtual-routers
	url := fmt.Sprintf("zstack/v1/multicast/virtual-routers")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateMulticastRouterFailed, err)
	}
	var result model.CreateMulticastRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除组播路由器
func DeleteMulticastRouter(params model.DeleteMulticastRouterRequest) mgresult.Result {
	//DELETE zstack/v1/multicast/virtual-routers/{uuid}
	url := fmt.Sprintf("zstack/v1/multicast/virtual-routers/%s", params.UUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteMulticastRouterFailed, err)
	}
	var result model.DeleteMulticastRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询组播路由器
func QueryMulticastRouter(params model.QueryMulticastRouterRequest) mgresult.Result {
	//GET zstack/v1/multicast/virtual-routers
	//GET zstack/v1/multicast/virtual-routers/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/multicast/virtual-routers")
	} else {
		url = fmt.Sprintf("zstack/v1/multicast/virtual-routers/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryMulticastRouterFailed, err)
	}
	var result model.QueryMulticastRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//改变组播路由器状态
func ChangeMulticastRouterState(params model.ChangeMulticastRouterStateRequest) mgresult.Result {
	//PUT zstack/v1/multicast/virtual-routers/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/multicast/virtual-routers/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeMulticastRouterStateFailed, err)
	}
	var result model.ChangeMulticastRouterStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取组播路由表
func GetVpcMulticastRoute(params model.GetVpcMulticastRouteRequest) mgresult.Result {
	//GET zstack/v1/multicast/virtual-routers/{uuid}/routes
	url := fmt.Sprintf("zstack/v1/multicast/virtual-routers/%s/routes", params.UUID)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVpcMulticastRouteFailed, err)
	}
	var result model.GetVpcMulticastRouteResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加组播路由静态RP地址
func AddRendezvousPointToMulticastRouter(params model.AddRendezvousPointToMulticastRouterRequest) mgresult.Result {
	//POST zstack/v1/multicast/virtual-routers/{uuid}/RendezvousPoint
	url := fmt.Sprintf("zstack/v1/multicast/virtual-routers/%s/RendezvousPoint", params.UUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddRendezvousPointToMulticastRouterFailed, err)
	}
	var result model.AddRendezvousPointToMulticastRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除组播路由器静态RP地址
func RemoveRendezvousPointFromMulticastRouter(params model.RemoveRendezvousPointFromMulticastRouterRequest) mgresult.Result {
	//DELETE zstack/v1/multicast/virtual-routers/{uuid}/RendezvousPoint
	url := fmt.Sprintf("zstack/v1/multicast/virtual-routers/%s/RendezvousPoint", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveRendezvousPointFromMulticastRouterFailed, err)
	}
	var result model.RemoveRendezvousPointFromMulticastRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
