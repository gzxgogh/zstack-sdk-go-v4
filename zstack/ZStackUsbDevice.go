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

//获取云主机usb重定向开关状态
func GetVmUsbRedirect(params model.GetVmUsbRedirectRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{uuid}/usbredirect
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/usbredirect", params.UUID)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetVmUsbRedirectFailed, err)
	}
	var result model.GetVmUsbRedirectResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//设置云主机usb重定向开关
func SetVmUsbRedirect(params model.SetVmUsbRedirectRequest) mgresult.Result {
	//PUT zstack/v1/vm-instances/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.SetVmUsbRedirectFailed, err)
	}
	var result model.SetVmUsbRedirectResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//云主机加载物理机USB设备
func AttachUsbDeviceToVm(params model.AttachUsbDeviceToVmRequest) mgresult.Result {
	//POST zstack/v1/usb-device/usb-devices/{usbDeviceUuid}/attach
	url := fmt.Sprintf("zstack/v1/usb-device/usb-devices/%s/attach", params.UsbDeviceUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.AttachUsbDeviceToVmFailed, err)
	}
	var result model.AttachUsbDeviceToVmParams
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//将云主机挂载的USB设备卸载
func DetachUsbDeviceFromVm(params model.DetachUsbDeviceFromVmRequest) mgresult.Result {
	//POST zstack/v1/usb-device/usb-devices/{usbDeviceUuid}/detach
	url := fmt.Sprintf("zstack/v1/usb-device/usb-devices/%s/attach", params.UsbDeviceUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.DetachUsbDeviceFromVmFailed, err)
	}
	var result model.DetachUsbDeviceFromVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//获取USB透传候选列表
func GetUsbDeviceCandidatesForAttachingVm(params model.GetUsbDeviceCandidatesForAttachingVmRequest) mgresult.Result {
	//GET zstack/v1/vm-instances/{vmInstanceUuid}/candidate-usb-devices
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/candidate-usb-devices", params.VmInstanceUuid)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.GetUsbDeviceCandidatesForAttachingVmFailed, err)
	}
	var result model.GetUsbDeviceCandidatesForAttachingVmResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//查询USB设备
func QueryUsbDevice(params model.QueryUsbDeviceRequest) mgresult.Result {
	//GET zstack/v1/usb-device/usb-devices
	//GET zstack/v1/usb-device/usb-devices/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/usb-device/usb-devices")
	} else {
		url = fmt.Sprintf("zstack/v1/usb-device/usb-devices/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.QueryUsbDeviceFailed, err)
	}
	var result model.QueryUsbDeviceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}

//更新USB设备
func UpdateUsbDevice(params model.UpdateUsbDeviceRequest) mgresult.Result {
	//PUT zstack/v1/usb-device/usb-devices/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/usb-device/usb-devices/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return mgerr.ErrorResultWithErr(errcode.UpdateUsbDeviceFailed, err)
	}
	var result model.UpdateUsbDeviceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return mgresult.Success(result)
}
