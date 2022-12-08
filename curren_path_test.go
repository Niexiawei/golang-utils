package utils

import "testing"

func Test_getTemp(t *testing.T) {
	t.Log(getTmpDir())
}

func Test_getCurrentAbPathByExecutable(t *testing.T) {
	t.Log(getCurrentAbPathByExecutable())
}

func Test_getCurrentAbPathByCaller(t *testing.T) {
	t.Log(getCurrentAbPathByCaller())
}

func TestCurrentAbPath(t *testing.T) {
	t.Log(CurrentAbPath())
}
