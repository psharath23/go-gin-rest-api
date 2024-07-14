package user

import "strings"

type User struct {
	Id string `json:"id"`
	AddUserBody
}

type Address struct {
	State string `json:"state"`
	City string `json:"city"`
	Street string `json:"street"`
	Hno string `json:"hno"`
	Pincode string `json:"pincode"`
}

type UserRouteParam struct {
	ID string `uri:"id" binding:"required"`
}

type AddUserBody struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Address Address `json:"address"`
}

func NewUser(addUserBody AddUserBody) *User {
	return &User {
		Id: strings.ToLower(strings.ReplaceAll(addUserBody.Name," ", "_")),
		AddUserBody: addUserBody,
	}
}

func (u *User) Update(user AddUserBody) {
	u.Address = user.Address
	u.Age = user.Age
	u.Name = user.Name
}