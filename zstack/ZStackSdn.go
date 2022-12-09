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

//添加SDN控制器
func AddSdnController(params model.AddSdnControllerRequest) models.Result {
	//POST zstack/v1/sdn-controllers
	url := fmt.Sprintf("zstack/v1/sdn-controllers")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddSdnControllerFailed, err.Error())
	}
	var result model.AddSdnControllerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询SDN控制器
func QuerySdnController(params model.QuerySdnControllerRequest) models.Result {
	//GET zstack/v1/sdn-controllers
	//GET zstack/v1/sdn-controllers/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/sdn-controllers")
	} else {
		url = fmt.Sprintf("zstack/v1/sdn-controllers/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QuerySdnControllerFailed, err.Error())
	}
	var result model.QuerySdnControllerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除SDN控制器
func RemoveSdnController(params model.RemoveSdnControllerRequest) models.Result {
	//DELETE zstack/v1/sdn-controllers/{uuid}
	url := fmt.Sprintf("zstack/v1/sdn-controllers/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveSdnControllerFailed, err.Error())
	}
	var result model.RemoveSdnControllerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新SDN控制器
func UpdateSdnController(params model.UpdateSdnControllerRequest) models.Result {
	//PUT zstack/v1/sdn-controllers/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/sdn-controllers/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateSdnControllerFailed, err.Error())
	}
	var result model.UpdateSdnControllerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
