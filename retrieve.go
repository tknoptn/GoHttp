// package main
// import packages needed to make API call and decode JSON response
package main
import (
	"encoding/json"
	"fmt"
	"net/http"
)

// define Employee struct to store employee data
type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"employee_name"`
	Age      int    `json:"employee_age"`
	Salary   int    `json:"employee_salary"`
	PhotoURL string `json:"profile_image"`
}

// define Employees struct to store multiple Employee structs
type Employees struct {
	Data []Employee `json:"data"`
}

func main() {
	// make API call to retrieve employee data
	resp, err := http.Get("http://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		fmt.Println("Error occurred while retrieving data:", err)
		return
	}
	defer resp.Body.Close()

	// decode JSON response
	var employees Employees
	err = json.NewDecoder(resp.Body).Decode(&employees)
	if err != nil {
		fmt.Println("Error occurred while decoding JSON:", err)
		return
	}

	// loop through employees and print data
	for _, emp := range employees.Data {
		fmt.Printf("ID: %d\nName: %s\nAge: %d\nSalary: %d\nPhoto URL: %s\n\n", emp.ID, emp.Name, emp.Age, emp.Salary, emp.PhotoURL)
	}
}
