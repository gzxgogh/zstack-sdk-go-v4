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

//创建弹性IP
func CreateEip(params model.CreateEipRequest) models.Result {
	//POST zstack/v1/eips
	url := fmt.Sprintf("zstack/v1/eips")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateEipFailed, err.Error())
	}
	var result model.CreateEipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除弹性IP
func DeleteEip(params model.DeleteEipRequest) models.Result {
	//DELETE zstack/v1/eips/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/eips/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteEipFailed, err.Error())
	}
	var result model.DeleteEipRequest
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询弹性IP
func QueryEip(params model.QueryEipRequest) models.Result {
	//GET zstack/v1/eips
	//GET zstack/v1/eips/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/eips")
	} else {
		url = fmt.Sprintf("zstack/v1/eips/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryEipFailed, err.Error())
	}
	var result model.QueryEipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新弹性IP
func UpdateEip(params model.UpdateEipRequest) models.Result {
	//PUT zstack/v1/eips/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/eips/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateEipFailed, err.Error())
	}
	var result model.UpdateEipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更改虚拟IP启用状态
func ChangeEipState(params model.ChangeEipStateRequest) models.Result {
	//PUT zstack/v1/eips/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/eips/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeEipStateFailed, err.Error())
	}
	var result model.ChangeEipStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取可绑定指定弹性IP的云主机网卡
func GetEipAttachableVmNics(params model.GetEipAttachableVmNicsRequest) models.Result {
	//GET zstack/v1/eips/{eipUuid}/vm-instances/candidate-nics
	url := fmt.Sprintf("zstack/v1/eips/%s/vm-instances/candidate-nics", params.EipUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetEipAttachableVmNicsFailed, err.Error())
	}
	var result model.GetEipAttachableVmNicsResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//绑定弹性IP
func AttachEip(params model.AttachEipRequest) models.Result {
	//POST zstack/v1/eips/{eipUuid}/vm-instances/nics/{vmNicUuid}
	url := fmt.Sprintf("zstack/v1/eips/%s/vm-instances/nics/%s", params.EipUuid, params.VmNicUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachEipFailed, err.Error())
	}
	var result model.AttachEipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//解绑弹性IP
func DetachEip(params model.DetachEipRequest) models.Result {
	//DELETE zstack/v1/eips/{uuid}/vm-instances/nics
	url := fmt.Sprintf("zstack/v1/eips/%s/vm-instances/nics", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachEipFailed, err.Error())
	}
	var result model.DetachEipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
