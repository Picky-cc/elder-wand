package auditors

import (
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/tasks/manager"
	"elder-wand/utils/log"
)

// 流水指纹匹配
type PluginMultiCashFlowPrintMatchingAuditor struct {
	manager.BasePluginBase
	PluginBaseMatching
}

type MultiCashFlowPrintMatchingTemporaryContext struct {
	EconomicMatterList          []interface{}
	DeliveryTransactionItemList []interface{}
}

func (auditor *PluginMultiCashFlowPrintMatchingAuditor) Process(agreement *models.Agreement, dataService *models.AgreementDataService) *errors.Error {
	log.Infof("[PluginMultiCashFlowPrintMatchingAuditor] agreement: %d, service: %d start", agreement.ID, dataService.ID)
	defer log.Infof("[PluginMultiCashFlowPrintMatchingAuditor] agreement: %d, service: %d end", agreement.ID, dataService.ID)

	return nil
}
