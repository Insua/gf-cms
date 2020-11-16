package orm

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

func IsNullOrEqual(field string, value interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if rv := reflect.ValueOf(value); rv.Kind() == reflect.Ptr && rv.IsNil() {
			return db.Where(fmt.Sprintf("%s IS NULL", field))
		}
		return db.Where(fmt.Sprintf("%s=(?)", field), value)
	}
}
