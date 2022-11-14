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
	if t.totalHours < OVERTIME_LIMIT {
		return 0
	}
	return RoundTo(t.totalHours-(OVERTIME_LIMIT+t.DoubleTime()), DECIMAL_PLACES)
}

func (t TotalHourlyData) DoubleTime() float64 {
	if t.totalHours < DBLTIME_LIMIT {
		return 0
	}
	return RoundTo(t.totalHours-DBLTIME_LIMIT, DECIMAL_PLACES)
}
