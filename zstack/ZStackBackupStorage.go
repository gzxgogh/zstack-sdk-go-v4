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

//删除镜像服务器
func DeleteBackupStorage(params model.DeleteBackupStorageRequest) mgresult.Result {
	//DELETE zstack/v1/backup-storage/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/backup-storage/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteBackupStorageFailed, err)
	}
	var result model.DeleteBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询镜像服务器
func QueryBackupStorage(params model.QueryBackupStorageRequest) mgresult.Result {
	//GET zstack/v1/backup-storage
	//GET zstack/v1/backup-storage/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/backup-storage")
	} else {
		url = fmt.Sprintf("zstack/v1/backup-storage/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryBackupStorageFailed, err)
	}
	var result model.QueryBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//重连镜像服务器
func ReconnectBackupStorage(params model.ReconnectBackupStorageRequest) mgresult.Result {
	//PUT zstack/v1/backup-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ReconnectBackupStorageFailed, err)
	}
	var result model.ReconnectBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更改镜像服务器可用状态
func ChangeBackupStorageState(params model.ChangeBackupStorageStateRequest) mgresult.Result {
	//PUT zstack/v1/backup-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeBackupStorageStateFailed, err)
	}
	var result model.ChangeBackupStorageStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取镜像服务器存储容量
func GetBackupStorageCapacity(params model.GetBackupStorageCapacityRequest) mgresult.Result {
	//GET zstack/v1/backup-storage/capacities
	url := fmt.Sprintf("zstack/v1/backup-storage/capacities")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetBackupStorageCapacityFailed, err)
	}
	var result model.GetBackupStorageCapacityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取镜像服务器类型列表
func GetBackupStorageTypes(params model.GetBackupStorageTypesRequest) mgresult.Result {
	//GET zstack/v1/backup-storage/types
	url := fmt.Sprintf("zstack/v1/backup-storage/types")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetBackupStorageTypesFailed, err)
	}
	var result model.GetBackupStorageTypesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新镜像服务器信息
func UpdateBackupStorage(params model.UpdateBackupStorageRequest) mgresult.Result {
	//PUT zstack/v1/backup-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateBackupStorageFailed, err)
	}
	var result model.UpdateBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从镜像服务器导出镜像
func ExportImageFromBackupStorage(params model.ExportImageFromBackupStorageRequest) mgresult.Result {
	//PUT zstack/v1/backup-storage/{backupStorageUuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/%s/actions", params.BackupStorageUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ExportImageFromBackupStorageFailed, err)
	}
	var result model.ExportImageFromBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从镜像服务器删除导出的镜像
func DeleteExportedImageFromBackupStorage(params model.DeleteExportedImageFromBackupStorageRequest) mgresult.Result {
	//DELETE zstack/v1/backup-storage/{backupStorageUuid}/exported-images/{imageUuid}
	url := fmt.Sprintf("zstack/v1/backup-storage/%s/exported-images/%s", params.BackupStorageUuid, params.ImageUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteExportedImageFromBackupStorageFailed, err)
	}
	var result model.DeleteExportedImageFromBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//挂载镜像服务器至区域
func AttachBackupStorageToZone(params model.AttachBackupStorageToZoneRequest) mgresult.Result {
	//POST zstack/v1/zones/{zoneUuid}/backup-storage/{backupStorageUuid}
	url := fmt.Sprintf("zstack/v1/zones/%s/backup-storage/%s", params.ZoneUUID, params.BackupStorageUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachBackupStorageToZoneFailed, err)
	}
	var result model.AttachBackupStorageToZoneResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从区域中卸载已经挂载的镜像服务器
func DetachBackupStorageFromZone(params model.DetachBackupStorageFromZoneRequest) mgresult.Result {
	//DELETE zstack/v1/zones/{zoneUuid}/backup-storage/{backupStorageUuid}
	url := fmt.Sprintf("zstack/v1/zones/%s/backup-storage/%s", params.ZoneUUID, params.BackupStorageUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachBackupStorageFromZoneFailed, err)
	}
	var result model.DetachBackupStorageFromZoneResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//跨镜像服务器迁移镜像
func BackupStorageMigrateImage(params model.BackupStorageMigrateImageRequest) mgresult.Result {
	//PUT zstack/v1/backup-storage/images/{imageUuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/images/%s/actions", params.ImageUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.BackupStorageMigrateImageFailed, err)
	}
	var result model.BackupStorageMigrateImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取候选列表
func GetBackupStorageCandidatesForImageMigration(params model.GetBackupStorageCandidatesForImageMigrationRequest) mgresult.Result {
	//GET zstack/v1/backup-storage/{srcBackupStorageUuid}/migration-candidates
	url := fmt.Sprintf("zstack/v1/backup-storage/%s/migration-candidates", params.SrcBackupStorageUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetBackupStorageCandidatesForImageMigrationFailed, err)
	}
	var result model.GetBackupStorageCandidatesForImageMigrationResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加镜像仓库服务器
func AddImageStoreBackupStorage(params model.AddImageStoreBackupStorageRequest) mgresult.Result {
	//POST zstack/v1/backup-storage/image-store
	url := fmt.Sprintf("zstack/v1/backup-storage/image-store")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddImageStoreBackupStorageFailed, err)
	}
	var result model.AddImageStoreBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询镜像仓库服务器
func QueryImageStoreBackupStorage(params model.QueryImageStoreBackupStorageRequest) mgresult.Result {
	//GET zstack/v1/backup-storage/image-store
	//GET zstack/v1/backup-storage/image-store/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/backup-storage/image-store")
	} else {
		url = fmt.Sprintf("zstack/v1/backup-storage/image-store/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryImageStoreBackupStorageFailed, err)
	}
	var result model.QueryImageStoreBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//重连镜像仓库服务器
func ReconnectImageStoreBackupStorage(params model.ReconnectImageStoreBackupStorageRequest) mgresult.Result {
	//PUT zstack/v1/backup-storage/image-store/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/image-store/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ReconnectImageStoreBackupStorageFailed, err)
	}
	var result model.ReconnectImageStoreBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新镜像仓库服务器信息
func UpdateImageStoreBackupStorage(params model.UpdateImageStoreBackupStorageRequest) mgresult.Result {
	//PUT zstack/v1/backup-storage/image-store/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/image-store/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateImageStoreBackupStorageFailed, err)
	}
	var result model.UpdateImageStoreBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从镜像仓库回收磁盘空间
func ReclaimSpaceFromImageStore(params model.ReclaimSpaceFromImageStoreRequest) mgresult.Result {
	//PUT zstack/v1/backup-storage/image-store/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/image-store/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ReclaimSpaceFromImageStoreFailed, err)
	}
	var result model.ReclaimSpaceFromImageStoreResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加Ceph镜像服务器
func AddCephBackupStorage(params model.AddCephBackupStorageRequest) mgresult.Result {
	//POST zstack/v1/backup-storage/ceph
	url := fmt.Sprintf("zstack/v1/backup-storage/ceph")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddCephBackupStorageFailed, err)
	}
	var result model.AddCephBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询Ceph镜像服务器
func QueryCephBackupStorage(params model.QueryCephBackupStorageRequest) mgresult.Result {
	//GET zstack/v1/backup-storage/ceph
	//GET zstack/v1/backup-storage/ceph/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/backup-storage/ceph")
	} else {
		url = fmt.Sprintf("zstack/v1/backup-storage/ceph/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryCephBackupStorageFailed, err)
	}
	var result model.QueryCephBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新Ceph镜像服务器mon节点
func UpdateCephBackupStorageMon(params model.UpdateCephBackupStorageMonRequest) mgresult.Result {
	//PUT zstack/v1/backup-storage/ceph/mons/{monUuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/ceph/mons/%s/actions", params.MonUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateCephBackupStorageMonFailed, err)
	}
	var result model.UpdateCephBackupStorageMonResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//Ceph镜像服务器添加mon节点
func AddMonToCephBackupStorage(params model.AddMonToCephBackupStorageRequest) mgresult.Result {
	//Post zstack/v1/backup-storage/ceph/{uuid}/mons
	url := fmt.Sprintf("zstack/v1/backup-storage/ceph/%s/mons", params.UUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddMonToCephBackupStorageFailed, err)
	}
	var result model.RemoveMonFromCephBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//Ceph镜像服务器删除mon节点
func RemoveMonFromCephBackupStorage(params model.RemoveMonFromCephBackupStorageRequest) mgresult.Result {
	//DELETE zstack/v1/backup-storage/ceph/{uuid}/mons?monHostnames={monHostnames}
	url := fmt.Sprintf("zstack/v1/backup-storage/ceph/%s/mons?monHostnames=%s", params.UUID, params.MonHostnames)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveMonFromCephBackupStorageFailed, err)
	}
	var result model.RemoveMonFromCephBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
