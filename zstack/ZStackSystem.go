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

//查询管理节点
func QueryManagementNode(params model.QueryManagementNodeRequest) models.Result {
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
		return models.Error(errcode.QueryManagementNodeFailed, err.Error())
	}
	var respResult model.QueryVirtualRouterOfferingResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取当前版本
func GetVersion(params model.GetVersionRequest) models.Result {
	//PUT zstack/v1/management-nodes/actions
	url := fmt.Sprintf("zstack/v1/management-nodes/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.GetVersionFailed, err.Error())
	}
	var respResult model.GetVersionResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取当前时间
func GetCurrentTime(params model.GetCurrentTimeRequest) models.Result {
	//PUT zstack/v1/management-nodes/actions
	url := fmt.Sprintf("zstack/v1/management-nodes/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.GetCurrentTimeFailed, err.Error())
	}
	var respResult model.GetCurrentTimeResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取管理节点当前的时区信息
func GetPlatformTimeZone(params model.GetPlatformTimeZoneRequest) models.Result {
	//PUT zstack/v1/management-nodes/actions
	url := fmt.Sprintf("zstack/v1/management-nodes/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.GetPlatformTimeZoneFailed, err.Error())
	}
	var respResult model.GetPlatformTimeZoneResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//检查管理节点是否能正常开始工作
func IsReadyToGo(params model.IsReadyToGoRequest) models.Result {
	//GET zstack/v1/management-nodes/ready
	url := fmt.Sprintf("zstack/v1/management-nodes/ready")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.IsReadyToGoFailed, err.Error())
	}
	var respResult model.IsReadyToGoResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取管理节点操作系统信息
func GetManagementNodeOS(params model.GetManagementNodeOSRequest) models.Result {
	//PUT zstack/v1/management/actions
	url := fmt.Sprintf("zstack/v1/management/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.GetManagementNodeOSFailed, err.Error())
	}
	var respResult model.GetManagementNodeOSResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取管理节点CPU架构类型
func GetManagementNodeArch(params model.GetManagementNodeArchRequest) models.Result {
	//PUT zstack/v1/management/actions
	url := fmt.Sprintf("zstack/v1/management/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.GetManagementNodeArchFailed, err.Error())
	}
	var respResult model.GetManagementNodeArchResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//创建系统标签
func CreateSystemTag(params model.CreateSystemTagRequest) models.Result {
	//POST zstack/v1/system-tags
	url := fmt.Sprintf("zstack/v1/system-tags")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateSystemTagFailed, err.Error())
	}
	var respResult model.CreateSystemTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询系统标签
func QuerySystemTag(params model.QuerySystemTagRequest) models.Result {
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
		return models.Error(errcode.QuerySystemTagFailed, err.Error())
	}
	var respResult model.QuerySystemTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//更新系统标签
func UpdateSystemTag(params model.UpdateSystemTagRequest) models.Result {
	//PUT zstack/v1/system-tags/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/system-tags/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateSystemTagFailed, err.Error())
	}
	var respResult model.UpdateSystemTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//创建用户标签
func CreateUserTag(params model.CreateUserTagRequest) models.Result {
	//POST zstack/v1/user-tags
	url := fmt.Sprintf("zstack/v1/user-tags")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateUserTagRequestFailed, err.Error())
	}
	var respResult model.CreateUserTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询用户标签
func QueryUserTag(params model.QueryUserTagRequest) models.Result {
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
		return models.Error(errcode.QueryUserTagFailed, err.Error())
	}
	var respResult model.QueryUserTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//删除标签
func DeleteTag(params model.DeleteTagRequest) models.Result {
	//DELETE zstack/v1/tags/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/tags/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteTagFailed, err.Error())
	}
	var respResult model.DeleteTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//创建资源标签
func CreateTag(params model.CreateTagRequest) models.Result {
	//POST zstack/v1/tags
	url := fmt.Sprintf("zstack/v1/tags")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateTagFailed, err.Error())
	}
	var respResult model.CreateTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询资源标签
func QueryTag(params model.QueryTagRequest) models.Result {
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
		return models.Error(errcode.QueryTagFailed, err.Error())
	}
	var respResult model.QueryTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//更新资源标签
func UpdateTag(params model.UpdateTagRequest) models.Result {
	//PUT zstack/v1/tags/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/tags/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateTagFailed, err.Error())
	}
	var respResult model.UpdateTagResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//加载资源标签
func AttachTagToResources(params model.AttachTagToResourcesRequest) models.Result {
	//POST zstack/v1/tags/{tagUuid}/resources
	url := fmt.Sprintf("zstack/v1/tags/%s/resources", params.TagUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachTagToResourcesFailed, err.Error())
	}
	var respResult model.AttachTagToResourcesResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//卸载资源标签
func DetachTagFromResources(params model.DetachTagFromResourcesRequest) models.Result {
	//DELETE zstack/v1/tags/{tagUuid}/resources?resourceUuids={resourceUuids}
	url := fmt.Sprintf("zstack/v1/tags/%s/resources?resourceUuids=%s", params.TagUuid, params.ResourceUuids)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachTagFromResourcesFailed, err.Error())
	}
	var respResult model.DetachTagFromResourcesResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取任务进度
func GetTaskProgress(params model.GetTaskProgressRequest) models.Result {
	//GET zstack/v1/task-progresses/{apiId}
	url := fmt.Sprintf("zstack/v1/task-progresses/%s", params.ApiId)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetTaskProgressFailed, err.Error())
	}
	var respResult model.GetTaskProgressResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取CPU和内存容量
func GetCpuMemoryCapacity(params model.GetCpuMemoryCapacityRequest) models.Result {
	//GET zstack/v1/hosts/capacities/cpu-memory
	url := fmt.Sprintf("zstack/v1/hosts/capacities/cpu-memory")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetCpuMemoryCapacityFailed, err.Error())
	}
	var respResult model.GetCpuMemoryCapacityResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//触发垃圾回收任务
func TriggerGCJob(params model.TriggerGCJobRequest) models.Result {
	//PUT zstack/v1/gc-jobs/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/gc-jobs/%s/actions", params.UUID)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.TriggerGCJobFailed, err.Error())
	}
	var respResult model.TriggerGCJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//删除垃圾回收任务
func DeleteGCJob(params model.DeleteGCJobRequest) models.Result {
	//DELETE zstack/v1/gc-jobs/{uuid}
	url := fmt.Sprintf("zstack/v1/gc-jobs/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteGCJobFailed, err.Error())
	}
	var respResult model.DeleteGCJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询垃圾回收任务
func QueryGCJob(params model.QueryGCJobRequest) models.Result {
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
		return models.Error(errcode.QueryGCJobFailed, err.Error())
	}
	var respResult model.QueryGCJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取许可证信息
func GetLicenseInfo(params model.GetLicenseInfoRequest) models.Result {
	//GET zstack/v1/licenses
	url := fmt.Sprintf("zstack/v1/licenses")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetLicenseInfoFailed, err.Error())
	}
	var respResult model.GetLicenseInfoResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取许可证历史授权信息
func GetLicenseRecords(params model.GetLicenseRecordsRequest) models.Result {
	//GET zstack/v1/licenses
	url := fmt.Sprintf("zstack/v1/licenses")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetLicenseRecordsFailed, err.Error())
	}
	var respResult model.GetLicenseRecordsResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//获取许可证容量
func GetLicenseCapabilities(params model.GetLicenseCapabilitiesRequest) models.Result {
	//GET zstack/v1/licenses/capabilities
	url := fmt.Sprintf("zstack/v1/licenses/capabilities")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetLicenseCapabilitiesFailed, err.Error())
	}
	var respResult model.GetLicenseCapabilitiesResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//删除许可证文件
func DeleteLicense(params model.DeleteLicenseRequest) models.Result {
	//DELETE /v1/licenses/mn/{managementNodeUuid}/actions?uuid={uuid}
	url := fmt.Sprintf("zstack/v1/licenses/mn/%s/actions?uuid=%s", params.ManagementNodeUuid, params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteLicenseFailed, err.Error())
	}
	var respResult model.DeleteLicenseResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//重新加载许可证
func ReloadLicense(params model.ReloadLicenseRequest) models.Result {
	//PUT zstack/v1/licenses/actions
	url := fmt.Sprintf("zstack/v1/licenses/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ReloadLicenseFailed, err.Error())
	}
	var respResult model.ReloadLicenseResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//更新许可证信息
func UpdateLicense(params model.UpdateLicenseRequest) models.Result {
	//PUT zstack/v1/licenses/mn/{managementNodeUuid}/actions
	url := fmt.Sprintf("zstack/v1/licenses/mn/%s/actions", params.ManagementNodeUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateLicenseFailed, err.Error())
	}
	var respResult model.UpdateLicenseResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//提交长时任务
func SubmitLongJob(params model.SubmitLongJobRequest) models.Result {
	//POST zstack/v1/longjobs
	url := fmt.Sprintf("zstack/v1/longjobs")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.SubmitLongJobFailed, err.Error())
	}
	var respResult model.SubmitLongJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//更新长时任务
func UpdateLongJob(params model.UpdateLongJobRequest) models.Result {
	//PUT zstack/v1/longjobs/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/longjobs/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateLongJobFailed, err.Error())
	}
	var respResult model.UpdateLongJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//删除长时任务
func DeleteLongJob(params model.DeleteLongJobRequest) models.Result {
	//DELETE zstack/v1/longjobs/{uuid}
	url := fmt.Sprintf("zstack/v1/longjobs/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteLongJobFailed, err.Error())
	}
	var respResult model.DeleteLongJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查询长时任务
func QueryLongJob(params model.QueryLongJobRequest) models.Result {
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
		return models.Error(errcode.QueryLongJobFailed, err.Error())
	}
	var respResult model.QueryLongJobResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//查看系统错误码
func GetElaborations(params model.GetElaborationsRequest) models.Result {
	//GET zstack/v1/errorcode/elaborations
	url := fmt.Sprintf("zstack/v1/errorcode/elaborations")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetElaborationsFailed, err.Error())
	}
	var respResult model.GetElaborationsResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}

//重新生成索引
func RefreshSearchIndexes(params model.RefreshSearchIndexesRequest) models.Result {
	//GET zstack/v1/errorcode/elaborations
	url := fmt.Sprintf("zstack/v1/errorcode/elaborations")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.RefreshSearchIndexesFailed, err.Error())
	}
	var respResult model.RefreshSearchIndexesResponse
	utils.FromJSON(dataStr, &respResult)
	logs.Debug("最终结果:{}", respResult)

	return models.Success(respResult)
}
