package auditors

import (
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/tasks/manager"
	"elder-wand/utils/log"
)

// 实收资金匹配
type PluginMultiReceiptsCashMatchingAuditor struct {
	manager.BasePluginBase
	PluginBaseMatching
}

func (auditor *PluginMultiReceiptsCashMatchingAuditor) Process(agreement *models.Agreement, dataService *models.AgreementDataService) *errors.Error {
	log.Infof("[PluginMultiReceiptsCashMatchingAuditor] agreement: %d, service: %d start", agreement.ID, dataService.ID)
	defer log.Infof("[PluginMultiReceiptsCashMatchingAuditor] agreement: %d, service: %d end", agreement.ID, dataService.ID)

	return nil
}
