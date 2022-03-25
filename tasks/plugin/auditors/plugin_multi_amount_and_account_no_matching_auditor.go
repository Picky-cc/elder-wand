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

func (auditor *PluginMultiAmountAndAccountNoMatchingAuditor) Process(task *models.Task, taskPlugin *models.TaskPlugin) *errors.Error {
	log.Infof("[PluginMultiAmountAndAccountNoMatchingAuditor] task: %d, plugin: %d start", task.ID, taskPlugin.ID)
	defer log.Infof("[PluginMultiAmountAndAccountNoMatchingAuditor] task: %d, plugin: %d end", task.ID, taskPlugin.ID)

	return nil
}
