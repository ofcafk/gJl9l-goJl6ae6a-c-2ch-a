package converter

import "testing"

func TestConvert(t *testing.T) {
	input := "Б"
	got := Convert(input)
	want := "6"

	if got != want {
		t.Errorf("missmatch. got: %q, want: %q", got, want)
	}
}
