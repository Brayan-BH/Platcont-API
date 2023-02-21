package date

import (
	"fmt"
	"strings"
	"time"
)

/**
 * Retorna el índice de un elemento en un arreglo de strings
 * date[string] = fecha a validar formato string dd/mm/yyyy
 * return [bool],[error] = [true or false],[descripción del error or  nil]
 */
func CheckDate(date string) error {
	_, err := time.Parse("02/01/2006", date)
	if err != nil {
		return err
	}
	return nil
}

/**
 * Retorna Fecha en formato yyyy-mm-dd hh:mm:ss en zona horaria  (America/Bogota)
 * Return [string] : fecha  formato string yyyy-mm-dd hh:mm:ss (America/Bogota)
 */
func GetDateLocationString() string {
	loc, _ := time.LoadLocation("America/Bogota")
	t := time.Now().In(loc).Format("2006-01-02 15:04:05")
	return t
}

/**
 * Retorna fecha en formato dd/mm/yyyy en zona horaria  (America/Bogota)
 * Return [string] : fecha  formato string dd/mm/yyyy (America/Bogota)
 */
func GetFechaLocationString() string {
	loc, _ := time.LoadLocation("America/Bogota")
	t := time.Now().In(loc).Format("02/01/2006")

	return t
}

// retorna fecha en formato string segun la location del usuario = (America/Bogota) 01:01:59
func GetHoraLocationString() string {
	loc, _ := time.LoadLocation("America/Bogota")
	t := time.Now().In(loc)
	now := fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())

	return now
}

// agrega dias a la fecha actual y retorna fecha en formato string segun la location del usuario = (America/Bogota) 01/01/2021
func GetFechaLocationStringAdd(days int) string {
	loc, _ := time.LoadLocation("America/Bogota")
	t := time.Now().In(loc)
	t = t.AddDate(0, 0, days)
	return t.Format("02/01/2006")

	// return now
}

// retorna la fecha segun la locacion del usurio es este caso esta configurado para (America/Bogota)
func GetDateLocationTime() time.Time {
	loc, _ := time.LoadLocation("America/Bogota")
	//set timezone,
	t := time.Now().In(loc)

	return t
}

// retorna la fecha segun la locacion del usurio es este caso esta configurado para (America/Bogota) 2021-01-01 12:00:00.000
func GetYearString() string {
	loc, _ := time.LoadLocation("America/Bogota")
	t := time.Now().In(loc).Format("2006")
	return t
}

// retorna la fecha segun la locacion del usurio es este caso esta configurado para (America/Bogota) 2021-01-01 12:00:00.000
func GetYear() int64 {
	loc, _ := time.LoadLocation("America/Bogota")
	t := time.Now().In(loc)
	return int64(t.Year())
}

func GetMonth() int64 {
	loc, _ := time.LoadLocation("America/Bogota")
	t := time.Now().In(loc)
	return int64(t.Month())
}

func GetMonthString() string {
	loc, _ := time.LoadLocation("America/Bogota")
	t := time.Now().In(loc).Format("01")
	return t
}

//suma días a una fecha
func SumarDate(date string, days int) string {
	date_time, _ := time.Parse("02/01/2006", date)
	date_time = date_time.AddDate(0, 0, days)
	return date_time.Format("02/01/2006")
}

/**
 * Retorna la diferencia entre dos fechas en días.
 * @param date {[]string}: fecha inicial y fecha final en formato dd/mm/yyyy, si no se pasa la fecha inicial se toma la fecha actual
 * @return {float64} = diferencia entre fechas en días
 */
func DiferenciaDate(date ...string) float64 {
	var date_time time.Time
	var date_time_restar time.Time
	if len(date) > 0 {
		date_time, _ = time.Parse("02/01/2006", date[0])
	} else {
		date_time = time.Now()
	}
	if len(date) == 2 {
		date_time_restar, _ = time.Parse("02/01/2006", date[1])
	} else {
		date_time_restar = time.Now()
	}
	diferencia := date_time.Sub(date_time_restar)

	return diferencia.Hours() / 24
}

func getYearDate(date string) (string, int) {
	date_time, _ := time.Parse("02/01/2006", date)
	now := fmt.Sprintf("%d", date_time.Year())
	yearInt := int(date_time.Year())
	return now, yearInt
}

func getMonthDate(date string) (string, int) {
	date_time, _ := time.Parse("02/01/2006", date)
	now := fmt.Sprintf("%02d", date_time.Month())
	monthInt := int(date_time.Month())
	return now, monthInt
}

func getDayDate(date string) (string, int) {
	date_time, _ := time.Parse("02/01/2006", date)
	now := fmt.Sprintf("%02d", date_time.Day())
	dayInt := int(date_time.Day())

	return now, dayInt
}

func IsItHoliday(date string, days_holidays []string) bool {
	date_time, _ := time.Parse("02/01/2006", date)
	dia := fmt.Sprintf("%s", date_time.Weekday())
	if dia == "Sunday" {
		return true
	}
	split_temp := strings.Split(date, "/")
	month_day_temp := split_temp[1] + "-" + split_temp[0]

	if IndexOf_String(days_holidays, month_day_temp) != -1 {
		return true
	}
	return false
}

func IndexOf_String(arreglo []string, search string) int {
	for indice, valor := range arreglo {
		if valor == search {
			return indice
		}
	}
	// -1 porque no existe
	return -1
}

//Se agrego lineas de codigo usuario Brayan

/**
 * Retorna fecha actual en zona horaria  (America/Bogota)
 * Return [time.time] : fecha  de ahora
 */
func GetDateLocation() time.Time {
	loc, _ := time.LoadLocation("America/Bogota")
	t := time.Now().In(loc)

	return t
}

/**
 * Retorna la fecha en tipo time
 * @param {string} date: fecha en foprmato DD/MM/YYYY
 * @return {time.Time} fecha en tipo time
 */
func GetDate(date string) time.Time {
	t, _ := time.Parse("02/01/2006", date)
	return t
}

// Obtenga el primer día del mes donde está la hora entrante, es decir, las 0 en punto el primer día de un mes. Si se pasa time.Now (), devuelve la hora a las 0 en punto el primer día del mes actual.
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// Obtenga el último día del mes donde está la hora entrante, es decir, las 0 en punto el último día de un mes. Si se pasa time.Now (), devuelve la hora de las 0 en punto el último día del mes actual.
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// Obtenga la hora a las 0 en punto de un día determinado
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}