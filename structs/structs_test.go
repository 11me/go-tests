package structs_test

import (
	"testing"

	"github.com/11me/go-tests/structs"
)

func TestPerimeter(t *testing.T) {
	rect := structs.Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, but want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rect := structs.Rectangle{10.0, 10.0}
	got := rect.Area()
	want := 100.0

	if got != want {
		t.Errorf("got %.2f, but want %.2f", got, want)
	}
}
