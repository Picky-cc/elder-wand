package auditors

import (
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/tasks/manager"
	"elder-wand/utils/log"
)

type PluginFreedomCashFlowAmountMatchingAuditor struct {
	manager.BasePluginBase
	PluginBaseMatching
}

type FreedomCashFlowAmountMatchingTemporaryContext struct {
	EconomicMatterList          []interface{}
	DeliveryTransactionItemList []interface{}
}

func (auditor *PluginFreedomCashFlowAmountMatchingAuditor) Process(agreement *models.Agreement, dataService *models.AgreementDataService) *errors.Error {
	log.Infof("[PluginFreedomCashFlowAmountMatchingAuditor] agreement: %d, service: %d start", agreement.ID, dataService.ID)
	defer log.Infof("[PluginFreedomCashFlowAmountMatchingAuditor] agreement: %d, service: %d end", agreement.ID, dataService.ID)
	return nil
}
