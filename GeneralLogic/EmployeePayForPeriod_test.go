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
