package enums

type ServicePlugin int

const (
	PluginFreedomCashFlowAmountMatching   ServicePlugin = 1001 //自由匹配
	PluginMultiAmountAndAccountNoMatching ServicePlugin = 1002 //账户加金额一致匹配
	PluginMultiAmountMatching             ServicePlugin = 1003 //金额一致匹配
	PluginMultiCashFlowPrintMatching      ServicePlugin = 1004 //流水指纹匹配
	PluginMultiReceiptsCashMatching       ServicePlugin = 1005 //实收资金匹配
)

var ServicePluginList = []map[string]interface{}{
	{"label": "自由匹配", "value": PluginFreedomCashFlowAmountMatching},        //自由匹配
	{"label": "账户加金额一致匹配", "value": PluginMultiAmountAndAccountNoMatching}, //实收资金匹配
	{"label": "金额一致匹配", "value": PluginMultiAmountMatching},                //金额一致匹配
	{"label": "流水指纹匹配", "value": PluginMultiCashFlowPrintMatching},         //流水指纹匹配
	{"label": "实收资金匹配", "value": PluginMultiReceiptsCashMatching},          //实收资金匹配
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
