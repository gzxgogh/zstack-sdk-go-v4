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

//添加vCenter资源
func AddVCenter(params model.AddVCenterRequest) mgresult.Result {
	//POST zstack/v1/vcenters
	url := fmt.Sprintf("zstack/v1/vcenters")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddVCenterFailed, err)
	}
	var result model.AddVCenterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询vCenter资源
func QueryVCenter(params model.QueryVCenterRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryVCenterFailed, err)
	}
	var result model.QueryVCenterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除vCenter资源
func DeleteVCenter(params model.DeleteVCenterRequest) mgresult.Result {
	//DELETE zstack/v1/vcenters/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/vcenters/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteVCenterFailed, err)
	}
	var result model.DeleteVCenterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新vCenter资源
func UpdateVCenter(params model.UpdateVCenterRequest) mgresult.Result {
	//PUT zstack/v1/vcenters/{uuid}/actions
	url := fmt.Sprintf("PUT zstack/v1/vcenters/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateVCenterFailed, err)
	}
	var result model.UpdateVCenterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//同步vCenter资源
func SyncVCenter(params model.SyncVCenterRequest) mgresult.Result {
	//PUT zstack/v1/vcenters/{uuid}/actions
	url := fmt.Sprintf("PUT zstack/v1/vcenters/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SyncVCenterFailed, err)
	}
	var result model.SyncVCenterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询vCenter集群
func QueryVCenterCluster(params model.QueryVCenterClusterRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryVCenterClusterFailed, err)
	}
	var result model.QueryVCenterClusterResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询vCenter主存储
func QueryVCenterPrimaryStorage(params model.QueryVCenterPrimaryStorageRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryVCenterPrimaryStorageFailed, err)
	}
	var result model.QueryVCenterPrimaryStorageRequest
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询vCenter镜像服务器
func QueryVCenterBackupStorage(params model.QueryVCenterBackupStorageRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryVCenterBackupStorageFailed, err)
	}
	var result model.QueryVCenterBackupStorageResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询vCenter资源池
func QueryVCenterResourcePool(params model.QueryVCenterResourcePoolRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryVCenterResourcePoolFailed, err)
	}
	var result model.QueryVCenterResourcePoolResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
