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

//查询管理节点
func QueryManagementNode(params model.QueryManagementNodeRequest) mgresult.Result {
	//GET zstack/v1/management-nodes
	//GET zstack/v1/management-nodes/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/management-nodes")
	} else {
		url = fmt.Sprintf("zstack/v1/management-nodes/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryManagementNodeFailed, err)
	}
	var respResult model.QueryVirtualRouterOfferingResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//获取当前版本
func GetVersion(params model.GetVersionRequest) mgresult.Result {
	//PUT zstack/v1/management-nodes/actions
	url := fmt.Sprintf("zstack/v1/management-nodes/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVersionFailed, err)
	}
	var respResult model.GetVersionResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//获取当前时间
func GetCurrentTime(params model.GetCurrentTimeRequest) mgresult.Result {
	//PUT zstack/v1/management-nodes/actions
	url := fmt.Sprintf("zstack/v1/management-nodes/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCurrentTimeFailed, err)
	}
	var respResult model.GetCurrentTimeResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//获取管理节点当前的时区信息
func GetPlatformTimeZone(params model.GetPlatformTimeZoneRequest) mgresult.Result {
	//PUT zstack/v1/management-nodes/actions
	url := fmt.Sprintf("zstack/v1/management-nodes/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetPlatformTimeZoneFailed, err)
	}
	var respResult model.GetPlatformTimeZoneResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//检查管理节点是否能正常开始工作
func IsReadyToGo(params model.IsReadyToGoRequest) mgresult.Result {
	//GET zstack/v1/management-nodes/ready
	url := fmt.Sprintf("zstack/v1/management-nodes/ready")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.IsReadyToGoFailed, err)
	}
	var respResult model.IsReadyToGoResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//获取管理节点操作系统信息
func GetManagementNodeOS(params model.GetManagementNodeOSRequest) mgresult.Result {
	//PUT zstack/v1/management/actions
	url := fmt.Sprintf("zstack/v1/management/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetManagementNodeOSFailed, err)
	}
	var respResult model.GetManagementNodeOSResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//获取管理节点CPU架构类型
func GetManagementNodeArch(params model.GetManagementNodeArchRequest) mgresult.Result {
	//PUT zstack/v1/management/actions
	url := fmt.Sprintf("zstack/v1/management/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetManagementNodeArchFailed, err)
	}
	var respResult model.GetManagementNodeArchResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//创建系统标签
func CreateSystemTag(params model.CreateSystemTagRequest) mgresult.Result {
	//POST zstack/v1/system-tags
	url := fmt.Sprintf("zstack/v1/system-tags")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateSystemTagFailed, err)
	}
	var respResult model.CreateSystemTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//查询系统标签
func QuerySystemTag(params model.QuerySystemTagRequest) mgresult.Result {
	//GET zstack/v1/system-tags
	//GET zstack/v1/system-tags/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/system-tags")
	} else {
		url = fmt.Sprintf("zstack/v1/system-tags/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QuerySystemTagFailed, err)
	}
	var respResult model.QuerySystemTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//更新系统标签
func UpdateSystemTag(params model.UpdateSystemTagRequest) mgresult.Result {
	//PUT zstack/v1/system-tags/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/system-tags/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateSystemTagFailed, err)
	}
	var respResult model.UpdateSystemTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//创建用户标签
func CreateUserTag(params model.CreateUserTagRequest) mgresult.Result {
	//POST zstack/v1/user-tags
	url := fmt.Sprintf("zstack/v1/user-tags")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateUserTagRequestFailed, err)
	}
	var respResult model.CreateUserTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//查询用户标签
func QueryUserTag(params model.QueryUserTagRequest) mgresult.Result {
	//GET zstack/v1/user-tags
	//GET zstack/v1/user-tags/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/user-tags")
	} else {
		url = fmt.Sprintf("zstack/v1/user-tags/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryUserTagFailed, err)
	}
	var respResult model.QueryUserTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//删除标签
func DeleteTag(params model.DeleteTagRequest) mgresult.Result {
	//DELETE zstack/v1/tags/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/tags/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteTagFailed, err)
	}
	var respResult model.DeleteTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//创建资源标签
func CreateTag(params model.CreateTagRequest) mgresult.Result {
	//POST zstack/v1/tags
	url := fmt.Sprintf("zstack/v1/tags")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateTagFailed, err)
	}
	var respResult model.CreateTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//查询资源标签
func QueryTag(params model.QueryTagRequest) mgresult.Result {
	//GET zstack/v1/tags
	//GET zstack/v1/tags/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/tags")
	} else {
		url = fmt.Sprintf("zstack/v1/tags/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryTagFailed, err)
	}
	var respResult model.QueryTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//更新资源标签
func UpdateTag(params model.UpdateTagRequest) mgresult.Result {
	//PUT zstack/v1/tags/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/tags/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateTagFailed, err)
	}
	var respResult model.UpdateTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//加载资源标签
func AttachTagToResources(params model.AttachTagToResourcesRequest) mgresult.Result {
	//POST zstack/v1/tags/{tagUuid}/resources
	url := fmt.Sprintf("zstack/v1/tags/%s/resources", params.TagUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachTagToResourcesFailed, err)
	}
	var respResult model.AttachTagToResourcesResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//卸载资源标签
func DetachTagFromResources(params model.DetachTagFromResourcesRequest) mgresult.Result {
	//DELETE zstack/v1/tags/{tagUuid}/resources?resourceUuids={resourceUuids}
	url := fmt.Sprintf("zstack/v1/tags/%s/resources?resourceUuids=%s", params.TagUuid, params.ResourceUuids)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachTagFromResourcesFailed, err)
	}
	var respResult model.DetachTagFromResourcesResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//获取任务进度
func GetTaskProgress(params model.GetTaskProgressRequest) mgresult.Result {
	//GET zstack/v1/task-progresses/{apiId}
	url := fmt.Sprintf("zstack/v1/task-progresses/%s", params.ApiId)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetTaskProgressFailed, err)
	}
	var respResult model.GetTaskProgressResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//获取CPU和内存容量
func GetCpuMemoryCapacity(params model.GetCpuMemoryCapacityRequest) mgresult.Result {
	//GET zstack/v1/hosts/capacities/cpu-memory
	url := fmt.Sprintf("zstack/v1/hosts/capacities/cpu-memory")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCpuMemoryCapacityFailed, err)
	}
	var respResult model.GetCpuMemoryCapacityResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//触发垃圾回收任务
func TriggerGCJob(params model.TriggerGCJobRequest) mgresult.Result {
	//PUT zstack/v1/gc-jobs/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/gc-jobs/%s/actions", params.UUID)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.TriggerGCJobFailed, err)
	}
	var respResult model.TriggerGCJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//删除垃圾回收任务
func DeleteGCJob(params model.DeleteGCJobRequest) mgresult.Result {
	//DELETE zstack/v1/gc-jobs/{uuid}
	url := fmt.Sprintf("zstack/v1/gc-jobs/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteGCJobFailed, err)
	}
	var respResult model.DeleteGCJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//查询垃圾回收任务
func QueryGCJob(params model.QueryGCJobRequest) mgresult.Result {
	//GET zstack/v1/gc-jobs
	//GET zstack/v1/gc-jobs/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/gc-jobs")
	} else {
		url = fmt.Sprintf("zstack/v1/gc-jobs/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryGCJobFailed, err)
	}
	var respResult model.QueryGCJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//获取许可证信息
func GetLicenseInfo(params model.GetLicenseInfoRequest) mgresult.Result {
	//GET zstack/v1/licenses
	url := fmt.Sprintf("zstack/v1/licenses")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetLicenseInfoFailed, err)
	}
	var respResult model.GetLicenseInfoResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//获取许可证历史授权信息
func GetLicenseRecords(params model.GetLicenseRecordsRequest) mgresult.Result {
	//GET zstack/v1/licenses
	url := fmt.Sprintf("zstack/v1/licenses")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetLicenseRecordsFailed, err)
	}
	var respResult model.GetLicenseRecordsResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//获取许可证容量
func GetLicenseCapabilities(params model.GetLicenseCapabilitiesRequest) mgresult.Result {
	//GET zstack/v1/licenses/capabilities
	url := fmt.Sprintf("zstack/v1/licenses/capabilities")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetLicenseCapabilitiesFailed, err)
	}
	var respResult model.GetLicenseCapabilitiesResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//删除许可证文件
func DeleteLicense(params model.DeleteLicenseRequest) mgresult.Result {
	//DELETE /v1/licenses/mn/{managementNodeUuid}/actions?uuid={uuid}
	url := fmt.Sprintf("zstack/v1/licenses/mn/%s/actions?uuid=%s", params.ManagementNodeUuid, params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteLicenseFailed, err)
	}
	var respResult model.DeleteLicenseResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//重新加载许可证
func ReloadLicense(params model.ReloadLicenseRequest) mgresult.Result {
	//PUT zstack/v1/licenses/actions
	url := fmt.Sprintf("zstack/v1/licenses/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ReloadLicenseFailed, err)
	}
	var respResult model.ReloadLicenseResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//更新许可证信息
func UpdateLicense(params model.UpdateLicenseRequest) mgresult.Result {
	//PUT zstack/v1/licenses/mn/{managementNodeUuid}/actions
	url := fmt.Sprintf("zstack/v1/licenses/mn/%s/actions", params.ManagementNodeUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateLicenseFailed, err)
	}
	var respResult model.UpdateLicenseResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//提交长时任务
func SubmitLongJob(params model.SubmitLongJobRequest) mgresult.Result {
	//POST zstack/v1/longjobs
	url := fmt.Sprintf("zstack/v1/longjobs")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SubmitLongJobFailed, err)
	}
	var respResult model.SubmitLongJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//更新长时任务
func UpdateLongJob(params model.UpdateLongJobRequest) mgresult.Result {
	//PUT zstack/v1/longjobs/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/longjobs/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateLongJobFailed, err)
	}
	var respResult model.UpdateLongJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//删除长时任务
func DeleteLongJob(params model.DeleteLongJobRequest) mgresult.Result {
	//DELETE zstack/v1/longjobs/{uuid}
	url := fmt.Sprintf("zstack/v1/longjobs/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteLongJobFailed, err)
	}
	var respResult model.DeleteLongJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//查询长时任务
func QueryLongJob(params model.QueryLongJobRequest) mgresult.Result {
	//GET zstack/v1/longjobs
	//GET zstack/v1/longjobs/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/longjobs")
	} else {
		url = fmt.Sprintf("zstack/v1/longjobs/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryLongJobFailed, err)
	}
	var respResult model.QueryLongJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//查看系统错误码
func GetElaborations(params model.GetElaborationsRequest) mgresult.Result {
	//GET zstack/v1/errorcode/elaborations
	url := fmt.Sprintf("zstack/v1/errorcode/elaborations")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetElaborationsFailed, err)
	}
	var respResult model.GetElaborationsResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}

//重新生成索引
func RefreshSearchIndexes(params model.RefreshSearchIndexesRequest) mgresult.Result {
	//GET zstack/v1/errorcode/elaborations
	url := fmt.Sprintf("zstack/v1/errorcode/elaborations")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RefreshSearchIndexesFailed, err)
	}
	var respResult model.RefreshSearchIndexesResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return mgresult.Success(respResult)
}
