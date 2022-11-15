package main

import (
	"encoding/json"
	"fmt"
	"log"

	"myapp/processEmployeePay"
)

const JSONC_FILE = "./jsonc/PunchLogicTest.jsonc"

func main() {

	data := processEmployeePay.ParseJSONCtoData(JSONC_FILE)
	jobsMap := processEmployeePay.GetJobsMap(data)
	empPayPeriods := []processEmployeePay.EmployeePayForPeriod{}

	for _, emp := range data.EmployeeData {
		empPayPeriods = append(empPayPeriods, processEmployeePay.GetEmployeePayForPeriod(jobsMap, emp))
	}

	employeePayJSON, err := json.Marshal(empPayPeriods)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(employeePayJSON))

}
