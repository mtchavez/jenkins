package jenkins


import (
  "testing"
)

func TestOneAtATime(t *testing.T) {
  result := one_at_a_time("Apple")
  var expected uint32 = 884782484
  if result != expected {
    t.Errorf("one_at_a_time expected %v but got %v", expected, result)
  }
}
