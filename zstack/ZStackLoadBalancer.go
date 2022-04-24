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

//创建负载均衡器
func CreateLoadBalancer(params model.CreateLoadBalancerRequest) mgresult.Result {
	//POST zstack/v1/load-balancers
	url := fmt.Sprintf("zstack/v1/load-balancers")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateLoadBalancerFailed, err)
	}
	var result model.CreateLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除负载均衡器
func DeleteLoadBalancer(params model.DeleteLoadBalancerRequest) mgresult.Result {
	//DELETE zstack/v1/load-balancers/listeners/{uuid}
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteLoadBalancerFailed, err)
	}
	var result model.DeleteLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询负载均衡器
func QueryLoadBalancer(params model.QueryLoadBalancerRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryLoadBalancerFailed, err)
	}
	var result model.QueryLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//刷新负载均衡器
func RefreshLoadBalancer(params model.RefreshLoadBalancerRequest) mgresult.Result {
	//PUT zstack/v1/load-balancers/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RefreshLoadBalancerFailed, err)
	}
	var result model.RefreshLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建负载均衡监听器
func CreateLoadBalancerListener(params model.CreateLoadBalancerListenerRequest) mgresult.Result {
	//POST zstack/v1/load-balancers/{loadBalancerUuid}/listeners
	url := fmt.Sprintf("zstack/v1/load-balancers/%s/listeners", params.LoadBalancerUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateLoadBalancerListenerFailed, err)
	}
	var result model.CreateLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除负载均衡监听器
func DeleteLoadBalancerListener(params model.DeleteLoadBalancerListenerRequest) mgresult.Result {
	//DELETE /v1/load-balancers/listeners/{uuid}
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteLoadBalancerListenerFailed, err)
	}
	var result model.DeleteLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询负载均衡监听器
func QueryLoadBalancerListener(params model.QueryLoadBalancerListenerRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryLoadBalancerListenerFailed, err)
	}
	var result model.QueryLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新负载均衡监听器
func UpdateLoadBalancerListener(params model.UpdateLoadBalancerListenerRequest) mgresult.Result {
	//PUT zstack/v1/load-balancers/listeners/{uuid}
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateLoadBalancerListenerFailed, err)
	}
	var result model.UpdateLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//修改负载均衡监听器参数
func ChangeLoadBalancerListener(params model.ChangeLoadBalancerListenerRequest) mgresult.Result {
	//PUT zstack/v1/load-balancers/listeners/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeLoadBalancerListenerFailed, err)
	}
	var result model.ChangeLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取云主机网卡
func GetCandidateVmNicsForLoadBalancer(params model.GetCandidateVmNicsForLoadBalancerRequest) mgresult.Result {
	//GET zstack/v1/load-balancers/listeners/{listenerUuid}/vm-instances/candidate-nics
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/vm-instances/candidate-nics", params.ListenerUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCandidateVmNicsForLoadBalancerFailed, err)
	}
	var result model.GetCandidateVmNicsForLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加云主机网卡到负载均衡器
func AddVmNicToLoadBalancer(params model.AddVmNicToLoadBalancerRequest) mgresult.Result {
	//POST zstack/v1/load-balancers/listeners/{listenerUuid}/vm-instances/nics
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/vm-instances/nics", params.ListenerUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddVmNicToLoadBalancerFailed, err)
	}
	var result model.AddVmNicToLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从负载均衡器移除云主机网卡
func RemoveVmNicFromLoadBalancer(params model.RemoveVmNicFromLoadBalancerRequest) mgresult.Result {
	//DELETE /v1/load-balancers/listeners/{listenerUuid}/vm-instances/nics?vmNicUuids={vmNicUuids}
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/vm-instances/nics?vmNicUuids=%s", params.ListenerUuid, params.VmNicUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveVmNicFromLoadBalancerFailed, err)
	}
	var result model.RemoveVmNicFromLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新负载均衡器
func UpdateLoadBalancer(params model.UpdateLoadBalancerRequest) mgresult.Result {
	//PUT zstack/v1/load-balancers/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateLoadBalancerFailed, err)
	}
	var result model.UpdateLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建证书
func CreateCertificate(params model.CreateCertificateRequest) mgresult.Result {
	//POST zstack/v1/certificates
	url := fmt.Sprintf("zstack/v1/certificates")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateCertificateFailed, err)
	}
	var result model.CreateCertificateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除证书
func DeleteCertificate(params model.DeleteCertificateRequest) mgresult.Result {
	//DELETE zstack/v1/certificates/{uuid}
	url := fmt.Sprintf("zstack/v1/certificates/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteCertificateFailed, err)
	}
	var result model.DeleteCertificateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询证书
func QueryCertificate(params model.QueryCertificateRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryCertificateFailed, err)
	}
	var result model.QueryCertificateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加证书到负载均衡
func AddCertificateToLoadBalancerListener(params model.AddCertificateToLoadBalancerListenerRequest) mgresult.Result {
	//POST zstack/v1/load-balancers/listeners/{listenerUuid}/certificate
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/certificate", params.ListenerUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddCertificateToLoadBalancerListenerFailed, err)
	}
	var result model.AddCertificateToLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从负载均衡移除证书
func RemoveCertificateFromLoadBalancerListener(params model.RemoveCertificateFromLoadBalancerListenerRequest) mgresult.Result {
	//DELETE zstack/v1/load-balancers/listeners/{listenerUuid}/certificate
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/certificate", params.ListenerUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveCertificateFromLoadBalancerListenerFailed, err)
	}
	var result model.RemoveCertificateFromLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新证书信息
func UpdateCertificate(params model.UpdateCertificateRequest) mgresult.Result {
	//PUT zstack/v1/certificates/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/certificates/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateCertificateFailed, err)
	}
	var result model.UpdateCertificateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加转发规则到访问控制策略组
func AddAccessControlListRedirectRule(params model.AddAccessControlListRedirectRuleRequest) mgresult.Result {
	//POST zstack/v1/access-control-lists/{aclUuid}/redirectRules
	url := fmt.Sprintf("zstack/v1/access-control-lists/%s/redirectRules", params.AclUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddAccessControlListRedirectRuleFailed, err)
	}
	var result model.AddAccessControlListRedirectRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//修改控制策略组转发规则名称
func ChangeAccessControlListRedirectRule(params model.ChangeAccessControlListRedirectRuleRequest) mgresult.Result {
	//PUT zstack/v1/access-control-lists/redirectRules/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/access-control-lists/redirectRules/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeAccessControlListRedirectRuleFailed, err)
	}
	var result model.ChangeAccessControlListRedirectRuleResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//修改访问控制组绑定的后端服务器组
func ChangeAccessControlListServerGroup(params model.ChangeAccessControlListServerGroupRequest) mgresult.Result {
	//PUT zstack/v1/load-balancers/listener/acl/{aclUuid}/servergroup/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/listener/acl/%s/servergroup/actions", params.AclUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeAccessControlListServerGroupFailed, err)
	}
	var result model.ChangeAccessControlListServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取负载均衡监听器访问控制策略条目
func GetLoadBalancerListenerACLEntries(params model.GetLoadBalancerListenerACLEntriesRequest) mgresult.Result {
	//GET zstack/v1/load-balancers/listeners/access-control-lists/entries
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/access-control-lists/entries")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetLoadBalancerListenerACLEntriesFailed, err)
	}
	var result model.GetLoadBalancerListenerACLEntriesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建访问控制策略组
func CreateAccessControlList(params model.CreateAccessControlListRequest) mgresult.Result {
	//POST zstack/v1/access-control-lists
	url := fmt.Sprintf("zstack/v1/access-control-lists")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateAccessControlListFailed, err)
	}
	var result model.CreateAccessControlListResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除访问控制策略组
func DeleteAccessControlList(params model.DeleteAccessControlListRequest) mgresult.Result {
	//DELETE zstack/v1/access-control-lists/{uuid}
	url := fmt.Sprintf("zstack/v1/access-control-lists/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteAccessControlListFailed, err)
	}
	var result model.DeleteAccessControlListResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询访问控制策略组
func QueryAccessControlList(params model.QueryAccessControlListRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryAccessControlListFailed, err)
	}
	var result model.QueryAccessControlListResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//向访问控制策略组添加IP组
func AddAccessControlListEntry(params model.AddAccessControlListEntryRequest) mgresult.Result {
	//POST zstack/v1/access-control-lists/{aclUuid}/ipentries
	url := fmt.Sprintf("zstack/v1/access-control-lists/%s/ipentries", params.AclUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddAccessControlListEntryFailed, err)
	}
	var result model.AddAccessControlListEntryResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除访问控制策略的IP组
func RemoveAccessControlListEntry(params model.RemoveAccessControlListEntryRequest) mgresult.Result {
	//DELETE zstack/v1/access-control-lists/{aclUuid}/ipentries/{uuid}
	url := fmt.Sprintf("zstack/v1/access-control-lists/%s/ipentries/%s", params.AclUuid, params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveAccessControlListEntryFailed, err)
	}
	var result model.RemoveAccessControlListEntryResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加监听器的访问控制策略
func AddAccessControlListToLoadBalancer(params model.AddAccessControlListToLoadBalancerRequest) mgresult.Result {
	//POST zstack/v1/load-balancers/listeners/{listenerUuid}/access-control-lists
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/access-control-lists", params.ListenerUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddAccessControlListToLoadBalancerFailed, err)
	}
	var result model.AddAccessControlListToLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除监听器访问控制策略
func RemoveAccessControlListFromLoadBalancer(params model.RemoveAccessControlListFromLoadBalancerRequest) mgresult.Result {
	//DELETE zstack/v1/load-balancers/listeners/{listenerUuid}/access-control-lists
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/access-control-lists", params.ListenerUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveAccessControlListFromLoadBalancerFailed, err)
	}
	var result model.RemoveAccessControlListFromLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取监听器可加载L3网络
func GetCandidateL3NetworksForLoadBalancer(params model.GetCandidateL3NetworksForLoadBalancerRequest) mgresult.Result {
	//GET zstack/v1/load-balancers/listeners/{listenerUuid}/networks/candidates
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/networks/candidates", params.ListenerUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCandidateL3NetworksForLoadBalancerFailed, err)
	}
	var result model.GetCandidateL3NetworksForLoadBalancerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建负载均衡器服务器组
func CreateLoadBalancerServerGroup(params model.CreateLoadBalancerServerGroupRequest) mgresult.Result {
	//POST zstack/v1/load-balancers/{loadBalancerUuid}/servergroups
	url := fmt.Sprintf("zstack/v1/load-balancers/%s/servergroups", params.LoadBalancerUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateLoadBalancerServerGroupFailed, err)
	}
	var result model.CreateLoadBalancerServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除负载均衡器服务器组
func DeleteLoadBalancerServerGroup(params model.DeleteLoadBalancerServerGroupRequest) mgresult.Result {
	//DELETE zstack/v1/load-balancers/servergroups/{uuid}
	url := fmt.Sprintf("zstack/v1/load-balancers/servergroups/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteLoadBalancerServerGroupFailed, err)
	}
	var result model.DeleteLoadBalancerServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新负载均衡器服务器组
func UpdateLoadBalancerServerGroup(params model.UpdateLoadBalancerServerGroupRequest) mgresult.Result {
	//PUT zstack/v1/load-balancers/servergroups/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/servergroups/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateLoadBalancerServerGroupFailed, err)
	}
	var result model.UpdateLoadBalancerServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询负载均衡器服务器组
func QueryLoadBalancerServerGroup(params model.QueryLoadBalancerServerGroupRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QueryLoadBalancerServerGroupFailed, err)
	}
	var result model.QueryLoadBalancerServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加后端服务器到服务器组
func AddBackendServerToServerGroup(params model.AddBackendServerToServerGroupRequest) mgresult.Result {
	//POST zstack/v1/load-balancers/servergroups/{serverGroupUuid}/backendservers
	url := fmt.Sprintf("zstack/v1/load-balancers/servergroups/%s/backendservers", params.ServerGroupUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddBackendServerToServerGroupFailed, err)
	}
	var result model.AddBackendServerToServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从服务器组移除后端服务器
func RemoveBackendServerFromServerGroup(params model.RemoveBackendServerFromServerGroupRequest) mgresult.Result {
	//PUT zstack/v1/load-balancers/servergroups/{serverGroupUuid}/backendservers/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/servergroups/%s/backendservers/actions", params.ServerGroupUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveBackendServerFromServerGroupFailed, err)
	}
	var result model.RemoveBackendServerFromServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//添加服务器组到负载均衡监听器
func AddServerGroupToLoadBalancerListener(params model.AddServerGroupToLoadBalancerListenerRequest) mgresult.Result {
	//POST zstack/v1/load-balancers/listeners/{listenerUuid}/servergroups
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/servergroups", params.ListenerUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddServerGroupToLoadBalancerListenerFailed, err)
	}
	var result model.AddBackendServerToServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//从负载均衡监听器移除服务器组
func RemoveServerGroupFromLoadBalancerListener(params model.RemoveServerGroupFromLoadBalancerListenerRequest) mgresult.Result {
	//DELETE zstack/v1/load-balancers/listeners/{listenerUuid}/servergroups
	url := fmt.Sprintf("zstack/v1/load-balancers/listeners/%s/servergroups", params.ListenerUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveServerGroupFromLoadBalancerListenerFailed, err)
	}
	var result model.RemoveServerGroupFromLoadBalancerListenerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新负载均衡后端服务器参数
func ChangeLoadBalancerBackendServer(params model.ChangeLoadBalancerBackendServerRequest) mgresult.Result {
	//PUT zstack/v1/load-balancers/servergroups/{serverGroupUuid}/backendserver/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/servergroups/%s/backendserver/actions", params.ServerGroupUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.ChangeLoadBalancerBackendServerFailed, err)
	}
	var result model.ChangeLoadBalancerBackendServerResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取可供负载均衡器服务器组添加的云主机网卡
func GetCandidateVmNicsForLoadBalancerServerGroup(params model.GetCandidateVmNicsForLoadBalancerServerGroupRequest) mgresult.Result {
	//GET zstack/v1/load-balancers/servergroups/candidate-nics
	url := fmt.Sprintf("zstack/v1/load-balancers/servergroups/candidate-nics")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetCandidateVmNicsForLoadBalancerServerGroupFailed, err)
	}
	var result model.GetCandidateVmNicsForLoadBalancerServerGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建性能独享型负载均衡器组
func CreateSlbGroup(params model.CreateSlbGroupRequest) mgresult.Result {
	//POST zstack/v1/load-balancers/slb/groups
	url := fmt.Sprintf("zstack/v1/load-balancers/slb/groups")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateSlbGroupFailed, err)
	}
	var result model.CreateSlbGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除性能独享型负载均衡器组
func DeleteSlbGroup(params model.DeleteSlbGroupRequest) mgresult.Result {
	//DELETE zstack/v1/load-balancers/slb/group/{uuid}
	url := fmt.Sprintf("zstack/v1/load-balancers/slb/group/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeleteSlbGroupFailed, err)
	}
	var result model.DeleteSlbGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新性能独享型负载均衡器组
func UpdateSlbGroup(params model.UpdateSlbGroupRequest) mgresult.Result {
	//PUT zstack/v1/load-balancers/slb/group/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/load-balancers/slb/group/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateSlbGroupFailed, err)
	}
	var result model.UpdateSlbGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询性能独享型负载均衡器组
func QuerySlbGroup(params model.QuerySlbGroupRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QuerySlbGroupFailed, err)
	}
	var result model.QuerySlbGroupResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建高性能实例性负载均衡器规格
func CreateSlbOffering(params model.CreateSlbOfferingRequest) mgresult.Result {
	//POST zstack/v1/instance-offerings/slb
	url := fmt.Sprintf("zstack/v1/instance-offerings/slb")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateSlbOfferingFailed, err)
	}
	var result model.CreateSlbOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询性能独享型负载均衡器规格
func QuerySlbOffering(params model.QuerySlbOfferingRequest) mgresult.Result {
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
		return mgerr.ErrorResultWithErr(errcode.QuerySlbOfferingFailed, err)
	}
	var result model.QuerySlbOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建高性能负载均衡器实例
func CreateSlbInstance(params model.CreateSlbInstanceRequest) mgresult.Result {
	//POST zstack/v1/load-balancers/slb/instances
	url := fmt.Sprintf("zstack/v1/load-balancers/slb/instances")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreateSlbInstanceFailed, err)
	}
	var result model.CreateSlbInstanceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
