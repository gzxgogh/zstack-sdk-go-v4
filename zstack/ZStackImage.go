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

//添加镜像
func AddImage(params model.AddImageRequest) mgresult.Result {
	//POST zstack/v1/images
	url := fmt.Sprintf("zstack/v1/images")
	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddImageFailed, err)
	}
	var result model.AddImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除镜像
func DeleteImage(params model.DeleteImageRequest) mgresult.Result {
	//DELETE zstack/v1/images/{uuid}?backupStorageUuids={backupStorageUuids}
	url := fmt.Sprintf("zstack/v1/images/%s", params.Uuid)
	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteImageFailed, err)
	}
	var result model.DeleteImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//彻底删除镜像
func ExpungeImage(params model.ExpungeImageRequest) mgresult.Result {
	//PUT zstack/v1/images/{imageUuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.ImageUuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ExpungeImageFailed, err)
	}
	var result model.ExpungeImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询镜像
func QueryImage(params model.QueryImageRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryImageFailed, err)
	}
	var result model.QueryImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//恢复镜像
func RecoverImage(params model.RecoverImageRequest) mgresult.Result {
	//PUT zstack/v1/images/{imageUuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.ImageUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RecoverImageFailed, err)
	}
	var result model.RecoverImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//修改镜像状态
func ChangeImageState(params model.ChangeImageStateRequest) mgresult.Result {
	//PUT zstack/v1/images/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeImageStateFailed, err)
	}
	var result model.ChangeImageStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新镜像信息
func UpdateImage(params model.UpdateImageRequest) mgresult.Result {
	//PUT zstack/v1/images/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateImageFailed, err)
	}
	var result model.UpdateImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//刷新镜像大小信息
func SyncImageSize(params model.SyncImageSizeRequest) mgresult.Result {
	//PUT zstack/v1/images/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SyncImageSizeFailed, err)
	}
	var result model.SyncImageSizeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取镜像服务器候选
func GetCandidateBackupStorageForCreatingImage(params model.GetCandidateBackupStorageForCreatingImageRequest) mgresult.Result {
	url := ""
	if params.VolumeUuid != "" {
		url = fmt.Sprintf("zstack/v1/images/volumes/%s/candidate-backup-storage", params.VolumeUuid)
	} else {
		url = fmt.Sprintf("zstack/v1/images/volume-snapshots/%s/candidate-backup-storage", params.VolumeSnapshotUuid)
	}
	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCandidateBackupStorageForCreatingImageFailed, err)
	}
	var result model.GetCandidateBackupStorageForCreatingImageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从根云盘创建根云盘镜像
func CreateRootVolumeTemplateFromRootVolume(params model.CreateRootVolumeTemplateFromRootVolumeRequest) mgresult.Result {
	//POST zstack/v1/images/root-volume-templates/from/volumes/{rootVolumeUuid}
	url := fmt.Sprintf("zstack/v1/images/root-volume-templates/from/volumes/%s", params.RootVolumeUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateRootVolumeTemplateFromRootVolumeFailed, err)
	}
	var result model.CreateRootVolumeTemplateFromRootVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从云盘快照创建根云盘镜像
func CreateRootVolumeTemplateFromVolumeSnapshot(params model.CreateRootVolumeTemplateFromVolumeSnapshotRequest) mgresult.Result {
	//POST zstack/v1/images/root-volume-templates/from/volume-snapshots/{snapshotUuid}
	url := fmt.Sprintf("zstack/v1/images/root-volume-templates/from/volume-snapshots/%s", params.SnapshotUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateRootVolumeTemplateFromVolumeSnapshotFailed, err)
	}
	var result model.CreateRootVolumeTemplateFromVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从云盘创建数据云盘镜像
func CreateDataVolumeTemplateFromVolume(params model.CreateDataVolumeTemplateFromVolumeRequest) mgresult.Result {
	//POST zstack/v1/images/data-volume-templates/from/volumes/{volumeUuid}
	url := fmt.Sprintf("zstack/v1/images/data-volume-templates/from/volumes/%s", params.VolumeUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateDataVolumeTemplateFromVolumeFailed, err)
	}
	var result model.CreateDataVolumeTemplateFromVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从云盘快照创建数据云盘镜像
func CreateDataVolumeTemplateFromVolumeSnapshot(params model.CreateDataVolumeTemplateFromVolumeSnapshotRequest) mgresult.Result {
	//POST zstack/v1/images/data-volume-templates/from/volume-snapshots/{snapshotUuid}
	url := fmt.Sprintf("zstack/v1/images/data-volume-templates/from/volumes/%s", params.SnapshotUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateDataVolumeTemplateFromVolumeSnapshotFailed, err)
	}
	var result model.CreateDataVolumeTemplateFromVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取镜像Qga
func GetImageQga(params model.GetImageQgaRequest) mgresult.Result {
	//GET zstack/v1/images/{uuid}/qga
	url := fmt.Sprintf("zstack/v1/images/%s/qga", params.Uuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetImageQgaFailed, err)
	}
	var result model.GetImageQgaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//设置镜像Qga
func SetImageQga(params model.SetImageQgaRequest) mgresult.Result {
	//PUT zstack/v1/images/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetImageQgaFailed, err)
	}
	var result model.SetImageQgaResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//设置镜像启动模式
func SetImageBootMode(params model.SetImageBootModeRequest) mgresult.Result {
	//PUT zstack/v1/images/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/images/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetImageBootModeFailed, err)
	}
	var result model.SetImageBootModeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取上传镜像任务详情
func GetUploadImageJobDetails(params model.GetUploadImageJobDetailsRequest) mgresult.Result {
	//PUT zstack/v1/images/upload-job/details/{imageId}
	url := fmt.Sprintf("zstack/v1/images/upload-job/details/%s", params.ImageId)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetUploadImageJobDetailsFailed, err)
	}
	var result model.GetUploadImageJobDetailsResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
