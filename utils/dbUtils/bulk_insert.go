package dbUtils

import (
	ew_error "elder-wand/errors"
	"elder-wand/utils/tracer"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

func BulkInsertWithRowsAffected(db *gorm.DB, tableName string, objects []interface{}, chunkSize int, excludeColumns ...string) (int64, *ew_error.Error) {
	// Split records with specified size not to exceed Database parameter limit
	if len(objects) == 0 {
		return 0, nil
	}
	if len(tableName) == 0 {
		return 0, ew_error.NewParamError(fmt.Sprintf("tableName must not be null!"))
	}
	if chunkSize <= 0 {
		chunkSize = 500
	}
	var rows int64

	for _, objSet := range splitObjects(objects, chunkSize) {
		n, err := insertObjSet(db, tableName, objSet, false, excludeColumns...)
		if err != nil {
			return rows, ew_error.NewDBError(err.Error())
		} else {
			rows += n
		}
	}
	return rows, nil
}

// BulkInsert multiple records at once
// [objects]        Must be a slice of struct
// [chunkSize]      Number of records to insert at once.
//                  Embedding a large number of variables at once will raise an *errors.Error beyond the limit of prepared statement.
//                  Larger size will normally lead the better performance, but 2000 to 3000 is reasonable.
// [excludeColumns] Columns you want to exclude from insert. You can omit if there is no column you want to exclude.
// fork from https://github.com/t-tiger/gorm-bulk-insert
func BulkInsert(db *gorm.DB, tableName string, objects []interface{}, chunkSize int, excludeColumns ...string) *ew_error.Error {
	// Split records with specified size not to exceed Database parameter limit
	if len(objects) == 0 {
		return nil
	}
	if len(tableName) == 0 {
		return ew_error.NewParamError(fmt.Sprintf("tableName must not be null!"))
	}
	if chunkSize <= 0 {
		chunkSize = 500
	}
	for _, objSet := range splitObjects(objects, chunkSize) {
		if _, err := insertObjSet(db, tableName, objSet, false, excludeColumns...); err != nil {
			return ew_error.NewDBError(err.Error())
		}
	}
	return nil
}

// BulkInsertIgnore multiple records at once
// [objects]        Must be a slice of struct
// [chunkSize]      Number of records to insert at once.
//                  Embedding a large number of variables at once will raise an *errors.Error beyond the limit of prepared statement.
//                  Larger size will normally lead the better performance, but 2000 to 3000 is reasonable.
// [excludeColumns] Columns you want to exclude from insert. You can omit if there is no column you want to exclude.
// fork from https://github.com/t-tiger/gorm-bulk-insert
func BulkInsertIgnore(db *gorm.DB, tableName string, objects []interface{}, chunkSize int, excludeColumns ...string) *ew_error.Error {
	// Split records with specified size not to exceed Database parameter limit
	if len(objects) == 0 {
		return nil
	}
	if len(tableName) == 0 {
		return ew_error.NewParamError(fmt.Sprintf("tableName must not be null!"))
	}
	if chunkSize <= 0 {
		chunkSize = 500
	}
	for _, objSet := range splitObjects(objects, chunkSize) {
		if _, err := insertObjSet(db, tableName, objSet, true, excludeColumns...); err != nil {
			return ew_error.NewDBError(err.Error())
		}
	}
	return nil
}

func BulkInsertMulti(db *gorm.DB, tableObjectMapping map[string][]interface{}, chunkSize int, excludeColumns ...string) *ew_error.Error {
	for tableName, objects := range tableObjectMapping {
		if tableName != "" && len(objects) > 0 {
			if err := BulkInsert(db, tableName, objects, chunkSize, excludeColumns...); err != nil {
				return err
			}
		}
	}
	return nil
}

func insertObjSet(db *gorm.DB, tableName string, objects []interface{}, insertIgnore bool, excludeColumns ...string) (int64, error) {
	if len(objects) == 0 {
		return 0, nil
	}

	firstAttrs, err := extractMapValue(objects[0], excludeColumns)
	if err != nil {
		return 0, err
	}

	attrSize := len(firstAttrs)

	// Scope to eventually run SQL
	mainScope := db.NewScope(objects[0])
	// Store placeholders for embedding variables
	placeholders := make([]string, 0, attrSize)

	// Replace with database column name
	dbColumns := make([]string, 0, attrSize)
	for _, key := range sortedKeys(firstAttrs) {
		dbColumns = append(dbColumns, gorm.ToColumnName(key))
	}

	for _, obj := range objects {
		objAttrs, err := extractMapValue(obj, excludeColumns)
		if err != nil {
			return 0, err
		}

		// If object sizes are different, SQL statement loses consistency
		if len(objAttrs) != attrSize {
			return 0, errors.New("attribute sizes are inconsistent")
		}

		scope := db.NewScope(obj)

		// Append variables
		variables := make([]string, 0, attrSize)
		for _, key := range sortedKeys(objAttrs) {
			variables = append(variables, scope.AddToVars(objAttrs[key]))
		}

		valueQuery := "(" + strings.Join(variables, ", ") + ")"
		placeholders = append(placeholders, valueQuery)

		// Also append variables to mainScope
		mainScope.SQLVars = append(mainScope.SQLVars, scope.SQLVars...)
	}

	if len(tableName) <= 0 {
		tableName = mainScope.QuotedTableName()
	}

	var insertIgnoreOption string
	if insertIgnore {
		insertIgnoreOption = "IGNORE"
	}

	mainScope.Raw(fmt.Sprintf("INSERT %s INTO %s (%s) VALUES %s",
		insertIgnoreOption,
		tableName,
		strings.Join(dbColumns, ", "),
		strings.Join(placeholders, ", "),
	))

	retConn := tracer.GormExec(db, mainScope.SQL, mainScope.SQLVars...)

	return retConn.RowsAffected, retConn.Error
}

// Obtain columns and values required for insert from interface
func extractMapValue(value interface{}, excludeColumns []string) (map[string]interface{}, error) {
	kind := reflect.ValueOf(value).Kind()
	if kind != reflect.Struct {
		return nil, errors.New("value must be kind of Struct")
	}

	var attrs = map[string]interface{}{}

	for _, field := range (&gorm.Scope{Value: value}).Fields() {
		// Exclude relational record because it's not directly contained in database columns
		_, hasForeignKey := field.TagSettingsGet("FOREIGNKEY")

		if !containString(excludeColumns, field.Struct.Name) && field.StructField.Relationship == nil && !hasForeignKey &&
			!field.IsIgnored {
			if field.Struct.Name == "CreatedAt" || field.Struct.Name == "UpdatedAt" {
				attrs[field.DBName] = time.Now()
			} else if field.StructField.HasDefaultValue && field.IsBlank {
				// If default value presents and field is empty, assign a default value
				if val, ok := field.TagSettingsGet("DEFAULT"); ok {
					attrs[field.DBName] = val
				} else {
					attrs[field.DBName] = field.Field.Interface()
				}
			} else {
				attrs[field.DBName] = field.Field.Interface()
			}
		}
	}
	return attrs, nil
}

// Separate objects into several size
func splitObjects(objArr []interface{}, size int) [][]interface{} {
	var chunkSet [][]interface{}
	var chunk []interface{}

	for len(objArr) > size {
		chunk, objArr = objArr[:size], objArr[size:]
		chunkSet = append(chunkSet, chunk)
	}
	if len(objArr) > 0 {
		chunkSet = append(chunkSet, objArr[:])
	}

	return chunkSet
}

// Enable map keys to be retrieved in same order when iterating
func sortedKeys(val map[string]interface{}) []string {
	var keys []string
	for key := range val {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// Check if string value is contained in slice
func containString(s []string, value string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}
