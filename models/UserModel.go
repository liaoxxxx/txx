package models

type User struct {
	Id        uint32
	Full_name string
	Nick_name string
	User_name string
	Phone     string
	Email     string
	Status    byte
	Password  string
	Salt      string
}
