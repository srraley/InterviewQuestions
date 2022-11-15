package main

const REG_WAGE_MULT = 1
const TIME_AND_HALF_MULT = 1.5
const DBL_WAGE_MULT = 2
const OVERTIME_LIMIT = 40
const DBLTIME_LIMIT = 48
const DECIMAL_PLACES = 4

type Data struct {
	JobMeta      `json:"jobMeta"`
	EmployeeData `json:"employeeData"`
}

type JobMeta []struct {
	Job          string  `json:"job"`
	Rate         float64 `json:"rate"`
	BenefitsRate float64 `json:"benefitsRate"`
}

type EmployeeData []Employee

type Employee struct {
	Employee      string `json:"employee"`
	TimePunchList `json:"timePunch"`
}

type TimePunchList []TimePunch

type TimePunch struct {
	Job   string `json:"job"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type Job struct {
	rate         float64
	benefitsRate float64
}

// job.Job is the title of the specific job for a key
func GetJobsMap(data Data) map[string]Job {
	jobs := make(map[string]Job)

	for _, job := range data.JobMeta {
		jobs[job.Job] = Job{rate: job.Rate, benefitsRate: job.BenefitsRate}
	}

	return jobs
}
