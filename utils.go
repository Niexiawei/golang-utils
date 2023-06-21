package golangutils

import (
	"net/http"
	"os"
	"reflect"
)

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}

func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

func IsBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func CookiesMerge(old, new []*http.Cookie) []*http.Cookie {
	newMap := map[string]*http.Cookie{}
	oldMap := map[string]*http.Cookie{}
	var resCookies []*http.Cookie
	for idx, v := range new {
		newMap[v.Name] = new[idx]
	}
	for idx, v := range old {
		oldMap[v.Name] = old[idx]
	}

	for k, v := range newMap {
		oldMap[k] = v
	}

	for _, v := range oldMap {
		resCookies = append(resCookies, v)
	}

	return resCookies
}
