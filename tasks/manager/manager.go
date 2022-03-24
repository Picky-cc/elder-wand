package manager

import (
	"elder-wand/enums"
	base "elder-wand/tasks/plugin"
	"elder-wand/utils"
)

type pluginManager struct {
	servicePlugins map[enums.ServicePlugin]base.ServicePlugin
}

func (m *pluginManager) init() {
	m.servicePlugins = make(map[enums.ServicePlugin]base.ServicePlugin, 0)
}

func (m *pluginManager) SetPlugin(service enums.ServicePlugin, plugin base.ServicePlugin) {
	m.servicePlugins[service] = plugin
}

func (m *pluginManager) GetPlugin(service enums.ServicePlugin) base.ServicePlugin {
	plugin, ok := m.servicePlugins[service]
	if !ok {
		return nil
	}
	return utils.ReflectNewStruct(plugin).(base.ServicePlugin)
}

var PluginManager pluginManager

func init() {
	PluginManager.init()
}
