package main

type TotalHourlyData struct {
	totalHours float64
}

func (t TotalHourlyData) Regular() float64 {
	if t.totalHours > OVERTIME_LIMIT {
		return OVERTIME_LIMIT
	}
	return t.totalHours
}

func (t TotalHourlyData) Overtime() float64 {
	if t.totalHours <= OVERTIME_LIMIT {
		return 0
	}
	return RoundTo(t.totalHours-(OVERTIME_LIMIT+t.DoubleTime()), DECIMAL_PLACES)
}

func (t TotalHourlyData) DoubleTime() float64 {
	if t.totalHours <= DBLTIME_LIMIT {
		return 0
	}
	return RoundTo(t.totalHours-DBLTIME_LIMIT, DECIMAL_PLACES)
}

func GetProcessedTimeData(totalHours, hours float64) []processedTimePunchData {
	if totalHours <= OVERTIME_LIMIT { //case of total hours under 40
		return []processedTimePunchData{{REG_WAGE_MULT, hours}}

	} else if totalHours > OVERTIME_LIMIT && totalHours <= DBLTIME_LIMIT { //case of total hours between 40 and 48
		return getProcessedDataOver40(totalHours, hours, OVERTIME_LIMIT, REG_WAGE_MULT, TIME_AND_HALF_MULT)

	} else { //case of total hours over 48
		return getProcessedDataOver40(totalHours, hours, DBLTIME_LIMIT, TIME_AND_HALF_MULT, DBL_WAGE_MULT)
	}
}

func getProcessedDataOver40(totalHours, hours, hourLimit, lowerWageMult, upperWageMult float64) []processedTimePunchData {
	if totalHours-hours < hourLimit { //case of partial time at reg/overtime pay and partial time at overtime/double pay
		hoursUnderLimit := hourLimit - (totalHours - hours)
		hoursOverLimit := hours - (hoursUnderLimit)
		return []processedTimePunchData{{lowerWageMult, hoursUnderLimit}, {upperWageMult, hoursOverLimit}}

	} else {
		return []processedTimePunchData{{upperWageMult, hours}}
	}
}

func SetTotalHours(t *TotalHourlyData, hours float64) {
	t.totalHours += hours
}
