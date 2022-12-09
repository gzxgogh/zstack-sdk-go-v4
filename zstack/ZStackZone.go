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

//创建区域
func CreateZone(params model.CreateZoneRequest) models.Result {
	//POST zstack/v1/zones
	url := fmt.Sprintf("zstack/v1/zones")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateZoneFailed, err.Error())
	}
	var result model.CreateZoneResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除区域
func DeleteZone(params model.DeleteZoneRequest) models.Result {
	//DELETE zstack/v1/zones/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/zones/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteZoneFailed, err.Error())
	}
	var result model.DeleteZoneResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询区域
func QueryZone(params model.QueryZoneRequest) models.Result {
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
		return models.Error(errcode.QueryZoneFailed, err.Error())
	}
	var result model.QueryZoneResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新区域
func UpdateZone(params model.UpdateZoneRequest) models.Result {
	//PUT zstack/v1/zones/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zones/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateZoneFailed, err.Error())
	}
	var result model.UpdateZoneResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//改变区域的可用状态
func ChangeZoneState(params model.ChangeZoneStateRequest) models.Result {
	//PUT zstack/v1/zones/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/zones/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeZoneStateFailed, err.Error())
	}
	var result model.ChangeZoneStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
