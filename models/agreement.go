package models

import (
	"elder-wand/constants"
	"elder-wand/utils/dbUtils"
)

type Agreement struct {
	BaseModel
}

func (*Agreement) TableName() string {
	return "t_clearing_agreement"
}

// 清算系统-协议工作参数
type AgreementAppendix struct {
	BaseModel
	AgreementID    dbUtils.SFID         `gorm:"index;not null"`            // 协议id
	ParamDictID    dbUtils.SFID         `gorm:"index;not null"`            // 参数对应的说明id
	AttributeKey   constants.ParamKey   `gorm:"type:varchar(64);not null"` // 参数名
	AttributeValue constants.ParamValue `gorm:"type:varchar(64);not null"` // 参数值
}

func (*AgreementAppendix) TableName() string {
	return "t_clearing_agreement_appendix"
}

// 清算系统-协议集合（所属同个信托实体的，属于同个集合）
type AgreementSet struct {
	BaseModel
	OwnerIdentity       string                  `gorm:"type:varchar(64);index;not null"`                                  // 通常是信托uuid
	AgreementType       constants.AgreementType `gorm:"type:varchar(64);not null"`                                        // 协议类型
	AgreementID         dbUtils.SFID            `gorm:"unique_index;not null"`                                            // 协议id
	FinancialContractNo string                  `gorm:"type:varchar(64);index:idx_t_clearing_agreement_set_project_code"` // 信托产品代码
}

func (*AgreementSet) TableName() string {
	return "t_clearing_agreement_set"
}
