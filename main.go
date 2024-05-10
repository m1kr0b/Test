package main

import (
	mult "Test/newfolder"
	"encoding/json"
	"fmt"
)

type Worker struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}

func main() {
	str := `{"name":"Andrey", "age": 28, "job":"none"}`
	fmt.Println(asd.Sum(1, 1))
	jsonInfo, err := json.Marshal(str)
	if err != nil {
		fmt.Println("Ошибка записи данных:", err)
	}
	ars := mult.mult(1, 2)
	fmt.Println(jsonInfo)
	var worker Worker
	err = json.Unmarshal(jsonInfo, &worker)
	if err != nil {
		fmt.Println("Ошибка записи данных:", err)
	}
	fmt.Println(worker)

}
