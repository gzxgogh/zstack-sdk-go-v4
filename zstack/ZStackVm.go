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

//创建云主机
func CreateVmInstances(params model.CreateVmInstanceRequest) mgresult.Result {
	url := "zstack/v1/vm-instances"
	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateVmInstancesFailed, err)
	}
	var vmInstanceResp model.CreateVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终的结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//从云盘创建云主机
func CreateVmInstanceFromVolume(params model.CreateVmFromVolumeRequest) mgresult.Result {
	url := "zstack/v1/vm-instances/from/volume"
	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateVmInstanceFromVolumeFailed, err)
	}
	var vmInstanceResp model.CreateVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终的结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//从快照创建云主机
func CreateVmInstanceFromVolumeSnapshot(params model.CreateVmInstanceFromVolumeSnapshotRequest) mgresult.Result {
	//POST zstack/v1/vm-instances/from/volume-snapshots/{volumeSnapshotUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/from/volume-snapshots/%s", params.VolumeSnapshotUuid)
	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateVmInstanceFromVolumeSnapshotFailed, err)
	}
	var vmInstanceResp model.CreateVmInstanceFromVolumeSnapshotResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终的结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//从快照组创建云主机
func CreateVmInstanceFromVolumeSnapshotGroup(params model.CreateVmInstanceFromVolumeSnapshotGroupRequest) mgresult.Result {
	//POST zstack/v1/vm-instances/from/volume-snapshots/group/{volumeSnapshotGroupUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/from/volume-snapshots/group/%s", params.VolumeSnapshotGroupUuid)
	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateVmInstanceFromVolumeSnapshotGroupFailed, err)
	}
	var vmInstanceResp model.CreateVmInstanceFromVolumeSnapshotGroupResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终的结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//删除云主机
func DestroyVmInstance(vmInstance model.DestroyVmInstanceRequest) mgresult.Result {
	//DELETE zstack/v1/vm-instances/{uuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s", vmInstance.UUID)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DestroyVmInstanceFailed, err)
	}
	var vmInstanceResp model.DestroyVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//恢复已删除云主机
func RecoverVmInstance(vmInstance model.RecoverVmInstanceRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RecoverVmInstanceFailed, err)
	}
	var vmInstanceResp model.RecoverVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//彻底删除云主机
func ExpungeVmInstance(vmInstance model.ExpungeVmInstanceRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ExpungeVmInstanceFailed, err)
	}
	var vmInstanceResp model.ExpungeVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//查询云主机
func QueryVmInstance(vmInstance model.QueryVmInstanceRequest) mgresult.Result {
	//GET zstack/v1/vm-instances
	//GET zstack/v1/vm-instances/{uuid}
	var url string
	if vmInstance.UUID == "" {
		url = "zstack/v1/vm-instances"
	} else {
		url = fmt.Sprintf("zstack/v1/vm-instances/%s", vmInstance.UUID)
	}
	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVmInstanceFailed, err)
	}
	var vmInstanceResp model.QueryVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//根据ip查询云主机
func QueryVmByIp(params model.QueryVmByIpRequest) mgresult.Result {

	var data model.QueryVmNicRequest
	data.ReqConfig.Host = params.Host
	data.ReqConfig.AccessKeyId = params.AccessKeyId
	data.ReqConfig.AccessKeySecret = params.AccessKeySecret

	url := fmt.Sprintf("zstack/v1/vm-instances/nics")
	dataStr, err := request.Get(url, data)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVmNicFailed, err)
	}
	var resp model.QueryVmNicResponse
	utils.FromJSON(dataStr, &resp)

	vmUuid := ""
	for _, vmNic := range resp.Inventories {
		if params.Ip == vmNic.IP {
			vmUuid = vmNic.VMInstanceUUID
		}
	}
	if vmUuid == "" {
		return mgresult.Success(nil)
	}

	var vmInstance model.QueryVmInstanceRequest
	vmInstance.ReqConfig.Host = params.Host
	vmInstance.ReqConfig.AccessKeyId = params.AccessKeyId
	vmInstance.ReqConfig.AccessKeySecret = params.AccessKeySecret
	vmInstance.UUID = vmUuid
	result := QueryVmInstance(vmInstance)

	return result
}

//启动云主机
func StartVmInstance(vmInstance model.StartVmInstanceRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.StartVmInstanceFailed, err)
	}
	var vmInstanceResp model.StartVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//停止云主机
func StopVmInstance(vmInstance model.StopVmInstanceRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.StopVmInstanceFailed, err)
	}
	var vmInstanceResp model.StopVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//重启云主机
func RebootVmInstance(vmInstance model.RebootVmInstanceRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RebootVmInstanceFailed, err)
	}
	var vmInstanceResp model.RebootVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//暂停云主机
func PauseVmInstance(vmInstance model.PauseVmInstanceRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.PauseVmInstanceFailed, err)
	}
	var vmInstanceResp model.PauseVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//恢复暂停的云主机
func ResumeVmInstance(vmInstance model.ResumeVmInstanceRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ResumeVmInstanceFailed, err)
	}
	var vmInstanceResp model.ResumeVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//重置云主机,执行该API后，云主机根云盘重置成最初创建的状态，意味着所有后续写入的数据都会被丢失，该操作不可逆！
func ReimageVmInstance(vmInstance model.ReimageVmInstanceRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.VmInstanceUuid)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ReimageVmInstanceFailed, err)
	}
	var vmInstanceResp model.ReimageVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//热迁移云主机
func MigrateVm(vmInstance model.MigrateVmRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.VmInstanceUuid)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.MigrateVmFailed, err)
	}

	var vmInstanceResp model.MigrateVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//查询共享磁盘所挂载的云主机
func QueryShareableVolumeVmInstanceRef(vmInstance model.QueryShareableVolumeVmInstanceRefRequest) mgresult.Result {
	//GET zstack/v1/volumes/vm-instances/refs
	url := "zstack/v1/volumes/vm-instances/refs"
	dataStr, err := request.Get(url, nil)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryShareableVolumeVmInstanceRefFailed, err)
	}

	var vmInstanceResp model.QueryShareableVolumeVmInstanceRefResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取可热迁移的物理机列表
func GetVmMigrationCandidateHosts(vmInstance model.GetVmMigrationCandidateHostsRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/migration-target-hosts
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/migration-target-hosts", vmInstance.VmInstanceUuid)
	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmMigrationCandidateHostsFailed, err)
	}

	var vmInstanceResp model.GetVmMigrationCandidateHostsResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取可选择的主存储
func GetCandidatePrimaryStoragesForCreatingVm(vmInstance model.GetCandidatePrimaryStoragesForCreatingVmRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/candidate-storages
	url := fmt.Sprintf("zstack/v1/vm-instances/candidate-storages")
	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCandidatePrimaryStoragesForCreatingVmFailed, err)
	}
	var vmInstanceResp model.GetCandidatePrimaryStoragesForCreatingVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机可加载ISO列表
func GetCandidateIsoForAttachingVm(vmInstance model.GetCandidateIsoForAttachingVmRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/iso-candidates
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/iso-candidates", vmInstance.VmInstanceUuid)
	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCandidateIsoForAttachingVmFailed, err)
	}

	var vmInstanceResp model.GetCandidateIsoForAttachingVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取ISO可加载云主机列表
func GetCandidateVmForAttachingIso(vmInstance model.GetCandidateVmForAttachingIsoRequest) mgresult.Result {
	//GET zstack/v1/images/iso/{isoUuid}/vm-candidates
	url := fmt.Sprintf("zstack/v1/images/iso/%s/vm-candidates", vmInstance.IsoUuid)
	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCandidateVmForAttachingIsoFailed, err)
	}

	var vmInstanceResp model.GetCandidateVmForAttachingIsoResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//加载ISO到云主机
func AttachIsoToVmInstance(vmInstance model.AttachIsoToVmInstanceRequest) mgresult.Result {
	//POST zstack/v1/vm-instances/{vmInstanceUuid}/iso/{isoUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/iso/%s", vmInstance.VmInstanceUuid, vmInstance.IsoUuid)
	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachIsoToVmInstanceFailed, err)
	}

	var vmInstanceResp model.AttachIsoToVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//卸载云主机上的ISO
func DetachIsoFromVmInstance(vmInstance model.DetachIsoFromVmInstanceRequest) mgresult.Result {
	//DELETE zstack/v1/vm-instances/{vmInstanceUuid}/iso?isoUuid={isoUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s", vmInstance.VmInstanceUuid)
	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachIsoFromVmInstanceFailed, err)
	}

	var vmInstanceResp model.DetachIsoFromVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机可加载云盘列表
func GetVmAttachableDataVolume(vmInstance model.GetVmAttachableDataVolumeRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/data-volume-candidates
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/data-volume-candidates", vmInstance.VmInstanceUuid)
	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmAttachableDataVolumeFailed, err)
	}

	var vmInstanceResp model.GetVmAttachableDataVolumeResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机可加载L3网络列表
func GetVmAttachableL3Network(vmInstance model.GetVmAttachableL3NetworkRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/l3-networks-candidates
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/l3-networks-candidates", vmInstance.VmInstanceUuid)
	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmAttachableL3NetworkFailed, err)
	}

	var vmInstanceResp model.GetVmAttachableL3NetworkResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//加载L3网络到云主机
func AttachL3NetworkToVm(vmInstance model.AttachL3NetworkToVmRequest) mgresult.Result {
	//POST zstack/v1/vm-instances/{vmInstanceUuid}/l3-networks/{l3NetworkUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/l3-networks/%s", vmInstance.VmInstanceUuid, vmInstance.L3NetworkUuid)
	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachL3NetworkToVmFailed, err)
	}

	var vmInstanceResp model.AttachL3NetworkToVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//从云主机卸载网络
func DetachL3NetworkFromVm(vmInstance model.DetachL3NetworkFromVmRequest) mgresult.Result {
	//DELETE zstack/v1/vm-instances/nics/{vmNicUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/nics/%s", vmInstance.VmNicUuid)
	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachL3NetworkFromVmFailed, err)
	}

	var vmInstanceResp model.DetachL3NetworkFromVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//创建云主机网卡
func CreateVmNic(vmInstance model.CreateVmNicRequest) mgresult.Result {
	//POST zstack/v1/nics
	url := fmt.Sprintf("zstack/v1/nics")
	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateVmNicFailed, err)
	}

	var vmInstanceResp model.CreateVmNicResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//加载网卡到云主机
func AttachVmNicToVm(vmInstance model.AttachVmNicToVmRequest) mgresult.Result {
	//POST zstack/v1/vm-instances/{vmInstanceUuid}/nices/{vmNicUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/nices/%s", vmInstance.VmInstanceUuid, vmInstance.VmNicUuid)
	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachVmNicToVmFailed, err)
	}

	var vmInstanceResp model.AttachVmNicToVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//删除云主机网卡
func DeleteVmNic(vmInstance model.DeleteVmNicRequest) mgresult.Result {
	//DELETE zstack/v1/nics/{uuid}
	url := fmt.Sprintf("zstack/v1/nics/%s", vmInstance.Uuid)
	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVmNicFailed, err)
	}

	var vmInstanceResp model.DeleteVmNicResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//查询云主机网卡
func QueryVmNic(vmInstance model.QueryVmNicRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/nics
	//GET zstack/v1/vm-instances/nics/{uuid}
	var url string
	if vmInstance.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/vm-instances/nics")
	} else {
		url = fmt.Sprintf("zstack/v1/vm-instances/nics/%s", vmInstance.Uuid)
	}

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVmNicFailed, err)
	}

	var vmInstanceResp model.QueryVmNicResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取网卡加载的网络服务名称
func GetVmNicAttachedNetworkService(vmInstance model.GetVmNicAttachedNetworkServiceRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/nics/{vmNicUuid}/attached-networkservices
	url := fmt.Sprintf("zstack/v1/vm-instances/nics/%s/attached-networkservices", vmInstance.VmNicUuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmNicAttachedNetworkServiceFailed, err)
	}

	var vmInstanceResp model.GetVmNicAttachedNetworkServiceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//修改云主机网卡三层网络
func ChangeVmNicNetwork(vmInstance model.ChangeVmNicNetworkRequest) mgresult.Result {
	//POST zstack/v1/vm-instances/nics/{vmNicUuid}/l3-networks/{l3NetworkUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/nics/%s/l3-networks/%s", vmInstance.VmNicUuid, vmInstance.DestL3NetworkUuid)

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeVmNicNetworkFailed, err)
	}

	var vmInstanceResp model.ChangeVmNicNetworkResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机网卡可挂载的三层网络
func GetCandidateL3NetworksForChangeVmNicNetwork(vmInstance model.GetCandidateL3NetworksForChangeVmNicNetworkRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/nics/{vmNicUuid}/l3-networks-candidates
	url := fmt.Sprintf("zstack/v1/vm-instances/nics/%s/l3-networks-candidates", vmInstance.VmNicUuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCandidateL3NetworksForChangeVmNicNetworkFailed, err)
	}

	var vmInstanceResp model.GetCandidateL3NetworksForChangeVmNicNetworkResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机网卡限速
func SetNicQoS(vmInstance model.SetNicQoSRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetNicQoSFailed, err)
	}

	var vmInstanceResp model.SetNicQoSResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机网卡限速
func GetNicQoS(vmInstance model.GetNicQoSRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/nic-qos
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/nic-qos", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetNicQoSFailed, err)
	}

	var vmInstanceResp model.GetNicQoSResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//取消云主机网卡限速
func DeleteNicQoS(vmInstance model.DeleteNicQoSRequest) mgresult.Result {
	//DELETE zstack/v1/vm-instances/{uuid}/nic-qos
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/nic-qos", vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteNicQoSFailed, err)
	}

	var vmInstanceResp model.DeleteNicQoSResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取相互依赖的镜像和L3网络
func GetInterdependentL3NetworksImages(vmInstance model.GetInterdependentL3NetworksImagesRequest) mgresult.Result {
	//GET zstack/v1/images-l3networks/dependencies
	url := fmt.Sprintf("zstack/v1/images-l3networks/dependencies")

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetInterdependentL3NetworksImagesFailed, err)
	}

	var vmInstanceResp model.GetInterdependentL3NetworksImagesResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机SSH Key
func SetVmSshKey(vmInstance model.SetVmSshKeyRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmSshKeyFailed, err)
	}

	var vmInstanceResp model.SetVmSshKeyResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机SSH Key
func GetVmSshKey(vmInstance model.GetVmSshKeyRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/ssh-keys
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/ssh-keys", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmSshKeyFailed, err)
	}

	var vmInstanceResp model.GetVmSshKeyResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//删除云主机SSH Key
func DeleteVmSshKey(vmInstance model.DeleteVmSshKeyRequest) mgresult.Result {
	//DELETE zstack/v1/vm-instances/{uuid}/ssh-keys
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/ssh-keys", vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVmSshKeyFailed, err)
	}

	var vmInstanceResp model.DeleteVmSshKeyResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//变更云主机密码
func ChangeVmPassword(vmInstance model.ChangeVmPasswordRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeVmPasswordFailed, err)
	}

	var vmInstanceResp model.ChangeVmPasswordResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机控制台密码
func SetVmConsolePassword(vmInstance model.SetVmConsolePasswordRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmConsolePasswordFailed, err)
	}

	var vmInstanceResp model.SetVmConsolePasswordResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机控制台密码
func GetVmConsolePassword(vmInstance model.GetVmConsolePasswordRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/console-passwords
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/console-passwords", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmConsolePasswordFailed, err)
	}

	var vmInstanceResp model.GetVmConsolePasswordResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//删除云主机控制台密码
func DeleteVmConsolePassword(vmInstance model.DeleteVmConsolePasswordRequest) mgresult.Result {
	//DELETE zstack/v1/vm-instances/{uuid}/console-password
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/console-passwords", vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVmConsolePasswordFailed, err)
	}

	var vmInstanceResp model.DeleteVmConsolePasswordResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机控制台地址和访问协议
func GetVmConsoleAddress(vmInstance model.GetVmConsoleAddressRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/console-addresses
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/console-addresses", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmConsoleAddressFailed, err)
	}

	var vmInstanceResp model.GetVmConsoleAddressResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机Hostname
func SetVmHostname(vmInstance model.SetVmHostnameRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmHostnameFailed, err)
	}

	var vmInstanceResp model.SetVmHostnameResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机Hostname
func GetVmHostname(vmInstance model.GetVmHostnameRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/hostnames
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/hostnames", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmHostnameFailed, err)
	}

	var vmInstanceResp model.GetVmHostnameResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//删除云主机Hostname
func DeleteVmHostname(vmInstance model.DeleteVmHostnameRequest) mgresult.Result {
	//DELETE zstack/v1/vm-instances/{uuid}/hostnames
	url := fmt.Sprintf("zstack/v1/vm-instances/%s", vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVmHostnameFailed, err)
	}

	var vmInstanceResp model.DeleteVmHostnameResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//创建启动云主机的定时任务
func CreateStartVmInstanceScheduler(vmInstance model.CreateStartVmInstanceSchedulerRequest) mgresult.Result {
	//POST zstack/v1/vm-instances/{vmUuid}/schedulers/starting
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/schedulers/starting", vmInstance.VmUuid)

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateStartVmInstanceSchedulerFailed, err)
	}

	var vmInstanceResp model.CreateStartVmInstanceSchedulerResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//创建停止云主机的定时任务
func CreateStopVmInstanceScheduler(vmInstance model.CreateStopVmInstanceSchedulerRequest) mgresult.Result {
	//POST zstack/v1/vm-instances/{vmUuid}/schedulers/stopping
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/schedulers/stopping", vmInstance.VmUuid)

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateStopVmInstanceSchedulerFailed, err)
	}

	var vmInstanceResp model.CreateStopVmInstanceSchedulerResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//创建重启云主机的定时任务
func CreateRebootVmInstanceScheduler(vmInstance model.CreateRebootVmInstanceSchedulerRequest) mgresult.Result {
	//POST zstack/v1/vm-instances/{vmUuid}/schedulers/rebooting
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/schedulers/rebooting", vmInstance.VmUuid)

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateRebootVmInstanceSchedulerFailed, err)
	}

	var vmInstanceResp model.CreateRebootVmInstanceSchedulerResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机启动设备列表
func GetVmBootOrder(vmInstance model.GetVmBootOrderRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/boot-orders
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/boot-orders", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmBootOrderFailed, err)
	}

	var vmInstanceResp model.GetVmBootOrderResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//指定云主机启动设备
func SetVmBootOrder(vmInstance model.SetVmBootOrderRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmBootOrderFailed, err)
	}

	var vmInstanceResp model.SetVmBootOrderResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取目的地列表
func GetCandidateZonesClustersHostsForCreatingVm(vmInstance model.GetCandidateZonesClustersHostsForCreatingVmRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/candidate-destinations
	url := fmt.Sprintf("zstack/v1/vm-instances/candidate-destinations")

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCandidateZonesClustersHostsForCreatingVmFailed, err)
	}

	var vmInstanceResp model.GetCandidateZonesClustersHostsForCreatingVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机可启动目的地列表
func GetVmStartingCandidateClustersHosts(vmInstance model.GetVmStartingCandidateClustersHostsRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/starting-target-hosts
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/starting-target-hosts", vmInstance.UUID)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmStartingCandidateClustersHostsFailed, err)
	}

	var vmInstanceResp model.GetVmStartingCandidateClustersHostsResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//指定云主机IP
func SetVmStaticIp(vmInstance model.SetVmStaticIpRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.VmInstanceUuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmStaticIpFailed, err)
	}

	var vmInstanceResp model.SetVmStaticIpResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//删除云主机指定IP
func DeleteVmStaticIp(vmInstance model.DeleteVmStaticIpRequest) mgresult.Result {
	//DELETE zstack/v1/vm-instances/{vmInstanceUuid}/static-ips?l3NetworkUuid={l3NetworkUuid}&deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/static-ips", vmInstance.VmInstanceUuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVmStaticIpFailed, err)
	}

	var vmInstanceResp model.DeleteVmStaticIpResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机能力
func GetVmCapabilities(vmInstance model.GetVmCapabilitiesRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/capabilities
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/capabilities", vmInstance.UUID)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmCapabilitiesFailed, err)
	}

	var vmInstanceResp model.GetVmCapabilitiesResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//更新云主机信息
func UpdateVmInstance(vmInstance model.UpdateVmInstanceRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateVmInstanceFailed, err)
	}

	var vmInstanceResp model.UpdateVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//克隆云主机到指定物理机
func CloneVmInstance(vmInstance model.CloneVmInstanceRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.VmInstanceUuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CloneVmInstanceFailed, err)
	}

	var vmInstanceResp model.CloneVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机时钟同步
func SetVmClockTrack(vmInstance model.SetVmClockTrackRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmClockTrackFailed, err)
	}

	var vmInstanceResp model.SetVmClockTrackResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机高可用级别
func SetVmInstanceHaLevel(vmInstance model.SetVmInstanceHaLevelRequest) mgresult.Result {
	//POST zstack/v1/vm-instances/{uuid}/ha-levels
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/ha-levels", vmInstance.UUID)

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmInstanceHaLevelFailed, err)
	}

	var vmInstanceResp model.SetVmInstanceHaLevelResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机高可用级别
func GetVmInstanceHaLevel(vmInstance model.GetVmInstanceHaLevelRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/ha-levels
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/ha-levels", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmInstanceHaLevelFailed, err)
	}

	var vmInstanceResp model.GetVmInstanceHaLevelResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//取消云主机高可用
func DeleteVmInstanceHaLevel(vmInstance model.DeleteVmInstanceHaLevelRequest) mgresult.Result {
	//DELETE zstack/v1/instances/{uuid}/ha-levels
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/ha-levels", vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVmInstanceHaLevelFailed, err)
	}

	var vmInstanceResp model.DeleteVmInstanceHaLevelResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机Qga
func GetVmQga(vmInstance model.GetVmQgaRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/qga
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/qga", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmQgaFailed, err)
	}

	var vmInstanceResp model.GetVmQgaResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机Qga
func SetVmQga(vmInstance model.SetVmQgaRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmQgaFailed, err)
	}

	var vmInstanceResp model.SetVmQgaResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机RDP开关状态
func GetVmRDP(vmInstance model.GetVmRDPRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/rdp
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/rdp", vmInstance.UUID)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmRDPFailed, err)
	}

	var vmInstanceResp model.GetVmRDPResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机RDP开关状态
func SetVmRDP(vmInstance model.SetVmRDPRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmRDPFailed, err)
	}

	var vmInstanceResp model.SetVmRDPResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机支持的屏幕数
func GetVmMonitorNumber(vmInstance model.GetVmMonitorNumberRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/monitorNumber
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/monitorNumber", vmInstance.UUID)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmMonitorNumberFailed, err)
	}

	var vmInstanceResp model.GetVmMonitorNumberResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机支持的屏幕数
func SetVmMonitorNumber(vmInstance model.SetVmMonitorNumberRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmMonitorNumberFailed, err)
	}

	var vmInstanceResp model.SetVmMonitorNumberResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//修改云主机根云盘
func ChangeVmImage(vmInstance model.ChangeVmImageRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.VmInstanceUuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeVmImageFailed, err)
	}

	var vmInstanceResp model.ChangeVmImageResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取候选镜像列表
func GetImageCandidatesForVmToChange(vmInstance model.GetImageCandidatesForVmToChangeRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/image-candidates
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/image-candidates", vmInstance.VmInstanceUuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetImageCandidatesForVmToChangeFailed, err)
	}

	var vmInstanceResp model.GetImageCandidatesForVmToChangeResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//更新云主机mac地址
func UpdateVmNicMac(vmInstance model.UpdateVmNicMacRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/nics/{vmNicUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/nics/%s/actions", vmInstance.VmNicUuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateVmNicMacFailed, err)
	}

	var vmInstanceResp model.UpdateVmNicMacResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机防IP欺骗启用状态
func SetVmCleanTraffic(vmInstance model.SetVmCleanTrafficRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmCleanTrafficFailed, err)
	}

	var vmInstanceResp model.SetVmCleanTrafficResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//为云主机创建CDROM
func CreateVmCdRom(vmInstance model.CreateVmCdRomRequest) mgresult.Result {
	//POST zstack/v1/vm-instances/cdroms
	url := fmt.Sprintf("zstack/v1/vm-instances/cdroms")

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateVmCdRomFailed, err)
	}

	var vmInstanceResp model.CreateVmCdRomResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//删除CDROM
func DeleteVmCdRom(vmInstance model.DeleteVmCdRomRequest) mgresult.Result {
	//DELETE zstack/v1/vm-instances/cdroms/{uuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/cdroms/%s", vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVmCdRomFailed, err)
	}

	var vmInstanceResp model.DeleteVmCdRomResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//修改CDROM
func UpdateVmCdRom(params model.UpdateVmCdRomRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/cdroms/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/cdroms/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateVmCdRomFailed, err)
	}
	var result model.UpdateVmCdRomResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//设置云主机默认CDROM
func SetVmInstanceDefaultCdRom(vmInstance model.SetVmInstanceDefaultCdRomRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/cdroms/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/cdroms/%s/actions", vmInstance.VmInstanceUuid, vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmInstanceDefaultCdRomFailed, err)
	}

	var vmInstanceResp model.SetVmInstanceDefaultCdRomResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//查询CDROM清单
func QueryVmCdRom(vmInstance model.QueryVmCdRomRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/cdroms
	//GET zstack/v1/vm-instances/cdroms/{uuid}
	var url string
	if vmInstance.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/vm-instances/cdroms")
	} else {
		url = fmt.Sprintf("zstack/v1/vm-instances/cdroms/%s", vmInstance.Uuid)
	}

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVmCdRomFailed, err)
	}

	var vmInstanceResp model.QueryVmCdRomResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//更改云主机优先级级别
func UpdateVmPriority(vmInstance model.UpdateVmPriorityRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateVmPriorityFailed, err)
	}

	var vmInstanceResp model.UpdateVmPriorityResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机显存
func SetVmQxlMemory(vmInstance model.SetVmQxlMemoryRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmQxlMemoryFailed, err)
	}

	var vmInstanceResp model.SetVmQxlMemoryResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机虚拟声卡类型
func SetVmSoundType(vmInstance model.SetVmSoundTypeRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmSoundTypeFailed, err)
	}

	var vmInstanceResp model.SetVmSoundTypeResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取spice的CA证书
func GetSpiceCertificates(vmInstance model.GetSpiceCertificatesRequest) mgresult.Result {
	//GET zstack/v1/spice/certificates
	url := fmt.Sprintf("zstack/v1/spice/certificates")

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetSpiceCertificatesFailed, err)
	}

	var vmInstanceResp model.GetSpiceCertificatesResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//为云主机加载增强工具镜像
func AttachGuestToolsIsoToVm(vmInstance model.AttachGuestToolsIsoToVmRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachGuestToolsIsoToVmFailed, err)
	}

	var vmInstanceResp model.AttachGuestToolsIsoToVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机可用的最新增强工具
func GetLatestGuestToolsForVm(vmInstance model.GetLatestGuestToolsForVmRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/latest-guest-tools
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/latest-guest-tools", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetLatestGuestToolsForVmFailed, err)
	}

	var vmInstanceResp model.GetLatestGuestToolsForVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机内部增强工具的信息
func GetVmGuestToolsInfo(vmInstance model.GetVmGuestToolsInfoRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/guest-tools-infos
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/guest-tools-infos", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmGuestToolsInfoFailed, err)
	}

	var vmInstanceResp model.GetVmGuestToolsInfoResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机第一启动项
func GetVmInstanceFirstBootDevice(vmInstance model.GetVmInstanceFirstBootDeviceRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/first-boot-device
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/first-boot-device", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmInstanceFirstBootDeviceFailed, err)
	}

	var vmInstanceResp model.GetVmInstanceFirstBootDeviceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置网卡型号
func UpdateVmNicDriver(vmInstance model.UpdateVmNicDriverRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions?vmNicUuid={vmNicUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.VmInstanceUuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateVmNicDriverFailed, err)
	}

	var vmInstanceResp model.UpdateVmNicDriverResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取云主机设备地址
func GetVmDeviceAddress(vmInstance model.GetVmDeviceAddressRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/devices
	url := fmt.Sprintf("zstack/v1/vm-instances/devices")

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmDeviceAddressFailed, err)
	}

	var vmInstanceResp model.GetVmDeviceAddressResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//批量获取云主机能力
func GetVmsCapabilities(vmInstance model.GetVmsCapabilitiesRequest) mgresult.Result {
	//POST zstack/v1/vm-instances/capabilities
	url := fmt.Sprintf("zstack/v1/vm-instances/capabilities")

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmsCapabilitiesFailed, err)
	}

	var vmInstanceResp model.GetVmsCapabilitiesResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机vNUMA
func SetVmNuma(vmInstance model.SetVmNumaRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmNumaFailed, err)
	}

	var vmInstanceResp model.SetVmNumaResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取VM Numa开关状态
func GetVmNuma(vmInstance model.GetVmNumaRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/vnuma
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/vnuma", vmInstance.UUID)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmNumaFailed, err)
	}

	var vmInstanceResp model.GetVmNumaResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//查询云主机的vNUMA拓扑信息
func GetVmvNUMATopology(vmInstance model.GetVmvNUMATopologyRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/vnuma-topology
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/vnuma-topology", vmInstance.UUID)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmvNUMATopologyFailed, err)
	}

	var vmInstanceResp model.GetVmvNUMATopologyResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//设置云主机Emulator Pinning功能
func SetVmEmulatorPinning(vmInstance model.SetVmEmulatorPinningRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmEmulatorPinningFailed, err)
	}

	var vmInstanceResp model.SetVmEmulatorPinningResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}

//获取VM Emulator Pin在那些物理机的cpu上
func GetVmEmulatorPinning(vmInstance model.GetVmEmulatorPinningRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/emulator-pinning
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/emulator-pinning", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmEmulatorPinningFailed, err)
	}

	var vmInstanceResp model.GetVmEmulatorPinningResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return mgresult.Success(vmInstanceResp)
}
