package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"myapp/processEmployeePay"

	"muzzammil.xyz/jsonc"
)

func main() {

	content, err := os.ReadFile("./jsonc/PunchLogicTest.jsonc")
	if err != nil {
		log.Fatal(err)
	}

	data := processEmployeePay.Data{}

	err = json.Unmarshal(jsonc.ToJSON(content), &data)
	if err != nil {
		log.Fatal(err)
	}

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
