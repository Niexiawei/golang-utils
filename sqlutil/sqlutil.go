package sqlutil

import (
	"database/sql"
	"fmt"
	"reflect"
	"slices"
	"time"
)

func RowsScanSlice(rows *sql.Rows) (result []any, columns []string, err error) {
	columns, err = rows.Columns()
	if err != nil {
		return
	}
	for rows.Next() {
		res := make([]any, len(columns))
		for i := range columns {
			res[i] = new(any)
		}
		_ = rows.Scan(res...)
		result = append(result, res)
	}
	return
}

func MapToStruct(dst []any, columns []string, src any) {
	refType := reflect.TypeOf(src).Elem()
	refValue := reflect.ValueOf(src).Elem()
	for i := range refType.NumField() {
		column, ok := refType.Field(i).Tag.Lookup("column")
		columnType := refType.Field(i).Type.String()
		if !ok {
			continue
		}
		columnIdx := slices.Index(columns, column)
		if columnIdx < 0 {
			continue
		}
		result := dst[columnIdx]
		switch columnType {
		case "string":
			refValue.Field(i).SetString(anyTypeToString(reflect.ValueOf(result).Elem()))
		case "time.Time":
			var timeVal reflect.Value
			val := reflect.ValueOf(result).Elem().Interface()
			if val == nil {
				zeroVal := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
				timeVal = reflect.ValueOf(zeroVal)
			} else {
				timeVal = reflect.ValueOf(val)
			}

			refValue.Field(i).Set(timeVal)
		}
	}
}

func anyTypeToString(value reflect.Value) string {
	switch value.Interface().(type) {
	case float64:
		return fmt.Sprintf("%.6f", value.Interface().(float64))
	case string:
		return value.Interface().(string)
	case nil:
		return ""
	default:
		return ""
	}
}
