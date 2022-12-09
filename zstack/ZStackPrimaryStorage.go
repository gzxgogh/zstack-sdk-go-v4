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

//删除主存储
func DeletePrimaryStorage(params model.DeletePrimaryStorageRequest) models.Result {
	//DELETE zstack/v1/primary-storage/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/primary-storage/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeletePrimaryStorageFailed, err.Error())
	}
	var result model.DeletePrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询物理机
func QueryPrimaryStorage(params model.QueryPrimaryStorageRequest) models.Result {
	//GET zstack/v1/primary-storage
	//GET zstack/v1/primary-storage/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/primary-storage")
	} else {
		url = fmt.Sprintf("zstack/v1/primary-storage/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPrimaryStorageFailed, err.Error())
	}
	var result model.QueryPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//向集群添加主存储
func AttachPrimaryStorageToCluster(params model.AttachPrimaryStorageToClusterRequest) models.Result {
	//POST zstack/v1/clusters/{clusterUuid}/primary-storage/{primaryStorageUuid}
	url := fmt.Sprintf("zstack/v1/clusters/%s/primary-storage/%s", params.ClusterUuid, params.PrimaryStorageUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachPrimaryStorageToClusterFailed, err.Error())
	}
	var result model.AttachPrimaryStorageToClusterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从集群卸载主存储
func DetachPrimaryStorageFromCluster(params model.DetachPrimaryStorageFromClusterRequest) models.Result {
	//DELETE	zstack/v1/clusters/{clusterUuid}/primary-storage/{primaryStorageUuid}
	url := fmt.Sprintf("zstack/v1/clusters/%s/primary-storage/%s", params.ClusterUuid, params.PrimaryStorageUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachPrimaryStorageFromClusterFailed, err.Error())
	}
	var result model.DetachPrimaryStorageFromClusterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//重连主存储
func ReconnectPrimaryStorage(params model.ReconnectPrimaryStorageRequest) models.Result {
	//PUT zstack/v1/primary-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ReconnectPrimaryStorageFailed, err.Error())
	}
	var result model.ReconnectPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取主存储容量
func GetPrimaryStorageCapacity(params model.GetPrimaryStorageCapacityRequest) models.Result {
	//GET zstack/v1/primary-storage/capacities
	url := fmt.Sprintf("zstack/v1/primary-storage/capacities")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetPrimaryStorageCapacityFailed, err.Error())
	}
	var result model.GetPrimaryStorageCapacityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//刷新主存储容量
func SyncPrimaryStorageCapacity(params model.SyncPrimaryStorageCapacityRequest) models.Result {
	//PUT zstack/v1/primary-storage/{primaryStorageUuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/%s/actions", params.PrimaryStorageUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.SyncPrimaryStorageCapacityFailed, err.Error())
	}
	var result model.SyncPrimaryStorageCapacityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更改主存储状态
func ChangePrimaryStorageState(params model.ChangePrimaryStorageStateRequest) models.Result {
	//PUT zstack/v1/primary-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangePrimaryStorageStateFailed, err.Error())
	}
	var result model.ChangePrimaryStorageStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新主存储信息
func UpdatePrimaryStorage(params model.UpdatePrimaryStorageRequest) models.Result {
	//PUT zstack/v1/primary-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdatePrimaryStorageFailed, err.Error())
	}
	var result model.UpdatePrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//清除主存储镜像缓存
func CleanUpImageCacheOnPrimaryStorage(params model.CleanUpImageCacheOnPrimaryStorageRequest) models.Result {
	//PUT zstack/v1/primary-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.CleanUpImageCacheOnPrimaryStorageFailed, err.Error())
	}
	var result model.CleanUpImageCacheOnPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取主存储分配策略清单
func GetPrimaryStorageAllocatorStrategies(params model.GetPrimaryStorageAllocatorStrategiesRequest) models.Result {
	//GET zstack/v1/primary-storage/allocators/strategies
	url := fmt.Sprintf("zstack/v1/primary-storage/allocators/strategies")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetPrimaryStorageAllocatorStrategiesFailed, err.Error())
	}
	var result model.GetPrimaryStorageAllocatorStrategiesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取主存储类型列表
func GetPrimaryStorageTypes(params model.GetPrimaryStorageTypesRequest) models.Result {
	//GET zstack/v1/primary-storage/types
	url := fmt.Sprintf("zstack/v1/primary-storage/types")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetPrimaryStorageTypesFailed, err.Error())
	}
	var result model.GetPrimaryStorageTypesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取候选列表
func GetPrimaryStorageCandidatesForVolumeMigration(params model.GetPrimaryStorageCandidatesForVolumeMigrationRequest) models.Result {
	//GET zstack/v1/primary-storage/volumes/{volumeUuid}/migration-candidates
	url := fmt.Sprintf("zstack/v1/primary-storage/volumes/%s/migration-candidates", params.VolumeUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetPrimaryStorageCandidatesForVolumeMigrationFailed, err.Error())
	}
	var result model.GetPrimaryStorageCandidatesForVolumeMigrationResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取跨存储迁移可选物理机列表
func GetHostCandidatesForVmMigration(params model.GetHostCandidatesForVmMigrationRequest) models.Result {
	//GET zstack/v1/primary-storage/hosts/{vmInstanceUuid}/migration-candidates
	url := fmt.Sprintf("zstack/v1/primary-storage/hosts/%s/migration-candidates", params.VmInstanceUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetHostCandidatesForVmMigrationFailed, err.Error())
	}
	var result model.GetHostCandidatesForVmMigrationResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//跨主存储迁移云盘
func PrimaryStorageMigrateVolume(params model.PrimaryStorageMigrateVolumeRequest) models.Result {
	//PUT zstack/v1/primary-storage/volumes/{volumeUuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/volumes/%s/actions", params.VolumeUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.PrimaryStorageMigrateVolumeFailed, err.Error())
	}
	var result model.PrimaryStorageMigrateVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//跨主存储迁移云主机
func PrimaryStorageMigrateVm(params model.PrimaryStorageMigrateVmRequest) models.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", params.VmInstanceUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.PrimaryStorageMigrateVmFailed, err.Error())
	}
	var result model.PrimaryStorageMigrateVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取存储迁移候选列表
func GetPrimaryStorageCandidatesForVmMigration(params model.GetPrimaryStorageCandidatesForVmMigrationRequest) models.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/storage-migration-candidates
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/storage-migration-candidates", params.VmInstanceUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetPrimaryStorageCandidatesForVmMigrationFailed, err.Error())
	}
	var result model.GetPrimaryStorageCandidatesForVmMigrationResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取主存储的License信息
func GetPrimaryStorageLicenseInfo(params model.GetPrimaryStorageLicenseInfoRequest) models.Result {
	//GET zstack/v1/primary-storage/{uuid}/license
	url := fmt.Sprintf("zstack/v1/primary-storage/%s/license", params.Uuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetPrimaryStorageLicenseInfoFailed, err.Error())
	}
	var result model.GetPrimaryStorageLicenseInfoResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加本地存储为主存储
func AddLocalPrimaryStorage(params model.AddLocalPrimaryStorageRequest) models.Result {
	//POST zstack/v1/primary-storage/local-storage
	url := fmt.Sprintf("zstack/v1/primary-storage/local-storage")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddLocalPrimaryStorageFailed, err.Error())
	}
	var result model.AddLocalPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询本地存储资源引用
func QueryLocalStorageResourceRef(params model.QueryLocalStorageResourceRefRequest) models.Result {
	//GET zstack/v1/primary-storage/local-storage/resource-refs
	url := fmt.Sprintf("zstack/v1/primary-storage/local-storage/resource-refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryLocalStorageResourceRefFailed, err.Error())
	}
	var result model.QueryLocalStorageResourceRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//迁移本地存储上存放的云盘
func LocalStorageMigrateVolume(params model.LocalStorageMigrateVolumeRequest) models.Result {
	//PUT zstack/v1/primary-storage/local-storage/volumes/{volumeUuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/local-storage/volumes/%s/actions", params.VolumeUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.LocalStorageMigrateVolumeFailed, err.Error())
	}
	var result model.LocalStorageMigrateVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取主机本地存储容量
func GetLocalStorageHostDiskCapacity(params model.GetLocalStorageHostDiskCapacityRequest) models.Result {
	//GET zstack/v1/primary-storage/local-storage/{primaryStorageUuid}/capacities
	url := fmt.Sprintf("zstack/v1/primary-storage/local-storage/%s/capacities", params.PrimaryStorageUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetLocalStorageHostDiskCapacityFailed, err.Error())
	}
	var result model.GetLocalStorageHostDiskCapacityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取迁移本地存储物理机
func LocalStorageGetVolumeMigratableHosts(params model.LocalStorageGetVolumeMigratableHostsRequest) models.Result {
	//GET zstack/v1/volumes/{volumeUuid}/migration-target-hosts
	url := fmt.Sprintf("zstack/v1/volumes/%s/migration-target-hosts", params.VolumeUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.LocalStorageGetVolumeMigratableHostsFailed, err.Error())
	}
	var result model.LocalStorageGetVolumeMigratableHostsResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加NFS主存储
func AddNfsPrimaryStorage(params model.AddNfsPrimaryStorageRequest) models.Result {
	//POST zstack/v1/primary-storage/nfs
	url := fmt.Sprintf("zstack/v1/primary-storage/nfs")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddNfsPrimaryStorageFailed, err.Error())
	}
	var result model.AddNfsPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加一个共享挂载点的主存储
func AddSharedMountPointPrimaryStorage(params model.AddSharedMountPointPrimaryStorageRequest) models.Result {
	//POST zstack/v1/primary-storage/smp
	url := fmt.Sprintf("zstack/v1/primary-storage/smp")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddSharedMountPointPrimaryStorageFailed, err.Error())
	}
	var result model.AddSharedMountPointPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加Ceph主存储
func AddCephPrimaryStorage(params model.AddCephPrimaryStorageRequest) models.Result {
	//POST zstack/v1/primary-storage/ceph
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddCephPrimaryStorageFailed, err.Error())
	}
	var result model.AddCephPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询Ceph主存储
func QueryCephPrimaryStorage(params model.QueryCephPrimaryStorageRequest) models.Result {
	//GET zstack/v1/primary-storage/ceph
	//GET zstack/v1/primary-storage/ceph/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/primary-storage/ceph")
	} else {
		url = fmt.Sprintf("zstack/v1/primary-storage/ceph/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryCephPrimaryStorageFailed, err.Error())
	}
	var result model.QueryCephPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//为Ceph主存储添加mon节点
func AddMonToCephPrimaryStorage(params model.AddMonToCephPrimaryStorageRequest) models.Result {
	//POST zstack/v1/primary-storage/ceph/{uuid}/mons
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph/%s/mons", params.Uuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddMonToCephPrimaryStorageFailed, err.Error())
	}
	var result model.AddMonToCephPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从Ceph主存储删除mon节点
func RemoveMonFromCephPrimaryStorage(params model.RemoveMonFromCephPrimaryStorageRequest) models.Result {
	//DELETE zstack/v1/primary-storage/ceph/{uuid}/mons?monHostnames={monHostnames}
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph/%s/mons", params.Uuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveMonFromCephPrimaryStorageFailed, err.Error())
	}
	var result model.RemoveMonFromCephPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新Ceph主存储mon节点
func UpdateCephPrimaryStorageMon(params model.UpdateCephPrimaryStorageMonRequest) models.Result {
	//PUT zstack/v1/primary-storage/ceph/mons/{monUuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph/mons/%s/actions", params.MonUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateCephPrimaryStorageMonFailed, err.Error())
	}
	var result model.UpdateCephPrimaryStorageMonResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加Ceph主存储池
func AddCephPrimaryStoragePool(params model.AddCephPrimaryStoragePoolRequest) models.Result {
	//POST zstack/v1/primary-storage/ceph/{primaryStorageUuid}/pools
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph/%s/pools", params.PrimaryStorageUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddCephPrimaryStoragePoolFailed, err.Error())
	}
	var result model.AddCephPrimaryStoragePoolResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除Ceph主存储池
func DeleteCephPrimaryStoragePool(params model.DeleteCephPrimaryStoragePoolRequest) models.Result {
	//DELETE zstack/v1/primary-storage/ceph/pools/{uuid}
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph/pools/%s", params.Uuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteCephPrimaryStoragePoolFailed, err.Error())
	}
	var result model.DeleteCephPrimaryStoragePoolResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询Ceph主存储池
func QueryCephPrimaryStoragePool(params model.QueryCephPrimaryStoragePoolRequest) models.Result {
	//GET zstack/v1/primary-storage/ceph/pools
	//GET zstack/v1/primary-storage/ceph/pools/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/primary-storage/ceph/pools")
	} else {
		url = fmt.Sprintf("zstack/v1/primary-storage/ceph/pools/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryCephPrimaryStoragePoolFailed, err.Error())
	}
	var result model.QueryCephPrimaryStoragePoolResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新Ceph主存储池
func UpdateCephPrimaryStoragePool(params model.UpdateCephPrimaryStoragePoolRequest) models.Result {
	//PUT zstack/v1/primary-storage/ceph/pools/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph/pools/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateCephPrimaryStoragePoolFailed, err.Error())
	}
	var result model.UpdateCephPrimaryStoragePoolResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加Shared Block主存储
func AddSharedBlockGroupPrimaryStorage(params model.AddSharedBlockGroupPrimaryStorageRequest) models.Result {
	//POST zstack/v1/primary-storage/sharedblockgroup
	url := fmt.Sprintf("zstack/v1/primary-storage/sharedblockgroup")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddSharedBlockGroupPrimaryStorageFailed, err.Error())
	}
	var result model.AddSharedBlockGroupPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询连接状态
func QuerySharedBlockGroupPrimaryStorageHostRef(params model.QuerySharedBlockGroupPrimaryStorageHostRefRequest) models.Result {
	//GET zstack/v1/sharedblock-group/host-refs
	//GET zstack/v1/sharedblock-group/host-refs/{primaryStorageUuid}
	var url string
	if params.PrimaryStorageUuid == "" {
		url = fmt.Sprintf("zstack/v1/sharedblock-group/host-refs")
	} else {
		url = fmt.Sprintf("zstack/v1/sharedblock-group/host-refs/%s", params.PrimaryStorageUuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QuerySharedBlockGroupPrimaryStorageHostRefFailed, err.Error())
	}
	var result model.QuerySharedBlockGroupPrimaryStorageHostRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询Shared Block主存储
func QuerySharedBlockGroupPrimaryStorage(params model.QuerySharedBlockGroupPrimaryStorageRequest) models.Result {
	//GET zstack/v1/primary-storage/sharedblockgroup
	//GET zstack/v1/primary-storage/sharedblockgroup/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/primary-storage/sharedblockgroup")
	} else {
		url = fmt.Sprintf("zstack/v1/primary-storage/sharedblockgroup/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QuerySharedBlockGroupPrimaryStorageFailed, err.Error())
	}
	var result model.QuerySharedBlockGroupPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加LUN到Shared Block主存储
func AddSharedBlockToSharedBlockGroup(params model.AddSharedBlockToSharedBlockGroupRequest) models.Result {
	//POST zstack/v1/primary-storage/sharedblockgroup/{uuid}/sharedblocks
	url := fmt.Sprintf("zstack/v1/primary-storage/sharedblockgroup/%s/sharedblocks", params.Uuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddSharedBlockToSharedBlockGroupFailed, err.Error())
	}
	var result model.AddSharedBlockToSharedBlockGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取共享块设备候选清单
func GetSharedBlockCandidate(params model.GetSharedBlockCandidateRequest) models.Result {
	//GET zstack/v1/primary-storage/sharedblockgroup/sharedblock-candidates
	url := fmt.Sprintf("zstack/v1/primary-storage/sharedblockgroup/sharedblock-candidates")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetSharedBlockCandidateFailed, err.Error())
	}
	var result model.GetSharedBlockCandidateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//刷新共享块设备容量
func RefreshSharedblockDeviceCapacity(params model.RefreshSharedblockDeviceCapacityRequest) models.Result {
	//POST zstack/v1/primary-storage/sharedblockgroup/{sharedBlockGroupUuid}/sharedblocks/{uuid}
	url := fmt.Sprintf("zstack/v1/primary-storage/sharedblockgroup/%s/sharedblocks/%s", params.SharedBlockGroupUuid, params.Uuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.RefreshSharedblockDeviceCapacityFailed, err.Error())
	}
	var result model.RefreshSharedblockDeviceCapacityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新共享块存储中共享块的diskUuid
func UpdateSharedBlockDiskUuid(params model.UpdateSharedBlockDiskUuidRequest) models.Result {
	//PUT zstack/v1/primary-storage/sharedblockgroup/{sharedBlockGroupUuid}/sharedblocks/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/sharedblockgroup/%s/sharedblocks/%s/actions", params.SharedBlockGroupUuid, params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateSharedBlockDiskUuidFailed, err.Error())
	}
	var result model.UpdateSharedBlockDiskUuidResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询共享块设备
func QuerySharedBlock(params model.QuerySharedBlockRequest) models.Result {
	//GET zstack/v1/sharedblock-group/sharedblocks
	//GET zstack/v1/sharedblock-group/sharedblock/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/sharedblock-group/sharedblocks")
	} else {
		url = fmt.Sprintf("zstack/v1/sharedblock-group/sharedblocks/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QuerySharedBlockFailed, err.Error())
	}
	var result model.QuerySharedBlockResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
