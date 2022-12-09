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

//创建组播路由器
func CreateMulticastRouter(params model.CreateMulticastRouterRequest) models.Result {
	//POST zstack/v1/multicast/virtual-routers
	url := fmt.Sprintf("zstack/v1/multicast/virtual-routers")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateMulticastRouterFailed, err.Error())
	}
	var result model.CreateMulticastRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除组播路由器
func DeleteMulticastRouter(params model.DeleteMulticastRouterRequest) models.Result {
	//DELETE zstack/v1/multicast/virtual-routers/{uuid}
	url := fmt.Sprintf("zstack/v1/multicast/virtual-routers/%s", params.UUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.DeleteMulticastRouterFailed, err.Error())
	}
	var result model.DeleteMulticastRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询组播路由器
func QueryMulticastRouter(params model.QueryMulticastRouterRequest) models.Result {
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
		return models.Error(errcode.QueryMulticastRouterFailed, err.Error())
	}
	var result model.QueryMulticastRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//改变组播路由器状态
func ChangeMulticastRouterState(params model.ChangeMulticastRouterStateRequest) models.Result {
	//PUT zstack/v1/multicast/virtual-routers/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/multicast/virtual-routers/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeMulticastRouterStateFailed, err.Error())
	}
	var result model.ChangeMulticastRouterStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取组播路由表
func GetVpcMulticastRoute(params model.GetVpcMulticastRouteRequest) models.Result {
	//GET zstack/v1/multicast/virtual-routers/{uuid}/routes
	url := fmt.Sprintf("zstack/v1/multicast/virtual-routers/%s/routes", params.UUID)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetVpcMulticastRouteFailed, err.Error())
	}
	var result model.GetVpcMulticastRouteResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加组播路由静态RP地址
func AddRendezvousPointToMulticastRouter(params model.AddRendezvousPointToMulticastRouterRequest) models.Result {
	//POST zstack/v1/multicast/virtual-routers/{uuid}/RendezvousPoint
	url := fmt.Sprintf("zstack/v1/multicast/virtual-routers/%s/RendezvousPoint", params.UUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddRendezvousPointToMulticastRouterFailed, err.Error())
	}
	var result model.AddRendezvousPointToMulticastRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除组播路由器静态RP地址
func RemoveRendezvousPointFromMulticastRouter(params model.RemoveRendezvousPointFromMulticastRouterRequest) models.Result {
	//DELETE zstack/v1/multicast/virtual-routers/{uuid}/RendezvousPoint
	url := fmt.Sprintf("zstack/v1/multicast/virtual-routers/%s/RendezvousPoint", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveRendezvousPointFromMulticastRouterFailed, err.Error())
	}
	var result model.RemoveRendezvousPointFromMulticastRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
