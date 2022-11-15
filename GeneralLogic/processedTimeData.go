package main

type processedTimeData struct {
	wageMultiplier float64
	hoursWorked    float64
}

func GetProcessedTimeData(totalHours, hours float64) []processedTimeData {
	if totalHours <= OVERTIME_LIMIT { //case of total hours under 40
		return []processedTimeData{{REG_WAGE_MULT, hours}}

	} else if totalHours > OVERTIME_LIMIT && totalHours <= DBLTIME_LIMIT { //case of total hours between 40 and 48
		return getProcessedDataOver40(totalHours, hours, OVERTIME_LIMIT, REG_WAGE_MULT, TIME_AND_HALF_MULT)

	} else { //case of total hours over 48
		return getProcessedDataOver40(totalHours, hours, DBLTIME_LIMIT, TIME_AND_HALF_MULT, DBL_WAGE_MULT)
	}
}

func getProcessedDataOver40(totalHours, hours, hourLimit, lowerWageMult, upperWageMult float64) []processedTimeData {
	if totalHours-hours < hourLimit { //case of partial time at reg/overtime pay and partial time at overtime/double pay
		hoursUnderLimit := hourLimit - (totalHours - hours)
		hoursOverLimit := hours - (hoursUnderLimit)
		return []processedTimeData{{lowerWageMult, hoursUnderLimit}, {upperWageMult, hoursOverLimit}}

	} else {
		return []processedTimeData{{upperWageMult, hours}}
	}
}
