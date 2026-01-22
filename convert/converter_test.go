package converter

import "testing"

// д -> g
func TestConvertToLatin(t *testing.T) {
	input := "Д"
	got := ConvertToLatin(input)
	want := "g"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestConvertToCyrillic(t *testing.T) {
	input := "lj"
	got := ConvertToCyrillic(input)
	want := "й"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// g -> д
// func TestConvertToCyrillic(t *testing.T) {
// 	input := "g"
// 	got := ConvertToLatin(input)
// 	want := TranslitMap[].value

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }// var latinWords string = "gaijme eMy cmyJl!"
// var cyrillicWords string = "Дайте ему стул!"
