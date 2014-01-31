package jenkins

import (
	"testing"
)

func TestHash(t *testing.T) {
	result := Hash([]byte("Apple"))
	var expected uint32 = 884782484
	if result != expected {
		t.Errorf("Hash expected %v but got %v", expected, result)
	}
}
