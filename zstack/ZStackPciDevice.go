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

//查询PCI设备
func QueryPciDevice(params model.QueryPciDeviceRequest) mgresult.Result {
	//GET zstack/v1/pci-device/pci-devices
	//GET zstack/v1/pci-device/pci-devices/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/pci-device/pci-devices")
	} else {
		url = fmt.Sprintf("zstack/v1/pci-device/pci-devices/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryPciDeviceFailed, err)
	}
	var result model.QueryPciDeviceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新PCI设备
func UpdatePciDevice(params model.UpdatePciDeviceRequest) mgresult.Result {
	//PUT zstack/v1/pci-device/pci-devices/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/pci-device/pci-devices/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdatePciDeviceFailed, err)
	}
	var result model.UpdatePciDeviceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除PCI设备
func DeletePciDevice(params model.DeletePciDeviceRequest) mgresult.Result {
	//DELETE zstack/v1/pci-device/pci-devices/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/pci-device/pci-devices/%s", params.UUID)
	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeletePciDeviceFailed, err)
	}
	var result model.DeletePciDeviceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取PCI设备列表
func GetPciDeviceCandidatesForAttachingVm(params model.GetPciDeviceCandidatesForAttachingVmRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/candidate-pci-devices
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/candidate-pci-devices", params.VmInstanceUuid)
	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetPciDeviceCandidatesForAttachingVmFailed, err)
	}
	var result model.GetPciDeviceCandidatesForAttachingVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取可加载PCI设备
func GetPciDeviceCandidatesForNewCreateVm(params model.GetPciDeviceCandidatesForNewCreateVmRequest) mgresult.Result {
	//GET zstack/v1/pci-device/candidate-pci-devices-for-new-create-vm
	url := fmt.Sprintf("zstack/v1/pci-device/candidate-pci-devices-for-new-create-vm")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetPciDeviceCandidatesForNewCreateVmFailed, err)
	}
	var result model.GetPciDeviceCandidatesForNewCreateVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//绑定PCI设备到云主机
func AttachPciDeviceToVm(params model.AttachPciDeviceToVmRequest) mgresult.Result {
	//POST zstack/v1/pci-device/pci-devices/{pciDeviceUuid}/attach
	url := fmt.Sprintf("zstack/v1/pci-device/pci-devices/%s/attach", params.PciDeviceUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachPciDeviceToVmFailed, err)
	}
	var result model.AttachPciDeviceToVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//卸载PCI设备
func DetachPciDeviceFromVm(params model.DetachPciDeviceFromVmRequest) mgresult.Result {
	//POST zstack/v1/pci-device/pci-devices/{pciDeviceUuid}/detach
	url := fmt.Sprintf("zstack/v1/pci-device/pci-devices/%s/attach", params.PciDeviceUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachPciDeviceFromVmFailed, err)
	}
	var result model.DetachPciDeviceFromVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//创建PCI设备规格
func CreatePciDeviceOffering(params model.CreatePciDeviceOfferingRequest) mgresult.Result {
	//POST zstack/v1/pci-device/pci-device-offerings
	url := fmt.Sprintf("zstack/v1/pci-device/pci-device-offerings")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.CreatePciDeviceOfferingFailed, err)
	}
	var result model.CreatePciDeviceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//删除PCI设备规格
func DeletePciDeviceOffering(params model.DeletePciDeviceOfferingRequest) mgresult.Result {
	//DELETE zstack/v1/pci-device/pci-device-offerings/{uuid}?deleteMode={deleteMode}
	url := fmt.Sprintf("zstack/v1/pci-device/pci-device-offerings/%s", params.Uuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DeletePciDeviceOfferingFailed, err)
	}
	var result model.DeletePciDeviceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询PCI设备规格信息
func QueryPciDeviceOffering(params model.QueryPciDeviceOfferingRequest) mgresult.Result {
	//GET zstack/v1/pci-device/pci-device-offerings
	//GET zstack/v1/pci-device/pci-device-offerings/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/pci-device/pci-device-offerings")
	} else {
		url = fmt.Sprintf("zstack/v1/pci-device/pci-device-offerings/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryPciDeviceOfferingFailed, err)
	}
	var result model.QueryPciDeviceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询PCI设备规格匹配
func QueryPciDevicePciDeviceOffering(params model.QueryPciDevicePciDeviceOfferingRequest) mgresult.Result {
	//GET zstack/v1/pci-devices/pci-devices/pci-device-offerings
	url := fmt.Sprintf("zstack/v1/pci-devices/pci-devices/pci-device-offerings")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryPciDevicePciDeviceOfferingFailed, err)
	}
	var result model.QueryPciDevicePciDeviceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取物理机Iommu启用状态
func GetHostIommuStatus(params model.GetHostIommuStatusRequest) mgresult.Result {
	//GET zstack/v1/pci-device/hosts/status/uuid}
	url := fmt.Sprintf("zstack/v1/pci-device/hosts/status/uuid")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetHostIommuStatusFailed, err)
	}
	var result model.GetHostIommuStatusResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新物理机Iommu启用状态
func UpdateHostIommuState(params model.UpdateHostIommuStateRequest) mgresult.Result {
	//PUT zstack/v1/pci-device/hosts/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/pci-device/hosts/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateHostIommuStateFailed, err)
	}
	var result model.UpdateHostIommuStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取物理机lommu就绪状态
func GetHostIommuState(params model.GetHostIommuStateRequest) mgresult.Result {
	//GET zstack/v1/pci-device/hosts/state/{uuid}
	url := fmt.Sprintf("zstack/v1/pci-device/hosts/state/%s", params.Uuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetHostIommuStateFailed, err)
	}
	var result model.GetHostIommuStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将PCI设备规格加载到云主机
func AddPciDeviceSpecToVmInstance(params model.AddPciDeviceSpecToVmInstanceRequest) mgresult.Result {
	//POST zstack/v1/pci-device-specs/{pciSpecUuid}/vm-instances/{vmInstanceUuid}
	url := fmt.Sprintf("zstack/v1/pci-device-specs/%s/vm-instances/%s", params.PciSpecUuid, params.VmInstanceUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddPciDeviceSpecToVmInstanceFailed, err)
	}
	var result model.AddPciDeviceSpecToVmInstanceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将PCI设备规格从云主机卸载
func RemovePciDeviceSpecFromVmInstance(params model.RemovePciDeviceSpecFromVmInstanceRequest) mgresult.Result {
	//DELETE zstack/v1/pci-device-specs/{pciSpecUuid}/vm-instances/{vmInstanceUuid}
	url := fmt.Sprintf("zstack/v1/pci-device-specs/%s/vm-instances/%s", params.PciSpecUuid, params.VmInstanceUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemovePciDeviceSpecFromVmInstanceFailed, err)
	}
	var result model.RemovePciDeviceSpecFromVmInstanceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新PCI设备规格
func UpdatePciDeviceSpec(params model.UpdatePciDeviceSpecRequest) mgresult.Result {
	//PUT zstack/v1/pci-device-specs/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/pci-device-specs/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdatePciDeviceSpecFailed, err)
	}
	var result model.UpdatePciDeviceSpecResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取PCI设备规格候选列表
func GetPciDeviceSpecCandidates(params model.GetPciDeviceSpecCandidatesRequest) mgresult.Result {
	//GET zstack/v1/pci-device-specs/candidates
	url := fmt.Sprintf("zstack/v1/pci-device-specs/candidates")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetPciDeviceSpecCandidatesFailed, err)
	}
	var result model.GetPciDeviceSpecCandidatesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询PCI设备规格
func QueryPciDeviceSpec(params model.QueryPciDeviceSpecRequest) mgresult.Result {
	//GET zstack/v1/pci-device-specs
	//GET zstack/v1/pci-device-specs/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/pci-device-specs")
	} else {
		url = fmt.Sprintf("zstack/v1/pci-device-specs/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryPciDeviceSpecFailed, err)
	}
	var result model.QueryPciDeviceSpecResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询云主机与PCI设备规格的关联关系
func QueryVmInstancePciDeviceSpecRef(params model.QueryVmInstancePciDeviceSpecRefRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/pci-device-specs
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/pci-device-specs/{pciSpecUuid}
	var url string
	if params.PciSpecUuid == "" {
		url = fmt.Sprintf("zstack/v1/vm-instances/%s/pci-device-specs", params.VmInstanceUuid)
	} else {
		url = fmt.Sprintf("zstack/v1/vm-instances/%s/pci-device-specs/%s", params.VmInstanceUuid, params.PciSpecUuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVmInstancePciDeviceSpecRefFailed, err)
	}
	var result model.QueryVmInstancePciDeviceSpecRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//虚拟化切分支持SR-IOV的PCI设备
func GenerateSriovPciDevices(params model.GenerateSriovPciDevicesRequest) mgresult.Result {
	//PUT zstack/v1/pci-devices/{pciDeviceUuid}/actions
	url := fmt.Sprintf("zstack/v1/pci-devices/%s/actions", params.PciDeviceUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GenerateSriovPciDevicesFailed, err)
	}
	var result model.GenerateSriovPciDevicesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//虚拟化还原支持SR-IOV的PCI设备
func UngenerateSriovPciDevices(params model.UngenerateSriovPciDevicesRequest) mgresult.Result {
	//PUT zstack/v1/pci-devices/{pciDeviceUuid}/actions
	url := fmt.Sprintf("zstack/v1/pci-devices/%s/actions", params.PciDeviceUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UngenerateSriovPciDevicesFailed, err)
	}
	var result model.UngenerateSriovPciDevicesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//虚拟化切分支持VFIO_MDEV的PCI设备
func GenerateMdevDevices(params model.GenerateMdevDevicesRequest) mgresult.Result {
	//PUT zstack/v1/pci-devices/{pciDeviceUuid}/actions
	url := fmt.Sprintf("zstack/v1/pci-devices/%s/actions", params.PciDeviceUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GenerateMdevDevicesFailed, err)
	}
	var result model.GenerateMdevDevicesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//虚拟化还原支持VFIO_MDEV的PCI设备
func UngenerateMdevDevices(params model.UngenerateMdevDevicesRequest) mgresult.Result {
	//PUT zstack/v1/pci-devices/{pciDeviceUuid}/actions
	url := fmt.Sprintf("zstack/v1/pci-devices/%s/actions", params.PciDeviceUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UngenerateMdevDevicesFailed, err)
	}
	var result model.UngenerateMdevDevicesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询PCI设备切分出的MDEV设备
func QueryMdevDeviceSpec(params model.QueryMdevDeviceSpecRequest) mgresult.Result {
	//GET zstack/v1/mdev-device-specs
	//GET zstack/v1/mdev-device-specs/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/mdev-device-specs")
	} else {
		url = fmt.Sprintf("zstack/v1/mdev-device-specs/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryMdevDeviceSpecFailed, err)
	}
	var result model.QueryMdevDeviceSpecResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将PCI设备切分出的MDEV设备绑定到云主机
func AttachMdevDeviceToVm(params model.AttachMdevDeviceToVmRequest) mgresult.Result {
	//POST zstack/v1/mdev-devices/{mdevDeviceUuid}/vm-instances/{vmInstanceUuid}
	url := fmt.Sprintf("zstack/v1/mdev-devices/%s/vm-instances/%s", params.MdevDeviceUuid, params.VmInstanceUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachMdevDeviceToVmFailed, err)
	}
	var result model.AttachMdevDeviceToVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将PCI设备切分出的MDEV设备从云主机卸载
func DetachMdevDeviceFromVm(params model.DetachMdevDeviceFromVmRequest) mgresult.Result {
	//DELETE zstack/v1/mdev-devices/{mdevDeviceUuid}/vm-instances/{vmInstanceUuid}
	url := fmt.Sprintf("zstack/v1/mdev-devices/%s/vm-instances/%s", params.MdevDeviceUuid, params.VmInstanceUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachMdevDeviceFromVmFailed, err)
	}
	var result model.DetachMdevDeviceFromVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新PCI设备切分出的MDEV设备
func UpdateMdevDevice(params model.UpdateMdevDeviceRequest) mgresult.Result {
	//PUT zstack/v1/mdev-devices/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/mdev-devices/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateMdevDeviceFailed, err)
	}
	var result model.UpdateMdevDeviceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取可用的MDEV设备
func GetMdevDeviceCandidates(params model.GetMdevDeviceCandidatesRequest) mgresult.Result {
	//GET zstack/v1/mdev-devices/candidates
	url := fmt.Sprintf("zstack/v1/mdev-devices/candidates")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetMdevDeviceCandidatesFailed, err)
	}
	var result model.GetMdevDeviceCandidatesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询PCI设备切分出的MDEV设备
func QueryMdevDevice(params model.QueryMdevDeviceRequest) mgresult.Result {
	//GET zstack/v1/mdev-devices
	//GET zstack/v1/mdev-devices/{uuid}
	var url string
	if params.Uuid == "" {
		url = fmt.Sprintf("zstack/v1/mdev-devices")
	} else {
		url = fmt.Sprintf("zstack/v1/mdev-devices/%s", params.Uuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryMdevDeviceFailed, err)
	}
	var result model.QueryMdevDeviceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将MDEV设备规格加载到云主机
func AddMdevDeviceSpecToVmInstance(params model.AddMdevDeviceSpecToVmInstanceRequest) mgresult.Result {
	//POST zstack/v1/mdev-device-specs/{mdevSpecUuid}/vm-instances/{vmInstanceUuid}
	url := fmt.Sprintf("zstack/v1/mdev-device-specs/%s/vm-instances/%s", params.MdevSpecUuid, params.VMInstanceUUID)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AddMdevDeviceSpecToVmInstanceFailed, err)
	}
	var result model.AddMdevDeviceSpecToVmInstanceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将MDEV设备规格从云主机卸载
func RemoveMdevDeviceSpecFromVmInstance(params model.RemoveMdevDeviceSpecFromVmInstanceRequest) mgresult.Result {
	//DELETE zstack/v1/mdev-device-specs/{mdevSpecUuid}/vm-instances/{vmInstanceUuid}
	url := fmt.Sprintf("zstack/v1/mdev-device-specs/%s/vm-instances/%s", params.MdevSpecUuid, params.VMInstanceUUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.RemoveMdevDeviceSpecFromVmInstanceFailed, err)
	}
	var result model.RemoveMdevDeviceSpecFromVmInstanceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新MDEV设备规格
func UpdateMdevDeviceSpec(params model.UpdateMdevDeviceSpecRequest) mgresult.Result {
	//PUT zstack/v1/mdev-device-specs/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/mdev-device-specs/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateMdevDeviceSpecFailed, err)
	}
	var result model.UpdateMdevDeviceSpecParams
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取可用的MDEV设备规格
func GetMdevDeviceSpecCandidates(params model.GetMdevDeviceSpecCandidatesRequest) mgresult.Result {
	//GET zstack/v1/mdev-device-specs/candidates
	url := fmt.Sprintf("zstack/v1/mdev-device-specs/candidates")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetMdevDeviceSpecCandidatesFailed, err)
	}
	var result model.GetMdevDeviceSpecCandidatesResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询云主机与MDEV设备规格的关联关系
func QueryVmInstanceMdevDeviceSpecRef(params model.QueryVmInstanceMdevDeviceSpecRefRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/mdev-device-specs
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/mdev-device-specs/{mdevSpecUuid}
	var url string
	if params.MdevSpecUuid == "" {
		url = fmt.Sprintf("zstack/v1/vm-instances/%s/mdev-device-specs", params.VmInstanceUuid)
	} else {
		url = fmt.Sprintf("zstack/v1/vm-instances/%s/mdev-device-specs/%s", params.VmInstanceUuid, params.MdevSpecUuid)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryVmInstanceMdevDeviceSpecRefFailed, err)
	}
	var result model.QueryVmInstanceMdevDeviceSpecRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
