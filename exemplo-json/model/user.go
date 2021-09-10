package model

import (
	"bytes"
	"encoding/json"
	"log"
	"time"
)

// para ignorar um campo do JSON usa-se `json:"-"`
type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u *User) ToJSON() string {
	user, err := json.Marshal(u)

	if err != nil {
		log.Fatal(err)
	}

	return bytes.NewBuffer(user).String()
}

func (u *User) ToEntity(JSON string) {
	err := json.Unmarshal([]byte(JSON), u)
	if err != nil {
		log.Fatal(err)
	}
}
