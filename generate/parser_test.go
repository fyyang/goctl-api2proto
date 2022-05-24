package generate

import (
	"strings"
	"testing"
)

var s = "getUserInfo"

func TestUp(t *testing.T) {
	t.Log(strings.ToUpper(string(s[0])))
}
