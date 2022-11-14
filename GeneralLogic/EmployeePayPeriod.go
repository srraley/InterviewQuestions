package main

import (
	"math"
	"time"
)

type EmployeePayForPeriod struct {
	Employee     string
	Regular      float64
	Overtime     float64
	Doubletime   float64
	WageTotal    float64
	BenefitTotal float64
}

func GetEmployeePayForPeriod(jobsMap map[string]Job, employee Employee) EmployeePayForPeriod {
	var totalWages float64
	var totalBenefits float64
	var totalHourlyData TotalHourlyData

	for _, tPunch := range employee.TimePunchList {
		hours := getHoursWorked(tPunch)
		totalHourlyData.totalHours += hours
		rate := jobsMap[tPunch.Job].rate
		benefitsRate := jobsMap[tPunch.Job].benefitsRate
		totalWages = getTotalWages(totalWages, getProcessedTimePunchData(totalHourlyData.totalHours, hours), rate)
		totalBenefits = getTotalBenefits(totalBenefits, hours, benefitsRate)
	}

	return EmployeePayForPeriod{
		employee.Employee,
		totalHourlyData.Total(),
		totalHourlyData.Overtime(),
		totalHourlyData.DoubleTime(),
		totalWages,
		totalBenefits,
	}
}

func getHoursWorked(tPunch TimePunch) float64 {
	layout := "2006-01-02 15:04:05"
	start, _ := time.Parse(layout, tPunch.Start)
	end, _ := time.Parse(layout, tPunch.End)
	return RoundTo(end.Sub(start).Hours(), 4)
}

func getTotalWages(runningWages float64, data []processedTimePunchData, rate float64) float64 {
	for _, d := range data {
		runningWages += calculateWages(d.hoursWorked, rate, d.wageMultiplier)
	}
	return runningWages
}

func getTotalBenefits(runningBenefits float64, hours float64, benefitsRate float64) float64 {
	return RoundTo(runningBenefits+calculateWages(hours, benefitsRate, REG_WAGE_MULT), 4)
}

func calculateWages(hoursWorked float64, rate float64, rateMultiplier float64) float64 {
	return RoundTo(hoursWorked*(rate*rateMultiplier), 4)

}

func RoundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func getProcessedTimePunchData(totalHours float64, hours float64) []processedTimePunchData {

	if totalHours < OVERTIME_LIMIT {
		return []processedTimePunchData{{REG_WAGE_MULT, hours}}

	} else if totalHours > OVERTIME_LIMIT && totalHours < DBLTIME_LIMIT {
		if totalHours-hours < OVERTIME_LIMIT {
			hoursUnder40 := OVERTIME_LIMIT - (totalHours - hours)
			hoursOver40 := hours - hoursUnder40
			return []processedTimePunchData{{REG_WAGE_MULT, hoursUnder40}, {TIME_AND_HALF_MULT, hoursOver40}}

		} else {
			return []processedTimePunchData{{TIME_AND_HALF_MULT, hours}}

		}
	} else {
		if totalHours-hours < DBLTIME_LIMIT {
			hoursUnder48 := DBLTIME_LIMIT - (totalHours - hours)
			hoursOver48 := hours - (hoursUnder48)
			return []processedTimePunchData{{TIME_AND_HALF_MULT, hoursUnder48}, {DBL_WAGE_MULT, hoursOver48}}

		} else {
			return []processedTimePunchData{{DBL_WAGE_MULT, hours}}

		}

	}
}
