package enums

type AgreementLifeCycle int

const (
	AgreementLifeCycleActive   AgreementLifeCycle = 1 // 有效
	AgreementLifeCycleFinished AgreementLifeCycle = 2 // 已终止
	AgreementLifeCycleCancel   AgreementLifeCycle = 3 // 作废
)

type AgreementType int

const (
	MulCashFlowAccountAndAmountOfAT AgreementType = 1   //多银行流水账户加金额匹配
	MulCashFlowAmountOfAT           AgreementType = 2   //多银行流水金额匹配
	MulCashFlowPrintOfAT            AgreementType = 3   //多银行流水指纹匹配
	FreedomCashFlowAmountOfAT       AgreementType = 4   //流水自由匹配
	MulReceiptsCashOfAT             AgreementType = 5   //多实收资金匹配
	ClearCashOfAT                   AgreementType = 101 //实收资金流水
)
