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

//创建快照组
func CreateVolumeSnapshotGroup(params model.CreateVolumeSnapshotGroupRequest) models.Result {
	//POST zstack/v1/volume-snapshots/group
	url := fmt.Sprintf("zstack/v1/volume-snapshots/group")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateVolumeSnapshotGroupFailed, err.Error())
	}
	var result model.CreateVolumeSnapshotGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除快照组
func DeleteVolumeSnapshotGroup(params model.DeleteVolumeSnapshotGroupRequest) models.Result {
	//DELETE zstack/v1/volume-snapshots/group/{uuid}
	url := fmt.Sprintf("zstack/v1/volume-snapshots/group/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteVolumeSnapshotGroupFailed, err.Error())
	}
	var result model.DeleteVolumeSnapshotGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新快照组信息
func UpdateVolumeSnapshotGroup(params model.UpdateVolumeSnapshotGroupRequest) models.Result {
	//PUT zstack/v1/volume-snapshots/group/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volume-snapshots/group/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateVolumeSnapshotGroupFailed, err.Error())
	}
	var result model.UpdateVolumeSnapshotGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询快照组
func QueryVolumeSnapshotGroup(params model.QueryVolumeSnapshotGroupRequest) models.Result {
	//GET zstack/v1/volume-snapshots/group
	//GET zstack/v1/volume-snapshots/group/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/volume-snapshots/group")
	} else {
		url = fmt.Sprintf("zstack/v1/volume-snapshots/group/%s", params.UUID)

	}
	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVolumeSnapshotGroupFailed, err.Error())
	}
	var result model.QueryVolumeSnapshotGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//检查快照组可用性
func CheckVolumeSnapshotGroupAvailability(params model.CheckVolumeSnapshotGroupAvailabilityRequest) models.Result {
	//GET zstack/v1/volume-snapshots/groups/availabilities
	url := fmt.Sprintf("zstack/v1/volume-snapshots/groups/availabilities")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.CheckVolumeSnapshotGroupAvailabilityFailed, err.Error())
	}
	var result model.CheckVolumeSnapshotGroupAvailabilityResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//恢复快照组
func RevertVmFromSnapshotGroup(params model.RevertVmFromSnapshotGroupRequest) models.Result {
	//PUT zstack/v1/volume-snapshots/group/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/volume-snapshots/group/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.RevertVmFromSnapshotGroupFailed, err.Error())
	}
	var result model.RevertVmFromSnapshotGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//解绑快照组
func UngroupVolumeSnapshotGroup(params model.UngroupVolumeSnapshotGroupRequest) models.Result {
	//DELETE zstack/v1/volume-snapshots/ungroup/{uuid}
	url := fmt.Sprintf("zstack/v1/volume-snapshots/ungroup/%s", params.Uuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.UngroupVolumeSnapshotGroupFailed, err.Error())
	}
	var result model.UngroupVolumeSnapshotGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
