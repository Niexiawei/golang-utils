package sqlutil

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_int(t *testing.T) {
	var a int8 = 12
	refV := reflect.ValueOf(a)
	t.Log(refV.Int())
	t.Log(strconv.FormatInt(refV.Int(), 10))
}
