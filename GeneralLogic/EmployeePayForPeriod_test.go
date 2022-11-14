package main

import "testing"

func TestGetHoursWorked(t *testing.T) {
	cases := []struct {
		tPunch   TimePunch
		expected float64
	}{
		{
			TimePunch{
				"Hospital - Painter",
				"2022-02-20 13:00:00",
				"2022-02-20 17:00:00",
			},
			4.0,
		},
		{
			TimePunch{
				"Hospital - Painter",
				"2022-02-20 13:00:00",
				"2022-02-21 13:00:00",
			},
			24.0,
		},
	}

	for _, c := range cases {

		result := getHoursWorked(c.tPunch)
		if result != c.expected {
			t.Log("error should be", c.expected, ", but got", result)
			t.Fail()
		}
	}
}

func TestTotalWages(t *testing.T) {
	cases := []struct {
		runningWages float64
		data         []processedTimePunchData
		rate         float64
		expected     float64
	}{
		{
			1000.00,
			[]processedTimePunchData{
				{
					1.0,
					8.0,
				},
			},
			10.00,
			1080.00,
		},
	}

	for _, c := range cases {

		result := getTotalWages(c.runningWages, c.data, c.rate)
		if result != c.expected {
			t.Log("error should be", c.expected, ", but got", result)
			t.Fail()
		}
	}
}

func TestTotalBenefits(t *testing.T) {
	cases := []struct {
		runningBenefits float64
		hours           float64
		rate            float64
		expected        float64
	}{
		{
			100.00,
			8.00,
			1.00,
			108.00,
		},
	}

	for _, c := range cases {

		result := getTotalBenefits(c.runningBenefits, c.hours, c.rate)
		if result != c.expected {
			t.Log("error should be", c.expected, ", but got", result)
			t.Fail()
		}
	}
}

func TestCalculateWages(t *testing.T) {
	cases := []struct {
		hoursWorked    float64
		rate           float64
		rateMultiplier float64
		expected       float64
	}{
		{
			8.00,
			8.00,
			1.00,
			64.00,
		},
		{
			8.00,
			8.00,
			1.5,
			96.00,
		},
		{
			8.00,
			8.00,
			2.0,
			128.00,
		},
	}

	for _, c := range cases {

		result := calculateWages(c.hoursWorked, c.rate, c.rateMultiplier)
		if result != c.expected {
			t.Log("error should be", c.expected, ", but got", result)
			t.Fail()
		}
	}
}

func TestRoundTo(t *testing.T) {
	cases := []struct {
		n        float64
		decimals uint32
		expected float64
	}{
		{
			1234.5678912345,
			2,
			1234.57,
		},
		{
			1234.5628912345,
			2,
			1234.56,
		},
	}

	for _, c := range cases {

		result := RoundTo(c.n, c.decimals)
		if result != c.expected {
			t.Log("error should be", c.expected, ", but got", result)
			t.Fail()
		}
	}
}

func TestGetProcessedTimePunchData(t *testing.T) {
	cases := []struct {
		totalHours float64
		hours      float64
		expected   []processedTimePunchData
	}{
		{
			39,
			2,
			[]processedTimePunchData{{REG_WAGE_MULT, 2}},
		},
		{
			40,
			2,
			[]processedTimePunchData{{REG_WAGE_MULT, 2}},
		},
		{
			41,
			2,
			[]processedTimePunchData{{REG_WAGE_MULT, 1}, {TIME_AND_HALF_MULT, 1}},
		},
		{
			47,
			2,
			[]processedTimePunchData{{TIME_AND_HALF_MULT, 2}},
		},
		{
			48,
			2,
			[]processedTimePunchData{{TIME_AND_HALF_MULT, 2}},
		},
		{
			49,
			2,
			[]processedTimePunchData{{TIME_AND_HALF_MULT, 1}, {DBL_WAGE_MULT, 1}},
		},
	}

	for _, c := range cases {

		result := getProcessedTimePunchData(c.totalHours, c.hours)
		if result[0] != c.expected[0] {
			t.Log("error should be", c.expected, ", but got", result)
			t.Fail()
		}
	}
}
