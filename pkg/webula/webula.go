package webula

import (
	"strings"
	"unicode/utf8"

	"github.com/microcosm-cc/bluemonday"
)

const (
	EmptyString       = ""
	Space             = " "
	Underline         = "_"
	Dot               = "."
	NewLine           = "\n"
	CarriageReturn    = "\r"
	Tab               = "\t"
	Comma             = ","
	Colon             = ":"
	Semicolon         = ";"
	Gradus            = "°"
	SingleQuote       = "'"
	DoubleQuote       = "\""
	QuoteLeft         = "«"
	QuoteRight        = "»"
	Mult              = "*"
	Div               = "/"
	Plus              = "+"
	Minus             = "-"
	Equal             = "="
	Percent           = "%"
	Number            = "№"
	Exclamation       = "!"
	RoundBracketLeft  = "("
	RoundBracketRight = ")"
	HTMLSpaceInUtf8   = "\xc2\xa0"
	HTMLSpace         = "&nbsp;"
	HTMLMDash         = "&mdash;"
)

// Сколько Символов в строке
func StrLen(s string) int {
	return utf8.RuneCountInString(s)
}

// Очистить строку от Html
func StripHTML(s string) string {
	strippedBytes := bluemonday.StrictPolicy().SanitizeBytes([]byte(s))

	return string(strippedBytes)
}

/**
 * Тримит элементы массива строк - если после трима длинна элемента=0
 * то элемент удаляется из массива
 *
 * &nbsp; как и другие xml-сущности (кажется, все) имеют аналог в utf-8.
 * так, &nbsp; можно представить в виде последовательности 0xC2 0xA0
 * (двумя байтами вместо шести), чем достаточно нередко пользуются разработчики.
 * поэтому такой символ мы тоже будем тримить.
 * Я с этим символом где то столкнулся а где не помню.
 */
func Trim(s []string, trimmers []string) []string {
	var result []string

	trimmed := ""

	for _, v := range s {
		trimmed = v

		for {
			lenBefore := len(trimmed)

			for _, trimmer := range trimmers {
				trimmed = strings.Trim(trimmed, trimmer)
			}

			lenAfter := len(trimmed)

			if lenAfter == lenBefore {
				break
			}
		}

		if utf8.RuneCountInString(trimmed) > 0 {
			result = append(result, trimmed)
		}
	}

	return result
}

// Функция возвращает Слайс из слов
func Words(text string) []string {
	splittersTrimmers := []string{HTMLSpaceInUtf8, Space, NewLine, CarriageReturn, Tab}

	for _, value := range splittersTrimmers {
		text = strings.ReplaceAll(text, value, Space)
	}

	words := strings.Split(text, Space)
	words = Trim(words, splittersTrimmers)

	return words
}

// Функция нормалицации текста - для устранения множественных повторяющихся символов
// По умолчанию повторяющимся символом является пробел.
//
// Функция делит строку на массив по указанным символам - пробелам
// затем Тримит - т.е. удаляет пустые элементы
// затем обратно склеивает с указанным символом
func NormalizeText(text string, glue string) string {
	words := Words(text)
	result := strings.Join(words, glue)

	return result
}

// Содержит ли строка HTML?
func IsHTML(s string) bool {
	fullLength := StrLen(s)
	strippedLength := StrLen(StripHTML(s))

	return fullLength != strippedLength
}

func RemoveDuplicatesString(source []string) []string {
	result := make([]string, 0)

	encountered := make(map[string]bool)

	for v := range source {
		_, found := encountered[source[v]]
		if !found {
			encountered[source[v]] = true

			result = append(result, source[v])
		}
	}

	return result
}

// Нормализуем символы URL
// надо сказать что здесь нормализуется не полный URL а его финальная часть name
// http://domain.name/some/url/<name>/
// соответственно аргумент функции принимает только финальную часть!
func NormalizeNameURL(name string) string {
	cleanName := strings.Trim(name, Space)
	cleanName = strings.ToLower(cleanName)
	cleanName = NormalizeText(cleanName, Space)

	mathReplaces := map[string]string{
		Mult:    EmptyString,
		Div:     EmptyString,
		Plus:    "_plus_",
		Minus:   "_minus_",
		Equal:   EmptyString,
		Number:  "_num_",
		Percent: EmptyString,
		Gradus:  "_gradus_",
	}

	for math, replace := range mathReplaces {
		cleanName = strings.ReplaceAll(cleanName, math, replace)
	}

	quotesReplaces := map[string]string{
		SingleQuote:       Space,
		DoubleQuote:       Space,
		QuoteLeft:         Space,
		QuoteRight:        Space,
		RoundBracketLeft:  Space,
		RoundBracketRight: Space,
		Comma:             Space,
		Colon:             Space,
		Semicolon:         Space,
		Dot:               Space,
		Space:             Space,
		HTMLMDash:         Space,
		Underline:         Space,
		Exclamation:       Space,
	}

	for quote, replace := range quotesReplaces {
		cleanName = strings.ReplaceAll(cleanName, quote, replace)
	}

	cleanName = NormalizeText(cleanName, Space)
	words := Words(cleanName)
	uniqueWords := RemoveDuplicatesString(words)

	return strings.Join(uniqueWords, Underline)
}

// Нормализуем фразы алфавитного указателя
func NormalizeAlphabet(name string) string {
	cleanName := strings.Trim(name, Space)
	phrases := strings.Split(cleanName, Comma)

	for i := range phrases {
		cleanPhrase := strings.Trim(phrases[i], Space)
		phrases[i] = NormalizeText(cleanPhrase, Space)
	}

	return strings.Join(phrases, Comma+Space)
}

func TranslitRusLat(in string) string {
	out := in

	translitReplaces := map[string]string{
		"А": "A",
		"Б": "B",
		"В": "V",
		"Г": "G",
		"Д": "D",
		"Е": "E",
		"Ё": "YO",
		"Ж": "ZH",
		"З": "Z",
		"И": "I",
		"Й": "Y",
		"К": "K",
		"Л": "L",
		"М": "M",
		"Н": "N",
		"О": "O",
		"П": "P",
		"Р": "R",
		"С": "S",
		"Т": "T",
		"У": "U",
		"Ф": "F",
		"Х": "H",
		"Ц": "C",
		"Ч": "CH",
		"Ш": "SH",
		"Щ": "SCH",
		"Ъ": "",
		"Ы": "Y",
		"Ь": "",
		"Э": "E",
		"Ю": "YU",
		"Я": "YA",
		"а": "a",
		"б": "b",
		"в": "v",
		"г": "g",
		"д": "d",
		"е": "e",
		"ё": "yo",
		"ж": "zh",
		"з": "z",
		"и": "i",
		"й": "y",
		"к": "k",
		"л": "l",
		"м": "m",
		"н": "n",
		"о": "o",
		"п": "p",
		"р": "r",
		"с": "s",
		"т": "t",
		"у": "u",
		"ф": "f",
		"х": "h",
		"ц": "c",
		"ч": "ch",
		"ш": "sh",
		"щ": "sch",
		"ъ": "",
		"ы": "y",
		"ь": "",
		"э": "e",
		"ю": "yu",
		"я": "ya",
	}

	for rus, lat := range translitReplaces {
		out = strings.ReplaceAll(out, rus, lat)
	}

	return out
}
