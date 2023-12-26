package data

import (
	"encoding/json"
	"os"
)

type AddressStruct struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}

type ContactStruct struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Employee struct {
	EmployeeID int           `json:"employeeId"`
	FirstName  string        `json:"firstName"`
	LastName   string        `json:"lastName"`
	Age        int           `json:"age"`
	Position   string        `json:"position"`
	Department string        `json:"department"`
	Salary     float64       `json:"salary"`
	Address    AddressStruct `json:"address"`
	Contact    ContactStruct `json:"contact"`
	IsActive   bool          `json:"isActive"`
	Avatar     string        `json:"avatar"`
}

func GetFinalPath() (string, error) {
	path, err := os.Getwd()
	if err == nil {
		return path + "/public/data.json", nil
	}
	return "", err
}

var finalPath, err = GetFinalPath()

func GetAllData() (data []Employee) {
	if err == nil {
		byteData, _ := os.ReadFile(finalPath)
		json.Unmarshal(byteData, &data)
	}
	return
}

func AddData(employee Employee) {
	data := GetAllData()
	data = append(data, employee)

	fileData, _ := json.MarshalIndent(data, "", " ")
	os.WriteFile(finalPath, fileData, 0644)
}
