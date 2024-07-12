package main

import (
	"fmt"
	"math/rand"
	"time"
	"weeklyTest2/models"
)

func main() {
	employees := GenerateEmployees(100)
	employees2 := GenerateEmployeesWithChannel(100)

	ShowAllEmployee(employees)
	ShowAllEmployee(employees2)
}

func SalaryRandomizer()float64{
		rand.Seed(time.Now().UnixNano())
		randomFloat := rand.Float64()
		randomSalary := 5000.0 + randomFloat*(15000.0-5000.0)

		return randomSalary
	
}

func ShowAllEmployee(employee []models.ChildInterface){
	for _, v := range employee {
		fmt.Println("-----------------------------------------")
		fmt.Printf("fullName = %s\n", v.GetfullName())
		fmt.Printf("salary = %.2f\n", v.GetSalary())
		fmt.Printf("status = %s\n", v.GetStatus())
		fmt.Printf("insurance = %.2f\n", v.GetInsurance())
		fmt.Printf("overtime = %.2f\n", v.GetOvertime())
		fmt.Printf("allowance = %.2f\n", v.GetAllowance())
		fmt.Printf("totalSalary = %.2f\n", v.GetTotalSalary())
		fmt.Println("-----------------------------------------")
	}
	
}

func GenerateEmployeesWithChannel(empCount int)[]models.ChildInterface{
	rand.Seed(time.Now().UnixNano())
	employees:= []models.ChildInterface{}
	names := []string {"Agathi Aiko","Argyro Yianna","Antonia Koralia","Marilena Ran","Fumie Klio", "Marika Keti","Kei Kasih","Shigeko Rin","Nomiki Akira","Sotos Kazumi","Kouji Kōki"}
	statuses := []models.Status{models.Permanents, models.Contracts, models.Trainees}
	employeeChannel := make(chan models.ChildInterface, empCount)
	nameRandomizer := 0

	for i:=0; i < empCount; i++ {

		nameRandomizer = rand.Intn(len(names))

		go func(name string, salary float64, status models.Status){

			//Conditional status
			if status == models.Permanents {
				employeeChannel <- models.NewPermanent(name, salary)
			
			}else if status == models.Contracts{
				employeeChannel <- models.NewContract(name, salary)
			
			}else {
				employeeChannel <- models.NewTrainee(name, salary)
			}

		}(names[nameRandomizer], SalaryRandomizer(), statuses[rand.Intn(len(statuses))]) //argument = (random name dari slice names, salary Randomizer, random status dari statuses)

		//Masukkan data pada employeeChannel ke slice
		employees = append(employees, <- employeeChannel)
	}

	close(employeeChannel)

	return employees
}

func GenerateEmployees(empCount int)[]models.ChildInterface{
	rand.Seed(time.Now().UnixNano())
	employees:= []models.ChildInterface{}
	names := []string {"Agathi Aiko","Argyro Yianna","Antonia Koralia","Marilena Ran","Fumie Klio", "Marika Keti","Kei Kasih","Shigeko Rin","Nomiki Akira","Sotos Kazumi","Kouji Kōki"}
	statuses := []models.Status{models.Permanents, models.Contracts, models.Trainees}
	nameRandomizer := 0

	for i:=0; i < empCount; i++ {
		nameRandomizer = rand.Intn(len(names))

		func(name string, salary float64, status models.Status){

			//Conditional status
			if status == models.Permanents {
				employees  =  append(employees,models.NewPermanent(name, salary))
			
			}else if status == models.Contracts{
				employees =  append(employees,models.NewContract(name, salary))
			
			}else {
				employees =  append(employees,models.NewTrainee(name, salary))
			}

		}(names[nameRandomizer], SalaryRandomizer(), statuses[rand.Intn(len(statuses))]) //argument = (random name dari slice names, salary Randomizer, random status dari statuses)

	}

	return employees
}