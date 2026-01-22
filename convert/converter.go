package converter

import (
	"strings"
	"unicode"
)

var translitMap = map[string]string{
	"А": "A", "а": "a",
	"Б": "6", "б": "6",
	"В": "B", "в": "B",
	"Г": "r", "г": "r",
	"Д": "g", "д": "g",
	"Е": "E", "е": "e",
	"Ж": "lll", "ж": "lll",
	"З": "3", "з": "3",
	"И": "u", "и": "u",
	"Й": "ij", "й": "ij",
	"К": "K", "к": "k",
	"Л": "Jl", "л": "Jl",
	"М": "M", "м": "M",
	"Н": "H", "н": "H",
	"О": "O", "о": "o",
	"П": "n", "п": "n",
	"Р": "P", "р": "p",
	"С": "C", "с": "c",
	"Т": "T", "т": "m",
	"У": "Y", "у": "y",
	"Ф": "qp", "ф": "qp",
	"Х": "X", "х": "x",
	"Ц": "LL", "ц": "LL",
	"Ч": "4", "ч": "4",
	"Ш": "LLl", "ш": "LLl",
	"Щ": "LLL", "щ": "LLL",
	"Ь": "b", "ь": "b",
	"Ы": "bl", "ы": "bl",
	"Э": "9", "э": "9",
	"Ю": "lO", "ю": "lO",
	"Я": "9l", "я": "9l",
}

func ConvertToLatin(cyrillic string) string {
	if val, ok := translitMap[cyrillic]; ok {
		return val
	}
	return cyrillic
}

func ConvertToCyrillic(latin string) string {
	if _, ok := translitMap[latin]; ok {
		return translitMap[latin]
	}
	return latin
}

func Convert(words string) string {
	var result string
	var builder strings.Builder

	for _, w := range words {
		if unicode.Is(unicode.Latin, w) {
			builder.WriteString(ConvertToCyrillic(string(w)))
		} else if unicode.Is(unicode.Cyrillic, w) {
			builder.WriteString(ConvertToLatin(string(w)))
		} else {
			builder.WriteString(string(w))
		}
	}

	result = builder.String()
	return result
}

// func Convert(letter string) string {
// 	switch letter {
// 	case "А":
// 		return "A"
// 	case "а":
// 		return "a"
// 	case "Б":
// 		return "6"
// 	case "б":
// 		return "6"
// 	case "В":
// 		return "B"
// 	case "в":
// 		return "B"
// 	case "Г":
// 		return "r"
// 	case "г":
// 		return "r"
// 	case "Д":
// 		return "g"
// 	case "д":
// 		return "g"
// 	case "Е":
// 		return "E"
// 	case "е":
// 		return "e"
// 	case "Ж":
// 		return "lll"
// 	case "ж":
// 		return "lll"
// 	case "З":
// 		return "3"
// 	case "з":
// 		return "3"
// 	case "И":
// 		return "u"
// 	case "и":
// 		return "u"
// 	case "Й":
// 		return "ij"
// 	case "й":
// 		return "ij"
// 	case "К":
// 		return "K"
// 	case "к":
// 		return "k"
// 	case "Л":
// 		return "Jl"
// 	case "л":
// 		return "Jl"
// 	case "М":
// 		return "M"
// 	case "м":
// 		return "M"
// 	case "Н":
// 		return "H"
// 	case "н":
// 		return "H"
// 	case "О":
// 		return "O"
// 	case "о":
// 		return "o"
// 	case "П":
// 		return "n"
// 	case "п":
// 		return "n"
// 	case "Р":
// 		return "P"
// 	case "р":
// 		return "p"
// 	case "С":
// 		return "C"
// 	case "с":
// 		return "c"
// 	case "Т":
// 		return "T"
// 	case "т":
// 		return "m"
// 	case "У":
// 		return "Y"
// 	case "у":
// 		return "y"
// 	case "Ф":
// 		return "qp"
// 	case "ф":
// 		return "qp"
// 	case "Х":
// 		return "X"
// 	case "х":
// 		return "x"
// 	case "Ц":
// 		return "LL"
// 	case "ц":
// 		return "LL"
// 	case "Ч":
// 		return "4"
// 	case "ч":
// 		return "4"
// 	case "Ш":
// 		return "LLl"
// 	case "ш":
// 		return "LLl"
// 	case "Щ":
// 		return "LLL"
// 	case "щ":
// 		return "LLL"
// 	case "Ь":
// 		return "b"
// 	case "ь":
// 		return "b"
// 	case "Ы":
// 		return "bl"
// 	case "ы":
// 		return "bl"
// 	case "Э":
// 		return "9"
// 	case "э":
// 		return "9"
// 	case "Ю":
// 		return "lO"
// 	case "ю":
// 		return "lO"
// 	case "Я":
// 		return "9l"
// 	case "я":
// 		return "9l"
// 	default:
// 		return letter
// 	}
// }

// A, a = A, a
// Б, б = 6
// В, в = B
// Г, г = r
// Д д = g
// Е,е = E, e
// Ж ж = lll
// З з = 3
// И и = u
// Й й = ij
// К к = K, k
// Л л = Jl
// М м = M m
// Н н = H
// О о = O o
// П п = n
// Р р = P p
// С с = C c
// Т т = T m
// У у = Y y
// Ф ф = qp
// Х х = X x
// Ц ц = LL
// Ч ч = 4
// Ш ш = LLl
// Щ щ = LLL
// Ь ь = b
// Ъ ъ =
// Ы ы = bl
// Э э = 9
// Ю ю = lO
// Я я = 9l
