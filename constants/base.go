package constants

import (
	"elder-wand/utils/dbUtils"
)

const (
	InvalidID dbUtils.SFID = 0
)
const (
	InvalidType = 0
)

type DataExportType int

const (
	DataExportTypeLinrun DataExportType = 1
)
