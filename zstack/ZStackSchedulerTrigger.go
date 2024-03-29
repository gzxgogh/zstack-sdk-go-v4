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

//创建定时器
func CreateSchedulerTrigger(params model.CreateSchedulerTriggerRequest) models.Result {
	//POST zstack/v1/scheduler/triggers
	url := fmt.Sprintf("zstack/v1/scheduler/triggers")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateSchedulerTriggerFailed, err.Error())
	}
	var respResult model.CreateSchedulerTriggerResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//删除定时器
func DeleteSchedulerTrigger(params model.DeleteResourcePriceRequest) models.Result {
	//DELETE zstack/v1/scheduler/triggers/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/scheduler/triggers/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteSchedulerTriggerFailed, err.Error())
	}
	var respResult model.DeleteSchedulerTriggerResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询定时器
func QuerySchedulerTrigger(params model.QuerySchedulerTriggerRequest) models.Result {
	//GET zstack/v1/scheduler/triggers
	//GET zstack/v1/scheduler/triggers/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/scheduler/triggers")
	} else {
		url = fmt.Sprintf("zstack/v1/scheduler/triggers/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QuerySchedulerTriggerFailed, err.Error())
	}
	var respResult model.QuerySchedulerTriggerResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//更新定时器
func UpdateSchedulerTrigger(params model.UpdateSchedulerTriggerRequest) models.Result {
	//PUT zstack/v1/scheduler/triggers/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/scheduler/triggers/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateSchedulerTriggerFailed, err.Error())
	}
	var respResult model.UpdateSchedulerTriggerResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//添加定时任务到定时器
func AddSchedulerJobToSchedulerTrigger(params model.AddSchedulerJobToSchedulerTriggerRequest) models.Result {
	//POST zstack/v1/scheduler/jobs/{schedulerJobUuid}/scheduler/triggers/{schedulerTriggerUuid}
	url := fmt.Sprintf("zstack/v1/scheduler/jobs/%s/scheduler/triggers/%s", params.SchedulerJobUuid, params.SchedulerTriggerUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddSchedulerJobToSchedulerTriggerFailed, err.Error())
	}
	var respResult model.AddSchedulerJobToSchedulerTriggerResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//从定时器移除定时任务
func RemoveSchedulerJobFromSchedulerTrigger(params model.RemoveSchedulerJobFromSchedulerTriggerRequest) models.Result {
	//DELETE zstack/v1/scheduler/jobs/{schedulerJobUuid}/scheduler/triggers/{schedulerTriggerUuid}
	url := fmt.Sprintf("zstack/v1/scheduler/jobs/%s/scheduler/triggers/%s", params.SchedulerJobUuid, params.SchedulerTriggerUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveSchedulerJobFromSchedulerTriggerFailed, err.Error())
	}
	var respResult model.RemoveSchedulerJobFromSchedulerTriggerResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//创建定时任务
func CreateSchedulerJob(params model.CreateSchedulerJobRequest) models.Result {
	//POST zstack/v1/scheduler/jobs
	url := fmt.Sprintf("zstack/v1/scheduler/jobs")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateSchedulerJobFailed, err.Error())
	}
	var respResult model.CreateSchedulerTriggerResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//删除定时任务
func DeleteSchedulerJob(params model.DeleteSchedulerJobRequest) models.Result {
	//DELETE zstack/v1/scheduler/jobs/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/scheduler/jobs/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteSchedulerJobFailed, err.Error())
	}
	var respResult model.DeleteSchedulerJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询定时任务
func QuerySchedulerJob(params model.QuerySchedulerJobRequest) models.Result {
	//GET zstack/v1/scheduler/jobs
	//GET zstack/v1/scheduler/jobs/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/scheduler/jobs")
	} else {
		url = fmt.Sprintf("zstack/v1/scheduler/jobs/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QuerySchedulerJobFailed, err.Error())
	}
	var respResult model.QuerySchedulerTriggerResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//更新定时任务
func UpdateSchedulerJob(params model.UpdateSchedulerJobRequest) models.Result {
	//PUT zstack/v1/scheduler/jobs/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/scheduler/jobs/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateSchedulerJobFailed, err.Error())
	}
	var respResult model.UpdateSchedulerJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取未挂载定时器的任务
func GetNoTriggerSchedulerJobs(params model.GetNoTriggerSchedulerJobsRequest) models.Result {
	//GET zstack/v1/scheduler/jobs/candidates
	url := fmt.Sprintf("zstack/v1/scheduler/jobs/candidates")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetNoTriggerSchedulerJobsFailed, err.Error())
	}
	var respResult model.GetNoTriggerSchedulerJobsResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//改变定时任务状态
func ChangeSchedulerState(params model.ChangeSchedulerStateRequest) models.Result {
	//PUT zstack/v1/schedulers/{uuid}
	url := fmt.Sprintf("zstack/v1/schedulers/%s", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeSchedulerStateFailed, err.Error())
	}
	var respResult model.ChangeSchedulerStateResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}
