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

//创建VPC路由器
func CreateVpcVRouter(params model.CreateVpcVRouterRequest) mgresult.Result {
	//POST zstack/v1/vpc/virtual-routers
	url := fmt.Sprintf("zstack/v1/vpc/virtual-routers")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateVpcVRouterFailed, err)
	}
	var result model.CreateVpcVRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询VPC路由器
func QueryVpcRouter(params model.QueryVpcRouterRequest) mgresult.Result {
	//GET zstack/v1/vpc/virtual-routers
	//GET zstack/v1/vpc/virtual-routers/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vpc/virtual-routers")
	} else {
		url = fmt.Sprintf("zstack/v1/vpc/virtual-routers/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVpcRouterFailed, err)
	}
	var result model.QueryVpcRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取VPC路由器可加载的三层网络
func GetAttachableVpcL3Network(params model.GetAttachableVpcL3NetworkRequest) mgresult.Result {
	//POST zstack/v1/vpc/virtual-routers/{uuid}/attachable-vpc-l3s
	url := fmt.Sprintf("zstack/v1/vpc/virtual-routers/%s/attachable-vpc-l3s", params.UUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetAttachableVpcL3NetworkFailed, err)
	}
	var result model.GetAttachableVpcL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取实时流量状态
func GetVpcVRouterDistributedRoutingConnections(params model.GetVpcVRouterDistributedRoutingConnectionsRequest) mgresult.Result {
	//GET zstack/v1/vpc/virtual-routers/{uuid}/tracked-connections
	url := fmt.Sprintf("zstack/v1/vpc/virtual-routers/%s/tracked-connections", params.UUID)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVpcVRouterDistributedRoutingConnectionsFailed, err)
	}
	var result model.GetVpcVRouterDistributedRoutingConnectionsResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取分布式路由是否打开
func GetVpcVRouterDistributedRoutingEnabled(params model.GetVpcVRouterDistributedRoutingEnabledRequest) mgresult.Result {
	//GET zstack/v1/vpc/virtual-routers/{uuid}/distributed-routing
	url := fmt.Sprintf("zstack/v1/vpc/virtual-routers/%s/distributed-routing", params.UUID)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVpcVRouterDistributedRoutingEnabledFailed, err)
	}
	var result model.GetVpcVRouterDistributedRoutingEnabledResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//设置分布式路由开关
func SetVpcVRouterDistributedRoutingEnabled(params model.SetVpcVRouterDistributedRoutingEnabledRequest) mgresult.Result {
	//POST zstack/v1/vpc/virtual-routers/{uuid}/distributed-routing
	url := fmt.Sprintf("zstack/v1/vpc/virtual-routers/%s/distributed-routing", params.UUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVpcVRouterDistributedRoutingEnabledFailed, err)
	}
	var result model.SetVpcVRouterDistributedRoutingEnabledResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//向VPC路由器添加DNS
func AddDnsToVpcRouter(params model.AddDnsToVpcRouterRequest) mgresult.Result {
	//POST zstack/v1/vpc/virtual-routers/{uuid}/dns
	url := fmt.Sprintf("zstack/v1/vpc/virtual-routers/%s/dns", params.UUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddDnsToVpcRouterFailed, err)
	}
	var result model.AddDnsToVpcRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从VPC路由器移除DNS
func RemoveDnsFromVpcRouter(params model.RemoveDnsFromVpcRouterRequest) mgresult.Result {
	//DELETE zstack/v1/vpc/virtual-routers/{uuid}/dns?dns={dns}
	url := fmt.Sprintf("zstack/v1/vpc/virtual-routers/%s/dns?dns=%s", params.UUID, params.Dns)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveDnsFromVpcRouterFailed, err)
	}
	var result model.RemoveDnsFromVpcRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取VPC路由器的网络服务状态
func GetVpcVRouterNetworkServiceState(params model.GetVpcVRouterNetworkServiceStateRequest) mgresult.Result {
	//GET zstack/v1/vpc/virtual-routers/{uuid}/networkservicestate
	url := fmt.Sprintf("zstack/v1/vpc/virtual-routers/%s/networkservicestate", params.UUID)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVpcVRouterNetworkServiceStateFailed, err)
	}
	var result model.GetVpcVRouterNetworkServiceStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//设置VPC路由器的网络服务
func SetVpcVRouterNetworkServiceState(params model.SetVpcVRouterNetworkServiceStateRequest) mgresult.Result {
	//POST zstack/v1/vpc/virtual-routers/{uuid}/networkservicestate
	url := fmt.Sprintf("zstack/v1/vpc/virtual-routers/%s/networkservicestate", params.UUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVpcVRouterNetworkServiceStateFailed, err)
	}
	var result model.SetVpcVRouterNetworkServiceStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新高可用组仲裁ip
func ChangeVpcHaGroupMonitorIps(params model.ChangeVpcHaGroupMonitorIpsRequest) mgresult.Result {
	//PUT zstack/v1/vpc/hagroups/{uuid}/monitorIps
	url := fmt.Sprintf("zstack/v1/vpc/hagroups/%s/monitorIps", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeVpcHaGroupMonitorIpsFailed, err)
	}
	var result model.ChangeVpcHaGroupMonitorIpsResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建高可用组
func CreateVpcHaGroup(params model.CreateVpcHaGroupRequest) mgresult.Result {
	//POST zstack/v1/vpc/hagroups
	url := fmt.Sprintf("zstack/v1/vpc/hagroups")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateVpcHaGroupFailed, err)
	}
	var result model.CreateVpcHaGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除高可用组
func DeleteVpcHaGroup(params model.DeleteVpcHaGroupRequest) mgresult.Result {
	//DELETE zstack/v1/vpc/hagroups/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/vpc/hagroups/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVpcHaGroupFailed, err)
	}
	var result model.DeleteVpcHaGroupRequest
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新高可用组
func UpdateVpcHaGroup(params model.UpdateVpcHaGroupRequest) mgresult.Result {
	//PUT zstack/v1/vpc/hagroups/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vpc/hagroups/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateVpcHaGroupFailed, err)
	}
	var result model.UpdateVpcHaGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询高可用
func QueryVpcHaGroup(params model.QueryVpcHaGroupRequest) mgresult.Result {
	//GET zstack/v1/vpc/hagroups
	//GET zstack/v1/vpc/hagroups/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vpc/hagroups")
	} else {
		url = fmt.Sprintf("zstack/v1/vpc/hagroups/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVpcHaGroupFailed, err)
	}
	var result model.QueryVpcHaGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新虚拟路由器
func UpdateVirtualRouter(params model.UpdateVirtualRouterRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/appliances/virtual-routers/%s/actions", params.VmInstanceUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateVirtualRouterFailed, err)
	}
	var result model.UpdateVirtualRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
