package processEmployeePay

import "testing"

func TestGetProcessedTimeData(t *testing.T) {
	cases := []struct {
		totalHours float64
		hours      float64
		expected   []processedTimeData
	}{
		{
			39,
			2,
			[]processedTimeData{{REG_WAGE_MULT, 2}},
		},
		{
			40,
			2,
			[]processedTimeData{{REG_WAGE_MULT, 2}},
		},
		{
			41,
			2,
			[]processedTimeData{{REG_WAGE_MULT, 1}, {TIME_AND_HALF_MULT, 1}},
		},
		{
			47,
			2,
			[]processedTimeData{{TIME_AND_HALF_MULT, 2}},
		},
		{
			48,
			2,
			[]processedTimeData{{TIME_AND_HALF_MULT, 2}},
		},
		{
			49,
			2,
			[]processedTimeData{{TIME_AND_HALF_MULT, 1}, {DBL_WAGE_MULT, 1}},
		},
	}

	for _, c := range cases {

		result := GetProcessedTimeData(c.totalHours, c.hours)
		if result[0] != c.expected[0] {
			t.Log("error should be", c.expected, ", but got", result)
			t.Fail()
		}
	}
}
