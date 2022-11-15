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
	empPaycheques := processEmployeePay.GetPayCheques(data, jobsMap)

	employeePayJSON, err := json.Marshal(empPaycheques)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(employeePayJSON))

}
