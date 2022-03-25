package enums

type ServicePlugin int

const (
	PluginOne   ServicePlugin = 1001 //插件一
	PluginTwo   ServicePlugin = 1002 //插件二
	PluginThree ServicePlugin = 1003 //插件三

)

var ServicePluginList = []map[string]interface{}{
	{"label": "插件一", "value": PluginOne},
	{"label": "插件二", "value": PluginTwo},
	{"label": "插件三", "value": PluginThree},
}
var ServicePluginMap map[ServicePlugin]string

func init() {
	ServicePluginMap = make(map[ServicePlugin]string)
	for _, m := range ServicePluginList {
		value := m["value"].(ServicePlugin)
		label := m["label"].(string)
		ServicePluginMap[value] = label
	}
}
