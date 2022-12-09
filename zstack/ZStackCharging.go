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

//创建资源价格
func CreateResourcePrice(params model.CreateResourcePriceRequest) models.Result {
	//POST zstack/v1/billings/prices
	url := fmt.Sprintf("zstack/v1/billings/prices")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateResourcePriceFailed, err.Error())
	}
	var result model.CreateResourcePriceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除资源价格
func DeleteResourcePrice(params model.DeleteResourcePriceRequest) models.Result {
	//DELETE zstack/v1/billings/prices/{uuid}
	url := fmt.Sprintf("zstack/v1/billings/prices/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteResourcePriceFailed, err.Error())
	}
	var result model.DeleteResourcePriceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改资源价格
func UpdateResourcePrice(params model.UpdateResourcePriceRequest) models.Result {
	//PUT zstack/v1/billings/prices/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/billings/prices/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateResourcePriceFailed, err.Error())
	}
	var result model.UpdateResourcePriceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询资源价格
func QueryResourcePrice(params model.QueryResourcePriceRequest) models.Result {
	//GET zstack/v1/billings/prices
	//GET zstack/v1/billing/prices/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/billings/prices")
	} else {
		url = fmt.Sprintf("zstack/v1/billings/prices/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryResourcePriceFailed, err.Error())
	}
	var result model.QueryResourcePriceResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//计算账户花费
func CalculateAccountSpending(params model.CalculateAccountSpendingRequest) models.Result {
	//PUT zstack/v1/billings/accounts/{accountUuid}/actions
	url := fmt.Sprintf("zstack/v1/billings/accounts/%s/actions", params.AccountUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.CalculateAccountSpendingFailed, err.Error())
	}
	var result model.CalculateAccountSpendingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询账户账单
func QueryAccountBilling(params model.QueryAccountBillingRequest) models.Result {
	//GET zstack/v1/billing/billings
	//GET zstack/v1/billing/billings/{id}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/billing/billings")
	} else {
		url = fmt.Sprintf("zstack/v1/billing/billings/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAccountBillingFailed, err.Error())
	}
	var result model.QueryAccountBillingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//创建计费价目
func CreatePriceTable(params model.CreatePriceTableRequest) models.Result {
	//POST zstack/v1/billings/price-tables
	url := fmt.Sprintf("zstack/v1/billings/price-tables")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreatePriceTableFailed, err.Error())
	}
	var result model.CreatePriceTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//删除计费价目
func DeletePriceTable(params model.DeletePriceTableRequest) models.Result {
	//DELETE zstack/v1/billings/price-tables/{uuid}
	url := fmt.Sprintf("zstack/v1/billings/price-tables/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeletePriceTableFailed, err.Error())
	}
	var result model.DeletePriceTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改计费价目
func UpdatePriceTable(params model.UpdatePriceTableRequest) models.Result {
	//PUT zstack/v1/billings/price-tables/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/billings/price-tables/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdatePriceTableFailed, err.Error())
	}
	var result model.UpdatePriceTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询计费价目
func QueryPriceTable(params model.QueryPriceTableRequest) models.Result {
	//GET zstack/v1/billings/price-tables
	//GET zstack/v1/billings/price-tables/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/billings/price-tables")
	} else {
		url = fmt.Sprintf("zstack/v1/billings/price-tables/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPriceTableFailed, err.Error())
	}
	var result model.QueryPriceTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//为账户项目指定计费价目
func AttachPriceTableToAccount(params model.AttachPriceTableToAccountRequest) models.Result {
	//POST zstack/v1/billings/price-tables/{tableUuid}/accounts/{accountUuid}
	url := fmt.Sprintf("zstack/v1/billings/price-tables/%s/accounts/%s", params.TableUuid, params.AccountUuid)

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.AttachPriceTableToAccountFailed, err.Error())
	}
	var result model.AttachPriceTableToAccountResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//取消账户项目关联的计费价目
func DetachPriceTableFromAccount(params model.DetachPriceTableFromAccountRequest) models.Result {
	//DELETE zstack/v1/billings/price-tables/{tableUuid}/accounts/{accountUuid}
	url := fmt.Sprintf("zstack/v1/billings/price-tables/%s/accounts/%s", params.TableUuid, params.AccountUuid)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DetachPriceTableFromAccountFailed, err.Error())
	}
	var result model.DetachPriceTableFromAccountResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//修改账户项目计费价目
func ChangeAccountPriceTableBinding(params model.ChangeAccountPriceTableBindingRequest) models.Result {
	//PUT zstack/v1/billings/price-tables/{tableUuid}/accounts/{accountUuid}
	url := fmt.Sprintf("zstack/v1/billings/price-tables/%s/accounts/%s", params.TableUuid, params.AccountUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeAccountPriceTableBindingFailed, err.Error())
	}
	var result model.ChangeAccountPriceTableBindingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//获得项目账户使用的计费价目
func GetAccountPriceTableRef(params model.GetAccountPriceTableRefRequest) models.Result {
	//GET zstack/v1/billings/price-tables/refs
	url := fmt.Sprintf("zstack/v1/billings/price-tables/refs")

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetAccountPriceTableRefFailed, err.Error())
	}
	var result model.GetAccountPriceTableRefResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//查询账户项目与计费价目的关联关系
func QueryAccountPriceTableRef(params model.QueryAccountPriceTableRefRequest) models.Result {
	//GET zstack/v1/accounts/price-tables/refs
	//GET zstack/v1/accounts/price-tables/refs/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/accounts/price-tables/refs")
	} else {
		url = fmt.Sprintf("zstack/v1/accounts/price-tables/refs/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryAccountPriceTableRefFailed, err.Error())
	}
	var result model.QueryPriceTableResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}

//生成账单
func GenerateAccountBilling(params model.GenerateAccountBillingRequest) models.Result {
	//PUT zstack/v1/billings/accounts/{accountUuid}/actions
	url := fmt.Sprintf("zstack/v1/billings/accounts/%s/actions")

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.GenerateAccountBillingFailed, err.Error())
	}
	var result model.GenerateAccountBillingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success(result)
}
