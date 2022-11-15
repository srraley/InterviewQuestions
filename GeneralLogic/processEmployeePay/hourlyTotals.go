package processEmployeePay

type HourlyTotals struct {
	totalHours float64
}

func SetTotalHours(t *HourlyTotals, hours float64) {
	t.totalHours += hours
}

func (t HourlyTotals) Regular() float64 {
	if t.totalHours > OVERTIME_LIMIT {
		return OVERTIME_LIMIT
	}
	return t.totalHours
}

func (t HourlyTotals) Overtime() float64 {
	if t.totalHours <= OVERTIME_LIMIT {
		return 0
	}
	return RoundTo(t.totalHours-(OVERTIME_LIMIT+t.DoubleTime()), DECIMAL_PLACES)
}

func (t HourlyTotals) DoubleTime() float64 {
	if t.totalHours <= DBLTIME_LIMIT {
		return 0
	}
	return RoundTo(t.totalHours-DBLTIME_LIMIT, DECIMAL_PLACES)
}
