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

//删除主存储
func DeletePrimaryStorage(params model.DeletePrimaryStorageRequest) mgresult.Result {
	//DELETE zstack/v1/primary-storage/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/primary-storage/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeletePrimaryStorageFailed, err)
	}
	var result model.DeletePrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询物理机
func QueryPrimaryStorage(params model.QueryPrimaryStorageRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryPrimaryStorageFailed, err)
	}
	var result model.QueryPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//向集群添加主存储
func AttachPrimaryStorageToCluster(params model.AttachPrimaryStorageToClusterRequest) mgresult.Result {
	//POST zstack/v1/clusters/{clusterUuid}/primary-storage/{primaryStorageUuid}
	url := fmt.Sprintf("zstack/v1/clusters/%s/primary-storage/%s", params.ClusterUuid, params.PrimaryStorageUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachPrimaryStorageToClusterFailed, err)
	}
	var result model.AttachPrimaryStorageToClusterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从集群卸载主存储
func DetachPrimaryStorageFromCluster(params model.DetachPrimaryStorageFromClusterRequest) mgresult.Result {
	//DELETE	zstack/v1/clusters/{clusterUuid}/primary-storage/{primaryStorageUuid}
	url := fmt.Sprintf("zstack/v1/clusters/%s/primary-storage/%s", params.ClusterUuid, params.PrimaryStorageUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachPrimaryStorageFromClusterFailed, err)
	}
	var result model.DetachPrimaryStorageFromClusterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//重连主存储
func ReconnectPrimaryStorage(params model.ReconnectPrimaryStorageRequest) mgresult.Result {
	//PUT zstack/v1/primary-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ReconnectPrimaryStorageFailed, err)
	}
	var result model.ReconnectPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取主存储容量
func GetPrimaryStorageCapacity(params model.GetPrimaryStorageCapacityRequest) mgresult.Result {
	//GET zstack/v1/primary-storage/capacities
	url := fmt.Sprintf("zstack/v1/primary-storage/capacities")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetPrimaryStorageCapacityFailed, err)
	}
	var result model.GetPrimaryStorageCapacityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//刷新主存储容量
func SyncPrimaryStorageCapacity(params model.SyncPrimaryStorageCapacityRequest) mgresult.Result {
	//PUT zstack/v1/primary-storage/{primaryStorageUuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/%s/actions", params.PrimaryStorageUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SyncPrimaryStorageCapacityFailed, err)
	}
	var result model.SyncPrimaryStorageCapacityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更改主存储状态
func ChangePrimaryStorageState(params model.ChangePrimaryStorageStateRequest) mgresult.Result {
	//PUT zstack/v1/primary-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangePrimaryStorageStateFailed, err)
	}
	var result model.ChangePrimaryStorageStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新主存储信息
func UpdatePrimaryStorage(params model.UpdatePrimaryStorageRequest) mgresult.Result {
	//PUT zstack/v1/primary-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdatePrimaryStorageFailed, err)
	}
	var result model.UpdatePrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//清除主存储镜像缓存
func CleanUpImageCacheOnPrimaryStorage(params model.CleanUpImageCacheOnPrimaryStorageRequest) mgresult.Result {
	//PUT zstack/v1/primary-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CleanUpImageCacheOnPrimaryStorageFailed, err)
	}
	var result model.CleanUpImageCacheOnPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取主存储分配策略清单
func GetPrimaryStorageAllocatorStrategies(params model.GetPrimaryStorageAllocatorStrategiesRequest) mgresult.Result {
	//GET zstack/v1/primary-storage/allocators/strategies
	url := fmt.Sprintf("zstack/v1/primary-storage/allocators/strategies")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetPrimaryStorageAllocatorStrategiesFailed, err)
	}
	var result model.GetPrimaryStorageAllocatorStrategiesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取主存储类型列表
func GetPrimaryStorageTypes(params model.GetPrimaryStorageTypesRequest) mgresult.Result {
	//GET zstack/v1/primary-storage/types
	url := fmt.Sprintf("zstack/v1/primary-storage/types")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetPrimaryStorageTypesFailed, err)
	}
	var result model.GetPrimaryStorageTypesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取候选列表
func GetPrimaryStorageCandidatesForVolumeMigration(params model.GetPrimaryStorageCandidatesForVolumeMigrationRequest) mgresult.Result {
	//GET zstack/v1/primary-storage/volumes/{volumeUuid}/migration-candidates
	url := fmt.Sprintf("zstack/v1/primary-storage/volumes/%s/migration-candidates", params.VolumeUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetPrimaryStorageCandidatesForVolumeMigrationFailed, err)
	}
	var result model.GetPrimaryStorageCandidatesForVolumeMigrationResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取跨存储迁移可选物理机列表
func GetHostCandidatesForVmMigration(params model.GetHostCandidatesForVmMigrationRequest) mgresult.Result {
	//GET zstack/v1/primary-storage/hosts/{vmInstanceUuid}/migration-candidates
	url := fmt.Sprintf("zstack/v1/primary-storage/hosts/%s/migration-candidates", params.VmInstanceUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetHostCandidatesForVmMigrationFailed, err)
	}
	var result model.GetHostCandidatesForVmMigrationResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//跨主存储迁移云盘
func PrimaryStorageMigrateVolume(params model.PrimaryStorageMigrateVolumeRequest) mgresult.Result {
	//PUT zstack/v1/primary-storage/volumes/{volumeUuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/volumes/%s/actions", params.VolumeUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.PrimaryStorageMigrateVolumeFailed, err)
	}
	var result model.PrimaryStorageMigrateVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//跨主存储迁移云主机
func PrimaryStorageMigrateVm(params model.PrimaryStorageMigrateVmRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", params.VmInstanceUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.PrimaryStorageMigrateVmFailed, err)
	}
	var result model.PrimaryStorageMigrateVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取存储迁移候选列表
func GetPrimaryStorageCandidatesForVmMigration(params model.GetPrimaryStorageCandidatesForVmMigrationRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/storage-migration-candidates
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/storage-migration-candidates", params.VmInstanceUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetPrimaryStorageCandidatesForVmMigrationFailed, err)
	}
	var result model.GetPrimaryStorageCandidatesForVmMigrationResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取主存储的License信息
func GetPrimaryStorageLicenseInfo(params model.GetPrimaryStorageLicenseInfoRequest) mgresult.Result {
	//GET zstack/v1/primary-storage/{uuid}/license
	url := fmt.Sprintf("zstack/v1/primary-storage/%s/license", params.Uuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetPrimaryStorageLicenseInfoFailed, err)
	}
	var result model.GetPrimaryStorageLicenseInfoResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加本地存储为主存储
func AddLocalPrimaryStorage(params model.AddLocalPrimaryStorageRequest) mgresult.Result {
	//POST zstack/v1/primary-storage/local-storage
	url := fmt.Sprintf("zstack/v1/primary-storage/local-storage")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddLocalPrimaryStorageFailed, err)
	}
	var result model.AddLocalPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询本地存储资源引用
func QueryLocalStorageResourceRef(params model.QueryLocalStorageResourceRefRequest) mgresult.Result {
	//GET zstack/v1/primary-storage/local-storage/resource-refs
	url := fmt.Sprintf("zstack/v1/primary-storage/local-storage/resource-refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryLocalStorageResourceRefFailed, err)
	}
	var result model.QueryLocalStorageResourceRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//迁移本地存储上存放的云盘
func LocalStorageMigrateVolume(params model.LocalStorageMigrateVolumeRequest) mgresult.Result {
	//PUT zstack/v1/primary-storage/local-storage/volumes/{volumeUuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/local-storage/volumes/%s/actions", params.VolumeUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.LocalStorageMigrateVolumeFailed, err)
	}
	var result model.LocalStorageMigrateVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取主机本地存储容量
func GetLocalStorageHostDiskCapacity(params model.GetLocalStorageHostDiskCapacityRequest) mgresult.Result {
	//GET zstack/v1/primary-storage/local-storage/{primaryStorageUuid}/capacities
	url := fmt.Sprintf("zstack/v1/primary-storage/local-storage/%s/capacities", params.PrimaryStorageUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetLocalStorageHostDiskCapacityFailed, err)
	}
	var result model.GetLocalStorageHostDiskCapacityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取迁移本地存储物理机
func LocalStorageGetVolumeMigratableHosts(params model.LocalStorageGetVolumeMigratableHostsRequest) mgresult.Result {
	//GET zstack/v1/volumes/{volumeUuid}/migration-target-hosts
	url := fmt.Sprintf("zstack/v1/volumes/%s/migration-target-hosts", params.VolumeUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.LocalStorageGetVolumeMigratableHostsFailed, err)
	}
	var result model.LocalStorageGetVolumeMigratableHostsResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加NFS主存储
func AddNfsPrimaryStorage(params model.AddNfsPrimaryStorageRequest) mgresult.Result {
	//POST zstack/v1/primary-storage/nfs
	url := fmt.Sprintf("zstack/v1/primary-storage/nfs")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddNfsPrimaryStorageFailed, err)
	}
	var result model.AddNfsPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加一个共享挂载点的主存储
func AddSharedMountPointPrimaryStorage(params model.AddSharedMountPointPrimaryStorageRequest) mgresult.Result {
	//POST zstack/v1/primary-storage/smp
	url := fmt.Sprintf("zstack/v1/primary-storage/smp")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddSharedMountPointPrimaryStorageFailed, err)
	}
	var result model.AddSharedMountPointPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加Ceph主存储
func AddCephPrimaryStorage(params model.AddCephPrimaryStorageRequest) mgresult.Result {
	//POST zstack/v1/primary-storage/ceph
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddCephPrimaryStorageFailed, err)
	}
	var result model.AddCephPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询Ceph主存储
func QueryCephPrimaryStorage(params model.QueryCephPrimaryStorageRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryCephPrimaryStorageFailed, err)
	}
	var result model.QueryCephPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//为Ceph主存储添加mon节点
func AddMonToCephPrimaryStorage(params model.AddMonToCephPrimaryStorageRequest) mgresult.Result {
	//POST zstack/v1/primary-storage/ceph/{uuid}/mons
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph/%s/mons", params.Uuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddMonToCephPrimaryStorageFailed, err)
	}
	var result model.AddMonToCephPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从Ceph主存储删除mon节点
func RemoveMonFromCephPrimaryStorage(params model.RemoveMonFromCephPrimaryStorageRequest) mgresult.Result {
	//DELETE zstack/v1/primary-storage/ceph/{uuid}/mons?monHostnames={monHostnames}
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph/%s/mons", params.Uuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveMonFromCephPrimaryStorageFailed, err)
	}
	var result model.RemoveMonFromCephPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新Ceph主存储mon节点
func UpdateCephPrimaryStorageMon(params model.UpdateCephPrimaryStorageMonRequest) mgresult.Result {
	//PUT zstack/v1/primary-storage/ceph/mons/{monUuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph/mons/%s/actions", params.MonUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateCephPrimaryStorageMonFailed, err)
	}
	var result model.UpdateCephPrimaryStorageMonResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加Ceph主存储池
func AddCephPrimaryStoragePool(params model.AddCephPrimaryStoragePoolRequest) mgresult.Result {
	//POST zstack/v1/primary-storage/ceph/{primaryStorageUuid}/pools
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph/%s/pools", params.PrimaryStorageUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddCephPrimaryStoragePoolFailed, err)
	}
	var result model.AddCephPrimaryStoragePoolResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除Ceph主存储池
func DeleteCephPrimaryStoragePool(params model.DeleteCephPrimaryStoragePoolRequest) mgresult.Result {
	//DELETE zstack/v1/primary-storage/ceph/pools/{uuid}
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph/pools/%s", params.Uuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteCephPrimaryStoragePoolFailed, err)
	}
	var result model.DeleteCephPrimaryStoragePoolResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询Ceph主存储池
func QueryCephPrimaryStoragePool(params model.QueryCephPrimaryStoragePoolRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryCephPrimaryStoragePoolFailed, err)
	}
	var result model.QueryCephPrimaryStoragePoolResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新Ceph主存储池
func UpdateCephPrimaryStoragePool(params model.UpdateCephPrimaryStoragePoolRequest) mgresult.Result {
	//PUT zstack/v1/primary-storage/ceph/pools/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/ceph/pools/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateCephPrimaryStoragePoolFailed, err)
	}
	var result model.UpdateCephPrimaryStoragePoolResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加Shared Block主存储
func AddSharedBlockGroupPrimaryStorage(params model.AddSharedBlockGroupPrimaryStorageRequest) mgresult.Result {
	//POST zstack/v1/primary-storage/sharedblockgroup
	url := fmt.Sprintf("zstack/v1/primary-storage/sharedblockgroup")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddSharedBlockGroupPrimaryStorageFailed, err)
	}
	var result model.AddSharedBlockGroupPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询连接状态
func QuerySharedBlockGroupPrimaryStorageHostRef(params model.QuerySharedBlockGroupPrimaryStorageHostRefRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QuerySharedBlockGroupPrimaryStorageHostRefFailed, err)
	}
	var result model.QuerySharedBlockGroupPrimaryStorageHostRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询Shared Block主存储
func QuerySharedBlockGroupPrimaryStorage(params model.QuerySharedBlockGroupPrimaryStorageRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QuerySharedBlockGroupPrimaryStorageFailed, err)
	}
	var result model.QuerySharedBlockGroupPrimaryStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加LUN到Shared Block主存储
func AddSharedBlockToSharedBlockGroup(params model.AddSharedBlockToSharedBlockGroupRequest) mgresult.Result {
	//POST zstack/v1/primary-storage/sharedblockgroup/{uuid}/sharedblocks
	url := fmt.Sprintf("zstack/v1/primary-storage/sharedblockgroup/%s/sharedblocks", params.Uuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddSharedBlockToSharedBlockGroupFailed, err)
	}
	var result model.AddSharedBlockToSharedBlockGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取共享块设备候选清单
func GetSharedBlockCandidate(params model.GetSharedBlockCandidateRequest) mgresult.Result {
	//GET zstack/v1/primary-storage/sharedblockgroup/sharedblock-candidates
	url := fmt.Sprintf("zstack/v1/primary-storage/sharedblockgroup/sharedblock-candidates")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetSharedBlockCandidateFailed, err)
	}
	var result model.GetSharedBlockCandidateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//刷新共享块设备容量
func RefreshSharedblockDeviceCapacity(params model.RefreshSharedblockDeviceCapacityRequest) mgresult.Result {
	//POST zstack/v1/primary-storage/sharedblockgroup/{sharedBlockGroupUuid}/sharedblocks/{uuid}
	url := fmt.Sprintf("zstack/v1/primary-storage/sharedblockgroup/%s/sharedblocks/%s", params.SharedBlockGroupUuid, params.Uuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RefreshSharedblockDeviceCapacityFailed, err)
	}
	var result model.RefreshSharedblockDeviceCapacityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新共享块存储中共享块的diskUuid
func UpdateSharedBlockDiskUuid(params model.UpdateSharedBlockDiskUuidRequest) mgresult.Result {
	//PUT zstack/v1/primary-storage/sharedblockgroup/{sharedBlockGroupUuid}/sharedblocks/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/primary-storage/sharedblockgroup/%s/sharedblocks/%s/actions", params.SharedBlockGroupUuid, params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateSharedBlockDiskUuidFailed, err)
	}
	var result model.UpdateSharedBlockDiskUuidResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询共享块设备
func QuerySharedBlock(params model.QuerySharedBlockRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QuerySharedBlockFailed, err)
	}
	var result model.QuerySharedBlockResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
