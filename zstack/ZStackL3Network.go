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

//创建三层网络
func CreateL3Network(params model.CreateL3NetworkRequest) models.Result {
	//POST zstack/v1/l3-networks
	url := fmt.Sprintf("zstack/v1/l3-networks")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateL3NetworkFailed, err.Error())
	}
	var result model.CreateL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除三层网络
func DeleteL3Network(params model.DeleteL3NetworkRequest) models.Result {
	//DELETE zstack/v1/l3-networks/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/l3-networks/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteL3NetworkFailed, err.Error())
	}
	var result model.DeleteL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询三层网络
func QueryL3Network(params model.QueryL3NetworkRequest) models.Result {
	//GET zstack/v1/l3-networks
	//GET zstack/v1/l3-networks/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/l3-networks")
	} else {
		url = fmt.Sprintf("zstack/v1/l3-networks/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryL3NetworkFailed, err.Error())
	}
	var result model.QueryL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新三层网络
func UpdateL3Network(params model.UpdateL3NetworkRequest) models.Result {
	//PUT zstack/v1/l3-networks/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateL3NetworkFailed, err.Error())
	}
	var result model.UpdateL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取三层网络类型
func GetL3NetworkTypes(params model.GetL3NetworkTypesRequest) models.Result {
	//GET zstack/v1/l3-networks/types
	url := fmt.Sprintf("zstack/v1/l3-networks/types")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetL3NetworkTypesFailed, err.Error())
	}
	var result model.GetL3NetworkTypesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//改变三层网络状态
func ChangeL3NetworkState(params model.ChangeL3NetworkStateRequest) models.Result {
	//PUT zstack/v1/l3-networks/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeL3NetworkStateFailed, err.Error())
	}
	var result model.ChangeL3NetworkStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取网络DHCP服务所用地址
func GetL3NetworkDhcpIpAddress(params model.GetL3NetworkDhcpIpAddressRequest) models.Result {
	//GET zstack/v1/l3-networks/{l3NetworkUuid}/dhcp-ip
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/dhcp-ip", params.L3NetworkUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetL3NetworkDhcpIpAddressFailed, err.Error())
	}
	var result model.GetL3NetworkDhcpIpAddressResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//向三层网络添加DNS
func AddDnsToL3Network(params model.AddDnsToL3NetworkRequest) models.Result {
	//POST zstack/v1/l3-networks/{l3NetworkUuid}/dns
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/dns", params.L3NetworkUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddDnsToL3NetworkFailed, err.Error())
	}
	var result model.AddDnsToL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从三层网络移除DNS
func RemoveDnsFromL3Network(params model.RemoveDnsFromL3NetworkRequest) models.Result {
	//DELETE /v1/l3-networks/{l3NetworkUuid}/dns/{dns}
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/dns/%s", params.L3NetworkUuid, params.Dns)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveDnsFromL3NetworkFailed, err.Error())
	}
	var result model.RemoveDnsFromL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//向三层网络添加主机路由
func AddHostRouteToL3Network(params model.AddHostRouteToL3NetworkRequest) models.Result {
	//POST zstack/v1/l3-networks/{l3NetworkUuid}/hostroute
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/hostroute", params.L3NetworkUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddHostRouteToL3NetworkFailed, err.Error())
	}
	var result model.AddHostRouteToL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从三层网络移除主机路由
func RemoveHostRouteFromL3Network(params model.RemoveHostRouteFromL3NetworkRequest) models.Result {
	//DELETE zstack/v1/l3-networks/{l3NetworkUuid}/hostroute?prefix={prefix}
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/hostroute?prefix=%s", params.L3NetworkUuid, params.Prefix)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveHostRouteFromL3NetworkFailed, err.Error())
	}
	var result model.RemoveHostRouteFromL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取空闲IP
func GetFreeIp(params model.GetFreeIpRequest) models.Result {
	//GET zstack/v1/l3-networks/{l3NetworkUuid}/ip/free
	//GET zstack/v1/l3-networks/ip-ranges/{ipRangeUuid}/ip/free
	var url string
	if params.L3NetworkUuid != "" {
		url = fmt.Sprintf("zstack/v1/l3-networks/%s/ip/free", params.L3NetworkUuid)
	} else if params.IpRangeUuid != "" {
		url = fmt.Sprintf("zstack/v1/l3-networks/ip-ranges/%s/ip/free", params.IpRangeUuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetFreeIpFailed, err.Error())
	}
	var result model.GetFreeIpResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//检查IP可用性
func CheckIpAvailability(params model.CheckIpAvailabilityRequest) models.Result {
	//GET zstack/v1/l3-networks/{l3NetworkUuid}/ip/{ip}/availability
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/ip/%s/availability", params.L3NetworkUuid, params.Ip)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.CheckIpAvailabilityFailed, err.Error())
	}
	var result model.CheckIpAvailabilityRequest
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取IP网络地址容量
func GetIpAddressCapacity(params model.GetIpAddressCapacityRequest) models.Result {
	//GET zstack/v1/ip-capacity
	url := fmt.Sprintf("zstack/v1/ip-capacity")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetIpAddressCapacityFailed, err.Error())
	}
	var result model.GetIpAddressCapacityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加IP地址范围
func AddIpRange(params model.AddIpRangeRequest) models.Result {
	//POST zstack/v1/l3-networks/{l3NetworkUuid}/ip-ranges
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/ip-ranges", params.L3NetworkUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddIpRangeFailed, err.Error())
	}
	var result model.AddIpRangeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除IP地址范围
func DeleteIpRange(params model.DeleteIpRangeRequest) models.Result {
	//DELETE /v1/l3-networks/ip-ranges/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/l3-networks/ip-ranges/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteIpRangeFailed, err.Error())
	}
	var result model.DeleteIpRangeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询IP地址范围
func QueryIpRange(params model.QueryIpRangeRequest) models.Result {
	//GET zstack/v1/l3-networks/ip-ranges
	//GET zstack /v1/l3-networks/ip-ranges/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/l3-networks/ip-ranges")
	} else {
		url = fmt.Sprintf("zstack/v1/l3-networks/ip-ranges/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryIpRangeFailed, err.Error())
	}
	var result model.QueryIpRangeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新IP地址范围
func UpdateIpRange(params model.UpdateIpRangeRequest) models.Result {
	//PUT zstack/v1/l3-networks/ip-ranges/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/l3-networks/ip-ranges/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateIpRangeFailed, err.Error())
	}
	var result model.UpdateIpRangeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//通过网络CIDR添加IP地址范围
func AddIpRangeByNetworkCidr(params model.AddIpRangeByNetworkCidrRequest) models.Result {
	//POST zstack/v1/l3-networks/{l3NetworkUuid}/ip-ranges/by-cidr
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/ip-ranges/by-cidr", params.L3NetworkUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.AddIpRangeByNetworkCidrFailed, err.Error())
	}
	var result model.AddIpRangeByNetworkCidrResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取三层网络Mtu值
func GetL3NetworkMtu(params model.GetL3NetworkMtuRequest) models.Result {
	//GET zstack/v1/l3-networks/{l3NetworkUuid}/mtu
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/mtu", params.L3NetworkUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetL3NetworkMtuFailed, err.Error())
	}
	var result model.GetL3NetworkMtuResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//设置三层网络Mtu值
func SetL3NetworkMtu(params model.SetL3NetworkMtuRequest) models.Result {
	//POST zstack/v1/l3-networks/{l3NetworkUuid}/mtu
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/mtu", params.L3NetworkUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.SetL3NetworkMtuFailed, err.Error())
	}
	var result model.SetL3NetworkMtuResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取三层网络上路由器的接口地址
func GetL3NetworkRouterInterfaceIp(params model.GetL3NetworkRouterInterfaceIpRequest) models.Result {
	//GET zstack/v1/l3-networks/{l3NetworkUuid}/router-interface-ip
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/router-interface-ip", params.L3NetworkUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetL3NetworkRouterInterfaceIpFailed, err.Error())
	}
	var result model.GetL3NetworkRouterInterfaceIpResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//设置三层网络路由器接口IP
func SetL3NetworkRouterInterfaceIp(params model.SetL3NetworkRouterInterfaceIpRequest) models.Result {
	//POST zstack/v1/l3-networks/{l3NetworkUuid}/router-interface-ip
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/router-interface-ip", params.L3NetworkUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.SetL3NetworkRouterInterfaceIpFailed, err.Error())
	}
	var result model.SetL3NetworkRouterInterfaceIpResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加IPv6地址范围
func AddIpv6Range(params model.AddIpv6RangeRequest) models.Result {
	//POST zstack/v1/l3-networks/{l3NetworkUuid}/ipv6-ranges
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/ipv6-ranges", params.L3NetworkUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddIpv6RangeFailed, err.Error())
	}
	var result model.AddIpv6RangeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//通过网络CIDR添加IPv6地址范围
func AddIpv6RangeByNetworkCidr(params model.AddIpv6RangeByNetworkCidrRequest) models.Result {
	//POST zstack/v1/l3-networks/{l3NetworkUuid}/ipv6-ranges/by-cidr
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/ipv6-ranges/by-cidr", params.L3NetworkUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddIpv6RangeByNetworkCidrFailed, err.Error())
	}
	var result model.AddIpv6RangeByNetworkCidrResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询IP地址
func QueryIpAddress(params model.QueryIpAddressRequest) models.Result {
	//GET zstack/v1/l3-networks/ip-address
	//GET zstack/v1/l3-networks/ip-address{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/l3-networks/ip-address")
	} else {
		url = fmt.Sprintf("zstack/v1/l3-networks/ip-address/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryIpAddressFailed, err.Error())
	}
	var result model.QueryIpAddressResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取三层网络IP地址使用情况统计
func GetL3NetworkIpStatistic(params model.GetL3NetworkIpStatisticRequest) models.Result {
	//GET zstack/v1/l3-networks/{l3NetworkUuid}/ip-statistic
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/ip-statistic", params.L3NetworkUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetL3NetworkIpStatisticFailed, err.Error())
	}
	var result model.GetL3NetworkIpStatisticResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询IP地址池
func QueryAddressPool(params model.QueryAddressPoolRequest) models.Result {
	//GET zstack/v1/l3-networks/address-pools
	//GET zstack/v1/l3-networks/address-pools/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/l3-networks/address-pools")
	} else {
		url = fmt.Sprintf("zstack/v1/l3-networks/address-pools/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAddressPoolFailed, err.Error())
	}
	var result model.QueryAddressPoolResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
