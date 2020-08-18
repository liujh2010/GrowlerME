package growler

import (
	"fmt"
	"testing"
)

func TestConfigGet(t *testing.T) {
	fmt.Print(GetConfig())
}
