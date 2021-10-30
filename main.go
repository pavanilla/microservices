package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/api", handler)
	fmt.Println("hai")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	q := "https://dummy.restapiexample.com/api/v1/employee/1"
	data, err := getData(q)
	if err != nil {
		fmt.Println(err)
		fmt.Println("there is a error")
	}
	res := APIResponse{
		Cache: false,
		Data:  data,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Println("there is a error connecting to server ")
	}
	fmt.Println("handler")
}

func getData(q string) ([]Fullresponse, error) {
	resp, err := http.Get(q)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	data := make([]Fullresponse, 0)
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil
}

type APIResponse struct {
	Cache bool           `json:"cache"`
	Data  []Fullresponse `json:"data"`
}

type Fullresponse struct {
	Status  string  `json:"string"`
	Details []Dummy `json:"details"`
	Message string  `json:"message"`
}
type Dummy struct {
	Id              int    `json:"id"`
	Employee_name   string `json:"employee_name"`
	Employee_salary string `json:"employee_salary"`
	Employee_age    int    `json:"employee_age"`
	Profile_image   string `json:"profile_image"`
}
