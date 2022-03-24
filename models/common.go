package models

//type Version struct {
//	App     string
//	Version string
//}
//
//func (*Version) TableName() string  {
//	return "t_clearing_version"
//}

type Counter struct {
	Tag     string `gorm:"type:varchar(64);unique_index;not null"`
	Counter uint64 `gorm:"type:bigint(20) unsigned;not null"`
}

func (*Counter) TableName() string {
	return "t_clearing_counter"
}
