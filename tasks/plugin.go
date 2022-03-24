package tasks

import (
	"elder-wand/enums"
	"elder-wand/tasks/manager"
	base "elder-wand/tasks/plugin"
	"elder-wand/tasks/plugin/auditors"
)

var ServicePlugins = map[enums.ServicePlugin]base.ServicePlugin{
	enums.PluginFreedomCashFlowAmountMatching:   &auditors.PluginFreedomCashFlowAmountMatchingAuditor{},
	enums.PluginMultiAmountAndAccountNoMatching: &auditors.PluginMultiAmountAndAccountNoMatchingAuditor{},
	enums.PluginMultiAmountMatching:             &auditors.PluginMultiAmountMatchingAuditor{},
	enums.PluginMultiCashFlowPrintMatching:      &auditors.PluginMultiCashFlowPrintMatchingAuditor{},
	enums.PluginMultiReceiptsCashMatching:       &auditors.PluginMultiReceiptsCashMatchingAuditor{},
}

func init() {
	for key, val := range ServicePlugins {
		manager.PluginManager.SetPlugin(key, val)
	}
}
