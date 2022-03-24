package constants

type AgreementType int

const (
	AgreementTypeClearing AgreementType = 1 // 清算
	AgreementTypeAudit    AgreementType = 2 // 审计
	AgreementTypeDeduct   AgreementType = 3 // 扣款
)

type AgreementLifeCycle int

const (
	AgreementLifeCycleActive   AgreementLifeCycle = 1 // 有效
	AgreementLifeCycleFinished AgreementLifeCycle = 2 // 已完成
	AgreementLifeCycleAborted  AgreementLifeCycle = 3 // 终止
)

type AgreementAccountLifeCycle int

const (
	AgreementAccountLifeCycleActive   AgreementAccountLifeCycle = 1 // 有效
	AgreementAccountLifeCycleCanceled AgreementAccountLifeCycle = 2 // 作废
)

type EntrustingPartyType int

const (
	EntrustingPartyTypeFinancialInstitution EntrustingPartyType = 1 // 金融机构（信托机构）
	EntrustingPartyTypePaymentInstitution   EntrustingPartyType = 2 // 三方支付机构
	EntrustingPartyTypeAssetsSide           EntrustingPartyType = 3 // 资产方
	EntrustingPartyTypeInsuranceInstitution EntrustingPartyType = 4 // 保险机构
)

type AgreementBusinessType int

const (
	ABTypeBusinessTransfer              AgreementBusinessType = 101 // 委托转付
	ABTypeThirdPartyClearing            AgreementBusinessType = 102 // 三方资金清算-线上扣款
	ABTypeActiveRepayment               AgreementBusinessType = 103 // 主动还款
	ABTypeBusinessDeduct                AgreementBusinessType = 104 // 商户代扣
	ABTypeCashFlowAccount               AgreementBusinessType = 105 // 流水日记账
	ABTypeFundAccount                   AgreementBusinessType = 106 // 资金账
	ABTypeEasyPayDeduct                 AgreementBusinessType = 107 // 快捷支付
	ABTypeBusinessPay                   AgreementBusinessType = 108 // 商户代偿
	ABTypeRefundRepayment               AgreementBusinessType = 109 // 退款还款
	ABTypeRepurchaseRepayment           AgreementBusinessType = 110 // 回购还款
	ABTypeOnlineDeductWithOnlinePay     AgreementBusinessType = 111 // 线上代扣-在线支付
	ABTypeMerchantThirdPartyClearing    AgreementBusinessType = 112 // 三方资金清算-商户代扣
	ABTypeOnlineDeduct                  AgreementBusinessType = 113 // 线上代扣
	ABTypeReceivablesPayable            AgreementBusinessType = 114 // 应收应付账
	ABTypeAccountingDocument            AgreementBusinessType = 115 // 会计凭证
	ABTypeSyncRemittance                AgreementBusinessType = 116 // 放款同步
	ABTypeRolloverContract              AgreementBusinessType = 117 // 转分期
	ABTypeCollectOnDelivery             AgreementBusinessType = 118 // 集合代收
	ABTypeCreditAssignment              AgreementBusinessType = 119 // 债权转让
	ABTypeCreditAssignmentCommonAsset   AgreementBusinessType = 120 // 债权转让标准资产账
	ABTypeRefundAmortization            AgreementBusinessType = 121 // 退款摊销(度小满场景,量化派)两个场景不一样，量化派不结转
	ABTypeOriginalConditionDistribution AgreementBusinessType = 122 // 原状分配
	ABTypeOffBalanceSheetPayment        AgreementBusinessType = 123 // 表外支付
	ABTypeActualReceiptAsset            AgreementBusinessType = 124 // 实收资金
)

var BusinessTypeList = []map[string]interface{}{
	{"label": "委托转付", "value": ABTypeBusinessTransfer},
	{"label": "三方资金清算-线上扣款", "value": ABTypeThirdPartyClearing},
	{"label": "主动还款", "value": ABTypeActiveRepayment},
	{"label": "商户代扣", "value": ABTypeBusinessDeduct},
	{"label": "流水日记账", "value": ABTypeCashFlowAccount},
	{"label": "资金账", "value": ABTypeFundAccount},
	{"label": "快捷支付", "value": ABTypeEasyPayDeduct},
	{"label": "商户代偿", "value": ABTypeBusinessPay},
	{"label": "退款还款", "value": ABTypeRefundRepayment},
	{"label": "回购还款", "value": ABTypeRepurchaseRepayment},
	{"label": "线上代扣", "value": ABTypeOnlineDeduct},
	{"label": "线上代扣-在线支付", "value": ABTypeOnlineDeductWithOnlinePay},
	{"label": "三方资金清算-商户代扣", "value": ABTypeMerchantThirdPartyClearing},
	{"label": "应收应付账", "value": ABTypeReceivablesPayable},
	{"label": "会计凭证", "value": ABTypeAccountingDocument},
	{"label": "放款同步", "value": ABTypeSyncRemittance},
	{"label": "转分期", "value": ABTypeRolloverContract},
	{"label": "集合代收", "value": ABTypeCollectOnDelivery},
	{"label": "债权转让", "value": ABTypeCreditAssignment},
	{"label": "债权转让标准资产账", "value": ABTypeCreditAssignmentCommonAsset},
	{"label": "退款摊销(度小满场景,量化派)两个场景不一样，量化派不结转", "value": ABTypeRefundAmortization},
	{"label": "原状分配", "value": ABTypeOriginalConditionDistribution},
	{"label": "表外支付", "value": ABTypeOffBalanceSheetPayment},
	{"label": "实收资金", "value": ABTypeActualReceiptAsset},
}

// 线上业务
var OnlineBusinessTypeList = []interface{}{ABTypeThirdPartyClearing, ABTypeMerchantThirdPartyClearing}

// 线下业务
var OfflineBusinessTypeList = []interface{}{ABTypeBusinessTransfer, ABTypeActiveRepayment, ABTypeBusinessPay, ABTypeRepurchaseRepayment, ABTypeCreditAssignment}

// 运营账业务
var OperationBusinessTypeList = []AgreementBusinessType{
	ABTypeThirdPartyClearing,
	ABTypeMerchantThirdPartyClearing,
	ABTypeBusinessTransfer,
	ABTypeActiveRepayment,
	ABTypeBusinessPay,
	ABTypeRepurchaseRepayment}

// 资金账业务
var FundBusinessTypeList = []AgreementBusinessType{ABTypeFundAccount}

// 流水账
var CashFlowBusinessTypeList = []AgreementBusinessType{ABTypeCashFlowAccount}

// 应收应付账
var RecPayBusinessTypeList = []AgreementBusinessType{ABTypeReceivablesPayable}

type DataScope int

const (
	DataScopeUnknown                                        DataScope = 0   // 未知
	DataScopeFundAccountRemittance                          DataScope = 101 // 资金账-放款
	DataScopeFundAccountCashFlow                            DataScope = 102 // 资金账-流水日记账
	DataScopeFundAccountOpsTDelivery                        DataScope = 103 // 资金账-运营账交割事务
	DataScopeFundAccountReserveFund                         DataScope = 104 // 资金账-外拨打款
	DataScopeFundAccountAssertTDelivery                     DataScope = 105 // 资金账-资产账交割事务
	DataScopeFundAccountPreparationTDelivery                DataScope = 106 // 资金账-计提交割事务
	DataScopeCollectOnDelivery                              DataScope = 107 // 资金账-集合代收
	DataScopeFundAccountActualReceiptAsset                  DataScope = 108 // 资金账-实收资金
	DataScopeCashFlowAccountCashFlow                        DataScope = 201 // 流水账-流水日记账
	DataScopeCashFlowAccountActualReceiptAsset              DataScope = 202 // 流水账-实收资金
	DataScopeOperationAccountCashFlow                       DataScope = 301 // 运营账-流水日记账
	DataScopeOperationAccountTDeliveryBusinessTransfer      DataScope = 302 // 运营账-交割事务-委托转付
	DataScopeOperationAccountTDeliveryBusinessPay           DataScope = 303 // 运营账-交割事务-商户代偿
	DataScopeOperationAccountTDeliveryRefundRepayment       DataScope = 304 // 运营账-交割事务-退款还款
	DataScopeOperationAccountTDeliveryMerchantDeduct        DataScope = 305 // 运营账-交割事务-三方代扣
	DataScopeOperationAccountTDeliveryRepurchase            DataScope = 306 // 运营账-交割事务-回购
	DataScopeDeductAccountEasyDeduct                        DataScope = 307 // 运营账-扣款-线上快捷
	DataScopeDeductAccountOnlineDeduct                      DataScope = 308 // 运营账-扣款-线上代扣
	DataScopeDeductAccountMerchantDeduct                    DataScope = 309 // 运营账-扣款-商户代扣
	DataScopeDeductAccountOnlinePayment                     DataScope = 310 // 运营账-扣款-在线支付
	DataScopeOnlineDeductRepaymentOrderSharePaymentType     DataScope = 311 // 运营账-扣款-分账支付
	DataScopeOnlineDeductRepaymentOrderPolymericPaymentType DataScope = 312 // 运营账-扣款-集合支付
	DataScopeOperationAccountTDeliveryOnlineDeductClear     DataScope = 313 // 运营账-交割事务-线上扣款清算
	DataScopeOperationAccountTDeliveryActiveRepayment       DataScope = 314 // 运营账-交割事务-主动还款
	DataScopeDeductAccountDeductInAdvance                   DataScope = 315 // 运营账-扣款-预收
	DataScopeRolloverContractRepayment                      DataScope = 316 // 运营账-交割事务-转分期还款
	DataScopeOperationAccountCreditAssignment               DataScope = 317 // 运营帐-债权转让
	DataScopeRefundAmortization                             DataScope = 318 // 运营帐-退款摊销
	DataScopeOriginalConditionDistribution                  DataScope = 319 // 运营帐-原状分配
	DataScopeOffBalanceSheetRepayment                       DataScope = 320 // 运营帐-表外支付
	DataScopeOperationAccountActualReceiptAsset             DataScope = 321 // 运营账-实收资金
	DataScopeOperationAccountCreditAssignmentReceipt        DataScope = 322 // 运营帐-债权转让应收账
	DataScopeAccountingDocumentsAsset                       DataScope = 401 // 会计凭证-应收应付账
	DataScopeAccountingDocumentsFund                        DataScope = 402 // 会计凭证-资金账
)

var RepaymentWayStrToAgreementBusinessTypeMap = map[string][]AgreementBusinessType{
	"1000": []AgreementBusinessType{ABTypeOnlineDeduct},
	"2001": []AgreementBusinessType{ABTypeBusinessTransfer},
	"2002": []AgreementBusinessType{},
	"2003": []AgreementBusinessType{ABTypeRepurchaseRepayment},
	"2004": []AgreementBusinessType{ABTypeBusinessTransfer},
	"2005": []AgreementBusinessType{ABTypeBusinessTransfer, ABTypeBusinessPay},
	"3001": []AgreementBusinessType{ABTypeActiveRepayment},
	"3002": []AgreementBusinessType{ABTypeActiveRepayment},
	"4001": []AgreementBusinessType{ABTypeBusinessDeduct},
	"5001": []AgreementBusinessType{ABTypeEasyPayDeduct},
	"6001": []AgreementBusinessType{ABTypeOnlineDeductWithOnlinePay},
	"7001": []AgreementBusinessType{},
	"8001": []AgreementBusinessType{},
	"2006": []AgreementBusinessType{},
}

var RepaymentWayStrToAgreementBusinessTypeMapV1 = map[string][]AgreementBusinessType{
	"1000": []AgreementBusinessType{ABTypeOnlineDeduct},
	"2001": []AgreementBusinessType{ABTypeBusinessTransfer},
	"2002": []AgreementBusinessType{},
	"2003": []AgreementBusinessType{ABTypeRepurchaseRepayment},
	"2004": []AgreementBusinessType{ABTypeBusinessTransfer},
	"2005": []AgreementBusinessType{ABTypeBusinessPay},
	"3001": []AgreementBusinessType{ABTypeActiveRepayment},
	"3002": []AgreementBusinessType{ABTypeActiveRepayment},
	"4001": []AgreementBusinessType{ABTypeBusinessDeduct},
	"5001": []AgreementBusinessType{ABTypeEasyPayDeduct},
	"6001": []AgreementBusinessType{ABTypeOnlineDeductWithOnlinePay},
	"7001": []AgreementBusinessType{},
	"8001": []AgreementBusinessType{},
	"2006": []AgreementBusinessType{},
}
