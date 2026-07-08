package maputil

import "reflect"

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))

	var i int
	for _, v := range m {
		values[i] = v
		i++
	}

	return values
}

// StructToMap 将 struct 中带 db 标签的字段转为 map[string]any。
// 指针字段为 nil 时跳过；非指针字段若有 db 标签直接写入。
// 用法：UpdateXxxReq 字段设为指针并加 `db:"column_name"` tag，调用此函数得到只含非 nil 字段的 updates map。
func StructToMap(src any) map[string]any {
	result := make(map[string]any)
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		col := t.Field(i).Tag.Get("db")
		if col == "" || col == "-" {
			continue
		}
		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				continue
			}
			result[col] = field.Elem().Interface()
		} else {
			result[col] = field.Interface()
		}
	}
	return result
}
