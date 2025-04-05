package gogeBasis

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	warning := ReturnWarningForTest()
	fmt.Println(warning.Warning())
}
