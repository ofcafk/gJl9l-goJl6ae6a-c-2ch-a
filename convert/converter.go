package converter

func Convert(letter string) string {
	switch letter {
	case "А", "а":
		return "A, a"
	case "Б", "б":
		return "6"
	case "В", "в":
		return "B"
	case "Г", "г":
		return "r"
	case "Д", "д":
		return "g"
	case "Е", "е":
		return "E, e"
	case "Ж", "ж":
		return "lll"
	case "З", "з":
		return "3"
	case "И", "и":
		return "u"
	case "Й", "й":
		return "ij"
	case "К", "к":
		return "K, k"
	case "Л", "л":
		return "Jl"
	case "М", "м":
		return "M m"
	case "Н", "н":
		return "H"
	case "О", "о":
		return "O o"
	case "П", "п":
		return "n"
	case "Р", "р":
		return "P p"
	case "С", "с":
		return "C c"
	case "Т", "т":
		return "T m"
	case "У", "у":
		return "Y y"
	case "Ф", "ф":
		return "qp"
	case "Х", "х":
		return "X x"
	case "Ц", "ц":
		return "LL"
	case "Ч", "ч":
		return "4"
	case "Ш", "ш":
		return "LLl"
	case "Щ", "щ":
		return "LLL"
	case "Ь", "ь":
		return "b"
	case "Ы", "ы":
		return "bl"
	case "Э", "э":
		return "9"
	case "Ю", "ю":
		return "lO"
	case "Я", "я":
		return "9l"

	default:
		return letter
	}
}

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
