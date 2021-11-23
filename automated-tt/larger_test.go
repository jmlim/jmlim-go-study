package automated_tt

import (
	"fmt"
	"testing"
)

func TestFirstLarger(t *testing.T) {
	want := 2
	got := Larger(2, 1)
	if got != want {
		t.Error(errorString1(got, want))
	}
}

func TestSecondLarger(t *testing.T) {
	want := 8
	got := Larger(4, 8)
	if got != want {
		t.Error(errorString1(got, want))
	}
}

// helper 함수로 분리
func errorString1(got int, want int) string {
	return fmt.Sprintf("Larger(%d, %d) = \"%d\", want \"%d\"", got, want, got, want)
}
