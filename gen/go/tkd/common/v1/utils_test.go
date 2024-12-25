package commonv1

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDayTimeRange_IncludesAt(t *testing.T) {
	cases := []struct {
		Range    string
		Query    string
		Includes bool
	}{
		{
			"06:00 - 07:00",
			"06:30",
			true,
		},
		{
			"06:00 - 07:00",
			"07:00",
			false,
		},
		{
			"06:00 - 07:00",
			"06:00",
			true,
		},
		{
			"19:30 - 06:00",
			"01:00",
			true,
		},
		{
			"23:59 - 23:59",
			"06:00",
			true,
		},
		{
			"06:00 - ",
			"06:30",
			true,
		},
		{
			"06:00 - ",
			"05:30",
			false,
		},
		{
			"- 06:00",
			"05:30",
			true,
		},
		{
			"- 06:00",
			"06:30",
			false,
		},
		{
			"-",
			"06:30",
			true,
		},
		{
			"00:00 - ",
			"23:59:59",
			true,
		},
		{
			"- 00:00:00",
			"00:00:00",
			false,
		},
	}

	for _, c := range cases {
		r, err := ParseDayTimeRange(c.Range)
		require.NoError(t, err)

		q, _ := ParseDayTime(c.Query)

		assert.Equal(t, c.Includes, r.IncludesAt(time.Now(), q))
	}
}
