package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Customer struct {
	Name string `json:"name"`
}

type UniversalDTO struct {
	Data interface{} `json:"data"`
	// more fields with important meta-data about the message...
}

func main() {
	// create a customer, add it to DTO object and marshal it
	customer := Customer{Name: "Ben"}
	dtoToSend := UniversalDTO{customer}
	byteData, _ := json.Marshal(dtoToSend)

	// unmarshal it (usually after receiving bytes from somewhere)
	receivedDTO := UniversalDTO{}
	json.Unmarshal(byteData, &receivedDTO)

	//Attempt to unmarshall our customer
	receivedCustomer := getCustomerFromDTO(receivedDTO.Data)
	fmt.Println(receivedCustomer)


	var res time.Time = time.Now()
	fmt.Println(res.Format("2006-01-02 15:04:05"),res.String())
	datetime := time.Unix(res.Unix(), 0).Format("2006-01-02 15:04:05")
	fmt.Println(datetime)




}

func getCustomerFromDTO(data interface{}) Customer {
	m := data.(map[string]interface{})
	customer := Customer{}
	if name, ok := m["name"].(string); ok {
		customer.Name = name
	}
	return customer
}

