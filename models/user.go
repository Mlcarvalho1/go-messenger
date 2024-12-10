package models

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}
