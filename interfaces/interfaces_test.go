package interfaces_test

import (
	"math"
	"testing"

	"github.com/11me/go-tests/interfaces"
	"github.com/11me/go-tests/structs"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape interfaces.Shape
		want  float64
	}{
		{structs.Rectangle{10.0, 10.0}, 100.0},
		{structs.Circle{10.0}, math.Pi * 10.0 * 10.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("got %.2f, want %.2f", got, tt.want)
		}
	}

}
