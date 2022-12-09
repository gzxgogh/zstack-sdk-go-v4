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

//重连VPC路由器
func ReconnectVirtualRouter(params model.ReconnectVirtualRouterRequest) models.Result {
	//PUT zstack/v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/appliances/virtual-routers/%s/actions", params.VmInstanceUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ReconnectVirtualRouterFailed, err.Error())
	}
	var result model.ReconnectVirtualRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询云路由器
func QueryVirtualRouterVm(params model.QueryVirtualRouterVmRequest) models.Result {
	//GET zstack/v1/vm-instances/appliances/virtual-routers
	//GET zstack/v1/vm-instances/appliances/virtual-routers/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/vm-instances/appliances/virtual-routers")
	} else {
		url = fmt.Sprintf("zstack/v1/vm-instances/appliances/virtual-routers/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVirtualRouterVmFailed, err.Error())
	}
	var result model.QueryVirtualRouterVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询服务云主机
func QueryApplianceVm(params model.QueryApplianceVmRequest) models.Result {
	//GET zstack/v1/vm-instances/appliances
	//GET zstack/v1/vm-instances/appliances/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/vm-instances/appliances")
	} else {
		url = fmt.Sprintf("zstack/v1/vm-instances/appliances/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryApplianceVmFailed, err.Error())
	}
	var result model.QueryApplianceVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建云路由规格
func CreateVirtualRouterOffering(params model.CreateVirtualRouterOfferingRequest) models.Result {
	//POST zstack/v1/instance-offerings/virtual-routers
	url := fmt.Sprintf("zstack/v1/instance-offerings/virtual-routers")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateVirtualRouterOfferingFailed, err.Error())
	}
	var result model.CreateVirtualRouterOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询云路由规格
func QueryVirtualRouterOffering(params model.QueryVirtualRouterOfferingRequest) models.Result {
	//GET zstack/v1/instance-offerings/virtual-routers
	//GET zstack/v1/instance-offerings/virtual-routers/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/instance-offerings/virtual-routers")
	} else {
		url = fmt.Sprintf("zstack/v1/instance-offerings/virtual-routers/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVirtualRouterOfferingFailed, err.Error())
	}
	var result model.QueryVirtualRouterOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新云路由规格
func UpdateVirtualRouterOffering(params model.UpdateVirtualRouterOfferingRequest) models.Result {
	//PUT zstack/v1/instance-offerings/virtual-routers/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/instance-offerings/virtual-routers/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateVirtualRouterOfferingFailed, err.Error())
	}
	var result model.UpdateVirtualRouterOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取云路由器可加载外部网络
func GetAttachablePublicL3ForVRouter(params model.GetAttachablePublicL3ForVRouterRequest) models.Result {
	//GET zstack/v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/attachable-public-l3s
	url := fmt.Sprintf("zstack/v1/vm-instances/appliances/virtual-routers/%s/attachable-public-l3s", params.VmInstanceUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetAttachablePublicL3ForVRouterFailed, err.Error())
	}
	var result model.GetAttachablePublicL3ForVRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建云路由路由表
func CreateVRouterRouteTable(params model.CreateVRouterRouteTableRequest) models.Result {
	//POST zstack/v1/vrouter-route-tables
	url := fmt.Sprintf("zstack/v1/vrouter-route-tables")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateVRouterRouteTableFailed, err.Error())
	}
	var result model.CreateVRouterRouteTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除云路由路由表
func DeleteVRouterRouteTable(params model.DeleteVRouterRouteTableRequest) models.Result {
	//DELETE zstack/v1/vrouter-route-tables/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/vrouter-route-tables/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteVRouterRouteTableFailed, err.Error())
	}
	var result model.DeleteVRouterRouteTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询云路由路由表
func QueryVRouterRouteTable(params model.QueryVRouterRouteTableRequest) models.Result {
	//GET zstack/v1/vrouter-route-tables
	//GET zstack/v1/vrouter-route-tables/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vrouter-route-tables")
	} else {
		url = fmt.Sprintf("zstack/v1/vrouter-route-tables/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVRouterRouteTableFailed, err.Error())
	}
	var result model.QueryVRouterRouteTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取路由器实时路由表
func GetVRouterRouteTable(params model.GetVRouterRouteTableRequest) models.Result {
	//GET zstack/v1/vrouter-route-tables/vrouter/{virtualRouterVmUuid}
	url := fmt.Sprintf("zstack/v1/vrouter-route-tables/vrouter/%s", params.VirtualRouterVmUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetVRouterRouteTableFailed, err.Error())
	}
	var result model.GetVRouterRouteTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加云路由路由条目
func AddVRouterRouteEntry(params model.AddVRouterRouteEntryRequest) models.Result {
	//POST zstack/v1/vrouter-route-tables/{routeTableUuid}/route-entries
	url := fmt.Sprintf("zstack/v1/vrouter-route-tables/%s/route-entries", params.RouteTableUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddVRouterRouteEntryFailed, err.Error())
	}
	var result model.AddVRouterRouteEntryResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除云路由路由条目
func DeleteVRouterRouteEntry(params model.DeleteVRouterRouteEntryRequest) models.Result {
	//DELETE zstack/v1/vrouter-route-tables/{routeTableUuid}/route-entries/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/vrouter-route-tables/%s/route-entries/%s", params.RouteTableUuid, params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteVRouterRouteEntryFailed, err.Error())
	}
	var result model.DeleteVRouterRouteEntryResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询云路由路由条目
func QueryVRouterRouteEntry(params model.QueryVRouterRouteEntryRequest) models.Result {
	//GET zstack/v1/vrouter-route-tables/route-entries
	//GET zstack/v1/vrouter-route-tables/route-entries/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vrouter-route-tables/route-entries")
	} else {
		url = fmt.Sprintf("zstack/v1/vrouter-route-tables/route-entries/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVRouterRouteEntryFailed, err.Error())
	}
	var result model.QueryVRouterRouteEntryResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//绑定云路由路由表到云路由器
func AttachVRouterRouteTableToVRouter(params model.AttachVRouterRouteTableToVRouterRequest) models.Result {
	//POST zstack/v1/vrouter-route-tables/{routeTableUuid}/attach
	url := fmt.Sprintf("zstack/v1/vrouter-route-tables/%s/attach", params.RouteTableUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachVRouterRouteTableToVRouterFailed, err.Error())
	}
	var result model.AttachVRouterRouteTableToVRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//解绑云路由路由表
func DetachVRouterRouteTableFromVRouter(params model.DetachVRouterRouteTableFromVRouterRequest) models.Result {
	//DELETE zstack/v1/vrouter-route-tables/{routeTableUuid}/detach/{virtualRouterVmUuid}
	url := fmt.Sprintf("zstack/v1/vrouter-route-tables/%s/detach/%s", params.RouteTableUuid, params.VirtualRouterVmUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachVRouterRouteTableFromVRouterFailed, err.Error())
	}
	var result model.DetachVRouterRouteTableFromVRouterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询绑定关系
func QueryVirtualRouterVRouterRouteTableRef(params model.QueryVirtualRouterVRouterRouteTableRefRequest) models.Result {
	//GET zstack/v1/vrouter-route-tables/virtual-router-refs
	url := fmt.Sprintf("zstack/v1/vrouter-route-tables/virtual-router-refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVirtualRouterVRouterRouteTableRefFailed, err.Error())
	}
	var result model.QueryVirtualRouterVRouterRouteTableRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
