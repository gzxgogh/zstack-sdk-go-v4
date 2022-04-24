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

//创建云主机规格
func CreateInstanceOffering(params model.CreateInstanceOfferingRequest) mgresult.Result {
	//POST zstack/v1/instance-offerings
	url := fmt.Sprintf("zstack/v1/instance-offerings")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateInstanceOfferingFailed, err)
	}
	var result model.CreateInstanceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除云主机规格
func DeleteInstanceOffering(params model.DeleteInstanceOfferingRequest) mgresult.Result {
	//DELETE zstack/v1/instance-offerings/{uuid}
	url := fmt.Sprintf("zstack/v1/instance-offerings/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteInstanceOfferingFailed, err)
	}
	var result model.DeleteInstanceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询云主机规格
func QueryInstanceOffering(params model.QueryInstanceOfferingRequest) mgresult.Result {
	//GET zstack/v1/instance-offerings
	//GET zstack/v1/instance-offerings/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/instance-offerings")
	} else {
		url = fmt.Sprintf("zstack/v1/instance-offerings/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryInstanceOfferingFailed, err)
	}
	var result model.QueryInstanceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更改云主机规格
func ChangeInstanceOffering(params model.ChangeInstanceOfferingRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", params.VmInstanceUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeInstanceOfferingFailed, err)
	}
	var result model.ChangeInstanceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新云主机规格
func UpdateInstanceOffering(params model.UpdateInstanceOfferingRequest) mgresult.Result {
	//PUT zstack/v1/instance-offerings/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/instance-offerings/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateInstanceOfferingFailed, err)
	}
	var result model.UpdateInstanceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更改云主机规格的启用状态
func ChangeInstanceOfferingState(params model.ChangeInstanceOfferingStateRequest) mgresult.Result {
	//PUT zstack/v1/instance-offerings/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/instance-offerings/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeInstanceOfferingStateFailed, err)
	}
	var result model.ChangeInstanceOfferingStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
