package webula

import (
	"strings"
	"unicode/utf8"

	"github.com/microcosm-cc/bluemonday"
)

const (
	EmptyString     = ""
	Space           = " "
	Underline       = "_"
	Dot             = "."
	NewLine         = "\n"
	CarriageReturn  = "\r"
	Tab             = "\t"
	Comma           = ","
	Semicolon       = ";"
	NbspSpaceInUtf8 = "\xc2\xa0"
	HtmlSpace       = "&nbsp;"
)

// Сколько Символов в строке
func StrLen(s string) int {
	return utf8.RuneCountInString(s)
}

// Очистить строку от Html
func StripHtml(s string) string {
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
		for _, trimmer := range trimmers {
			trimmed = strings.Trim(trimmed, trimmer)
		}
		if utf8.RuneCountInString(trimmed) > 0 {
			result = append(result, trimmed)
		}
	}
	return result
}

// Функция возвращает Слайс из слов
func Words(text string) []string {
	words := strings.Split(text, Space)
	words = Trim(words, []string{NbspSpaceInUtf8, Space, NewLine, CarriageReturn, Tab})
	return words
}

// Функция нормалицации текста - для устранения множественных повторяющихся символов
// По умолчанию повторяющимся символом является пробел.
//
// Функция делит строку на массив по указанным символам - пробелам
// затем Тримит - т.е. удаляет пустые элементы
// затем обратно склеивает с указанным символом
func NormalizeText(text string) string {
	words := Words(text)
	result := strings.Join(words, Space)
	return result
}

// Содержит ли строка HTML?
func IsHtml(s string) bool {
	fullLength := StrLen(s)
	strippedLength := StrLen(StripHtml(s))
	return fullLength != strippedLength
}
