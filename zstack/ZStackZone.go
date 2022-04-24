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

//创建区域
func CreateZone(params model.CreateZoneRequest) mgresult.Result {
	//POST zstack/v1/zones
	url := fmt.Sprintf("zstack/v1/zones")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateZoneFailed, err)
	}
	var result model.CreateZoneResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除区域
func DeleteZone(params model.DeleteZoneRequest) mgresult.Result {
	//DELETE zstack/v1/zones/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/zones/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteZoneFailed, err)
	}
	var result model.DeleteZoneResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询区域
func QueryZone(params model.QueryZoneRequest) mgresult.Result {
	//GET zstack/v1/zones
	//GET zstack/v1/zones/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/zones")
	} else {
		url = fmt.Sprintf("zstack/v1/zones/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryZoneFailed, err)
	}
	var result model.QueryZoneResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新区域
func UpdateZone(params model.UpdateZoneRequest) mgresult.Result {
	//PUT zstack/v1/zones/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zones/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateZoneFailed, err)
	}
	var result model.UpdateZoneResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//改变区域的可用状态
func ChangeZoneState(params model.ChangeZoneStateRequest) mgresult.Result {
	//PUT zstack/v1/zones/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zones/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeZoneStateFailed, err)
	}
	var result model.ChangeZoneStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
