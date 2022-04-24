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

//查询SCSI Lun
func QueryScsiLun(vmInstance model.QueryScsiLunRequest) mgresult.Result {
	//GET zstack/v1/storage-devices/scsi-lun/luns
	//GET zstack/v1/storage-devices/scsi-lun/luns/{uuid}
	var url string
	if vmInstance.UUID == "" {
		url = fmt.Sprintf("zstack/v1/storage-devices/scsi-lun/luns")
	} else {
		url = fmt.Sprintf("zstack/v1/storage-devices/scsi-lun/luns/%s", vmInstance.UUID)
	}

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryScsiLunFailed, err)
	}
	var result model.QueryScsiLunResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将LUN设备从物理机卸载
func DetachScsiLunFromHost(vmInstance model.DetachScsiLunFromHostRequest) mgresult.Result {
	//PUT zstack/v1/storage-devices/scsi-lun/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/storage-devices/scsi-lun/%s/actions")

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachScsiLunFromHostFailed, err)
	}
	var result model.DetachScsiLunFromHostResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加iSCSI服务器
func AddIscsiServer(vmInstance model.AddIscsiServerRequest) mgresult.Result {
	//POST zstack/v1/storage-devices/iscsi/servers
	url := fmt.Sprintf("zstack/v1/storage-devices/iscsi/servers")

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddIscsiServerFailed, err)
	}
	var result model.AddIscsiServerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除iSCSI服务器
func DeleteIscsiServer(vmInstance model.DeleteIscsiServerRequest) mgresult.Result {
	//DELETE zstack/v1/storage-devices/iscsi/servers/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/storage-devices/iscsi/servers/%s", vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteIscsiServerFailed, err)
	}
	var result model.DeleteIscsiServerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询iSCSI服务器
func QueryIscsiServer(vmInstance model.QueryIscsiServerRequest) mgresult.Result {
	//GET zstack/v1/storage-devices/iscsi/servers
	//GET zstack/v1/storage-devices/iscsi/servers/{uuid}
	var url string
	if vmInstance.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/storage-devices/iscsi/servers")
	} else {
		url = fmt.Sprintf("zstack/v1/storage-devices/iscsi/servers/%s", vmInstance.Uuid)
	}

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryIscsiServerFailed, err)
	}
	var result model.QueryIscsiServerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//刷新iSCSI服务器
func RefreshIscsiServer(vmInstance model.RefreshIscsiServerRequest) mgresult.Result {
	//POST zstack/v1/storage-devices/iscsi/servers/{uuid}
	url := fmt.Sprintf("zstack/v1/storage-devices/iscsi/servers/%s", vmInstance.Uuid)

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RefreshIscsiServerFailed, err)
	}
	var result model.RefreshIscsiServerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新iSCSI服务器配置
func UpdateIscsiServer(vmInstance model.UpdateIscsiServerRequest) mgresult.Result {
	//POST zstack/v1/storage-devices/iscsi/servers
	url := fmt.Sprintf("zstack/v1/storage-devices/iscsi/servers")

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateIscsiServerFailed, err)
	}
	var result model.UpdateIscsiServerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将iSCSI服务器加载到集群
func AttachIscsiServerToCluster(vmInstance model.AttachIscsiServerToClusterRequest) mgresult.Result {
	//POST zstack/v1/clusters/{clusterUuid}/storage-devices/iscsi/servers/{uuid}
	url := fmt.Sprintf("zstack/v1/clusters/%s/storage-devices/iscsi/servers/%s", vmInstance.ClusterUuid, vmInstance.Uuid)

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachIscsiServerToClusterFailed, err)
	}
	var result model.AttachIscsiServerToClusterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将iSCSI服务器从集群卸载
func DetachIscsiServerFromCluster(vmInstance model.DetachIscsiServerFromClusterRequest) mgresult.Result {
	//DELETE zstack/v1/clusters/{clusterUuid}/storage-devices/iscsi/servers/{uuid}
	url := fmt.Sprintf("zstack/v1/clusters/%s/storage-devices/iscsi/servers/%s", vmInstance.ClusterUuid, vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachIscsiServerFromClusterFailed, err)
	}
	var result model.DetachIscsiServerFromClusterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询iSCSI磁盘
func QueryIscsiLun(vmInstance model.QueryIscsiLunRequest) mgresult.Result {
	//GET zstack/v1/storage-devices/iscsi/luns
	//GET zstack/v1/storage-devices/iscsi/luns/{uuid}
	var url string
	if vmInstance.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/storage-devices/iscsi/luns")
	} else {
		url = fmt.Sprintf("zstack/v1/storage-devices/iscsi/luns/%s", vmInstance.Uuid)
	}

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryIscsiLunFailed, err)
	}
	var result model.QueryIscsiLunResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询FC SAN存储
func QueryFiberChannelStorage(vmInstance model.QueryFiberChannelStorageRequest) mgresult.Result {
	//GET zstack/v1/storage-devices/fiber-channel/controllers
	//GET zstack/v1/storage-devices/fiber-channel/controllers/{uuid}
	var url string
	if vmInstance.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/storage-devices/fiber-channel/controllers")
	} else {
		url = fmt.Sprintf("zstack/v1/storage-devices/fiber-channel/controllers/%s", vmInstance.Uuid)
	}

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryFiberChannelStorageFailed, err)
	}
	var result model.QueryFiberChannelStorageRequest
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//刷新FC SAN存储
func RefreshFiberChannelStorage(vmInstance model.RefreshFiberChannelStorageRequest) mgresult.Result {
	//POST zstack/v1/storage-devices/fiber-channel/controllers
	url := fmt.Sprintf("zstack/v1/storage-devices/fiber-channel/controllers")

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RefreshFiberChannelStorageFailed, err)
	}
	var result model.RefreshFiberChannelStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新SCSI LUN配置
func UpdateScsiLun(vmInstance model.UpdateScsiLunRequest) mgresult.Result {
	//PUT zstack/v1/storage-devices/scsi-lun/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/storage-devices/scsi-lun/%s/actions", vmInstance.Uuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateScsiLunFailed, err)
	}
	var result model.UpdateScsiLunResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将SCSI LUN加载到云主机
func AttachScsiLunToVmInstance(vmInstance model.AttachScsiLunToVmInstanceRequest) mgresult.Result {
	//POST zstack/v1/vm-instances/{vmInstanceUuid}/scsi-lun/{uuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/scsi-lun/%s", vmInstance.VmInstanceUuid, vmInstance.Uuid)

	dataStr, err := request.Post(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachScsiLunToVmInstanceFailed, err)
	}
	var result model.AttachScsiLunToVmInstanceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将SCSI LUN从云主机卸载
func DetachScsiLunFromVmInstance(vmInstance model.DetachScsiLunFromVmInstanceRequest) mgresult.Result {
	//DELETE zstack/v1/vm-instances/{vmInstanceUuid}/scsi-lun/{uuid}
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/scsi-lun/%s", vmInstance.VmInstanceUuid, vmInstance.Uuid)

	dataStr, err := request.Delete(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachScsiLunFromVmInstanceFailed, err)
	}
	var result model.DetachScsiLunFromVmInstanceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//检查SCSI Lun与集群连接关系
func CheckScsiLunClusterStatus(vmInstance model.CheckScsiLunClusterStatusRequest) mgresult.Result {
	//PUT zstack/v1/storage-devices/scsi-lun/{uuid}/cluster/{clusterUuid}
	url := fmt.Sprintf("zstack/v1/storage-devices/scsi-lun/%s/cluster/%s", vmInstance.Uuid, vmInstance.CusterUuid)

	dataStr, err := request.Put(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CheckScsiLunClusterStatusFailed, err)
	}
	var result model.CheckScsiLunClusterStatusResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取可加载的SCSI Lun
func GetScsiLunCandidatesForAttachingVm(vmInstance model.GetScsiLunCandidatesForAttachingVmRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/candidate-storage-devices
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/candidate-storage-devices", vmInstance.VmInstanceUuid)

	dataStr, err := request.Get(url, vmInstance)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetScsiLunCandidatesForAttachingVmFailed, err)
	}
	var result model.GetScsiLunCandidatesForAttachingVmRequest
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
