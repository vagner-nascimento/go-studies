package architecture

import (
	"runtime"
	"testing"
)

func TestDependent(t *testing.T) {
	t.Parallel()

	if runtime.GOARCH == "amd64" { // It brings the
		t.Skip("Doesn't works on amd64") // It leaves the test, doesn't execute the next lines
	}

	t.Log("Architecture is valid") // It will be not executed if arch was amd64
}
