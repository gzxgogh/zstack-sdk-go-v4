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

//添加镜像
func AddImage(params model.AddImageRequest) models.Result {
	//POST zstack/v1/images
	url := fmt.Sprintf("zstack/v1/images")
	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddImageFailed, err.Error())
	}
	var result model.AddImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除镜像
func DeleteImage(params model.DeleteImageRequest) models.Result {
	//DELETE zstack/v1/images/{uuid}?backupStorageUuids={backupStorageUuids}
	url := fmt.Sprintf("zstack/v1/images/%s", params.Uuid)
	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteImageFailed, err.Error())
	}
	var result model.DeleteImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//彻底删除镜像
func ExpungeImage(params model.ExpungeImageRequest) models.Result {
	//PUT zstack/v1/images/{imageUuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.ImageUuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ExpungeImageFailed, err.Error())
	}
	var result model.ExpungeImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询镜像
func QueryImage(params model.QueryImageRequest) models.Result {
	//GET zstack/v1/images
	//GET zstack/v1/images/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/images")
	} else {
		url = fmt.Sprintf("zstack/v1/images/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryImageFailed, err.Error())
	}
	var result model.QueryImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//恢复镜像
func RecoverImage(params model.RecoverImageRequest) models.Result {
	//PUT zstack/v1/images/{imageUuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.ImageUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.RecoverImageFailed, err.Error())
	}
	var result model.RecoverImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改镜像状态
func ChangeImageState(params model.ChangeImageStateRequest) models.Result {
	//PUT zstack/v1/images/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeImageStateFailed, err.Error())
	}
	var result model.ChangeImageStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新镜像信息
func UpdateImage(params model.UpdateImageRequest) models.Result {
	//PUT zstack/v1/images/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateImageFailed, err.Error())
	}
	var result model.UpdateImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//刷新镜像大小信息
func SyncImageSize(params model.SyncImageSizeRequest) models.Result {
	//PUT zstack/v1/images/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.SyncImageSizeFailed, err.Error())
	}
	var result model.SyncImageSizeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取镜像服务器候选
func GetCandidateBackupStorageForCreatingImage(params model.GetCandidateBackupStorageForCreatingImageRequest) models.Result {
	url := ""
	if params.VolumeUuid != "" {
		url = fmt.Sprintf("zstack/v1/images/volumes/%s/candidate-backup-storage", params.VolumeUuid)
	} else {
		url = fmt.Sprintf("zstack/v1/images/volume-snapshots/%s/candidate-backup-storage", params.VolumeSnapshotUuid)
	}
	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetCandidateBackupStorageForCreatingImageFailed, err.Error())
	}
	var result model.GetCandidateBackupStorageForCreatingImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从根云盘创建根云盘镜像
func CreateRootVolumeTemplateFromRootVolume(params model.CreateRootVolumeTemplateFromRootVolumeRequest) models.Result {
	//POST zstack/v1/images/root-volume-templates/from/volumes/{rootVolumeUuid}
	url := fmt.Sprintf("zstack/v1/images/root-volume-templates/from/volumes/%s", params.RootVolumeUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateRootVolumeTemplateFromRootVolumeFailed, err.Error())
	}
	var result model.CreateRootVolumeTemplateFromRootVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从云盘快照创建根云盘镜像
func CreateRootVolumeTemplateFromVolumeSnapshot(params model.CreateRootVolumeTemplateFromVolumeSnapshotRequest) models.Result {
	//POST zstack/v1/images/root-volume-templates/from/volume-snapshots/{snapshotUuid}
	url := fmt.Sprintf("zstack/v1/images/root-volume-templates/from/volume-snapshots/%s", params.SnapshotUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateRootVolumeTemplateFromVolumeSnapshotFailed, err.Error())
	}
	var result model.CreateRootVolumeTemplateFromVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从云盘创建数据云盘镜像
func CreateDataVolumeTemplateFromVolume(params model.CreateDataVolumeTemplateFromVolumeRequest) models.Result {
	//POST zstack/v1/images/data-volume-templates/from/volumes/{volumeUuid}
	url := fmt.Sprintf("zstack/v1/images/data-volume-templates/from/volumes/%s", params.VolumeUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateDataVolumeTemplateFromVolumeFailed, err.Error())
	}
	var result model.CreateDataVolumeTemplateFromVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从云盘快照创建数据云盘镜像
func CreateDataVolumeTemplateFromVolumeSnapshot(params model.CreateDataVolumeTemplateFromVolumeSnapshotRequest) models.Result {
	//POST zstack/v1/images/data-volume-templates/from/volume-snapshots/{snapshotUuid}
	url := fmt.Sprintf("zstack/v1/images/data-volume-templates/from/volumes/%s", params.SnapshotUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.CreateDataVolumeTemplateFromVolumeSnapshotFailed, err.Error())
	}
	var result model.CreateDataVolumeTemplateFromVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取镜像Qga
func GetImageQga(params model.GetImageQgaRequest) models.Result {
	//GET zstack/v1/images/{uuid}/qga
	url := fmt.Sprintf("zstack/v1/images/%s/qga", params.Uuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetImageQgaFailed, err.Error())
	}
	var result model.GetImageQgaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//设置镜像Qga
func SetImageQga(params model.SetImageQgaRequest) models.Result {
	//PUT zstack/v1/images/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.SetImageQgaFailed, err.Error())
	}
	var result model.SetImageQgaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//设置镜像启动模式
func SetImageBootMode(params model.SetImageBootModeRequest) models.Result {
	//PUT zstack/v1/images/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.SetImageBootModeFailed, err.Error())
	}
	var result model.SetImageBootModeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取上传镜像任务详情
func GetUploadImageJobDetails(params model.GetUploadImageJobDetailsRequest) models.Result {
	//PUT zstack/v1/images/upload-job/details/{imageId}
	url := fmt.Sprintf("zstack/v1/images/upload-job/details/%s", params.ImageId)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.GetUploadImageJobDetailsFailed, err.Error())
	}
	var result model.GetUploadImageJobDetailsResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
