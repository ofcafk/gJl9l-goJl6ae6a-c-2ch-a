package main

import (
	"fmt"

	converter "github.com/ofcafk/gJl9l-goJl6ae6a-c-2ch-a/convert"
)

func main() {

	cyrillicSentence := "Дайте ему стул!"
	latinSentence := "gaijme eMy cmyJl!"

	fmt.Println(converter.Convert(cyrillicSentence))
	fmt.Println(converter.Convert(latinSentence))
}
