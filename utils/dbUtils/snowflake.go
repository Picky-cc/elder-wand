package dbUtils

import (
	berrors "elder-wand/errors"
	"elder-wand/settings"
	"errors"
	"fmt"
	"strconv"

	uuid "github.com/satori/go.uuid"

	"github.com/bwmarrin/snowflake"
)

var sfNode *snowflake.Node

// SFID snowflakeID
type SFID uint64

// MarshalJSON returns a json byte array string of the snowflake ID.
func (f SFID) MarshalJSON() ([]byte, error) {
	buff := make([]byte, 0, 22)
	buff = append(buff, '"')
	buff = strconv.AppendInt(buff, int64(f), 10)
	buff = append(buff, '"')
	return buff, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (f *SFID) UnmarshalJSON(b []byte) error {
	if len(b) == 2 && b[0] == '"' && b[1] == '"' {
		*f = SFID(0)
		return nil
	}
	if len(b) < 3 || b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("invalid snowflake ID %q", string(b))
	}

	i, err := strconv.ParseUint(string(b[1:len(b)-1]), 10, 64)
	if err != nil {
		return err
	}

	*f = SFID(i)
	return nil
}

func (f *SFID) IsValid() bool {
	return *f != SFID(0)
}

func (f *SFID) ToString() string {
	return strconv.FormatUint(uint64(*f), 10)
}

func String2SFID(idStr string) (SFID, *berrors.Error) {
	i, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, berrors.NewValueError(err.Error())
	}
	return SFID(i), nil
}

func StringList2SFIDList(idStrList []string) ([]SFID, *berrors.Error) {
	idList := make([]SFID, len(idStrList))
	for idx, idStr := range idStrList {
		id, err := String2SFID(idStr)
		if err != nil {
			return nil, err
		}
		idList[idx] = id
	}
	return idList, nil
}

func SFIDList2StringList(idList []SFID) []string {
	idStrList := make([]string, len(idList))
	for idx, id := range idList {
		idStrList[idx] = id.ToString()
	}
	return idStrList
}

// Init 注意！！！同一机构下的不同实例一定要区分开，通过环境变量来配置
func Init() {
	snowflake.NodeBits = 5
	snowflake.StepBits = 17
	nodeID := getNodeID()
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		panic(err)
	}
	sfNode = node
}

func Generate() snowflake.ID {
	if sfNode == nil {
		panic(errors.New("snowflake node not inited"))
	}
	return sfNode.Generate()
}

func GenerateID() SFID {
	id := Generate()
	return SFID(id)
}

func GenerateInt64() uint64 {
	id := Generate()
	return uint64(id)
}

func GenerateString() string {
	id := GenerateID()
	return id.ToString()
}

func UUID() string {
	return uuid.NewV4().String()
}

func getNodeID() int64 {
	nodeID := settings.Config.EwNodeID
	if nodeID == 0 || nodeID > 31 {
		panic(fmt.Errorf("config `EwNodeID: (%d)` invalid", nodeID))
	}
	return int64(nodeID)
}
