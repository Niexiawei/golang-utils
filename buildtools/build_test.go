package build

import (
	"fmt"
	"testing"
)

func TestIsDev(t *testing.T) {
	fmt.Println(IsProd())
	fmt.Println(VersionType)
}
