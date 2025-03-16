package user

var global int = 42

type User struct {
	name string
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

type Address struct {
	city string
}

func (a *Address) City() string {
	return a.city
}

func (a *Address) SetCity(city string) {
	a.city = city
}
