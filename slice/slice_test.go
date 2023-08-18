package slice

import "testing"

func TestDelByIndex(t *testing.T) {
	testSlice := []string{
		"index0",
		"index1",
		"index2",
		"index3",
	}
	DelByIndex(&testSlice, 2)
	t.Log(testSlice)
}

func TestDel(t *testing.T) {
	testSlice := []string{
		"index0",
		"index1",
		"index2",
		"index3",
	}
	Del(&testSlice, "index3")
	Del(&testSlice, "index4")
	t.Log(testSlice)
}
