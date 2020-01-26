package mymath

import (
	"testing"
)

// On console, is possible to run all tests inside the tests folder with this command: go test ./... (... means all other folders and subfolders)
// You can run a single file as well: go test math/math_test.go

const errorMessage = "Expected value %v is different to the calculated value %v"

func TestAvarage(t *testing.T) {
	t.Parallel() // It means that this test will run parallel with other tests which use t.Parallel too
	expectedVal := 7.28
	val := Avarage(7.2, 9.9, 6.1, 5.9)

	if expectedVal != val {
		t.Errorf(errorMessage, expectedVal, val)
	}
}
