package main

import (
	"fmt"
)

func main() {
	var companyName = "Mafer Construction"
	var totalVacancy = 50
	var remainingVacancy = 50

	fmt.Printf("Welcome to %v vacancy form.\n", companyName)
	fmt.Printf("We have a total of %v vacancies and %v are still available.\n", totalVacancy, remainingVacancy)

	for {
		var firstName, lastName string
		var email string
		var age uint
		var grade float64
		var salary = 15000

		fmt.Println("Enter your first name and last name: ")
		fmt.Scan(&firstName, &lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter your age: ")
		fmt.Scan(&age)

		fmt.Println("Enter your CGPA: ")
		fmt.Scan(&grade)

		if age >= 27 && grade >= 3.5 {
			fmt.Printf("Congradulations %v %v. you are employeed in %v with a salary of %v.\n", firstName, lastName, companyName, salary)
			remainingVacancy--
			if remainingVacancy > 0 {
				fmt.Printf("we have a remaining of %v vacancies.\n", remainingVacancy)
			} else {
				fmt.Println("We have finished our vacancies.")
			}

		} else if age < 27 && grade >= 3.5 {
			remainingAge := 27 - age
			fmt.Printf("you are %v. you can not employeed now please come back after %v years when you are 27.\n", age, remainingAge)
		} else {
			fmt.Printf("sorry %v %v you can not be employeed at %v. your grade is %v, it doesn't match with 3.5\n", firstName, lastName, companyName, grade)
		}
	}

}
