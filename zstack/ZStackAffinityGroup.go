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

//创建亲和组
func CreateAffinityGroup(params model.CreateAffinityGroupRequest) mgresult.Result {
	//POST zstack/v1/affinity-groups
	url := fmt.Sprintf("zstack/v1/affinity-groups")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateAffinityGroupFailed, err)
	}
	var result model.CreateAffinityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除亲和组
func DeleteAffinityGroup(params model.DeleteAffinityGroupRequest) mgresult.Result {
	//DELETE zstack/v1/affinity-groups/{uuid}
	url := fmt.Sprintf("zstack/v1/affinity-groups/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteAffinityGroupFailed, err)
	}
	var result model.DeleteAffinityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询亲和组
func QueryAffinityGroup(params model.QueryAffinityGroupRequest) mgresult.Result {
	//GET zstack/v1/affinity-groups
	//GET zstack/v1/affinity-groups/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/affinity-groups")
	} else {
		url = fmt.Sprintf("zstack/v1/affinity-groups/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryAffinityGroupFailed, err)
	}
	var result model.QueryAffinityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新亲和组
func UpdateAffinityGroup(params model.UpdateAffinityGroupRequest) mgresult.Result {
	//PUT zstack/v1/affinity-groups/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/affinity-groups/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateAffinityGroupFailed, err)
	}
	var result model.UpdateAffinityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加云主机到亲和组
func AddVmToAffinityGroup(params model.AddVmToAffinityGroupRequest) mgresult.Result {
	//POST zstack/v1/affinity-groups/{affinityGroupUuid}/vm-instances/{uuid}
	url := fmt.Sprintf("zstack/v1/affinity-groups/%s/vm-instances/%s", params.AffinityGroupUuid, params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddVmToAffinityGroupFailed, err)
	}
	var result model.AddVmToAffinityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从亲和组移除云主机
func RemoveVmFromAffinityGroup(params model.RemoveVmFromAffinityGroupRequest) mgresult.Result {
	//DELETE zstack/v1/affinity-groups/{affinityGroupUuid}/vm-instances?uuid={uuid}
	url := fmt.Sprintf("zstack/v1/affinity-groups/%s/vm-instances", params.AffinityGroupUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveVmFromAffinityGroupFailed, err)
	}
	var result model.RemoveVmFromAffinityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//改变亲和组的使用状态
func ChangeAffinityGroupState(params model.ChangeAffinityGroupStateRequest) mgresult.Result {
	//PUT zstack/v1/affinity-groups/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/affinity-groups/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeAffinityGroupStateFailed, err)
	}
	var result model.ChangeAffinityGroupStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取可绑定云主机的亲和组
func GetCandidateAffinityGroupForAttachingVm(params model.GetCandidateAffinityGroupForAttachingVmRequest) mgresult.Result {
	//GET zstack/v1/affinityGroup/attachingVm
	url := fmt.Sprintf("zstack/v1/affinityGroup/attachingVm")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCandidateAffinityGroupForAttachingVmFailed, err)
	}
	var result model.GetCandidateAffinityGroupForAttachingVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取可绑定亲和组的云主机
func GetCandidateVMForAttachingAffinityGroup(params model.GetCandidateVMForAttachingAffinityGroupRequest) mgresult.Result {
	//GET zstack/v1/VM/attachingGroup
	url := fmt.Sprintf("zstack/v1/VM/attachingGroup")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCandidateVMForAttachingAffinityGroupFailed, err)
	}
	var result model.GetCandidateVMForAttachingAffinityGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建云主机时获取可加入非亲和组
func GetCandidateAffinityGroupForCreatingVm(params model.GetCandidateAffinityGroupForCreatingVmRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/candidate-affinityGroup
	url := fmt.Sprintf("zstack/v1/vm-instances/candidate-affinityGroup")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCandidateAffinityGroupForCreatingVmFailed, err)
	}
	var result model.GetCandidateAffinityGroupForCreatingVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
