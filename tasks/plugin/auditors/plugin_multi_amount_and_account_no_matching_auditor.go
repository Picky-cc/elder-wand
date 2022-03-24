package auditors

import (
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/tasks/manager"
	"elder-wand/utils/log"
)

// 账户加金额一致匹配
type PluginMultiAmountAndAccountNoMatchingAuditor struct {
	manager.BasePluginBase
	PluginBaseMatching
}

type MultiAmountAndAccountNoMatchingTemporaryContext struct {
	EconomicMatterList          []interface{}
	DeliveryTransactionItemList []interface{}
}

func (auditor *PluginMultiAmountAndAccountNoMatchingAuditor) Process(agreement *models.Agreement, dataService *models.AgreementDataService) *errors.Error {
	log.Infof("[PluginMultiAmountAndAccountNoMatchingAuditor] agreement: %d, service: %d start", agreement.ID, dataService.ID)
	defer log.Infof("[PluginMultiAmountAndAccountNoMatchingAuditor] agreement: %d, service: %d end", agreement.ID, dataService.ID)

	return nil
}
