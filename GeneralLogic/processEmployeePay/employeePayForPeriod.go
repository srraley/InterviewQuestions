package processEmployeePay

import (
	"log"
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
	var totalWages, totalBenefits float64
	var totalHourlyData TotalHourlyData

	for _, tPunch := range employee.TimePunchList {
		hours := getHoursWorked(tPunch)
		SetTotalHours(&totalHourlyData, hours)
		rate := jobsMap[tPunch.Job].rate
		benefitsRate := jobsMap[tPunch.Job].benefitsRate
		totalWages = getTotalWages(totalWages, GetProcessedTimeData(totalHourlyData.totalHours, hours), rate)
		totalBenefits = getTotalBenefits(totalBenefits, hours, benefitsRate)
	}

	return EmployeePayForPeriod{
		employee.Employee,
		totalHourlyData.Regular(),
		totalHourlyData.Overtime(),
		totalHourlyData.DoubleTime(),
		totalWages,
		totalBenefits,
	}
}

func getHoursWorked(tPunch TimePunch) float64 {
	layout := "2006-01-02 15:04:05"
	start, err := time.Parse(layout, tPunch.Start)
	if err != nil {
		log.Fatal(err)
	}
	end, err := time.Parse(layout, tPunch.End)
	if err != nil {
		log.Fatal(err)
	}
	return RoundTo(end.Sub(start).Hours(), DECIMAL_PLACES)
}

func getTotalWages(runningWages float64, data []processedTimeData, rate float64) float64 {
	for _, d := range data {
		runningWages += calculateWages(d.hoursWorked, rate, d.wageMultiplier)
	}
	return runningWages
}

func getTotalBenefits(runningBenefits, hours, benefitsRate float64) float64 {
	return RoundTo(runningBenefits+calculateWages(hours, benefitsRate, REG_WAGE_MULT), DECIMAL_PLACES)
}

func calculateWages(hoursWorked, rate, rateMultiplier float64) float64 {
	return RoundTo(hoursWorked*(rate*rateMultiplier), DECIMAL_PLACES)
}

func RoundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}
