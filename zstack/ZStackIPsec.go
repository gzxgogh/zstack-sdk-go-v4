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

//创建IPsec连接
func CreateIPsecConnection(params model.CreateIPsecConnectionRequest) models.Result {
	//POST zstack/v1/ipsec
	url := fmt.Sprintf("zstack/v1/ipsec")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateIPsecConnectionFailed, err.Error())
	}
	var result model.CreateIPsecConnectionRequest
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除IPsec连接
func DeleteIPsecConnection(params model.DeleteIPsecConnectionRequest) models.Result {
	//DELETE /v1/ipsec/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/ipsec/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteIPsecConnectionFailed, err.Error())
	}
	var result model.DeleteIPsecConnectionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新IPsec连接
func UpdateIPsecConnection(params model.UpdateIPsecConnectionRequest) models.Result {
	//PUT zstack/v1/ipsec/{uuid}
	url := fmt.Sprintf("zstack/v1/ipsec/%s", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateIPsecConnectionFailed, err.Error())
	}
	var result model.UpdateIPsecConnectionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询IPsec连接
func QueryIPSecConnection(params model.QueryIPSecConnectionRequest) models.Result {
	//GET zstack/v1/ipsec
	//GET zstack/v1/ipsec/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/ipsec")
	} else {
		url = fmt.Sprintf("zstack/v1/ipsec/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryIPSecConnectionFailed, err.Error())
	}
	var result model.QueryIPSecConnectionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更改IPsec连接状态
func ChangeIPSecConnectionState(params model.ChangeIPSecConnectionStateRequest) models.Result {
	//PUT zstack/v1/ipsec/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/ipsec/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeIPSecConnectionStateFailed, err.Error())
	}
	var result model.ChangeIPSecConnectionStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加三层网络到IPsec连接
func AttachL3NetworksToIPsecConnection(params model.AttachL3NetworksToIPsecConnectionRequest) models.Result {
	//POST zstack/v1/ipsec/{uuid}/l3networks
	url := fmt.Sprintf("zstack/v1/ipsec/%s/l3networks", params.UUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachL3NetworksToIPsecConnectionFailed, err.Error())
	}
	var result model.AttachL3NetworksToIPsecConnectionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从IPsec连接删除三层网络
func DetachL3NetworksFromIPsecConnection(params model.DetachL3NetworksFromIPsecConnectionRequest) models.Result {
	//DELETE zstack/v1/ipsec/{uuid}/l3networks?l3NetworkUuids={l3NetworkUuids}
	url := fmt.Sprintf("zstack/v1/ipsec/%s/l3networks?l3NetworkUuids=%s", params.UUID, params.L3NetworkUuids)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachL3NetworksFromIPsecConnectionFailed, err.Error())
	}
	var result model.DetachL3NetworksFromIPsecConnectionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加远端CIDR到IPsec连接
func AddRemoteCidrsToIPsecConnection(params model.AddRemoteCidrsToIPsecConnectionRequest) models.Result {
	//POST zstack/v1/ipsec/{uuid}/remote-cidrs
	url := fmt.Sprintf("zstack/v1/ipsec/%s/remote-cidrs", params.UUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddRemoteCidrsToIPsecConnectionFailed, err.Error())
	}
	var result model.AddRemoteCidrsToIPsecConnectionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除远端CIDR
func RemoveRemoteCidrsFromIPsecConnection(params model.RemoveRemoteCidrsFromIPsecConnectionRequest) models.Result {
	//DELETE zstack/v1/ipsec/{uuid}/remote-cidrs?peerCidrs={peerCidrs}
	url := fmt.Sprintf("zstack/v1/ipsec/%s/remote-cidrs?peerCidrs=%s", params.UUID, params.PeerCidrs)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveRemoteCidrsFromIPsecConnectionFailed, err.Error())
	}
	var result model.RemoveRemoteCidrsFromIPsecConnectionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
