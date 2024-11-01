package commonv1

import (
	"errors"
	"log/slog"
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

func (dt *DayTime) At(t time.Time) time.Time {
	year, month, date := t.Date()

	return time.Date(year, month, date, int(dt.Hour), int(dt.Minute), int(dt.Second), 0, t.Location())
}
