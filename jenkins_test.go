package jenkins

import (
	"testing"
)

func TestHash(t *testing.T) {
	result := Hash("Apple")
	var expected uint32 = 884782484
	if result != expected {
		t.Errorf("one_at_a_time expected %v but got %v", expected, result)
	}
}
