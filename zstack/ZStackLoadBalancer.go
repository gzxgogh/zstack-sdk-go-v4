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

//创建负载均衡器
func CreateLoadBalancer(params model.CreateLoadBalancerRequest) models.Result {
	//POST zstack/v1/load-balancers
	url := fmt.Sprintf("zstack/v1/load-balancers")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateLoadBalancerFailed, err.Error())
	}
	var result model.CreateLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除负载均衡器
func DeleteLoadBalancer(params model.DeleteLoadBalancerRequest) models.Result {
	//DELETE zstack/v1/load-balancers/listeners/{uuid}
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteLoadBalancerFailed, err.Error())
	}
	var result model.DeleteLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询负载均衡器
func QueryLoadBalancer(params model.QueryLoadBalancerRequest) models.Result {
	//GET zstack/v1/load-balancers
	//GET zstack/v1/load-balancers/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/load-balancers")
	} else {
		url = fmt.Sprintf("zstack/v1/load-balancers/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryLoadBalancerFailed, err.Error())
	}
	var result model.QueryLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//刷新负载均衡器
func RefreshLoadBalancer(params model.RefreshLoadBalancerRequest) models.Result {
	//PUT zstack/v1/load-balancers/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.RefreshLoadBalancerFailed, err.Error())
	}
	var result model.RefreshLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建负载均衡监听器
func CreateLoadBalancerListener(params model.CreateLoadBalancerListenerRequest) models.Result {
	//POST zstack/v1/load-balancers/{loadBalancerUuid}/listeners
	url := fmt.Sprintf("zstack/v1/load-balancers/%s/listeners", params.LoadBalancerUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateLoadBalancerListenerFailed, err.Error())
	}
	var result model.CreateLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除负载均衡监听器
func DeleteLoadBalancerListener(params model.DeleteLoadBalancerListenerRequest) models.Result {
	//DELETE /v1/load-balancers/listeners/{uuid}
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteLoadBalancerListenerFailed, err.Error())
	}
	var result model.DeleteLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询负载均衡监听器
func QueryLoadBalancerListener(params model.QueryLoadBalancerListenerRequest) models.Result {
	//GET zstack/v1/load-balancers/listeners
	//GET zstack/v1/load-balancers/listeners/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/load-balancers/listeners")
	} else {
		url = fmt.Sprintf("zstack/v1/load-balancers/listeners/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryLoadBalancerListenerFailed, err.Error())
	}
	var result model.QueryLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新负载均衡监听器
func UpdateLoadBalancerListener(params model.UpdateLoadBalancerListenerRequest) models.Result {
	//PUT zstack/v1/load-balancers/listeners/{uuid}
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateLoadBalancerListenerFailed, err.Error())
	}
	var result model.UpdateLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改负载均衡监听器参数
func ChangeLoadBalancerListener(params model.ChangeLoadBalancerListenerRequest) models.Result {
	//PUT zstack/v1/load-balancers/listeners/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeLoadBalancerListenerFailed, err.Error())
	}
	var result model.ChangeLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取云主机网卡
func GetCandidateVmNicsForLoadBalancer(params model.GetCandidateVmNicsForLoadBalancerRequest) models.Result {
	//GET zstack/v1/load-balancers/listeners/{listenerUuid}/vm-instances/candidate-nics
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/vm-instances/candidate-nics", params.ListenerUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetCandidateVmNicsForLoadBalancerFailed, err.Error())
	}
	var result model.GetCandidateVmNicsForLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加云主机网卡到负载均衡器
func AddVmNicToLoadBalancer(params model.AddVmNicToLoadBalancerRequest) models.Result {
	//POST zstack/v1/load-balancers/listeners/{listenerUuid}/vm-instances/nics
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/vm-instances/nics", params.ListenerUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddVmNicToLoadBalancerFailed, err.Error())
	}
	var result model.AddVmNicToLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从负载均衡器移除云主机网卡
func RemoveVmNicFromLoadBalancer(params model.RemoveVmNicFromLoadBalancerRequest) models.Result {
	//DELETE /v1/load-balancers/listeners/{listenerUuid}/vm-instances/nics?vmNicUuids={vmNicUuids}
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/vm-instances/nics?vmNicUuids=%s", params.ListenerUuid, params.VmNicUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveVmNicFromLoadBalancerFailed, err.Error())
	}
	var result model.RemoveVmNicFromLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新负载均衡器
func UpdateLoadBalancer(params model.UpdateLoadBalancerRequest) models.Result {
	//PUT zstack/v1/load-balancers/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateLoadBalancerFailed, err.Error())
	}
	var result model.UpdateLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建证书
func CreateCertificate(params model.CreateCertificateRequest) models.Result {
	//POST zstack/v1/certificates
	url := fmt.Sprintf("zstack/v1/certificates")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateCertificateFailed, err.Error())
	}
	var result model.CreateCertificateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除证书
func DeleteCertificate(params model.DeleteCertificateRequest) models.Result {
	//DELETE zstack/v1/certificates/{uuid}
	url := fmt.Sprintf("zstack/v1/certificates/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteCertificateFailed, err.Error())
	}
	var result model.DeleteCertificateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询证书
func QueryCertificate(params model.QueryCertificateRequest) models.Result {
	//GET zstack/v1/certificates
	//GET zstack/v1/certificates/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/certificates")
	} else {
		url = fmt.Sprintf("zstack/v1/certificates/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryCertificateFailed, err.Error())
	}
	var result model.QueryCertificateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加证书到负载均衡
func AddCertificateToLoadBalancerListener(params model.AddCertificateToLoadBalancerListenerRequest) models.Result {
	//POST zstack/v1/load-balancers/listeners/{listenerUuid}/certificate
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/certificate", params.ListenerUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddCertificateToLoadBalancerListenerFailed, err.Error())
	}
	var result model.AddCertificateToLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从负载均衡移除证书
func RemoveCertificateFromLoadBalancerListener(params model.RemoveCertificateFromLoadBalancerListenerRequest) models.Result {
	//DELETE zstack/v1/load-balancers/listeners/{listenerUuid}/certificate
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/certificate", params.ListenerUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveCertificateFromLoadBalancerListenerFailed, err.Error())
	}
	var result model.RemoveCertificateFromLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新证书信息
func UpdateCertificate(params model.UpdateCertificateRequest) models.Result {
	//PUT zstack/v1/certificates/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/certificates/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateCertificateFailed, err.Error())
	}
	var result model.UpdateCertificateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加转发规则到访问控制策略组
func AddAccessControlListRedirectRule(params model.AddAccessControlListRedirectRuleRequest) models.Result {
	//POST zstack/v1/access-control-lists/{aclUuid}/redirectRules
	url := fmt.Sprintf("zstack/v1/access-control-lists/%s/redirectRules", params.AclUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddAccessControlListRedirectRuleFailed, err.Error())
	}
	var result model.AddAccessControlListRedirectRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改控制策略组转发规则名称
func ChangeAccessControlListRedirectRule(params model.ChangeAccessControlListRedirectRuleRequest) models.Result {
	//PUT zstack/v1/access-control-lists/redirectRules/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/access-control-lists/redirectRules/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeAccessControlListRedirectRuleFailed, err.Error())
	}
	var result model.ChangeAccessControlListRedirectRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改访问控制组绑定的后端服务器组
func ChangeAccessControlListServerGroup(params model.ChangeAccessControlListServerGroupRequest) models.Result {
	//PUT zstack/v1/load-balancers/listener/acl/{aclUuid}/servergroup/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/listener/acl/%s/servergroup/actions", params.AclUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeAccessControlListServerGroupFailed, err.Error())
	}
	var result model.ChangeAccessControlListServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取负载均衡监听器访问控制策略条目
func GetLoadBalancerListenerACLEntries(params model.GetLoadBalancerListenerACLEntriesRequest) models.Result {
	//GET zstack/v1/load-balancers/listeners/access-control-lists/entries
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/access-control-lists/entries")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetLoadBalancerListenerACLEntriesFailed, err.Error())
	}
	var result model.GetLoadBalancerListenerACLEntriesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建访问控制策略组
func CreateAccessControlList(params model.CreateAccessControlListRequest) models.Result {
	//POST zstack/v1/access-control-lists
	url := fmt.Sprintf("zstack/v1/access-control-lists")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateAccessControlListFailed, err.Error())
	}
	var result model.CreateAccessControlListResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除访问控制策略组
func DeleteAccessControlList(params model.DeleteAccessControlListRequest) models.Result {
	//DELETE zstack/v1/access-control-lists/{uuid}
	url := fmt.Sprintf("zstack/v1/access-control-lists/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteAccessControlListFailed, err.Error())
	}
	var result model.DeleteAccessControlListResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询访问控制策略组
func QueryAccessControlList(params model.QueryAccessControlListRequest) models.Result {
	//GET zstack/v1/access-control-lists
	//GET zstack/v1/access-control-lists/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/access-control-lists")
	} else {
		url = fmt.Sprintf("zstack/v1/access-control-lists/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAccessControlListFailed, err.Error())
	}
	var result model.QueryAccessControlListResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//向访问控制策略组添加IP组
func AddAccessControlListEntry(params model.AddAccessControlListEntryRequest) models.Result {
	//POST zstack/v1/access-control-lists/{aclUuid}/ipentries
	url := fmt.Sprintf("zstack/v1/access-control-lists/%s/ipentries", params.AclUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddAccessControlListEntryFailed, err.Error())
	}
	var result model.AddAccessControlListEntryResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除访问控制策略的IP组
func RemoveAccessControlListEntry(params model.RemoveAccessControlListEntryRequest) models.Result {
	//DELETE zstack/v1/access-control-lists/{aclUuid}/ipentries/{uuid}
	url := fmt.Sprintf("zstack/v1/access-control-lists/%s/ipentries/%s", params.AclUuid, params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveAccessControlListEntryFailed, err.Error())
	}
	var result model.RemoveAccessControlListEntryResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加监听器的访问控制策略
func AddAccessControlListToLoadBalancer(params model.AddAccessControlListToLoadBalancerRequest) models.Result {
	//POST zstack/v1/load-balancers/listeners/{listenerUuid}/access-control-lists
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/access-control-lists", params.ListenerUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddAccessControlListToLoadBalancerFailed, err.Error())
	}
	var result model.AddAccessControlListToLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除监听器访问控制策略
func RemoveAccessControlListFromLoadBalancer(params model.RemoveAccessControlListFromLoadBalancerRequest) models.Result {
	//DELETE zstack/v1/load-balancers/listeners/{listenerUuid}/access-control-lists
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/access-control-lists", params.ListenerUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveAccessControlListFromLoadBalancerFailed, err.Error())
	}
	var result model.RemoveAccessControlListFromLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取监听器可加载L3网络
func GetCandidateL3NetworksForLoadBalancer(params model.GetCandidateL3NetworksForLoadBalancerRequest) models.Result {
	//GET zstack/v1/load-balancers/listeners/{listenerUuid}/networks/candidates
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/networks/candidates", params.ListenerUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetCandidateL3NetworksForLoadBalancerFailed, err.Error())
	}
	var result model.GetCandidateL3NetworksForLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建负载均衡器服务器组
func CreateLoadBalancerServerGroup(params model.CreateLoadBalancerServerGroupRequest) models.Result {
	//POST zstack/v1/load-balancers/{loadBalancerUuid}/servergroups
	url := fmt.Sprintf("zstack/v1/load-balancers/%s/servergroups", params.LoadBalancerUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateLoadBalancerServerGroupFailed, err.Error())
	}
	var result model.CreateLoadBalancerServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除负载均衡器服务器组
func DeleteLoadBalancerServerGroup(params model.DeleteLoadBalancerServerGroupRequest) models.Result {
	//DELETE zstack/v1/load-balancers/servergroups/{uuid}
	url := fmt.Sprintf("zstack/v1/load-balancers/servergroups/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteLoadBalancerServerGroupFailed, err.Error())
	}
	var result model.DeleteLoadBalancerServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新负载均衡器服务器组
func UpdateLoadBalancerServerGroup(params model.UpdateLoadBalancerServerGroupRequest) models.Result {
	//PUT zstack/v1/load-balancers/servergroups/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/servergroups/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateLoadBalancerServerGroupFailed, err.Error())
	}
	var result model.UpdateLoadBalancerServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询负载均衡器服务器组
func QueryLoadBalancerServerGroup(params model.QueryLoadBalancerServerGroupRequest) models.Result {
	//GET zstack/v1/load-balancers/servergroups
	//GET zstack/v1/load-balancers/servergroups/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/load-balancers/servergroups")
	} else {
		url = fmt.Sprintf("zstack/v1/load-balancers/servergroups/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryLoadBalancerServerGroupFailed, err.Error())
	}
	var result model.QueryLoadBalancerServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加后端服务器到服务器组
func AddBackendServerToServerGroup(params model.AddBackendServerToServerGroupRequest) models.Result {
	//POST zstack/v1/load-balancers/servergroups/{serverGroupUuid}/backendservers
	url := fmt.Sprintf("zstack/v1/load-balancers/servergroups/%s/backendservers", params.ServerGroupUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddBackendServerToServerGroupFailed, err.Error())
	}
	var result model.AddBackendServerToServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从服务器组移除后端服务器
func RemoveBackendServerFromServerGroup(params model.RemoveBackendServerFromServerGroupRequest) models.Result {
	//PUT zstack/v1/load-balancers/servergroups/{serverGroupUuid}/backendservers/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/servergroups/%s/backendservers/actions", params.ServerGroupUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.RemoveBackendServerFromServerGroupFailed, err.Error())
	}
	var result model.RemoveBackendServerFromServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//添加服务器组到负载均衡监听器
func AddServerGroupToLoadBalancerListener(params model.AddServerGroupToLoadBalancerListenerRequest) models.Result {
	//POST zstack/v1/load-balancers/listeners/{listenerUuid}/servergroups
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/servergroups", params.ListenerUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AddServerGroupToLoadBalancerListenerFailed, err.Error())
	}
	var result model.AddBackendServerToServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//从负载均衡监听器移除服务器组
func RemoveServerGroupFromLoadBalancerListener(params model.RemoveServerGroupFromLoadBalancerListenerRequest) models.Result {
	//DELETE zstack/v1/load-balancers/listeners/{listenerUuid}/servergroups
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/servergroups", params.ListenerUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.RemoveServerGroupFromLoadBalancerListenerFailed, err.Error())
	}
	var result model.RemoveServerGroupFromLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新负载均衡后端服务器参数
func ChangeLoadBalancerBackendServer(params model.ChangeLoadBalancerBackendServerRequest) models.Result {
	//PUT zstack/v1/load-balancers/servergroups/{serverGroupUuid}/backendserver/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/servergroups/%s/backendserver/actions", params.ServerGroupUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeLoadBalancerBackendServerFailed, err.Error())
	}
	var result model.ChangeLoadBalancerBackendServerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获取可供负载均衡器服务器组添加的云主机网卡
func GetCandidateVmNicsForLoadBalancerServerGroup(params model.GetCandidateVmNicsForLoadBalancerServerGroupRequest) models.Result {
	//GET zstack/v1/load-balancers/servergroups/candidate-nics
	url := fmt.Sprintf("zstack/v1/load-balancers/servergroups/candidate-nics")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetCandidateVmNicsForLoadBalancerServerGroupFailed, err.Error())
	}
	var result model.GetCandidateVmNicsForLoadBalancerServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建性能独享型负载均衡器组
func CreateSlbGroup(params model.CreateSlbGroupRequest) models.Result {
	//POST zstack/v1/load-balancers/slb/groups
	url := fmt.Sprintf("zstack/v1/load-balancers/slb/groups")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateSlbGroupFailed, err.Error())
	}
	var result model.CreateSlbGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除性能独享型负载均衡器组
func DeleteSlbGroup(params model.DeleteSlbGroupRequest) models.Result {
	//DELETE zstack/v1/load-balancers/slb/group/{uuid}
	url := fmt.Sprintf("zstack/v1/load-balancers/slb/group/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteSlbGroupFailed, err.Error())
	}
	var result model.DeleteSlbGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//更新性能独享型负载均衡器组
func UpdateSlbGroup(params model.UpdateSlbGroupRequest) models.Result {
	//PUT zstack/v1/load-balancers/slb/group/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/slb/group/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateSlbGroupFailed, err.Error())
	}
	var result model.UpdateSlbGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询性能独享型负载均衡器组
func QuerySlbGroup(params model.QuerySlbGroupRequest) models.Result {
	//GET zstack/v1/load-balancers/slb/groups
	//GET zstack/v1/load-balancers/slb/groups/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/load-balancers/slb/groups")
	} else {
		url = fmt.Sprintf("zstack/v1/load-balancers/slb/groups/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QuerySlbGroupFailed, err.Error())
	}
	var result model.QuerySlbGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建高性能实例性负载均衡器规格
func CreateSlbOffering(params model.CreateSlbOfferingRequest) models.Result {
	//POST zstack/v1/instance-offerings/slb
	url := fmt.Sprintf("zstack/v1/instance-offerings/slb")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateSlbOfferingFailed, err.Error())
	}
	var result model.CreateSlbOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询性能独享型负载均衡器规格
func QuerySlbOffering(params model.QuerySlbOfferingRequest) models.Result {
	//GET zstack/v1/instance-offerings/slb
	//GET zstack/v1/instance-offerings/slb/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/instance-offerings/slb")
	} else {
		url = fmt.Sprintf("zstack/v1/instance-offerings/slb/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QuerySlbOfferingFailed, err.Error())
	}
	var result model.QuerySlbOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建高性能负载均衡器实例
func CreateSlbInstance(params model.CreateSlbInstanceRequest) models.Result {
	//POST zstack/v1/load-balancers/slb/instances
	url := fmt.Sprintf("zstack/v1/load-balancers/slb/instances")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateSlbInstanceFailed, err.Error())
	}
	var result model.CreateSlbInstanceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
