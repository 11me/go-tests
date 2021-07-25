package arrays_test

import (
	"reflect"
	"testing"

	"github.com/11me/go-tests/arrays"
)

func TestSum(t *testing.T) {

	t.Run("Collection of any size array (slice)", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		got := arrays.Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("want %d, but got %d, given %v\n", want, got, numbers)
		}

	})

}

func TestSumAll(t *testing.T) {

	got := arrays.SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, but got %v", want, got)
	}

}
