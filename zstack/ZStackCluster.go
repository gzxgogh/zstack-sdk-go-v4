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

//创建一个集群
func CreateCluster(params model.CreateClusterRequest) models.Result {
	//POST zstack/v1/clusters
	url := fmt.Sprintf("zstack/v1/clusters")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateClusterFailed, err.Error())
	}
	var result model.CreateClusterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除一个集群
func DeleteCluster(params model.DeleteClusterRequest) models.Result {
	//DELETE zstack/v1/clusters/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/clusters/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteClusterFailed, err.Error())
	}
	var result model.DeleteClusterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询集群
func QueryCluster(params model.QueryClusterRequest) models.Result {
	//GET zstack/v1/clusters
	//GET zstack/v1/clusters/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/clusters")
	} else {
		url = fmt.Sprintf("zstack/v1/clusters/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryClusterFailed, err.Error())
	}
	var result model.QueryClusterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新集群
func UpdateCluster(params model.UpdateClusterRequest) models.Result {
	//PUT zstack/v1/clusters/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/clusters/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateClusterFailed, err.Error())
	}
	var result model.UpdateClusterParams
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//改变一个集群的可用状态
func ChangeClusterState(params model.ChangeClusterStateRequest) models.Result {
	//PUT zstack/v1/clusters/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/clusters/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeClusterStateFailed, err.Error())
	}
	var result model.ChangeClusterStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//升级集群内物理机的操作系统
func UpdateClusterOS(params model.UpdateClusterOSRequest) models.Result {
	//PUT zstack/v1/clusters/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/clusters/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateClusterOSFailed, err.Error())
	}
	var result model.UpdateClusterOSResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
