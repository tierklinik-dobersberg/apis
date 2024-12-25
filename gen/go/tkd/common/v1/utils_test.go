package commonv1

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDayTimeRange_IncludesAt(t *testing.T) {
	cases := []struct {
		Range    []string
		Query    string
		Includes bool
	}{
		{
			[]string{"06:00", "07:00"},
			"06:30",
			true,
		},
		{
			[]string{"06:00", "07:00"},
			"07:00",
			false,
		},
		{
			[]string{"06:00", "07:00"},
			"06:00",
			true,
		},
		{
			[]string{"19:30", "06:00"},
			"01:00",
			true,
		},
		{
			[]string{"23:59", "23:59"},
			"06:00",
			true,
		},
		{
			[]string{"06:00", ""},
			"06:30",
			true,
		},
		{
			[]string{"06:00", ""},
			"05:30",
			false,
		},
		{
			[]string{"", "06:00"},
			"05:30",
			true,
		},
		{
			[]string{"", "06:00"},
			"06:30",
			false,
		},
		{
			[]string{"", ""},
			"06:30",
			true,
		},
		{
			[]string{"00:00", ""},
			"23:59:59",
			true,
		},
		{
			[]string{"", "00:00:00"},
			"00:00:00",
			false,
		},
	}

	for _, c := range cases {
		var (
			r1 *DayTime
			r2 *DayTime
		)

		if len(c.Range[0]) > 0 {
			r1, _ = ParseDayTime(c.Range[0])
		}
		if len(c.Range[1]) > 0 {
			r2, _ = ParseDayTime(c.Range[1])
		}

		q, _ := ParseDayTime(c.Query)

		r := &DayTimeRange{
			Start: r1,
			End:   r2,
		}

		assert.Equal(t, c.Includes, r.IncludesAt(time.Now(), q))
	}
}
