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

//查询物理机
func QueryHost(params model.QueryHostRequest) mgresult.Result {
	//GET zstack/v1/hosts
	//GET zstack/v1/hosts/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/hosts")
	} else {
		url = fmt.Sprintf("zstack/v1/hosts/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryHostFailed, err)
	}
	var result model.QueryHostResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除物理机
func DeleteHost(params model.DeleteHostRequest) mgresult.Result {
	//DELETE zstack/v1/hosts/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/hosts/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteHostFailed, err)
	}
	var result model.DeleteHostResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新物理机信息
func UpdateHost(params model.UpdateHostRequest) mgresult.Result {
	//PUT zstack/v1/hosts/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/hosts/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateHostFailed, err)
	}
	var result model.UpdateHostResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新物理机启用状态
func ChangeHostState(params model.ChangeHostStateRequest) mgresult.Result {
	//PUT zstack/v1/hosts/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/hosts/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeHostStateFailed, err)
	}
	var result model.ChangeHostStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//重连物理机
func ReconnectHost(params model.ReconnectHostRequest) mgresult.Result {
	//PUT zstack/v1/hosts/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/hosts/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ReconnectHostFailed, err)
	}
	var result model.ReconnectHostResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取物理机分配策略
func GetHostAllocatorStrategies(params model.GetHostAllocatorStrategiesRequest) mgresult.Result {
	//GET zstack/v1/hosts/allocators/strategies
	url := fmt.Sprintf("zstack/v1/hosts/allocators/strategies")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetHostAllocatorStrategiesFailed, err)
	}
	var result model.GetHostAllocatorStrategiesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取云主机虚拟化技术类型
func GetHypervisorTypes(params model.GetHypervisorTypesRequest) mgresult.Result {
	//GET zstack/v1/hosts/hypervisor-types
	url := fmt.Sprintf("zstack/v1/hosts/hypervisor-types")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetHypervisorTypesFailed, err)
	}
	var result model.GetHypervisorTypesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新KVM主机信息
func UpdateKVMHost(params model.UpdateKVMHostRequest) mgresult.Result {
	//PUT zstack/v1/hosts/kvm/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/hosts/kvm/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateKVMHostFailed, err)
	}
	var result model.UpdateKVMHostResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加KVM主机
func AddKVMHost(params model.AddKVMHostRequest) mgresult.Result {
	//POST zstack/v1/hosts/kvm
	url := fmt.Sprintf("zstack/v1/hosts/kvm")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddKVMHostFailed, err)
	}
	var result model.AddKVMHostResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//KVM运行Shell命令
func KvmRunShell(params model.KvmRunShellRequest) mgresult.Result {
	//PUT zstack/v1/hosts/kvm/actions
	url := fmt.Sprintf("zstack/v1/hosts/kvm/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.KvmRunShellFailed, err)
	}
	var result model.KvmRunShellResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//通过文件导入方式添加物理机
func AddKVMHostFromConfigFile(params model.AddKVMHostFromConfigFileRequest) mgresult.Result {
	//POST zstack/v1/hosts/kvm/from-file
	url := fmt.Sprintf("zstack/v1/hosts/kvm/from-file")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddKVMHostFromConfigFileFailed, err)
	}
	var result model.AddKVMHostFromConfigFileResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加物理机文件语法检查
func CheckKVMHostConfigFile(params model.CheckKVMHostConfigFileRequest) mgresult.Result {
	//POST /v1/hosts/kvm/from-file/check
	url := fmt.Sprintf("zstack/v1/hosts/kvm/from-file/check")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CheckKVMHostConfigFileFailed, err)
	}
	var result model.CheckKVMHostConfigFileResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取物理机物理网络信息
func GetHostNetworkFacts(params model.GetHostNetworkFactsRequest) mgresult.Result {
	//GET zstack/v1/hosts/network-facts/{hostUuid}
	url := fmt.Sprintf("zstack/v1/hosts/network-facts/%s", params.HostUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetHostNetworkFactsFailed, err)
	}
	var result model.GetHostNetworkFactsResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询物理机Bond信息
func QueryHostNetworkBonding(params model.QueryHostNetworkBondingRequest) mgresult.Result {
	//GET zstack/v1/hosts/bondings
	//GET zstack/v1/hosts/bondings/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/hosts/bondings")
	} else {
		url = fmt.Sprintf("zstack/v1/hosts/bondings/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryHostNetworkBondingFailed, err)
	}
	var result model.QueryHostNetworkBondingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询物理机网卡信息
func QueryHostNetworkInterface(params model.QueryHostNetworkInterfaceRequest) mgresult.Result {
	//GET zstack/v1/hosts/nics
	//GET zstack/v1/hosts/nics/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/hosts/nics")
	} else {
		url = fmt.Sprintf("zstack/v1/hosts/nics/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryHostNetworkInterfaceFailed, err)
	}
	var result model.QueryHostNetworkInterfaceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询物理机的NUMA拓扑信息
func GetHostNUMATopology(params model.GetHostNUMATopologyRequest) mgresult.Result {
	//POST zstack/v1/hosts/{uuid}/numa
	url := fmt.Sprintf("zstack/v1/hosts/%s/numa", params.Uuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetHostNUMATopologyFailed, err)
	}
	var result model.GetHostNUMATopologyResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//根据云主机需求获取物理机资源的分配信息
func GetHostResourceAllocation(params model.GetHostResourceAllocationRequest) mgresult.Result {
	//GET zstack/v1/hosts/nics
	//GET zstack/v1/hosts/nics/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/hosts/nics")
	} else {
		url = fmt.Sprintf("zstack/v1/hosts/nics/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetHostResourceAllocationFailed, err)
	}
	var result model.GetHostResourceAllocationResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
