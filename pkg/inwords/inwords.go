package inwords

import (
	"fmt"
	"strings"

	"github.com/nobuenhombre/suikat/pkg/chunks"
	"github.com/nobuenhombre/suikat/pkg/converter"
	"github.com/nobuenhombre/suikat/pkg/ge"
	"github.com/nobuenhombre/suikat/pkg/webula"
)

const (
	GenderMale   = "male"
	GenderFemale = "female"
)

const (
	ClassLength      = 3
	MaxNumberClasses = 5
)

type gender string

type numberClass int

type genderLessWords []string

type genderWords map[gender]genderLessWords

type classGenderWords struct {
	gender gender
	words  genderLessWords
}

type unitsWords map[numberClass]classGenderWords

type numberWords struct {
	null           string
	firstTen       genderWords
	secondTen      genderLessWords
	tens           genderLessWords
	hundreds       genderLessWords
	maxNumberClass numberClass
	units          unitsWords
}

func newNumberWords() *numberWords {
	return &numberWords{
		null: "ноль",
		firstTen: genderWords{
			GenderMale:   {"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"},
			GenderFemale: {"", "одна", "две", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"},
		},
		secondTen: genderLessWords{
			"десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать",
			"пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать",
		},
		tens: genderLessWords{
			"", "", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто",
		},
		hundreds: genderLessWords{
			"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот",
		},
		maxNumberClass: MaxNumberClasses,
		units: unitsWords{
			0: {gender: GenderFemale, words: genderLessWords{"копейка", "копейки", "копеек"}},
			1: {gender: GenderMale, words: genderLessWords{"рубль", "рубля", "рублей"}},
			2: {gender: GenderFemale, words: genderLessWords{"тысяча", "тысячи", "тысяч"}},
			3: {gender: GenderMale, words: genderLessWords{"миллион", "миллиона", "миллионов"}},
			4: {gender: GenderMale, words: genderLessWords{"миллиард", "миллиарда", "миллиардов"}},
			5: {gender: GenderMale, words: genderLessWords{"триллион", "триллиона", "триллионов"}},
		},
	}
}

func getClassDigits(classDigits string) (i3, i2, i1 int64, err error) {
	if len(classDigits) != ClassLength {
		return 0, 0, 0, ge.Pin(&ge.MismatchError{
			ComparedItems: "len(classDigits)",
			Expected:      ClassLength,
			Actual:        len(classDigits),
		})
	}

	digits := chunks.SplitStr(classDigits, 1)

	i3, err = converter.StringToInt64(digits[0])
	if err != nil {
		return 0, 0, 0, ge.Pin(err)
	}

	i2, err = converter.StringToInt64(digits[1])
	if err != nil {
		return 0, 0, 0, ge.Pin(err)
	}

	i1, err = converter.StringToInt64(digits[2])
	if err != nil {
		return 0, 0, 0, ge.Pin(err)
	}

	return i3, i2, i1, nil
}

const (
	NumberHundred = 100
	NumberTwenty  = 20
	NumberTen     = 10
	NumberFive    = 5
)

// get - склонение словоформы
func (glw genderLessWords) get(n int64) string {
	nx := n % NumberHundred
	if nx > NumberTen && nx < NumberTwenty {
		return glw[2]
	}

	nxy := nx % NumberTen
	if nxy > 1 && nxy < NumberFive {
		return glw[1]
	}

	if nxy == 1 {
		return glw[0]
	}

	return glw[2]
}

func formatSingleClass(numberWords *numberWords, chunkKey int, chunkValue string) (out []string, err error) {
	out = make([]string, 0)

	chunkValueInt, err := converter.StringToInt64(chunkValue)
	if err != nil {
		return out, ge.Pin(err)
	}

	if chunkValueInt == 0 {
		return out, nil
	}

	class := numberWords.maxNumberClass - numberClass(chunkKey) - 1

	gender := numberWords.units[class].gender

	i3, i2, i1, err := getClassDigits(chunkValue)
	if err != nil {
		return out, ge.Pin(err)
	}

	out = append(out, numberWords.hundreds[i3]) // 1xx-9xx
	if i2 > 1 {
		out = append(out, numberWords.tens[i2]+" "+numberWords.firstTen[gender][i1]) // 20-99
	} else if i2 == 1 {
		out = append(out, numberWords.secondTen[i1]) // 10-19
	} else {
		out = append(out, numberWords.firstTen[gender][i1]) // 1-9
	}

	// units without rub & kop
	if class > 1 {
		out = append(out, numberWords.units[class].words.get(chunkValueInt))
	}

	return out, nil
}

func Format(num float64) (string, error) {
	numberWords := newNumberWords()

	list := strings.Split(fmt.Sprintf("%015.2f", num), ".")

	rubStr := list[0]
	kopStr := list[1]

	rubInt, err := converter.StringToInt64(rubStr)
	if err != nil {
		return "", ge.Pin(err)
	}

	kopInt, err := converter.StringToInt64(kopStr)
	if err != nil {
		return "", ge.Pin(err)
	}

	out := make([]string, 0)

	if rubInt > 0 {
		strChunks := chunks.SplitStr(rubStr, ClassLength)

		for chunkKey, chunkValue := range strChunks {
			classOut, err := formatSingleClass(numberWords, chunkKey, chunkValue)
			if err != nil {
				return "", ge.Pin(err)
			}

			out = append(out, classOut...)
		}
	} else {
		out = append(out, numberWords.null)
	}

	morphedRub := numberWords.units[1].words.get(rubInt)
	out = append(out, morphedRub)

	morphedKop := kopStr + " " + numberWords.units[0].words.get(kopInt)
	out = append(out, morphedKop)

	gluedOut := strings.Join(out, " ")

	return webula.NormalizeText(gluedOut, " "), nil
}
