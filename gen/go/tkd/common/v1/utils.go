package commonv1

import (
	"errors"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (m Month) AsMonth() time.Month {
	return time.Month(m)
}

func FromMonth(m time.Month) Month {
	return Month(m)
}

func FromTime(t time.Time) *Date {
	year, month, day := t.Date()

	return &Date{
		Year:  int64(year),
		Month: FromMonth(month),
		Day:   int32(day),
	}
}

var ErrInvalidDate = errors.New("invalid date, supported format is 2006-01-02 or 2006/01/02")

func ParseDate(date string) (*Date, error) {
	t, err := time.Parse("2006-01-02", date)
	if err == nil {
		return FromTime(t), nil
	}

	t, err = time.Parse("2006/01/02", date)
	if err == nil {
		return FromTime(t), nil
	}

	return nil, ErrInvalidDate
}

func (d *Date) AsTime() time.Time {
	return d.AsTimeInLocation(time.UTC)
}

func (d *Date) AsTimeInLocation(loc *time.Location) time.Time {
	if d.Timezone != "" {
		t, err := time.LoadLocation(d.Timezone)
		if err != nil {
			slog.Error("invalid timezone specifier in tkd.common.v1.Date", "timezone", d.Timezone)
		} else {
			loc = t
		}
	}

	return time.Date(int(d.Year), d.Month.AsMonth(), int(d.Day), 0, 0, 0, 0, loc)
}

func (dow DayOfWeek) ToWeekday() time.Weekday {
	return time.Weekday(dow)
}

func FromWeekday(d time.Weekday) DayOfWeek {
	return DayOfWeek(d)
}

func (d *TimeRange) Includes(t time.Time) bool {
	if d == nil {
		return false
	}

	if !d.From.IsValid() && !d.To.IsValid() {
		return false
	}

	if d.From.IsValid() && t.Before(d.From.AsTime()) {
		return false
	}

	if d.To.IsValid() && t.After(d.To.AsTime()) {
		return false
	}

	return true
}

func NewTimeRange(start, end time.Time) *TimeRange {
	tr := &TimeRange{}

	if !start.IsZero() {
		tr.From = timestamppb.New(start)
	}

	if !end.IsZero() {
		tr.To = timestamppb.New(end)
	}

	return tr
}

func (dt *DayTimeRange) At(t time.Time) *TimeRange {
	return NewTimeRange(
		dt.Start.At(t),
		dt.End.At(t),
	)
}

// IncludesAt checks whether the DayTimeRange, applied to ref, includes the DayTime dt.
// If dtr is nil or both Start and End are nil, true is returned.
// If dtr.Start is nil, it is assumed to be 00:00:00
// If dtr.End is nil, it is assumed to be 00:00:00 on the next day.
// Start is inclusive while end is exclusive.
func (dtr *DayTimeRange) IncludesAt(ref time.Time, dt *DayTime) bool {
	if dtr == nil {
		return true
	}

	var (
		startTime time.Time
		endTime   time.Time
		query     = dt.At(ref)
	)

	if dtr.Start != nil {
		startTime = dtr.Start.At(ref)
	}
	if dtr.End != nil {
		endTime = dtr.End.At(ref)
	}

	if startTime.IsZero() && endTime.IsZero() {
		return true
	}

	if startTime.IsZero() {
		return query.Before(endTime)
	}

	if endTime.IsZero() {
		return query.Equal(startTime) || query.After(startTime)
	}

	// if end is before start the previous day is meant
	if startTime.After(endTime) || startTime.Equal(endTime) {
		startTime = startTime.Add(-24 * time.Hour)
	}

	return (query.Equal(startTime) || query.After(startTime)) && query.Before(endTime)
}

func (dt *DayTime) At(t time.Time) time.Time {
	year, month, date := t.Date()

	return time.Date(year, month, date, int(dt.Hour), int(dt.Minute), int(dt.Second), 0, t.Location())
}

func (dt *DayTime) AsDuration() time.Duration {
	return (time.Hour * time.Duration(dt.Hour)) + (time.Minute * time.Duration(dt.Minute)) + (time.Second * time.Duration(dt.Second))
}

func DayTimeFromDuration(d time.Duration) *DayTime {
	du := float64(d)
	hours := math.Floor(du / float64(time.Hour))
	du -= (hours * float64(time.Hour))

	minutes := math.Floor(du / float64(time.Minute))
	du -= (minutes * float64(time.Minute))

	seconds := math.Floor(du / float64(time.Second))

	return &DayTime{
		Hour:   int32(hours),
		Minute: int32(minutes),
		Second: int32(seconds),
	}
}

func ParseDayTime(s string) (*DayTime, error) {
	parts := strings.Split(s, ":")
	if len(parts) < 2 || len(parts) > 3 {
		return nil, fmt.Errorf("invalid day-time format, expected HH:MM[:SS]")
	}

	var (
		hours   int
		minutes int
		seconds int
		err     error
	)

	hours, err = strconv.Atoi(strings.TrimPrefix(parts[0], "0"))
	if err != nil {
		return nil, err
	}

	minutes, err = strconv.Atoi(strings.TrimPrefix(parts[1], "0"))
	if err != nil {
		return nil, err
	}

	if len(parts) == 3 {
		seconds, err = strconv.Atoi(strings.TrimPrefix(parts[2], "0"))
		if err != nil {
			return nil, err
		}
	}

	return &DayTime{
		Hour:   int32(hours),
		Minute: int32(minutes),
		Second: int32(seconds),
	}, nil
}
