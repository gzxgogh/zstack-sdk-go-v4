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

//创建VXLAN网络池
func CreateL2VxlanNetworkPool(params model.CreateL2VxlanNetworkPoolRequest) mgresult.Result {
	//POST zstack/v1/l2-networks/vxlan-pool
	url := fmt.Sprintf("zstack/v1/l2-networks/vxlan-pool")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateL2VxlanNetworkPoolFailed, err)
	}
	var result model.CreateL2VxlanNetworkPoolResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询VXLAN网络池
func QueryL2VxlanNetworkPool(params model.QueryL2VxlanNetworkPoolRequest) mgresult.Result {
	//GET zstack/v1/l2-networks/vxlan-pool
	//GET zstack/v1/l2-networks/vxlan-pool/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/l2-networks/vxlan-pool")
	} else {
		url = fmt.Sprintf("zstack/v1/l2-networks/vxlan-pool/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryL2VxlanNetworkPoolFailed, err)
	}
	var result model.QueryL2VxlanNetworkPoolResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建VXLAN网络池
func CreateL2VxlanNetwork(params model.CreateL2VxlanNetworkRequest) mgresult.Result {
	//POST zstack/v1/l2-networks/vxlan
	url := fmt.Sprintf("zstack/v1/l2-networks/vxlan")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateL2VxlanNetworkFailed, err)
	}
	var result model.CreateL2VxlanNetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询VXLAN网络
func QueryL2VxlanNetwork(params model.QueryL2VxlanNetworkRequest) mgresult.Result {
	//GET zstack/v1/l2-networks/vxlan
	//GET zstack/v1/l2-networks/vxlan/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/l2-networks/vxlan")
	} else {
		url = fmt.Sprintf("zstack/v1/l2-networks/vxlan/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryL2VxlanNetworkFailed, err)
	}
	var result model.QueryL2VxlanNetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建普通二层网络
func CreateL2NoVlanNetwork(params model.CreateL2NoVlanNetworkRequest) mgresult.Result {
	//POST zstack/v1/l2-networks/no-vlan
	url := fmt.Sprintf("zstack/v1/l2-networks/no-vlan")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateL2NoVlanNetworkFailed, err)
	}
	var result model.CreateL2NoVlanNetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建二层VLAN网络
func CreateL2VlanNetwork(params model.CreateL2VlanNetworkRequest) mgresult.Result {
	//POST zstack/v1/l2-networks/vlan
	url := fmt.Sprintf("zstack/v1/l2-networks/vlan")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateL2VlanNetworkFailed, err)
	}
	var result model.CreateL2VlanNetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询二层VLAN网络
func QueryL2VlanNetwork(params model.QueryL2VlanNetworkRequest) mgresult.Result {
	//GET zstack/v1/l2-networks/vlan
	//GET zstack/v1/l2-networks/vlan/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/l2-networks/vlan")
	} else {
		url = fmt.Sprintf("zstack/v1/l2-networks/vlan/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryL2VlanNetworkFailed, err)
	}
	var result model.QueryL2VlanNetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除二层网络
func DeleteL2Network(params model.DeleteL2NetworkRequest) mgresult.Result {
	//DELETE zstack/v1/l2-networks/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/l2-networks/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteL2NetworkFailed, err)
	}
	var result model.DeleteL2NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询二层网络
func QueryL2Network(params model.QueryL2NetworkRequest) mgresult.Result {
	//GET zstack/v1/l2-networks
	//GET zstack/v1/l2-networks/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/l2-networks")
	} else {
		url = fmt.Sprintf("zstack/v1/l2-networks/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryL2NetworkFailed, err)
	}
	var result model.QueryL2NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新二层网络
func UpdateL2Network(params model.UpdateL2NetworkRequest) mgresult.Result {
	//PUT zstack/v1/l2-networks/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/l2-networks/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateL2NetworkFailed, err)
	}
	var result model.UpdateL2NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取二层网络类型
func GetL2NetworkTypes(params model.GetL2NetworkTypesRequest) mgresult.Result {
	//GET zstack/v1/l2-networks/types
	url := fmt.Sprintf("zstack/v1/l2-networks/types")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetL2NetworkTypesFailed, err)
	}
	var result model.GetL2NetworkTypesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//挂载二层网络到集群
func AttachL2NetworkToCluster(params model.AttachL2NetworkToClusterRequest) mgresult.Result {
	//POST zstack/v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}
	url := fmt.Sprintf("zstack/v1/l2-networks/%s/clusters/%s", params.L2NetworkUuid, params.ClusterUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachL2NetworkToClusterFailed, err)
	}
	var result model.AttachL2NetworkToClusterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从集群上卸载二层网络
func DetachL2NetworkFromCluster(params model.DetachL2NetworkFromClusterRequest) mgresult.Result {
	//DELETE zstack/v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}
	url := fmt.Sprintf("zstack/v1/l2-networks/%s/clusters/%s", params.L2NetworkUuid, params.ClusterUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachL2NetworkFromClusterFailed, err)
	}
	var result model.DetachL2NetworkFromClusterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建Vni Range
func CreateVniRange(params model.CreateVniRangeRequest) mgresult.Result {
	//POST zstack/v1/l2-networks/vxlan-pool/{l2NetworkUuid}/vni-ranges
	url := fmt.Sprintf("zstack/v1/l2-networks/vxlan-pool/%s/vni-ranges", params.L2NetworkUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateVniRangeFailed, err)
	}
	var result model.CreateVniRangeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询Vni Range
func QueryVniRange(params model.QueryVniRangeRequest) mgresult.Result {
	//GET zstack/v1/l2-networks/vxlan-pool/vni-range
	//GET zstack/v1/l2-networks/vxlan-pool/vni-range/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/l2-networks/vxlan-pool/vni-range")
	} else {
		url = fmt.Sprintf("zstack/v1/l2-networks/vxlan-pool/vni-range/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVniRangeFailed, err)
	}
	var result model.QueryVniRangeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除Vni  Range
func DeleteVniRange(params model.DeleteVniRangeRequest) mgresult.Result {
	//DELETE zstack/v1/l2-networks/vxlan-pool/vni-ranges/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/l2-networks/vxlan-pool/vni-ranges/%s", params.Uuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVniRangeFailed, err)
	}
	var result model.DeleteVniRangeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//修改Vni  Range
func UpdateVniRange(params model.UpdateVniRangeRequest) mgresult.Result {
	//PUT zstack/v1/l2-networks/vxlan-pool/vni-ranges/{uuid}
	url := fmt.Sprintf("zstack/v1/l2-networks/vxlan-pool/vni-ranges/%s", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateVniRangeFailed, err)
	}
	var result model.UpdateVniRangeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
