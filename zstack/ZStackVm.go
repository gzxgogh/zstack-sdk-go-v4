package zstack

import (
	"fmt"
	"github.com/gzxgogh/zstack-sdk-go-v4/errcode"
	"github.com/gzxgogh/zstack-sdk-go-v4/model"
	"github.com/gzxgogh/zstack-sdk-go-v4/request"
	"github.com/maczh/mgin/models"

	"github.com/maczh/mgin/logs"
	"github.com/maczh/mgin/utils"
)

//创建云主机
func CreateVmInstances(params model.CreateVmInstanceRequest) models.Result {
	url := "zstack/v1/vm-instances"
	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateVmInstancesFailed, err.Error())
	}
	var vmInstanceResp model.CreateVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终的结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//从云盘创建云主机
func CreateVmInstanceFromVolume(params model.CreateVmFromVolumeRequest) models.Result {
	url := "zstack/v1/vm-instances/from/volume"
	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateVmInstanceFromVolumeFailed, err.Error())
	}
	var vmInstanceResp model.CreateVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终的结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//从快照创建云主机
func CreateVmInstanceFromVolumeSnapshot(params model.CreateVmInstanceFromVolumeSnapshotRequest) models.Result {
	//POST zstack/v1/vm-instances/from/volume-snapshots/{volumeSnapshotUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/from/volume-snapshots/%s", params.VolumeSnapshotUuid)
	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateVmInstanceFromVolumeSnapshotFailed, err.Error())
	}
	var vmInstanceResp model.CreateVmInstanceFromVolumeSnapshotResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终的结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//从快照组创建云主机
func CreateVmInstanceFromVolumeSnapshotGroup(params model.CreateVmInstanceFromVolumeSnapshotGroupRequest) models.Result {
	//POST zstack/v1/vm-instances/from/volume-snapshots/group/{volumeSnapshotGroupUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/from/volume-snapshots/group/%s", params.VolumeSnapshotGroupUuid)
	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateVmInstanceFromVolumeSnapshotGroupFailed, err.Error())
	}
	var vmInstanceResp model.CreateVmInstanceFromVolumeSnapshotGroupResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终的结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//删除云主机
func DestroyVmInstance(vmInstance model.DestroyVmInstanceRequest) models.Result {
	//DELETE zstack/v1/vm-instances/{uuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s", vmInstance.UUID)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return models.Error(errcode.DestroyVmInstanceFailed, err.Error())
	}
	var vmInstanceResp model.DestroyVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//恢复已删除云主机
func RecoverVmInstance(vmInstance model.RecoverVmInstanceRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.RecoverVmInstanceFailed, err.Error())
	}
	var vmInstanceResp model.RecoverVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//彻底删除云主机
func ExpungeVmInstance(vmInstance model.ExpungeVmInstanceRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.ExpungeVmInstanceFailed, err.Error())
	}
	var vmInstanceResp model.ExpungeVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//查询云主机
func QueryVmInstance(vmInstance model.QueryVmInstanceRequest) models.Result {
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
		return models.Error(errcode.QueryVmInstanceFailed, err.Error())
	}
	var vmInstanceResp model.QueryVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//根据ip查询云主机
func QueryVmByIp(params model.QueryVmByIpRequest) models.Result {

	var data model.QueryVmNicRequest
	data.ReqConfig.Host = params.Host
	data.ReqConfig.AccessKeyId = params.AccessKeyId
	data.ReqConfig.AccessKeySecret = params.AccessKeySecret

	url := fmt.Sprintf("zstack/v1/vm-instances/nics")
	dataStr, err := request.Get(url, data)
	if err != nil {
		return models.Error(errcode.QueryVmNicFailed, err.Error())
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
		return models.Success(nil)
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
func StartVmInstance(vmInstance model.StartVmInstanceRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.StartVmInstanceFailed, err.Error())
	}
	var vmInstanceResp model.StartVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//停止云主机
func StopVmInstance(vmInstance model.StopVmInstanceRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.StopVmInstanceFailed, err.Error())
	}
	var vmInstanceResp model.StopVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//重启云主机
func RebootVmInstance(vmInstance model.RebootVmInstanceRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.RebootVmInstanceFailed, err.Error())
	}
	var vmInstanceResp model.RebootVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//暂停云主机
func PauseVmInstance(vmInstance model.PauseVmInstanceRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.PauseVmInstanceFailed, err.Error())
	}
	var vmInstanceResp model.PauseVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//恢复暂停的云主机
func ResumeVmInstance(vmInstance model.ResumeVmInstanceRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.ResumeVmInstanceFailed, err.Error())
	}
	var vmInstanceResp model.ResumeVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//重置云主机,执行该API后，云主机根云盘重置成最初创建的状态，意味着所有后续写入的数据都会被丢失，该操作不可逆！
func ReimageVmInstance(vmInstance model.ReimageVmInstanceRequest) models.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.VmInstanceUuid)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.ReimageVmInstanceFailed, err.Error())
	}
	var vmInstanceResp model.ReimageVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//热迁移云主机
func MigrateVm(vmInstance model.MigrateVmRequest) models.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.VmInstanceUuid)
	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.MigrateVmFailed, err.Error())
	}

	var vmInstanceResp model.MigrateVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//查询共享磁盘所挂载的云主机
func QueryShareableVolumeVmInstanceRef(vmInstance model.QueryShareableVolumeVmInstanceRefRequest) models.Result {
	//GET zstack/v1/volumes/vm-instances/refs
	url := "zstack/v1/volumes/vm-instances/refs"
	dataStr, err := request.Get(url, nil)
	if err != nil {
		return models.Error(errcode.QueryShareableVolumeVmInstanceRefFailed, err.Error())
	}

	var vmInstanceResp model.QueryShareableVolumeVmInstanceRefResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取可热迁移的物理机列表
func GetVmMigrationCandidateHosts(vmInstance model.GetVmMigrationCandidateHostsRequest) models.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/migration-target-hosts
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/migration-target-hosts", vmInstance.VmInstanceUuid)
	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmMigrationCandidateHostsFailed, err.Error())
	}

	var vmInstanceResp model.GetVmMigrationCandidateHostsResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取可选择的主存储
func GetCandidatePrimaryStoragesForCreatingVm(vmInstance model.GetCandidatePrimaryStoragesForCreatingVmRequest) models.Result {
	//GET zstack/v1/vm-instances/candidate-storages
	url := fmt.Sprintf("zstack/v1/vm-instances/candidate-storages")
	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetCandidatePrimaryStoragesForCreatingVmFailed, err.Error())
	}
	var vmInstanceResp model.GetCandidatePrimaryStoragesForCreatingVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机可加载ISO列表
func GetCandidateIsoForAttachingVm(vmInstance model.GetCandidateIsoForAttachingVmRequest) models.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/iso-candidates
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/iso-candidates", vmInstance.VmInstanceUuid)
	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetCandidateIsoForAttachingVmFailed, err.Error())
	}

	var vmInstanceResp model.GetCandidateIsoForAttachingVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取ISO可加载云主机列表
func GetCandidateVmForAttachingIso(vmInstance model.GetCandidateVmForAttachingIsoRequest) models.Result {
	//GET zstack/v1/images/iso/{isoUuid}/vm-candidates
	url := fmt.Sprintf("zstack/v1/images/iso/%s/vm-candidates", vmInstance.IsoUuid)
	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetCandidateVmForAttachingIsoFailed, err.Error())
	}

	var vmInstanceResp model.GetCandidateVmForAttachingIsoResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//加载ISO到云主机
func AttachIsoToVmInstance(vmInstance model.AttachIsoToVmInstanceRequest) models.Result {
	//POST zstack/v1/vm-instances/{vmInstanceUuid}/iso/{isoUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/iso/%s", vmInstance.VmInstanceUuid, vmInstance.IsoUuid)
	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return models.Error(errcode.AttachIsoToVmInstanceFailed, err.Error())
	}

	var vmInstanceResp model.AttachIsoToVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//卸载云主机上的ISO
func DetachIsoFromVmInstance(vmInstance model.DetachIsoFromVmInstanceRequest) models.Result {
	//DELETE zstack/v1/vm-instances/{vmInstanceUuid}/iso?isoUuid={isoUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/iso", vmInstance.VmInstanceUuid)
	dataStr, err := request.DeleteUrlWithParams(url, vmInstance)
	if err != nil {
		return models.Error(errcode.DetachIsoFromVmInstanceFailed, err.Error())
	}

	var vmInstanceResp model.DetachIsoFromVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机可加载云盘列表
func GetVmAttachableDataVolume(vmInstance model.GetVmAttachableDataVolumeRequest) models.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/data-volume-candidates
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/data-volume-candidates", vmInstance.VmInstanceUuid)
	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmAttachableDataVolumeFailed, err.Error())
	}

	var vmInstanceResp model.GetVmAttachableDataVolumeResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机可加载L3网络列表
func GetVmAttachableL3Network(vmInstance model.GetVmAttachableL3NetworkRequest) models.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/l3-networks-candidates
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/l3-networks-candidates", vmInstance.VmInstanceUuid)
	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmAttachableL3NetworkFailed, err.Error())
	}

	var vmInstanceResp model.GetVmAttachableL3NetworkResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//加载L3网络到云主机
func AttachL3NetworkToVm(vmInstance model.AttachL3NetworkToVmRequest) models.Result {
	//POST zstack/v1/vm-instances/{vmInstanceUuid}/l3-networks/{l3NetworkUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/l3-networks/%s", vmInstance.VmInstanceUuid, vmInstance.L3NetworkUuid)
	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return models.Error(errcode.AttachL3NetworkToVmFailed, err.Error())
	}

	var vmInstanceResp model.AttachL3NetworkToVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//从云主机卸载网络
func DetachL3NetworkFromVm(vmInstance model.DetachL3NetworkFromVmRequest) models.Result {
	//DELETE zstack/v1/vm-instances/nics/{vmNicUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/nics/%s", vmInstance.VmNicUuid)
	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return models.Error(errcode.DetachL3NetworkFromVmFailed, err.Error())
	}

	var vmInstanceResp model.DetachL3NetworkFromVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//创建云主机网卡
func CreateVmNic(vmInstance model.CreateVmNicRequest) models.Result {
	//POST zstack/v1/nics
	url := fmt.Sprintf("zstack/v1/nics")
	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return models.Error(errcode.CreateVmNicFailed, err.Error())
	}

	var vmInstanceResp model.CreateVmNicResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//加载网卡到云主机
func AttachVmNicToVm(vmInstance model.AttachVmNicToVmRequest) models.Result {
	//POST zstack/v1/vm-instances/{vmInstanceUuid}/nices/{vmNicUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/nices/%s", vmInstance.VmInstanceUuid, vmInstance.VmNicUuid)
	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return models.Error(errcode.AttachVmNicToVmFailed, err.Error())
	}

	var vmInstanceResp model.AttachVmNicToVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//删除云主机网卡
func DeleteVmNic(vmInstance model.DeleteVmNicRequest) models.Result {
	//DELETE zstack/v1/nics/{uuid}
	url := fmt.Sprintf("zstack/v1/nics/%s", vmInstance.Uuid)
	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return models.Error(errcode.DeleteVmNicFailed, err.Error())
	}

	var vmInstanceResp model.DeleteVmNicResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//查询云主机网卡
func QueryVmNic(vmInstance model.QueryVmNicRequest) models.Result {
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
		return models.Error(errcode.QueryVmNicFailed, err.Error())
	}

	var vmInstanceResp model.QueryVmNicResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取网卡加载的网络服务名称
func GetVmNicAttachedNetworkService(vmInstance model.GetVmNicAttachedNetworkServiceRequest) models.Result {
	//GET zstack/v1/vm-instances/nics/{vmNicUuid}/attached-networkservices
	url := fmt.Sprintf("zstack/v1/vm-instances/nics/%s/attached-networkservices", vmInstance.VmNicUuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmNicAttachedNetworkServiceFailed, err.Error())
	}

	var vmInstanceResp model.GetVmNicAttachedNetworkServiceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//修改云主机网卡三层网络
func ChangeVmNicNetwork(vmInstance model.ChangeVmNicNetworkRequest) models.Result {
	//POST zstack/v1/vm-instances/nics/{vmNicUuid}/l3-networks/{l3NetworkUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/nics/%s/l3-networks/%s", vmInstance.VmNicUuid, vmInstance.DestL3NetworkUuid)

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return models.Error(errcode.ChangeVmNicNetworkFailed, err.Error())
	}

	var vmInstanceResp model.ChangeVmNicNetworkResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机网卡可挂载的三层网络
func GetCandidateL3NetworksForChangeVmNicNetwork(vmInstance model.GetCandidateL3NetworksForChangeVmNicNetworkRequest) models.Result {
	//GET zstack/v1/vm-instances/nics/{vmNicUuid}/l3-networks-candidates
	url := fmt.Sprintf("zstack/v1/vm-instances/nics/%s/l3-networks-candidates", vmInstance.VmNicUuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetCandidateL3NetworksForChangeVmNicNetworkFailed, err.Error())
	}

	var vmInstanceResp model.GetCandidateL3NetworksForChangeVmNicNetworkResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机网卡限速
func SetNicQoS(vmInstance model.SetNicQoSRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetNicQoSFailed, err.Error())
	}

	var vmInstanceResp model.SetNicQoSResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机网卡限速
func GetNicQoS(vmInstance model.GetNicQoSRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/nic-qos
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/nic-qos", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetNicQoSFailed, err.Error())
	}

	var vmInstanceResp model.GetNicQoSResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//取消云主机网卡限速
func DeleteNicQoS(vmInstance model.DeleteNicQoSRequest) models.Result {
	//DELETE zstack/v1/vm-instances/{uuid}/nic-qos
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/nic-qos", vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return models.Error(errcode.DeleteNicQoSFailed, err.Error())
	}

	var vmInstanceResp model.DeleteNicQoSResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取相互依赖的镜像和L3网络
func GetInterdependentL3NetworksImages(vmInstance model.GetInterdependentL3NetworksImagesRequest) models.Result {
	//GET zstack/v1/images-l3networks/dependencies
	url := fmt.Sprintf("zstack/v1/images-l3networks/dependencies")

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetInterdependentL3NetworksImagesFailed, err.Error())
	}

	var vmInstanceResp model.GetInterdependentL3NetworksImagesResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机SSH Key
func SetVmSshKey(vmInstance model.SetVmSshKeyRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmSshKeyFailed, err.Error())
	}

	var vmInstanceResp model.SetVmSshKeyResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机SSH Key
func GetVmSshKey(vmInstance model.GetVmSshKeyRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/ssh-keys
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/ssh-keys", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmSshKeyFailed, err.Error())
	}

	var vmInstanceResp model.GetVmSshKeyResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//删除云主机SSH Key
func DeleteVmSshKey(vmInstance model.DeleteVmSshKeyRequest) models.Result {
	//DELETE zstack/v1/vm-instances/{uuid}/ssh-keys
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/ssh-keys", vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return models.Error(errcode.DeleteVmSshKeyFailed, err.Error())
	}

	var vmInstanceResp model.DeleteVmSshKeyResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//变更云主机密码
func ChangeVmPassword(vmInstance model.ChangeVmPasswordRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.ChangeVmPasswordFailed, err.Error())
	}

	var vmInstanceResp model.ChangeVmPasswordResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机控制台密码
func SetVmConsolePassword(vmInstance model.SetVmConsolePasswordRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmConsolePasswordFailed, err.Error())
	}

	var vmInstanceResp model.SetVmConsolePasswordResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机控制台密码
func GetVmConsolePassword(vmInstance model.GetVmConsolePasswordRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/console-passwords
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/console-passwords", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmConsolePasswordFailed, err.Error())
	}

	var vmInstanceResp model.GetVmConsolePasswordResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//删除云主机控制台密码
func DeleteVmConsolePassword(vmInstance model.DeleteVmConsolePasswordRequest) models.Result {
	//DELETE zstack/v1/vm-instances/{uuid}/console-password
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/console-passwords", vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return models.Error(errcode.DeleteVmConsolePasswordFailed, err.Error())
	}

	var vmInstanceResp model.DeleteVmConsolePasswordResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机控制台地址和访问协议
func GetVmConsoleAddress(vmInstance model.GetVmConsoleAddressRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/console-addresses
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/console-addresses", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmConsoleAddressFailed, err.Error())
	}

	var vmInstanceResp model.GetVmConsoleAddressResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机Hostname
func SetVmHostname(vmInstance model.SetVmHostnameRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmHostnameFailed, err.Error())
	}

	var vmInstanceResp model.SetVmHostnameResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机Hostname
func GetVmHostname(vmInstance model.GetVmHostnameRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/hostnames
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/hostnames", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmHostnameFailed, err.Error())
	}

	var vmInstanceResp model.GetVmHostnameResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//删除云主机Hostname
func DeleteVmHostname(vmInstance model.DeleteVmHostnameRequest) models.Result {
	//DELETE zstack/v1/vm-instances/{uuid}/hostnames
	url := fmt.Sprintf("zstack/v1/vm-instances/%s", vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return models.Error(errcode.DeleteVmHostnameFailed, err.Error())
	}

	var vmInstanceResp model.DeleteVmHostnameResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//创建启动云主机的定时任务
func CreateStartVmInstanceScheduler(vmInstance model.CreateStartVmInstanceSchedulerRequest) models.Result {
	//POST zstack/v1/vm-instances/{vmUuid}/schedulers/starting
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/schedulers/starting", vmInstance.VmUuid)

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return models.Error(errcode.CreateStartVmInstanceSchedulerFailed, err.Error())
	}

	var vmInstanceResp model.CreateStartVmInstanceSchedulerResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//创建停止云主机的定时任务
func CreateStopVmInstanceScheduler(vmInstance model.CreateStopVmInstanceSchedulerRequest) models.Result {
	//POST zstack/v1/vm-instances/{vmUuid}/schedulers/stopping
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/schedulers/stopping", vmInstance.VmUuid)

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return models.Error(errcode.CreateStopVmInstanceSchedulerFailed, err.Error())
	}

	var vmInstanceResp model.CreateStopVmInstanceSchedulerResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//创建重启云主机的定时任务
func CreateRebootVmInstanceScheduler(vmInstance model.CreateRebootVmInstanceSchedulerRequest) models.Result {
	//POST zstack/v1/vm-instances/{vmUuid}/schedulers/rebooting
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/schedulers/rebooting", vmInstance.VmUuid)

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return models.Error(errcode.CreateRebootVmInstanceSchedulerFailed, err.Error())
	}

	var vmInstanceResp model.CreateRebootVmInstanceSchedulerResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机启动设备列表
func GetVmBootOrder(vmInstance model.GetVmBootOrderRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/boot-orders
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/boot-orders", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmBootOrderFailed, err.Error())
	}

	var vmInstanceResp model.GetVmBootOrderResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//指定云主机启动设备
func SetVmBootOrder(vmInstance model.SetVmBootOrderRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmBootOrderFailed, err.Error())
	}

	var vmInstanceResp model.SetVmBootOrderResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取目的地列表
func GetCandidateZonesClustersHostsForCreatingVm(vmInstance model.GetCandidateZonesClustersHostsForCreatingVmRequest) models.Result {
	//GET zstack/v1/vm-instances/candidate-destinations
	url := fmt.Sprintf("zstack/v1/vm-instances/candidate-destinations")

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetCandidateZonesClustersHostsForCreatingVmFailed, err.Error())
	}

	var vmInstanceResp model.GetCandidateZonesClustersHostsForCreatingVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机可启动目的地列表
func GetVmStartingCandidateClustersHosts(vmInstance model.GetVmStartingCandidateClustersHostsRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/starting-target-hosts
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/starting-target-hosts", vmInstance.UUID)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmStartingCandidateClustersHostsFailed, err.Error())
	}

	var vmInstanceResp model.GetVmStartingCandidateClustersHostsResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//指定云主机IP
func SetVmStaticIp(vmInstance model.SetVmStaticIpRequest) models.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.VmInstanceUuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmStaticIpFailed, err.Error())
	}

	var vmInstanceResp model.SetVmStaticIpResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//删除云主机指定IP
func DeleteVmStaticIp(vmInstance model.DeleteVmStaticIpRequest) models.Result {
	//DELETE zstack/v1/vm-instances/{vmInstanceUuid}/static-ips?l3NetworkUuid={l3NetworkUuid}&deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/static-ips", vmInstance.VmInstanceUuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return models.Error(errcode.DeleteVmStaticIpFailed, err.Error())
	}

	var vmInstanceResp model.DeleteVmStaticIpResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机能力
func GetVmCapabilities(vmInstance model.GetVmCapabilitiesRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/capabilities
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/capabilities", vmInstance.UUID)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmCapabilitiesFailed, err.Error())
	}

	var vmInstanceResp model.GetVmCapabilitiesResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//更新云主机信息
func UpdateVmInstance(vmInstance model.UpdateVmInstanceRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.UpdateVmInstanceFailed, err.Error())
	}

	var vmInstanceResp model.UpdateVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//克隆云主机到指定物理机
func CloneVmInstance(vmInstance model.CloneVmInstanceRequest) models.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.VmInstanceUuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.CloneVmInstanceFailed, err.Error())
	}

	var vmInstanceResp model.CloneVmInstanceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机时钟同步
func SetVmClockTrack(vmInstance model.SetVmClockTrackRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmClockTrackFailed, err.Error())
	}

	var vmInstanceResp model.SetVmClockTrackResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机高可用级别
func SetVmInstanceHaLevel(vmInstance model.SetVmInstanceHaLevelRequest) models.Result {
	//POST zstack/v1/vm-instances/{uuid}/ha-levels
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/ha-levels", vmInstance.UUID)

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmInstanceHaLevelFailed, err.Error())
	}

	var vmInstanceResp model.SetVmInstanceHaLevelResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机高可用级别
func GetVmInstanceHaLevel(vmInstance model.GetVmInstanceHaLevelRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/ha-levels
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/ha-levels", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmInstanceHaLevelFailed, err.Error())
	}

	var vmInstanceResp model.GetVmInstanceHaLevelResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//取消云主机高可用
func DeleteVmInstanceHaLevel(vmInstance model.DeleteVmInstanceHaLevelRequest) models.Result {
	//DELETE zstack/v1/instances/{uuid}/ha-levels
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/ha-levels", vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return models.Error(errcode.DeleteVmInstanceHaLevelFailed, err.Error())
	}

	var vmInstanceResp model.DeleteVmInstanceHaLevelResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机Qga
func GetVmQga(vmInstance model.GetVmQgaRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/qga
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/qga", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmQgaFailed, err.Error())
	}

	var vmInstanceResp model.GetVmQgaResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机Qga
func SetVmQga(vmInstance model.SetVmQgaRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmQgaFailed, err.Error())
	}

	var vmInstanceResp model.SetVmQgaResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机RDP开关状态
func GetVmRDP(vmInstance model.GetVmRDPRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/rdp
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/rdp", vmInstance.UUID)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmRDPFailed, err.Error())
	}

	var vmInstanceResp model.GetVmRDPResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机RDP开关状态
func SetVmRDP(vmInstance model.SetVmRDPRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmRDPFailed, err.Error())
	}

	var vmInstanceResp model.SetVmRDPResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机支持的屏幕数
func GetVmMonitorNumber(vmInstance model.GetVmMonitorNumberRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/monitorNumber
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/monitorNumber", vmInstance.UUID)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmMonitorNumberFailed, err.Error())
	}

	var vmInstanceResp model.GetVmMonitorNumberResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机支持的屏幕数
func SetVmMonitorNumber(vmInstance model.SetVmMonitorNumberRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmMonitorNumberFailed, err.Error())
	}

	var vmInstanceResp model.SetVmMonitorNumberResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//修改云主机根云盘
func ChangeVmImage(vmInstance model.ChangeVmImageRequest) models.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.VmInstanceUuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.ChangeVmImageFailed, err.Error())
	}

	var vmInstanceResp model.ChangeVmImageResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取候选镜像列表
func GetImageCandidatesForVmToChange(vmInstance model.GetImageCandidatesForVmToChangeRequest) models.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/image-candidates
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/image-candidates", vmInstance.VmInstanceUuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetImageCandidatesForVmToChangeFailed, err.Error())
	}

	var vmInstanceResp model.GetImageCandidatesForVmToChangeResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//更新云主机mac地址
func UpdateVmNicMac(vmInstance model.UpdateVmNicMacRequest) models.Result {
	//PUT zstack/v1/vm-instances/nics/{vmNicUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/nics/%s/actions", vmInstance.VmNicUuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.UpdateVmNicMacFailed, err.Error())
	}

	var vmInstanceResp model.UpdateVmNicMacResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机防IP欺骗启用状态
func SetVmCleanTraffic(vmInstance model.SetVmCleanTrafficRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmCleanTrafficFailed, err.Error())
	}

	var vmInstanceResp model.SetVmCleanTrafficResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//为云主机创建CDROM
func CreateVmCdRom(vmInstance model.CreateVmCdRomRequest) models.Result {
	//POST zstack/v1/vm-instances/cdroms
	url := fmt.Sprintf("zstack/v1/vm-instances/cdroms")

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return models.Error(errcode.CreateVmCdRomFailed, err.Error())
	}

	var vmInstanceResp model.CreateVmCdRomResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//删除CDROM
func DeleteVmCdRom(vmInstance model.DeleteVmCdRomRequest) models.Result {
	//DELETE zstack/v1/vm-instances/cdroms/{uuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/cdroms/%s", vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return models.Error(errcode.DeleteVmCdRomFailed, err.Error())
	}

	var vmInstanceResp model.DeleteVmCdRomResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//修改CDROM
func UpdateVmCdRom(params model.UpdateVmCdRomRequest) models.Result {
	//PUT zstack/v1/vm-instances/cdroms/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/cdroms/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateVmCdRomFailed, err.Error())
	}
	var result model.UpdateVmCdRomResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//设置云主机默认CDROM
func SetVmInstanceDefaultCdRom(vmInstance model.SetVmInstanceDefaultCdRomRequest) models.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/cdroms/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/cdroms/%s/actions", vmInstance.VmInstanceUuid, vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmInstanceDefaultCdRomFailed, err.Error())
	}

	var vmInstanceResp model.SetVmInstanceDefaultCdRomResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//查询CDROM清单
func QueryVmCdRom(vmInstance model.QueryVmCdRomRequest) models.Result {
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
		return models.Error(errcode.QueryVmCdRomFailed, err.Error())
	}

	var vmInstanceResp model.QueryVmCdRomResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//更改云主机优先级级别
func UpdateVmPriority(vmInstance model.UpdateVmPriorityRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.UpdateVmPriorityFailed, err.Error())
	}

	var vmInstanceResp model.UpdateVmPriorityResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机显存
func SetVmQxlMemory(vmInstance model.SetVmQxlMemoryRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmQxlMemoryFailed, err.Error())
	}

	var vmInstanceResp model.SetVmQxlMemoryResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机虚拟声卡类型
func SetVmSoundType(vmInstance model.SetVmSoundTypeRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmSoundTypeFailed, err.Error())
	}

	var vmInstanceResp model.SetVmSoundTypeResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取spice的CA证书
func GetSpiceCertificates(vmInstance model.GetSpiceCertificatesRequest) models.Result {
	//GET zstack/v1/spice/certificates
	url := fmt.Sprintf("zstack/v1/spice/certificates")

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetSpiceCertificatesFailed, err.Error())
	}

	var vmInstanceResp model.GetSpiceCertificatesResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//为云主机加载增强工具镜像
func AttachGuestToolsIsoToVm(vmInstance model.AttachGuestToolsIsoToVmRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.AttachGuestToolsIsoToVmFailed, err.Error())
	}

	var vmInstanceResp model.AttachGuestToolsIsoToVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机可用的最新增强工具
func GetLatestGuestToolsForVm(vmInstance model.GetLatestGuestToolsForVmRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/latest-guest-tools
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/latest-guest-tools", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetLatestGuestToolsForVmFailed, err.Error())
	}

	var vmInstanceResp model.GetLatestGuestToolsForVmResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机内部增强工具的信息
func GetVmGuestToolsInfo(vmInstance model.GetVmGuestToolsInfoRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/guest-tools-infos
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/guest-tools-infos", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmGuestToolsInfoFailed, err.Error())
	}

	var vmInstanceResp model.GetVmGuestToolsInfoResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机第一启动项
func GetVmInstanceFirstBootDevice(vmInstance model.GetVmInstanceFirstBootDeviceRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/first-boot-device
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/first-boot-device", vmInstance.Uuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmInstanceFirstBootDeviceFailed, err.Error())
	}

	var vmInstanceResp model.GetVmInstanceFirstBootDeviceResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置网卡型号
func UpdateVmNicDriver(vmInstance model.UpdateVmNicDriverRequest) models.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions?vmNicUuid={vmNicUuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.VmInstanceUuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.UpdateVmNicDriverFailed, err.Error())
	}

	var vmInstanceResp model.UpdateVmNicDriverResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取云主机设备地址
func GetVmDeviceAddress(vmInstance model.GetVmDeviceAddressRequest) models.Result {
	//GET zstack/v1/vm-instances/devices
	url := fmt.Sprintf("zstack/v1/vm-instances/devices")

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmDeviceAddressFailed, err.Error())
	}

	var vmInstanceResp model.GetVmDeviceAddressResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//批量获取云主机能力
func GetVmsCapabilities(vmInstance model.GetVmsCapabilitiesRequest) models.Result {
	//POST zstack/v1/vm-instances/capabilities
	url := fmt.Sprintf("zstack/v1/vm-instances/capabilities")

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmsCapabilitiesFailed, err.Error())
	}

	var vmInstanceResp model.GetVmsCapabilitiesResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机vNUMA
func SetVmNuma(vmInstance model.SetVmNumaRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmNumaFailed, err.Error())
	}

	var vmInstanceResp model.SetVmNumaResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取VM Numa开关状态
func GetVmNuma(vmInstance model.GetVmNumaRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/vnuma
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/vnuma", vmInstance.UUID)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmNumaFailed, err.Error())
	}

	var vmInstanceResp model.GetVmNumaResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//查询云主机的vNUMA拓扑信息
func GetVmvNUMATopology(vmInstance model.GetVmvNUMATopologyRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/vnuma-topology
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/vnuma-topology", vmInstance.UUID)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmvNUMATopologyFailed, err.Error())
	}

	var vmInstanceResp model.GetVmvNUMATopologyResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//设置云主机Emulator Pinning功能
func SetVmEmulatorPinning(vmInstance model.SetVmEmulatorPinningRequest) models.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.SetVmEmulatorPinningFailed, err.Error())
	}

	var vmInstanceResp model.SetVmEmulatorPinningResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}

//获取VM Emulator Pin在那些物理机的cpu上
func GetVmEmulatorPinning(vmInstance model.GetVmEmulatorPinningRequest) models.Result {
	//GET zstack/v1/vm-instances/{uuid}/emulator-pinning
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/emulator-pinning", vmInstance.UUID)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return models.Error(errcode.GetVmEmulatorPinningFailed, err.Error())
	}

	var vmInstanceResp model.GetVmEmulatorPinningResponse
	utils.FromJSON(dataStr, &vmInstanceResp)
	logs.Debug("最终结果:{}", vmInstanceResp)

	return models.Success(vmInstanceResp)
}
