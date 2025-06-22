package dates

import (
	"fmt"
	"pdate/internal/job"
	"strings"
	"time"
)

var weekdayNames = map[job.Language][]string{
	job.English:    {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
	job.Spanish:    {"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"},
	job.French:     {"Dimanche", "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi"},
	job.Swiss:      {"suntig", "mäntig", "zistig", "mittwuch", "donstig", "fritig", "samstig"},
	job.German:     {"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"},
	job.Italian:    {"Domenica", "Lunedì", "Martedì", "Mercoledì", "Giovedì", "Venerdì", "Sabato"},
	job.Portuguese: {"Domingo", "Segunda-feira", "Terça-feira", "Quarta-feira", "Quinta-feira", "Sexta-feira", "Sábado"},
	job.Dutch:      {"Zondag", "Maandag", "Dinsdag", "Woensdag", "Donderdag", "Vrijdag", "Zaterdag"},
	job.Russian:    {"Воскресенье", "Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота"},
	job.Chinese:    {"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
	job.Arabic:     {"الأحد", "الاثنين", "الثلاثاء", "الأربعاء", "الخميس", "الجمعة", "السبت"},
	job.Hindi:      {"रविवार", "सोमवार", "मंगलवार", "बुधवार", "गुरुवार", "शुक्रवार", "शनिवार"},
}

var monthNames = map[job.Language][]string{
	job.English:    {"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	job.Spanish:    {"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"},
	job.French:     {"Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"},
	job.Swiss:      {"januar", "februar", "märz", "apriu", "mai", "juni", "july", "august", "september", "oktober", "november", "dezember"},
	job.German:     {"Januar", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"},
	job.Italian:    {"Gennaio", "Febbraio", "Marzo", "Aprile", "Maggio", "Giugno", "Luglio", "Agosto", "Settembre", "Ottobre", "Novembre", "Dicembre"},
	job.Portuguese: {"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"},
	job.Dutch:      {"Januari", "Februari", "Maart", "April", "Mei", "Juni", "Juli", "Augustus", "September", "Oktober", "November", "December"},
	job.Russian:    {"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"},
	job.Chinese:    {"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
	job.Arabic:     {"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
	job.Hindi:      {"जनवरी", "फ़रवरी", "मार्च", "अप्रैल", "मई", "जून", "जुलाई", "अगस्त", "सितंबर", "अक्टूबर", "नवंबर", "दिसंबर"},
}

var hasShortForm = map[job.Language]bool{
	job.English:    true,
	job.Spanish:    true,
	job.French:     true,
	job.Swiss:      true,
	job.German:     true,
	job.Italian:    true,
	job.Portuguese: true,
	job.Dutch:      true,
}

func FormatDates(dates []time.Time, format string, lang job.Language) []string {
	var formattedDates []string
	for _, date := range dates {
		formattedDates = append(formattedDates, ReplaceDatePlaceholdersWithDate(format, date, lang))
	}
	return formattedDates
}

func ReplaceDatePlaceholdersWithDate(input string, date time.Time, lang job.Language) string {
	wdFull := weekdayNames[lang][int(date.Weekday())]
	wdShort := GetShortFormName(wdFull, lang)
	mnFull := monthNames[lang][int(date.Month())-1]
	mnShort := GetShortFormName(mnFull, lang)
	replacer := strings.NewReplacer(
		"{YYYY}", fmt.Sprintf("%04d", date.Year()),
		"{YY}", fmt.Sprintf("%02d", date.Year()%100),
		"{MM}", fmt.Sprintf("%02d", int(date.Month())),
		"{DD}", fmt.Sprintf("%02d", date.Day()),
		"{WD}", wdFull,
		"{wd}", wdShort,
		"{MN}", mnFull,
		"{mn}", mnShort,
		"{M}", fmt.Sprintf("%d", int(date.Month())),
		"{D}", fmt.Sprintf("%d", date.Day()),
	)
	return replacer.Replace(input)
}

func GetShortFormName(input string, lang job.Language) string {
	if hasShortForm[lang] {
		return input[:3]
	}
	return input
}
