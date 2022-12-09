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

//创建云盘规格
func CreateDiskOffering(params model.CreateDiskOfferingRequest) models.Result {
	//POST zstack/v1/disk-offerings
	url := fmt.Sprintf("zstack/v1/disk-offerings")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateDiskOfferingFailed, err.Error())
	}
	var result model.CreateDiskOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除云盘规格
func DeleteDiskOffering(params model.DeleteDiskOfferingRequest) models.Result {
	//DELETE zstack/v1/disk-offerings/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/disk-offerings/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteDiskOfferingFailed, err.Error())
	}
	var result model.DeleteDiskOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询云盘规格
func QueryDiskOffering(params model.QueryDiskOfferingRequest) models.Result {
	//GET zstack/v1/disk-offerings
	//GET zstack/v1/disk-offerings/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/disk-offerings")
	} else {
		url = fmt.Sprintf("zstack/v1/disk-offerings/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryDiskOfferingFailed, err.Error())
	}
	var result model.QueryDiskOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更改云盘规格的启用状态
func ChangeDiskOfferingState(params model.ChangeDiskOfferingStateRequest) models.Result {
	//PUT zstack/v1/disk-offerings/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/disk-offerings/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeDiskOfferingStateFailed, err.Error())
	}
	var result model.ChangeDiskOfferingStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新云盘规格
func UpdateDiskOffering(params model.UpdateDiskOfferingRequest) models.Result {
	//PUT zstack/v1/disk-offerings/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/disk-offerings/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateDiskOfferingFailed, err.Error())
	}
	var result model.UpdateDiskOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
