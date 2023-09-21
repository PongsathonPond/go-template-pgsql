package pgsql

import (
	"fmt"
	"github.com/uniplaces/carbon"
	conf "idev-cms-service/config"
	"log"
	"reflect"
	"strings"
	"time"
)

func (repo *Repository) makeFilters(filters []string) (query string, args []interface{}) {
	var key, operations, value, start, end string

	args = make([]interface{}, 0)
	for _, v := range filters {
		slFilter := strings.SplitN(v, ":", 3)
		if len(slFilter) == 3 {
			key = slFilter[0]
			operations = slFilter[1]
			value = slFilter[2]

			// find if value is date range (between)
			slFilterDate := strings.SplitN(value, "|", 2)
			if len(slFilterDate) == 2 {
				start, end = makeValueDateTimeFilter(slFilter[2])
				//start = slFilterDate[0]
				//end = slFilterDate[1]
			}
		} else {
			key = slFilter[0]
			operations = slFilter[1]
		}

		if query != "" {
			query += " AND "
		}

		switch operations {
		case "eq":
			query += key + " = ?"
			args = append(args, value)
			break
		case "ne":
			query += key + " <> ?"
			args = append(args, value)
			break
		case "gt":
			query += key + " > ?"
			args = append(args, value)
			break
		case "gte":
			query += key + " >= ?"
			args = append(args, value)
			break
		case "lt":
			query += key + " < ?"
			args = append(args, value)
			break
		case "lte":
			query += key + " <= ?"
			args = append(args, value)
			break
		case "like":
			query += key + " LIKE ?"
			args = append(args, "%"+value+"%")
			break
		case "isNull":
			query += key + " IS NULL"
			//args = append(args, "")
			break
		case "isNotNull":
			query += key + " <> ?"
			args = append(args, "")
			break
		case "btw":
			query += key + " BETWEEN ? AND ?"
			args = append(args, start, end)
			break
		case "in":
			query += key + " IN (?)"
			in := makeIn(value)
			args = append(args, in)
			break
		case "eqDate":
			start, end = makeValueDateTimeFilter(slFilter[2])
			query += key + " BETWEEN ? AND ?"
			args = append(args, start, end)
		default:
			query += key + " = ?"
			args = append(args, value)
			break
		}
	}
	return query, args
}

func (repo *Repository) makeSorts(sorts []string) (strSort string) {
	strSort = ""
	for i, v := range sorts {
		slFilter := strings.Split(v, ":")
		if i > 0 {
			strSort += ", "
		}
		strSort += fmt.Sprintf("%s %s", slFilter[0], slFilter[1])
	}
	return strSort
}

func (repo *Repository) interfaceToSlice(slice interface{}) (ret []interface{}) {
	s := reflect.Indirect(reflect.ValueOf(slice))

	log.Println(s)
	if s.Kind() != reflect.Slice {
		log.Panic("InterfaceSlice() given a non-slice type")
	}
	if s.IsNil() {
		return nil
	}

	ret = make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func makeIn(val string) (result []string) {
	result = strings.Split(val, "|")
	return
}

func makeValueDateTimeFilter(val string) (startDate, endDate string) {
	slFilterDate := strings.Split(val, "|")
	if len(slFilterDate) == 2 {
		sd, err := checkDateTimeFormatString(slFilterDate[0])
		if err != nil {
			return slFilterDate[0], slFilterDate[1]
		}

		ed, err := checkDateTimeFormatString(slFilterDate[1])
		if err != nil {
			return slFilterDate[0], slFilterDate[1]
		}

		sdCarbon := carbon.NewCarbon(sd)
		_ = sdCarbon.SetTimeZone(conf.Config{}.AppTimeZone) // ใช้โซนเวลาจาก conf.Config

		edCarbon := carbon.NewCarbon(ed)
		_ = edCarbon.SetTimeZone(conf.Config{}.AppTimeZone)

		startDate = sdCarbon.UTC().Format("2006-01-02 15:04:05")
		endDate = edCarbon.UTC().Format("2006-01-02 15:04:05")
	} else {
		d, err := checkDateTimeFormatString(slFilterDate[0])
		if err != nil {
			return startDate, endDate
		}

		t := carbon.NewCarbon(d)
		_ = t.SetTimeZone(conf.Config{}.AppTimeZone)

		startDate = t.StartOfDay().UTC().Format("2006-01-02 15:04:05")
		endDate = t.EndOfDay().UTC().Format("2006-01-02 15:04:05")
	}

	return startDate, endDate
}

func checkDateTimeFormatString(val string) (d time.Time, err error) {
	layouts := []string{"2006-01-02 15:04:05", "2006-01-02"}
	for _, l := range layouts {
		d, err = time.Parse(l, val)
		if !d.IsZero() {
			return
		}
	}

	return
}
