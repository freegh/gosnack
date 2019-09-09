package main

import(
	"fmt"
)

type User struct {
	Name string
	Email string
}

func (u *User) Notify() {
	fmt.Println("User notify na ka")
}

func (u User) SendEmail() {
	fmt.Println("user send email.")
}

type Notifier interface {
	Notify()
}

type Emailer interface {
	SendEmail()
	Notify()
}

func main() {
	u := User{
		Name : "Knot",
		Email : "Ch@g.com",
	}

	SendNotify(&u)
	SendToEmail(&u)
}

func SendNotify(n Notifier) {
	n.Notify()
}

func SendToEmail(e Emailer) {
	e.SendEmail()
	e.Notify()
}