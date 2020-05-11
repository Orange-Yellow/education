package controller

import "github.com/orangey.com/education/service"

type Application struct {
	Setup *service.ServiceSetup
}

type User struct {
	LoginName	string
	Password	string
	IsAdmin	string
}


var users []User

func init() {

	admin := User{LoginName:"admin", Password:"123456", IsAdmin:"T"}
	alice := User{LoginName:"bipt", Password:"123456", IsAdmin:"T"}
	bob := User{LoginName:"user", Password:"123456", IsAdmin:"F"}
	jack := User{LoginName:"bob", Password:"123456", IsAdmin:"F"}

	users = append(users, admin)
	users = append(users, alice)
	users = append(users, bob)
	users = append(users, jack)

}