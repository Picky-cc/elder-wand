package auditors

import (
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/tasks/manager"
	"elder-wand/utils/log"
)

// 金额一致匹配
type PluginMultiAmountMatchingAuditor struct {
	manager.BasePluginBase
	PluginBaseMatching
}

type MultiAmountMatchingDataServiceParams struct {
	HostAccountNoList []string
}

type MultiAmountMatchingTemporaryContext struct {
	EconomicMatterList          []interface{}
	DeliveryTransactionItemList []interface{}
}

func (auditor *PluginMultiAmountMatchingAuditor) Process(agreement *models.Agreement, dataService *models.AgreementDataService) *errors.Error {
	log.Infof("[PluginMultiAmountMatchingAuditor] agreement: %d, service: %d start", agreement.ID, dataService.ID)
	defer log.Infof("[PluginMultiAmountMatchingAuditor] agreement: %d, service: %d end", agreement.ID, dataService.ID)

	return nil
}
