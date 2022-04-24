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

//创建云盘规格
func CreateDiskOffering(params model.CreateDiskOfferingRequest) mgresult.Result {
	//POST zstack/v1/disk-offerings
	url := fmt.Sprintf("zstack/v1/disk-offerings")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateDiskOfferingFailed, err)
	}
	var result model.CreateDiskOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除云盘规格
func DeleteDiskOffering(params model.DeleteDiskOfferingRequest) mgresult.Result {
	//DELETE zstack/v1/disk-offerings/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/disk-offerings/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteDiskOfferingFailed, err)
	}
	var result model.DeleteDiskOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询云盘规格
func QueryDiskOffering(params model.QueryDiskOfferingRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryDiskOfferingFailed, err)
	}
	var result model.QueryDiskOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更改云盘规格的启用状态
func ChangeDiskOfferingState(params model.ChangeDiskOfferingStateRequest) mgresult.Result {
	//PUT zstack/v1/disk-offerings/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/disk-offerings/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeDiskOfferingStateFailed, err)
	}
	var result model.ChangeDiskOfferingStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新云盘规格
func UpdateDiskOffering(params model.UpdateDiskOfferingRequest) mgresult.Result {
	//PUT zstack/v1/disk-offerings/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/disk-offerings/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateDiskOfferingFailed, err)
	}
	var result model.UpdateDiskOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
