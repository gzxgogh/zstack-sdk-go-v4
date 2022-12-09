package zstack

import (
	"fmt"
	"github.com/gzxgogh/zstack-sdk-go-v4/errcode"
	"github.com/gzxgogh/zstack-sdk-go-v4/model"
	"github.com/gzxgogh/zstack-sdk-go-v4/request"
	"github.com/maczh/mgin/logs"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
	"strings"
)

//创建云盘
func CreateDataVolume(params model.CreateDataVolumeRequest) models.Result {
	//POST zstack/v1/volumes/data
	url := "zstack/v1/volumes/data"
	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateDataVolumeFailed, err.Error())
	}
	var result model.CreateDataVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//删除云盘
func DeleteDataVolume(params model.DeleteDataVolumeRequest) models.Result {
	//DELETE zstack/v1/volumes/{uuid}
	url := fmt.Sprintf("zstack/v1/volumes/%s", params.Uuid)
	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteDataVolumeFailed, err.Error())
	}
	var result model.DeleteDataVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//彻底删除云盘
func ExpungeDataVolume(params model.ExpungeDataVolumeRequest) models.Result {
	//PUT zstack/v1/volumes/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ExpungeDataVolumeFailed, err.Error())
	}
	var result model.ExpungeDataVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//恢复云盘
func RecoverDataVolume(params model.RecoverDataVolumeRequest) models.Result {
	//PUT zstack/v1/volumes/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.RecoverDataVolumeFailed, err.Error())
	}
	var result model.RecoverDataVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//开启或关闭云盘
func ChangeVolumeState(params model.ChangeVolumeStateRequest) models.Result {
	//PUT zstack/v1/volumes/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeVolumeStateFailed, err.Error())
	}
	var result model.ChangeVolumeStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//从镜像创建云盘
func CreateDataVolumeFromVolumeTemplate(params model.CreateDataVolumeFromVolumeTemplateRequest) models.Result {
	//POST zstack/v1/volumes/data/from/data-volume-templates/{imageUuid}
	url := fmt.Sprintf("zstack/v1/volumes/data/from/data-volume-templates/%s", params.ImageUUID)
	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateDataVolumeFromVolumeTemplateFailed, err.Error())
	}
	var result model.CreateDataVolumeFromVolumeTemplateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//从快照创建云盘
func CreateDataVolumeFromVolumeSnapshot(params model.CreateDataVolumeFromVolumeSnapshotRequest) models.Result {
	//POST zstack/v1/volumes/data/from/volume-snapshots/{volumeSnapshotUuid}
	url := fmt.Sprintf("zstack/v1/volumes/data/from/volume-snapshots/%s", params.VolumeSnapshotUuid)
	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateDataVolumeFromVolumeSnapshotFailed, err.Error())
	}
	var result model.CreateDataVolumeFromVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//获取云盘清单
func QueryVolume(params model.QueryVolumeRequest) models.Result {
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
		return models.Error(errcode.QueryVolumeFailed, err.Error())
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

	return models.Success(result)
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
func GetVolumeFormat(params model.GetVolumeFormatRequest) models.Result {
	//GET zstack/v1/volumes/formats
	url := fmt.Sprintf("zstack/v1/volumes/formats")
	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetVolumeFormatFailed, err.Error())
	}
	var result model.GetVolumeFormatResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//获取云盘支持的类型的能力
func GetVolumeCapabilities(params model.GetVolumeCapabilitiesRequest) models.Result {
	//GET zstack/v1/volumes/formats
	url := fmt.Sprintf("zstack/v1/volumes/formats")
	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetVolumeCapabilitiesFailed, err.Error())
	}
	var result model.GetVolumeCapabilitiesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//同步云盘大小
func SyncVolumeSize(params model.SyncVolumeSizeRequest) models.Result {
	//PUT zstack/v1/volumes/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.SyncVolumeSizeFailed, err.Error())
	}
	var result model.SyncVolumeSizeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//扩展根云盘
func ResizeRootVolume(params model.ResizeRootVolumeRequest) models.Result {
	//PUT zstack/v1/volumes/resize/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/resize/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ResizeRootVolumeFailed, err.Error())
	}
	var result model.ResizeRootVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//扩展数据云盘
func ResizeDataVolume(params model.ResizeDataVolumeRequest) models.Result {
	//PUT zstack/v1/volumes/data/resize/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/data/resize/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ResizeDataVolumeFailed, err.Error())
	}
	var result model.ResizeDataVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//修改云盘属性
func UpdateVolume(params model.UpdateVolumeRequest) models.Result {
	//PUT zstack/v1/volumes/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateVolumeFailed, err.Error())
	}
	var result model.UpdateVolumeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//设置云盘限速
func SetVolumeQoS(params model.SetVolumeQoSRequest) models.Result {
	//PUT zstack/v1/volumes/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volumes/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.SetVolumeQoSFailed, err.Error())
	}
	var result model.SetVolumeQoSResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//获取云盘限速
func GetVolumeQoS(params model.GetVolumeQoSRequest) models.Result {
	//GET zstack/v1/volumes/{uuid}/qos
	url := fmt.Sprintf("zstack/v1/volumes/%s/qos", params.Uuid)
	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetVolumeQoSFailed, err.Error())
	}
	var result model.GetVolumeQoSResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//取消云盘网卡限速
func DeleteVolumeQoS(params model.DeleteVolumeQoSRequest) models.Result {
	//DELETE zstack/v1/volumes/{uuid}/qos
	url := fmt.Sprintf("zstack/v1/volumes/%s/qos", params.Uuid)
	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteVolumeQoSFailed, err.Error())
	}
	var result model.DeleteVolumeQoSResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//获取云盘可被加载的云主机列表
func GetDataVolumeAttachableVm(params model.GetDataVolumeAttachableVmRequest) models.Result {
	//GET zstack/v1/volumes/{volumeUuid}/candidate-vm-instances
	url := fmt.Sprintf("zstack/v1/volumes/%s/candidate-vm-instances", params.VolumeUuid)
	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetDataVolumeAttachableVmFailed, err.Error())
	}
	var result model.GetDataVolumeAttachableVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//挂载云盘到云主机上
func AttachDataVolumeToVm(params model.AttachDataVolumeToVmRequest) models.Result {
	//POST zstack/v1/volumes/{volumeUuid}/vm-instances/{vmInstanceUuid}
	url := fmt.Sprintf("zstack/v1/volumes/%s/vm-instances/%s", params.VolumeUuid, params.VmInstanceUuid)
	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachDataVolumeToVmFailed, err.Error())
	}
	var result model.AttachDataVolumeToVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//从云主机上卸载云盘
func DetachDataVolumeFromVm(params model.DetachDataVolumeFromVmRequest) models.Result {
	//DELETE zstack/v1/volumes/{uuid}/vm-instances?vmUuid={vmUuid}
	url := fmt.Sprintf("zstack/v1/volumes/%s/vm-instances", params.Uuid)
	dataStr, err := request.DeleteUrlWithParams(url, params)
	if err != nil {
		return models.Error(errcode.DetachDataVolumeFromVmFailed, err.Error())
	}
	var result model.DetachDataVolumeFromVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//从云盘创建快照
func CreateVolumeSnapshot(params model.CreateVolumeSnapshotRequest) models.Result {
	//POST zstack/v1/volumes/{volumeUuid}/volume-snapshots
	url := fmt.Sprintf("zstack/v1/volumes/%s/volume-snapshots", params.VolumeUuid)
	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateVolumeSnapshotFailed, err.Error())
	}
	var result model.CreateVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//查询云盘快照
func QueryVolumeSnapshot(params model.QueryVolumeSnapshotRequest) models.Result {
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
		return models.Error(errcode.QueryVolumeSnapshotFailed, err.Error())
	}
	var result model.QueryVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//查询快照树
func QueryVolumeSnapshotTree(params model.QueryVolumeSnapshotTreeRequest) models.Result {
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
		return models.Error(errcode.QueryVolumeSnapshotTreeFailed, err.Error())
	}
	var result model.QueryVolumeSnapshotTreeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//更新云盘快照信息
func UpdateVolumeSnapshot(params model.UpdateVolumeSnapshotRequest) models.Result {
	//PUT zstack/v1/volume-snapshots/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volume-snapshots/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateVolumeSnapshotFailed, err.Error())
	}
	var result model.UpdateVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//删除云盘快照
func DeleteVolumeSnapshot(params model.DeleteVolumeSnapshotRequest) models.Result {
	//DELETE zstack/v1/volume-snapshots/{uuid}
	url := fmt.Sprintf("zstack/v1/volume-snapshots/%s", params.Uuid)
	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteVolumeSnapshotFailed, err.Error())
	}
	var result model.DeleteVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//将云盘回滚至指定快照
func RevertVolumeFromSnapshot(params model.RevertVolumeFromSnapshotRequest) models.Result {
	//PUT zstack/v1/volume-snapshots/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volume-snapshots/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.RevertVolumeFromSnapshotFailed, err.Error())
	}
	var result model.RevertVolumeFromSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//获取快照容量
func GetVolumeSnapshotSize(params model.GetVolumeSnapshotSizeRequest) models.Result {
	//PUT zstack/v1/volume-snapshots/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volume-snapshots/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.GetVolumeSnapshotSizeFailed, err.Error())
	}
	var result model.GetVolumeSnapshotSizeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}

//快照瘦身
func ShrinkVolumeSnapshot(params model.ShrinkVolumeSnapshotRequest) models.Result {
	//PUT zstack/v1/volume-snapshots/shrink/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volume-snapshots/shrink/%s/actions", params.Uuid)
	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ShrinkVolumeSnapshotFailed, err.Error())
	}
	var result model.ShrinkVolumeSnapshotResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终的结果:{}", result)

	return models.Success(result)
}
