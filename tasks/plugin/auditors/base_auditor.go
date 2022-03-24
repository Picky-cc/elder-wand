package auditors

import (
	"elder-wand/tasks/manager"
	"elder-wand/utils/dbUtils"
	"github.com/shopspring/decimal"
)

type PluginBaseMatching struct {
	manager.BasePluginBase
}

type MatchingResultContext struct {
	PlanID                dbUtils.SFID
	AssetAmount           decimal.Decimal
	CancelMatterID        dbUtils.SFID
	DeliveryTransactionID dbUtils.SFID
	DetailList            []MatchingResultDetail
	Full                  bool
}

type MatchingResultDetail struct {
	FlowPlanID dbUtils.SFID
	Amount     decimal.Decimal
}
