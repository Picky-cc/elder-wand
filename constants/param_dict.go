package constants

// 参数类型
type ParamType int

const (
	ParamTypeAudit    ParamType = 1 // 审计协议
	ParamTypeClearing ParamType = 2 // 清算协议
)

type ParamKey string

const ()

type ParamValue string

const ()

//var AuditAgreementParamValues = map[ParamKey][]ParamValue{
//
//}
//
//var ClearingAgreementParamValues = map[ParamKey][]ParamValue{
//	ParamKeyFetchDebitCashFlowMaxID:  {ParamValueEmpty},
//	ParamKeyFetchCreditCashFlowMaxID: {ParamValueEmpty},
//}
