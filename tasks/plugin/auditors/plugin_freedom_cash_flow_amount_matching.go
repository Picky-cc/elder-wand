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

func (auditor *PluginFreedomCashFlowAmountMatchingAuditor) Process(task *models.Task, taskPlugin *models.TaskPlugin) *errors.Error {
	log.Infof("[PluginFreedomCashFlowAmountMatchingAuditor] task: %d, plugin: %d start", task.ID, taskPlugin.ID)
	defer log.Infof("[PluginFreedomCashFlowAmountMatchingAuditor] task: %d, plugin: %d end", task.ID, taskPlugin.ID)
	return nil
}
