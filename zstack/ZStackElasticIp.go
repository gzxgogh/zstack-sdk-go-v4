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

//创建弹性IP
func CreateEip(params model.CreateEipRequest) mgresult.Result {
	//POST zstack/v1/eips
	url := fmt.Sprintf("zstack/v1/eips")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateEipFailed, err)
	}
	var result model.CreateEipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除弹性IP
func DeleteEip(params model.DeleteEipRequest) mgresult.Result {
	//DELETE zstack/v1/eips/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/eips/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteEipFailed, err)
	}
	var result model.DeleteEipRequest
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询弹性IP
func QueryEip(params model.QueryEipRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryEipFailed, err)
	}
	var result model.QueryEipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新弹性IP
func UpdateEip(params model.UpdateEipRequest) mgresult.Result {
	//PUT zstack/v1/eips/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/eips/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateEipFailed, err)
	}
	var result model.UpdateEipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更改虚拟IP启用状态
func ChangeEipState(params model.ChangeEipStateRequest) mgresult.Result {
	//PUT zstack/v1/eips/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/eips/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeEipStateFailed, err)
	}
	var result model.ChangeEipStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取可绑定指定弹性IP的云主机网卡
func GetEipAttachableVmNics(params model.GetEipAttachableVmNicsRequest) mgresult.Result {
	//GET zstack/v1/eips/{eipUuid}/vm-instances/candidate-nics
	url := fmt.Sprintf("zstack/v1/eips/%s/vm-instances/candidate-nics", params.EipUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetEipAttachableVmNicsFailed, err)
	}
	var result model.GetEipAttachableVmNicsResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//绑定弹性IP
func AttachEip(params model.AttachEipRequest) mgresult.Result {
	//POST zstack/v1/eips/{eipUuid}/vm-instances/nics/{vmNicUuid}
	url := fmt.Sprintf("zstack/v1/eips/%s/vm-instances/nics/%s", params.EipUuid, params.VmNicUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachEipFailed, err)
	}
	var result model.AttachEipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//解绑弹性IP
func DetachEip(params model.DetachEipRequest) mgresult.Result {
	//DELETE zstack/v1/eips/{uuid}/vm-instances/nics
	url := fmt.Sprintf("zstack/v1/eips/%s/vm-instances/nics", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachEipFailed, err)
	}
	var result model.DetachEipResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
