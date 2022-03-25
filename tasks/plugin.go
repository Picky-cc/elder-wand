package tasks

import (
	"elder-wand/enums"
	"elder-wand/tasks/manager"
	base "elder-wand/tasks/plugin"
	"elder-wand/tasks/plugin/auditors"
)

var ServicePlugins = map[enums.ServicePlugin]base.ServicePlugin{
	enums.PluginOne: &auditors.PluginFreedomCashFlowAmountMatchingAuditor{},
	enums.PluginTwo: &auditors.PluginMultiAmountAndAccountNoMatchingAuditor{},
}

func init() {
	for key, val := range ServicePlugins {
		manager.PluginManager.SetPlugin(key, val)
	}
}
