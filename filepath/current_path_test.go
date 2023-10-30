package filepath

import "testing"

func TestGetCurrentAbPathByExecutable(t *testing.T) {
	t.Log(GetCurrentAbPathByExecutable())
}

func TestGetTmpDir(t *testing.T) {
	t.Log(GetTmpDir())
}
