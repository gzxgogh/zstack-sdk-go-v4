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

//删除镜像服务器
func DeleteBackupStorage(params model.DeleteBackupStorageRequest) models.Result {
	//DELETE zstack/v1/backup-storage/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/backup-storage/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteBackupStorageFailed, err.Error())
	}
	var result model.DeleteBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询镜像服务器
func QueryBackupStorage(params model.QueryBackupStorageRequest) models.Result {
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
		return models.Error(errcode.QueryBackupStorageFailed, err.Error())
	}
	var result model.QueryBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//重连镜像服务器
func ReconnectBackupStorage(params model.ReconnectBackupStorageRequest) models.Result {
	//PUT zstack/v1/backup-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ReconnectBackupStorageFailed, err.Error())
	}
	var result model.ReconnectBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更改镜像服务器可用状态
func ChangeBackupStorageState(params model.ChangeBackupStorageStateRequest) models.Result {
	//PUT zstack/v1/backup-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeBackupStorageStateFailed, err.Error())
	}
	var result model.ChangeBackupStorageStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取镜像服务器存储容量
func GetBackupStorageCapacity(params model.GetBackupStorageCapacityRequest) models.Result {
	//GET zstack/v1/backup-storage/capacities
	url := fmt.Sprintf("zstack/v1/backup-storage/capacities")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetBackupStorageCapacityFailed, err.Error())
	}
	var result model.GetBackupStorageCapacityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取镜像服务器类型列表
func GetBackupStorageTypes(params model.GetBackupStorageTypesRequest) models.Result {
	//GET zstack/v1/backup-storage/types
	url := fmt.Sprintf("zstack/v1/backup-storage/types")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetBackupStorageTypesFailed, err.Error())
	}
	var result model.GetBackupStorageTypesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新镜像服务器信息
func UpdateBackupStorage(params model.UpdateBackupStorageRequest) models.Result {
	//PUT zstack/v1/backup-storage/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateBackupStorageFailed, err.Error())
	}
	var result model.UpdateBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从镜像服务器导出镜像
func ExportImageFromBackupStorage(params model.ExportImageFromBackupStorageRequest) models.Result {
	//PUT zstack/v1/backup-storage/{backupStorageUuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/%s/actions", params.BackupStorageUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ExportImageFromBackupStorageFailed, err.Error())
	}
	var result model.ExportImageFromBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从镜像服务器删除导出的镜像
func DeleteExportedImageFromBackupStorage(params model.DeleteExportedImageFromBackupStorageRequest) models.Result {
	//DELETE zstack/v1/backup-storage/{backupStorageUuid}/exported-images/{imageUuid}
	url := fmt.Sprintf("zstack/v1/backup-storage/%s/exported-images/%s", params.BackupStorageUuid, params.ImageUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteExportedImageFromBackupStorageFailed, err.Error())
	}
	var result model.DeleteExportedImageFromBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//挂载镜像服务器至区域
func AttachBackupStorageToZone(params model.AttachBackupStorageToZoneRequest) models.Result {
	//POST zstack/v1/zones/{zoneUuid}/backup-storage/{backupStorageUuid}
	url := fmt.Sprintf("zstack/v1/zones/%s/backup-storage/%s", params.ZoneUUID, params.BackupStorageUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachBackupStorageToZoneFailed, err.Error())
	}
	var result model.AttachBackupStorageToZoneResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从区域中卸载已经挂载的镜像服务器
func DetachBackupStorageFromZone(params model.DetachBackupStorageFromZoneRequest) models.Result {
	//DELETE zstack/v1/zones/{zoneUuid}/backup-storage/{backupStorageUuid}
	url := fmt.Sprintf("zstack/v1/zones/%s/backup-storage/%s", params.ZoneUUID, params.BackupStorageUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachBackupStorageFromZoneFailed, err.Error())
	}
	var result model.DetachBackupStorageFromZoneResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//跨镜像服务器迁移镜像
func BackupStorageMigrateImage(params model.BackupStorageMigrateImageRequest) models.Result {
	//PUT zstack/v1/backup-storage/images/{imageUuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/images/%s/actions", params.ImageUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.BackupStorageMigrateImageFailed, err.Error())
	}
	var result model.BackupStorageMigrateImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取候选列表
func GetBackupStorageCandidatesForImageMigration(params model.GetBackupStorageCandidatesForImageMigrationRequest) models.Result {
	//GET zstack/v1/backup-storage/{srcBackupStorageUuid}/migration-candidates
	url := fmt.Sprintf("zstack/v1/backup-storage/%s/migration-candidates", params.SrcBackupStorageUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetBackupStorageCandidatesForImageMigrationFailed, err.Error())
	}
	var result model.GetBackupStorageCandidatesForImageMigrationResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加镜像仓库服务器
func AddImageStoreBackupStorage(params model.AddImageStoreBackupStorageRequest) models.Result {
	//POST zstack/v1/backup-storage/image-store
	url := fmt.Sprintf("zstack/v1/backup-storage/image-store")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddImageStoreBackupStorageFailed, err.Error())
	}
	var result model.AddImageStoreBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询镜像仓库服务器
func QueryImageStoreBackupStorage(params model.QueryImageStoreBackupStorageRequest) models.Result {
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
		return models.Error(errcode.QueryImageStoreBackupStorageFailed, err.Error())
	}
	var result model.QueryImageStoreBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//重连镜像仓库服务器
func ReconnectImageStoreBackupStorage(params model.ReconnectImageStoreBackupStorageRequest) models.Result {
	//PUT zstack/v1/backup-storage/image-store/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/image-store/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ReconnectImageStoreBackupStorageFailed, err.Error())
	}
	var result model.ReconnectImageStoreBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新镜像仓库服务器信息
func UpdateImageStoreBackupStorage(params model.UpdateImageStoreBackupStorageRequest) models.Result {
	//PUT zstack/v1/backup-storage/image-store/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/image-store/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateImageStoreBackupStorageFailed, err.Error())
	}
	var result model.UpdateImageStoreBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从镜像仓库回收磁盘空间
func ReclaimSpaceFromImageStore(params model.ReclaimSpaceFromImageStoreRequest) models.Result {
	//PUT zstack/v1/backup-storage/image-store/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/image-store/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ReclaimSpaceFromImageStoreFailed, err.Error())
	}
	var result model.ReclaimSpaceFromImageStoreResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加Ceph镜像服务器
func AddCephBackupStorage(params model.AddCephBackupStorageRequest) models.Result {
	//POST zstack/v1/backup-storage/ceph
	url := fmt.Sprintf("zstack/v1/backup-storage/ceph")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddCephBackupStorageFailed, err.Error())
	}
	var result model.AddCephBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询Ceph镜像服务器
func QueryCephBackupStorage(params model.QueryCephBackupStorageRequest) models.Result {
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
		return models.Error(errcode.QueryCephBackupStorageFailed, err.Error())
	}
	var result model.QueryCephBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新Ceph镜像服务器mon节点
func UpdateCephBackupStorageMon(params model.UpdateCephBackupStorageMonRequest) models.Result {
	//PUT zstack/v1/backup-storage/ceph/mons/{monUuid}/actions
	url := fmt.Sprintf("zstack/v1/backup-storage/ceph/mons/%s/actions", params.MonUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateCephBackupStorageMonFailed, err.Error())
	}
	var result model.UpdateCephBackupStorageMonResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//Ceph镜像服务器添加mon节点
func AddMonToCephBackupStorage(params model.AddMonToCephBackupStorageRequest) models.Result {
	//Post zstack/v1/backup-storage/ceph/{uuid}/mons
	url := fmt.Sprintf("zstack/v1/backup-storage/ceph/%s/mons", params.UUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddMonToCephBackupStorageFailed, err.Error())
	}
	var result model.RemoveMonFromCephBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//Ceph镜像服务器删除mon节点
func RemoveMonFromCephBackupStorage(params model.RemoveMonFromCephBackupStorageRequest) models.Result {
	//DELETE zstack/v1/backup-storage/ceph/{uuid}/mons?monHostnames={monHostnames}
	url := fmt.Sprintf("zstack/v1/backup-storage/ceph/%s/mons?monHostnames=%s", params.UUID, params.MonHostnames)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveMonFromCephBackupStorageFailed, err.Error())
	}
	var result model.RemoveMonFromCephBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
