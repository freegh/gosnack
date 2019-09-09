package main

import (
	"fmt"
)

type User struct {
	Name string
	Email string
}

type Notifier interface {
	Notify(s string)
	// SendEmail(cc string) (result string)
	Emailer //embedded
}

type Emailer interface {
	SendEmail(cc string) (result string)
}

func SendNotify(u Notifier, s string) {
	u.Notify(s)
	e := u.SendEmail("KnLo")
	fmt.Println("result", e)
}

func SendtoEmail(e Emailer) {
	fmt.Println(e.SendEmail("P Senit"))
}

func (u User) SendEmail(s string) string {
	return s
}

func (a Admin) SendEmail(s string) string {
	return s
}

func (u User) Notify(s string) {
	fmt.Println("Nong to ", u.Email, s)
}

//override embedded
func (a Admin) Notify(n string ) {
	fmt.Println("secret ", a.Email, n)
}

type Admin struct{
	//embedded struct
	User
	Level string
}

func main() {
	u := User{
		Name: "Knot",
		Email: "Chayapol.ph@gmail.com",
	}

	u.Notify("aaa")
	User.Notify(u,"io")

	// a := Admin{
	// 	User: User{
	// 		Name: "Pang",
	// 		Email: "Pijak@gmail.com",
	// 	},
	// 	Level: "GOD user",
	// }

	a := Admin{}
	a.Name = "Nong"
	a.Email = "admin@g.com"
	a.Level = "supper user"

	a.Notify("123")
	a.User.Notify("BB")

	SendNotify(u, "1")
	SendtoEmail(u)

	// a.User.Name = "ss"
	// a.User.Email ="admi@g.com"
	// a.Level = "High user"

	fmt.Println("user", u)
	fmt.Println("admin", a)
	fmt.Println("admin Name", a.Name)
	fmt.Println("admin Email", a.Email)
}