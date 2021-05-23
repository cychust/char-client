package models

type User struct {
	UserId string
	Name   string
	Age    int
	Sex    int
}

func (u User) Info() string {
	info := u.Name + " " + string(u.Age) + "years old "
	return info
}
