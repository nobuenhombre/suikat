package datesi18nru

import "time"

func GetLongDayNames() map[time.Weekday]string {
	return map[time.Weekday]string{
		time.Sunday:    "Воскресенье",
		time.Monday:    "Понедельник",
		time.Tuesday:   "Вторник",
		time.Wednesday: "Среда",
		time.Thursday:  "Четверг",
		time.Friday:    "Пятница",
		time.Saturday:  "Суббота",
	}
}

func GetLongDayName(day time.Weekday) string {
	name, found := GetLongDayNames()[day]
	if !found {
		return "Unknown"
	}

	return name
}

func GetShortDayNames() map[time.Weekday]string {
	return map[time.Weekday]string{
		time.Sunday:    "Вс",
		time.Monday:    "Пн",
		time.Tuesday:   "Вт",
		time.Wednesday: "Ср",
		time.Thursday:  "Чт",
		time.Friday:    "Пт",
		time.Saturday:  "Сб",
	}
}

func GetShortDayName(day time.Weekday) string {
	name, found := GetShortDayNames()[day]
	if !found {
		return "Unknown"
	}

	return name
}

func GetLongMonthNames() map[time.Month]string {
	return map[time.Month]string{
		time.January:   "Январь",
		time.February:  "Февраль",
		time.March:     "Март",
		time.April:     "Апрель",
		time.May:       "Май",
		time.June:      "Июнь",
		time.July:      "Июль",
		time.August:    "Август",
		time.September: "Сентябрь",
		time.October:   "Октябрь",
		time.November:  "Ноябрь",
		time.December:  "Декабрь",
	}
}

func GetLongMonthName(month time.Month) string {
	name, found := GetLongMonthNames()[month]
	if !found {
		return "Unknown"
	}

	return name
}

func GetLongMonthNamesExt() map[time.Month]string {
	return map[time.Month]string{
		time.January:   "Января",
		time.February:  "Февраля",
		time.March:     "Марта",
		time.April:     "Апреля",
		time.May:       "Мая",
		time.June:      "Июня",
		time.July:      "Июля",
		time.August:    "Августа",
		time.September: "Сентября",
		time.October:   "Октября",
		time.November:  "Ноября",
		time.December:  "Декабря",
	}
}

func GetLongMonthNameExt(month time.Month) string {
	name, found := GetLongMonthNamesExt()[month]
	if !found {
		return "Unknown"
	}

	return name
}

func GetShortMonthNames() map[time.Month]string {
	return map[time.Month]string{
		time.January:   "Янв",
		time.February:  "Фев",
		time.March:     "Мар",
		time.April:     "Апр",
		time.May:       "Май",
		time.June:      "Июн",
		time.July:      "Июл",
		time.August:    "Авг",
		time.September: "Сен",
		time.October:   "Окт",
		time.November:  "Ноя",
		time.December:  "Дек",
	}
}

func GetShortMonthName(month time.Month) string {
	name, found := GetShortMonthNames()[month]
	if !found {
		return "Unknown"
	}

	return name
}
