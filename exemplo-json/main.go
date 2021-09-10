package main

import (
	"exemplo-json/model"
	"fmt"
	"time"
)

func main() {
	u := model.User{
		Id:        1,
		Name:      "Matheus",
		Email:     "matheus-pg1@hotmail.com",
		CreatedAt: time.Now(),
	}

	userJSON := u.ToJSON()
	fmt.Println(userJSON)

	userInJSON := `{"id":1,"name":"Matheus","email":"matheus-pg1@hotmail.com","createdAt":"2021-08-31T08:08:55.0917678-03:00"}`
	var u2 model.User
	u2.ToEntity(userInJSON)
	fmt.Println(u2)
}
