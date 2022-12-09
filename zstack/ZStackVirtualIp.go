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

//创建虚拟IP
func CreateVip(params model.CreateVipRequest) models.Result {
	//POST zstack/v1/vips
	url := fmt.Sprintf("zstack/v1/vips")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateVipFailed, err.Error())
	}
	var result model.CreateVipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除虚拟IP
func DeleteVip(params model.DeleteVipRequest) models.Result {
	//DELETE zstack/v1/vips/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/vips/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteVipFailed, err.Error())
	}
	var result model.DeleteVipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询虚拟IP
func QueryVip(params model.QueryVipRequest) models.Result {
	//GET zstack/v1/vips
	//GET zstack/v1/vips/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vips")
	} else {
		url = fmt.Sprintf("zstack/v1/vips/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVipFailed, err.Error())
	}
	var result model.QueryVipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新虚拟IP
func UpdateVip(params model.UpdateVipRequest) models.Result {
	//PUT zstack/v1/vips/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vips/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateVipFailed, err.Error())
	}
	var result model.UpdateVipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更改虚拟IP启用状态
func ChangeVipState(params model.ChangeVipStateRequest) models.Result {
	//PUT zstack/v1/vips/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vips/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeVipStateFailed, err.Error())
	}
	var result model.ChangeVipStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取虚拟IP所有业务端口列表
func GetVipUsedPorts(params model.GetVipUsedPortsRequest) models.Result {
	//GET zstack/v1/vips/{uuid}/usedports
	url := fmt.Sprintf("zstack/v1/vips/%s/usedports", params.UUID)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetVipUsedPortsFailed, err.Error())
	}
	var result model.GetVipUsedPortsResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//设置虚拟IP限速
func SetVipQos(params model.SetVipQosRequest) models.Result {
	//PUT zstack/v1/vips/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vips/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.SetVipQosFailed, err.Error())
	}
	var result model.SetVipQosResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取虚拟IP限速
func GetVipQos(params model.GetVipQosRequest) models.Result {
	//GET zstack/v1/vip/{uuid}/vip-qos
	url := fmt.Sprintf("zstack/v1/vip/%s/vip-qos", params.UUID)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetVipQosFailed, err.Error())
	}
	var result model.GetVipQosResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//取消虚拟IP限速
func DeleteVipQos(params model.DeleteVipQosRequest) models.Result {
	//DELETE zstack/v1/vips/{uuid}/vip-qos?port={port}
	url := fmt.Sprintf("zstack/v1/vips/%s/vip-qos?port=%s", params.UUID, params.Port)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteVipQosFailed, err.Error())
	}
	var result model.DeleteVipQosResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
