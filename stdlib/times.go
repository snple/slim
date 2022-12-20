package stdlib

import (
	"time"

	"github.com/snple/slim"
)

var timesModule = map[string]slim.Object{
	"format_ansic":        &slim.String{Value: time.ANSIC},
	"format_unix_date":    &slim.String{Value: time.UnixDate},
	"format_ruby_date":    &slim.String{Value: time.RubyDate},
	"format_rfc822":       &slim.String{Value: time.RFC822},
	"format_rfc822z":      &slim.String{Value: time.RFC822Z},
	"format_rfc850":       &slim.String{Value: time.RFC850},
	"format_rfc1123":      &slim.String{Value: time.RFC1123},
	"format_rfc1123z":     &slim.String{Value: time.RFC1123Z},
	"format_rfc3339":      &slim.String{Value: time.RFC3339},
	"format_rfc3339_nano": &slim.String{Value: time.RFC3339Nano},
	"format_kitchen":      &slim.String{Value: time.Kitchen},
	"format_stamp":        &slim.String{Value: time.Stamp},
	"format_stamp_milli":  &slim.String{Value: time.StampMilli},
	"format_stamp_micro":  &slim.String{Value: time.StampMicro},
	"format_stamp_nano":   &slim.String{Value: time.StampNano},
	"nanosecond":          &slim.Int{Value: int64(time.Nanosecond)},
	"microsecond":         &slim.Int{Value: int64(time.Microsecond)},
	"millisecond":         &slim.Int{Value: int64(time.Millisecond)},
	"second":              &slim.Int{Value: int64(time.Second)},
	"minute":              &slim.Int{Value: int64(time.Minute)},
	"hour":                &slim.Int{Value: int64(time.Hour)},
	"january":             &slim.Int{Value: int64(time.January)},
	"february":            &slim.Int{Value: int64(time.February)},
	"march":               &slim.Int{Value: int64(time.March)},
	"april":               &slim.Int{Value: int64(time.April)},
	"may":                 &slim.Int{Value: int64(time.May)},
	"june":                &slim.Int{Value: int64(time.June)},
	"july":                &slim.Int{Value: int64(time.July)},
	"august":              &slim.Int{Value: int64(time.August)},
	"september":           &slim.Int{Value: int64(time.September)},
	"october":             &slim.Int{Value: int64(time.October)},
	"november":            &slim.Int{Value: int64(time.November)},
	"december":            &slim.Int{Value: int64(time.December)},
	"sleep": &slim.UserFunction{
		Name:  "sleep",
		Value: timesSleep,
	}, // sleep(int)
	"parse_duration": &slim.UserFunction{
		Name:  "parse_duration",
		Value: timesParseDuration,
	}, // parse_duration(str) => int
	"since": &slim.UserFunction{
		Name:  "since",
		Value: timesSince,
	}, // since(time) => int
	"until": &slim.UserFunction{
		Name:  "until",
		Value: timesUntil,
	}, // until(time) => int
	"duration_hours": &slim.UserFunction{
		Name:  "duration_hours",
		Value: timesDurationHours,
	}, // duration_hours(int) => float
	"duration_minutes": &slim.UserFunction{
		Name:  "duration_minutes",
		Value: timesDurationMinutes,
	}, // duration_minutes(int) => float
	"duration_nanoseconds": &slim.UserFunction{
		Name:  "duration_nanoseconds",
		Value: timesDurationNanoseconds,
	}, // duration_nanoseconds(int) => int
	"duration_seconds": &slim.UserFunction{
		Name:  "duration_seconds",
		Value: timesDurationSeconds,
	}, // duration_seconds(int) => float
	"duration_string": &slim.UserFunction{
		Name:  "duration_string",
		Value: timesDurationString,
	}, // duration_string(int) => string
	"month_string": &slim.UserFunction{
		Name:  "month_string",
		Value: timesMonthString,
	}, // month_string(int) => string
	"date": &slim.UserFunction{
		Name:  "date",
		Value: timesDate,
	}, // date(year, month, day, hour, min, sec, nsec) => time
	"now": &slim.UserFunction{
		Name:  "now",
		Value: timesNow,
	}, // now() => time
	"parse": &slim.UserFunction{
		Name:  "parse",
		Value: timesParse,
	}, // parse(format, str) => time
	"unix": &slim.UserFunction{
		Name:  "unix",
		Value: timesUnix,
	}, // unix(sec, nsec) => time
	"add": &slim.UserFunction{
		Name:  "add",
		Value: timesAdd,
	}, // add(time, int) => time
	"add_date": &slim.UserFunction{
		Name:  "add_date",
		Value: timesAddDate,
	}, // add_date(time, years, months, days) => time
	"sub": &slim.UserFunction{
		Name:  "sub",
		Value: timesSub,
	}, // sub(t time, u time) => int
	"after": &slim.UserFunction{
		Name:  "after",
		Value: timesAfter,
	}, // after(t time, u time) => bool
	"before": &slim.UserFunction{
		Name:  "before",
		Value: timesBefore,
	}, // before(t time, u time) => bool
	"time_year": &slim.UserFunction{
		Name:  "time_year",
		Value: timesTimeYear,
	}, // time_year(time) => int
	"time_month": &slim.UserFunction{
		Name:  "time_month",
		Value: timesTimeMonth,
	}, // time_month(time) => int
	"time_day": &slim.UserFunction{
		Name:  "time_day",
		Value: timesTimeDay,
	}, // time_day(time) => int
	"time_weekday": &slim.UserFunction{
		Name:  "time_weekday",
		Value: timesTimeWeekday,
	}, // time_weekday(time) => int
	"time_hour": &slim.UserFunction{
		Name:  "time_hour",
		Value: timesTimeHour,
	}, // time_hour(time) => int
	"time_minute": &slim.UserFunction{
		Name:  "time_minute",
		Value: timesTimeMinute,
	}, // time_minute(time) => int
	"time_second": &slim.UserFunction{
		Name:  "time_second",
		Value: timesTimeSecond,
	}, // time_second(time) => int
	"time_nanosecond": &slim.UserFunction{
		Name:  "time_nanosecond",
		Value: timesTimeNanosecond,
	}, // time_nanosecond(time) => int
	"time_unix": &slim.UserFunction{
		Name:  "time_unix",
		Value: timesTimeUnix,
	}, // time_unix(time) => int
	"time_unix_nano": &slim.UserFunction{
		Name:  "time_unix_nano",
		Value: timesTimeUnixNano,
	}, // time_unix_nano(time) => int
	"time_format": &slim.UserFunction{
		Name:  "time_format",
		Value: timesTimeFormat,
	}, // time_format(time, format) => string
	"time_location": &slim.UserFunction{
		Name:  "time_location",
		Value: timesTimeLocation,
	}, // time_location(time) => string
	"time_string": &slim.UserFunction{
		Name:  "time_string",
		Value: timesTimeString,
	}, // time_string(time) => string
	"is_zero": &slim.UserFunction{
		Name:  "is_zero",
		Value: timesIsZero,
	}, // is_zero(time) => bool
	"to_local": &slim.UserFunction{
		Name:  "to_local",
		Value: timesToLocal,
	}, // to_local(time) => time
	"to_utc": &slim.UserFunction{
		Name:  "to_utc",
		Value: timesToUTC,
	}, // to_utc(time) => time
}

func timesSleep(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	i1, ok := slim.ToInt64(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	time.Sleep(time.Duration(i1))
	ret = slim.UndefinedValue

	return
}

func timesParseDuration(args ...slim.Object) (
	ret slim.Object,
	err error,
) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	s1, ok := slim.ToString(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	dur, err := time.ParseDuration(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &slim.Int{Value: int64(dur)}

	return
}

func timesSince(args ...slim.Object) (
	ret slim.Object,
	err error,
) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: int64(time.Since(t1))}

	return
}

func timesUntil(args ...slim.Object) (
	ret slim.Object,
	err error,
) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: int64(time.Until(t1))}

	return
}

func timesDurationHours(args ...slim.Object) (
	ret slim.Object,
	err error,
) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	i1, ok := slim.ToInt64(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Float{Value: time.Duration(i1).Hours()}

	return
}

func timesDurationMinutes(args ...slim.Object) (
	ret slim.Object,
	err error,
) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	i1, ok := slim.ToInt64(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Float{Value: time.Duration(i1).Minutes()}

	return
}

func timesDurationNanoseconds(args ...slim.Object) (
	ret slim.Object,
	err error,
) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	i1, ok := slim.ToInt64(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: time.Duration(i1).Nanoseconds()}

	return
}

func timesDurationSeconds(args ...slim.Object) (
	ret slim.Object,
	err error,
) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	i1, ok := slim.ToInt64(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Float{Value: time.Duration(i1).Seconds()}

	return
}

func timesDurationString(args ...slim.Object) (
	ret slim.Object,
	err error,
) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	i1, ok := slim.ToInt64(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.String{Value: time.Duration(i1).String()}

	return
}

func timesMonthString(args ...slim.Object) (
	ret slim.Object,
	err error,
) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	i1, ok := slim.ToInt64(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.String{Value: time.Month(i1).String()}

	return
}

func timesDate(args ...slim.Object) (
	ret slim.Object,
	err error,
) {
	if len(args) != 7 {
		err = slim.ErrWrongNumArguments
		return
	}

	i1, ok := slim.ToInt(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}
	i2, ok := slim.ToInt(args[1])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}
	i3, ok := slim.ToInt(args[2])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}
	i4, ok := slim.ToInt(args[3])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}
	i5, ok := slim.ToInt(args[4])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "fifth",
			Expected: "int(compatible)",
			Found:    args[4].TypeName(),
		}
		return
	}
	i6, ok := slim.ToInt(args[5])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "sixth",
			Expected: "int(compatible)",
			Found:    args[5].TypeName(),
		}
		return
	}
	i7, ok := slim.ToInt(args[6])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "seventh",
			Expected: "int(compatible)",
			Found:    args[6].TypeName(),
		}
		return
	}

	ret = &slim.Time{
		Value: time.Date(i1,
			time.Month(i2), i3, i4, i5, i6, i7, time.Now().Location()),
	}

	return
}

func timesNow(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 0 {
		err = slim.ErrWrongNumArguments
		return
	}

	ret = &slim.Time{Value: time.Now()}

	return
}

func timesParse(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 2 {
		err = slim.ErrWrongNumArguments
		return
	}

	s1, ok := slim.ToString(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := slim.ToString(args[1])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	parsed, err := time.Parse(s1, s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &slim.Time{Value: parsed}

	return
}

func timesUnix(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 2 {
		err = slim.ErrWrongNumArguments
		return
	}

	i1, ok := slim.ToInt64(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := slim.ToInt64(args[1])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &slim.Time{Value: time.Unix(i1, i2)}

	return
}

func timesAdd(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 2 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := slim.ToInt64(args[1])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &slim.Time{Value: t1.Add(time.Duration(i2))}

	return
}

func timesSub(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 2 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := slim.ToTime(args[1])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: int64(t1.Sub(t2))}

	return
}

func timesAddDate(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 4 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := slim.ToInt(args[1])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	i3, ok := slim.ToInt(args[2])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := slim.ToInt(args[3])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	ret = &slim.Time{Value: t1.AddDate(i2, i3, i4)}

	return
}

func timesAfter(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 2 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := slim.ToTime(args[1])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if t1.After(t2) {
		ret = slim.TrueValue
	} else {
		ret = slim.FalseValue
	}

	return
}

func timesBefore(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 2 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := slim.ToTime(args[1])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.Before(t2) {
		ret = slim.TrueValue
	} else {
		ret = slim.FalseValue
	}

	return
}

func timesTimeYear(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: int64(t1.Year())}

	return
}

func timesTimeMonth(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: int64(t1.Month())}

	return
}

func timesTimeDay(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: int64(t1.Day())}

	return
}

func timesTimeWeekday(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: int64(t1.Weekday())}

	return
}

func timesTimeHour(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: int64(t1.Hour())}

	return
}

func timesTimeMinute(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: int64(t1.Minute())}

	return
}

func timesTimeSecond(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: int64(t1.Second())}

	return
}

func timesTimeNanosecond(args ...slim.Object) (
	ret slim.Object,
	err error,
) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: int64(t1.Nanosecond())}

	return
}

func timesTimeUnix(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: t1.Unix()}

	return
}

func timesTimeUnixNano(args ...slim.Object) (
	ret slim.Object,
	err error,
) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Int{Value: t1.UnixNano()}

	return
}

func timesTimeFormat(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 2 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := slim.ToString(args[1])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s := t1.Format(s2)
	if len(s) > slim.MaxStringLen {

		return nil, slim.ErrStringLimit
	}

	ret = &slim.String{Value: s}

	return
}

func timesIsZero(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.IsZero() {
		ret = slim.TrueValue
	} else {
		ret = slim.FalseValue
	}

	return
}

func timesToLocal(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Time{Value: t1.Local()}

	return
}

func timesToUTC(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.Time{Value: t1.UTC()}

	return
}

func timesTimeLocation(args ...slim.Object) (
	ret slim.Object,
	err error,
) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.String{Value: t1.Location().String()}

	return
}

func timesTimeString(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		err = slim.ErrWrongNumArguments
		return
	}

	t1, ok := slim.ToTime(args[0])
	if !ok {
		err = slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &slim.String{Value: t1.String()}

	return
}
