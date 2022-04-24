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

//L3网络中是否存在可用VF网卡
func IsVfNicAvailableInL3Network(params model.IsVfNicAvailableInL3NetworkRequest) mgresult.Result {
	//GET zstack/v1/l3-networks/{l3NetworkUuid}/hosts/{hostUuid}/vfnicavailable
	url := fmt.Sprintf("zstack/v1/l3-networks/%s/hosts/%s/vfnicavailable", params.L3NetworkUUID, params.HostUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.IsVfNicAvailableInL3NetworkFailed, err)
	}
	var result model.IsVfNicAvailableInL3NetworkResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//修改云主机网卡类型
func ChangeVmNicType(params model.ChangeVmNicTypeRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/nics/{vmNicUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/nics/%s/actions", params.VmNicUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeVmNicTypeFailed, err)
	}
	var result model.ChangeVmNicTypeResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
