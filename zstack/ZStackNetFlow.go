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

//创建流量监控搜集器
func CreateFlowCollector(params model.CreateFlowCollectorRequest) mgresult.Result {
	//POST zstack/v1/flowmeters/{flowMeterUuid}/collectors
	url := fmt.Sprintf("zstack/v1/flowmeters/%s/collectors", params.FlowMeterUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateFlowCollectorFailed, err)
	}
	var result model.CreateFlowCollectorResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询流量监控的搜集器
func QueryFlowCollector(params model.QueryFlowCollectorRequest) mgresult.Result {
	//GET zstack/v1/flowmeters/collectors
	//GET zstack/v1/flowmeters/collectors/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/flowmeters/collectors")
	} else {
		url = fmt.Sprintf("zstack/v1/flowmeters/collectors/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryFlowCollectorFailed, err)
	}
	var result model.QueryFlowCollectorResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新搜集器信息
func UpdateFlowCollector(params model.UpdateFlowCollectorRequest) mgresult.Result {
	//PUT zstack/v1/flowmeters/collectors/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/flowmeters/collectors/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateFlowCollectorFailed, err)
	}
	var result model.UpdateFlowCollectorResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除流量监控搜集器
func DeleteFlowCollector(params model.DeleteFlowCollectorRequest) mgresult.Result {
	//DELETE zstack/v1/flowmeters/collectors/{uuid}
	url := fmt.Sprintf("zstack/v1/flowmeters/collectors/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteFlowCollectorDeleteFlowCollectorFailed, err)
	}
	var result model.DeleteFlowCollectorResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建流量监控资源
func CreateFlowMeter(params model.CreateFlowMeterRequest) mgresult.Result {
	//POST zstack/v1/flowmeters
	url := fmt.Sprintf("zstack/v1/flowmeters")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateFlowMeterFailed, err)
	}
	var result model.CreateFlowMeterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询流量监控资源
func QueryFlowMeter(params model.QueryFlowMeterRequest) mgresult.Result {
	//GET zstack/v1/flowmeters
	//GET zstack/v1/flowmeters/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/flowmeters")
	} else {
		url = fmt.Sprintf("zstack/v1/flowmeters/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryFlowMeterFailed, err)
	}
	var result model.QueryFlowMeterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新流量监控的信息
func UpdateFlowMeter(params model.UpdateFlowMeterRequest) mgresult.Result {
	//PUT zstack/v1/flowmeters/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/flowmeters/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateFlowMeterFailed, err)
	}
	var result model.UpdateFlowMeterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除流量监控资源
func DeleteFlowMeter(params model.DeleteFlowMeterRequest) mgresult.Result {
	//DELETE zstack/v1/flowmeters/{uuid}
	url := fmt.Sprintf("zstack/v1/flowmeters/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteFlowMeterFailed, err)
	}
	var result model.DeleteFlowMeterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加VPC路由器网络到FlowMeter中
func AddVRouterNetworksToFlowMeter(params model.AddVRouterNetworksToFlowMeterRequest) mgresult.Result {
	//POST zstack/v1/flowmeters/{flowMeterUuid}/router/{vRouterUuid}/addnetworks
	url := fmt.Sprintf("zstack/v1/flowmeters/%s/router/%s/addnetworks", params.FlowMeterUuid, params.VRouterUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddVRouterNetworksToFlowMeterFailed, err)
	}
	var result model.AddVRouterNetworksToFlowMeterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询开启流量监控的VPC网络
func QueryVRouterFlowMeterNetwork(params model.QueryVRouterFlowMeterNetworkRequest) mgresult.Result {
	//GET zstack/v1/flowmeters/networks
	//GET zstack/v1/flowmeters/networks/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/flowmeters/networks")
	} else {
		url = fmt.Sprintf("zstack/v1/flowmeters/networks/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVRouterFlowMeterNetworkFailed, err)
	}
	var result model.QueryVRouterFlowMeterNetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//流量监控的统计信息
func GetVRouterFlowCounter(params model.GetVRouterFlowCounterRequest) mgresult.Result {
	//GET zstack/v1/flowmeters/{vRouterUuid}/counter
	url := fmt.Sprintf("zstack/v1/flowmeters/%s/counter", params.VRouterUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVRouterFlowCounterFailed, err)
	}
	var result model.GetVRouterFlowCounterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//清除VPC路由器网络的流量监控
func RemoveVRouterNetworksFromFlowMeter(params model.RemoveVRouterNetworksFromFlowMeterRequest) mgresult.Result {
	//DELETE zstack/v1/flowmeters/networks
	url := fmt.Sprintf("zstack/v1/flowmeters/networks")

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveVRouterNetworksFromFlowMeterFailed, err)
	}
	var result model.RemoveVRouterNetworksFromFlowMeterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取流量监控协议中使用的系统标识
func GetFlowMeterRouterId(params model.GetFlowMeterRouterIdRequest) mgresult.Result {
	//GET zstack/v1/flowmeters/{vRouterUuid}/routerid
	url := fmt.Sprintf("zstack/v1/flowmeters/%s/routerid", params.VRouterUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetFlowMeterRouterIdFailed, err)
	}
	var result model.GetFlowMeterRouterIdResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//设置流量监控协议使用的VPC标识
func SetFlowMeterRouterId(params model.SetFlowMeterRouterIdRequest) mgresult.Result {
	//POST zstack/v1/flowmeters/{vRouterUuid}/routerid
	url := fmt.Sprintf("zstack/v1/flowmeters/%s/routerid", params.VRouterUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetFlowMeterRouterIdFailed, err)
	}
	var result model.SetFlowMeterRouterIdResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
