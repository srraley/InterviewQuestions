package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"muzzammil.xyz/jsonc"
)

func main() {

	content, err := os.ReadFile("./data/PunchLogicTest.jsonc")
	if err != nil {
		log.Fatal(err)
	}

	data := Data{}

	err = json.Unmarshal(jsonc.ToJSON(content), &data)
	if err != nil {
		log.Fatal(err)
	}

	jobsMap := GetJobsMap(data)

	empPayPeriods := []EmployeePayForPeriod{}

	for _, emp := range data.EmployeeData {
		empPayPeriods = append(empPayPeriods, GetEmployeePayForPeriod(jobsMap, emp))
	}

	employeePayJSON, _ := json.Marshal(empPayPeriods)

	fmt.Println(string(employeePayJSON))

}
