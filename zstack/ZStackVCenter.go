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

//添加vCenter资源
func AddVCenter(params model.AddVCenterRequest) models.Result {
	//POST zstack/v1/vcenters
	url := fmt.Sprintf("zstack/v1/vcenters")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddVCenterFailed, err.Error())
	}
	var result model.AddVCenterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询vCenter资源
func QueryVCenter(params model.QueryVCenterRequest) models.Result {
	//GET zstack/v1/vcenters
	//GET zstack/v1/vcenters/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vcenters")
	} else {
		url = fmt.Sprintf("zstack/v1/vcenters/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVCenterFailed, err.Error())
	}
	var result model.QueryVCenterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除vCenter资源
func DeleteVCenter(params model.DeleteVCenterRequest) models.Result {
	//DELETE zstack/v1/vcenters/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/vcenters/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteVCenterFailed, err.Error())
	}
	var result model.DeleteVCenterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新vCenter资源
func UpdateVCenter(params model.UpdateVCenterRequest) models.Result {
	//PUT zstack/v1/vcenters/{uuid}/actions
	url := fmt.Sprintf("PUT zstack/v1/vcenters/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateVCenterFailed, err.Error())
	}
	var result model.UpdateVCenterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//同步vCenter资源
func SyncVCenter(params model.SyncVCenterRequest) models.Result {
	//PUT zstack/v1/vcenters/{uuid}/actions
	url := fmt.Sprintf("PUT zstack/v1/vcenters/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.SyncVCenterFailed, err.Error())
	}
	var result model.SyncVCenterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询vCenter集群
func QueryVCenterCluster(params model.QueryVCenterClusterRequest) models.Result {
	//GET zstack/v1/vcenters/clusters
	//GET zstack/v1/vcenters/clusters/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vcenters/clusters")
	} else {
		url = fmt.Sprintf("zstack/v1/vcenters/clusters/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVCenterClusterFailed, err.Error())
	}
	var result model.QueryVCenterClusterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询vCenter主存储
func QueryVCenterPrimaryStorage(params model.QueryVCenterPrimaryStorageRequest) models.Result {
	//GET zstack/v1/vcenters/primary-storage
	//GET zstack/v1/vcenters/primary-storage/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vcenters/primary-storage")
	} else {
		url = fmt.Sprintf("zstack/v1/vcenters/primary-storage/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVCenterPrimaryStorageFailed, err.Error())
	}
	var result model.QueryVCenterPrimaryStorageRequest
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询vCenter镜像服务器
func QueryVCenterBackupStorage(params model.QueryVCenterBackupStorageRequest) models.Result {
	//GET zstack/v1/vcenters/backup-storage
	//GET zstack/v1/vcenters/backup-storage/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vcenters/backup-storage")
	} else {
		url = fmt.Sprintf("zstack/v1/vcenters/backup-storage/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVCenterBackupStorageFailed, err.Error())
	}
	var result model.QueryVCenterBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询vCenter资源池
func QueryVCenterResourcePool(params model.QueryVCenterResourcePoolRequest) models.Result {
	//GET zstack/v1/vcenters/clusters/resourcepools
	//GET zstack/v1/vcenters/clusters/resourcepools/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/vcenters/clusters/resourcepools")
	} else {
		url = fmt.Sprintf("zstack/v1/vcenters/clusters/resourcepools/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryVCenterResourcePoolFailed, err.Error())
	}
	var result model.QueryVCenterResourcePoolResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
