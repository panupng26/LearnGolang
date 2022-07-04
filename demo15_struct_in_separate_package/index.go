package main

import (
	"demo15/employee"
)

func main() {
	e := employee.Employee{
		FirstName: "Bank",
		LastName: "Panupong",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}

	e.LeavesRemaining()
}