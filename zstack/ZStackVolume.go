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
	"strings"
)

//创建云盘
func CreateDataVolume(params model.CreateDataVolumeRequest) mgresult.Result {
	//POST zstack/v1/volumes/data
	url := "zstack/v1/volumes/data"
	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateDataVolumeFailed, err)
	}
	var result model.CreateDataVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//删除云盘
func DeleteDataVolume(params model.DeleteDataVolumeRequest) mgresult.Result {
	//DELETE zstack/v1/volumes/{uuid}
	url := fmt.Sprintf("zstack/v1/volumes/%s", params.Uuid)
	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteDataVolumeFailed, err)
	}
	var result model.DeleteDataVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//彻底删除云盘
func ExpungeDataVolume(params model.ExpungeDataVolumeRequest) mgresult.Result {
	//PUT zstack/v1/volumes/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ExpungeDataVolumeFailed, err)
	}
	var result model.ExpungeDataVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//恢复云盘
func RecoverDataVolume(params model.RecoverDataVolumeRequest) mgresult.Result {
	//PUT zstack/v1/volumes/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RecoverDataVolumeFailed, err)
	}
	var result model.RecoverDataVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//开启或关闭云盘
func ChangeVolumeState(params model.ChangeVolumeStateRequest) mgresult.Result {
	//PUT zstack/v1/volumes/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeVolumeStateFailed, err)
	}
	var result model.ChangeVolumeStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//从镜像创建云盘
func CreateDataVolumeFromVolumeTemplate(params model.CreateDataVolumeFromVolumeTemplateRequest) mgresult.Result {
	//POST zstack/v1/volumes/data/from/data-volume-templates/{imageUuid}
	url := fmt.Sprintf("zstack/v1/volumes/data/from/data-volume-templates/%s", params.ImageUUID)
	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateDataVolumeFromVolumeTemplateFailed, err)
	}
	var result model.CreateDataVolumeFromVolumeTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//从快照创建云盘
func CreateDataVolumeFromVolumeSnapshot(params model.CreateDataVolumeFromVolumeSnapshotRequest) mgresult.Result {
	//POST zstack/v1/volumes/data/from/volume-snapshots/{volumeSnapshotUuid}
	url := fmt.Sprintf("zstack/v1/volumes/data/from/volume-snapshots/%s", params.VolumeSnapshotUuid)
	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateDataVolumeFromVolumeSnapshotFailed, err)
	}
	var result model.CreateDataVolumeFromVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//获取云盘清单
func QueryVolume(params model.QueryVolumeRequest) mgresult.Result {
	//GET zstack/v1/volumes
	//GET zstack/v1/volumes/{uuid}
	var url string
	if params.Uuid == "" {
		url = "zstack/v1/volumes"
	} else {
		url = fmt.Sprintf("zstack/v1/volumes/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVolumeFailed, err)
	}
	var result model.QueryVolumeResponse
	utils.FromJSON(dataStr, &result)
	tagKeyLst := QueryTags(params)

	var inventories []model.VolumeInventory
	for _, volumeInventory := range result.Inventories {
		uuid := volumeInventory.UUID
		if tagKeyLst[uuid] != nil {
			tagArr := strings.Split(fmt.Sprint(tagKeyLst[uuid]["tag"]), "::")
			volumeInventory.WWN = tagArr[len(tagArr)-1]
		}
		inventories = append(inventories, volumeInventory)
	}
	result.Inventories = inventories
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

func QueryTags(params model.QueryVolumeRequest) map[string]map[string]interface{} {
	params.Uuid = ""
	url := "zstack/v1/system-tags"
	dataStr, err := request.Get(url, params)
	if err != nil {
		return nil
	}
	var result map[string]interface{}
	utils.FromJSON(dataStr, &result)
	tagLst := result["inventories"].([]interface{})
	tagKeyLst := make(map[string]map[string]interface{})
	for _, item := range tagLst {
		obj := item.(map[string]interface{})
		resourceUuid := fmt.Sprint(obj["resourceUuid"])
		if fmt.Sprint(obj["inherent"]) == "true" {
			tagKeyLst[resourceUuid] = obj
		}
	}
	return tagKeyLst
}

//获取云盘格式
func GetVolumeFormat(params model.GetVolumeFormatRequest) mgresult.Result {
	//GET zstack/v1/volumes/formats
	url := fmt.Sprintf("zstack/v1/volumes/formats")
	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVolumeFormatFailed, err)
	}
	var result model.GetVolumeFormatResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//获取云盘支持的类型的能力
func GetVolumeCapabilities(params model.GetVolumeCapabilitiesRequest) mgresult.Result {
	//GET zstack/v1/volumes/formats
	url := fmt.Sprintf("zstack/v1/volumes/formats")
	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVolumeCapabilitiesFailed, err)
	}
	var result model.GetVolumeCapabilitiesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//同步云盘大小
func SyncVolumeSize(params model.SyncVolumeSizeRequest) mgresult.Result {
	//PUT zstack/v1/volumes/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SyncVolumeSizeFailed, err)
	}
	var result model.SyncVolumeSizeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//扩展根云盘
func ResizeRootVolume(params model.ResizeRootVolumeRequest) mgresult.Result {
	//PUT zstack/v1/volumes/resize/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/resize/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ResizeRootVolumeFailed, err)
	}
	var result model.ResizeRootVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//扩展数据云盘
func ResizeDataVolume(params model.ResizeDataVolumeRequest) mgresult.Result {
	//PUT zstack/v1/volumes/data/resize/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/data/resize/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ResizeDataVolumeFailed, err)
	}
	var result model.ResizeDataVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//修改云盘属性
func UpdateVolume(params model.UpdateVolumeRequest) mgresult.Result {
	//PUT zstack/v1/volumes/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateVolumeFailed, err)
	}
	var result model.UpdateVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//设置云盘限速
func SetVolumeQoS(params model.SetVolumeQoSRequest) mgresult.Result {
	//PUT zstack/v1/volumes/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVolumeQoSFailed, err)
	}
	var result model.SetVolumeQoSResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//获取云盘限速
func GetVolumeQoS(params model.GetVolumeQoSRequest) mgresult.Result {
	//GET zstack/v1/volumes/{uuid}/qos
	url := fmt.Sprintf("zstack/v1/volumes/%s/qos", params.Uuid)
	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVolumeQoSFailed, err)
	}
	var result model.GetVolumeQoSResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//取消云盘网卡限速
func DeleteVolumeQoS(params model.DeleteVolumeQoSRequest) mgresult.Result {
	//DELETE zstack/v1/volumes/{uuid}/qos
	url := fmt.Sprintf("zstack/v1/volumes/%s/qos", params.Uuid)
	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVolumeQoSFailed, err)
	}
	var result model.DeleteVolumeQoSResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//获取云盘可被加载的云主机列表
func GetDataVolumeAttachableVm(params model.GetDataVolumeAttachableVmRequest) mgresult.Result {
	//GET zstack/v1/volumes/{volumeUuid}/candidate-vm-instances
	url := fmt.Sprintf("zstack/v1/volumes/%s/candidate-vm-instances", params.VolumeUuid)
	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetDataVolumeAttachableVmFailed, err)
	}
	var result model.GetDataVolumeAttachableVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//挂载云盘到云主机上
func AttachDataVolumeToVm(params model.AttachDataVolumeToVmRequest) mgresult.Result {
	//POST zstack/v1/volumes/{volumeUuid}/vm-instances/{vmInstanceUuid}
	url := fmt.Sprintf("zstack/v1/volumes/%s/vm-instances/%s", params.VolumeUuid, params.VmInstanceUuid)
	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachDataVolumeToVmFailed, err)
	}
	var result model.AttachDataVolumeToVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//从云主机上卸载云盘
func DetachDataVolumeFromVm(params model.DetachDataVolumeFromVmRequest) mgresult.Result {
	//DELETE zstack/v1/volumes/{uuid}/vm-instances?vmUuid={vmUuid}
	url := fmt.Sprintf("zstack/v1/volumes/%s/vm-instances", params.Uuid)
	dataStr, err := request.DeleteUrlWithParams(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachDataVolumeFromVmFailed, err)
	}
	var result model.DetachDataVolumeFromVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//从云盘创建快照
func CreateVolumeSnapshot(params model.CreateVolumeSnapshotRequest) mgresult.Result {
	//POST zstack/v1/volumes/{volumeUuid}/volume-snapshots
	url := fmt.Sprintf("zstack/v1/volumes/%s/volume-snapshots", params.VolumeUuid)
	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateVolumeSnapshotFailed, err)
	}
	var result model.CreateVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//查询云盘快照
func QueryVolumeSnapshot(params model.QueryVolumeSnapshotRequest) mgresult.Result {
	//GET zstack/v1/volume-snapshots
	//GET zstack/v1/volume-snapshots/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/volume-snapshots")
	} else {
		url = fmt.Sprintf("zstack/v1/volume-snapshots/%s", params.Uuid)
	}
	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVolumeSnapshotFailed, err)
	}
	var result model.QueryVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//查询快照树
func QueryVolumeSnapshotTree(params model.QueryVolumeSnapshotTreeRequest) mgresult.Result {
	//GET zstack/v1/volume-snapshots/trees
	//GET zstack/v1/volume-snapshots/trees/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/volume-snapshots/trees")
	} else {
		url = fmt.Sprintf("zstack/v1/volume-snapshots/trees/%s", params.Uuid)
	}
	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVolumeSnapshotTreeFailed, err)
	}
	var result model.QueryVolumeSnapshotTreeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//更新云盘快照信息
func UpdateVolumeSnapshot(params model.UpdateVolumeSnapshotRequest) mgresult.Result {
	//PUT zstack/v1/volume-snapshots/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volume-snapshots/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateVolumeSnapshotFailed, err)
	}
	var result model.UpdateVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//删除云盘快照
func DeleteVolumeSnapshot(params model.DeleteVolumeSnapshotRequest) mgresult.Result {
	//DELETE zstack/v1/volume-snapshots/{uuid}
	url := fmt.Sprintf("zstack/v1/volume-snapshots/%s", params.Uuid)
	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVolumeSnapshotFailed, err)
	}
	var result model.DeleteVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//将云盘回滚至指定快照
func RevertVolumeFromSnapshot(params model.RevertVolumeFromSnapshotRequest) mgresult.Result {
	//PUT zstack/v1/volume-snapshots/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volume-snapshots/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RevertVolumeFromSnapshotFailed, err)
	}
	var result model.RevertVolumeFromSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//获取快照容量
func GetVolumeSnapshotSize(params model.GetVolumeSnapshotSizeRequest) mgresult.Result {
	//PUT zstack/v1/volume-snapshots/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volume-snapshots/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVolumeSnapshotSizeFailed, err)
	}
	var result model.GetVolumeSnapshotSizeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}

//快照瘦身
func ShrinkVolumeSnapshot(params model.ShrinkVolumeSnapshotRequest) mgresult.Result {
	//PUT zstack/v1/volume-snapshots/shrink/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volume-snapshots/shrink/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ShrinkVolumeSnapshotFailed, err)
	}
	var result model.ShrinkVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return mgresult.Success(result)
}
