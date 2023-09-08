package timeUtils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	timeUtils "github.com/vasvass/Pizza-Factory/utils"
)

func TestSecondsToMilliSeconds(t *testing.T) {
	testCases := []struct {
		desc     string
		in       timeUtils.Seconds
		expected timeUtils.MilliSeconds
	}{
		{
			desc:     "1 second should be 1000 milliseconds",
			in:       1,
			expected: 1000,
		},
		{
			desc:     "2 second should be 2000 milliseconds",
			in:       2,
			expected: 2000,
		},
		{
			desc:     "3 second should be 3000 milliseconds",
			in:       3,
			expected: 3000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out := tC.in.ToMilliSeconds()

			assert.Equal(t, tC.expected, out)
		})
	}
}

func TestSecondsToNanoSeconds(t *testing.T) {
	testCases := []struct {
		desc     string
		in       timeUtils.Seconds
		expected timeUtils.NanoSeconds
	}{
		{
			desc:     "1 second should be 100000 nanoseconds",
			in:       1,
			expected: 1000000000,
		},
		{
			desc:     "2 second should be 2000000000 nanoseconds",
			in:       2,
			expected: 2000000000,
		},
		{
			desc:     "3 second should be 3000000000 nanoseconds",
			in:       3,
			expected: 3000000000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out := tC.in.ToNanoSeconds()

			assert.Equal(t, tC.expected, out)
		})
	}
}
